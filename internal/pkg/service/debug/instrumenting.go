package debug

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var promVersionLabels = prometheus.Labels{"method": "version"}
var promBuildTimestampLabels = prometheus.Labels{"method": "uppercase"}
var promEnvironmentLabels = prometheus.Labels{"method": "environment"}

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

func (is *instrumentingService) Version() string {
	defer func(begin time.Time) {
		is.requestCount.With(promVersionLabels).Add(1)
		is.requestLatency.With(promVersionLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Version()
}

func (is *instrumentingService) BuildTimestamp() string {
	defer func(begin time.Time) {
		is.requestCount.With(promBuildTimestampLabels).Add(1)
		is.requestLatency.With(promBuildTimestampLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.BuildTimestamp()
}

func (is *instrumentingService) Environment() []string {
	defer func(begin time.Time) {
		is.requestCount.With(promEnvironmentLabels).Add(1)
		is.requestLatency.With(promEnvironmentLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Environment()
}
