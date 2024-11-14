// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: config/v1/service.proto

package config

import (
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

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service name for service discovery
	Name         string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AutoEndpoint bool              `protobuf:"varint,2,opt,name=auto_endpoint,proto3" json:"auto_endpoint,omitempty"`
	Grpc         *Service_GRPC     `protobuf:"bytes,10,opt,name=grpc,proto3" json:"grpc,omitempty"`
	Http         *Service_HTTP     `protobuf:"bytes,20,opt,name=http,proto3" json:"http,omitempty"`
	Gins         *Service_GINS     `protobuf:"bytes,30,opt,name=gins,proto3" json:"gins,omitempty"`
	Websocket    *WebSocket        `protobuf:"bytes,100,opt,name=websocket,proto3" json:"websocket,omitempty"`
	Message      *Message          `protobuf:"bytes,200,opt,name=message,proto3" json:"message,omitempty"`
	Task         *Task             `protobuf:"bytes,300,opt,name=task,proto3" json:"task,omitempty"`
	Middleware   *Middleware       `protobuf:"bytes,400,opt,name=middleware,proto3" json:"middleware,omitempty"`
	Selector     *Service_Selector `protobuf:"bytes,500,opt,name=selector,proto3" json:"selector,omitempty"`
	Host         string            `protobuf:"bytes,99,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	mi := &file_config_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_config_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetAutoEndpoint() bool {
	if x != nil {
		return x.AutoEndpoint
	}
	return false
}

func (x *Service) GetGrpc() *Service_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Service) GetHttp() *Service_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Service) GetGins() *Service_GINS {
	if x != nil {
		return x.Gins
	}
	return nil
}

func (x *Service) GetWebsocket() *WebSocket {
	if x != nil {
		return x.Websocket
	}
	return nil
}

func (x *Service) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Service) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *Service) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

func (x *Service) GetSelector() *Service_Selector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *Service) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

// GINS
type Service_GINS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network         string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr            string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	UseTls          bool                 `protobuf:"varint,3,opt,name=use_tls,proto3" json:"use_tls,omitempty"`
	CertFile        string               `protobuf:"bytes,4,opt,name=cert_file,proto3" json:"cert_file,omitempty"`
	KeyFile         string               `protobuf:"bytes,5,opt,name=key_file,proto3" json:"key_file,omitempty"`
	Timeout         *durationpb.Duration `protobuf:"bytes,6,opt,name=timeout,proto3,oneof" json:"timeout,omitempty"`
	ShutdownTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=shutdown_timeout,proto3,oneof" json:"shutdown_timeout,omitempty"`
	ReadTimeout     *durationpb.Duration `protobuf:"bytes,8,opt,name=read_timeout,proto3,oneof" json:"read_timeout,omitempty"`
	WriteTimeout    *durationpb.Duration `protobuf:"bytes,9,opt,name=write_timeout,proto3,oneof" json:"write_timeout,omitempty"`
	IdleTimeout     *durationpb.Duration `protobuf:"bytes,10,opt,name=idle_timeout,proto3,oneof" json:"idle_timeout,omitempty"`
	Endpoint        string               `protobuf:"bytes,11,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Service_GINS) Reset() {
	*x = Service_GINS{}
	mi := &file_config_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service_GINS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service_GINS) ProtoMessage() {}

func (x *Service_GINS) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service_GINS.ProtoReflect.Descriptor instead.
func (*Service_GINS) Descriptor() ([]byte, []int) {
	return file_config_v1_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Service_GINS) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Service_GINS) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Service_GINS) GetUseTls() bool {
	if x != nil {
		return x.UseTls
	}
	return false
}

func (x *Service_GINS) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *Service_GINS) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

func (x *Service_GINS) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Service_GINS) GetShutdownTimeout() *durationpb.Duration {
	if x != nil {
		return x.ShutdownTimeout
	}
	return nil
}

func (x *Service_GINS) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Service_GINS) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Service_GINS) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *Service_GINS) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

// HTTP
type Service_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network         string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr            string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	UseTls          bool                 `protobuf:"varint,3,opt,name=use_tls,proto3" json:"use_tls,omitempty"`
	CertFile        string               `protobuf:"bytes,4,opt,name=cert_file,proto3" json:"cert_file,omitempty"`
	KeyFile         string               `protobuf:"bytes,5,opt,name=key_file,proto3" json:"key_file,omitempty"`
	Timeout         *durationpb.Duration `protobuf:"bytes,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	ShutdownTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=shutdown_timeout,proto3" json:"shutdown_timeout,omitempty"`
	ReadTimeout     *durationpb.Duration `protobuf:"bytes,8,opt,name=read_timeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout    *durationpb.Duration `protobuf:"bytes,9,opt,name=write_timeout,proto3" json:"write_timeout,omitempty"`
	IdleTimeout     *durationpb.Duration `protobuf:"bytes,10,opt,name=idle_timeout,proto3" json:"idle_timeout,omitempty"`
	Endpoint        string               `protobuf:"bytes,11,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Service_HTTP) Reset() {
	*x = Service_HTTP{}
	mi := &file_config_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service_HTTP) ProtoMessage() {}

func (x *Service_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service_HTTP.ProtoReflect.Descriptor instead.
func (*Service_HTTP) Descriptor() ([]byte, []int) {
	return file_config_v1_service_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Service_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Service_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Service_HTTP) GetUseTls() bool {
	if x != nil {
		return x.UseTls
	}
	return false
}

func (x *Service_HTTP) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *Service_HTTP) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

func (x *Service_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Service_HTTP) GetShutdownTimeout() *durationpb.Duration {
	if x != nil {
		return x.ShutdownTimeout
	}
	return nil
}

func (x *Service_HTTP) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Service_HTTP) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Service_HTTP) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *Service_HTTP) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

// GRPC
type Service_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network         string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr            string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	UseTls          bool                 `protobuf:"varint,3,opt,name=use_tls,proto3" json:"use_tls,omitempty"`
	CertFile        string               `protobuf:"bytes,4,opt,name=cert_file,proto3" json:"cert_file,omitempty"`
	KeyFile         string               `protobuf:"bytes,5,opt,name=key_file,proto3" json:"key_file,omitempty"`
	Timeout         *durationpb.Duration `protobuf:"bytes,6,opt,name=timeout,proto3,oneof" json:"timeout,omitempty"`
	ShutdownTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=shutdown_timeout,proto3,oneof" json:"shutdown_timeout,omitempty"`
	ReadTimeout     *durationpb.Duration `protobuf:"bytes,8,opt,name=read_timeout,proto3,oneof" json:"read_timeout,omitempty"`
	WriteTimeout    *durationpb.Duration `protobuf:"bytes,9,opt,name=write_timeout,proto3,oneof" json:"write_timeout,omitempty"`
	IdleTimeout     *durationpb.Duration `protobuf:"bytes,10,opt,name=idle_timeout,proto3,oneof" json:"idle_timeout,omitempty"`
	Endpoint        string               `protobuf:"bytes,11,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Service_GRPC) Reset() {
	*x = Service_GRPC{}
	mi := &file_config_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service_GRPC) ProtoMessage() {}

func (x *Service_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service_GRPC.ProtoReflect.Descriptor instead.
func (*Service_GRPC) Descriptor() ([]byte, []int) {
	return file_config_v1_service_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Service_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Service_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Service_GRPC) GetUseTls() bool {
	if x != nil {
		return x.UseTls
	}
	return false
}

func (x *Service_GRPC) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *Service_GRPC) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

func (x *Service_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Service_GRPC) GetShutdownTimeout() *durationpb.Duration {
	if x != nil {
		return x.ShutdownTimeout
	}
	return nil
}

func (x *Service_GRPC) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Service_GRPC) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Service_GRPC) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *Service_GRPC) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

// Selector
type Service_Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Builder string `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
}

