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
	SyncService_SyncConfig_FullMethodName          = "/sync.v1.SyncService/SyncConfig"
	SyncService_UpdateConfig_FullMethodName        = "/sync.v1.SyncService/UpdateConfig"
	SyncService_SyncGroup_FullMethodName           = "/sync.v1.SyncService/SyncGroup"
	SyncService_SyncUserKeyWallet_FullMethodName   = "/sync.v1.SyncService/SyncUserKeyWallet"
	SyncService_UpdateUserKeyWallet_FullMethodName = "/sync.v1.SyncService/UpdateUserKeyWallet"
)

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncServiceClient interface {
	// 拉取指定时间点之后的配置变动信息
	SyncConfig(ctx context.Context, in *SyncConfigRequest, opts ...grpc.CallOption) (*SyncConfigResponse, error)
	// 提交最新配置 若配置的ID为空，则创建新配置。若配置的删除时间不为空，则代表该配置已被删除。这里只负责配置内容修改。
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error)
	// 通过UID获取所有所在组变动信息
	SyncGroup(ctx context.Context, in *SyncGroupRequest, opts ...grpc.CallOption) (*SyncGroupResponse, error)
	// 获取用户密钥对
	SyncUserKeyWallet(ctx context.Context, in *SyncUserKeyWalletRequest, opts ...grpc.CallOption) (*SyncUserKeyWalletResponse, error)
	// 修改用户密钥对，修改的时候所有相关组的加密密钥均要替换
	UpdateUserKeyWallet(ctx context.Context, in *UpdateUserKeyWalletRequest, opts ...grpc.CallOption) (*UpdateUserKeyWalletResponse, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) SyncConfig(ctx context.Context, in *SyncConfigRequest, opts ...grpc.CallOption) (*SyncConfigResponse, error) {
	out := new(SyncConfigResponse)
	err := c.cc.Invoke(ctx, SyncService_SyncConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error) {
	out := new(UpdateConfigResponse)
	err := c.cc.Invoke(ctx, SyncService_UpdateConfig_FullMethodName, in, out, opts...)
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

func (c *syncServiceClient) SyncUserKeyWallet(ctx context.Context, in *SyncUserKeyWalletRequest, opts ...grpc.CallOption) (*SyncUserKeyWalletResponse, error) {
	out := new(SyncUserKeyWalletResponse)
	err := c.cc.Invoke(ctx, SyncService_SyncUserKeyWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) UpdateUserKeyWallet(ctx context.Context, in *UpdateUserKeyWalletRequest, opts ...grpc.CallOption) (*UpdateUserKeyWalletResponse, error) {
	out := new(UpdateUserKeyWalletResponse)
	err := c.cc.Invoke(ctx, SyncService_UpdateUserKeyWallet_FullMethodName, in, out, opts...)
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
	SyncConfig(context.Context, *SyncConfigRequest) (*SyncConfigResponse, error)
	// 提交最新配置 若配置的ID为空，则创建新配置。若配置的删除时间不为空，则代表该配置已被删除。这里只负责配置内容修改。
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error)
	// 通过UID获取所有所在组变动信息
	SyncGroup(context.Context, *SyncGroupRequest) (*SyncGroupResponse, error)
	// 获取用户密钥对
	SyncUserKeyWallet(context.Context, *SyncUserKeyWalletRequest) (*SyncUserKeyWalletResponse, error)
	// 修改用户密钥对，修改的时候所有相关组的加密密钥均要替换
	UpdateUserKeyWallet(context.Context, *UpdateUserKeyWalletRequest) (*UpdateUserKeyWalletResponse, error)
	mustEmbedUnimplementedSyncServiceServer()
}

// UnimplementedSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSyncServiceServer struct {
}

func (UnimplementedSyncServiceServer) SyncConfig(context.Context, *SyncConfigRequest) (*SyncConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncConfig not implemented")
}
func (UnimplementedSyncServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedSyncServiceServer) SyncGroup(context.Context, *SyncGroupRequest) (*SyncGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncGroup not implemented")
}
func (UnimplementedSyncServiceServer) SyncUserKeyWallet(context.Context, *SyncUserKeyWalletRequest) (*SyncUserKeyWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncUserKeyWallet not implemented")
}
func (UnimplementedSyncServiceServer) UpdateUserKeyWallet(context.Context, *UpdateUserKeyWalletRequest) (*UpdateUserKeyWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserKeyWallet not implemented")
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

func _SyncService_SyncConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SyncConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SyncConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SyncConfig(ctx, req.(*SyncConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_UpdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
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

func _SyncService_SyncUserKeyWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUserKeyWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SyncUserKeyWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SyncUserKeyWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SyncUserKeyWallet(ctx, req.(*SyncUserKeyWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_UpdateUserKeyWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserKeyWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).UpdateUserKeyWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_UpdateUserKeyWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).UpdateUserKeyWallet(ctx, req.(*UpdateUserKeyWalletRequest))
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
			MethodName: "SyncConfig",
			Handler:    _SyncService_SyncConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _SyncService_UpdateConfig_Handler,
		},
		{
			MethodName: "SyncGroup",
			Handler:    _SyncService_SyncGroup_Handler,
		},
		{
			MethodName: "SyncUserKeyWallet",
			Handler:    _SyncService_SyncUserKeyWallet_Handler,
		},
		{
			MethodName: "UpdateUserKeyWallet",
			Handler:    _SyncService_UpdateUserKeyWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync/v1/sync_service.proto",
}
