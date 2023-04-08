package version

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
