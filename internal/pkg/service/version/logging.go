package version

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
