package string

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var promUppercaseLabels = prometheus.Labels{"method": "uppercase"}
var promCountLabels = prometheus.Labels{"method": "count"}

type instrumentingService struct {
	requestCount   *prometheus.CounterVec
	requestLatency *prometheus.HistogramVec
	svc            Service
}

func NewInstrumentingService(counter *prometheus.CounterVec, latency *prometheus.HistogramVec, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		svc:            s,
	}
}

func (is *instrumentingService) Uppercase(s string) (string, error) {
	defer func(begin time.Time) {
		is.requestCount.With(promUppercaseLabels).Add(1)
		is.requestLatency.With(promUppercaseLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Uppercase(s)
}

func (is *instrumentingService) Count(s string) int {
	defer func(begin time.Time) {
		is.requestCount.With(promCountLabels).Add(1)
		is.requestLatency.With(promCountLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Count(s)
}
