// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: config/v1/data.proto

package configv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Data
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Database
	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// Cache
	Cache *Data_Cache `protobuf:"bytes,2,opt,name=cache,proto3" json:"cache,omitempty"`
	// Storage
	Storage *Data_Storage `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_config_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetCache() *Data_Cache {
	if x != nil {
		return x.Cache
	}
	return nil
}

func (x *Data) GetStorage() *Data_Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

// Database
type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Debugging
	Debug bool `protobuf:"varint,1,opt,name=debug,proto3" json:"debug,omitempty"`
	// Driver name: mysql, postgresql, mongodb, sqlite......
	Driver string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	// Data source (DSN string)
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// Data migration switch
	Migrate bool `protobuf:"varint,10,opt,name=migrate,proto3" json:"migrate,omitempty"`
	// Link tracking switch
	EnableTrace bool `protobuf:"varint,12,opt,name=enable_trace,proto3" json:"enable_trace,omitempty"`
	// Performance analysis switch
	EnableMetrics bool `protobuf:"varint,13,opt,name=enable_metrics,proto3" json:"enable_metrics,omitempty"`
	// Maximum number of free connections in the connection pool
	MaxIdleConnections int32 `protobuf:"varint,20,opt,name=max_idle_connections,proto3" json:"max_idle_connections,omitempty"`
	// Maximum number of open connections in the connection pool
	MaxOpenConnections int32 `protobuf:"varint,21,opt,name=max_open_connections,proto3" json:"max_open_connections,omitempty"`
	// Maximum length of time that the connection can be reused
	ConnectionMaxLifetime *durationpb.Duration `protobuf:"bytes,22,opt,name=connection_max_lifetime,proto3" json:"connection_max_lifetime,omitempty"`
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	mi := &file_config_v1_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Data_Database) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Data_Database) GetMigrate() bool {
	if x != nil {
		return x.Migrate
	}
	return false
}

func (x *Data_Database) GetEnableTrace() bool {
	if x != nil {
		return x.EnableTrace
	}
	return false
}

func (x *Data_Database) GetEnableMetrics() bool {
	if x != nil {
		return x.EnableMetrics
	}
	return false
}

func (x *Data_Database) GetMaxIdleConnections() int32 {
	if x != nil {
		return x.MaxIdleConnections
	}
	return 0
}

func (x *Data_Database) GetMaxOpenConnections() int32 {
	if x != nil {
		return x.MaxOpenConnections
	}
	return 0
}

func (x *Data_Database) GetConnectionMaxLifetime() *durationpb.Duration {
	if x != nil {
		return x.ConnectionMaxLifetime
	}
	return nil
}

// Redis
type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Password     string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Db           int32                `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`
	DialTimeout  *durationpb.Duration `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,6,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	mi := &file_config_v1_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Data_Redis) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

// Memcached
type Data_Memcached struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr     string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Username string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	MaxIdle  int32                `protobuf:"varint,4,opt,name=max_idle,proto3" json:"max_idle,omitempty"`
	Timeout  *durationpb.Duration `protobuf:"bytes,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Data_Memcached) Reset() {
	*x = Data_Memcached{}
	mi := &file_config_v1_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Memcached) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Memcached) ProtoMessage() {}

func (x *Data_Memcached) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Memcached.ProtoReflect.Descriptor instead.
func (*Data_Memcached) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Data_Memcached) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Memcached) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Data_Memcached) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Memcached) GetMaxIdle() int32 {
	if x != nil {
		return x.MaxIdle
	}
	return 0
}

func (x *Data_Memcached) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// Memory
type Data_Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size            int32                `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Capacity        int32                `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Expiration      *durationpb.Duration `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	CleanupInterval *durationpb.Duration `protobuf:"bytes,4,opt,name=cleanup_interval,proto3" json:"cleanup_interval,omitempty"`
}

func (x *Data_Memory) Reset() {
	*x = Data_Memory{}
	mi := &file_config_v1_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Memory) ProtoMessage() {}

func (x *Data_Memory) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Memory.ProtoReflect.Descriptor instead.
func (*Data_Memory) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Data_Memory) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Data_Memory) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Data_Memory) GetExpiration() *durationpb.Duration {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *Data_Memory) GetCleanupInterval() *durationpb.Duration {
	if x != nil {
		return x.CleanupInterval
	}
	return nil
}

