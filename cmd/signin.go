package cmd

import (
	"github.com/alexandresoro/netatmo/signin"
	"github.com/alexandresoro/netatmo/utils"

	"github.com/spf13/cobra"
)

const (
	defaultHostname = "localhost"
)

var hostname string
var outputTokenPath string

var signinCmd = &cobra.Command{
	Use:   "signin clientId clientSecret",
	Short: "Sign-in with an existing Netatmo account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		signInParams := signin.SignInParameters{
			ClientId:     args[0],
			ClientSecret: args[1],
			OutputPath:   outputTokenPath,
		}

		code, redirectUri := signin.RequestNewCode(signInParams, hostname)

		tokenBytes, err := signin.RetrieveTokenFromCode(code, redirectUri, args[0], args[1])
		if err != nil {
			panic(err)
		}

		err = utils.WriteToDestinationPath(signInParams.OutputPath, tokenBytes)
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(signinCmd)

	signinCmd.Flags().StringVar(&outputTokenPath, "tokenPath", defaultTokenPath, "Path where the token should be written")
	signinCmd.Flags().StringVar(&hostname, "hostname", defaultHostname, "Address where this server is accessible")

}
