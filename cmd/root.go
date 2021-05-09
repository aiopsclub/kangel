package cmd

import (
	"github.com/spf13/cobra"
	"os"

	"fmt"
)

var rootCmd = &cobra.Command{
	Use:   "kangel",
	Short: "kangel is angel daemon for kubelet.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
