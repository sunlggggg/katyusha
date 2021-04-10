package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
  Use:   "katyusha",
  Short: "Katyusha is a Network Disk.",
  Long: `This is Katyusha`,
  Run: func(cmd *cobra.Command, args []string) {
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

