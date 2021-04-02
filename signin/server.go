package signin

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/alexandresoro/netatmo/netatmoApi"
)

const (
	serverPort = 3000
	loginPath  = "/login"
	returnPath = "/callback"
)

type SignInParameters struct {
	ClientId     string
	ClientSecret string
	OutputPath   string
}

type AuthorizationCallbackHandler struct {
	signInParameters SignInParameters
	codeCh           chan string
}

func (authHandler AuthorizationCallbackHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	code := req.URL.Query().Get("code")
	fmt.Fprintf(rw, "Code retrieved, you can now close this tab!")

	authHandler.codeCh <- code
}

func RequestNewCode(signInParameters SignInParameters, hostname string) (string, string) {

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: mux,
	}

	redirectUri := fmt.Sprintf("http://%s:%d%s", hostname, serverPort, returnPath)

	authorizeParams := url.Values{}
	authorizeParams.Set("client_id", signInParameters.ClientId)
	authorizeParams.Set("redirect_uri", redirectUri)
	authorizeParams.Set("scope", netatmoApi.Scope)
	authorizeParams.Set("state", "state")

	loginHandler := http.RedirectHandler(netatmoApi.NetatmoApiUrl+netatmoApi.AuthorizeEndpoint+"?"+authorizeParams.Encode(), http.StatusFound)
	mux.Handle(loginPath, loginHandler)

	codeCh := make(chan string)

	authorizationHandler := AuthorizationCallbackHandler{
		signInParameters: signInParameters,
		codeCh:           codeCh,
	}
	mux.Handle(returnPath, authorizationHandler)

	go func() {
		fmt.Printf("Please open http://%s:%d/login in a browser to continue\n", hostname, serverPort)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	receivedCode := <-codeCh

	srv.Shutdown(context.Background())

	return receivedCode, redirectUri
}
