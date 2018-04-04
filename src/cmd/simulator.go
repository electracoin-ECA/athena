package cmd

import (
	"github.com/ivangabriele/athena/src/simulator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(simulatorCommand)
}

var simulatorCommand = &cobra.Command{
	Use:   "simulator",
	Short: "Athena network terminal simulator.",
	Run: func(cmd *cobra.Command, args []string) {
		simulator.Start()
	},
}
