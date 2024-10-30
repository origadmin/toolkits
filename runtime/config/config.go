package config

import (
	"github.com/go-kratos/kratos/v2/config"
)

type (
	Config   = config.Config
	Source   = config.Source
	Option   = config.Option
	Decoder  = config.Decoder
	Resolver = config.Resolver
	Merge    = config.Merge
)

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	File   Type = "file"
	Consul Type = "consul"
	ETCD   Type = "etcd"
	//Nacos  Type = "nacos"
	// Apollo Type = "apollo"
	// Kubernetes Type = "kubernetes"
	// Polaris Type = "polaris"
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

func DefaultConsulConfig() *SourceConfig {
	return &SourceConfig{
		Type: Consul.String(),
		Consul: &SourceConfig_Consul{
			Address: "127.0.0.1:8500",
			Scheme:  "http",
		},
	}
}

func NewFileConfig(path string) *SourceConfig {
	return &SourceConfig{
		Type: File.String(),
		File: &SourceConfig_File{
			Path: path,
		},
	}
}
