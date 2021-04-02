package signin

import (
	"io"
	"net/http"
	"net/url"

	"github.com/alexandresoro/netatmo/netatmoApi"
)

func RetrieveTokenFromCode(code string, redirectUri string, clientId string, clientSecret string) ([]byte, error) {

	response, err := http.PostForm(netatmoApi.NetatmoApiUrl+netatmoApi.AuthenticateEndpoint, url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"code":          {code},
		"redirect_uri":  {redirectUri},
		"scope":         {netatmoApi.Scope},
	})

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)

}
