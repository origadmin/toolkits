// Copyright (config) 2024 OrigAdmin. All rights reserved.

// Package prometheus is the data access object
package prometheus

import (
	"net/http"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/origadmin/toolkits/metrics"
)

// Prometheus is a Prometheus wrapper
// counter only statistics count value in prometheus
// requestDurationSeconds statistics bucket,count,sum in prometheus
// summary statistics count,sum in prometheus
type Prometheus struct {
	config                 *Config
	module                 *atomic.Value
	registry               *prometheus.Registry
	requestsInFlightVec    *prometheus.GaugeVec
	requestDurationSeconds *prometheus.HistogramVec
	summary                *prometheus.SummaryVec
	requestTotalVec        *prometheus.CounterVec
	slowRequestsTotal      *prometheus.CounterVec
	errorsTotal            *prometheus.CounterVec
	responseSize           *prometheus.HistogramVec
	requestSize            *prometheus.HistogramVec
	event                  *prometheus.CounterVec
	siteEvent              *prometheus.CounterVec
}

// register initializes the Prometheus wrapper
func (p *Prometheus) register() {
	p.requestTotalVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricRequestsTotal.String(),
			Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		p.config.MetricLabels[metrics.MetricRequestsTotal],
	)

	p.event = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricEvent.String(),
			Help:      "number of module event",
		},
		p.config.MetricLabels[metrics.MetricEvent],
	)
	p.registry.MustRegister(p.event)

	p.siteEvent = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricSiteEvent.String(),
			Help:      "number of module site event",
		},
		p.config.MetricLabels[metrics.MetricSiteEvent],
	)
	p.registry.MustRegister(p.siteEvent)

	p.errorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricErrorsTotal.String(),
			Help:      "The HTTP request errors counter",
		},
		p.config.MetricLabels[metrics.MetricErrorsTotal],
	)
	p.registry.MustRegister(p.errorsTotal)

	p.slowRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricRequestsSlowTotal.String(),
			Help:      "The HTTP request slow counter",
		},
		p.config.MetricLabels[metrics.MetricRequestsSlowTotal],
	)
	p.registry.MustRegister(p.slowRequestsTotal)

	// Request size requestDurationSeconds
	p.requestSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricRequestSizeBytes.String(),
			Help:      "The HTTP request sizes in bytes.",
		},
		p.config.MetricLabels[metrics.MetricRequestSizeBytes],
	)
	p.registry.MustRegister(p.requestSize)

	// Response size requestDurationSeconds
	p.responseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricResponseSizeBytes.String(),
			Help:      "The HTTP response sizes in bytes.",
		},
		p.config.MetricLabels[metrics.MetricResponseSizeBytes],
	)
	p.registry.MustRegister(p.responseSize)

	// Request Duration Seconds for module latency
	p.requestDurationSeconds = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricRequestDurationSeconds.String(),
			Help:      "The HTTP request latencies in seconds.",
			Buckets:   p.config.Buckets,
		},
		p.config.MetricLabels[metrics.MetricRequestDurationSeconds],
	)
	p.registry.MustRegister(p.requestDurationSeconds)

	p.slowRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricErrorsTotal.String(),
			Help:      "The HTTP request errors counter",
		},
		p.config.MetricLabels[metrics.MetricErrorsTotal],
	)
	p.registry.MustRegister(p.slowRequestsTotal)

	// Summary for module latency
	p.summary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  p.config.Namespace,
			Subsystem:  p.config.SubSystem,
			Name:       metrics.MetricSummaryLatency.String(),
			Help:       "summary of module latency",
			Objectives: p.config.Objectives,
		},
		p.config.MetricLabels[metrics.MetricSummaryLatency],
	)
	p.registry.MustRegister(p.summary)

	// Gauge for app state
	p.requestsInFlightVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: p.config.Namespace,
			Subsystem: p.config.SubSystem,
			Name:      metrics.MetricRequestsInFlight.String(),
			Help:      "The HTTP requests in flight, partitioned by status code and HTTP method.",
		},
		p.config.MetricLabels[metrics.MetricRequestsInFlight],
	)
	p.registry.MustRegister(p.requestsInFlightVec)

	// Register default collectors if enabled
	if p.config.DefaultCollect {
		p.registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		p.registry.MustRegister(collectors.NewGoCollector())
	}
}

// Observe uses the Prometheus metric collector to observe and record metrics.
//
// This method takes a reporter parameter of type metrics.Reporter interface,
// which is used to report metric data. Through this method, a Prometheus instance
// can receive and process metric information from various sources, enabling
// real-time monitoring and data collection of system or application performance.
//
// The method does not return any value; its primary purpose is to update and maintain
// the internal metric collectors of Prometheus. By calling this method, one can ensure
// that relevant metric data is correctly collected and stored for subsequent analysis and querying.
func (p *Prometheus) Observe(reporter metrics.Reporter) {
	p.Log(reporter.Handler(), reporter.Method(), reporter.Code(), float64(reporter.WriteSize()), float64(reporter.ReadSize()), float64(reporter.Latency()))
}

