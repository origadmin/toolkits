package config

import (
	kratosconfig "github.com/go-kratos/kratos/v2/config"

	"github.com/origadmin/toolkits/runtime/internal/config/v1"
)

// Import the kratos config package and the internal config package
type (
	Metric                                       = config.Metric
	Cors                                         = config.Cors
	Data                                         = config.Data
	Data_Database                                = config.Data_Database
	Data_Storage                                 = config.Data_Storage
	Data_Storage_File                            = config.Data_Storage_File
	Data_Storage_Redis                           = config.Data_Storage_Redis
	Data_Storage_Mongo                           = config.Data_Storage_Mongo
	Task                                         = config.Task
	Task_Asynq                                   = config.Task_Asynq
	Task_Machinery                               = config.Task_Machinery
	Task_Cron                                    = config.Task_Cron
	Logger_Level                                 = config.Logger_Level
	Logger                                       = config.Logger
	SourceConfig                                 = config.SourceConfig
	SourceConfig_File                            = config.SourceConfig_File
	SourceConfig_Consul                          = config.SourceConfig_Consul
	SourceConfig_ETCD                            = config.SourceConfig_ETCD
	SourceConfig_Custom                          = config.SourceConfig_Custom
	Trace                                        = config.Trace
	Message                                      = config.Message
	Message_MQTT                                 = config.Message_MQTT
	Message_Kafka                                = config.Message_Kafka
	Message_RabbitMQ                             = config.Message_RabbitMQ
	Message_ActiveMQ                             = config.Message_ActiveMQ
	Message_NATS                                 = config.Message_NATS
	Message_NSQ                                  = config.Message_NSQ
	Message_Pulsar                               = config.Message_Pulsar
	Message_Redis                                = config.Message_Redis
	Message_RocketMQ                             = config.Message_RocketMQ
	EntrySelector                                = config.EntrySelector
	Service                                      = config.Service
	Service_Entry                                = config.Service_Entry
	Service_GINS                                 = config.Service_GINS
	Service_HTTP                                 = config.Service_HTTP
	Service_GRPC                                 = config.Service_GRPC
	Service_Websocket                            = config.Service_Websocket
	Service_Middleware                           = config.Service_Middleware
	Service_Middleware_Metrics                   = config.Service_Middleware_Metrics
	Service_Middleware_Traces                    = config.Service_Middleware_Traces
	Service_Middleware_Logger                    = config.Service_Middleware_Logger
	Service_Middleware_Cors                      = config.Service_Middleware_Cors
	Registry                                     = config.Registry
	Registry_Consul                              = config.Registry_Consul
	Registry_ETCD                                = config.Registry_ETCD
	Registry_Custom                              = config.Registry_Custom
	Authorization                                = config.Authorization
	Casbin                                       = config.Casbin
	Security                                     = config.Security
	Middleware                                   = config.Middleware
	Middleware_RateLimiter                       = config.Middleware_RateLimiter
	Middleware_RateLimiter_Redis                 = config.Middleware_RateLimiter_Redis
	Middleware_RateLimiter_Memory                = config.Middleware_RateLimiter_Memory
	CorsMultiError                               = config.CorsMultiError
	CorsValidationError                          = config.CorsValidationError
	DataMultiError                               = config.DataMultiError
	DataValidationError                          = config.DataValidationError
	Data_DatabaseMultiError                      = config.Data_DatabaseMultiError
	Data_DatabaseValidationError                 = config.Data_DatabaseValidationError
	Data_StorageMultiError                       = config.Data_StorageMultiError
	Data_StorageValidationError                  = config.Data_StorageValidationError
	Data_Storage_FileMultiError                  = config.Data_Storage_FileMultiError
	Data_Storage_FileValidationError             = config.Data_Storage_FileValidationError
	Data_Storage_RedisMultiError                 = config.Data_Storage_RedisMultiError
	Data_Storage_RedisValidationError            = config.Data_Storage_RedisValidationError
	Data_Storage_MongoMultiError                 = config.Data_Storage_MongoMultiError
	Data_Storage_MongoValidationError            = config.Data_Storage_MongoValidationError
	TaskMultiError                               = config.TaskMultiError
	TaskValidationError                          = config.TaskValidationError
	Task_AsynqMultiError                         = config.Task_AsynqMultiError
	Task_AsynqValidationError                    = config.Task_AsynqValidationError
	Task_MachineryMultiError                     = config.Task_MachineryMultiError
	Task_MachineryValidationError                = config.Task_MachineryValidationError
	Task_CronMultiError                          = config.Task_CronMultiError
	Task_CronValidationError                     = config.Task_CronValidationError
	LoggerMultiError                             = config.LoggerMultiError
	LoggerValidationError                        = config.LoggerValidationError
	SourceConfigMultiError                       = config.SourceConfigMultiError
	SourceConfigValidationError                  = config.SourceConfigValidationError
	SourceConfig_FileMultiError                  = config.SourceConfig_FileMultiError
	SourceConfig_FileValidationError             = config.SourceConfig_FileValidationError
	SourceConfig_ConsulMultiError                = config.SourceConfig_ConsulMultiError
	SourceConfig_ConsulValidationError           = config.SourceConfig_ConsulValidationError
	SourceConfig_ETCDMultiError                  = config.SourceConfig_ETCDMultiError
	SourceConfig_ETCDValidationError             = config.SourceConfig_ETCDValidationError
	SourceConfig_CustomMultiError                = config.SourceConfig_CustomMultiError
	SourceConfig_CustomValidationError           = config.SourceConfig_CustomValidationError
	TraceMultiError                              = config.TraceMultiError
	TraceValidationError                         = config.TraceValidationError
	MessageMultiError                            = config.MessageMultiError
	MessageValidationError                       = config.MessageValidationError
	Message_MQTTMultiError                       = config.Message_MQTTMultiError
	Message_MQTTValidationError                  = config.Message_MQTTValidationError
	Message_KafkaMultiError                      = config.Message_KafkaMultiError
	Message_KafkaValidationError                 = config.Message_KafkaValidationError
	Message_RabbitMQMultiError                   = config.Message_RabbitMQMultiError
	Message_RabbitMQValidationError              = config.Message_RabbitMQValidationError
	Message_ActiveMQMultiError                   = config.Message_ActiveMQMultiError
	Message_ActiveMQValidationError              = config.Message_ActiveMQValidationError
	Message_NATSMultiError                       = config.Message_NATSMultiError
	Message_NATSValidationError                  = config.Message_NATSValidationError
	Message_NSQMultiError                        = config.Message_NSQMultiError
	Message_NSQValidationError                   = config.Message_NSQValidationError
	Message_PulsarMultiError                     = config.Message_PulsarMultiError
	Message_PulsarValidationError                = config.Message_PulsarValidationError
	Message_RedisMultiError                      = config.Message_RedisMultiError
	Message_RedisValidationError                 = config.Message_RedisValidationError
	Message_RocketMQMultiError                   = config.Message_RocketMQMultiError
	Message_RocketMQValidationError              = config.Message_RocketMQValidationError
	MetricMultiError                             = config.MetricMultiError
	MetricValidationError                        = config.MetricValidationError
	EntrySelectorMultiError                      = config.EntrySelectorMultiError
	EntrySelectorValidationError                 = config.EntrySelectorValidationError
	ServiceMultiError                            = config.ServiceMultiError
	ServiceValidationError                       = config.ServiceValidationError
	Service_EntryMultiError                      = config.Service_EntryMultiError
	Service_EntryValidationError                 = config.Service_EntryValidationError
	Service_GINSMultiError                       = config.Service_GINSMultiError
	Service_GINSValidationError                  = config.Service_GINSValidationError
	Service_HTTPMultiError                       = config.Service_HTTPMultiError
	Service_HTTPValidationError                  = config.Service_HTTPValidationError
	Service_GRPCMultiError                       = config.Service_GRPCMultiError
	Service_GRPCValidationError                  = config.Service_GRPCValidationError
	Service_WebsocketMultiError                  = config.Service_WebsocketMultiError
	Service_WebsocketValidationError             = config.Service_WebsocketValidationError
	Service_MiddlewareMultiError                 = config.Service_MiddlewareMultiError
	Service_MiddlewareValidationError            = config.Service_MiddlewareValidationError
	Service_Middleware_MetricsMultiError         = config.Service_Middleware_MetricsMultiError
	Service_Middleware_MetricsValidationError    = config.Service_Middleware_MetricsValidationError
	Service_Middleware_TracesMultiError          = config.Service_Middleware_TracesMultiError
	Service_Middleware_TracesValidationError     = config.Service_Middleware_TracesValidationError
	Service_Middleware_LoggerMultiError          = config.Service_Middleware_LoggerMultiError
	Service_Middleware_LoggerValidationError     = config.Service_Middleware_LoggerValidationError
	Service_Middleware_CorsMultiError            = config.Service_Middleware_CorsMultiError
	Service_Middleware_CorsValidationError       = config.Service_Middleware_CorsValidationError
	RegistryMultiError                           = config.RegistryMultiError
	RegistryValidationError                      = config.RegistryValidationError
	Registry_ConsulMultiError                    = config.Registry_ConsulMultiError
	Registry_ConsulValidationError               = config.Registry_ConsulValidationError
	Registry_ETCDMultiError                      = config.Registry_ETCDMultiError
	Registry_ETCDValidationError                 = config.Registry_ETCDValidationError
	Registry_CustomMultiError                    = config.Registry_CustomMultiError
	Registry_CustomValidationError               = config.Registry_CustomValidationError
	AuthorizationMultiError                      = config.AuthorizationMultiError
	AuthorizationValidationError                 = config.AuthorizationValidationError
	CasbinMultiError                             = config.CasbinMultiError
	CasbinValidationError                        = config.CasbinValidationError
	SecurityMultiError                           = config.SecurityMultiError
	SecurityValidationError                      = config.SecurityValidationError
	MiddlewareMultiError                         = config.MiddlewareMultiError
	MiddlewareValidationError                    = config.MiddlewareValidationError
	Middleware_RateLimiterMultiError             = config.Middleware_RateLimiterMultiError
	Middleware_RateLimiterValidationError        = config.Middleware_RateLimiterValidationError
	Middleware_RateLimiter_RedisMultiError       = config.Middleware_RateLimiter_RedisMultiError
	Middleware_RateLimiter_RedisValidationError  = config.Middleware_RateLimiter_RedisValidationError
	Middleware_RateLimiter_MemoryMultiError      = config.Middleware_RateLimiter_MemoryMultiError
	Middleware_RateLimiter_MemoryValidationError = config.Middleware_RateLimiter_MemoryValidationError
)

