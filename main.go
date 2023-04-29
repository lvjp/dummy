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
	"github.com/rs/zerolog"
	"github.com/sourcegraph/conc/pool"
)

func main() {
	logger := zerolog.New(os.Stdout)

	// Check for local debugging
	if os.Getenv("SCW_APPLICATION_NAME") == "" {
		logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		})
	}

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s := newServer(logger)

	p := pool.New().WithErrors().WithContext(mainCtx)
	p.Go(func(ctx context.Context) error {
		s.BaseContext = func(l net.Listener) context.Context {
			return ctx
		}

		logger.Info().Str("addr", s.Addr).Msg("Start HTTP server")

		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error().Err(err).Msg("Cannot server listen and serve")
			return err
		}

		logger.Info().Msg("HTTP server gracefully shutdown")
		return nil
	})

	p.Go(func(ctx context.Context) error {
		<-ctx.Done()

		logger.Info().Msg("Shutdown sequence started")
		return s.Shutdown(context.Background())
	})

	if err := p.Wait(); err != nil {
		logger.Error().Err(err).Msg("Unexpected error on shutdown")
		return
	}

	logger.Info().Msg("Shutdown sequence finished")
}

func newServer(logger zerolog.Logger) *http.Server {
	mux := http.NewServeMux()

	mux.Handle("/string/", http.StripPrefix("/string", initStringService(logger)))
	mux.Handle("/debug/", http.StripPrefix("/debug", initDebugService(logger)))
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

func initStringService(logger zerolog.Logger) http.Handler {
	fieldKeys := []string{"method"}

	svc := stringsvc.NewService()
	svc = stringsvc.NewLoggingService(logger.With().Str("service", "string").Logger(), svc)
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

func initDebugService(logger zerolog.Logger) http.Handler {
	fieldKeys := []string{"method"}
	svc := debugsvc.NewService()
	svc = debugsvc.NewLoggingservice(logger.With().Str("service", "string").Logger(), svc)
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
