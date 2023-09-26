// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sync/v1/sync_service.proto

package syncV1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"` // 获取该时间之后的设置变化，若不设置该字段，则拉取全量配置信息
	GroupId string                 `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_service_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRequest) GetAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.After
	}
	return nil
}

func (x *SyncRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTime   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"` // 当前服务器时间，用于给客户端下一次拉取变动时做一个参照，避免因客户端时间偏差导致同步混乱
	HostSet      []*Host                `protobuf:"bytes,11,rep,name=host_set,json=hostSet,proto3" json:"host_set,omitempty"`
	KnownHostSet []*KnownHost           `protobuf:"bytes,12,rep,name=known_host_set,json=knownHostSet,proto3" json:"known_host_set,omitempty"`
	SshKeySet    []*SshKey              `protobuf:"bytes,13,rep,name=ssh_key_set,json=sshKeySet,proto3" json:"ssh_key_set,omitempty"`
	IdentitySet  []*Identity            `protobuf:"bytes,14,rep,name=identity_set,json=identitySet,proto3" json:"identity_set,omitempty"`
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_service_proto_rawDescGZIP(), []int{1}
}

func (x *SyncResponse) GetServerTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ServerTime
	}
	return nil
}

func (x *SyncResponse) GetHostSet() []*Host {
	if x != nil {
		return x.HostSet
	}
	return nil
}

func (x *SyncResponse) GetKnownHostSet() []*KnownHost {
	if x != nil {
		return x.KnownHostSet
	}
	return nil
}

func (x *SyncResponse) GetSshKeySet() []*SshKey {
	if x != nil {
		return x.SshKeySet
	}
	return nil
}

func (x *SyncResponse) GetIdentitySet() []*Identity {
	if x != nil {
		return x.IdentitySet
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Types that are assignable to Data:
	//
	//	*UpdateRequest_Host
	//	*UpdateRequest_KnownHost
	//	*UpdateRequest_SshKey
	//	*UpdateRequest_Identity
	Data isUpdateRequest_Data `protobuf_oneof:"data"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (m *UpdateRequest) GetData() isUpdateRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UpdateRequest) GetHost() *Host {
	if x, ok := x.GetData().(*UpdateRequest_Host); ok {
		return x.Host
	}
	return nil
}

func (x *UpdateRequest) GetKnownHost() *KnownHost {
	if x, ok := x.GetData().(*UpdateRequest_KnownHost); ok {
		return x.KnownHost
	}
	return nil
}

func (x *UpdateRequest) GetSshKey() *SshKey {
	if x, ok := x.GetData().(*UpdateRequest_SshKey); ok {
		return x.SshKey
	}
	return nil
}

func (x *UpdateRequest) GetIdentity() *Identity {
	if x, ok := x.GetData().(*UpdateRequest_Identity); ok {
		return x.Identity
	}
	return nil
}

type isUpdateRequest_Data interface {
	isUpdateRequest_Data()
}

type UpdateRequest_Host struct {
	Host *Host `protobuf:"bytes,11,opt,name=host,proto3,oneof"`
}

type UpdateRequest_KnownHost struct {
	KnownHost *KnownHost `protobuf:"bytes,12,opt,name=known_host,json=knownHost,proto3,oneof"`
}

type UpdateRequest_SshKey struct {
	SshKey *SshKey `protobuf:"bytes,13,opt,name=ssh_key,json=sshKey,proto3,oneof"`
}

type UpdateRequest_Identity struct {
	Identity *Identity `protobuf:"bytes,14,opt,name=identity,proto3,oneof"`
}

func (*UpdateRequest_Host) isUpdateRequest_Data() {}

func (*UpdateRequest_KnownHost) isUpdateRequest_Data() {}

func (*UpdateRequest_SshKey) isUpdateRequest_Data() {}

func (*UpdateRequest_Identity) isUpdateRequest_Data() {}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UpdateResponse_Host
	//	*UpdateResponse_KnownHost
	//	*UpdateResponse_SshKey
	//	*UpdateResponse_Identity
	Data isUpdateResponse_Data `protobuf_oneof:"data"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_service_proto_rawDescGZIP(), []int{3}
}

func (m *UpdateResponse) GetData() isUpdateResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UpdateResponse) GetHost() *Host {
	if x, ok := x.GetData().(*UpdateResponse_Host); ok {
		return x.Host
	}
	return nil
}

func (x *UpdateResponse) GetKnownHost() *KnownHost {
	if x, ok := x.GetData().(*UpdateResponse_KnownHost); ok {
		return x.KnownHost
	}
	return nil
}

func (x *UpdateResponse) GetSshKey() *SshKey {
	if x, ok := x.GetData().(*UpdateResponse_SshKey); ok {
		return x.SshKey
	}
	return nil
}

