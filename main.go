package main

import (
	"fmt"
	"os"
	"pavlovic265/265-gt/commands"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gt",
	Short: "",
}

func main() {
	rootCmd.AddCommand(commands.Status())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
