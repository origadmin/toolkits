// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: config/v1/registry.proto

package config

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// Registration registry center
type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string           `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`                 // Type
	ServiceName string           `protobuf:"bytes,2,opt,name=service_name,proto3" json:"service_name,omitempty"` // ServiceName
	Consul      *Registry_Consul `protobuf:"bytes,300,opt,name=consul,proto3,oneof" json:"consul,omitempty"`     // Consul
	Etcd        *Registry_ETCD   `protobuf:"bytes,400,opt,name=etcd,proto3,oneof" json:"etcd,omitempty"`         // ETCD
	Custom      *Registry_Custom `protobuf:"bytes,500,opt,name=custom,proto3,oneof" json:"custom,omitempty"`     // Custom
}

func (x *Registry) Reset() {
	*x = Registry{}
	mi := &file_config_v1_registry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_registry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_config_v1_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Registry) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Registry) GetConsul() *Registry_Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Registry) GetEtcd() *Registry_ETCD {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *Registry) GetCustom() *Registry_Custom {
	if x != nil {
		return x.Custom
	}
	return nil
}

// Consul
type Registry_Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Scheme      string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Token       string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	HeartBeat   bool   `protobuf:"varint,4,opt,name=heart_beat,proto3" json:"heart_beat,omitempty"`
	HealthCheck bool   `protobuf:"varint,5,opt,name=health_check,proto3" json:"health_check,omitempty"`
	Datacenter  string `protobuf:"bytes,6,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	//  string tag = 7 [json_name = "tag"];
	HealthCheckInterval uint32 `protobuf:"varint,8,opt,name=health_check_interval,proto3" json:"health_check_interval,omitempty"`
	//  string health_check_timeout = 9[json_name = "health_check_timeout"];
	Timeout                        *durationpb.Duration `protobuf:"bytes,10,opt,name=timeout,proto3" json:"timeout,omitempty"`
	DeregisterCriticalServiceAfter uint32               `protobuf:"varint,11,opt,name=deregister_critical_service_after,proto3" json:"deregister_critical_service_after,omitempty"`
}

func (x *Registry_Consul) Reset() {
	*x = Registry_Consul{}
	mi := &file_config_v1_registry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry_Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Consul) ProtoMessage() {}

func (x *Registry_Consul) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_registry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Consul.ProtoReflect.Descriptor instead.
func (*Registry_Consul) Descriptor() ([]byte, []int) {
	return file_config_v1_registry_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Registry_Consul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Registry_Consul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Registry_Consul) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Registry_Consul) GetHeartBeat() bool {
	if x != nil {
		return x.HeartBeat
	}
	return false
}

func (x *Registry_Consul) GetHealthCheck() bool {
	if x != nil {
		return x.HealthCheck
	}
	return false
}

func (x *Registry_Consul) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

func (x *Registry_Consul) GetHealthCheckInterval() uint32 {
	if x != nil {
		return x.HealthCheckInterval
	}
	return 0
}

func (x *Registry_Consul) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Registry_Consul) GetDeregisterCriticalServiceAfter() uint32 {
	if x != nil {
		return x.DeregisterCriticalServiceAfter
	}
	return 0
}

// ETCD
type Registry_ETCD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *Registry_ETCD) Reset() {
	*x = Registry_ETCD{}
	mi := &file_config_v1_registry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry_ETCD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_ETCD) ProtoMessage() {}

func (x *Registry_ETCD) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_registry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_ETCD.ProtoReflect.Descriptor instead.
func (*Registry_ETCD) Descriptor() ([]byte, []int) {
	return file_config_v1_registry_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Registry_ETCD) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type Registry_Custom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *anypb.Any `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Registry_Custom) Reset() {
	*x = Registry_Custom{}
	mi := &file_config_v1_registry_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry_Custom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Custom) ProtoMessage() {}

func (x *Registry_Custom) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_registry_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Custom.ProtoReflect.Descriptor instead.
func (*Registry_Custom) Descriptor() ([]byte, []int) {
	return file_config_v1_registry_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Registry_Custom) GetConfig() *anypb.Any {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_config_v1_registry_proto protoreflect.FileDescriptor

var file_config_v1_registry_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x06,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x59, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x45, 0xba, 0x48, 0x42, 0x72, 0x40, 0x52,
	0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x52, 0x04, 0x65,
	0x74, 0x63, 0x64, 0x52, 0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x52, 0x06, 0x61, 0x70, 0x6f, 0x6c,
	0x6c, 0x6f, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x07,
	0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x90, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x54, 0x43, 0x44, 0x48, 0x01, 0x52, 0x04,
	0x65, 0x74, 0x63, 0x64, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x48, 0x02, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x88, 0x01,
	0x01, 0x1a, 0xed, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x65,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f,
	0x62, 0x65, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x15, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x33,
	0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x4c, 0x0a, 0x21, 0x64, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x21,
	0x64, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x1a, 0x24, 0x0a, 0x04, 0x45, 0x54, 0x43, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x36, 0x0a, 0x06, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x65,
	0x74, 0x63, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x9b,
	0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72,
	0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x73,
	0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_v1_registry_proto_rawDescOnce sync.Once
	file_config_v1_registry_proto_rawDescData = file_config_v1_registry_proto_rawDesc
)

func file_config_v1_registry_proto_rawDescGZIP() []byte {
	file_config_v1_registry_proto_rawDescOnce.Do(func() {
		file_config_v1_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1_registry_proto_rawDescData)
	})
	return file_config_v1_registry_proto_rawDescData
}

var file_config_v1_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_v1_registry_proto_goTypes = []any{
	(*Registry)(nil),            // 0: config.v1.Registry
	(*Registry_Consul)(nil),     // 1: config.v1.Registry.Consul
	(*Registry_ETCD)(nil),       // 2: config.v1.Registry.ETCD
	(*Registry_Custom)(nil),     // 3: config.v1.Registry.Custom
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
	(*anypb.Any)(nil),           // 5: google.protobuf.Any
}
var file_config_v1_registry_proto_depIdxs = []int32{
	1, // 0: config.v1.Registry.consul:type_name -> config.v1.Registry.Consul
	2, // 1: config.v1.Registry.etcd:type_name -> config.v1.Registry.ETCD
	3, // 2: config.v1.Registry.custom:type_name -> config.v1.Registry.Custom
	4, // 3: config.v1.Registry.Consul.timeout:type_name -> google.protobuf.Duration
	5, // 4: config.v1.Registry.Custom.config:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_v1_registry_proto_init() }
func file_config_v1_registry_proto_init() {
	if File_config_v1_registry_proto != nil {
		return
	}
	file_config_v1_registry_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_v1_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_v1_registry_proto_goTypes,
		DependencyIndexes: file_config_v1_registry_proto_depIdxs,
		MessageInfos:      file_config_v1_registry_proto_msgTypes,
	}.Build()
	File_config_v1_registry_proto = out.File
	file_config_v1_registry_proto_rawDesc = nil
	file_config_v1_registry_proto_goTypes = nil
	file_config_v1_registry_proto_depIdxs = nil
}
