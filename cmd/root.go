package cmd

import (
	"github.com/spf13/cobra"
)

const (
	defaultTokenPath  = "~/.config/netatmo/token"
	defaultOutputFile = "~/.config/netatmo/data"
)

var rootCmd = &cobra.Command{
	Use:   "netatmo",
	Short: "A Go application to update the weather data of your Netatmo station",
	Long: `A Go application to update the weather data of your Netatmo station.
To use it, you will need to create a client ID/secret pair from dev.netatmo.com,
then retrieve an initial token with the signin command.
After that, you will be able to reuse this token to update your weather data automatically`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
