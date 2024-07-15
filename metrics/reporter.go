package metrics

import (
	"github.com/origadmin/toolkits/context"
)

// Reporter is an interface that defines methods for reporting metrics.
type Reporter interface {
	// Context returns the context of the reporter.
	Context() context.Context
	// API returns the API endpoint associated with the metric.
	API() string
	// Method returns the HTTP method used for the API call.
	Method() string
	// Code returns the status code of the API response.
	Code() string
	// BytesWritten returns the number of bytes written in the response.
	BytesWritten() int64
	// BytesReceived returns the number of bytes received in the request.
	BytesReceived() int64
	// Succeed returns a boolean indicating if the request was successful.
	Succeed() bool
	// TraceID returns the unique trace ID associated with the request.
	TraceID() string
	// Latency returns the latency of the request in milliseconds.
	Latency() int64
}