func (x *UpdateResponse) GetIdentity() *Identity {
	if x, ok := x.GetData().(*UpdateResponse_Identity); ok {
		return x.Identity
	}
	return nil
}

type isUpdateResponse_Data interface {
	isUpdateResponse_Data()
}

type UpdateResponse_Host struct {
	Host *Host `protobuf:"bytes,11,opt,name=host,proto3,oneof"`
}

type UpdateResponse_KnownHost struct {
	KnownHost *KnownHost `protobuf:"bytes,12,opt,name=known_host,json=knownHost,proto3,oneof"`
}

type UpdateResponse_SshKey struct {
	SshKey *SshKey `protobuf:"bytes,13,opt,name=ssh_key,json=sshKey,proto3,oneof"`
}

type UpdateResponse_Identity struct {
	Identity *Identity `protobuf:"bytes,14,opt,name=identity,proto3,oneof"`
}

func (*UpdateResponse_Host) isUpdateResponse_Data() {}

func (*UpdateResponse_KnownHost) isUpdateResponse_Data() {}

func (*UpdateResponse_SshKey) isUpdateResponse_Data() {}

func (*UpdateResponse_Identity) isUpdateResponse_Data() {}

type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGroupRequest) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGroupResponse) Reset() {
	*x = UpdateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupResponse) ProtoMessage() {}

func (x *UpdateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateGroupResponse) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_service_proto_rawDescGZIP(), []int{5}
}

type SyncGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *SyncGroupRequest) Reset() {
	*x = SyncGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncGroupRequest) ProtoMessage() {}

func (x *SyncGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncGroupRequest.ProtoReflect.Descriptor instead.
func (*SyncGroupRequest) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_service_proto_rawDescGZIP(), []int{6}
}

func (x *SyncGroupRequest) GetAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.After
	}
	return nil
}

type SyncGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	Groups     []*Group               `protobuf:"bytes,11,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *SyncGroupResponse) Reset() {
	*x = SyncGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncGroupResponse) ProtoMessage() {}

func (x *SyncGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncGroupResponse.ProtoReflect.Descriptor instead.
func (*SyncGroupResponse) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_service_proto_rawDescGZIP(), []int{7}
}

func (x *SyncGroupResponse) GetServerTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ServerTime
	}
	return nil
}

func (x *SyncGroupResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_sync_v1_sync_service_proto protoreflect.FileDescriptor

var file_sync_v1_sync_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x65, 0x79, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x79, 0x6e,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5a, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x96, 0x02, 0x0a,
	0x0c, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x07, 0x68, 0x6f, 0x73,
	0x74, 0x53, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x0e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x0c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x74, 0x12, 0x2f,
	0x0a, 0x0b, 0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x73,
	0x68, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x12,
	0x34, 0x0a, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x53, 0x65, 0x74, 0x22, 0xe9, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07,
	0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x6f, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x07, 0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x73, 0x68, 0x4b, 0x65, 0x79,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48,
	0x00, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22,
	0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x11,
	0x53, 0x79, 0x6e, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x32, 0x91, 0x03, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x14,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x5a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f,
	0x67, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x6f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x1b, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x64, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x12, 0x18, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x61, 0x6e, 0x6a, 0x69, 0x54,
	0x65, 0x63, 0x68, 0x2f, 0x6a, 0x54, 0x65, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x79, 0x6e, 0x63, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_v1_sync_service_proto_rawDescOnce sync.Once
	file_sync_v1_sync_service_proto_rawDescData = file_sync_v1_sync_service_proto_rawDesc
)

func file_sync_v1_sync_service_proto_rawDescGZIP() []byte {
	file_sync_v1_sync_service_proto_rawDescOnce.Do(func() {
		file_sync_v1_sync_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_v1_sync_service_proto_rawDescData)
	})
	return file_sync_v1_sync_service_proto_rawDescData
}

var file_sync_v1_sync_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sync_v1_sync_service_proto_goTypes = []interface{}{
	(*SyncRequest)(nil),           // 0: sync.v1.SyncRequest
	(*SyncResponse)(nil),          // 1: sync.v1.SyncResponse
	(*UpdateRequest)(nil),         // 2: sync.v1.UpdateRequest
	(*UpdateResponse)(nil),        // 3: sync.v1.UpdateResponse
	(*UpdateGroupRequest)(nil),    // 4: sync.v1.UpdateGroupRequest
	(*UpdateGroupResponse)(nil),   // 5: sync.v1.UpdateGroupResponse
	(*SyncGroupRequest)(nil),      // 6: sync.v1.SyncGroupRequest
	(*SyncGroupResponse)(nil),     // 7: sync.v1.SyncGroupResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*Host)(nil),                  // 9: sync.v1.Host
	(*KnownHost)(nil),             // 10: sync.v1.KnownHost
	(*SshKey)(nil),                // 11: sync.v1.SshKey
	(*Identity)(nil),              // 12: sync.v1.Identity
	(*Group)(nil),                 // 13: sync.v1.Group
}
var file_sync_v1_sync_service_proto_depIdxs = []int32{
	8,  // 0: sync.v1.SyncRequest.after:type_name -> google.protobuf.Timestamp
	8,  // 1: sync.v1.SyncResponse.server_time:type_name -> google.protobuf.Timestamp
	9,  // 2: sync.v1.SyncResponse.host_set:type_name -> sync.v1.Host
	10, // 3: sync.v1.SyncResponse.known_host_set:type_name -> sync.v1.KnownHost
	11, // 4: sync.v1.SyncResponse.ssh_key_set:type_name -> sync.v1.SshKey
	12, // 5: sync.v1.SyncResponse.identity_set:type_name -> sync.v1.Identity
	9,  // 6: sync.v1.UpdateRequest.host:type_name -> sync.v1.Host
	10, // 7: sync.v1.UpdateRequest.known_host:type_name -> sync.v1.KnownHost
	11, // 8: sync.v1.UpdateRequest.ssh_key:type_name -> sync.v1.SshKey
	12, // 9: sync.v1.UpdateRequest.identity:type_name -> sync.v1.Identity
	9,  // 10: sync.v1.UpdateResponse.host:type_name -> sync.v1.Host
	10, // 11: sync.v1.UpdateResponse.known_host:type_name -> sync.v1.KnownHost
	11, // 12: sync.v1.UpdateResponse.ssh_key:type_name -> sync.v1.SshKey
	12, // 13: sync.v1.UpdateResponse.identity:type_name -> sync.v1.Identity
	13, // 14: sync.v1.UpdateGroupRequest.group:type_name -> sync.v1.Group
	8,  // 15: sync.v1.SyncGroupRequest.after:type_name -> google.protobuf.Timestamp
	8,  // 16: sync.v1.SyncGroupResponse.server_time:type_name -> google.protobuf.Timestamp
	13, // 17: sync.v1.SyncGroupResponse.groups:type_name -> sync.v1.Group
	0,  // 18: sync.v1.SyncService.Sync:input_type -> sync.v1.SyncRequest
	2,  // 19: sync.v1.SyncService.Update:input_type -> sync.v1.UpdateRequest
	4,  // 20: sync.v1.SyncService.UpdateGroup:input_type -> sync.v1.UpdateGroupRequest
	6,  // 21: sync.v1.SyncService.SyncGroup:input_type -> sync.v1.SyncGroupRequest
	1,  // 22: sync.v1.SyncService.Sync:output_type -> sync.v1.SyncResponse
	3,  // 23: sync.v1.SyncService.Update:output_type -> sync.v1.UpdateResponse
	5,  // 24: sync.v1.SyncService.UpdateGroup:output_type -> sync.v1.UpdateGroupResponse
	7,  // 25: sync.v1.SyncService.SyncGroup:output_type -> sync.v1.SyncGroupResponse
	22, // [22:26] is the sub-list for method output_type
	18, // [18:22] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_sync_v1_sync_service_proto_init() }
func file_sync_v1_sync_service_proto_init() {
	if File_sync_v1_sync_service_proto != nil {
		return
	}
	file_sync_v1_host_proto_init()
	file_sync_v1_keychain_proto_init()
	file_sync_v1_known_hosts_proto_init()
	file_sync_v1_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sync_v1_sync_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_sync_v1_sync_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
		file_sync_v1_sync_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_sync_v1_sync_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_sync_v1_sync_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
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
		file_sync_v1_sync_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupResponse); i {
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
		file_sync_v1_sync_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncGroupRequest); i {
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
		file_sync_v1_sync_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncGroupResponse); i {
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
	file_sync_v1_sync_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*UpdateRequest_Host)(nil),
		(*UpdateRequest_KnownHost)(nil),
		(*UpdateRequest_SshKey)(nil),
		(*UpdateRequest_Identity)(nil),
	}
	file_sync_v1_sync_service_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*UpdateResponse_Host)(nil),
		(*UpdateResponse_KnownHost)(nil),
		(*UpdateResponse_SshKey)(nil),
		(*UpdateResponse_Identity)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sync_v1_sync_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sync_v1_sync_service_proto_goTypes,
		DependencyIndexes: file_sync_v1_sync_service_proto_depIdxs,
		MessageInfos:      file_sync_v1_sync_service_proto_msgTypes,
	}.Build()
	File_sync_v1_sync_service_proto = out.File
	file_sync_v1_sync_service_proto_rawDesc = nil
	file_sync_v1_sync_service_proto_goTypes = nil
	file_sync_v1_sync_service_proto_depIdxs = nil
}
