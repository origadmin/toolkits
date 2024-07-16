// Copyright (config) 2024 OrigAdmin. All rights reserved.

// Package prometheus is the data access object
package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/origadmin/toolkits/metrics"
)

// Prometheus is a Prometheus wrapper
type Prometheus struct {
	config    metrics.Config
	registry  *prometheus.Registry
	gauge     *prometheus.GaugeVec
	histogram *prometheus.HistogramVec
	summary   *prometheus.SummaryVec
	counters  map[metrics.MetricType]*prometheus.CounterVec
}

func (p *Prometheus) Enabled() bool {
	return p.config.Enable
}

func (p *Prometheus) Observe(reporter metrics.Reporter) {
	p.Log(reporter.API(), reporter.Method(), reporter.Code(), float64(reporter.BytesWritten()), float64(reporter.BytesReceived()), float64(reporter.Latency()))
}

// init initializes the Prometheus wrapper
func (p *Prometheus) init() {
	labels := metrics.MetricLabels()
	for i, l := range p.config.MetricLabels {
		if len(l) > 0 {
			labels[i] = l
		}
	}
	p.counters = map[metrics.MetricType]*prometheus.CounterVec{
		metrics.MetricRequestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: p.config.Namespace,
				Subsystem: p.config.SubSystem,
				Name:      metrics.MetricRequestTotal.String(),
				Help:      "number of module requests",
			},
			labels[metrics.MetricRequestTotal],
		),
		metrics.MetricSendBytes: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: p.config.Namespace,
				Subsystem: p.config.SubSystem,
				Name:      metrics.MetricSendBytes.String(),
				Help:      "number of module send bytes",
			},
			labels[metrics.MetricSendBytes],
		),
		metrics.MetricRecvBytes: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: p.config.Namespace,
				Subsystem: p.config.SubSystem,
				Name:      metrics.MetricRecvBytes.String(),
				Help:      "number of module recv bytes",
			},
			labels[metrics.MetricRecvBytes],
		),
		metrics.MetricException: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: p.config.Namespace,
				Subsystem: p.config.SubSystem,
				Name:      metrics.MetricException.String(),
				Help:      "number of module exception",
			},
			labels[metrics.MetricException],
		),
		metrics.MetricEvent: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: p.config.Namespace,
				Subsystem: p.config.SubSystem,
				Name:      metrics.MetricEvent.String(),
				Help:      "number of module event",
			},
			labels[metrics.MetricEvent],
		),
		metrics.MetricSiteEvent: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: p.config.Namespace,
				Subsystem: p.config.SubSystem,
				Name:      metrics.MetricSiteEvent.String(),
				Help:      "number of module site event",
			},
			labels[metrics.MetricSiteEvent],
		),
	}

	p.registry.MustRegister(CollectorsFromMap(p.counters)...)

	// Histogram for module latency
	p.histogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricHistogramLatency.String(),
			Help:      "histogram of module latency",
			Buckets:   p.config.Buckets,
		},
		labels[metrics.MetricHistogramLatency],
	)
	p.registry.MustRegister(p.histogram)

	// Summary for module latency
	p.summary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  p.config.Namespace,
			Subsystem:  p.config.SubSystem,
			Name:       metrics.MetricSummaryLatency.String(),
			Help:       "summary of module latency",
			Objectives: p.config.Objectives,
		},
		labels[metrics.MetricSummaryLatency],
	)
	p.registry.MustRegister(p.summary)

	// Gauge for app state
	p.gauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricGaugeState.String(),
			Help:      "gauge of app state",
		},
		labels[metrics.MetricGaugeState],
	)
	p.registry.MustRegister(p.gauge)

	// Register default collectors if enabled
	if p.config.DefaultCollect {
		p.registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		p.registry.MustRegister(collectors.NewGoCollector())
	}
}

func (p *Prometheus) Handler() http.Handler {
	if p.config.ListenPort == 0 {
		return nil
	}

	mux := http.NewServeMux()
	handle := promhttp.HandlerFor(p.registry, promhttp.HandlerOpts{})
	mux.Handle("/metrics", promhttp.InstrumentMetricHandler(
		p.registry,
		http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			username, pwd, ok := req.BasicAuth()
			if !ok || !(username == p.config.BasicUserName && pwd == p.config.BasicPassword) {
				resp.WriteHeader(http.StatusUnauthorized)
				_, _ = resp.Write([]byte("401 Unauthorized"))
				return
			}
			handle.ServeHTTP(resp, req)
		})),
	)
	return mux

}

// Log logs the API request and its details.
//
// Parameters: api string, method string, code string, sendBytes float64, recvBytes float64, latency float64.
func (p *Prometheus) Log(api, method, code string, sendBytes, recvBytes, latency float64) {
	if len(p.config.LogMethod) > 0 {
		if _, ok := p.config.LogMethod[method]; !ok {
			method = ""
		}
	}
	if len(p.config.LogAPI) > 0 {
		if _, ok := p.config.LogAPI[api]; !ok {
			api = ""
		}
	}

	p.CounterInc(metrics.MetricRequestTotal, p.config.Application, "this", api, method, code)
	if sendBytes > 0 {
		p.CounterAdd(metrics.MetricSendBytes, sendBytes, p.config.Application, "this", api, method, code)
	}
	if recvBytes > 0 {
		p.CounterAdd(metrics.MetricRecvBytes, recvBytes, p.config.Application, "this", api, method, code)
	}
	if len(p.config.Buckets) > 0 {
		p.histogram.WithLabelValues(p.config.Application, "this", api, method).Observe(latency)
	}
	if len(p.config.Objectives) > 0 {
		p.summary.WithLabelValues(p.config.Application, "this", api, method).Observe(latency)
	}
}

