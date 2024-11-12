// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: config/v1/logger.proto

package config

import (
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

type Logger_Level int32

const (
	Logger_LEVEL_UNSPECIFIED Logger_Level = 0
	Logger_LEVEL_DEBUG       Logger_Level = 1
	Logger_LEVEL_INFO        Logger_Level = 2
	Logger_LEVEL_WARN        Logger_Level = 3
	Logger_LEVEL_ERROR       Logger_Level = 4
	Logger_LEVEL_FATAL       Logger_Level = 5
)

// Enum value maps for Logger_Level.
var (
	Logger_Level_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "LEVEL_DEBUG",
		2: "LEVEL_INFO",
		3: "LEVEL_WARN",
		4: "LEVEL_ERROR",
		5: "LEVEL_FATAL",
	}
	Logger_Level_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"LEVEL_DEBUG":       1,
		"LEVEL_INFO":        2,
		"LEVEL_WARN":        3,
		"LEVEL_ERROR":       4,
		"LEVEL_FATAL":       5,
	}
)

func (x Logger_Level) Enum() *Logger_Level {
	p := new(Logger_Level)
	*p = x
	return p
}

func (x Logger_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Logger_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_config_v1_logger_proto_enumTypes[0].Descriptor()
}

func (Logger_Level) Type() protoreflect.EnumType {
	return &file_config_v1_logger_proto_enumTypes[0]
}

func (x Logger_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Logger_Level.Descriptor instead.
func (Logger_Level) EnumDescriptor() ([]byte, []int) {
	return file_config_v1_logger_proto_rawDescGZIP(), []int{0, 0}
}

// logger config.
type Logger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level Logger_Level `protobuf:"varint,3,opt,name=level,proto3,enum=config.v1.Logger_Level" json:"level,omitempty"`
}

func (x *Logger) Reset() {
	*x = Logger{}
	mi := &file_config_v1_logger_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger) ProtoMessage() {}

func (x *Logger) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_logger_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger.ProtoReflect.Descriptor instead.
func (*Logger) Descriptor() ([]byte, []int) {
	return file_config_v1_logger_proto_rawDescGZIP(), []int{0}
}

func (x *Logger) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Logger) GetLevel() Logger_Level {
	if x != nil {
		return x.Level
	}
	return Logger_LEVEL_UNSPECIFIED
}

var File_config_v1_logger_proto protoreflect.FileDescriptor

var file_config_v1_logger_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x22, 0xbe, 0x01, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x71, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e,
	0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x46, 0x41, 0x54,
	0x41, 0x4c, 0x10, 0x05, 0x42, 0x99, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x6f, 0x6f, 0x6c,
	0x6b, 0x69, 0x74, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03,
	0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_v1_logger_proto_rawDescOnce sync.Once
	file_config_v1_logger_proto_rawDescData = file_config_v1_logger_proto_rawDesc
)

func file_config_v1_logger_proto_rawDescGZIP() []byte {
	file_config_v1_logger_proto_rawDescOnce.Do(func() {
		file_config_v1_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1_logger_proto_rawDescData)
	})
	return file_config_v1_logger_proto_rawDescData
}

var file_config_v1_logger_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_v1_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_v1_logger_proto_goTypes = []any{
	(Logger_Level)(0), // 0: config.v1.Logger.Level
	(*Logger)(nil),    // 1: config.v1.Logger
}
var file_config_v1_logger_proto_depIdxs = []int32{
	0, // 0: config.v1.Logger.level:type_name -> config.v1.Logger.Level
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_v1_logger_proto_init() }
func file_config_v1_logger_proto_init() {
	if File_config_v1_logger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_v1_logger_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_v1_logger_proto_goTypes,
		DependencyIndexes: file_config_v1_logger_proto_depIdxs,
		EnumInfos:         file_config_v1_logger_proto_enumTypes,
		MessageInfos:      file_config_v1_logger_proto_msgTypes,
	}.Build()
	File_config_v1_logger_proto = out.File
	file_config_v1_logger_proto_rawDesc = nil
	file_config_v1_logger_proto_goTypes = nil
	file_config_v1_logger_proto_depIdxs = nil
}
