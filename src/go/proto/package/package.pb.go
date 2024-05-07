// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: package/package.proto

package _package

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "market.goravel.dev/proto/base"
	user "market.goravel.dev/proto/user"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string     `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Summary       string     `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	Description   string     `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Link          string     `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	Version       string     `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	LastUpdatedAt string     `protobuf:"bytes,8,opt,name=last_updated_at,json=lastUpdatedAt,proto3" json:"last_updated_at,omitempty"`
	CreatedAt     string     `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string     `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	User          *user.User `protobuf:"bytes,11,opt,name=user,proto3" json:"user,omitempty"`
	Tags          []*Tag     `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	ViewCount     uint32     `protobuf:"varint,13,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_package_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_package_package_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Package) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Package) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Package) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Package) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Package) GetLastUpdatedAt() string {
	if x != nil {
		return x.LastUpdatedAt
	}
	return ""
}

func (x *Package) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Package) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Package) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Package) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Package) GetViewCount() uint32 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

type GetPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auto-injected by the API Gateway.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPackageRequest) Reset() {
	*x = GetPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageRequest) ProtoMessage() {}

func (x *GetPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageRequest.ProtoReflect.Descriptor instead.
func (*GetPackageRequest) Descriptor() ([]byte, []int) {
	return file_package_package_proto_rawDescGZIP(), []int{1}
}

func (x *GetPackageRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetPackageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *base.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Package *Package     `protobuf:"bytes,2,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *GetPackageResponse) Reset() {
	*x = GetPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageResponse) ProtoMessage() {}

func (x *GetPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageResponse.ProtoReflect.Descriptor instead.
func (*GetPackageResponse) Descriptor() ([]byte, []int) {
	return file_package_package_proto_rawDescGZIP(), []int{2}
}

func (x *GetPackageResponse) GetStatus() *base.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetPackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type PackagesQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId   string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *PackagesQuery) Reset() {
	*x = PackagesQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackagesQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackagesQuery) ProtoMessage() {}

func (x *PackagesQuery) ProtoReflect() protoreflect.Message {
	mi := &file_package_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackagesQuery.ProtoReflect.Descriptor instead.
func (*PackagesQuery) Descriptor() ([]byte, []int) {
	return file_package_package_proto_rawDescGZIP(), []int{3}
}

func (x *PackagesQuery) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *PackagesQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PackagesQuery) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *base.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Query      *PackagesQuery   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *GetPackagesRequest) Reset() {
	*x = GetPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackagesRequest) ProtoMessage() {}

func (x *GetPackagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackagesRequest.ProtoReflect.Descriptor instead.
func (*GetPackagesRequest) Descriptor() ([]byte, []int) {
	return file_package_package_proto_rawDescGZIP(), []int{4}
}

func (x *GetPackagesRequest) GetPagination() *base.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetPackagesRequest) GetQuery() *PackagesQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

type GetPackagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *base.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Packages []*Package   `protobuf:"bytes,2,rep,name=packages,proto3" json:"packages,omitempty"`
	Total    int64        `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetPackagesResponse) Reset() {
	*x = GetPackagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_package_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackagesResponse) ProtoMessage() {}

func (x *GetPackagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_package_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackagesResponse.ProtoReflect.Descriptor instead.
func (*GetPackagesResponse) Descriptor() ([]byte, []int) {
	return file_package_package_proto_rawDescGZIP(), []int{5}
}

func (x *GetPackagesResponse) GetStatus() *base.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetPackagesResponse) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *GetPackagesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreatePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auto-injected by the API Gateway.
	UserId      string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover       string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Url         string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Summary     string   `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Tags        []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreatePackageRequest) Reset() {
	*x = CreatePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_package_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageRequest) ProtoMessage() {}