// Define types from kratos config package
type (
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

// Define variables for proto files
var (
	Logger_LEVEL_UNSPECIFIED        = config.Logger_LEVEL_UNSPECIFIED
	Logger_LEVEL_DEBUG              = config.Logger_LEVEL_DEBUG
	Logger_LEVEL_INFO               = config.Logger_LEVEL_INFO
	Logger_LEVEL_WARN               = config.Logger_LEVEL_WARN
	Logger_LEVEL_ERROR              = config.Logger_LEVEL_ERROR
	Logger_LEVEL_FATAL              = config.Logger_LEVEL_FATAL
	File_config_v1_metrics_proto    = config.File_config_v1_metrics_proto
	File_config_v1_cors_proto       = config.File_config_v1_cors_proto
	File_config_v1_data_proto       = config.File_config_v1_data_proto
	File_config_v1_task_proto       = config.File_config_v1_task_proto
	Logger_Level_name               = config.Logger_Level_name
	Logger_Level_value              = config.Logger_Level_value
	File_config_v1_logger_proto     = config.File_config_v1_logger_proto
	File_config_v1_source_proto     = config.File_config_v1_source_proto
	File_config_v1_tracer_proto     = config.File_config_v1_tracer_proto
	File_config_v1_message_proto    = config.File_config_v1_message_proto
	File_config_v1_service_proto    = config.File_config_v1_service_proto
	File_config_v1_registry_proto   = config.File_config_v1_registry_proto
	File_config_v1_security_proto   = config.File_config_v1_security_proto
	File_config_v1_middleware_proto = config.File_config_v1_middleware_proto
)

var (
	// ErrNotFound defined error from kratos config package
	ErrNotFound = kratosconfig.ErrNotFound
)

// New returns a new config instance
func New(opts ...Option) Config {
	return kratosconfig.New(opts...)
}

// WithDecoder sets the decoder
func WithDecoder(d Decoder) Option {
	return kratosconfig.WithDecoder(d)
}

// WithMergeFunc sets the merge function
func WithMergeFunc(m Merge) Option {
	return kratosconfig.WithMergeFunc(m)
}

// WithResolveActualTypes enables resolving actual types
func WithResolveActualTypes(enableConvertToType bool) Option {
	return kratosconfig.WithResolveActualTypes(enableConvertToType)
}

// WithResolver sets the resolver
func WithResolver(r Resolver) Option {
	return kratosconfig.WithResolver(r)
}

// WithSource sets the source
func WithSource(s ...Source) Option {
	return kratosconfig.WithSource(s...)
}
