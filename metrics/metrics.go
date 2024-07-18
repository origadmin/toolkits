// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package metrics provides the metrics recorder
package metrics

type MetricType string

const (
	MetricRequestSizeBytes       MetricType = "request_size_bytes"
	MetricRequestDurationSeconds MetricType = "request_duration_seconds"
	MetricRequestsTotal          MetricType = "requests_total"
	MetricRequestsSlowTotal      MetricType = "requests_slow_total"
	MetricRequestsInFlight       MetricType = "requests_in_flight"
	MetricResponseSizeBytes      MetricType = "response_size_bytes"
	MetricErrorsTotal            MetricType = "errors_total"
	MetricEvent                  MetricType = "event"
	MetricSiteEvent              MetricType = "site_event"
	MetricSummaryLatency         MetricType = "summary_latency"
)

func (m MetricType) String() string {
	return string(m)
}

type dummyMetrics struct{}

// Metrics is the metrics interface for metrics
type Metrics interface {
	Enabled() bool // Enabled returns whether metrics is enabled
	Observe(reporter Reporter)
	Log(handler, method, code string, sendBytes, recvBytes, latency float64)
	RequestTotal(module, handler, method, code string)
	ResponseSize(module, handler, method, code string, length float64)
	RequestSize(module, handler, method, code string, length float64)
	RequestDurationSeconds(module, handler, method string, latency float64)
	SummaryLatencyLog(module, handler, method string, latency float64)
	ErrorsTotal(module, errors string)
	Event(module, event string)
	SiteEvent(module, event, site string)
	RequestsInFlight(module, state string, value float64)
}

const (
	MetricLabelInstance = "instance"
	MetricLabelHandler  = "handler"
	MetricLabelCode     = "code"
	MetricLabelMethod   = "method"
	MetricLabelModule   = "module"
	MetricLabelError    = "error"
)

var metricLabelNames = map[MetricType][]string{
	MetricRequestsTotal:          {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricRequestSizeBytes:       {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricResponseSizeBytes:      {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricRequestDurationSeconds: {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod},
	MetricSummaryLatency:         {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod},
	MetricErrorsTotal:            {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod, MetricLabelError},
	MetricEvent:                  {MetricLabelInstance, MetricLabelModule, "event"},
	MetricSiteEvent:              {MetricLabelInstance, MetricLabelModule, "event", "site"},
	MetricRequestsInFlight:       {MetricLabelInstance, MetricLabelModule, "state"},
}

// MetricLabelNames returns the labels for the given metric type
func MetricLabelNames(metricType MetricType) []string {
	if v, ok := metricLabelNames[metricType]; ok {
		return v
	}
	return []string{}
}

// MetricLabels returns the labels for the given metric type
func MetricLabels() map[MetricType][]string {
	return metricLabelNames
}

func (d dummyMetrics) Enabled() bool {
	return false
}
func (d dummyMetrics) Observe(reporter Reporter) {}

func (d dummyMetrics) Log(handler, method, code string, sendBytes, recvBytes, latency float64) {}

func (d dummyMetrics) RequestTotal(module, handler, method, code string) {}

func (d dummyMetrics) RequestDurationSeconds(module, handler, method string, latency float64) {}

func (d dummyMetrics) SummaryLatencyLog(module, handler, method string, latency float64) {}

func (d dummyMetrics) ErrorsTotal(module, errors string) {}

func (d dummyMetrics) Event(module, event string) {}

func (d dummyMetrics) SiteEvent(module, event, site string) {}

func (d dummyMetrics) RequestsInFlight(module, state string, value float64) {}

func (d dummyMetrics) ResponseSize(module, handler, method, code string, length float64) {}

func (d dummyMetrics) RequestSize(module, handler, method, code string, length float64) {}

func (d dummyMetrics) ResetCounter() {}

var DummyMetrics Metrics = dummyMetrics{}
