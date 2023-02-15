package cobra

import (
	"github.com/balena-labs-research/go-cli/cmd/go-cli/devices"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var lookupAddressCmd = &cobra.Command{
	Use:    "lookup",
	Short:  "Find devices on the network by reverse lookup. Requires a query to all devices in your network IP range.",
	Long:   "Uses a reverse lookup on all IP ranges in the range `your-systems-subnet/24`. Specify your own range by passing a range as an argument: 192.168.1.0/24.",
	PreRun: toggleDebug,
	Args:   cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		devices.Scan("lookup", args)
	},
}

func init() {
	rootCmd.AddCommand(lookupAddressCmd)
}
