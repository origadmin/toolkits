// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package metrics provides the metrics recorder
package metrics

// Metrics is the metrics interface for metrics
type Metrics interface {
	// Enabled returns whether metrics is enabled
	Enabled() bool

	// Observe records a metric using the provided Reporter
	Observe(reporter Reporter)

	// Recorder returns the Recorder for this Metrics instance
	Recorder() Recorder
}
