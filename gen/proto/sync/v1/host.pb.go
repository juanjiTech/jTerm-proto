// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: sync/v1/host.proto

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

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`               // 内部唯一ID ulid
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"` // 更新时间
	DeletedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"` // 删除时间
	Label      string                 `protobuf:"bytes,11,opt,name=label,proto3" json:"label,omitempty"`        // 标签
	Tags       []string               `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	Address    string                 `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"` // ip或host
	Port       int64                  `protobuf:"varint,14,opt,name=port,proto3" json:"port,omitempty"`
	Charset    string                 `protobuf:"bytes,15,opt,name=charset,proto3" json:"charset,omitempty"`       // 字符集
	Username   string                 `protobuf:"bytes,21,opt,name=username,proto3" json:"username,omitempty"`     // 用户名
	Password   string                 `protobuf:"bytes,22,opt,name=password,proto3" json:"password,omitempty"`     // 密码
	KeyId      string                 `protobuf:"bytes,23,opt,name=keyId,proto3" json:"keyId,omitempty"`           // 密钥ID
	IdentityId string                 `protobuf:"bytes,24,opt,name=identityId,proto3" json:"identityId,omitempty"` // 身份ID 若身份ID不为空，则覆盖手动配置
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_sync_v1_host_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Host) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Host) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Host) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Host) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Host) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Host) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Host) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Host) GetCharset() string {
	if x != nil {
		return x.Charset
	}
	return ""
}

func (x *Host) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Host) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Host) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *Host) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

var File_sync_v1_host_proto protoreflect.FileDescriptor

var file_sync_v1_host_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4,
	0x03, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x61, 0x6e, 0x6a, 0x69, 0x54, 0x65, 0x63, 0x68, 0x2f, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x75, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x79, 0x6e, 0x63, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_v1_host_proto_rawDescOnce sync.Once
	file_sync_v1_host_proto_rawDescData = file_sync_v1_host_proto_rawDesc
)

func file_sync_v1_host_proto_rawDescGZIP() []byte {
	file_sync_v1_host_proto_rawDescOnce.Do(func() {
		file_sync_v1_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_v1_host_proto_rawDescData)
	})
	return file_sync_v1_host_proto_rawDescData
}

var file_sync_v1_host_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sync_v1_host_proto_goTypes = []interface{}{
	(*Host)(nil),                  // 0: sync.v1.Host
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_sync_v1_host_proto_depIdxs = []int32{
	1, // 0: sync.v1.Host.createdAt:type_name -> google.protobuf.Timestamp
	1, // 1: sync.v1.Host.updatedAt:type_name -> google.protobuf.Timestamp
	1, // 2: sync.v1.Host.deletedAt:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sync_v1_host_proto_init() }
func file_sync_v1_host_proto_init() {
	if File_sync_v1_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_v1_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
			RawDescriptor: file_sync_v1_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_v1_host_proto_goTypes,
		DependencyIndexes: file_sync_v1_host_proto_depIdxs,
		MessageInfos:      file_sync_v1_host_proto_msgTypes,
	}.Build()
	File_sync_v1_host_proto = out.File
	file_sync_v1_host_proto_rawDesc = nil
	file_sync_v1_host_proto_goTypes = nil
	file_sync_v1_host_proto_depIdxs = nil
}
