package cobra

import (
	"github.com/balena-labs-research/go-cli/cmd/go-cli/devices"
	"github.com/spf13/cobra"
)

// arpScanCmd represents the arp scan command
var arpScanCmd = &cobra.Command{
	Use:    "arpscan",
	Short:  "Perform Arp scan for balenaOS development devices on your local network",
	PreRun: toggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		devices.ArpScan()
	},
}

func init() {
	rootCmd.AddCommand(arpScanCmd)
}
