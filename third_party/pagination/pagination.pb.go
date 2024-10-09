// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: pagination/pagination.proto

package pagination

import (
	rpcerr "github.com/origadmin/toolkits/third_party/errors/rpcerr"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Paging sort
type SortOrder int32

const (
	// No sort
	SortOrder_UNSORTED SortOrder = 0
	// Ascending order
	SortOrder_ASCENDING SortOrder = 1
	// Descending order
	SortOrder_DESCENDING SortOrder = 2
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "UNSORTED",
		1: "ASCENDING",
		2: "DESCENDING",
	}
	SortOrder_value = map[string]int32{
		"UNSORTED":   0,
		"ASCENDING":  1,
		"DESCENDING": 2,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_pagination_pagination_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_pagination_pagination_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_pagination_pagination_proto_rawDescGZIP(), []int{0}
}

// Query parameters
type QueryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *QueryParam) Reset() {
	*x = QueryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagination_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParam) ProtoMessage() {}

func (x *QueryParam) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParam.ProtoReflect.Descriptor instead.
func (*QueryParam) Descriptor() ([]byte, []int) {
	return file_pagination_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *QueryParam) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// Paging general request
type PagingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current page
	Current *int32 `protobuf:"varint,1,opt,name=current,proto3,oneof" json:"current,omitempty"`
	// The number of lines per page
	PageSize *int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	// Query parameter
	Query map[string]*QueryParam `protobuf:"bytes,3,rep,name=query,proto3" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Sort
	OrderBy map[string]SortOrder `protobuf:"bytes,4,rep,name=order_by,json=orderBy,proto3" json:"order_by,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=pagination.SortOrder"`
	// Whether not paging
	NoPaging *bool `protobuf:"varint,5,opt,name=no_paging,json=noPaging,proto3,oneof" json:"no_paging,omitempty"`
}

func (x *PagingRequest) Reset() {
	*x = PagingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagination_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingRequest) ProtoMessage() {}

func (x *PagingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingRequest.ProtoReflect.Descriptor instead.
func (*PagingRequest) Descriptor() ([]byte, []int) {
	return file_pagination_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PagingRequest) GetCurrent() int32 {
	if x != nil && x.Current != nil {
		return *x.Current
	}
	return 0
}

func (x *PagingRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *PagingRequest) GetQuery() map[string]*QueryParam {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *PagingRequest) GetOrderBy() map[string]SortOrder {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

func (x *PagingRequest) GetNoPaging() bool {
	if x != nil && x.NoPaging != nil {
		return *x.NoPaging
	}
	return false
}

// Paging general result
type PagingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Total   int32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Error   *rpcerr.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Data    *anypb.Any    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Extra   *anypb.Any    `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *PagingResponse) Reset() {
	*x = PagingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagination_pagination_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingResponse) ProtoMessage() {}

func (x *PagingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_pagination_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingResponse.ProtoReflect.Descriptor instead.
func (*PagingResponse) Descriptor() ([]byte, []int) {
	return file_pagination_pagination_proto_rawDescGZIP(), []int{2}
}

func (x *PagingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PagingResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PagingResponse) GetError() *rpcerr.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PagingResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PagingResponse) GetExtra() *anypb.Any {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_pagination_pagination_proto protoreflect.FileDescriptor

var file_pagination_pagination_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x72, 0x70, 0x63,
	0x65, 0x72, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a,
	0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0xbe, 0x03, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x41, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x6f, 0x5f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x08, 0x6e, 0x6f, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x1a, 0x50, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x6f, 0x5f, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x22, 0xc2, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x65, 0x72, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2a, 0x38, 0x0a, 0x09, 0x53, 0x6f, 0x72,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x53, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x42, 0x20, 0x5a, 0x1e, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x73, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pagination_pagination_proto_rawDescOnce sync.Once
	file_pagination_pagination_proto_rawDescData = file_pagination_pagination_proto_rawDesc
)

func file_pagination_pagination_proto_rawDescGZIP() []byte {
	file_pagination_pagination_proto_rawDescOnce.Do(func() {
		file_pagination_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_pagination_pagination_proto_rawDescData)
	})
	return file_pagination_pagination_proto_rawDescData
}

var file_pagination_pagination_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pagination_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pagination_pagination_proto_goTypes = []any{
	(SortOrder)(0),         // 0: pagination.SortOrder
	(*QueryParam)(nil),     // 1: pagination.QueryParam
	(*PagingRequest)(nil),  // 2: pagination.PagingRequest
	(*PagingResponse)(nil), // 3: pagination.PagingResponse
	nil,                    // 4: pagination.PagingRequest.QueryEntry
	nil,                    // 5: pagination.PagingRequest.OrderByEntry
	(*rpcerr.Error)(nil),   // 6: errors.rpcerr.Error
	(*anypb.Any)(nil),      // 7: google.protobuf.Any
}
var file_pagination_pagination_proto_depIdxs = []int32{
	4, // 0: pagination.PagingRequest.query:type_name -> pagination.PagingRequest.QueryEntry
	5, // 1: pagination.PagingRequest.order_by:type_name -> pagination.PagingRequest.OrderByEntry
	6, // 2: pagination.PagingResponse.error:type_name -> errors.rpcerr.Error
	7, // 3: pagination.PagingResponse.data:type_name -> google.protobuf.Any
	7, // 4: pagination.PagingResponse.extra:type_name -> google.protobuf.Any
	1, // 5: pagination.PagingRequest.QueryEntry.value:type_name -> pagination.QueryParam
	0, // 6: pagination.PagingRequest.OrderByEntry.value:type_name -> pagination.SortOrder
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pagination_pagination_proto_init() }
func file_pagination_pagination_proto_init() {
	if File_pagination_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pagination_pagination_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*QueryParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pagination_pagination_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PagingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pagination_pagination_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PagingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pagination_pagination_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pagination_pagination_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pagination_pagination_proto_goTypes,
		DependencyIndexes: file_pagination_pagination_proto_depIdxs,
		EnumInfos:         file_pagination_pagination_proto_enumTypes,
		MessageInfos:      file_pagination_pagination_proto_msgTypes,
	}.Build()
	File_pagination_pagination_proto = out.File
	file_pagination_pagination_proto_rawDesc = nil
	file_pagination_pagination_proto_goTypes = nil
	file_pagination_pagination_proto_depIdxs = nil
}
