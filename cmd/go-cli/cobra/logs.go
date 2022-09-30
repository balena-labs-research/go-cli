package cobra

import (
	"github.com/balena-community/go-cli/cmd/go-cli/devices"
	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:    "logs [optional: device address]",
	Short:  "Stream logs of running containers",
	PreRun: toggleDebug,
	Args:   cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		devices.StreamAllLogs(args)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
