// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: github.com/openconfig/gnsi/pathz/pathz.proto

package pathz

import (
	gnmi "github.com/openconfig/gnmi/proto/gnmi"
	_ "github.com/openconfig/gnsi/version"
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

type PolicyInstance int32

const (
	PolicyInstance_POLICY_INSTANCE_UNSPECIFIED PolicyInstance = 0
	PolicyInstance_POLICY_INSTANCE_ACTIVE      PolicyInstance = 1
	PolicyInstance_POLICY_INSTANCE_SANDBOX     PolicyInstance = 2
)

// Enum value maps for PolicyInstance.
var (
	PolicyInstance_name = map[int32]string{
		0: "POLICY_INSTANCE_UNSPECIFIED",
		1: "POLICY_INSTANCE_ACTIVE",
		2: "POLICY_INSTANCE_SANDBOX",
	}
	PolicyInstance_value = map[string]int32{
		"POLICY_INSTANCE_UNSPECIFIED": 0,
		"POLICY_INSTANCE_ACTIVE":      1,
		"POLICY_INSTANCE_SANDBOX":     2,
	}
)

func (x PolicyInstance) Enum() *PolicyInstance {
	p := new(PolicyInstance)
	*p = x
	return p
}

func (x PolicyInstance) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyInstance) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_enumTypes[0].Descriptor()
}

func (PolicyInstance) Type() protoreflect.EnumType {
	return &file_github_com_openconfig_gnsi_pathz_pathz_proto_enumTypes[0]
}

func (x PolicyInstance) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyInstance.Descriptor instead.
func (PolicyInstance) EnumDescriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{0}
}

type RotateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RotateRequest:
	//	*RotateRequest_UploadRequest
	//	*RotateRequest_FinalizeRotation
	RotateRequest  isRotateRequest_RotateRequest `protobuf_oneof:"rotate_request"`
	ForceOverwrite bool                          `protobuf:"varint,3,opt,name=force_overwrite,json=forceOverwrite,proto3" json:"force_overwrite,omitempty"`
}

func (x *RotateRequest) Reset() {
	*x = RotateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateRequest) ProtoMessage() {}

func (x *RotateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateRequest.ProtoReflect.Descriptor instead.
func (*RotateRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{0}
}

func (m *RotateRequest) GetRotateRequest() isRotateRequest_RotateRequest {
	if m != nil {
		return m.RotateRequest
	}
	return nil
}

func (x *RotateRequest) GetUploadRequest() *UploadRequest {
	if x, ok := x.GetRotateRequest().(*RotateRequest_UploadRequest); ok {
		return x.UploadRequest
	}
	return nil
}

func (x *RotateRequest) GetFinalizeRotation() *FinalizeRequest {
	if x, ok := x.GetRotateRequest().(*RotateRequest_FinalizeRotation); ok {
		return x.FinalizeRotation
	}
	return nil
}

func (x *RotateRequest) GetForceOverwrite() bool {
	if x != nil {
		return x.ForceOverwrite
	}
	return false
}

type isRotateRequest_RotateRequest interface {
	isRotateRequest_RotateRequest()
}

type RotateRequest_UploadRequest struct {
	UploadRequest *UploadRequest `protobuf:"bytes,1,opt,name=upload_request,json=uploadRequest,proto3,oneof"`
}

type RotateRequest_FinalizeRotation struct {
	FinalizeRotation *FinalizeRequest `protobuf:"bytes,2,opt,name=finalize_rotation,json=finalizeRotation,proto3,oneof"`
}

func (*RotateRequest_UploadRequest) isRotateRequest_RotateRequest() {}

func (*RotateRequest_FinalizeRotation) isRotateRequest_RotateRequest() {}

type RotateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*RotateResponse_Upload
	Response isRotateResponse_Response `protobuf_oneof:"response"`
}

func (x *RotateResponse) Reset() {
	*x = RotateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateResponse) ProtoMessage() {}

func (x *RotateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateResponse.ProtoReflect.Descriptor instead.
func (*RotateResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{1}
}

func (m *RotateResponse) GetResponse() isRotateResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *RotateResponse) GetUpload() *UploadResponse {
	if x, ok := x.GetResponse().(*RotateResponse_Upload); ok {
		return x.Upload
	}
	return nil
}

type isRotateResponse_Response interface {
	isRotateResponse_Response()
}

type RotateResponse_Upload struct {
	Upload *UploadResponse `protobuf:"bytes,1,opt,name=upload,proto3,oneof"`
}

