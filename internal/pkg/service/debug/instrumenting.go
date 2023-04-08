package debug

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	svc            Service
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		svc:            s,
	}
}

func (is *instrumentingService) Version() string {
	defer func(begin time.Time) {
		is.requestCount.With("method", "version").Add(1)
		is.requestLatency.With("method", "version").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Version()
}

func (is *instrumentingService) BuildTimestamp() string {
	defer func(begin time.Time) {
		is.requestCount.With("method", "buildtimestamp").Add(1)
		is.requestLatency.With("method", "buildtimestamp").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.BuildTimestamp()
}

func (is *instrumentingService) Environment() []string {
	defer func(begin time.Time) {
		is.requestCount.With("method", "environment").Add(1)
		is.requestLatency.With("method", "environment").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Environment()
}
