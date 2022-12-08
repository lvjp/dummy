package cmd

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"time"

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
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			data := bytes.NewBufferString("Hello from the dummy authentication server")
			length := data.Len()

			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			written, err := io.Copy(w, data)
			if err != nil {
				slog.Error("Cannot write HTTP response", err,
					slog.Int64("written", written),
					slog.Int("length", length),
				)
			}
		})

		s := &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Cannot listen and serve", err)
			return
		}

		slog.Info("Server gracefully shutdowned")
	},
}

func init() {
	authCmd.AddCommand(
		authServeCmd,
	)

	rootCmd.AddCommand(authCmd)
}
