// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sync/v1/sync_service.proto

package syncV1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SyncService_Sync_FullMethodName               = "/sync.v1.SyncService/Sync"
	SyncService_Update_FullMethodName             = "/sync.v1.SyncService/Update"
	SyncService_UpdateGroup_FullMethodName        = "/sync.v1.SyncService/UpdateGroup"
	SyncService_SyncGroup_FullMethodName          = "/sync.v1.SyncService/SyncGroup"
	SyncService_GenUserKeyChain_FullMethodName    = "/sync.v1.SyncService/GenUserKeyChain"
	SyncService_GetEncryptedKey_FullMethodName    = "/sync.v1.SyncService/GetEncryptedKey"
	SyncService_UpdateEncryptedKey_FullMethodName = "/sync.v1.SyncService/UpdateEncryptedKey"
)

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncServiceClient interface {
	// 拉取指定时间点之后的配置变动信息
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	// 提交最新配置 若配置的ID为空，则创建新配置。若配置的删除时间不为空，则代表该配置已被删除。
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// 更新组信息 如果创建一个新的组，需要同时处理生成新的密钥
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error)
	// 通过UID获取所有组信息
	SyncGroup(ctx context.Context, in *SyncGroupRequest, opts ...grpc.CallOption) (*SyncGroupResponse, error)
	// 生成用户密钥对
	GenUserKeyChain(ctx context.Context, in *GenUserKeyChainRequest, opts ...grpc.CallOption) (*GenUserKeyChainResponse, error)
	// 获取mater key加密后的解密密钥,对应组
	GetEncryptedKey(ctx context.Context, in *GetEncryptedKeyRequest, opts ...grpc.CallOption) (*GetEncryptedKeyResponse, error)
	// 更新mater key后修改加密后密钥,提交多组上去，将全部相关信息修改，不能有缺失。
	UpdateEncryptedKey(ctx context.Context, in *UpdateEncryptedKeyRequest, opts ...grpc.CallOption) (*UpdateEncryptedKeyResponse, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, SyncService_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, SyncService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error) {
	out := new(UpdateGroupResponse)
	err := c.cc.Invoke(ctx, SyncService_UpdateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SyncGroup(ctx context.Context, in *SyncGroupRequest, opts ...grpc.CallOption) (*SyncGroupResponse, error) {
	out := new(SyncGroupResponse)
	err := c.cc.Invoke(ctx, SyncService_SyncGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GenUserKeyChain(ctx context.Context, in *GenUserKeyChainRequest, opts ...grpc.CallOption) (*GenUserKeyChainResponse, error) {
	out := new(GenUserKeyChainResponse)
	err := c.cc.Invoke(ctx, SyncService_GenUserKeyChain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetEncryptedKey(ctx context.Context, in *GetEncryptedKeyRequest, opts ...grpc.CallOption) (*GetEncryptedKeyResponse, error) {
	out := new(GetEncryptedKeyResponse)
	err := c.cc.Invoke(ctx, SyncService_GetEncryptedKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) UpdateEncryptedKey(ctx context.Context, in *UpdateEncryptedKeyRequest, opts ...grpc.CallOption) (*UpdateEncryptedKeyResponse, error) {
	out := new(UpdateEncryptedKeyResponse)
	err := c.cc.Invoke(ctx, SyncService_UpdateEncryptedKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServiceServer is the server API for SyncService service.
// All implementations must embed UnimplementedSyncServiceServer
// for forward compatibility
type SyncServiceServer interface {
	// 拉取指定时间点之后的配置变动信息
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	// 提交最新配置 若配置的ID为空，则创建新配置。若配置的删除时间不为空，则代表该配置已被删除。
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// 更新组信息 如果创建一个新的组，需要同时处理生成新的密钥
	UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error)
	// 通过UID获取所有组信息
	SyncGroup(context.Context, *SyncGroupRequest) (*SyncGroupResponse, error)
	// 生成用户密钥对
	GenUserKeyChain(context.Context, *GenUserKeyChainRequest) (*GenUserKeyChainResponse, error)
	// 获取mater key加密后的解密密钥,对应组
	GetEncryptedKey(context.Context, *GetEncryptedKeyRequest) (*GetEncryptedKeyResponse, error)
	// 更新mater key后修改加密后密钥,提交多组上去，将全部相关信息修改，不能有缺失。
	UpdateEncryptedKey(context.Context, *UpdateEncryptedKeyRequest) (*UpdateEncryptedKeyResponse, error)
	mustEmbedUnimplementedSyncServiceServer()
}

// UnimplementedSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSyncServiceServer struct {
}

func (UnimplementedSyncServiceServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedSyncServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSyncServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedSyncServiceServer) SyncGroup(context.Context, *SyncGroupRequest) (*SyncGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncGroup not implemented")
}
func (UnimplementedSyncServiceServer) GenUserKeyChain(context.Context, *GenUserKeyChainRequest) (*GenUserKeyChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenUserKeyChain not implemented")
}
func (UnimplementedSyncServiceServer) GetEncryptedKey(context.Context, *GetEncryptedKeyRequest) (*GetEncryptedKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncryptedKey not implemented")
}
func (UnimplementedSyncServiceServer) UpdateEncryptedKey(context.Context, *UpdateEncryptedKeyRequest) (*UpdateEncryptedKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEncryptedKey not implemented")
}
func (UnimplementedSyncServiceServer) mustEmbedUnimplementedSyncServiceServer() {}

// UnsafeSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncServiceServer will
// result in compilation errors.
type UnsafeSyncServiceServer interface {
	mustEmbedUnimplementedSyncServiceServer()
}

func RegisterSyncServiceServer(s grpc.ServiceRegistrar, srv SyncServiceServer) {
	s.RegisterService(&SyncService_ServiceDesc, srv)
}

func _SyncService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SyncGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SyncGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SyncGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SyncGroup(ctx, req.(*SyncGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GenUserKeyChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenUserKeyChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GenUserKeyChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GenUserKeyChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GenUserKeyChain(ctx, req.(*GenUserKeyChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetEncryptedKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEncryptedKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetEncryptedKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetEncryptedKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetEncryptedKey(ctx, req.(*GetEncryptedKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_UpdateEncryptedKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEncryptedKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).UpdateEncryptedKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_UpdateEncryptedKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).UpdateEncryptedKey(ctx, req.(*UpdateEncryptedKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncService_ServiceDesc is the grpc.ServiceDesc for SyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sync.v1.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _SyncService_Sync_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SyncService_Update_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _SyncService_UpdateGroup_Handler,
		},
		{
			MethodName: "SyncGroup",
			Handler:    _SyncService_SyncGroup_Handler,
		},
		{
			MethodName: "GenUserKeyChain",
			Handler:    _SyncService_GenUserKeyChain_Handler,
		},
		{
			MethodName: "GetEncryptedKey",
			Handler:    _SyncService_GetEncryptedKey_Handler,
		},
		{
			MethodName: "UpdateEncryptedKey",
			Handler:    _SyncService_UpdateEncryptedKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync/v1/sync_service.proto",
}