func (*RotateResponse_Upload) isRotateResponse_Response() {}

type FinalizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinalizeRequest) Reset() {
	*x = FinalizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeRequest) ProtoMessage() {}

func (x *FinalizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeRequest.ProtoReflect.Descriptor instead.
func (*FinalizeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{2}
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string               `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CreatedOn uint64               `protobuf:"varint,2,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	Policy    *AuthorizationPolicy `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{3}
}

func (x *UploadRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UploadRequest) GetCreatedOn() uint64 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

func (x *UploadRequest) GetPolicy() *AuthorizationPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{4}
}

type ProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User           string         `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Path           *gnmi.Path     `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Mode           Mode           `protobuf:"varint,3,opt,name=mode,proto3,enum=gnsi.pathz.Mode" json:"mode,omitempty"`
	PolicyInstance PolicyInstance `protobuf:"varint,4,opt,name=policy_instance,json=policyInstance,proto3,enum=gnsi.pathz.PolicyInstance" json:"policy_instance,omitempty"`
}

func (x *ProbeRequest) Reset() {
	*x = ProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeRequest) ProtoMessage() {}

func (x *ProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeRequest.ProtoReflect.Descriptor instead.
func (*ProbeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{5}
}

func (x *ProbeRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ProbeRequest) GetPath() *gnmi.Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *ProbeRequest) GetMode() Mode {
	if x != nil {
		return x.Mode
	}
	return Mode_MODE_UNSPECIFIED
}

func (x *ProbeRequest) GetPolicyInstance() PolicyInstance {
	if x != nil {
		return x.PolicyInstance
	}
	return PolicyInstance_POLICY_INSTANCE_UNSPECIFIED
}

type ProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action  Action `protobuf:"varint,1,opt,name=action,proto3,enum=gnsi.pathz.Action" json:"action,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ProbeResponse) Reset() {
	*x = ProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeResponse) ProtoMessage() {}

func (x *ProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeResponse.ProtoReflect.Descriptor instead.
func (*ProbeResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{6}
}

func (x *ProbeResponse) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

func (x *ProbeResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyInstance PolicyInstance `protobuf:"varint,1,opt,name=policy_instance,json=policyInstance,proto3,enum=gnsi.pathz.PolicyInstance" json:"policy_instance,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{7}
}

func (x *GetRequest) GetPolicyInstance() PolicyInstance {
	if x != nil {
		return x.PolicyInstance
	}
	return PolicyInstance_POLICY_INSTANCE_UNSPECIFIED
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string               `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CreatedOn uint64               `protobuf:"varint,2,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	Policy    *AuthorizationPolicy `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP(), []int{8}
}

func (x *GetResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetResponse) GetCreatedOn() uint64 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

func (x *GetResponse) GetPolicy() *AuthorizationPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_github_com_openconfig_gnsi_pathz_pathz_proto protoreflect.FileDescriptor

var file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x73, 0x69, 0x2f, 0x70, 0x61, 0x74,
	0x68, 0x7a, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x1a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x67, 0x6e, 0x6d, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6e, 0x6d,
	0x69, 0x2f, 0x67, 0x6e, 0x6d, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x73, 0x69, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x73, 0x69, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x7a,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x11, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74,
	0x68, 0x7a, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f,
	0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42,
	0x10, 0x0a, 0x0e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x52, 0x0a, 0x0e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x4f, 0x6e, 0x12, 0x37, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x10, 0x0a, 0x0e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xad,
	0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x6e, 0x6d, 0x69, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x55,
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67,
	0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x7f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e,
	0x12, 0x37, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2a, 0x6a, 0x0a, 0x0e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x50,
	0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4f, 0x4c, 0x49,
	0x43, 0x59, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x41, 0x4e, 0x44,
	0x42, 0x4f, 0x58, 0x10, 0x02, 0x32, 0xc2, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x68, 0x7a, 0x12,
	0x43, 0x0a, 0x06, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x67, 0x6e, 0x73, 0x69,
	0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68,
	0x7a, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x18, 0x2e,
	0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70,
	0x61, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6e, 0x73, 0x69,
	0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x73, 0x69, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x7a, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescOnce sync.Once
	file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescData = file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDesc
)

func file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescGZIP() []byte {
	file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescOnce.Do(func() {
		file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescData)
	})
	return file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDescData
}

