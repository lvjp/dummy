package debug

import (
	"time"

	"golang.org/x/exp/slog"
)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    Service
}

func NewLoggingservice(logger *slog.Logger, s Service) Service {
	return loggingMiddleware{logger, s}
}

func (lm loggingMiddleware) Version() (output string) {
	defer func(begin time.Time) {
		lm.logger.Info("Request processed",
			"method", "version",
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	return lm.svc.Version()
}

func (lm loggingMiddleware) BuildTimestamp() (output string) {
	defer func(begin time.Time) {
		lm.logger.Info("Request processed",
			"method", "buildtimestamp",
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	return lm.svc.BuildTimestamp()
}

func (lm loggingMiddleware) Environment() (output []string) {
	defer func(begin time.Time) {
		lm.logger.Info("Request processed",
			"method", "environment",
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	return lm.svc.Environment()
}
