//go:build linux
// +build linux

package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "docker [Command]",
	}
	rootCmd.AddCommand(InitRunCmd())
	rootCmd.AddCommand()
	rootCmd.AddCommand()

	rootCmd.Execute()

}
