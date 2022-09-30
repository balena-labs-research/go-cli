package cobra

import (
	"github.com/balena-community/go-cli/cmd/go-cli/devices"
	"github.com/spf13/cobra"
)

// sshCmd represents the ssh command
var (
	port   string
	sshCmd = &cobra.Command{
		Use:   "ssh [optional: device address]",
		Short: "Connect to a balenaOS development device via SSH",
		Long: `Connect to a balenaOS development device via SSH.
		
Scan for devices and prompt which to connect to
  $ balena ssh 
Connect via SSH to the specified device
  $ balena ssh c938a7a.local
`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			devices.Ssh(args, "", port)
		},
	}
)

func init() {
	rootCmd.AddCommand(sshCmd)

	sshCmd.Flags().StringVarP(&port, "port", "p", "", "Specify ssh port. Default is 22222")
}
