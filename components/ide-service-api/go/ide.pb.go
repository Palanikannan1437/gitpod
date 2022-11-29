// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: ide.proto

package api

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

// WorkspaceType specifies the purpose/use of a workspace. Different workspace types are handled differently by all parts of the system.
// copied from https://github.com/gitpod-io/gitpod/blob/a7f35378326ca5ec41aab1a48418949070a9b37a/components/ws-manager-api/core.proto#L660-L675
type WorkspaceType int32

const (
	WorkspaceType_REGULAR    WorkspaceType = 0
	WorkspaceType_PREBUILD   WorkspaceType = 1
	WorkspaceType_IMAGEBUILD WorkspaceType = 4
)

// Enum value maps for WorkspaceType.
var (
	WorkspaceType_name = map[int32]string{
		0: "REGULAR",
		1: "PREBUILD",
		4: "IMAGEBUILD",
	}
	WorkspaceType_value = map[string]int32{
		"REGULAR":    0,
		"PREBUILD":   1,
		"IMAGEBUILD": 4,
	}
)

func (x WorkspaceType) Enum() *WorkspaceType {
	p := new(WorkspaceType)
	*p = x
	return p
}

func (x WorkspaceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceType) Descriptor() protoreflect.EnumDescriptor {
	return file_ide_proto_enumTypes[0].Descriptor()
}

func (WorkspaceType) Type() protoreflect.EnumType {
	return &file_ide_proto_enumTypes[0]
}

func (x WorkspaceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceType.Descriptor instead.
func (WorkspaceType) EnumDescriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{0}
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{0}
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfigResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// EnvironmentVariable describes an env var as key/value pair
type EnvironmentVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EnvironmentVariable) Reset() {
	*x = EnvironmentVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentVariable) ProtoMessage() {}

func (x *EnvironmentVariable) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentVariable.ProtoReflect.Descriptor instead.
func (*EnvironmentVariable) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{2}
}

