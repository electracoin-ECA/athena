package cmd

import (
	"fmt"
	"os"

	"github.com/ivangabriele/athena/src/simulator"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simulator",
	Short: "Athena network terminal simulator.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		simulator.Start()
	},
}

// Execute do that.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
