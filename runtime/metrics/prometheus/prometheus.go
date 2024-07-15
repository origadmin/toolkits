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

const (
	defaultApplication = "application_orig_admin"
	defaultNamespace   = "namespace_orig_admin"
	defaultSubSystem   = "admin"
	defaultListenPort  = 9100
)

// Config holds the configuration settings for the Prometheus wrapper
type Config struct {
	Enable         bool                // Enable flag
	Application    string              // Application name
	Namespace      string              // Namespace for Prometheus
	SubSystem      string              // Subsystem for Prometheus
	ListenPort     int                 // Port to listen on
	BasicUserName  string              // Basic authentication username
	BasicPassword  string              // Basic authentication password
	LogAPI         map[string]struct{} // Map of APIs to log
	LogMethod      map[string]struct{} // Map of methods to log
	Buckets        []float64           // Buckets for histogram
	Objectives     map[float64]float64 // Objectives for summary
	DefaultCollect bool                // Flag to enable default collectors
}

// Metrics is a Prometheus wrapper
type Metrics struct {
	config                             Config
	gatherer                           prometheus.Gatherer
	registerer                         prometheus.Registerer
	registry                           *prometheus.Registry
	gaugeState                         *prometheus.GaugeVec
	histogramLatency                   *prometheus.HistogramVec
	summaryLatency                     *prometheus.SummaryVec
	counterRequests, counterSendBytes  *prometheus.CounterVec
	counterRcevBytes, counterException *prometheus.CounterVec
	counterEvent, counterSiteEvent     *prometheus.CounterVec
}

func (m *Metrics) Enabled() bool {
	return m.config.Enable
}

func (m *Metrics) Observe(reporter metrics.Reporter) {
	m.Log(reporter.API(), reporter.Method(), reporter.Code(), float64(reporter.BytesWritten()), float64(reporter.BytesReceived()), float64(reporter.Latency()))
}

func (m *Metrics) Recorder() metrics.Recorder {
	return m
}

// init initializes the Prometheus wrapper
func (m *Metrics) init() {
	// Counter for module requests
	m.counterRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: m.config.Namespace,
			Subsystem: m.config.SubSystem,
			Name:      "counter_requests",
			Help:      "number of module requests",
		},
		[]string{"app", "module", "api", "method", "code"},
	)
	m.registry.MustRegister(m.counterRequests)

	// Counter for module send bytes
	m.counterSendBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: m.config.Namespace,
			Subsystem: m.config.SubSystem,
			Name:      "counter_send_bytes",
			Help:      "number of module send bytes",
		},
		[]string{"app", "module", "api", "method", "code"},
	)
	m.registry.MustRegister(m.counterSendBytes)

	// Counter for module receive bytes
	m.counterRcevBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: m.config.Namespace,
			Subsystem: m.config.SubSystem,
			Name:      "counter_recv_bytes",
			Help:      "number of module receive bytes",
		},
		[]string{"app", "module", "api", "method", "code"},
	)
	m.registry.MustRegister(m.counterRcevBytes)

	// Histogram for module latency
	m.histogramLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: m.config.Namespace,
			Subsystem: m.config.SubSystem,
			Name:      "histogram_latency",
			Help:      "histogram of module latency",
			Buckets:   m.config.Buckets,
		},
		[]string{"app", "module", "api", "method"},
	)
	m.registry.MustRegister(m.histogramLatency)

	// Summary for module latency
	m.summaryLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  m.config.Namespace,
			Subsystem:  m.config.SubSystem,
			Name:       "summary_latency",
			Help:       "summary of module latency",
			Objectives: m.config.Objectives,
		},
		[]string{"app", "module", "api", "method"},
	)
	m.registry.MustRegister(m.summaryLatency)

	// Gauge for app state
	m.gaugeState = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: m.config.Namespace,
			Subsystem: m.config.SubSystem,
			Name:      "gauge_state",
			Help:      "gauge of app state",
		},
		[]string{"app", "module", "state"},
	)
	m.registry.MustRegister(m.gaugeState)

	// Counter for module exception
	m.counterException = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: m.config.Namespace,
			Subsystem: m.config.SubSystem,
			Name:      "counter_exception",
			Help:      "number of module exception",
		},
		[]string{"app", "module", "exception"},
	)
	m.registry.MustRegister(m.counterException)

	// Counter for module event
	m.counterEvent = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: m.config.Namespace,
			Subsystem: m.config.SubSystem,
			Name:      "counter_event",
			Help:      "number of module event",
		},
		[]string{"app", "module", "event"},
	)
	m.registry.MustRegister(m.counterEvent)

	// Counter for module site event
	m.counterSiteEvent = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: m.config.Namespace,
			Subsystem: m.config.SubSystem,
			Name:      "counter_site_event",
			Help:      "number of module site event",
		},
		[]string{"app", "module", "event", "site"},
	)
	m.registry.MustRegister(m.counterSiteEvent)

	// Register default collectors if enabled
	if m.config.DefaultCollect {
		m.registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		m.registry.MustRegister(collectors.NewGoCollector())
	}
}

