package string

import (
	"time"

	"golang.org/x/exp/slog"
)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    Service
}

func NewLoggingService(logger *slog.Logger, s Service) Service {
	return &loggingMiddleware{logger, s}
}

func (lm loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		lm.logger.Info("Request processed",
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return lm.svc.Uppercase(s)
}

func (lm loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		lm.logger.Info("Request processed",
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	return lm.svc.Count(s)
}
