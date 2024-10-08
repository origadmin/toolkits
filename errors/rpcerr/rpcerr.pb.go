// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.2
// source: errors/rpcerr/rpcerr.proto

package rpcerr

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Detail string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_errors_rpcerr_rpcerr_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_errors_rpcerr_rpcerr_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_errors_rpcerr_rpcerr_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type MultiError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []*Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *MultiError) Reset() {
	*x = MultiError{}
	mi := &file_errors_rpcerr_rpcerr_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiError) ProtoMessage() {}

func (x *MultiError) ProtoReflect() protoreflect.Message {
	mi := &file_errors_rpcerr_rpcerr_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiError.ProtoReflect.Descriptor instead.
func (*MultiError) Descriptor() ([]byte, []int) {
	return file_errors_rpcerr_rpcerr_proto_rawDescGZIP(), []int{1}
}

func (x *MultiError) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_errors_rpcerr_rpcerr_proto protoreflect.FileDescriptor

var file_errors_rpcerr_rpcerr_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x65, 0x72, 0x72, 0x2f,
	0x72, 0x70, 0x63, 0x65, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x65, 0x72, 0x72, 0x22, 0x43, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x22, 0x3a, 0x0a, 0x0a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2c,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x65, 0x72, 0x72, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x72, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6f, 0x72, 0x69, 0x67, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x65,
	0x72, 0x72, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x6b,
	0x69, 0x74, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x65, 0x72,
	0x72, 0x3b, 0x72, 0x70, 0x63, 0x65, 0x72, 0x72, 0xa2, 0x02, 0x15, 0x4f, 0x72, 0x69, 0x67, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x70, 0x63, 0x65, 0x72, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_rpcerr_rpcerr_proto_rawDescOnce sync.Once
	file_errors_rpcerr_rpcerr_proto_rawDescData = file_errors_rpcerr_rpcerr_proto_rawDesc
)

func file_errors_rpcerr_rpcerr_proto_rawDescGZIP() []byte {
	file_errors_rpcerr_rpcerr_proto_rawDescOnce.Do(func() {
		file_errors_rpcerr_rpcerr_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_rpcerr_rpcerr_proto_rawDescData)
	})
	return file_errors_rpcerr_rpcerr_proto_rawDescData
}

var file_errors_rpcerr_rpcerr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_errors_rpcerr_rpcerr_proto_goTypes = []any{
	(*Error)(nil),      // 0: errors.rpcerr.Error
	(*MultiError)(nil), // 1: errors.rpcerr.MultiError
}
var file_errors_rpcerr_rpcerr_proto_depIdxs = []int32{
	0, // 0: errors.rpcerr.MultiError.errors:type_name -> errors.rpcerr.Error
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_errors_rpcerr_rpcerr_proto_init() }
func file_errors_rpcerr_rpcerr_proto_init() {
	if File_errors_rpcerr_rpcerr_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_rpcerr_rpcerr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_rpcerr_rpcerr_proto_goTypes,
		DependencyIndexes: file_errors_rpcerr_rpcerr_proto_depIdxs,
		MessageInfos:      file_errors_rpcerr_rpcerr_proto_msgTypes,
	}.Build()
	File_errors_rpcerr_rpcerr_proto = out.File
	file_errors_rpcerr_rpcerr_proto_rawDesc = nil
	file_errors_rpcerr_rpcerr_proto_goTypes = nil
	file_errors_rpcerr_rpcerr_proto_depIdxs = nil
}