func (m *Metrics) Handler() http.Handler {
	if m.config.ListenPort == 0 {
		return nil
	}

	mux := http.NewServeMux()
	handle := promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{})
	mux.Handle("/metrics", promhttp.InstrumentMetricHandler(
		m.registry,
		http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			username, pwd, ok := req.BasicAuth()
			if !ok || !(username == m.config.BasicUserName && pwd == m.config.BasicPassword) {
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
func (m *Metrics) Log(api, method, code string, sendBytes, recvBytes, latency float64) {
	if len(m.config.LogMethod) > 0 {
		if _, ok := m.config.LogMethod[method]; !ok {
			return
		}
	}
	if len(m.config.LogAPI) > 0 {
		if _, ok := m.config.LogAPI[api]; !ok {
			return
		}
	}

	m.counterRequests.WithLabelValues(m.config.Application, "this", api, method, code).Inc()
	if sendBytes > 0 {
		m.counterSendBytes.WithLabelValues(m.config.Application, "this", api, method, code).Add(sendBytes)
	}
	if recvBytes > 0 {
		m.counterRcevBytes.WithLabelValues(m.config.Application, "this", api, method, code).Add(recvBytes)
	}
	if len(m.config.Buckets) > 0 {
		m.histogramLatency.WithLabelValues(m.config.Application, "this", api, method).Observe(latency)
	}
	if len(m.config.Objectives) > 0 {
		m.summaryLatency.WithLabelValues(m.config.Application, "this", api, method).Observe(latency)
	}
}

// RequestLog logs the request with the given module, api, method, and code.
//
// Parameters: module string, api string, method string, code string.
// Return type: none.
func (m *Metrics) RequestLog(module, api, method, code string) {
	m.counterRequests.WithLabelValues(m.config.Application, module, api, method, code).Inc()
}

// SendBytesLog logs the byte count for a specific module, API, method, and code.
//
// Parameters: module, api, method, code string, byte float64
func (m *Metrics) SendBytesLog(module, api, method, code string, byte float64) {
	m.counterSendBytes.WithLabelValues(m.config.Application, module, api, method, code).Add(byte)
}

// RecvBytesLog is a Go function that logs received bytes.
//
// It takes the following parameters: module (string), api (string), method (string), code (string), byte (float64).
// It does not return anything.
func (m *Metrics) RecvBytesLog(module, api, method, code string, byte float64) {
	m.counterRcevBytes.WithLabelValues(m.config.Application, module, api, method, code).Add(byte)
}

// HistogramLatencyLog logs the latency of a histogram module, API, and method.
//
// module: the name of the module.
// api: the name of the API.
// method: the name of the method.
// latency: the latency of the API call.
func (m *Metrics) HistogramLatencyLog(module, api, method string, latency float64) {
	m.histogramLatency.WithLabelValues(m.config.Application, module, api, method).Observe(latency)
}

// SummaryLatencyLog logs the latency of a summary module, API, and method.
//
// module: the name of the module.
// api: the name of the API.
// method: the name of the method.
// latency: the latency of the API call.
func (m *Metrics) SummaryLatencyLog(module, api, method string, latency float64) {
	m.summaryLatency.WithLabelValues(m.config.Application, module, api, method).Observe(latency)
}

// ExceptionLog logs the occurrence of an exception in a module.
//
// module: the name of the module.
// exception: the name of the exception.
func (m *Metrics) ExceptionLog(module, exception string) {
	m.counterException.WithLabelValues(m.config.Application, module, exception).Inc()
}

// EventLog logs an event in a module.
//
// module: the name of the module.
// event: the name of the event.
func (m *Metrics) EventLog(module, event string) {
	m.counterEvent.WithLabelValues(m.config.Application, module, event).Inc()
}

// SiteEventLog logs an event in a module for a specific site.
//
// module: the name of the module.
// event: the name of the event.
// site: the name of the site.
func (m *Metrics) SiteEventLog(module, event, site string) {
	m.counterSiteEvent.WithLabelValues(m.config.Application, module, event, site).Inc()
}

// StateLog logs a state in a module.
//
// module: the name of the module.
// state: the name of the state.
// value: the value of the state.
func (m *Metrics) StateLog(module, state string, value float64) {
	m.gaugeState.WithLabelValues(m.config.Application, module, state).Set(value)
}

// ResetCounter resets the counters to zero.
func (m *Metrics) ResetCounter() {
	// Reset site event counters.
	m.counterSiteEvent.Reset()

	// Reset event counters.
	m.counterEvent.Reset()

	// Reset exception counters.
	m.counterException.Reset()

	// Reset receive bytes counters.
	m.counterRcevBytes.Reset()

	// Reset send bytes counters.
	m.counterSendBytes.Reset()
}

// RegisterCustomCollector registers a custom collector.
//
// config: the collector to be registered.
func (m *Metrics) RegisterCustomCollector(c prometheus.Collector) {
	m.registry.MustRegister(c)
}

// NewMetrics creates a Prometheus wrapper with given config.
//
// conf: the config for the wrapper.
func NewMetrics(conf *Config) *Metrics {
	// Initialize and run the metrics if enabled.
	if !conf.Enable {
		return nil
	}

	if conf.Application == "" {
		conf.Application = defaultApplication
	}

	if conf.Namespace == "" {
		conf.Namespace = defaultNamespace
	}

	if conf.SubSystem == "" {
		conf.SubSystem = defaultSubSystem
	}

	// Set default listen port if not provided and enable Prometheus.
	if conf.Enable && conf.ListenPort == 0 {
		conf.ListenPort = defaultListenPort
	}

	// Create wrapper with given config.
	m := &Metrics{
		config:   *conf,
		registry: prometheus.NewRegistry(),
	}
	m.init()

	return m
}
