package cobra

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	debug   bool
	rootCmd = &cobra.Command{
		Use:    "balena-go",
		Short:  "balena-go is a CLI for balenaOS development",
		Long:   "This is a community maintained tool providing a series of development related functions for balena devices",
		PreRun: toggleDebug,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your global flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug logging")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func toggleDebug(cmd *cobra.Command, args []string) {
	if debug {
		log.Info("Debug logs enabled")
		log.SetLevel(log.DebugLevel)
	}
}
