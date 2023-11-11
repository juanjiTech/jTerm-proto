// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sync/v1/group.proto

package syncV1

import (
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

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupBasic    *GroupBasic      `protobuf:"bytes,1,opt,name=group_basic,json=groupBasic,proto3" json:"group_basic,omitempty"`
	GroupKeychain []*GroupKeyChain `protobuf:"bytes,2,rep,name=group_keychain,json=groupKeychain,proto3" json:"group_keychain,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_sync_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetGroupBasic() *GroupBasic {
	if x != nil {
		return x.GroupBasic
	}
	return nil
}

func (x *Group) GetGroupKeychain() []*GroupKeyChain {
	if x != nil {
		return x.GroupKeychain
	}
	return nil
}

type GroupBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Name        string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Avatar      string                 `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Uid         string                 `protobuf:"bytes,8,opt,name=uid,proto3" json:"uid,omitempty"`             // 创建者
	PublicKey   string                 `protobuf:"bytes,9,opt,name=publicKey,proto3" json:"publicKey,omitempty"` // 公钥,公钥存在服务器上，加密的私钥与相关用户关联存储
}

func (x *GroupBasic) Reset() {
	*x = GroupBasic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupBasic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupBasic) ProtoMessage() {}

func (x *GroupBasic) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupBasic.ProtoReflect.Descriptor instead.
func (*GroupBasic) Descriptor() ([]byte, []int) {
	return file_sync_v1_group_proto_rawDescGZIP(), []int{1}
}

func (x *GroupBasic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroupBasic) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GroupBasic) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *GroupBasic) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *GroupBasic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupBasic) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GroupBasic) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GroupBasic) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GroupBasic) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type GroupKeyChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId   string `protobuf:"bytes,11,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	PrivateKey string `protobuf:"bytes,12,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
}

func (x *GroupKeyChain) Reset() {
	*x = GroupKeyChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupKeyChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupKeyChain) ProtoMessage() {}

func (x *GroupKeyChain) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupKeyChain.ProtoReflect.Descriptor instead.
func (*GroupKeyChain) Descriptor() ([]byte, []int) {
	return file_sync_v1_group_proto_rawDescGZIP(), []int{2}
}

func (x *GroupKeyChain) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *GroupKeyChain) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

var File_sync_v1_group_proto protoreflect.FileDescriptor

var file_sync_v1_group_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7c, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x34, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x3d,
	0x0a, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x0d,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0xcb, 0x02,
	0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x4c, 0x0a, 0x0d, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x61, 0x6e, 0x6a, 0x69, 0x54, 0x65,
	0x63, 0x68, 0x2f, 0x6a, 0x54, 0x65, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x79, 0x6e, 0x63, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_v1_group_proto_rawDescOnce sync.Once
	file_sync_v1_group_proto_rawDescData = file_sync_v1_group_proto_rawDesc
)

func file_sync_v1_group_proto_rawDescGZIP() []byte {
	file_sync_v1_group_proto_rawDescOnce.Do(func() {
		file_sync_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_v1_group_proto_rawDescData)
	})
	return file_sync_v1_group_proto_rawDescData
}

var file_sync_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sync_v1_group_proto_goTypes = []interface{}{
	(*Group)(nil),                 // 0: sync.v1.Group
	(*GroupBasic)(nil),            // 1: sync.v1.GroupBasic
	(*GroupKeyChain)(nil),         // 2: sync.v1.GroupKeyChain
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_sync_v1_group_proto_depIdxs = []int32{
	1, // 0: sync.v1.Group.group_basic:type_name -> sync.v1.GroupBasic
	2, // 1: sync.v1.Group.group_keychain:type_name -> sync.v1.GroupKeyChain
	3, // 2: sync.v1.GroupBasic.created_at:type_name -> google.protobuf.Timestamp
	3, // 3: sync.v1.GroupBasic.updated_at:type_name -> google.protobuf.Timestamp
	3, // 4: sync.v1.GroupBasic.deleted_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sync_v1_group_proto_init() }
func file_sync_v1_group_proto_init() {
	if File_sync_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_v1_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_sync_v1_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupBasic); i {
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
		file_sync_v1_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupKeyChain); i {
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
			RawDescriptor: file_sync_v1_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_v1_group_proto_goTypes,
		DependencyIndexes: file_sync_v1_group_proto_depIdxs,
		MessageInfos:      file_sync_v1_group_proto_msgTypes,
	}.Build()
	File_sync_v1_group_proto = out.File
	file_sync_v1_group_proto_rawDesc = nil
	file_sync_v1_group_proto_goTypes = nil
	file_sync_v1_group_proto_depIdxs = nil
}
