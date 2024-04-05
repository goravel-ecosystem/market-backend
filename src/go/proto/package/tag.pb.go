// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: package/tag.proto

package _package

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "market.goravel.dev/proto/base"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_package_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_package_tag_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Tag) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Tag) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type TagsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId string `protobuf:"bytes,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TagsQuery) Reset() {
	*x = TagsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagsQuery) ProtoMessage() {}

func (x *TagsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_package_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagsQuery.ProtoReflect.Descriptor instead.
func (*TagsQuery) Descriptor() ([]byte, []int) {
	return file_package_tag_proto_rawDescGZIP(), []int{1}
}

func (x *TagsQuery) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *TagsQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *base.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Query      *TagsQuery       `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *GetTagsRequest) Reset() {
	*x = GetTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagsRequest) ProtoMessage() {}

func (x *GetTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagsRequest.ProtoReflect.Descriptor instead.
func (*GetTagsRequest) Descriptor() ([]byte, []int) {
	return file_package_tag_proto_rawDescGZIP(), []int{2}
}

func (x *GetTagsRequest) GetPagination() *base.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetTagsRequest) GetQuery() *TagsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

type GetTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *base.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Tags   []*Tag       `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Total  int64        `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetTagsResponse) Reset() {
	*x = GetTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagsResponse) ProtoMessage() {}

func (x *GetTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagsResponse.ProtoReflect.Descriptor instead.
func (*GetTagsResponse) Descriptor() ([]byte, []int) {
	return file_package_tag_proto_rawDescGZIP(), []int{3}
}

func (x *GetTagsResponse) GetStatus() *base.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetTagsResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetTagsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_package_tag_proto protoreflect.FileDescriptor

var file_package_tag_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x3e, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x6c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x61, 0x67,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x6f, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x22,
	0x5a, 0x20, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x67, 0x6f, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_tag_proto_rawDescOnce sync.Once
	file_package_tag_proto_rawDescData = file_package_tag_proto_rawDesc
)

func file_package_tag_proto_rawDescGZIP() []byte {
	file_package_tag_proto_rawDescOnce.Do(func() {
		file_package_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_tag_proto_rawDescData)
	})
	return file_package_tag_proto_rawDescData
}

var file_package_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_package_tag_proto_goTypes = []interface{}{
	(*Tag)(nil),             // 0: package.Tag
	(*TagsQuery)(nil),       // 1: package.TagsQuery
	(*GetTagsRequest)(nil),  // 2: package.GetTagsRequest
	(*GetTagsResponse)(nil), // 3: package.GetTagsResponse
	(*base.Pagination)(nil), // 4: base.Pagination
	(*base.Status)(nil),     // 5: base.Status
}
var file_package_tag_proto_depIdxs = []int32{
	4, // 0: package.GetTagsRequest.pagination:type_name -> base.Pagination
	1, // 1: package.GetTagsRequest.query:type_name -> package.TagsQuery
	5, // 2: package.GetTagsResponse.status:type_name -> base.Status
	0, // 3: package.GetTagsResponse.tags:type_name -> package.Tag
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_package_tag_proto_init() }
func file_package_tag_proto_init() {
	if File_package_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_package_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_package_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagsQuery); i {
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
		file_package_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagsRequest); i {
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
		file_package_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagsResponse); i {
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
			RawDescriptor: file_package_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_package_tag_proto_goTypes,
		DependencyIndexes: file_package_tag_proto_depIdxs,
		MessageInfos:      file_package_tag_proto_msgTypes,
	}.Build()
	File_package_tag_proto = out.File
	file_package_tag_proto_rawDesc = nil
	file_package_tag_proto_goTypes = nil
	file_package_tag_proto_depIdxs = nil
}
