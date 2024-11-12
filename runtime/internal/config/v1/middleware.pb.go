// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: config/v1/middleware.proto

package config

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

type Middleware struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout *durationpb.Duration `protobuf:"bytes,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Middleware) Reset() {
	*x = Middleware{}
	mi := &file_config_v1_middleware_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware) ProtoMessage() {}

func (x *Middleware) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_middleware_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware.ProtoReflect.Descriptor instead.
func (*Middleware) Descriptor() ([]byte, []int) {
	return file_config_v1_middleware_proto_rawDescGZIP(), []int{0}
}

func (x *Middleware) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// Rate limiter
type Middleware_RateLimiter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // rate limiter name, supported: bbr.
	Period int64  `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// The number of requests allowed in a window of time
	XRatelimitLimit int64 `protobuf:"varint,5,opt,name=x_ratelimit_limit,proto3" json:"x_ratelimit_limit,omitempty"`
	// The number of requests that can still be made in the current window of time
	XRatelimitRemaining int64 `protobuf:"varint,6,opt,name=x_ratelimit_remaining,proto3" json:"x_ratelimit_remaining,omitempty"`
	// The number of seconds until the current rate limit window completely resets
	XRatelimitReset int64 `protobuf:"varint,7,opt,name=x_ratelimit_reset,proto3" json:"x_ratelimit_reset,omitempty"`
	// When rate limited, the number of seconds to wait before another request will be accepted
	RetryAfter int64 `protobuf:"varint,8,opt,name=retry_after,proto3" json:"retry_after,omitempty"`
	// memory/redis
	StoreType string                         `protobuf:"bytes,100,opt,name=store_type,proto3" json:"store_type,omitempty"`
	Memory    *Middleware_RateLimiter_Memory `protobuf:"bytes,101,opt,name=memory,proto3" json:"memory,omitempty"`
	Redis     *Middleware_RateLimiter_Redis  `protobuf:"bytes,102,opt,name=redis,proto3" json:"redis,omitempty"`
}

func (x *Middleware_RateLimiter) Reset() {
	*x = Middleware_RateLimiter{}
	mi := &file_config_v1_middleware_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_RateLimiter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_RateLimiter) ProtoMessage() {}

func (x *Middleware_RateLimiter) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_middleware_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_RateLimiter.ProtoReflect.Descriptor instead.
func (*Middleware_RateLimiter) Descriptor() ([]byte, []int) {
	return file_config_v1_middleware_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Middleware_RateLimiter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Middleware_RateLimiter) GetPeriod() int64 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *Middleware_RateLimiter) GetXRatelimitLimit() int64 {
	if x != nil {
		return x.XRatelimitLimit
	}
	return 0
}

func (x *Middleware_RateLimiter) GetXRatelimitRemaining() int64 {
	if x != nil {
		return x.XRatelimitRemaining
	}
	return 0
}

func (x *Middleware_RateLimiter) GetXRatelimitReset() int64 {
	if x != nil {
		return x.XRatelimitReset
	}
	return 0
}

func (x *Middleware_RateLimiter) GetRetryAfter() int64 {
	if x != nil {
		return x.RetryAfter
	}
	return 0
}

func (x *Middleware_RateLimiter) GetStoreType() string {
	if x != nil {
		return x.StoreType
	}
	return ""
}

func (x *Middleware_RateLimiter) GetMemory() *Middleware_RateLimiter_Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *Middleware_RateLimiter) GetRedis() *Middleware_RateLimiter_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type Middleware_RateLimiter_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr     string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Db       int64  `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`
}

func (x *Middleware_RateLimiter_Redis) Reset() {
	*x = Middleware_RateLimiter_Redis{}
	mi := &file_config_v1_middleware_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_RateLimiter_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_RateLimiter_Redis) ProtoMessage() {}

func (x *Middleware_RateLimiter_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_middleware_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_RateLimiter_Redis.ProtoReflect.Descriptor instead.
func (*Middleware_RateLimiter_Redis) Descriptor() ([]byte, []int) {
	return file_config_v1_middleware_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Middleware_RateLimiter_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Middleware_RateLimiter_Redis) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Middleware_RateLimiter_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Middleware_RateLimiter_Redis) GetDb() int64 {
	if x != nil {
		return x.Db
	}
	return 0
}

