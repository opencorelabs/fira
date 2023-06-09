// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: protos/fira/v1/app-requests.proto

package v1

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

type CreateAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateAppRequest) Reset() {
	*x = CreateAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppRequest) ProtoMessage() {}

func (x *CreateAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppRequest.ProtoReflect.Descriptor instead.
func (*CreateAppRequest) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAppRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App *App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *CreateAppResponse) Reset() {
	*x = CreateAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppResponse) ProtoMessage() {}

func (x *CreateAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppResponse.ProtoReflect.Descriptor instead.
func (*CreateAppResponse) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAppResponse) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type GetAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetAppRequest) Reset() {
	*x = GetAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRequest) ProtoMessage() {}

func (x *GetAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRequest.ProtoReflect.Descriptor instead.
func (*GetAppRequest) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{2}
}

func (x *GetAppRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App *App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *GetAppResponse) Reset() {
	*x = GetAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppResponse) ProtoMessage() {}

func (x *GetAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppResponse.ProtoReflect.Descriptor instead.
func (*GetAppResponse) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{3}
}

func (x *GetAppResponse) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type ListAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAppsRequest) Reset() {
	*x = ListAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppsRequest) ProtoMessage() {}

func (x *ListAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppsRequest.ProtoReflect.Descriptor instead.
func (*ListAppsRequest) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{4}
}

type ListAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*App `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
}

func (x *ListAppsResponse) Reset() {
	*x = ListAppsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppsResponse) ProtoMessage() {}

func (x *ListAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppsResponse.ProtoReflect.Descriptor instead.
func (*ListAppsResponse) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{5}
}

func (x *ListAppsResponse) GetApps() []*App {
	if x != nil {
		return x.Apps
	}
	return nil
}

type RotateAppTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId       string      `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Environment Environment `protobuf:"varint,2,opt,name=environment,proto3,enum=protos.fira.v1.Environment" json:"environment,omitempty"`
}

func (x *RotateAppTokenRequest) Reset() {
	*x = RotateAppTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotateAppTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateAppTokenRequest) ProtoMessage() {}

func (x *RotateAppTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateAppTokenRequest.ProtoReflect.Descriptor instead.
func (*RotateAppTokenRequest) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{6}
}

func (x *RotateAppTokenRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *RotateAppTokenRequest) GetEnvironment() Environment {
	if x != nil {
		return x.Environment
	}
	return Environment_ENVIRONMENT_UNSPECIFIED
}

type RotateAppTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App *App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *RotateAppTokenResponse) Reset() {
	*x = RotateAppTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotateAppTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateAppTokenResponse) ProtoMessage() {}

func (x *RotateAppTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateAppTokenResponse.ProtoReflect.Descriptor instead.
func (*RotateAppTokenResponse) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{7}
}