// File
type Data_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Root string `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *Data_File) Reset() {
	*x = Data_File{}
	mi := &file_config_v1_data_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_File) ProtoMessage() {}

func (x *Data_File) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_File.ProtoReflect.Descriptor instead.
func (*Data_File) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Data_File) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

// OSS
type Data_Oss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Data_Oss) Reset() {
	*x = Data_Oss{}
	mi := &file_config_v1_data_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Oss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Oss) ProtoMessage() {}

func (x *Data_Oss) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Oss.ProtoReflect.Descriptor instead.
func (*Data_Oss) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 5}
}

// Mongo
type Data_Mongo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Data_Mongo) Reset() {
	*x = Data_Mongo{}
	mi := &file_config_v1_data_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Mongo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Mongo) ProtoMessage() {}

func (x *Data_Mongo) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Mongo.ProtoReflect.Descriptor instead.
func (*Data_Mongo) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 6}
}

// Storage
type Data_Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string      `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	File  *Data_File  `protobuf:"bytes,10,opt,name=file,proto3" json:"file,omitempty"`
	Redis *Data_Redis `protobuf:"bytes,11,opt,name=redis,proto3" json:"redis,omitempty"`
	Mongo *Data_Mongo `protobuf:"bytes,12,opt,name=mongo,proto3" json:"mongo,omitempty"`
	Oss   *Data_Oss   `protobuf:"bytes,13,opt,name=oss,proto3" json:"oss,omitempty"`
}

func (x *Data_Storage) Reset() {
	*x = Data_Storage{}
	mi := &file_config_v1_data_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Storage) ProtoMessage() {}

func (x *Data_Storage) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Storage.ProtoReflect.Descriptor instead.
func (*Data_Storage) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 7}
}

func (x *Data_Storage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Data_Storage) GetFile() *Data_File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Data_Storage) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Data_Storage) GetMongo() *Data_Mongo {
	if x != nil {
		return x.Mongo
	}
	return nil
}

func (x *Data_Storage) GetOss() *Data_Oss {
	if x != nil {
		return x.Oss
	}
	return nil
}

// Cache
type Data_Cache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Driver name: redis, memcached, etc.
	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	// Redis
	Redis *Data_Redis `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	// Memcached
	Memcached *Data_Memcached `protobuf:"bytes,3,opt,name=memcached,proto3" json:"memcached,omitempty"`
	// Memory cache
	Memory *Data_Memory `protobuf:"bytes,4,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *Data_Cache) Reset() {
	*x = Data_Cache{}
	mi := &file_config_v1_data_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Cache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Cache) ProtoMessage() {}

func (x *Data_Cache) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_data_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Cache.ProtoReflect.Descriptor instead.
func (*Data_Cache) Descriptor() ([]byte, []int) {
	return file_config_v1_data_proto_rawDescGZIP(), []int{0, 8}
}

func (x *Data_Cache) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Cache) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Data_Cache) GetMemcached() *Data_Memcached {
	if x != nil {
		return x.Memcached
	}
	return nil
}

func (x *Data_Cache) GetMemory() *Data_Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

var File_config_v1_data_proto protoreflect.FileDescriptor

var file_config_v1_data_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde,
	0x0d, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0xc1, 0x03,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x12, 0x64, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x4c, 0xba, 0x48, 0x49, 0x72, 0x47, 0x52, 0x05, 0x6d, 0x73, 0x73, 0x71, 0x6c, 0x52, 0x05,
	0x6d, 0x79, 0x73, 0x71, 0x6c, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71,
	0x6c, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x52, 0x06, 0x73, 0x71, 0x6c, 0x69,
	0x74, 0x65, 0x52, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x71, 0x6c, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x33, 0x52, 0x06,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f,
	0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x53, 0x0a, 0x17,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c,
	0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x1a, 0x9d, 0x02, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x64, 0x62, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x1a, 0xa8, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d,
	0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0xba, 0x01, 0x0a,
	0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x1a, 0x1a, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6f, 0x74, 0x1a, 0x05, 0x0a, 0x03, 0x4f, 0x73, 0x73, 0x1a, 0x07, 0x0a, 0x05,
	0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x1a, 0xee, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x24, 0xba, 0x48, 0x21, 0x72, 0x1f, 0x52, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x52, 0x03, 0x6f, 0x73, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x12,
	0x25, 0x0a, 0x03, 0x6f, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x73,
	0x73, 0x52, 0x03, 0x6f, 0x73, 0x73, 0x1a, 0xd6, 0x01, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x37, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1f, 0xba, 0x48, 0x1c, 0x72, 0x1a, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x52, 0x09,
	0x6d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52,
	0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x6d, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12,
	0x2e, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42,
	0xa3, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x73, 0x2f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0xf8,
	0x01, 0x01, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_v1_data_proto_rawDescOnce sync.Once
	file_config_v1_data_proto_rawDescData = file_config_v1_data_proto_rawDesc
)

