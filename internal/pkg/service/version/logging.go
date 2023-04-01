package version

import (
	"time"

	"golang.org/x/exp/slog"
)

func LoggingMiddleware(logger *slog.Logger) ServiceMiddleware {
	return func(next VersionService) VersionService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger *slog.Logger
	VersionService
}

func (lm loggingMiddleware) Version() (output string) {
	defer func(begin time.Time) {
		lm.logger.Info("Request processed",
			"method", "version",
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = lm.VersionService.Version()
	return
}
