package cmd

import (
	"encoding/json"
	"fmt"
	"netatmo/update"
	"netatmo/utils"

	"github.com/spf13/cobra"
)

var (
	tokenPath  string
	outputFile string
)

var updateCmd = &cobra.Command{
	Use:   "update deviceMacAddress clientId clientSecret",
	Short: "Update the weather data",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		// Extract the existing token information
		tokenInfo, err := update.ExtractTokenInfoFromFile(tokenPath)
		if err != nil {
			fmt.Printf("The provided file at location %s that contains the token can not be found or read\n", tokenPath)
			panic(err)
		}

		// Refresh the token
		refreshedTokenInfo, err := update.RefreshToken(tokenInfo, args[1], args[2])
		if err != nil {
			panic(err)
		}

		// Save the updated token to the token file
		err = utils.WriteToDestinationPath(tokenPath, refreshedTokenInfo)
		if err != nil {
			panic(err)
		}

		// Now we can call the API for weather data
		updatedTokenInfo := update.TokenData{}
		err = json.Unmarshal(refreshedTokenInfo, &updatedTokenInfo)
		if err != nil {
			panic(err)
		}
		weatherData, err := update.GetWeatherData(updatedTokenInfo.AccessToken, args[0])
		if err != nil {
			panic(err)
		}

		// Write the weather data to an output file
		err = update.WriteWeatherDataToFile(weatherData, outputFile)
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVar(&tokenPath, "tokenPath", defaultTokenPath, "Path from where the existing token should be read and updated")
	updateCmd.Flags().StringVar(&outputFile, "outputFile", defaultOutputFile, "Path where the output result should be written")

}
