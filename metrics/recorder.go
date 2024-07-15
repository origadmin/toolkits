// Copyright (c) 2024 GodCong. All rights reserved.

// Package metrics provides the metrics recorder
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Recorder is an interface for recording various metrics.
type Recorder interface {
	// Log records the general metrics like API, method, code, sendBytes, recvBytes, and latency.
	Log(api, method, code string, sendBytes, recvBytes, latency float64)

	// RequestLog records the request details like module, API, method, and code.
	RequestLog(module, api, method, code string)

	// SendBytesLog records the bytes sent for a specific module, API, method, and code.
	SendBytesLog(module, api, method, code string, byte float64)

	// RecvBytesLog records the bytes received for a specific module, API, method, and code.
	RecvBytesLog(module, api, method, code string, byte float64)

	// HistogramLatencyLog records the latency in a histogram format for a specific module, API, and method.
	HistogramLatencyLog(module, api, method string, latency float64)

	// SummaryLatencyLog records the latency in a summary format for a specific module, API, and method.
	SummaryLatencyLog(module, api, method string, latency float64)

	// ExceptionLog records an exception event for a specific module and exception.
	ExceptionLog(module, exception string)

	// EventLog records a general event for a specific module and event.
	EventLog(module, event string)

	// SiteEventLog records a site-specific event for a specific module, event, and site.
	SiteEventLog(module, event, site string)

	// StateLog records the state value for a specific module and state.
	StateLog(module, state string, value float64)

	// ResetCounter resets any internal counters maintained by the recorder.
	ResetCounter()

	// RegisterCustomCollector allows registering a custom prometheus.Collector.
	RegisterCustomCollector(c prometheus.Collector)
}

type dummyRecorder struct{}

func (d dummyRecorder) Log(api, method, code string, sendBytes, recvBytes, latency float64) {

}

func (d dummyRecorder) RequestLog(module, api, method, code string) {

}

func (d dummyRecorder) SendBytesLog(module, api, method, code string, byte float64) {

}

func (d dummyRecorder) RecvBytesLog(module, api, method, code string, byte float64) {

}

func (d dummyRecorder) HistogramLatencyLog(module, api, method string, latency float64) {

}

func (d dummyRecorder) SummaryLatencyLog(module, api, method string, latency float64) {

}

func (d dummyRecorder) ExceptionLog(module, exception string) {

}

func (d dummyRecorder) EventLog(module, event string) {

}

func (d dummyRecorder) SiteEventLog(module, event, site string) {

}

func (d dummyRecorder) StateLog(module, state string, value float64) {

}

func (d dummyRecorder) ResetCounter() {

}

func (d dummyRecorder) RegisterCustomCollector(c prometheus.Collector) {

}

var DummyRecorder Recorder = dummyRecorder{}