var file_github_com_openconfig_gnsi_pathz_pathz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_openconfig_gnsi_pathz_pathz_proto_goTypes = []interface{}{
	(PolicyInstance)(0),         // 0: gnsi.pathz.PolicyInstance
	(*RotateRequest)(nil),       // 1: gnsi.pathz.RotateRequest
	(*RotateResponse)(nil),      // 2: gnsi.pathz.RotateResponse
	(*FinalizeRequest)(nil),     // 3: gnsi.pathz.FinalizeRequest
	(*UploadRequest)(nil),       // 4: gnsi.pathz.UploadRequest
	(*UploadResponse)(nil),      // 5: gnsi.pathz.UploadResponse
	(*ProbeRequest)(nil),        // 6: gnsi.pathz.ProbeRequest
	(*ProbeResponse)(nil),       // 7: gnsi.pathz.ProbeResponse
	(*GetRequest)(nil),          // 8: gnsi.pathz.GetRequest
	(*GetResponse)(nil),         // 9: gnsi.pathz.GetResponse
	(*AuthorizationPolicy)(nil), // 10: gnsi.pathz.AuthorizationPolicy
	(*gnmi.Path)(nil),           // 11: gnmi.Path
	(Mode)(0),                   // 12: gnsi.pathz.Mode
	(Action)(0),                 // 13: gnsi.pathz.Action
}
var file_github_com_openconfig_gnsi_pathz_pathz_proto_depIdxs = []int32{
	4,  // 0: gnsi.pathz.RotateRequest.upload_request:type_name -> gnsi.pathz.UploadRequest
	3,  // 1: gnsi.pathz.RotateRequest.finalize_rotation:type_name -> gnsi.pathz.FinalizeRequest
	5,  // 2: gnsi.pathz.RotateResponse.upload:type_name -> gnsi.pathz.UploadResponse
	10, // 3: gnsi.pathz.UploadRequest.policy:type_name -> gnsi.pathz.AuthorizationPolicy
	11, // 4: gnsi.pathz.ProbeRequest.path:type_name -> gnmi.Path
	12, // 5: gnsi.pathz.ProbeRequest.mode:type_name -> gnsi.pathz.Mode
	0,  // 6: gnsi.pathz.ProbeRequest.policy_instance:type_name -> gnsi.pathz.PolicyInstance
	13, // 7: gnsi.pathz.ProbeResponse.action:type_name -> gnsi.pathz.Action
	0,  // 8: gnsi.pathz.GetRequest.policy_instance:type_name -> gnsi.pathz.PolicyInstance
	10, // 9: gnsi.pathz.GetResponse.policy:type_name -> gnsi.pathz.AuthorizationPolicy
	1,  // 10: gnsi.pathz.Pathz.Rotate:input_type -> gnsi.pathz.RotateRequest
	6,  // 11: gnsi.pathz.Pathz.Probe:input_type -> gnsi.pathz.ProbeRequest
	8,  // 12: gnsi.pathz.Pathz.Get:input_type -> gnsi.pathz.GetRequest
	2,  // 13: gnsi.pathz.Pathz.Rotate:output_type -> gnsi.pathz.RotateResponse
	7,  // 14: gnsi.pathz.Pathz.Probe:output_type -> gnsi.pathz.ProbeResponse
	9,  // 15: gnsi.pathz.Pathz.Get:output_type -> gnsi.pathz.GetResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_github_com_openconfig_gnsi_pathz_pathz_proto_init() }
func file_github_com_openconfig_gnsi_pathz_pathz_proto_init() {
	if File_github_com_openconfig_gnsi_pathz_pathz_proto != nil {
		return
	}
	file_github_com_openconfig_gnsi_pathz_authorization_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotateRequest); i {
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
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotateResponse); i {
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
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeRequest); i {
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
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeRequest); i {
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
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeResponse); i {
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
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
	file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RotateRequest_UploadRequest)(nil),
		(*RotateRequest_FinalizeRotation)(nil),
	}
	file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RotateResponse_Upload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_openconfig_gnsi_pathz_pathz_proto_goTypes,
		DependencyIndexes: file_github_com_openconfig_gnsi_pathz_pathz_proto_depIdxs,
		EnumInfos:         file_github_com_openconfig_gnsi_pathz_pathz_proto_enumTypes,
		MessageInfos:      file_github_com_openconfig_gnsi_pathz_pathz_proto_msgTypes,
	}.Build()
	File_github_com_openconfig_gnsi_pathz_pathz_proto = out.File
	file_github_com_openconfig_gnsi_pathz_pathz_proto_rawDesc = nil
	file_github_com_openconfig_gnsi_pathz_pathz_proto_goTypes = nil
	file_github_com_openconfig_gnsi_pathz_pathz_proto_depIdxs = nil
}
