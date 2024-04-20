// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: office_svc.proto

package office_svc

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

type PaginationMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalResults int32 `protobuf:"varint,1,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	TotalPages   int32 `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Page         int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage      int32 `protobuf:"varint,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *PaginationMeta) Reset() {
	*x = PaginationMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationMeta) ProtoMessage() {}

func (x *PaginationMeta) ProtoReflect() protoreflect.Message {
	mi := &file_office_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationMeta.ProtoReflect.Descriptor instead.
func (*PaginationMeta) Descriptor() ([]byte, []int) {
	return file_office_svc_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationMeta) GetTotalResults() int32 {
	if x != nil {
		return x.TotalResults
	}
	return 0
}

func (x *PaginationMeta) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *PaginationMeta) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationMeta) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_office_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_office_svc_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PaginationRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_office_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_office_svc_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type OfficeOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *PaginationMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data []*Data         `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OfficeOutput) Reset() {
	*x = OfficeOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeOutput) ProtoMessage() {}

func (x *OfficeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_office_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeOutput.ProtoReflect.Descriptor instead.
func (*OfficeOutput) Descriptor() ([]byte, []int) {
	return file_office_svc_proto_rawDescGZIP(), []int{3}
}

func (x *OfficeOutput) GetMeta() *PaginationMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *OfficeOutput) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type OfficeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *OfficeInput) Reset() {
	*x = OfficeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_svc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeInput) ProtoMessage() {}

func (x *OfficeInput) ProtoReflect() protoreflect.Message {
	mi := &file_office_svc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeInput.ProtoReflect.Descriptor instead.
func (*OfficeInput) Descriptor() ([]byte, []int) {
	return file_office_svc_proto_rawDescGZIP(), []int{4}
}

func (x *OfficeInput) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_office_svc_proto protoreflect.FileDescriptor

var file_office_svc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x22, 0x86,
	0x01, 0x0a, 0x0f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x0c, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x76, 0x63, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f,
	0x73, 0x76, 0x63, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4d,
	0x0a, 0x0b, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3e, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x54, 0x0a,
	0x0d, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43,
	0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x17,
	0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x5f, 0x73, 0x76, 0x63, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f,
	0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_office_svc_proto_rawDescOnce sync.Once
	file_office_svc_proto_rawDescData = file_office_svc_proto_rawDesc
)

func file_office_svc_proto_rawDescGZIP() []byte {
	file_office_svc_proto_rawDescOnce.Do(func() {
		file_office_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_office_svc_proto_rawDescData)
	})
	return file_office_svc_proto_rawDescData
}

var file_office_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_office_svc_proto_goTypes = []interface{}{
	(*PaginationMeta)(nil),    // 0: office_svc.pagination_meta
	(*PaginationRequest)(nil), // 1: office_svc.pagination_request
	(*Data)(nil),              // 2: office_svc.data
	(*OfficeOutput)(nil),      // 3: office_svc.OfficeOutput
	(*OfficeInput)(nil),       // 4: office_svc.OfficeInput
}
var file_office_svc_proto_depIdxs = []int32{
	0, // 0: office_svc.OfficeOutput.meta:type_name -> office_svc.pagination_meta
	2, // 1: office_svc.OfficeOutput.data:type_name -> office_svc.data
	1, // 2: office_svc.OfficeInput.pagination:type_name -> office_svc.pagination_request
	4, // 3: office_svc.OfficeService.SearchOffice:input_type -> office_svc.OfficeInput
	3, // 4: office_svc.OfficeService.SearchOffice:output_type -> office_svc.OfficeOutput
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_office_svc_proto_init() }
func file_office_svc_proto_init() {
	if File_office_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_office_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationMeta); i {
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
		file_office_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationRequest); i {
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
		file_office_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_office_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeOutput); i {
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
		file_office_svc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeInput); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_office_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_office_svc_proto_goTypes,
		DependencyIndexes: file_office_svc_proto_depIdxs,
		MessageInfos:      file_office_svc_proto_msgTypes,
	}.Build()
	File_office_svc_proto = out.File
	file_office_svc_proto_rawDesc = nil
	file_office_svc_proto_goTypes = nil
	file_office_svc_proto_depIdxs = nil
}
