// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: review/review_service.proto

package review

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReviewService_UpdateReview_FullMethodName = "/review.ReviewService/UpdateReview"
	ReviewService_NewReview_FullMethodName    = "/review.ReviewService/NewReview"
	ReviewService_GetReview_FullMethodName    = "/review.ReviewService/GetReview"
)

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewServiceClient interface {
	// UPDATE
	UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error)
	// POST
	NewReview(ctx context.Context, in *NewReviewRequest, opts ...grpc.CallOption) (*NewReviewResponse, error)
	// GET
	GetReview(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error)
}

type reviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewServiceClient(cc grpc.ClientConnInterface) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReviewResponse)
	err := c.cc.Invoke(ctx, ReviewService_UpdateReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) NewReview(ctx context.Context, in *NewReviewRequest, opts ...grpc.CallOption) (*NewReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewReviewResponse)
	err := c.cc.Invoke(ctx, ReviewService_NewReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetReview(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReviewsResponse)
	err := c.cc.Invoke(ctx, ReviewService_GetReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
// All implementations must embed UnimplementedReviewServiceServer
// for forward compatibility.
type ReviewServiceServer interface {
	// UPDATE
	UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error)
	// POST
	NewReview(context.Context, *NewReviewRequest) (*NewReviewResponse, error)
	// GET
	GetReview(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error)
	mustEmbedUnimplementedReviewServiceServer()
}

// UnimplementedReviewServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReviewServiceServer struct{}

func (UnimplementedReviewServiceServer) UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedReviewServiceServer) NewReview(context.Context, *NewReviewRequest) (*NewReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewReview not implemented")
}
func (UnimplementedReviewServiceServer) GetReview(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReview not implemented")
}
func (UnimplementedReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {}
func (UnimplementedReviewServiceServer) testEmbeddedByValue()                       {}

// UnsafeReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServiceServer will
// result in compilation errors.
type UnsafeReviewServiceServer interface {
	mustEmbedUnimplementedReviewServiceServer()
}

func RegisterReviewServiceServer(s grpc.ServiceRegistrar, srv ReviewServiceServer) {
	// If the following call pancis, it indicates UnimplementedReviewServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReviewService_ServiceDesc, srv)
}

func _ReviewService_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewService_UpdateReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).UpdateReview(ctx, req.(*UpdateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_NewReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).NewReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewService_NewReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).NewReview(ctx, req.(*NewReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewService_GetReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReview(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewService_ServiceDesc is the grpc.ServiceDesc for ReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "review.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateReview",
			Handler:    _ReviewService_UpdateReview_Handler,
		},
		{
			MethodName: "NewReview",
			Handler:    _ReviewService_NewReview_Handler,
		},
		{
			MethodName: "GetReview",
			Handler:    _ReviewService_GetReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review/review_service.proto",
}
