package update

import (
	"io"
	"net/http"
	"net/url"

	"github.com/alexandresoro/netatmo/netatmoApi"
)

func RefreshToken(tokenData TokenData, clientId string, clientSecret string) ([]byte, error) {

	response, err := http.PostForm(netatmoApi.NetatmoApiUrl+netatmoApi.AuthenticateEndpoint, url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {tokenData.RefreshToken},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
	})

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
