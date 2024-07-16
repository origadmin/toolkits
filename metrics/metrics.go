// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package metrics provides the metrics recorder
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricType string

const (
	MetricRequestTotal     MetricType = "request_total"
	MetricSendBytes        MetricType = "send_bytes"
	MetricRecvBytes        MetricType = "recv_bytes"
	MetricException        MetricType = "exception"
	MetricEvent            MetricType = "event"
	MetricSiteEvent        MetricType = "site_event"
	MetricHistogramLatency MetricType = "histogram_latency"
	MetricSummaryLatency   MetricType = "summary_latency"
	MetricGaugeState       MetricType = "gauge_state"
)

func (m MetricType) String() string {
	return string(m)
}

type dummyMetrics struct{}

// Metrics is the metrics interface for metrics
type Metrics interface {
	Enabled() bool // Enabled returns whether metrics is enabled
	Observe(reporter Reporter)
	Log(api, method, code string, sendBytes, recvBytes, latency float64)
	RequestLog(module, api, method, code string)
	SendBytesLog(module, api, method, code string, length float64)
	RecvBytesLog(module, api, method, code string, length float64)
	HistogramLatencyLog(module, api, method string, latency float64)
	SummaryLatencyLog(module, api, method string, latency float64)
	ExceptionLog(module, exception string)
	EventLog(module, event string)
	SiteEventLog(module, event, site string)
	StateLog(module, state string, value float64)
}

var metricLabels = map[MetricType][]string{
	MetricRequestTotal:     {"app", "module", "api", "method", "code"},
	MetricSendBytes:        {"app", "module", "api", "method", "code"},
	MetricRecvBytes:        {"app", "module", "api", "method", "code"},
	MetricException:        {"app", "module", "exception"},
	MetricEvent:            {"app", "module", "event"},
	MetricSiteEvent:        {"app", "module", "event", "site"},
	MetricHistogramLatency: {"app", "module", "api", "method"},
	MetricSummaryLatency:   {"app", "module", "api", "method"},
	MetricGaugeState:       {"app", "module", "state"},
}

// MetricTypeLabels returns the labels for the given metric type
func MetricTypeLabels(metricType MetricType) []string {
	if v, ok := metricLabels[metricType]; ok {
		return v
	}
	return []string{}
}

// MetricLabels returns the labels for the given metric type
func MetricLabels() map[MetricType][]string {
	return metricLabels
}

func (d dummyMetrics) Enabled() bool {
	return false
}
func (d dummyMetrics) Observe(reporter Reporter) {}

func (d dummyMetrics) Log(api, method, code string, sendBytes, recvBytes, latency float64) {}

func (d dummyMetrics) RequestLog(module, api, method, code string) {}

func (d dummyMetrics) SendBytesLog(module, api, method, code string, byte float64) {}

func (d dummyMetrics) RecvBytesLog(module, api, method, code string, byte float64) {}

func (d dummyMetrics) HistogramLatencyLog(module, api, method string, latency float64) {}

func (d dummyMetrics) SummaryLatencyLog(module, api, method string, latency float64) {}

func (d dummyMetrics) ExceptionLog(module, exception string) {}

func (d dummyMetrics) EventLog(module, event string) {}

func (d dummyMetrics) SiteEventLog(module, event, site string) {}

func (d dummyMetrics) StateLog(module, state string, value float64) {}

func (d dummyMetrics) ResetCounter() {}

func (d dummyMetrics) RegisterCustomCollector(c prometheus.Collector) {}

var DummyMetrics Metrics = dummyMetrics{}
