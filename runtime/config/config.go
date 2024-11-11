package config

import (
	kratosconfig "github.com/go-kratos/kratos/v2/config"

	"github.com/origadmin/toolkits/runtime/internal/config/v1"
)

type (
	AuthorizationConfig                             = config.AuthorizationConfig
	AuthorizationConfigMultiError                   = config.AuthorizationConfigMultiError
	AuthorizationConfigValidationError              = config.AuthorizationConfigValidationError
	CasbinConfig                                    = config.CasbinConfig
	CasbinConfigMultiError                          = config.CasbinConfigMultiError
	CasbinConfigValidationError                     = config.CasbinConfigValidationError
	CorsConfig                                      = config.CorsConfig
	CorsConfigMultiError                            = config.CorsConfigMultiError
	CorsConfigValidationError                       = config.CorsConfigValidationError
	DataConfig                                      = config.DataConfig
	DataConfig_Database                             = config.DataConfig_Database
	DataConfig_DatabaseMultiError                   = config.DataConfig_DatabaseMultiError
	DataConfig_DatabaseValidationError              = config.DataConfig_DatabaseValidationError
	DataConfigMultiError                            = config.DataConfigMultiError
	DataConfigValidationError                       = config.DataConfigValidationError
	EntrySelectorConfig                             = config.EntrySelectorConfig
	EntrySelectorConfigMultiError                   = config.EntrySelectorConfigMultiError
	EntrySelectorConfigValidationError              = config.EntrySelectorConfigValidationError
	LoggerConfig                                    = config.LoggerConfig
	LoggerConfigMultiError                          = config.LoggerConfigMultiError
	LoggerConfigValidationError                     = config.LoggerConfigValidationError
	MessageConfig                                   = config.MessageConfig
	MessageConfig_ActiveMQ                          = config.MessageConfig_ActiveMQ
	MessageConfig_ActiveMQMultiError                = config.MessageConfig_ActiveMQMultiError
	MessageConfig_ActiveMQValidationError           = config.MessageConfig_ActiveMQValidationError
	MessageConfig_Kafka                             = config.MessageConfig_Kafka
	MessageConfig_KafkaMultiError                   = config.MessageConfig_KafkaMultiError
	MessageConfig_KafkaValidationError              = config.MessageConfig_KafkaValidationError
	MessageConfig_MQTT                              = config.MessageConfig_MQTT
	MessageConfig_MQTTMultiError                    = config.MessageConfig_MQTTMultiError
	MessageConfig_MQTTValidationError               = config.MessageConfig_MQTTValidationError
	MessageConfig_NATS                              = config.MessageConfig_NATS
	MessageConfig_NATSMultiError                    = config.MessageConfig_NATSMultiError
	MessageConfig_NATSValidationError               = config.MessageConfig_NATSValidationError
	MessageConfig_NSQ                               = config.MessageConfig_NSQ
	MessageConfig_NSQMultiError                     = config.MessageConfig_NSQMultiError
	MessageConfig_NSQValidationError                = config.MessageConfig_NSQValidationError
	MessageConfig_Pulsar                            = config.MessageConfig_Pulsar
	MessageConfig_PulsarMultiError                  = config.MessageConfig_PulsarMultiError
	MessageConfig_PulsarValidationError             = config.MessageConfig_PulsarValidationError
	MessageConfig_RabbitMQ                          = config.MessageConfig_RabbitMQ
	MessageConfig_RabbitMQMultiError                = config.MessageConfig_RabbitMQMultiError
	MessageConfig_RabbitMQValidationError           = config.MessageConfig_RabbitMQValidationError
	MessageConfig_Redis                             = config.MessageConfig_Redis
	MessageConfig_RedisMultiError                   = config.MessageConfig_RedisMultiError
	MessageConfig_RedisValidationError              = config.MessageConfig_RedisValidationError
	MessageConfig_RocketMQ                          = config.MessageConfig_RocketMQ
	MessageConfig_RocketMQMultiError                = config.MessageConfig_RocketMQMultiError
	MessageConfig_RocketMQValidationError           = config.MessageConfig_RocketMQValidationError
	MessageConfigMultiError                         = config.MessageConfigMultiError
	MessageConfigValidationError                    = config.MessageConfigValidationError
	MetricConfig                                    = config.MetricConfig
	MetricConfigMultiError                          = config.MetricConfigMultiError
	MetricConfigValidationError                     = config.MetricConfigValidationError
	MiddlewareConfig                                = config.MiddlewareConfig
	MiddlewareConfigMultiError                      = config.MiddlewareConfigMultiError
	MiddlewareConfigValidationError                 = config.MiddlewareConfigValidationError
	RegistryConfig                                  = config.RegistryConfig
	RegistryConfig_Consul                           = config.RegistryConfig_Consul
	RegistryConfig_ConsulMultiError                 = config.RegistryConfig_ConsulMultiError
	RegistryConfig_ConsulValidationError            = config.RegistryConfig_ConsulValidationError
	RegistryConfig_ETCD                             = config.RegistryConfig_ETCD
	RegistryConfig_ETCDMultiError                   = config.RegistryConfig_ETCDMultiError
	RegistryConfig_ETCDValidationError              = config.RegistryConfig_ETCDValidationError
	RegistryConfigMultiError                        = config.RegistryConfigMultiError
	RegistryConfigValidationError                   = config.RegistryConfigValidationError
	SecurityConfig                                  = config.SecurityConfig
	SecurityConfigMultiError                        = config.SecurityConfigMultiError
	SecurityConfigValidationError                   = config.SecurityConfigValidationError
	ServiceConfig                                   = config.ServiceConfig
	ServiceConfig_Entry                             = config.ServiceConfig_Entry
	ServiceConfig_EntryMultiError                   = config.ServiceConfig_EntryMultiError
	ServiceConfig_EntryValidationError              = config.ServiceConfig_EntryValidationError
	ServiceConfig_GINS                              = config.ServiceConfig_GINS
	ServiceConfig_GINSMultiError                    = config.ServiceConfig_GINSMultiError
	ServiceConfig_GINSValidationError               = config.ServiceConfig_GINSValidationError
	ServiceConfig_GRPC                              = config.ServiceConfig_GRPC
	ServiceConfig_GRPCMultiError                    = config.ServiceConfig_GRPCMultiError
	ServiceConfig_GRPCValidationError               = config.ServiceConfig_GRPCValidationError
	ServiceConfig_HTTP                              = config.ServiceConfig_HTTP
	ServiceConfig_HTTPMultiError                    = config.ServiceConfig_HTTPMultiError
	ServiceConfig_HTTPValidationError               = config.ServiceConfig_HTTPValidationError
	ServiceConfig_Middleware                        = config.ServiceConfig_Middleware
	ServiceConfig_Middleware_Cors                   = config.ServiceConfig_Middleware_Cors
	ServiceConfig_Middleware_CorsMultiError         = config.ServiceConfig_Middleware_CorsMultiError
	ServiceConfig_Middleware_CorsValidationError    = config.ServiceConfig_Middleware_CorsValidationError
	ServiceConfig_Middleware_Logger                 = config.ServiceConfig_Middleware_Logger
	ServiceConfig_Middleware_LoggerMultiError       = config.ServiceConfig_Middleware_LoggerMultiError
	ServiceConfig_Middleware_LoggerValidationError  = config.ServiceConfig_Middleware_LoggerValidationError
	ServiceConfig_Middleware_Metrics                = config.ServiceConfig_Middleware_Metrics
	ServiceConfig_Middleware_MetricsMultiError      = config.ServiceConfig_Middleware_MetricsMultiError
	ServiceConfig_Middleware_MetricsValidationError = config.ServiceConfig_Middleware_MetricsValidationError
	ServiceConfig_Middleware_Traces                 = config.ServiceConfig_Middleware_Traces
	ServiceConfig_Middleware_TracesMultiError       = config.ServiceConfig_Middleware_TracesMultiError
	ServiceConfig_Middleware_TracesValidationError  = config.ServiceConfig_Middleware_TracesValidationError
	ServiceConfig_MiddlewareMultiError              = config.ServiceConfig_MiddlewareMultiError
	ServiceConfig_MiddlewareValidationError         = config.ServiceConfig_MiddlewareValidationError
	ServiceConfig_Websocket                         = config.ServiceConfig_Websocket
	ServiceConfig_WebsocketMultiError               = config.ServiceConfig_WebsocketMultiError
	ServiceConfig_WebsocketValidationError          = config.ServiceConfig_WebsocketValidationError
	ServiceConfigMultiError                         = config.ServiceConfigMultiError
	ServiceConfigValidationError                    = config.ServiceConfigValidationError
	SourceConfig                                    = config.SourceConfig
	SourceConfig_Consul                             = config.SourceConfig_Consul
	SourceConfig_ConsulMultiError                   = config.SourceConfig_ConsulMultiError
	SourceConfig_ConsulValidationError              = config.SourceConfig_ConsulValidationError
	SourceConfig_ETCD                               = config.SourceConfig_ETCD
	SourceConfig_ETCDMultiError                     = config.SourceConfig_ETCDMultiError
	SourceConfig_ETCDValidationError                = config.SourceConfig_ETCDValidationError
	SourceConfig_Custom                             = config.SourceConfig_Custom
	SourceConfig_CustomMultiError                   = config.SourceConfig_CustomMultiError
	SourceConfig_CustomValidationError              = config.SourceConfig_CustomValidationError
	SourceConfig_File                               = config.SourceConfig_File
	SourceConfig_FileMultiError                     = config.SourceConfig_FileMultiError
	SourceConfig_FileValidationError                = config.SourceConfig_FileValidationError
	SourceConfigMultiError                          = config.SourceConfigMultiError
	SourceConfigValidationError                     = config.SourceConfigValidationError
	TaskConfig                                      = config.TaskConfig
	TaskConfig_Asynq                                = config.TaskConfig_Asynq
	TaskConfig_AsynqMultiError                      = config.TaskConfig_AsynqMultiError
	TaskConfig_AsynqValidationError                 = config.TaskConfig_AsynqValidationError
	TaskConfig_Cron                                 = config.TaskConfig_Cron
	TaskConfig_CronMultiError                       = config.TaskConfig_CronMultiError
	TaskConfig_CronValidationError                  = config.TaskConfig_CronValidationError
	TaskConfig_Machinery                            = config.TaskConfig_Machinery
	TaskConfig_MachineryMultiError                  = config.TaskConfig_MachineryMultiError
	TaskConfig_MachineryValidationError             = config.TaskConfig_MachineryValidationError
	TaskConfigMultiError                            = config.TaskConfigMultiError
	TaskConfigValidationError                       = config.TaskConfigValidationError
	TraceConfig                                     = config.TraceConfig
	TraceConfigMultiError                           = config.TraceConfigMultiError
	TraceConfigValidationError                      = config.TraceConfigValidationError

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
	File_config_v1_cors_proto       = config.File_config_v1_cors_proto
	File_config_v1_data_proto       = config.File_config_v1_data_proto
	File_config_v1_logger_proto     = config.File_config_v1_logger_proto
	File_config_v1_metrics_proto    = config.File_config_v1_metrics_proto
	File_config_v1_middleware_proto = config.File_config_v1_middleware_proto
	File_config_v1_registry_proto   = config.File_config_v1_registry_proto
	File_config_v1_security_proto   = config.File_config_v1_security_proto
	File_config_v1_service_proto    = config.File_config_v1_service_proto
	File_config_v1_source_proto     = config.File_config_v1_source_proto
	File_config_v1_tracer_proto     = config.File_config_v1_tracer_proto

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
