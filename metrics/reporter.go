package metrics

import (
	"github.com/origadmin/toolkits/context"
)

// Reporter is an interface that defines methods for reporting metrics.
type Reporter interface {
	// Context returns the context of the reporter.
	Context() context.Context
	// Handler returns the Handler Handler endpoint associated with the metric.
	Handler() string
	// Method returns the HTTP method used for the Handler call.
	Method() string
	// Code returns the status code of the Handler response.
	Code() string
	// WriteSize returns the number of bytes write in the response.
	WriteSize() int64
	// ReadSize returns the number of bytes read in the request.
	ReadSize() int64
	// Succeed returns a boolean indicating if the request was successful.
	Succeed() bool
	// TraceID returns the unique trace ID associated with the request.
	TraceID() string
	// Latency returns the latency of the request in milliseconds.
	Latency() int64
}
