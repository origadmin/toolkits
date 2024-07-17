// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package metrics provides the metrics recorder
package metrics

type MetricType string

const (
	MetricRequestsTotal          MetricType = "requests_total"
	MetricSlowRequestsTotal      MetricType = "requests_slow_total"
	MetricResponseSizeBytes      MetricType = "response_size_bytes"
	MetricRequestSizeBytes       MetricType = "request_size_bytes"
	MetricErrorsTotal            MetricType = "errors_total"
	MetricEvent                  MetricType = "event"
	MetricSiteEvent              MetricType = "site_event"
	MetricRequestDurationSeconds MetricType = "request_duration_seconds"
	MetricSummaryLatency         MetricType = "summary_latency"
	MetricRequestsInFlight       MetricType = "requests_in_flight"
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
	MetricLabelJob      = "job"
	MetricLabelError    = "error"
)

var metricLabelNames = map[MetricType][]string{
	MetricRequestsTotal:          {MetricLabelInstance, MetricLabelJob, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricRequestSizeBytes:       {MetricLabelInstance, MetricLabelJob, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricResponseSizeBytes:      {MetricLabelInstance, MetricLabelJob, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricRequestDurationSeconds: {MetricLabelInstance, MetricLabelJob, MetricLabelHandler, MetricLabelMethod},
	MetricSummaryLatency:         {MetricLabelInstance, MetricLabelJob, MetricLabelHandler, MetricLabelMethod},
	MetricErrorsTotal:            {MetricLabelInstance, MetricLabelJob, MetricLabelError},
	MetricEvent:                  {MetricLabelInstance, MetricLabelJob, "event"},
	MetricSiteEvent:              {MetricLabelInstance, MetricLabelJob, "event", "site"},
	MetricRequestsInFlight:       {MetricLabelInstance, MetricLabelJob, "state"},
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
