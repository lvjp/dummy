package main

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

	"github.com/sourcegraph/conc/pool"
	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

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

	// PORT is definied by Scaleway serverless containers
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			return mainCtx
		},
	}

	p := pool.New().WithErrors().WithContext(mainCtx)
	p.Go(func(ctx context.Context) error {
		s.BaseContext = func(l net.Listener) context.Context {
			return ctx
		}

		logger := slog.With(slog.String("module", "authserver"))

		logger.With("port", port).Info("Start HTTP server")

		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Cannot server listen and serve", err)
			return err
		}

		slog.Info("HTTP server gracefully shutdown")
		return nil
	})

	p.Go(func(ctx context.Context) error {
		<-ctx.Done()

		slog.Info("Shutdown sequence started")
		return s.Shutdown(context.Background())
	})

	if err := p.Wait(); err != nil {
		slog.Error("Unexpected error on shutdown", err)
		return
	}

	slog.Info("Shutdown sequence finished")
}
