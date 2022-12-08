package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage user authentication",
}

// authServeCmd represents the `auth serve` command
var authServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the authentication server",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("auth serve called")
	},
}

func init() {
	authCmd.AddCommand(
		authServeCmd,
	)

	rootCmd.AddCommand(authCmd)
}