func (x *Service_Selector) Reset() {
	*x = Service_Selector{}
	mi := &file_config_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service_Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service_Selector) ProtoMessage() {}

func (x *Service_Selector) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service_Selector.ProtoReflect.Descriptor instead.
func (*Service_Selector) Descriptor() ([]byte, []int) {
	return file_config_v1_service_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Service_Selector) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Service_Selector) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

var File_config_v1_service_proto protoreflect.FileDescriptor

var file_config_v1_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x11, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75,
	0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x2b, 0x0a,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x67, 0x69,
	0x6e, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x49, 0x4e,
	0x53, 0x52, 0x04, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x77, 0x65, 0x62, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x09, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0x36, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x18, 0x90,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0xcd, 0x04, 0x0a, 0x04, 0x47, 0x49, 0x4e, 0x53, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x72, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x38, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a, 0x10, 0x73,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x01, 0x52, 0x10, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52,
	0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x42, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x04, 0x52, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x13, 0x0a,
	0x11, 0x5f, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0xdf, 0x03, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x72, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f,
	0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x68, 0x75,
	0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d, 0x0a,
	0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3f, 0x0a, 0x0d,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d, 0x0a,
	0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0xcd, 0x04, 0x0a, 0x04, 0x47, 0x52, 0x50,
	0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x65, 0x72,
	0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65,
	0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a,
	0x10, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x01, 0x52, 0x10, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x0c, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a,
	0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x03, 0x52, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x04, 0x52, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42,
	0x13, 0x0a, 0x11, 0x5f, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x64, 0x6c, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x3e, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x42, 0x9a, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xf8,
	0x01, 0x01, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_v1_service_proto_rawDescOnce sync.Once
	file_config_v1_service_proto_rawDescData = file_config_v1_service_proto_rawDesc
)

