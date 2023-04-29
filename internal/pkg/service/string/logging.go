package string

import (
	"time"

	"github.com/rs/zerolog"
)

type loggingMiddleware struct {
	logger zerolog.Logger
	svc    Service
}

func NewLoggingService(logger zerolog.Logger, s Service) Service {
	return &loggingMiddleware{logger, s}
}

func (lm loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		lm.logger.Info().
			Str("method", "uppercase").
			Str("input", s).
			Str("output", output).
			Err(err).
			Dur("took", time.Since(begin)).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.Uppercase(s)
}

func (lm loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		lm.logger.Info().
			Str("method", "count").
			Str("input", s).
			Int("n", n).
			Dur("took", time.Since(begin)).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.Count(s)
}
