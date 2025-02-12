// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: creation/creation_service.proto

package creation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CreationService_UploadCreation_FullMethodName            = "/creation.CreationService/UploadCreation"
	CreationService_GetCreation_FullMethodName               = "/creation.CreationService/GetCreation"
	CreationService_GetSpaceCreationList_FullMethodName      = "/creation.CreationService/GetSpaceCreationList"
	CreationService_GetCollectionCreationList_FullMethodName = "/creation.CreationService/GetCollectionCreationList"
	CreationService_GetHomeCreationList_FullMethodName       = "/creation.CreationService/GetHomeCreationList"
	CreationService_GetSimilarCreationList_FullMethodName    = "/creation.CreationService/GetSimilarCreationList"
	CreationService_DeleteCreation_FullMethodName            = "/creation.CreationService/DeleteCreation"
	CreationService_UpdateCreation_FullMethodName            = "/creation.CreationService/UpdateCreation"
	CreationService_UpdateCreationStatus_FullMethodName      = "/creation.CreationService/UpdateCreationStatus"
)

// CreationServiceClient is the client API for CreationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreationServiceClient interface {
	// POST
	UploadCreation(ctx context.Context, in *UploadCreationRequest, opts ...grpc.CallOption) (*UploadCreationResponse, error)
	// GET
	GetCreation(ctx context.Context, in *GetCreationRequest, opts ...grpc.CallOption) (*GetCreationResponse, error)
	// 作者的作品
	GetSpaceCreationList(ctx context.Context, in *GetPublicCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	// 收藏夹
	GetCollectionCreationList(ctx context.Context, in *GetSpecificCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	// 推荐列表在creation处获取，但写入由interaction负责，不够则发布事件进行重新计算
	// 主页
	GetHomeCreationList(ctx context.Context, in *GetSpecificCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	// 相似列表在creation处获取，但写入由interaction负责，不存在则直接调取api然后直接返回，同时存入redis
	// 作品的相似作品列表
	GetSimilarCreationList(ctx context.Context, in *GetPublicCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	// DELETE
	DeleteCreation(ctx context.Context, in *DeleteCreationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UPDATE
	UpdateCreation(ctx context.Context, in *UpdateCreationRequest, opts ...grpc.CallOption) (*UpdateCreationResponse, error)
	UpdateCreationStatus(ctx context.Context, in *UpdateCreationStatusRequest, opts ...grpc.CallOption) (*UpdateCreationResponse, error)
}

type creationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreationServiceClient(cc grpc.ClientConnInterface) CreationServiceClient {
	return &creationServiceClient{cc}
}

func (c *creationServiceClient) UploadCreation(ctx context.Context, in *UploadCreationRequest, opts ...grpc.CallOption) (*UploadCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_UploadCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetCreation(ctx context.Context, in *GetCreationRequest, opts ...grpc.CallOption) (*GetCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_GetCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetSpaceCreationList(ctx context.Context, in *GetPublicCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_GetSpaceCreationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetCollectionCreationList(ctx context.Context, in *GetSpecificCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_GetCollectionCreationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetHomeCreationList(ctx context.Context, in *GetSpecificCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_GetHomeCreationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetSimilarCreationList(ctx context.Context, in *GetPublicCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_GetSimilarCreationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) DeleteCreation(ctx context.Context, in *DeleteCreationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CreationService_DeleteCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) UpdateCreation(ctx context.Context, in *UpdateCreationRequest, opts ...grpc.CallOption) (*UpdateCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_UpdateCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) UpdateCreationStatus(ctx context.Context, in *UpdateCreationStatusRequest, opts ...grpc.CallOption) (*UpdateCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_UpdateCreationStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreationServiceServer is the server API for CreationService service.
// All implementations must embed UnimplementedCreationServiceServer
// for forward compatibility.
type CreationServiceServer interface {
	// POST
	UploadCreation(context.Context, *UploadCreationRequest) (*UploadCreationResponse, error)
	// GET
	GetCreation(context.Context, *GetCreationRequest) (*GetCreationResponse, error)
	// 作者的作品
	GetSpaceCreationList(context.Context, *GetPublicCreationListRequest) (*GetCreationListResponse, error)
	// 收藏夹
	GetCollectionCreationList(context.Context, *GetSpecificCreationListRequest) (*GetCreationListResponse, error)
	// 推荐列表在creation处获取，但写入由interaction负责，不够则发布事件进行重新计算
	// 主页
	GetHomeCreationList(context.Context, *GetSpecificCreationListRequest) (*GetCreationListResponse, error)
	// 相似列表在creation处获取，但写入由interaction负责，不存在则直接调取api然后直接返回，同时存入redis
	// 作品的相似作品列表
	GetSimilarCreationList(context.Context, *GetPublicCreationListRequest) (*GetCreationListResponse, error)
	// DELETE
	DeleteCreation(context.Context, *DeleteCreationRequest) (*emptypb.Empty, error)
	// UPDATE
	UpdateCreation(context.Context, *UpdateCreationRequest) (*UpdateCreationResponse, error)
	UpdateCreationStatus(context.Context, *UpdateCreationStatusRequest) (*UpdateCreationResponse, error)
	mustEmbedUnimplementedCreationServiceServer()
}

// UnimplementedCreationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCreationServiceServer struct{}

func (UnimplementedCreationServiceServer) UploadCreation(context.Context, *UploadCreationRequest) (*UploadCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCreation not implemented")
}
func (UnimplementedCreationServiceServer) GetCreation(context.Context, *GetCreationRequest) (*GetCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreation not implemented")
}
func (UnimplementedCreationServiceServer) GetSpaceCreationList(context.Context, *GetPublicCreationListRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpaceCreationList not implemented")
}
func (UnimplementedCreationServiceServer) GetCollectionCreationList(context.Context, *GetSpecificCreationListRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionCreationList not implemented")
}
func (UnimplementedCreationServiceServer) GetHomeCreationList(context.Context, *GetSpecificCreationListRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeCreationList not implemented")
}
func (UnimplementedCreationServiceServer) GetSimilarCreationList(context.Context, *GetPublicCreationListRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimilarCreationList not implemented")
}
func (UnimplementedCreationServiceServer) DeleteCreation(context.Context, *DeleteCreationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCreation not implemented")
}
func (UnimplementedCreationServiceServer) UpdateCreation(context.Context, *UpdateCreationRequest) (*UpdateCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCreation not implemented")
}
func (UnimplementedCreationServiceServer) UpdateCreationStatus(context.Context, *UpdateCreationStatusRequest) (*UpdateCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCreationStatus not implemented")
}
func (UnimplementedCreationServiceServer) mustEmbedUnimplementedCreationServiceServer() {}
func (UnimplementedCreationServiceServer) testEmbeddedByValue()                         {}

// UnsafeCreationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreationServiceServer will
// result in compilation errors.
type UnsafeCreationServiceServer interface {
	mustEmbedUnimplementedCreationServiceServer()
}

func RegisterCreationServiceServer(s grpc.ServiceRegistrar, srv CreationServiceServer) {
	// If the following call pancis, it indicates UnimplementedCreationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CreationService_ServiceDesc, srv)
}

func _CreationService_UploadCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).UploadCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_UploadCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).UploadCreation(ctx, req.(*UploadCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetCreation(ctx, req.(*GetCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetSpaceCreationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicCreationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetSpaceCreationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetSpaceCreationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetSpaceCreationList(ctx, req.(*GetPublicCreationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetCollectionCreationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecificCreationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetCollectionCreationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetCollectionCreationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetCollectionCreationList(ctx, req.(*GetSpecificCreationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetHomeCreationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecificCreationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetHomeCreationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetHomeCreationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetHomeCreationList(ctx, req.(*GetSpecificCreationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetSimilarCreationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicCreationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetSimilarCreationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetSimilarCreationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetSimilarCreationList(ctx, req.(*GetPublicCreationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_DeleteCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).DeleteCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_DeleteCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).DeleteCreation(ctx, req.(*DeleteCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_UpdateCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).UpdateCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_UpdateCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).UpdateCreation(ctx, req.(*UpdateCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_UpdateCreationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCreationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).UpdateCreationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_UpdateCreationStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).UpdateCreationStatus(ctx, req.(*UpdateCreationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreationService_ServiceDesc is the grpc.ServiceDesc for CreationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "creation.CreationService",
	HandlerType: (*CreationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadCreation",
			Handler:    _CreationService_UploadCreation_Handler,
		},
		{
			MethodName: "GetCreation",
			Handler:    _CreationService_GetCreation_Handler,
		},
		{
			MethodName: "GetSpaceCreationList",
			Handler:    _CreationService_GetSpaceCreationList_Handler,
		},
		{
			MethodName: "GetCollectionCreationList",
			Handler:    _CreationService_GetCollectionCreationList_Handler,
		},
		{
			MethodName: "GetHomeCreationList",
			Handler:    _CreationService_GetHomeCreationList_Handler,
		},
		{
			MethodName: "GetSimilarCreationList",
			Handler:    _CreationService_GetSimilarCreationList_Handler,
		},
		{
			MethodName: "DeleteCreation",
			Handler:    _CreationService_DeleteCreation_Handler,
		},
		{
			MethodName: "UpdateCreation",
			Handler:    _CreationService_UpdateCreation_Handler,
		},
		{
			MethodName: "UpdateCreationStatus",
			Handler:    _CreationService_UpdateCreationStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "creation/creation_service.proto",
}
