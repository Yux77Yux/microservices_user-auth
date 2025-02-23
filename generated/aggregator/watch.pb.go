// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: aggregator/methods/watch.proto

package aggregator

import (
	comment "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
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

// 视频和用户信息   专注于返回视频信息，用户数据处理延后
type WatchCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId int64 `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
}

func (x *WatchCreationRequest) Reset() {
	*x = WatchCreationRequest{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchCreationRequest) ProtoMessage() {}

func (x *WatchCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchCreationRequest.ProtoReflect.Descriptor instead.
func (*WatchCreationRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{0}
}

func (x *WatchCreationRequest) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

type WatchCreationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg          *common.ApiResponse         `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`                                       // 返回的响应
	CreationUser *common.UserCreationComment `protobuf:"bytes,2,opt,name=creation_user,json=creationUser,proto3" json:"creation_user,omitempty"` // 作品的用户信息
	CreationInfo *creation.CreationInfo      `protobuf:"bytes,3,opt,name=creation_info,json=creationInfo,proto3" json:"creation_info,omitempty"` // 作品的详细信息
}

func (x *WatchCreationResponse) Reset() {
	*x = WatchCreationResponse{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchCreationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchCreationResponse) ProtoMessage() {}

func (x *WatchCreationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchCreationResponse.ProtoReflect.Descriptor instead.
func (*WatchCreationResponse) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{1}
}

func (x *WatchCreationResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *WatchCreationResponse) GetCreationUser() *common.UserCreationComment {
	if x != nil {
		return x.CreationUser
	}
	return nil
}

func (x *WatchCreationResponse) GetCreationInfo() *creation.CreationInfo {
	if x != nil {
		return x.CreationInfo
	}
	return nil
}

// 第一次加载作品的评论
type InitialCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *comment.InitialCommentsRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *InitialCommentsRequest) Reset() {
	*x = InitialCommentsRequest{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitialCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialCommentsRequest) ProtoMessage() {}

func (x *InitialCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialCommentsRequest.ProtoReflect.Descriptor instead.
func (*InitialCommentsRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{2}
}

func (x *InitialCommentsRequest) GetRequest() *comment.InitialCommentsRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type InitialCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments  []*TopCommentInfo    `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Area      *comment.CommentArea `protobuf:"bytes,2,opt,name=area,proto3" json:"area,omitempty"`
	PageCount int32                `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	Msg       *common.ApiResponse  `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"` // 返回的响应
}

func (x *InitialCommentsResponse) Reset() {
	*x = InitialCommentsResponse{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitialCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialCommentsResponse) ProtoMessage() {}

func (x *InitialCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialCommentsResponse.ProtoReflect.Descriptor instead.
func (*InitialCommentsResponse) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{3}
}

func (x *InitialCommentsResponse) GetComments() []*TopCommentInfo {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *InitialCommentsResponse) GetArea() *comment.CommentArea {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *InitialCommentsResponse) GetPageCount() int32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *InitialCommentsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

type GetTopCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *comment.GetTopCommentsRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetTopCommentsRequest) Reset() {
	*x = GetTopCommentsRequest{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopCommentsRequest) ProtoMessage() {}

func (x *GetTopCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetTopCommentsRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{4}
}

func (x *GetTopCommentsRequest) GetRequest() *comment.GetTopCommentsRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetTopCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*TopCommentInfo   `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Msg      *common.ApiResponse `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // 返回的响应
}

func (x *GetTopCommentsResponse) Reset() {
	*x = GetTopCommentsResponse{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopCommentsResponse) ProtoMessage() {}

func (x *GetTopCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetTopCommentsResponse) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{5}
}

func (x *GetTopCommentsResponse) GetComments() []*TopCommentInfo {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetTopCommentsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 一级评论内的评论
type GetSecondCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request     *comment.GetSecondCommentsRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	AccessToken *common.AccessToken               `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetSecondCommentsRequest) Reset() {
	*x = GetSecondCommentsRequest{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecondCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecondCommentsRequest) ProtoMessage() {}

func (x *GetSecondCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecondCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetSecondCommentsRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{6}
}

func (x *GetSecondCommentsRequest) GetRequest() *comment.GetSecondCommentsRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *GetSecondCommentsRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type GetSecondCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*SecondCommentInfo `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Msg      *common.ApiResponse  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // 返回的响应
}

