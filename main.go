package main

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/commands"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gt",
	Short: "",
}

func main() {
	rootCmd.AddCommand(commands.Status())
	rootCmd.AddCommand(commands.Create())
	rootCmd.AddCommand(commands.Checkout())
	rootCmd.AddCommand(commands.Swich())
	rootCmd.AddCommand(commands.Move())
	rootCmd.AddCommand(commands.Add())
	rootCmd.AddCommand(commands.Cont())
	rootCmd.AddCommand(commands.Push())
	rootCmd.AddCommand(commands.Pull())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
