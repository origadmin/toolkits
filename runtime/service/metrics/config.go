package metrics

import (
	"net/http"
	"time"
)

const (
	defaultListenPort = 9100
)

type Config struct {
	UseSecure      bool
	BasicUserName  string
	BasicPassword  string
	Enable         bool
	HandlerFunc    http.HandlerFunc
	ListenPort     int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func (c *Config) setup() {
	// Set default listen port if not provided and enable Prometheus.
	if c.Enable && c.ListenPort == 0 {
		c.ListenPort = defaultListenPort
	}
	// Set default read and write timeout if not provided.
	if c.ReadTimeout == 0 {
		c.ReadTimeout = 10 * time.Second
	}
	// Set default read and write timeout if not provided.
	if c.WriteTimeout == 0 {
		c.WriteTimeout = 10 * time.Second
	}
	// Set default max header bytes if not provided.
	if c.MaxHeaderBytes == 0 {
		c.MaxHeaderBytes = 1 << 20
	}
}
