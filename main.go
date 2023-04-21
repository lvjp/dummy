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

	debugsvc "github.com/lvjp/dummy/internal/pkg/service/debug"
	stringsvc "github.com/lvjp/dummy/internal/pkg/service/string"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	mux.Handle("/string/", http.StripPrefix("/string", initStringService()))
	mux.Handle("/debug/", http.StripPrefix("/debug", initDebugService()))
	mux.Handle("/metrics", promhttp.Handler())

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

func initStringService() http.Handler {
	fieldKeys := []string{"method"}

	svc := stringsvc.NewService()
	svc = stringsvc.NewLoggingService(slog.Default().With("service", "string"), svc)
	svc = stringsvc.NewInstrumentingService(
		promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "string_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "api",
			Subsystem: "string_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		svc,
	)
	return stringsvc.MakeHandler(svc)
}

func initDebugService() http.Handler {
	fieldKeys := []string{"method"}
	svc := debugsvc.NewService()
	svc = debugsvc.NewLoggingservice(slog.Default().With("service", "debug"), svc)
	svc = debugsvc.NewInstrumentingService(
		promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "debug_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "api",
			Subsystem: "debug_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		svc,
	)
	return debugsvc.MakeHandler(svc)
}