type Middleware_RateLimiter_Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expiration      *durationpb.Duration `protobuf:"bytes,1,opt,name=expiration,proto3" json:"expiration,omitempty"`
	CleanupInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=cleanup_interval,proto3" json:"cleanup_interval,omitempty"`
}

func (x *Middleware_RateLimiter_Memory) Reset() {
	*x = Middleware_RateLimiter_Memory{}
	mi := &file_config_v1_middleware_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_RateLimiter_Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_RateLimiter_Memory) ProtoMessage() {}

func (x *Middleware_RateLimiter_Memory) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_middleware_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_RateLimiter_Memory.ProtoReflect.Descriptor instead.
func (*Middleware_RateLimiter_Memory) Descriptor() ([]byte, []int) {
	return file_config_v1_middleware_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Middleware_RateLimiter_Memory) GetExpiration() *durationpb.Duration {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *Middleware_RateLimiter_Memory) GetCleanupInterval() *durationpb.Duration {
	if x != nil {
		return x.CleanupInterval
	}
	return nil
}

var File_config_v1_middleware_proto protoreflect.FileDescriptor

var file_config_v1_middleware_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x05, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x96, 0x05, 0x0a, 0x0b, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x15, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x11, 0x78, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xba, 0x48,
	0x11, 0x72, 0x0f, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x12, 0x3d, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a,
	0x63, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x64, 0x62, 0x1a, 0x8a, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x10, 0x63, 0x6c,
	0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x42, 0x9d, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x42, 0x0f, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x6f, 0x6f,
	0x6c, 0x6b, 0x69, 0x74, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_v1_middleware_proto_rawDescOnce sync.Once
	file_config_v1_middleware_proto_rawDescData = file_config_v1_middleware_proto_rawDesc
)

func file_config_v1_middleware_proto_rawDescGZIP() []byte {
	file_config_v1_middleware_proto_rawDescOnce.Do(func() {
		file_config_v1_middleware_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1_middleware_proto_rawDescData)
	})
	return file_config_v1_middleware_proto_rawDescData
}

var file_config_v1_middleware_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_v1_middleware_proto_goTypes = []any{
	(*Middleware)(nil),                    // 0: config.v1.Middleware
	(*Middleware_RateLimiter)(nil),        // 1: config.v1.Middleware.RateLimiter
	(*Middleware_RateLimiter_Redis)(nil),  // 2: config.v1.Middleware.RateLimiter.Redis
	(*Middleware_RateLimiter_Memory)(nil), // 3: config.v1.Middleware.RateLimiter.Memory
	(*durationpb.Duration)(nil),           // 4: google.protobuf.Duration
}
var file_config_v1_middleware_proto_depIdxs = []int32{
	4, // 0: config.v1.Middleware.timeout:type_name -> google.protobuf.Duration
	3, // 1: config.v1.Middleware.RateLimiter.memory:type_name -> config.v1.Middleware.RateLimiter.Memory
	2, // 2: config.v1.Middleware.RateLimiter.redis:type_name -> config.v1.Middleware.RateLimiter.Redis
	4, // 3: config.v1.Middleware.RateLimiter.Memory.expiration:type_name -> google.protobuf.Duration
	4, // 4: config.v1.Middleware.RateLimiter.Memory.cleanup_interval:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_v1_middleware_proto_init() }
func file_config_v1_middleware_proto_init() {
	if File_config_v1_middleware_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_v1_middleware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_v1_middleware_proto_goTypes,
		DependencyIndexes: file_config_v1_middleware_proto_depIdxs,
		MessageInfos:      file_config_v1_middleware_proto_msgTypes,
	}.Build()
	File_config_v1_middleware_proto = out.File
	file_config_v1_middleware_proto_rawDesc = nil
	file_config_v1_middleware_proto_goTypes = nil
	file_config_v1_middleware_proto_depIdxs = nil
}
