// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.2
// source: middlewares/cors.proto

package middlewares

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

// Cors middleware config.
type CorsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool `protobuf:"varint,2,opt,name=allow_credentials,proto3" json:"allow_credentials,omitempty"`
	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is [*]
	AllowOrigins []string `protobuf:"bytes,3,rep,name=allow_origins,proto3" json:"allow_origins,omitempty"`
	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS)
	AllowMethods []string `protobuf:"bytes,4,rep,name=allow_methods,proto3" json:"allow_methods,omitempty"`
	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	AllowHeaders []string `protobuf:"bytes,5,rep,name=allow_headers,proto3" json:"allow_headers,omitempty"`
	// ExposeHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposeHeaders []string `protobuf:"bytes,6,rep,name=expose_headers,proto3" json:"expose_headers,omitempty"`
	// MaxAge indicates how long (with second-precision) the results of a preflight request
	// can be cached
	MaxAge *durationpb.Duration `protobuf:"bytes,7,opt,name=max_age,proto3" json:"max_age,omitempty"`
	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
	AllowWildcard bool `protobuf:"varint,8,opt,name=allow_wildcard,proto3" json:"allow_wildcard,omitempty"`
	// Allows usage of popular browser extensions schemas
	AllowBrowserExtensions bool `protobuf:"varint,9,opt,name=allow_browser_extensions,proto3" json:"allow_browser_extensions,omitempty"`
	// Allows usage of WebSocket protocol
	AllowWebSockets bool `protobuf:"varint,10,opt,name=allow_web_sockets,proto3" json:"allow_web_sockets,omitempty"`
	// Allows usage of private network addresses (127.0.0.1, [::1], localhost)
	AllowPrivateNetwork bool `protobuf:"varint,11,opt,name=allow_private_network,proto3" json:"allow_private_network,omitempty"`
	// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
	AllowFiles bool `protobuf:"varint,12,opt,name=allow_files,proto3" json:"allow_files,omitempty"`
}

func (x *CorsConfig) Reset() {
	*x = CorsConfig{}
	mi := &file_middlewares_cors_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CorsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorsConfig) ProtoMessage() {}

func (x *CorsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_middlewares_cors_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorsConfig.ProtoReflect.Descriptor instead.
func (*CorsConfig) Descriptor() ([]byte, []int) {
	return file_middlewares_cors_proto_rawDescGZIP(), []int{0}
}

func (x *CorsConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CorsConfig) GetAllowCredentials() bool {
	if x != nil {
		return x.AllowCredentials
	}
	return false
}

func (x *CorsConfig) GetAllowOrigins() []string {
	if x != nil {
		return x.AllowOrigins
	}
	return nil
}

func (x *CorsConfig) GetAllowMethods() []string {
	if x != nil {
		return x.AllowMethods
	}
	return nil
}

func (x *CorsConfig) GetAllowHeaders() []string {
	if x != nil {
		return x.AllowHeaders
	}
	return nil
}

func (x *CorsConfig) GetExposeHeaders() []string {
	if x != nil {
		return x.ExposeHeaders
	}
	return nil
}

func (x *CorsConfig) GetMaxAge() *durationpb.Duration {
	if x != nil {
		return x.MaxAge
	}
	return nil
}

func (x *CorsConfig) GetAllowWildcard() bool {
	if x != nil {
		return x.AllowWildcard
	}
	return false
}

func (x *CorsConfig) GetAllowBrowserExtensions() bool {
	if x != nil {
		return x.AllowBrowserExtensions
	}
	return false
}

func (x *CorsConfig) GetAllowWebSockets() bool {
	if x != nil {
		return x.AllowWebSockets
	}
	return false
}

func (x *CorsConfig) GetAllowPrivateNetwork() bool {
	if x != nil {
		return x.AllowPrivateNetwork
	}
	return false
}

func (x *CorsConfig) GetAllowFiles() bool {
	if x != nil {
		return x.AllowFiles
	}
	return false
}

var File_middlewares_cors_proto protoreflect.FileDescriptor

var file_middlewares_cors_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x04, 0x0a,
	0x0a, 0x43, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65,
	0x78, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x33, 0x0a,
	0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61,
	0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x77, 0x69, 0x6c, 0x64,
	0x63, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x12, 0x3a, 0x0a, 0x18, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x77, 0x65, 0x62, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x77, 0x65, 0x62, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x73, 0x2f, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_middlewares_cors_proto_rawDescOnce sync.Once
	file_middlewares_cors_proto_rawDescData = file_middlewares_cors_proto_rawDesc
)

func file_middlewares_cors_proto_rawDescGZIP() []byte {
	file_middlewares_cors_proto_rawDescOnce.Do(func() {
		file_middlewares_cors_proto_rawDescData = protoimpl.X.CompressGZIP(file_middlewares_cors_proto_rawDescData)
	})
	return file_middlewares_cors_proto_rawDescData
}

var file_middlewares_cors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_middlewares_cors_proto_goTypes = []any{
	(*CorsConfig)(nil),          // 0: middleware.cors.v1.CorsConfig
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
}
var file_middlewares_cors_proto_depIdxs = []int32{
	1, // 0: middleware.cors.v1.CorsConfig.max_age:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_middlewares_cors_proto_init() }
func file_middlewares_cors_proto_init() {
	if File_middlewares_cors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_middlewares_cors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_middlewares_cors_proto_goTypes,
		DependencyIndexes: file_middlewares_cors_proto_depIdxs,
		MessageInfos:      file_middlewares_cors_proto_msgTypes,
	}.Build()
	File_middlewares_cors_proto = out.File
	file_middlewares_cors_proto_rawDesc = nil
	file_middlewares_cors_proto_goTypes = nil
	file_middlewares_cors_proto_depIdxs = nil
}
