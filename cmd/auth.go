package cmd

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"
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
		mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			slog.Info("Incoming request",
				slog.Time("time", time.Now()),
				slog.String("host", r.Host),
				slog.String("remote", r.RemoteAddr),
				slog.String("proto", r.Proto),
				slog.String("URI", r.RequestURI),
				slog.String("Referer", r.Header.Get("Referer")),
				slog.String("User-Agent", r.Header.Get("User-Agent")),
			)

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
			BaseContext: func(_ net.Listener) context.Context {
				return mainCtx
			},
		}

		g, gCtx := errgroup.WithContext(mainCtx)
		g.Go(func() error {
			logger := slog.With(slog.String("module", "authserver"))

			logger.Info("Start HTTP server")

			if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				slog.Error("Cannot server listen and serve", err)
				return err
			}

			slog.Info("HTTP server gracefully shutdown")
			return nil
		})

		g.Go(func() error {
			<-gCtx.Done()

			slog.Info("Shutdown sequence started")
			return s.Shutdown(context.Background())
		})

		if err := g.Wait(); err != nil {
			slog.Error("Unexpected error on shutdown", err)
			return
		}

		slog.Info("Shutdown sequence finished")
	},
}

func init() {
	authCmd.AddCommand(
		authServeCmd,
	)

	rootCmd.AddCommand(authCmd)
}
