package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	stringsvc "github.com/lvjp/dummy/internal/pkg/service/string"
	versionsvc "github.com/lvjp/dummy/internal/pkg/service/version"
	"github.com/sourcegraph/conc/pool"
	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s := newServer()

	p := pool.New().WithErrors().WithContext(mainCtx)
	p.Go(func(ctx context.Context) error {
		s.BaseContext = func(l net.Listener) context.Context {
			return ctx
		}

		slog.With("addr", s.Addr).Info("Start HTTP server")

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

func newServer() *http.Server {
	mux := http.NewServeMux()

	{
		svc := stringsvc.NewStringService()
		svc = stringsvc.LoggingMiddleware(slog.Default())(svc)

		mux.Handle("/string/uppercase", stringsvc.NewUppercaseHandler(svc))
		mux.Handle("/string/count", stringsvc.NewCountHandler(svc))
	}

	{
		svc := versionsvc.NewVersionService()
		svc = versionsvc.LoggingMiddleware(slog.Default())(svc)
		mux.Handle("/version", versionsvc.NewVersionHandler(svc))
	}

	// PORT is definied by Scaleway serverless containers
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
