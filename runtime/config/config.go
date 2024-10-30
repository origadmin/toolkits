package config

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

type (
	Config   = config.Config
	Source   = config.Source
	Option   = config.Option
	Decoder  = config.Decoder
	Resolver = config.Resolver
	Merge    = config.Merge
)

func New(opts ...Option) Config {
	return config.New(opts...)
}

func WithSource(sources ...Source) Option {
	return config.WithSource(sources...)
}

func WithDecoder(decoder Decoder) Option {
	return config.WithDecoder(decoder)
}

func WithResolver(resolver Resolver) Option {
	return config.WithResolver(resolver)
}

func WithMergeFunc(merge Merge) Option {
	return config.WithMergeFunc(merge)
}

func LoadConfig(path string) Config {
	return config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
}

func DefaultConsulConfig() *SourceConfig {
	return &SourceConfig{
		Type: "consul",
		Consul: &SourceConfig_Consul{
			Address: "127.0.0.1:8500",
			Scheme:  "http",
		},
	}
}