// RequestLog logs the request with the given module, api, method, and code.
//
// Parameters: module string, api string, method string, code string.
// Return type: none.
func (p *Prometheus) RequestLog(module, api, method, code string) {
	p.CounterInc(metrics.MetricRequestTotal, p.config.Application, module, api, method, code)
}

// SendBytesLog logs the byte count for a specific module, API, method, and code.
//
// Parameters: module, api, method, code string, byte float64
func (p *Prometheus) SendBytesLog(module, api, method, code string, byte float64) {
	p.CounterAdd(metrics.MetricSendBytes, byte, p.config.Application, module, api, method, code)
}

// RecvBytesLog is a Go function that logs received bytes.
//
// It takes the following parameters: module (string), api (string), method (string), code (string), byte (float64).
// It does not return anything.
func (p *Prometheus) RecvBytesLog(module, api, method, code string, byte float64) {
	p.CounterAdd(metrics.MetricRecvBytes, byte, p.config.Application, module, api, method, code)
}

// HistogramLatencyLog logs the latency of a histogram module, API, and method.
//
// module: the name of the module.
// api: the name of the API.
// method: the name of the method.
// latency: the latency of the API call.
func (p *Prometheus) HistogramLatencyLog(module, api, method string, latency float64) {
	p.histogram.WithLabelValues(p.config.Application, module, api, method).Observe(latency)
}

// SummaryLatencyLog logs the latency of a summary module, API, and method.
//
// module: the name of the module.
// api: the name of the API.
// method: the name of the method.
// latency: the latency of the API call.
func (p *Prometheus) SummaryLatencyLog(module, api, method string, latency float64) {
	p.summary.WithLabelValues(p.config.Application, module, api, method).Observe(latency)
}

// ExceptionLog logs the occurrence of an exception in a module.
//
// module: the name of the module.
// exception: the name of the exception.
func (p *Prometheus) ExceptionLog(module, exception string) {
	p.CounterInc(metrics.MetricException, p.config.Application, module, exception)
}

// EventLog logs an event in a module.
//
// module: the name of the module.
// event: the name of the event.
func (p *Prometheus) EventLog(module, event string) {
	p.CounterInc(metrics.MetricEvent, p.config.Application, module, event)
}

// SiteEventLog logs an event in a module for a specific site.
//
// module: the name of the module.
// event: the name of the event.
// site: the name of the site.
func (p *Prometheus) SiteEventLog(module, event, site string) {
	p.CounterInc(metrics.MetricEvent, p.config.Application, module, event, site)
}

// CounterInc increments the counter for a given metric type and labels.
// The counter is incremented using the WithLabelValues method of the corresponding counter vector.
//
// Parameters:
// - metricType: the type of metric to increment the counter for
// - labels: the labels to associate with the counter
func (p *Prometheus) CounterInc(metricType metrics.MetricType, labels ...string) {
	// Check if the metric type is valid
	if c, ok := p.counters[metricType]; ok {
		// Increment the counter for the given metric type and labels
		c.WithLabelValues(labels...).Inc()
		return
	}
}

// CounterAdd increments the counter for a given metric type and labels.
// The counter is incremented using the WithLabelValues method of the corresponding counter vector.
//
// Parameters:
// - metricType: the type of metric to increment the counter for
// - value: the value to add to the counter
// - labels: the labels to associate with the counter
func (p *Prometheus) CounterAdd(metricType metrics.MetricType, value float64, labels ...string) {
	// Check if the metric type is valid
	if c, ok := p.counters[metricType]; ok {
		// Increment the counter for the given metric type and labels
		c.WithLabelValues(labels...).Add(value)
		return
	}
}

// StateLog logs a state in a module.
//
// module: the name of the module.
// state: the name of the state.
// value: the value of the state.
func (p *Prometheus) StateLog(module, state string, value float64) {
	p.gauge.WithLabelValues(p.config.Application, module, state).Set(value)
}

// ResetCounter resets the counters to zero.
func (p *Prometheus) ResetCounter() {
	for i := range p.counters {
		p.counters[i].Reset()
	}
}

// RegisterCustomCollector registers a custom collector.
//
// config: the collector to be registered.
func (p *Prometheus) RegisterCustomCollector(c prometheus.Collector) {
	p.registry.MustRegister(c)
}

// WithPrometheus creates a Prometheus metrics with given config.
//
// conf: the config for the wrapper.
func WithPrometheus(conf *metrics.Config) metrics.Metrics {
	// Initialize and run the metrics if enabled.
	if !conf.Enable {
		return metrics.DummyMetrics
	}

	// if conf.Application == "" {
	// 	conf.Application = defaultApplication
	// }
	//
	// if conf.Namespace == "" {
	// 	conf.Namespace = defaultNamespace
	// }
	//
	// if conf.SubSystem == "" {
	// 	conf.SubSystem = defaultSubSystem
	// }
	//
	// // Set default listen port if not provided and enable Prometheus.
	// if conf.Enable && conf.ListenPort == 0 {
	// 	conf.ListenPort = defaultListenPort
	// }

	// Create wrapper with given config.
	m := &Prometheus{
		config:   *conf,
		registry: prometheus.NewRegistry(),
	}
	m.init()

	return m
}
