package cmd

import (
	"github.com/ivangabriele/athena/src/daemon"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(daemonCommand)
}

var daemonCommand = &cobra.Command{
	Use:   "daemon",
	Short: "CLI Daemon.",
	Run: func(cmd *cobra.Command, args []string) {
		daemon.Start()
	},
}
