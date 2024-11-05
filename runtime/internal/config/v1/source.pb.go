// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: config/v1/source.proto

package config

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Registration source center
type SourceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"` // Type
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Name
	File        *SourceConfig_File   `protobuf:"bytes,3,opt,name=file,proto3,oneof" json:"file,omitempty"`
	Consul      *SourceConfig_Consul `protobuf:"bytes,4,opt,name=consul,proto3,oneof" json:"consul,omitempty"` // Consul
	Etcd        *SourceConfig_ETCD   `protobuf:"bytes,5,opt,name=etcd,proto3,oneof" json:"etcd,omitempty"`     // ETCD
	Formats     []string             `protobuf:"bytes,6,rep,name=formats,proto3" json:"formats,omitempty"`
	EnvPrefixes []string             `protobuf:"bytes,8,rep,name=env_prefixes,proto3" json:"env_prefixes,omitempty"`
	EnvArgs     map[string]string    `protobuf:"bytes,7,rep,name=env_args,proto3" json:"env_args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SourceConfig) Reset() {
	*x = SourceConfig{}
	mi := &file_config_v1_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceConfig) ProtoMessage() {}

func (x *SourceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceConfig.ProtoReflect.Descriptor instead.
func (*SourceConfig) Descriptor() ([]byte, []int) {
	return file_config_v1_source_proto_rawDescGZIP(), []int{0}
}

func (x *SourceConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SourceConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SourceConfig) GetFile() *SourceConfig_File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *SourceConfig) GetConsul() *SourceConfig_Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *SourceConfig) GetEtcd() *SourceConfig_ETCD {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *SourceConfig) GetFormats() []string {
	if x != nil {
		return x.Formats
	}
	return nil
}

func (x *SourceConfig) GetEnvPrefixes() []string {
	if x != nil {
		return x.EnvPrefixes
	}
	return nil
}

func (x *SourceConfig) GetEnvArgs() map[string]string {
	if x != nil {
		return x.EnvArgs
	}
	return nil
}

// File
type SourceConfig_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Format string `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *SourceConfig_File) Reset() {
	*x = SourceConfig_File{}
	mi := &file_config_v1_source_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceConfig_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceConfig_File) ProtoMessage() {}

func (x *SourceConfig_File) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_source_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceConfig_File.ProtoReflect.Descriptor instead.
func (*SourceConfig_File) Descriptor() ([]byte, []int) {
	return file_config_v1_source_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SourceConfig_File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SourceConfig_File) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

// Consul
type SourceConfig_Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Scheme  string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Path    string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *SourceConfig_Consul) Reset() {
	*x = SourceConfig_Consul{}
	mi := &file_config_v1_source_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceConfig_Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceConfig_Consul) ProtoMessage() {}

func (x *SourceConfig_Consul) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_source_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceConfig_Consul.ProtoReflect.Descriptor instead.
func (*SourceConfig_Consul) Descriptor() ([]byte, []int) {
	return file_config_v1_source_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SourceConfig_Consul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SourceConfig_Consul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *SourceConfig_Consul) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SourceConfig_Consul) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type SourceConfig_ETCD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *SourceConfig_ETCD) Reset() {
	*x = SourceConfig_ETCD{}
	mi := &file_config_v1_source_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceConfig_ETCD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceConfig_ETCD) ProtoMessage() {}

func (x *SourceConfig_ETCD) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_source_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceConfig_ETCD.ProtoReflect.Descriptor instead.
func (*SourceConfig_ETCD) Descriptor() ([]byte, []int) {
	return file_config_v1_source_proto_rawDescGZIP(), []int{0, 2}
}

func (x *SourceConfig_ETCD) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

var File_config_v1_source_proto protoreflect.FileDescriptor

var file_config_v1_source_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x05, 0x0a,
	0x0c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3d, 0xfa, 0x42, 0x3a,
	0x72, 0x38, 0x52, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x52, 0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x52, 0x06, 0x61,
	0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x48, 0x01, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x45, 0x54, 0x43, 0x44, 0x48, 0x02, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x76,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x6e, 0x76, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x40, 0x0a,
	0x08, 0x65, 0x6e, 0x76, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x41, 0x72, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x1a,
	0x32, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x1a, 0x64, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x24, 0x0a, 0x04, 0x45, 0x54, 0x43,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a,
	0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x42, 0x96, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa2, 0x02,
	0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_v1_source_proto_rawDescOnce sync.Once
	file_config_v1_source_proto_rawDescData = file_config_v1_source_proto_rawDesc
)

func file_config_v1_source_proto_rawDescGZIP() []byte {
	file_config_v1_source_proto_rawDescOnce.Do(func() {
		file_config_v1_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1_source_proto_rawDescData)
	})
	return file_config_v1_source_proto_rawDescData
}

var file_config_v1_source_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_config_v1_source_proto_goTypes = []any{
	(*SourceConfig)(nil),        // 0: config.v1.SourceConfig
	(*SourceConfig_File)(nil),   // 1: config.v1.SourceConfig.File
	(*SourceConfig_Consul)(nil), // 2: config.v1.SourceConfig.Consul
	(*SourceConfig_ETCD)(nil),   // 3: config.v1.SourceConfig.ETCD
	nil,                         // 4: config.v1.SourceConfig.EnvArgsEntry
}
var file_config_v1_source_proto_depIdxs = []int32{
	1, // 0: config.v1.SourceConfig.file:type_name -> config.v1.SourceConfig.File
	2, // 1: config.v1.SourceConfig.consul:type_name -> config.v1.SourceConfig.Consul
	3, // 2: config.v1.SourceConfig.etcd:type_name -> config.v1.SourceConfig.ETCD
	4, // 3: config.v1.SourceConfig.env_args:type_name -> config.v1.SourceConfig.EnvArgsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_config_v1_source_proto_init() }
func file_config_v1_source_proto_init() {
	if File_config_v1_source_proto != nil {
		return
	}
	file_config_v1_source_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_v1_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_v1_source_proto_goTypes,
		DependencyIndexes: file_config_v1_source_proto_depIdxs,
		MessageInfos:      file_config_v1_source_proto_msgTypes,
	}.Build()
	File_config_v1_source_proto = out.File
	file_config_v1_source_proto_rawDesc = nil
	file_config_v1_source_proto_goTypes = nil
	file_config_v1_source_proto_depIdxs = nil
}