func (x *EnvironmentVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvironmentVariable) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ResolveStartWorkspaceSpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        WorkspaceType `protobuf:"varint,1,opt,name=type,proto3,enum=ide_service_api.WorkspaceType" json:"type,omitempty"`
	Context     string        `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	UserSetting string        `protobuf:"bytes,3,opt,name=user_setting,json=userSetting,proto3" json:"user_setting,omitempty"`
}

func (x *ResolveStartWorkspaceSpecRequest) Reset() {
	*x = ResolveStartWorkspaceSpecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveStartWorkspaceSpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveStartWorkspaceSpecRequest) ProtoMessage() {}

func (x *ResolveStartWorkspaceSpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveStartWorkspaceSpecRequest.ProtoReflect.Descriptor instead.
func (*ResolveStartWorkspaceSpecRequest) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveStartWorkspaceSpecRequest) GetType() WorkspaceType {
	if x != nil {
		return x.Type
	}
	return WorkspaceType_REGULAR
}

func (x *ResolveStartWorkspaceSpecRequest) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *ResolveStartWorkspaceSpecRequest) GetUserSetting() string {
	if x != nil {
		return x.UserSetting
	}
	return ""
}

type ResolveStartWorkspaceSpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envvars         []*EnvironmentVariable `protobuf:"bytes,1,rep,name=envvars,proto3" json:"envvars,omitempty"`
	SupervisorImage string                 `protobuf:"bytes,2,opt,name=supervisor_image,json=supervisorImage,proto3" json:"supervisor_image,omitempty"`
	WebImage        string                 `protobuf:"bytes,3,opt,name=web_image,json=webImage,proto3" json:"web_image,omitempty"`
	IdeImageLayers  []string               `protobuf:"bytes,4,rep,name=ide_image_layers,json=ideImageLayers,proto3" json:"ide_image_layers,omitempty"`
}

func (x *ResolveStartWorkspaceSpecResponse) Reset() {
	*x = ResolveStartWorkspaceSpecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveStartWorkspaceSpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveStartWorkspaceSpecResponse) ProtoMessage() {}

func (x *ResolveStartWorkspaceSpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveStartWorkspaceSpecResponse.ProtoReflect.Descriptor instead.
func (*ResolveStartWorkspaceSpecResponse) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{4}
}

func (x *ResolveStartWorkspaceSpecResponse) GetEnvvars() []*EnvironmentVariable {
	if x != nil {
		return x.Envvars
	}
	return nil
}

func (x *ResolveStartWorkspaceSpecResponse) GetSupervisorImage() string {
	if x != nil {
		return x.SupervisorImage
	}
	return ""
}

func (x *ResolveStartWorkspaceSpecResponse) GetWebImage() string {
	if x != nil {
		return x.WebImage
	}
	return ""
}

func (x *ResolveStartWorkspaceSpecResponse) GetIdeImageLayers() []string {
	if x != nil {
		return x.IdeImageLayers
	}
	return nil
}

var File_ide_proto protoreflect.FileDescriptor

var file_ide_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x69, 0x64, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x3f, 0x0a, 0x13, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x93, 0x01, 0x0a, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xd5, 0x01, 0x0a, 0x21, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07,
	0x65, 0x6e, 0x76, 0x76, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x76, 0x61, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x10,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73,
	0x6f, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x62, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x62, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x64, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2a, 0x3a,
	0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x50, 0x52, 0x45, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d,
	0x41, 0x47, 0x45, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x04, 0x32, 0xe9, 0x01, 0x0a, 0x0a, 0x49,
	0x44, 0x45, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x2e, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x64, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x84, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x31, 0x2e,
	0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x70, 0x6f, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f,
	0x69, 0x64, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ide_proto_rawDescOnce sync.Once
	file_ide_proto_rawDescData = file_ide_proto_rawDesc
)

func file_ide_proto_rawDescGZIP() []byte {
	file_ide_proto_rawDescOnce.Do(func() {
		file_ide_proto_rawDescData = protoimpl.X.CompressGZIP(file_ide_proto_rawDescData)
	})
	return file_ide_proto_rawDescData
}

var file_ide_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ide_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ide_proto_goTypes = []interface{}{
	(WorkspaceType)(0),                        // 0: ide_service_api.WorkspaceType
	(*GetConfigRequest)(nil),                  // 1: ide_service_api.GetConfigRequest
	(*GetConfigResponse)(nil),                 // 2: ide_service_api.GetConfigResponse
	(*EnvironmentVariable)(nil),               // 3: ide_service_api.EnvironmentVariable
	(*ResolveStartWorkspaceSpecRequest)(nil),  // 4: ide_service_api.ResolveStartWorkspaceSpecRequest
	(*ResolveStartWorkspaceSpecResponse)(nil), // 5: ide_service_api.ResolveStartWorkspaceSpecResponse
}
var file_ide_proto_depIdxs = []int32{
	0, // 0: ide_service_api.ResolveStartWorkspaceSpecRequest.type:type_name -> ide_service_api.WorkspaceType
	3, // 1: ide_service_api.ResolveStartWorkspaceSpecResponse.envvars:type_name -> ide_service_api.EnvironmentVariable
	1, // 2: ide_service_api.IDEService.GetConfig:input_type -> ide_service_api.GetConfigRequest
	4, // 3: ide_service_api.IDEService.ResolveStartWorkspaceSpec:input_type -> ide_service_api.ResolveStartWorkspaceSpecRequest
	2, // 4: ide_service_api.IDEService.GetConfig:output_type -> ide_service_api.GetConfigResponse
	5, // 5: ide_service_api.IDEService.ResolveStartWorkspaceSpec:output_type -> ide_service_api.ResolveStartWorkspaceSpecResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ide_proto_init() }
func file_ide_proto_init() {
	if File_ide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_ide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigResponse); i {
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
		file_ide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentVariable); i {
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
		file_ide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveStartWorkspaceSpecRequest); i {
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
		file_ide_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveStartWorkspaceSpecResponse); i {
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
			RawDescriptor: file_ide_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ide_proto_goTypes,
		DependencyIndexes: file_ide_proto_depIdxs,
		EnumInfos:         file_ide_proto_enumTypes,
		MessageInfos:      file_ide_proto_msgTypes,
	}.Build()
	File_ide_proto = out.File
	file_ide_proto_rawDesc = nil
	file_ide_proto_goTypes = nil
	file_ide_proto_depIdxs = nil
}
