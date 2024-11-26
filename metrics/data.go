/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package metrics

type MetricData struct {
	Type     MetricType
	TraceID  string
	Endpoint string
	Method   string
	Code     int
	SendSize int64
	RecvSize int64
	Latency  float64
	Succeed  bool
}
