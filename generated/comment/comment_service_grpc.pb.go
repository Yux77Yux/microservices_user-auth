// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: comment/comment_service.proto

package comment

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
	CommentService_PublishComment_FullMethodName    = "/comment.CommentService/PublishComment"
	CommentService_InitialComments_FullMethodName   = "/comment.CommentService/InitialComments"
	CommentService_DeleteComment_FullMethodName     = "/comment.CommentService/DeleteComment"
	CommentService_GetTopComments_FullMethodName    = "/comment.CommentService/GetTopComments"
	CommentService_GetSecondComments_FullMethodName = "/comment.CommentService/GetSecondComments"
	CommentService_GetReplyComments_FullMethodName  = "/comment.CommentService/GetReplyComments"
	CommentService_GetComment_FullMethodName        = "/comment.CommentService/GetComment"
	CommentService_GetComments_FullMethodName       = "/comment.CommentService/GetComments"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	PublishComment(ctx context.Context, in *PublishCommentRequest, opts ...grpc.CallOption) (*PublishCommentResponse, error)
	InitialComments(ctx context.Context, in *InitialCommentsRequest, opts ...grpc.CallOption) (*InitialCommentsResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTopComments(ctx context.Context, in *GetTopCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	GetSecondComments(ctx context.Context, in *GetSecondCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	GetReplyComments(ctx context.Context, in *GetReplyCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) PublishComment(ctx context.Context, in *PublishCommentRequest, opts ...grpc.CallOption) (*PublishCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_PublishComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) InitialComments(ctx context.Context, in *InitialCommentsRequest, opts ...grpc.CallOption) (*InitialCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitialCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_InitialComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CommentService_DeleteComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetTopComments(ctx context.Context, in *GetTopCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetTopComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetSecondComments(ctx context.Context, in *GetSecondCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetSecondComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetReplyComments(ctx context.Context, in *GetReplyCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetReplyComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_GetComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility.
type CommentServiceServer interface {
	PublishComment(context.Context, *PublishCommentRequest) (*PublishCommentResponse, error)
	InitialComments(context.Context, *InitialCommentsRequest) (*InitialCommentsResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error)
	GetTopComments(context.Context, *GetTopCommentsRequest) (*GetCommentsResponse, error)
	GetSecondComments(context.Context, *GetSecondCommentsRequest) (*GetCommentsResponse, error)
	GetReplyComments(context.Context, *GetReplyCommentsRequest) (*GetCommentsResponse, error)
	GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentServiceServer struct{}

func (UnimplementedCommentServiceServer) PublishComment(context.Context, *PublishCommentRequest) (*PublishCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishComment not implemented")
}
func (UnimplementedCommentServiceServer) InitialComments(context.Context, *InitialCommentsRequest) (*InitialCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialComments not implemented")
}
func (UnimplementedCommentServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentServiceServer) GetTopComments(context.Context, *GetTopCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopComments not implemented")
}
func (UnimplementedCommentServiceServer) GetSecondComments(context.Context, *GetSecondCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecondComments not implemented")
}
func (UnimplementedCommentServiceServer) GetReplyComments(context.Context, *GetReplyCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplyComments not implemented")
}
func (UnimplementedCommentServiceServer) GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCommentServiceServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}
func (UnimplementedCommentServiceServer) testEmbeddedByValue()                        {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_PublishComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).PublishComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_PublishComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).PublishComment(ctx, req.(*PublishCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_InitialComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitialCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).InitialComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_InitialComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).InitialComments(ctx, req.(*InitialCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetTopComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetTopComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetTopComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetTopComments(ctx, req.(*GetTopCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetSecondComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecondCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetSecondComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetSecondComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetSecondComments(ctx, req.(*GetSecondCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetReplyComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReplyCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetReplyComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetReplyComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetReplyComments(ctx, req.(*GetReplyCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishComment",
			Handler:    _CommentService_PublishComment_Handler,
		},
		{
			MethodName: "InitialComments",
			Handler:    _CommentService_InitialComments_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentService_DeleteComment_Handler,
		},
		{
			MethodName: "GetTopComments",
			Handler:    _CommentService_GetTopComments_Handler,
		},
		{
			MethodName: "GetSecondComments",
			Handler:    _CommentService_GetSecondComments_Handler,
		},
		{
			MethodName: "GetReplyComments",
			Handler:    _CommentService_GetReplyComments_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _CommentService_GetComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _CommentService_GetComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment/comment_service.proto",
}
