package fortune

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var promCreateLabels = prometheus.Labels{"method": "create"}
var promReadLabels = prometheus.Labels{"method": "read"}
var promUpdateLabels = prometheus.Labels{"method": "update"}
var promDeleteLabels = prometheus.Labels{"method": "delete"}

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

func (is *instrumentingService) Create(fortune string) (string, error) {
	defer func(begin time.Time) {
		is.requestCount.With(promCreateLabels).Add(1)
		is.requestLatency.With(promCreateLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Create(fortune)
}

func (is *instrumentingService) Read(uuid string) (fortune string, err error) {
	defer func(begin time.Time) {
		is.requestCount.With(promReadLabels).Add(1)
		is.requestLatency.With(promReadLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Read(uuid)
}

func (is *instrumentingService) Update(uuid, fortune string) (err error) {
	defer func(begin time.Time) {
		is.requestCount.With(promUpdateLabels).Add(1)
		is.requestLatency.With(promUpdateLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Update(uuid, fortune)
}

func (is *instrumentingService) Delete(uuid string) (err error) {
	defer func(begin time.Time) {
		is.requestCount.With(promDeleteLabels).Add(1)
		is.requestLatency.With(promDeleteLabels).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.svc.Delete(uuid)
}
