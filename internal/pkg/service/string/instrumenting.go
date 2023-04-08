package string

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

func (is *instrumentingService) Uppercase(s string) (string, error) {
	defer func(begin time.Time) {
		is.requestCount.With("method", "uppercase").Add(1)
		is.requestLatency.With("method", "uppercase").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Uppercase(s)
}

func (is *instrumentingService) Count(s string) int {
	defer func(begin time.Time) {
		is.requestCount.With("method", "count").Add(1)
		is.requestLatency.With("method", "count").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Count(s)
}
