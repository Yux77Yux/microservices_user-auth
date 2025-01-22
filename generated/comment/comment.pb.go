// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: comment/messages/comment.proto

package comment

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

type CommentArea_Status int32

const (
	CommentArea_DEFAULT CommentArea_Status = 0
	CommentArea_HIDING  CommentArea_Status = 1
	CommentArea_CLOSED  CommentArea_Status = 3
)

// Enum value maps for CommentArea_Status.
var (
	CommentArea_Status_name = map[int32]string{
		0: "DEFAULT",
		1: "HIDING",
		3: "CLOSED",
	}
	CommentArea_Status_value = map[string]int32{
		"DEFAULT": 0,
		"HIDING":  1,
		"CLOSED":  3,
	}
)

func (x CommentArea_Status) Enum() *CommentArea_Status {
	p := new(CommentArea_Status)
	*p = x
	return p
}

func (x CommentArea_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentArea_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_messages_comment_proto_enumTypes[0].Descriptor()
}

func (CommentArea_Status) Type() protoreflect.EnumType {
	return &file_comment_messages_comment_proto_enumTypes[0]
}

func (x CommentArea_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentArea_Status.Descriptor instead.
func (CommentArea_Status) EnumDescriptor() ([]byte, []int) {
	return file_comment_messages_comment_proto_rawDescGZIP(), []int{1, 0}
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId  int32                  `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	Root       int32                  `protobuf:"varint,2,opt,name=root,proto3" json:"root,omitempty"`
	Parent     int32                  `protobuf:"varint,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Dialog     int32                  `protobuf:"varint,4,opt,name=dialog,proto3" json:"dialog,omitempty"`
	UserId     int64                  `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreationId int64                  `protobuf:"varint,6,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Content    string                 `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	Media      string                 `protobuf:"bytes,9,opt,name=media,proto3" json:"media,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	mi := &file_comment_messages_comment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_messages_comment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comment_messages_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetCommentId() int32 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *Comment) GetRoot() int32 {
	if x != nil {
		return x.Root
	}
	return 0
}

func (x *Comment) GetParent() int32 {
	if x != nil {
		return x.Parent
	}
	return 0
}

func (x *Comment) GetDialog() int32 {
	if x != nil {
		return x.Dialog
	}
	return 0
}

func (x *Comment) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

type CommentArea struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId    int64              `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
	TotalComments int32              `protobuf:"varint,2,opt,name=total_comments,json=totalComments,proto3" json:"total_comments,omitempty"`
	AreaStatus    CommentArea_Status `protobuf:"varint,3,opt,name=area_status,json=areaStatus,proto3,enum=comment.messages.CommentArea_Status" json:"area_status,omitempty"`
}

func (x *CommentArea) Reset() {
	*x = CommentArea{}
	mi := &file_comment_messages_comment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommentArea) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentArea) ProtoMessage() {}

func (x *CommentArea) ProtoReflect() protoreflect.Message {
	mi := &file_comment_messages_comment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentArea.ProtoReflect.Descriptor instead.
func (*CommentArea) Descriptor() ([]byte, []int) {
	return file_comment_messages_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentArea) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

func (x *CommentArea) GetTotalComments() int32 {
	if x != nil {
		return x.TotalComments
	}
	return 0
}

func (x *CommentArea) GetAreaStatus() CommentArea_Status {
	if x != nil {
		return x.AreaStatus
	}
	return CommentArea_DEFAULT
}

type AnyComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyComment []*Comment `protobuf:"bytes,1,rep,name=any_comment,json=anyComment,proto3" json:"any_comment,omitempty"`
}

func (x *AnyComment) Reset() {
	*x = AnyComment{}
	mi := &file_comment_messages_comment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnyComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyComment) ProtoMessage() {}

func (x *AnyComment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_messages_comment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyComment.ProtoReflect.Descriptor instead.
func (*AnyComment) Descriptor() ([]byte, []int) {
	return file_comment_messages_comment_proto_rawDescGZIP(), []int{2}
}

func (x *AnyComment) GetAnyComment() []*Comment {
	if x != nil {
		return x.AnyComment
	}
	return nil
}

type AnyCommentArea struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyCommentArea []*CommentArea `protobuf:"bytes,1,rep,name=any_comment_area,json=anyCommentArea,proto3" json:"any_comment_area,omitempty"`
}

func (x *AnyCommentArea) Reset() {
	*x = AnyCommentArea{}
	mi := &file_comment_messages_comment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnyCommentArea) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyCommentArea) ProtoMessage() {}

func (x *AnyCommentArea) ProtoReflect() protoreflect.Message {
	mi := &file_comment_messages_comment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyCommentArea.ProtoReflect.Descriptor instead.
func (*AnyCommentArea) Descriptor() ([]byte, []int) {
	return file_comment_messages_comment_proto_rawDescGZIP(), []int{3}
}

func (x *AnyCommentArea) GetAnyCommentArea() []*CommentArea {
	if x != nil {
		return x.AnyCommentArea
	}
	return nil
}

var File_comment_messages_comment_proto protoreflect.FileDescriptor

var file_comment_messages_comment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0xcb, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x45, 0x0a, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x72, 0x65, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x61, 0x72, 0x65, 0x61,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x48, 0x49, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x44, 0x10, 0x03, 0x22, 0x48, 0x0a, 0x0a, 0x41, 0x6e, 0x79, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x6e, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x59, 0x0a, 0x0e, 0x41, 0x6e, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x65,
	0x61, 0x12, 0x47, 0x0a, 0x10, 0x61, 0x6e, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x0e, 0x61, 0x6e, 0x79, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x65, 0x61, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75,
	0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_messages_comment_proto_rawDescOnce sync.Once
	file_comment_messages_comment_proto_rawDescData = file_comment_messages_comment_proto_rawDesc
)

func file_comment_messages_comment_proto_rawDescGZIP() []byte {
	file_comment_messages_comment_proto_rawDescOnce.Do(func() {
		file_comment_messages_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_messages_comment_proto_rawDescData)
	})
	return file_comment_messages_comment_proto_rawDescData
}

var file_comment_messages_comment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_comment_messages_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_comment_messages_comment_proto_goTypes = []any{
	(CommentArea_Status)(0),       // 0: comment.messages.CommentArea.Status
	(*Comment)(nil),               // 1: comment.messages.Comment
	(*CommentArea)(nil),           // 2: comment.messages.CommentArea
	(*AnyComment)(nil),            // 3: comment.messages.AnyComment
	(*AnyCommentArea)(nil),        // 4: comment.messages.AnyCommentArea
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_comment_messages_comment_proto_depIdxs = []int32{
	5, // 0: comment.messages.Comment.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: comment.messages.CommentArea.area_status:type_name -> comment.messages.CommentArea.Status
	1, // 2: comment.messages.AnyComment.any_comment:type_name -> comment.messages.Comment
	2, // 3: comment.messages.AnyCommentArea.any_comment_area:type_name -> comment.messages.CommentArea
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_comment_messages_comment_proto_init() }
func file_comment_messages_comment_proto_init() {
	if File_comment_messages_comment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comment_messages_comment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comment_messages_comment_proto_goTypes,
		DependencyIndexes: file_comment_messages_comment_proto_depIdxs,
		EnumInfos:         file_comment_messages_comment_proto_enumTypes,
		MessageInfos:      file_comment_messages_comment_proto_msgTypes,
	}.Build()
	File_comment_messages_comment_proto = out.File
	file_comment_messages_comment_proto_rawDesc = nil
	file_comment_messages_comment_proto_goTypes = nil
	file_comment_messages_comment_proto_depIdxs = nil
}