// Log logs the Handler request and its details.
//
// Parameters: code string, method string, handler string, sendBytes float64, recvBytes float64, latency float64.
func (p *Prometheus) Log(code string, method, handler string, sendBytes, recvBytes, latency float64) {
	if len(p.config.LogMethod) > 0 {
		if _, ok := p.config.LogMethod[method]; !ok {
			return // ignore
		}
	}
	if len(p.config.LogHandler) > 0 {
		if _, ok := p.config.LogHandler[handler]; !ok {
			return // ignore
		}
	}

	p.RequestTotal("this", handler, method, code)
	p.ResponseSize("this", handler, method, code, sendBytes)
	p.RequestSize("this", handler, method, code, recvBytes)
	p.RequestDurationSeconds("this", handler, method, latency)
	p.SummaryLatencyLog("this", handler, method, latency)
}

// RequestTotal logs the request with the given module, handler, method, and code.
//
// Parameters: module string, handler string, method string, code string.
// Return type: none.
func (p *Prometheus) RequestTotal(module, handler, method, code string) {
	p.requestTotalVec.WithLabelValues(p.config.Application, module, handler, method, code)
}

// ResponseSize logs the byte count for a specific module, Handler, method, and code.
//
// It takes the following parameters: module (string), handler (string), method (string), code (string), length (float64).
// It does not return anything.
func (p *Prometheus) ResponseSize(module, handler, method, code string, length float64) {
	if length > 0 {
		p.responseSize.WithLabelValues(p.config.Application, module, handler, method, code).Observe(length)
	}
}

// RequestSize is a Go function that logs received bytes.
//
// It takes the following parameters: module (string), handler (string), method (string), code (string), length (float64).
// It does not return anything.
func (p *Prometheus) RequestSize(module, handler, method, code string, length float64) {
	if length > 0 {
		p.requestSize.WithLabelValues(p.config.Application, module, handler, method, code).Observe(length)
	}
}

// RequestDurationSeconds logs the latency of a requestDurationSeconds module, Handler, and method.
//
// module: the name of the module.
// handler: the name of the Handler handler.
// method: the name of the method.
// latency: the latency of the Handler call.
func (p *Prometheus) RequestDurationSeconds(module, handler, method string, latency float64) {
	if len(p.config.Buckets) > 0 {
		p.requestDurationSeconds.WithLabelValues(p.config.Application, module, handler, method).Observe(latency)
	}
}

// SummaryLatencyLog logs the latency of a summary module, Handler, and method.
//
// module: the name of the module.
// handler: the name of the Handler handler.
// method: the name of the method.
// latency: the latency of the Handler call.
func (p *Prometheus) SummaryLatencyLog(module, handler, method string, latency float64) {
	p.summary.WithLabelValues(p.config.Application, module, handler, method).Observe(latency)
}

// ErrorsTotal logs the occurrence of an exception in a module.
//
// module: the name of the module.
// errors: the name of the errors.
func (p *Prometheus) ErrorsTotal(module, handler, method, errors string) {
	p.errorsTotal.WithLabelValues(p.config.Application, module, handler, method, errors).Inc()
}

// Event logs an event in a module.
//
// module: the name of the module.
// event: the name of the event.
func (p *Prometheus) Event(module, event string) {
	p.event.WithLabelValues(p.config.Application, module, event)
}

// SiteEvent logs an event in a module for a specific site.
//
// module: the name of the module.
// event: the name of the event.
// site: the name of the site.
func (p *Prometheus) SiteEvent(module, event, site string) {
	p.event.WithLabelValues(p.config.Application, module, event, site)
}

// RequestsInFlight logs a state in a module.
//
// module: the name of the module.
// state: the name of the state.
// value: the value of the state.
func (p *Prometheus) RequestsInFlight(module, state string, value float64) {
	p.requestsInFlightVec.WithLabelValues(p.config.Application, module, state).Set(value)
}

func (p *Prometheus) SetModule(module string) {
	p.module.Store(module)
}

func (p *Prometheus) GetModule() string {
	return p.module.Load().(string)
	// if v, ok := p.module.Load().(string); ok {
	// 	return v
	// }
	// return "this"
}

// WithPrometheus creates a Prometheus metrics with given config.
//
// conf: the config for the metrics.Metrics.
func WithPrometheus(conf *Config) *Prometheus {
	conf.setup()

	// Create Prometheus metrics with given config.
	m := &Prometheus{
		config:   conf,
		module:   &atomic.Value{},
		registry: prometheus.NewRegistry(),
	}

	// Set default module if not provided.
	m.module.Store("this")

	m.register()
	return m
}

func HTTPHandler(prom *Prometheus) http.Handler {
	handle := promhttp.HandlerFor(prom.registry, promhttp.HandlerOpts{})
	return promhttp.InstrumentMetricHandler(
		prom.registry,
		handle)
}
