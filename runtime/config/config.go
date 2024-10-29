package config

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

type (
	Config = config.Config
	Source = config.Source
)

func LoadConfig(path string) Config {
	return config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
}

func DefaultConsulConfig() *DiscoveryConfig {
	return &DiscoveryConfig{
		Type: "consul",
		Consul: &Consul{
			Address: "127.0.0.1:8500",
			Scheme:  "http",
		},
	}
}
