package cobra

import (
	"github.com/balena-community/go-cli/cmd/go-cli/devices"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:    "scan",
	Short:  "Scan for balenaOS development devices on your local network",
	PreRun: toggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		devices.Scan()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