func (x *GetSecondCommentsResponse) Reset() {
	*x = GetSecondCommentsResponse{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecondCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecondCommentsResponse) ProtoMessage() {}

func (x *GetSecondCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecondCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetSecondCommentsResponse) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{7}
}

func (x *GetSecondCommentsResponse) GetComments() []*SecondCommentInfo {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetSecondCommentsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 回复我的评论，在页面的消息内显示, 权限类 没做
type GetReplyCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request     *comment.GetReplyCommentsRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	AccessToken *common.AccessToken              `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetReplyCommentsRequest) Reset() {
	*x = GetReplyCommentsRequest{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReplyCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReplyCommentsRequest) ProtoMessage() {}

func (x *GetReplyCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReplyCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetReplyCommentsRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{8}
}

func (x *GetReplyCommentsRequest) GetRequest() *comment.GetReplyCommentsRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *GetReplyCommentsRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type GetCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*CommentInfo      `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Msg      *common.ApiResponse `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // 返回的响应
}

func (x *GetCommentsResponse) Reset() {
	*x = GetCommentsResponse{}
	mi := &file_aggregator_methods_watch_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsResponse) ProtoMessage() {}

func (x *GetCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_methods_watch_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetCommentsResponse) Descriptor() ([]byte, []int) {
	return file_aggregator_methods_watch_proto_rawDescGZIP(), []int{9}
}

func (x *GetCommentsResponse) GetComments() []*CommentInfo {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetCommentsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_aggregator_methods_watch_proto protoreflect.FileDescriptor

var file_aggregator_methods_watch_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x1a, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x14, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xc6, 0x01, 0x0a,
	0x15, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x40, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x44, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5b, 0x0a, 0x16, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x17, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x31, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x59, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x97, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x86, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x95, 0x01, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x7a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x3b, 0x5a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37,
	0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_aggregator_methods_watch_proto_rawDescOnce sync.Once
	file_aggregator_methods_watch_proto_rawDescData = file_aggregator_methods_watch_proto_rawDesc
)

func file_aggregator_methods_watch_proto_rawDescGZIP() []byte {
	file_aggregator_methods_watch_proto_rawDescOnce.Do(func() {
		file_aggregator_methods_watch_proto_rawDescData = protoimpl.X.CompressGZIP(file_aggregator_methods_watch_proto_rawDescData)
	})
	return file_aggregator_methods_watch_proto_rawDescData
}

var file_aggregator_methods_watch_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_aggregator_methods_watch_proto_goTypes = []any{
	(*WatchCreationRequest)(nil),             // 0: aggregator.methods.WatchCreationRequest
	(*WatchCreationResponse)(nil),            // 1: aggregator.methods.WatchCreationResponse
	(*InitialCommentsRequest)(nil),           // 2: aggregator.methods.InitialCommentsRequest
	(*InitialCommentsResponse)(nil),          // 3: aggregator.methods.InitialCommentsResponse
	(*GetTopCommentsRequest)(nil),            // 4: aggregator.methods.GetTopCommentsRequest
	(*GetTopCommentsResponse)(nil),           // 5: aggregator.methods.GetTopCommentsResponse
	(*GetSecondCommentsRequest)(nil),         // 6: aggregator.methods.GetSecondCommentsRequest
	(*GetSecondCommentsResponse)(nil),        // 7: aggregator.methods.GetSecondCommentsResponse
	(*GetReplyCommentsRequest)(nil),          // 8: aggregator.methods.GetReplyCommentsRequest
	(*GetCommentsResponse)(nil),              // 9: aggregator.methods.GetCommentsResponse
	(*common.ApiResponse)(nil),               // 10: common.ApiResponse
	(*common.UserCreationComment)(nil),       // 11: common.UserCreationComment
	(*creation.CreationInfo)(nil),            // 12: creation.messages.CreationInfo
	(*comment.InitialCommentsRequest)(nil),   // 13: comment.methods.InitialCommentsRequest
	(*TopCommentInfo)(nil),                   // 14: aggregator.messages.TopCommentInfo
	(*comment.CommentArea)(nil),              // 15: comment.messages.CommentArea
	(*comment.GetTopCommentsRequest)(nil),    // 16: comment.methods.GetTopCommentsRequest
	(*comment.GetSecondCommentsRequest)(nil), // 17: comment.methods.GetSecondCommentsRequest
	(*common.AccessToken)(nil),               // 18: common.AccessToken
	(*SecondCommentInfo)(nil),                // 19: aggregator.messages.SecondCommentInfo
	(*comment.GetReplyCommentsRequest)(nil),  // 20: comment.methods.GetReplyCommentsRequest
	(*CommentInfo)(nil),                      // 21: aggregator.messages.CommentInfo
}
var file_aggregator_methods_watch_proto_depIdxs = []int32{
	10, // 0: aggregator.methods.WatchCreationResponse.msg:type_name -> common.ApiResponse
	11, // 1: aggregator.methods.WatchCreationResponse.creation_user:type_name -> common.UserCreationComment
	12, // 2: aggregator.methods.WatchCreationResponse.creation_info:type_name -> creation.messages.CreationInfo
	13, // 3: aggregator.methods.InitialCommentsRequest.request:type_name -> comment.methods.InitialCommentsRequest
	14, // 4: aggregator.methods.InitialCommentsResponse.comments:type_name -> aggregator.messages.TopCommentInfo
	15, // 5: aggregator.methods.InitialCommentsResponse.area:type_name -> comment.messages.CommentArea
	10, // 6: aggregator.methods.InitialCommentsResponse.msg:type_name -> common.ApiResponse
	16, // 7: aggregator.methods.GetTopCommentsRequest.request:type_name -> comment.methods.GetTopCommentsRequest
	14, // 8: aggregator.methods.GetTopCommentsResponse.comments:type_name -> aggregator.messages.TopCommentInfo
	10, // 9: aggregator.methods.GetTopCommentsResponse.msg:type_name -> common.ApiResponse
	17, // 10: aggregator.methods.GetSecondCommentsRequest.request:type_name -> comment.methods.GetSecondCommentsRequest
	18, // 11: aggregator.methods.GetSecondCommentsRequest.access_token:type_name -> common.AccessToken
	19, // 12: aggregator.methods.GetSecondCommentsResponse.comments:type_name -> aggregator.messages.SecondCommentInfo
	10, // 13: aggregator.methods.GetSecondCommentsResponse.msg:type_name -> common.ApiResponse
	20, // 14: aggregator.methods.GetReplyCommentsRequest.request:type_name -> comment.methods.GetReplyCommentsRequest
	18, // 15: aggregator.methods.GetReplyCommentsRequest.access_token:type_name -> common.AccessToken
	21, // 16: aggregator.methods.GetCommentsResponse.comments:type_name -> aggregator.messages.CommentInfo
	10, // 17: aggregator.methods.GetCommentsResponse.msg:type_name -> common.ApiResponse
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_aggregator_methods_watch_proto_init() }
func file_aggregator_methods_watch_proto_init() {
	if File_aggregator_methods_watch_proto != nil {
		return
	}
	file_aggregator_messages_comment_info_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aggregator_methods_watch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aggregator_methods_watch_proto_goTypes,
		DependencyIndexes: file_aggregator_methods_watch_proto_depIdxs,
		MessageInfos:      file_aggregator_methods_watch_proto_msgTypes,
	}.Build()
	File_aggregator_methods_watch_proto = out.File
	file_aggregator_methods_watch_proto_rawDesc = nil
	file_aggregator_methods_watch_proto_goTypes = nil
	file_aggregator_methods_watch_proto_depIdxs = nil
}
