package cobra

import (
	"github.com/balena-labs-research/go-cli/cmd/go-cli/devices"
	"github.com/spf13/cobra"
)

var (
	port string
)

// arpScanCmd represents the arp scan command
var arpScanCmd = &cobra.Command{
	Use:    "arpscan",
	Short:  "Perform Arp scan for balenaOS development devices on your local network",
	PreRun: toggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		devices.Scan("arp", nil)
	},
}

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

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh [optional: device address]",
	Short: "Connect to a balenaOS development device via SSH",
	Long: `Connect to a balenaOS development device via SSH.
		
Scan for devices and prompt which to connect to
  $ balena ssh
Connect via SSH to the specified device
  $ balena ssh c938a7a.local
`,
	PreRun: toggleDebug,
	Args:   cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		devices.Ssh(args, "root", port)
	},
}

func init() {
	rootCmd.AddCommand(arpScanCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(lookupAddressCmd)
	rootCmd.AddCommand(sshCmd)
	sshCmd.Flags().StringVarP(&port, "port", "p", "22222", "Specify ssh port. Default is 22222")
}
