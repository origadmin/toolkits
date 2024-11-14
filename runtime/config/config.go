package config

import (
	kratosconfig "github.com/go-kratos/kratos/v2/config"

	"github.com/origadmin/toolkits/runtime/internal/config/v1"
)

// Import the kratos config package and the internal config package
type (
	Data                                         = config.Data
	Data_Database                                = config.Data_Database
	Data_Storage                                 = config.Data_Storage
	Data_Storage_File                            = config.Data_Storage_File
	Data_Storage_Redis                           = config.Data_Storage_Redis
	Data_Storage_Mongo                           = config.Data_Storage_Mongo
	Data_Storage_Oss                             = config.Data_Storage_Oss
	Cors                                         = config.Cors
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
	Service                                      = config.Service
	Service_GINS                                 = config.Service_GINS
	Service_HTTP                                 = config.Service_HTTP
	Service_GRPC                                 = config.Service_GRPC
	Registry                                     = config.Registry
	Registry_Consul                              = config.Registry_Consul
	Registry_ETCD                                = config.Registry_ETCD
	Registry_Custom                              = config.Registry_Custom
	WebSocket                                    = config.WebSocket
	UserMetric_MetricType                        = config.UserMetric_MetricType
	UserMetric                                   = config.UserMetric
	Security                                     = config.Security
	Middleware                                   = config.Middleware
	Security_Casbin                              = config.Security_Casbin
	Security_JWT                                 = config.Security_JWT
	Middleware_RateLimiter                       = config.Middleware_RateLimiter
	Middleware_Metrics                           = config.Middleware_Metrics
	Middleware_Metadata                          = config.Middleware_Metadata
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
	Data_Storage_OssMultiError                   = config.Data_Storage_OssMultiError
	Data_Storage_OssValidationError              = config.Data_Storage_OssValidationError
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
	ServiceMultiError                            = config.ServiceMultiError
	ServiceValidationError                       = config.ServiceValidationError
	Service_GINSMultiError                       = config.Service_GINSMultiError
	Service_GINSValidationError                  = config.Service_GINSValidationError
	Service_HTTPMultiError                       = config.Service_HTTPMultiError
	Service_HTTPValidationError                  = config.Service_HTTPValidationError
	Service_GRPCMultiError                       = config.Service_GRPCMultiError
	Service_GRPCValidationError                  = config.Service_GRPCValidationError
	Service_Selector                             = config.Service_Selector
	RegistryMultiError                           = config.RegistryMultiError
	RegistryValidationError                      = config.RegistryValidationError
	Registry_ConsulMultiError                    = config.Registry_ConsulMultiError
	Registry_ConsulValidationError               = config.Registry_ConsulValidationError
	Registry_ETCDMultiError                      = config.Registry_ETCDMultiError
	Registry_ETCDValidationError                 = config.Registry_ETCDValidationError
	Registry_CustomMultiError                    = config.Registry_CustomMultiError
	Registry_CustomValidationError               = config.Registry_CustomValidationError
	WebSocketMultiError                          = config.WebSocketMultiError
	WebSocketValidationError                     = config.WebSocketValidationError
	UserMetricMultiError                         = config.UserMetricMultiError
	UserMetricValidationError                    = config.UserMetricValidationError
	SecurityMultiError                           = config.SecurityMultiError
	SecurityValidationError                      = config.SecurityValidationError
	MiddlewareMultiError                         = config.MiddlewareMultiError
	MiddlewareValidationError                    = config.MiddlewareValidationError
	Security_CasbinMultiError                    = config.Security_CasbinMultiError
	Security_CasbinValidationError               = config.Security_CasbinValidationError
	Security_JWTMultiError                       = config.Security_JWTMultiError
	Security_JWTValidationError                  = config.Security_JWTValidationError
	Middleware_RateLimiterMultiError             = config.Middleware_RateLimiterMultiError
	Middleware_RateLimiterValidationError        = config.Middleware_RateLimiterValidationError
	Middleware_MetricsMultiError                 = config.Middleware_MetricsMultiError
	Middleware_MetricsValidationError            = config.Middleware_MetricsValidationError
	Middleware_MetadataMultiError                = config.Middleware_MetadataMultiError
	Middleware_MetadataValidationError           = config.Middleware_MetadataValidationError
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
	Logger_LEVEL_UNSPECIFIED = config.Logger_LEVEL_UNSPECIFIED
	Logger_LEVEL_DEBUG       = config.Logger_LEVEL_DEBUG
	Logger_LEVEL_INFO        = config.Logger_LEVEL_INFO
	Logger_LEVEL_WARN        = config.Logger_LEVEL_WARN
	Logger_LEVEL_ERROR       = config.Logger_LEVEL_ERROR
	Logger_LEVEL_FATAL       = config.Logger_LEVEL_FATAL
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
