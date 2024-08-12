// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/provider_service.proto

package provider_service

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
	Provider_DiscoverAssets_FullMethodName  = "/provider.Provider/DiscoverAssets"
	Provider_RunAssetScan_FullMethodName    = "/provider.Provider/RunAssetScan"
	Provider_RemoveAssetScan_FullMethodName = "/provider.Provider/RemoveAssetScan"
)

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	DiscoverAssets(ctx context.Context, in *DiscoverAssetsParams, opts ...grpc.CallOption) (*DiscoverAssetsResult, error)
	// The VMClarity CLI image is provided and can be used as a scanning tool, or otherwise, you can use your own scanning tool.
	// In case you don't use the VMClarity CLI, you should update AssetScan state in the backend: Ready, Aborted, Done, InProgress etc.
	RunAssetScan(ctx context.Context, in *RunAssetScanParams, opts ...grpc.CallOption) (*RunAssetScanResult, error)
	RemoveAssetScan(ctx context.Context, in *RemoveAssetScanParams, opts ...grpc.CallOption) (*RemoveAssetScanResult, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) DiscoverAssets(ctx context.Context, in *DiscoverAssetsParams, opts ...grpc.CallOption) (*DiscoverAssetsResult, error) {
	out := new(DiscoverAssetsResult)
	err := c.cc.Invoke(ctx, Provider_DiscoverAssets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) RunAssetScan(ctx context.Context, in *RunAssetScanParams, opts ...grpc.CallOption) (*RunAssetScanResult, error) {
	out := new(RunAssetScanResult)
	err := c.cc.Invoke(ctx, Provider_RunAssetScan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) RemoveAssetScan(ctx context.Context, in *RemoveAssetScanParams, opts ...grpc.CallOption) (*RemoveAssetScanResult, error) {
	out := new(RemoveAssetScanResult)
	err := c.cc.Invoke(ctx, Provider_RemoveAssetScan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	DiscoverAssets(context.Context, *DiscoverAssetsParams) (*DiscoverAssetsResult, error)
	// The VMClarity CLI image is provided and can be used as a scanning tool, or otherwise, you can use your own scanning tool.
	// In case you don't use the VMClarity CLI, you should update AssetScan state in the backend: Ready, Aborted, Done, InProgress etc.
	RunAssetScan(context.Context, *RunAssetScanParams) (*RunAssetScanResult, error)
	RemoveAssetScan(context.Context, *RemoveAssetScanParams) (*RemoveAssetScanResult, error)
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) DiscoverAssets(context.Context, *DiscoverAssetsParams) (*DiscoverAssetsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverAssets not implemented")
}
func (UnimplementedProviderServer) RunAssetScan(context.Context, *RunAssetScanParams) (*RunAssetScanResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunAssetScan not implemented")
}
func (UnimplementedProviderServer) RemoveAssetScan(context.Context, *RemoveAssetScanParams) (*RemoveAssetScanResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAssetScan not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_DiscoverAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverAssetsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).DiscoverAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_DiscoverAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).DiscoverAssets(ctx, req.(*DiscoverAssetsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_RunAssetScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunAssetScanParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).RunAssetScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_RunAssetScan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).RunAssetScan(ctx, req.(*RunAssetScanParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_RemoveAssetScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAssetScanParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).RemoveAssetScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_RemoveAssetScan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).RemoveAssetScan(ctx, req.(*RemoveAssetScanParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "provider.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiscoverAssets",
			Handler:    _Provider_DiscoverAssets_Handler,
		},
		{
			MethodName: "RunAssetScan",
			Handler:    _Provider_RunAssetScan_Handler,
		},
		{
			MethodName: "RemoveAssetScan",
			Handler:    _Provider_RemoveAssetScan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/provider_service.proto",
}