func (x *RotateAppTokenResponse) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type InvalidateAppTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Jwt   string `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *InvalidateAppTokenRequest) Reset() {
	*x = InvalidateAppTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidateAppTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidateAppTokenRequest) ProtoMessage() {}

func (x *InvalidateAppTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidateAppTokenRequest.ProtoReflect.Descriptor instead.
func (*InvalidateAppTokenRequest) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{8}
}

func (x *InvalidateAppTokenRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *InvalidateAppTokenRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type InvalidateAppTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InvalidateAppTokenResponse) Reset() {
	*x = InvalidateAppTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fira_v1_app_requests_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidateAppTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidateAppTokenResponse) ProtoMessage() {}

func (x *InvalidateAppTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fira_v1_app_requests_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidateAppTokenResponse.ProtoReflect.Descriptor instead.
func (*InvalidateAppTokenResponse) Descriptor() ([]byte, []int) {
	return file_protos_fira_v1_app_requests_proto_rawDescGZIP(), []int{9}
}

var File_protos_fira_v1_app_requests_proto protoreflect.FileDescriptor

var file_protos_fira_v1_app_requests_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x69, 0x72, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x70, 0x2d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x69, 0x72, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x26, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x22, 0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x11, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22, 0x6d, 0x0a, 0x15, 0x52,
	0x6f, 0x74, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3f, 0x0a, 0x16, 0x52, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x44, 0x0a, 0x19, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77,
	0x74, 0x22, 0x1c, 0x0a, 0x1a, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x66, 0x69, 0x72, 0x61, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_fira_v1_app_requests_proto_rawDescOnce sync.Once
	file_protos_fira_v1_app_requests_proto_rawDescData = file_protos_fira_v1_app_requests_proto_rawDesc
)

func file_protos_fira_v1_app_requests_proto_rawDescGZIP() []byte {
	file_protos_fira_v1_app_requests_proto_rawDescOnce.Do(func() {
		file_protos_fira_v1_app_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_fira_v1_app_requests_proto_rawDescData)
	})
	return file_protos_fira_v1_app_requests_proto_rawDescData
}

var file_protos_fira_v1_app_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_fira_v1_app_requests_proto_goTypes = []interface{}{
	(*CreateAppRequest)(nil),           // 0: protos.fira.v1.CreateAppRequest
	(*CreateAppResponse)(nil),          // 1: protos.fira.v1.CreateAppResponse
	(*GetAppRequest)(nil),              // 2: protos.fira.v1.GetAppRequest
	(*GetAppResponse)(nil),             // 3: protos.fira.v1.GetAppResponse
	(*ListAppsRequest)(nil),            // 4: protos.fira.v1.ListAppsRequest
	(*ListAppsResponse)(nil),           // 5: protos.fira.v1.ListAppsResponse
	(*RotateAppTokenRequest)(nil),      // 6: protos.fira.v1.RotateAppTokenRequest
	(*RotateAppTokenResponse)(nil),     // 7: protos.fira.v1.RotateAppTokenResponse
	(*InvalidateAppTokenRequest)(nil),  // 8: protos.fira.v1.InvalidateAppTokenRequest
	(*InvalidateAppTokenResponse)(nil), // 9: protos.fira.v1.InvalidateAppTokenResponse
	(*App)(nil),                        // 10: protos.fira.v1.App
	(Environment)(0),                   // 11: protos.fira.v1.Environment
}
var file_protos_fira_v1_app_requests_proto_depIdxs = []int32{
	10, // 0: protos.fira.v1.CreateAppResponse.app:type_name -> protos.fira.v1.App
	10, // 1: protos.fira.v1.GetAppResponse.app:type_name -> protos.fira.v1.App
	10, // 2: protos.fira.v1.ListAppsResponse.apps:type_name -> protos.fira.v1.App
	11, // 3: protos.fira.v1.RotateAppTokenRequest.environment:type_name -> protos.fira.v1.Environment
	10, // 4: protos.fira.v1.RotateAppTokenResponse.app:type_name -> protos.fira.v1.App
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_protos_fira_v1_app_requests_proto_init() }
func file_protos_fira_v1_app_requests_proto_init() {
	if File_protos_fira_v1_app_requests_proto != nil {
		return
	}
	file_protos_fira_v1_app_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_fira_v1_app_requests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppRequest); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppResponse); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRequest); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppResponse); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppsRequest); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppsResponse); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotateAppTokenRequest); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotateAppTokenResponse); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidateAppTokenRequest); i {
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
		file_protos_fira_v1_app_requests_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidateAppTokenResponse); i {
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
			RawDescriptor: file_protos_fira_v1_app_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_fira_v1_app_requests_proto_goTypes,
		DependencyIndexes: file_protos_fira_v1_app_requests_proto_depIdxs,
		MessageInfos:      file_protos_fira_v1_app_requests_proto_msgTypes,
	}.Build()
	File_protos_fira_v1_app_requests_proto = out.File
	file_protos_fira_v1_app_requests_proto_rawDesc = nil
	file_protos_fira_v1_app_requests_proto_goTypes = nil
	file_protos_fira_v1_app_requests_proto_depIdxs = nil
}
