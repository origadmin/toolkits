package config

import (
	kratosconfig "github.com/go-kratos/kratos/v2/config"

	"github.com/origadmin/toolkits/runtime/internal/config/v1"
	"github.com/origadmin/toolkits/runtime/internal/middlewares/v1"
)

type (
	RegistryConfig                       = config.RegistryConfig
	RegistryConfig_Consul                = config.RegistryConfig_Consul
	RegistryConfig_ConsulMultiError      = config.RegistryConfig_ConsulMultiError
	RegistryConfig_ConsulValidationError = config.RegistryConfig_ConsulValidationError
	RegistryConfig_ETCD                  = config.RegistryConfig_ETCD
	RegistryConfig_ETCDMultiError        = config.RegistryConfig_ETCDMultiError
	RegistryConfig_ETCDValidationError   = config.RegistryConfig_ETCDValidationError
	RegistryConfigMultiError             = config.RegistryConfigMultiError
	RegistryConfigValidationError        = config.RegistryConfigValidationError

	SourceConfig                       = config.SourceConfig
	SourceConfig_Consul                = config.SourceConfig_Consul
	SourceConfig_ConsulMultiError      = config.SourceConfig_ConsulMultiError
	SourceConfig_ConsulValidationError = config.SourceConfig_ConsulValidationError
	SourceConfig_ETCD                  = config.SourceConfig_ETCD
	SourceConfig_ETCDMultiError        = config.SourceConfig_ETCDMultiError
	SourceConfig_ETCDValidationError   = config.SourceConfig_ETCDValidationError
	SourceConfig_File                  = config.SourceConfig_File
	SourceConfig_FileMultiError        = config.SourceConfig_FileMultiError
	SourceConfig_FileValidationError   = config.SourceConfig_FileValidationError
	SourceConfigMultiError             = config.SourceConfigMultiError
	SourceConfigValidationError        = config.SourceConfigValidationError

	AuthorizationConfig                = middlewares.AuthorizationConfig
	AuthorizationConfigMultiError      = middlewares.AuthorizationConfigMultiError
	AuthorizationConfigValidationError = middlewares.AuthorizationConfigValidationError
	CasbinConfig                       = middlewares.CasbinConfig
	CasbinConfigMultiError             = middlewares.CasbinConfigMultiError
	CasbinConfigValidationError        = middlewares.CasbinConfigValidationError
	CorsConfig                         = middlewares.CorsConfig
	CorsConfigMultiError               = middlewares.CorsConfigMultiError
	CorsConfigValidationError          = middlewares.CorsConfigValidationError
	LoggerConfig                       = middlewares.LoggerConfig
	LoggerConfigMultiError             = middlewares.LoggerConfigMultiError
	LoggerConfigValidationError        = middlewares.LoggerConfigValidationError
	MetricConfig                       = middlewares.MetricConfig
	MetricConfigMultiError             = middlewares.MetricConfigMultiError
	MetricConfigValidationError        = middlewares.MetricConfigValidationError
	MiddlewareConfig                   = middlewares.MiddlewareConfig
	MiddlewareConfigMultiError         = middlewares.MiddlewareConfigMultiError
	MiddlewareConfigValidationError    = middlewares.MiddlewareConfigValidationError
	SecurityConfig                     = middlewares.SecurityConfig
	SecurityConfigMultiError           = middlewares.SecurityConfigMultiError
	SecurityConfigValidationError      = middlewares.SecurityConfigValidationError

	Config   = kratosconfig.Config
	Decoder  = kratosconfig.Decoder
	KeyValue = kratosconfig.KeyValue
	Merge    = kratosconfig.Merge
	Observer = kratosconfig.Observer
	Option   = kratosconfig.Option
	Reader   = kratosconfig.Reader
	Resolver = kratosconfig.Resolver
	Source   = kratosconfig.Source
	Value    = kratosconfig.Value
	Watcher  = kratosconfig.Watcher
)

var (
	ErrNotFound = kratosconfig.ErrNotFound
)

func New(opts ...Option) Config {
	return kratosconfig.New(opts...)
}

func WithDecoder(d Decoder) Option {
	return kratosconfig.WithDecoder(d)
}

func WithMergeFunc(m Merge) Option {
	return kratosconfig.WithMergeFunc(m)
}

func WithResolveActualTypes(enableConvertToType bool) Option {
	return kratosconfig.WithResolveActualTypes(enableConvertToType)
}
func WithResolver(r Resolver) Option {
	return kratosconfig.WithResolver(r)
}
func WithSource(s ...Source) Option {
	return kratosconfig.WithSource(s...)
}
