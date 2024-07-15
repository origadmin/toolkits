package metrics

import (
	"net/http"
	"time"
)

// Config represents the configuration for the Service server.
type Config struct {
	// ListenPort is the port on which the Service server will listen.
	ListenPort int

	// ReadTimeout is the maximum duration for reading the entire request, including the body.
	ReadTimeout time.Duration

	// WriteTimeout is the maximum duration before timing out writes of the response.
	WriteTimeout time.Duration

	// MaxHeaderBytes is the maximum number of bytes in the request header.
	MaxHeaderBytes int

	// Handler is the HTTP handler that will handle incoming requests.
	Handler http.Handler
}
