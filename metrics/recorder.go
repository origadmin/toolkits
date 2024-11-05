// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package metrics provides the metrics recorder
package metrics

// Recorder is an interface for recording various metrics.
type Recorder interface {
	RequestTotal(module, handler, method, code string)
	CounterSendBytes(module, handler, method, code string, length int64)
	CounterRecvBytes(module, handler, method, code string, length int64)
	RequestDurationSeconds(module, handler, method string, latency int64)
	SummaryLatencyLog(module, handler, method string, latency int64)
	CounterException(module, errors string)
	CounterEvent(module, event string)
	CounterSiteEvent(module, event, site string)
	RequestsInFlight(module, state string, value int64)
}