func (x *CreatePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_package_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackageRequest.ProtoReflect.Descriptor instead.
func (*CreatePackageRequest) Descriptor() ([]byte, []int) {
	return file_package_package_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePackageRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreatePackageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePackageRequest) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *CreatePackageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreatePackageRequest) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CreatePackageRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePackageRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CreatePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *base.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Package *Package     `protobuf:"bytes,2,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *CreatePackageResponse) Reset() {
	*x = CreatePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_package_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageResponse) ProtoMessage() {}

func (x *CreatePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_package_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackageResponse.ProtoReflect.Descriptor instead.
func (*CreatePackageResponse) Descriptor() ([]byte, []int) {
	return file_package_package_proto_rawDescGZIP(), []int{7}
}

func (x *CreatePackageResponse) GetStatus() *base.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CreatePackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

var File_package_package_proto protoreflect.FileDescriptor

var file_package_package_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x74, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x7f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0xbb, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x22, 0x69, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x32, 0x8c, 0x03,
	0x0a, 0x0e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12,
	0x0e, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x54, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x5b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x68, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x77, 0x42, 0x22, 0x5a, 0x20,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x67, 0x6f, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_package_proto_rawDescOnce sync.Once
	file_package_package_proto_rawDescData = file_package_package_proto_rawDesc
)

func file_package_package_proto_rawDescGZIP() []byte {
	file_package_package_proto_rawDescOnce.Do(func() {
		file_package_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_package_proto_rawDescData)
	})
	return file_package_package_proto_rawDescData
}

var file_package_package_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_package_package_proto_goTypes = []interface{}{
	(*Package)(nil),               // 0: package.Package
	(*GetPackageRequest)(nil),     // 1: package.GetPackageRequest
	(*GetPackageResponse)(nil),    // 2: package.GetPackageResponse
	(*PackagesQuery)(nil),         // 3: package.PackagesQuery
	(*GetPackagesRequest)(nil),    // 4: package.GetPackagesRequest
	(*GetPackagesResponse)(nil),   // 5: package.GetPackagesResponse
	(*CreatePackageRequest)(nil),  // 6: package.CreatePackageRequest
	(*CreatePackageResponse)(nil), // 7: package.CreatePackageResponse
	(*user.User)(nil),             // 8: user.User
	(*Tag)(nil),                   // 9: package.Tag
	(*base.Status)(nil),           // 10: base.Status
	(*base.Pagination)(nil),       // 11: base.Pagination
	(*GetTagsRequest)(nil),        // 12: package.GetTagsRequest
	(*GetTagsResponse)(nil),       // 13: package.GetTagsResponse
}
var file_package_package_proto_depIdxs = []int32{
	8,  // 0: package.Package.user:type_name -> user.User
	9,  // 1: package.Package.tags:type_name -> package.Tag
	10, // 2: package.GetPackageResponse.status:type_name -> base.Status
	0,  // 3: package.GetPackageResponse.package:type_name -> package.Package
	11, // 4: package.GetPackagesRequest.pagination:type_name -> base.Pagination
	3,  // 5: package.GetPackagesRequest.query:type_name -> package.PackagesQuery
	10, // 6: package.GetPackagesResponse.status:type_name -> base.Status
	0,  // 7: package.GetPackagesResponse.packages:type_name -> package.Package
	10, // 8: package.CreatePackageResponse.status:type_name -> base.Status
	0,  // 9: package.CreatePackageResponse.package:type_name -> package.Package
	1,  // 10: package.PackageService.GetPackage:input_type -> package.GetPackageRequest
	12, // 11: package.PackageService.GetTags:input_type -> package.GetTagsRequest
	4,  // 12: package.PackageService.GetPackages:input_type -> package.GetPackagesRequest
	6,  // 13: package.PackageService.CreatePackage:input_type -> package.CreatePackageRequest
	2,  // 14: package.PackageService.GetPackage:output_type -> package.GetPackageResponse
	13, // 15: package.PackageService.GetTags:output_type -> package.GetTagsResponse
	5,  // 16: package.PackageService.GetPackages:output_type -> package.GetPackagesResponse
	7,  // 17: package.PackageService.CreatePackage:output_type -> package.CreatePackageResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_package_package_proto_init() }
func file_package_package_proto_init() {
	if File_package_package_proto != nil {
		return
	}
	file_package_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_package_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_package_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageRequest); i {
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
		file_package_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageResponse); i {
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
		file_package_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackagesQuery); i {
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
		file_package_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackagesRequest); i {
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
		file_package_package_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackagesResponse); i {
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
		file_package_package_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackageRequest); i {
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
		file_package_package_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackageResponse); i {
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
			RawDescriptor: file_package_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_package_package_proto_goTypes,
		DependencyIndexes: file_package_package_proto_depIdxs,
		MessageInfos:      file_package_package_proto_msgTypes,
	}.Build()
	File_package_package_proto = out.File
	file_package_package_proto_rawDesc = nil
	file_package_package_proto_goTypes = nil
	file_package_package_proto_depIdxs = nil
}
