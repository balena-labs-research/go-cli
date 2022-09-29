package cobra

import (
	"github.com/balena-community/go-cli/cmd/go-cli/devices"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var (
	port   string
	sshCmd = &cobra.Command{
		Use:   "ssh [device address]",
		Short: "Connect to a balenaOS development device via SSH",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			devices.Ssh(args[0], "", port)
		},
	}
)

func init() {
	rootCmd.AddCommand(sshCmd)

	sshCmd.Flags().StringVarP(&port, "port", "p", "", "Specify ssh port. Default is 22222")
}
