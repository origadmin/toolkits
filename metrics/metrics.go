// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package metrics provides the metrics recorder
package metrics

import (
	"github.com/origadmin/toolkits/context"
)

type MetricType string

const (
	//MetricRequestSizeBytes       MetricType = "request_size_bytes"       // MetricRequestSizeBytes is the request size in bytes
	//MetricRequestDurationSeconds MetricType = "request_duration_seconds" // MetricRequestDurationSeconds is the request duration in seconds
	//MetricRequestsTotal          MetricType = "requests_total"           // MetricRequestsTotal is the total number of requests
	//MetricRequestsSlowTotal      MetricType = "requests_slow_total"      // MetricRequestsSlowTotal is the total number of slow requests
	//MetricRequestsInFlight       MetricType = "requests_in_flight"       // MetricRequestsInFlight is the number of requests in flight
	//MetricResponseSizeBytes      MetricType = "response_size_bytes"      // MetricResponseSizeBytes is the response size in bytes
	//MetricErrorsTotal            MetricType = "errors_total"             // MetricErrorsTotal is the total number of errors
	//MetricEvent                  MetricType = "event"                    // MetricEvent is the event
	//MetricSiteEvent              MetricType = "site_event"               // MetricSiteEvent is the site event
	//MetricSummaryLatency         MetricType = "summary_latency"          // MetricSummaryLatency is the summary latency
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
