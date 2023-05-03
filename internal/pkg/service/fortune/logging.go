package fortune

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

func (lm loggingMiddleware) Create(fortune string) (uuid string, err error) {
	defer func(begin time.Time) {
		took := time.Since(begin)
		lm.logger.Info().
			Str("method", "create").
			Str("fortune", fortune).
			Str("uuid", uuid).
			Err(err).
			Dur("took", took).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.Create(fortune)
}

func (lm loggingMiddleware) Read(uuid string) (fortune string, err error) {
	defer func(begin time.Time) {
		took := time.Since(begin)
		lm.logger.Info().
			Str("method", "read").
			Str("uuid", uuid).
			Err(err).
			Dur("took", took).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.Read(uuid)
}

func (lm loggingMiddleware) Update(uuid, fortune string) (err error) {
	defer func(begin time.Time) {
		took := time.Since(begin)
		lm.logger.Info().
			Str("method", "update").
			Str("fortune", fortune).
			Str("uuid", uuid).
			Err(err).
			Dur("took", took).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.Update(uuid, fortune)
}

func (lm loggingMiddleware) Delete(uuid string) (err error) {
	defer func(begin time.Time) {
		took := time.Since(begin)
		lm.logger.Info().
			Str("method", "delete").
			Str("uuid", uuid).
			Err(err).
			Dur("took", took).
			Msg("Request processed")
	}(time.Now())

	return lm.svc.Delete(uuid)
}
