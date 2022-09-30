package cobra

import (
	"github.com/balena-community/go-cli/cmd/go-cli/api"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:    "api",
	Short:  "Start the API and listen for requests",
	PreRun: toggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		api.Start()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