func file_config_v1_data_proto_rawDescGZIP() []byte {
	file_config_v1_data_proto_rawDescOnce.Do(func() {
		file_config_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1_data_proto_rawDescData)
	})
	return file_config_v1_data_proto_rawDescData
}

var file_config_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_config_v1_data_proto_goTypes = []any{
	(*Data)(nil),                // 0: config.v1.Data
	(*Data_Database)(nil),       // 1: config.v1.Data.Database
	(*Data_Redis)(nil),          // 2: config.v1.Data.Redis
	(*Data_Memcached)(nil),      // 3: config.v1.Data.Memcached
	(*Data_Memory)(nil),         // 4: config.v1.Data.Memory
	(*Data_File)(nil),           // 5: config.v1.Data.File
	(*Data_Oss)(nil),            // 6: config.v1.Data.Oss
	(*Data_Mongo)(nil),          // 7: config.v1.Data.Mongo
	(*Data_Storage)(nil),        // 8: config.v1.Data.Storage
	(*Data_Cache)(nil),          // 9: config.v1.Data.Cache
	(*durationpb.Duration)(nil), // 10: google.protobuf.Duration
}
var file_config_v1_data_proto_depIdxs = []int32{
	1,  // 0: config.v1.Data.database:type_name -> config.v1.Data.Database
	9,  // 1: config.v1.Data.cache:type_name -> config.v1.Data.Cache
	8,  // 2: config.v1.Data.storage:type_name -> config.v1.Data.Storage
	10, // 3: config.v1.Data.Database.connection_max_lifetime:type_name -> google.protobuf.Duration
	10, // 4: config.v1.Data.Redis.dial_timeout:type_name -> google.protobuf.Duration
	10, // 5: config.v1.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	10, // 6: config.v1.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	10, // 7: config.v1.Data.Memcached.timeout:type_name -> google.protobuf.Duration
	10, // 8: config.v1.Data.Memory.expiration:type_name -> google.protobuf.Duration
	10, // 9: config.v1.Data.Memory.cleanup_interval:type_name -> google.protobuf.Duration
	5,  // 10: config.v1.Data.Storage.file:type_name -> config.v1.Data.File
	2,  // 11: config.v1.Data.Storage.redis:type_name -> config.v1.Data.Redis
	7,  // 12: config.v1.Data.Storage.mongo:type_name -> config.v1.Data.Mongo
	6,  // 13: config.v1.Data.Storage.oss:type_name -> config.v1.Data.Oss
	2,  // 14: config.v1.Data.Cache.redis:type_name -> config.v1.Data.Redis
	3,  // 15: config.v1.Data.Cache.memcached:type_name -> config.v1.Data.Memcached
	4,  // 16: config.v1.Data.Cache.memory:type_name -> config.v1.Data.Memory
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_config_v1_data_proto_init() }
func file_config_v1_data_proto_init() {
	if File_config_v1_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_v1_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_v1_data_proto_goTypes,
		DependencyIndexes: file_config_v1_data_proto_depIdxs,
		MessageInfos:      file_config_v1_data_proto_msgTypes,
	}.Build()
	File_config_v1_data_proto = out.File
	file_config_v1_data_proto_rawDesc = nil
	file_config_v1_data_proto_goTypes = nil
	file_config_v1_data_proto_depIdxs = nil
}
