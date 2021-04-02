package update

import (
	"encoding/json"
	"netatmo/utils"
	"os"
)

type TokenData struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func ExtractTokenInfoFromFile(tokenPath string) (TokenData, error) {

	path := utils.ResolvePathWithTilde(tokenPath)

	token := TokenData{}

	tokenData, err := os.ReadFile(path)
	if err != nil {
		return TokenData{}, err
	}

	err = json.Unmarshal([]byte(tokenData), &token)
	if err != nil {
		return TokenData{}, err
	}

	return token, nil
}
