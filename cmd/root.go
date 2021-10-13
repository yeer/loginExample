package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "loginExample",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to use loginExample, use -h or --help to see more detail")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
