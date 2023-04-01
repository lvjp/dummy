package string

import (
	"time"

	"golang.org/x/exp/slog"
)

func LoggingMiddleware(logger *slog.Logger) ServiceMiddleware {
	return func(next StringService) StringService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger *slog.Logger
	StringService
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

	output, err = lm.StringService.Uppercase(s)
	return
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

	n = lm.StringService.Count(s)
	return
}
