package debug

import (
	"time"

	"github.com/rs/zerolog"
)

type loggingMiddleware struct {
	logger zerolog.Logger
	svc    Service
}

func NewLoggingservice(logger zerolog.Logger, s Service) Service {
	return loggingMiddleware{logger, s}
}

func (lm loggingMiddleware) Version() (output string) {
	defer func(begin time.Time) {
		lm.logger.Info().
			Str("method", "version").
			Str("output", output).
			Dur("took", time.Since(begin)).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.Version()
}

func (lm loggingMiddleware) BuildTimestamp() (output string) {
	defer func(begin time.Time) {
		lm.logger.Info().
			Str("method", "buildtimestamp").
			Str("output", output).
			Dur("took", time.Since(begin)).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.BuildTimestamp()
}

func (lm loggingMiddleware) Environment() (output []string) {
	defer func(begin time.Time) {
		lm.logger.Info().
			Str("method", "environment").
			Strs("output", output).
			Dur("took", time.Since(begin)).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.Environment()
}
