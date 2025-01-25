package main

import (
	"fmt"
	"os"
	"pavlovic265/265-gt/commands"
	"pavlovic265/265-gt/commands/branch"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gt",
	Short: "",
}

func main() {
	rootCmd.AddCommand(commands.Status())

	rootCmd.AddCommand(branch.Create())
	rootCmd.AddCommand(branch.Checkout())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
