package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dummy",
	Short: "A simple dummy app",
}

// version variable is here to be modified at compile time.
var version = "(development)"

func init() {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("{{ .Version }}\n")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Execution failed", err)
		os.Exit(1)
	}
}