func file_config_v1_service_proto_rawDescGZIP() []byte {
	file_config_v1_service_proto_rawDescOnce.Do(func() {
		file_config_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1_service_proto_rawDescData)
	})
	return file_config_v1_service_proto_rawDescData
}

var file_config_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_config_v1_service_proto_goTypes = []any{
	(*Service)(nil),             // 0: config.v1.Service
	(*Service_GINS)(nil),        // 1: config.v1.Service.GINS
	(*Service_HTTP)(nil),        // 2: config.v1.Service.HTTP
	(*Service_GRPC)(nil),        // 3: config.v1.Service.GRPC
	(*Service_Selector)(nil),    // 4: config.v1.Service.Selector
	(*WebSocket)(nil),           // 5: config.v1.WebSocket
	(*Message)(nil),             // 6: config.v1.Message
	(*Task)(nil),                // 7: config.v1.Task
	(*Middleware)(nil),          // 8: config.v1.Middleware
	(*durationpb.Duration)(nil), // 9: google.protobuf.Duration
}
var file_config_v1_service_proto_depIdxs = []int32{
	3,  // 0: config.v1.Service.grpc:type_name -> config.v1.Service.GRPC
	2,  // 1: config.v1.Service.http:type_name -> config.v1.Service.HTTP
	1,  // 2: config.v1.Service.gins:type_name -> config.v1.Service.GINS
	5,  // 3: config.v1.Service.websocket:type_name -> config.v1.WebSocket
	6,  // 4: config.v1.Service.message:type_name -> config.v1.Message
	7,  // 5: config.v1.Service.task:type_name -> config.v1.Task
	8,  // 6: config.v1.Service.middleware:type_name -> config.v1.Middleware
	4,  // 7: config.v1.Service.selector:type_name -> config.v1.Service.Selector
	9,  // 8: config.v1.Service.GINS.timeout:type_name -> google.protobuf.Duration
	9,  // 9: config.v1.Service.GINS.shutdown_timeout:type_name -> google.protobuf.Duration
	9,  // 10: config.v1.Service.GINS.read_timeout:type_name -> google.protobuf.Duration
	9,  // 11: config.v1.Service.GINS.write_timeout:type_name -> google.protobuf.Duration
	9,  // 12: config.v1.Service.GINS.idle_timeout:type_name -> google.protobuf.Duration
	9,  // 13: config.v1.Service.HTTP.timeout:type_name -> google.protobuf.Duration
	9,  // 14: config.v1.Service.HTTP.shutdown_timeout:type_name -> google.protobuf.Duration
	9,  // 15: config.v1.Service.HTTP.read_timeout:type_name -> google.protobuf.Duration
	9,  // 16: config.v1.Service.HTTP.write_timeout:type_name -> google.protobuf.Duration
	9,  // 17: config.v1.Service.HTTP.idle_timeout:type_name -> google.protobuf.Duration
	9,  // 18: config.v1.Service.GRPC.timeout:type_name -> google.protobuf.Duration
	9,  // 19: config.v1.Service.GRPC.shutdown_timeout:type_name -> google.protobuf.Duration
	9,  // 20: config.v1.Service.GRPC.read_timeout:type_name -> google.protobuf.Duration
	9,  // 21: config.v1.Service.GRPC.write_timeout:type_name -> google.protobuf.Duration
	9,  // 22: config.v1.Service.GRPC.idle_timeout:type_name -> google.protobuf.Duration
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_config_v1_service_proto_init() }
func file_config_v1_service_proto_init() {
	if File_config_v1_service_proto != nil {
		return
	}
	file_config_v1_message_proto_init()
	file_config_v1_middleware_proto_init()
	file_config_v1_task_proto_init()
	file_config_v1_websocket_proto_init()
	file_config_v1_service_proto_msgTypes[1].OneofWrappers = []any{}
	file_config_v1_service_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_v1_service_proto_goTypes,
		DependencyIndexes: file_config_v1_service_proto_depIdxs,
		MessageInfos:      file_config_v1_service_proto_msgTypes,
	}.Build()
	File_config_v1_service_proto = out.File
	file_config_v1_service_proto_rawDesc = nil
	file_config_v1_service_proto_goTypes = nil
	file_config_v1_service_proto_depIdxs = nil
}
