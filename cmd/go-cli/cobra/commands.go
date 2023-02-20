package cobra

import (
	"github.com/balena-labs-research/go-cli/cmd/go-cli/devices"
	"github.com/mutagen-io/mutagen/cmd"
	"github.com/mutagen-io/mutagen/cmd/mutagen/daemon"
	"github.com/mutagen-io/mutagen/cmd/mutagen/sync"
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
	Short:  "Find devices on the network by reverse lookup. Requires a query to all devices in your network IP range",
	Long:   "Uses a reverse lookup on all IP ranges in the range `your-systems-subnet/24`. Specify your own range by passing a range as an argument: 192.168.1.0/24",
	PreRun: toggleDebug,
	Args:   cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		devices.Scan("lookup", args)
	},
}

var mount = &cobra.Command{
	Use:          "mount",
	Short:        "For mounting balena containers to your local system",
	SilenceUsage: true,
}

var mutagenCreate = &cobra.Command{
	Use:          "create <container_name> <container_path> <local_path> <host_address>",
	Short:        "Create a new synchronised mount session with a balena container",
	PreRun:       toggleDebug,
	RunE:         sync.CreateMain,
	SilenceUsage: true,
}

var mutagenList = &cobra.Command{
	Use:          "list [<session>...]",
	Short:        "List existing synchronization sessions and their statuses",
	PreRun:       toggleDebug,
	RunE:         sync.ListMain,
	SilenceUsage: true,
}

// runCommand is the run command.
var mutagenRun = &cobra.Command{
	Use:          "run",
	Short:        "Run the Mutagen daemon",
	Args:         cmd.DisallowArguments,
	Hidden:       true,
	RunE:         daemon.RunMain,
	SilenceUsage: true,
}

var mutagenTerminate = &cobra.Command{
	Use:          "terminate [<session>...]",
	Short:        "Permanently terminate a synchronization session",
	PreRun:       toggleDebug,
	RunE:         sync.TerminateMain,
	SilenceUsage: true,
}

// stopCommand is the stop command.
var mutagenStop = &cobra.Command{
	Use:          "stop",
	Short:        "Stop the Mutagen daemon if it's running",
	Args:         cmd.DisallowArguments,
	RunE:         daemon.StopMain,
	SilenceUsage: true,
}

// startCommand is the start command.
var mutagenStart = &cobra.Command{
	Use:          "start",
	Short:        "Start the Mutagen daemon if it's not already running",
	Args:         cmd.DisallowArguments,
	RunE:         daemon.StartMain,
	SilenceUsage: true,
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
	rootCmd.AddCommand(mount)
	mount.AddCommand(mutagenCreate)
	mount.AddCommand(mutagenList)
	mount.AddCommand(mutagenRun)
	mount.AddCommand(mutagenTerminate)
	mount.AddCommand(mutagenStop)
	mount.AddCommand(mutagenStart)
}
