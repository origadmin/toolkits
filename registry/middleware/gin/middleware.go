package gin

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/metrics"
)

type ginReporter struct {
	ctx           context.Context
	api           string
	method        string
	code          int
	succeed       bool
	latency       float64
	receivedBytes int64
	writtenBytes  int64
}

func (g ginReporter) Code() string {
	return fmt.Sprintf("%d", g.code)
}

func (g ginReporter) Latency() int64 {
	return int64(g.latency)
}

func (g ginReporter) API() string {
	return g.api
}

func (g ginReporter) Method() string {
	return g.method
}

func (g ginReporter) Context() context.Context {
	return g.ctx
}

// BytesWritten returns the number of bytes sent in the response
func (g ginReporter) BytesWritten() int64 {
	return g.writtenBytes
}

// BytesReceived returns the number of bytes received in the request
func (g ginReporter) BytesReceived() int64 {
	return g.receivedBytes
}

// Succeed returns whether the request was successful or not
func (g ginReporter) Succeed() bool {
	return g.succeed
}

// TraceID returns the trace ID of the request
func (g ginReporter) TraceID() string {
	return context.FromTraceID(g.ctx)
}

// Adapter is the Gin Adapter.
type Adapter struct {
	RequestBytes  func(ctx *gin.Context) int64
	ResponseBytes func(ctx *gin.Context) int64
	API           func(ctx *gin.Context) string
	Method        func(ctx *gin.Context) string
	Code          func(ctx *gin.Context) int
}

// Middleware is the Gin Adapter function that logs API requests and their details.
//
// It takes the following parameters:
// enable: bool, whether to enable the middleware or not.
// reqKey: string, the key to retrieve the received bytes from the context.
func (obj *Adapter) Middleware(metrics metrics.Metrics) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !metrics.Enabled() {
			ctx.Next()
			return
		}

		metrics.Observe(obj.Reporter(ctx))
	}
}

// Reporter is the Gin Adapter function that logs API requests and their details.
func (obj *Adapter) Reporter(ctx *gin.Context) metrics.Reporter {
	start := time.Now()
	received := obj.RequestBytes(ctx)

	ctx.Next()
	latency := float64(time.Since(start).Milliseconds())
	api := obj.API(ctx)
	method := obj.Method(ctx)
	code := obj.Code(ctx)
	succeed := ctx.Writer.Status() < 400
	written := obj.ResponseBytes(ctx)
	// metrics.Recorder().Log(p, ctx.Request.Method, fmt.Sprintf("%d", ctx.Writer.Status()), float64(writeBytes),
	// float64(), latency)
	return &ginReporter{
		ctx:           ctx.Request.Context(),
		api:           api,
		method:        method,
		code:          code,
		succeed:       succeed,
		latency:       latency,
		writtenBytes:  written,
		receivedBytes: received,
	}
}

// Default is the default Gin Adapter.
func Default(reqKey string) *Adapter {
	return &Adapter{
		RequestBytes: func(ctx *gin.Context) int64 {
			if v, ok := ctx.Get(reqKey); ok {
				if b, ok := v.([]byte); ok {
					return int64(len(b))
				}
			}
			return 0
		},
		ResponseBytes: func(ctx *gin.Context) int64 {
			return int64(ctx.Writer.Size())
		},
		API: func(ctx *gin.Context) string {
			path := ctx.FullPath()
			for _, param := range ctx.Params {
				path = strings.ReplaceAll(path, param.Value, ":"+param.Key)
			}
			return path
		},
		Method: func(ctx *gin.Context) string {
			return ctx.Request.Method
		},
		Code: func(ctx *gin.Context) int {
			return ctx.Writer.Status()
		},
	}
}
