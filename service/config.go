package service

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Addrs     []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config
	Options   map[string]any
}

type DiscoveryConfig struct {
	TTL time.Duration
}

type Setting func(*Config)

func NewConfig(opts ...Setting) *Config {
	cfg := new(Config)
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
