/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package metrics provides the metrics recorder
package metrics

import (
	"context"
)

type MetricType string

const (
	MetricUptime                  MetricType = "uptime"
	MetricRequestsTotal           MetricType = "requests_total"
	MetricRequestsDurationSeconds MetricType = "requests_duration_seconds"
	MetricRequestsInFlight        MetricType = "requests_in_flight"
	MetricRequestsSlowTotal       MetricType = "requests_slow_total"
	MetricCounterSendBytes        MetricType = "counter_send_bytes"
	MetricCounterRecvBytes        MetricType = "counter_recv_bytes"
	MetricHistogramLatency        MetricType = "histogram_latency"
	MetricSummaryLatency          MetricType = "summary_latency"
	MetricCounterException        MetricType = "counter_exception"
	MetricCounterEvent            MetricType = "counter_event"
	MetricCounterSiteEvent        MetricType = "counter_site_event"
)

func (m MetricType) String() string {
	return string(m)
}

type dummyMetrics struct{}

// Metrics is the metrics interface for metrics
type Metrics interface {
	Enabled() bool // Enabled returns whether metrics is enabled
	Disable()
	Observe(ctx context.Context, data MetricData)
	Log(ctx context.Context, handler, method string, code int, sendBytes, recvBytes int64, latency float64)
}

const (
	MetricLabelInstance = "instance"
	MetricLabelHandler  = "handler"
	MetricLabelCode     = "code"
	MetricLabelMethod   = "method"
	MetricLabelModule   = "module"
	MetricLabelErrors   = "errors"
	MetricLabelEvent    = "event"
	MetricLabelSite     = "site"
	MetricLabelState    = "state"
)

var metricLabelNames = map[MetricType][]string{
	MetricRequestsTotal:           {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricCounterSendBytes:        {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricCounterRecvBytes:        {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod, MetricLabelCode},
	MetricRequestsDurationSeconds: {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod},
	MetricSummaryLatency:          {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod},
	MetricCounterException:        {MetricLabelInstance, MetricLabelModule, MetricLabelHandler, MetricLabelMethod, MetricLabelErrors},
	MetricCounterEvent:            {MetricLabelInstance, MetricLabelModule, MetricLabelEvent},
	MetricCounterSiteEvent:        {MetricLabelInstance, MetricLabelModule, MetricLabelEvent, MetricLabelSite},
	MetricRequestsInFlight:        {MetricLabelInstance, MetricLabelModule, MetricLabelState},
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
