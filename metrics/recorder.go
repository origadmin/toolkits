// Copyright (c) 2024 GodCong. All rights reserved.

// Package metrics provides the metrics recorder
package metrics

// Recorder is an interface for recording various metrics.
type Recorder interface {
	Record(metrics Metrics)
}
