package metrics

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/origadmin/toolkits/metrics"
)

type Metrics struct {
	serv *http.Server
}

// Start starts the Metrics by listening for incoming connections.
func (s Metrics) Start(_ context.Context) error {
	// Start the HTTP server in a goroutine to allow for concurrent connections.
	go func() {
		err := s.serv.ListenAndServe()
		if err != nil {
			return
		}
	}()
	return nil
}

// Stop stops the Metrics gracefully by shutting down the HTTP server.
func (s Metrics) Stop(ctx context.Context) error {
	return s.serv.Shutdown(ctx)
}

// New creates a new instance of the Metrics based on the provided configuration.
func New(conf *metrics.Config) (*Metrics, error) {
	if conf.ListenPort == 0 {
		return nil, fmt.Errorf("listen port is empty")
	}
	if conf.ReadTimeout == 0 {
		conf.ReadTimeout = 10 * time.Second
	}
	if conf.WriteTimeout == 0 {
		conf.WriteTimeout = 10 * time.Second
	}
	if conf.MaxHeaderBytes == 0 {
		conf.MaxHeaderBytes = 1 << 20
	}

	serv := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.ListenPort),
		Handler:        conf.Handler,
		ReadTimeout:    conf.ReadTimeout,
		WriteTimeout:   conf.WriteTimeout,
		MaxHeaderBytes: conf.MaxHeaderBytes,
	}

	return &Metrics{
		serv: serv,
	}, nil
}
