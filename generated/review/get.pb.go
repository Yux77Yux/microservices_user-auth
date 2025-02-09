// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: review/methods/get.proto

package review

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
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

type GetReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ReviewStatus `protobuf:"varint,1,opt,name=status,proto3,enum=review.messages.ReviewStatus" json:"status,omitempty"` // 获取需要审核的，或已被审核的
	Type   TargetType   `protobuf:"varint,2,opt,name=type,proto3,enum=review.messages.TargetType" json:"type,omitempty"`       // 获取类型，评论，用户，作品
}

func (x *GetReviewsRequest) Reset() {
	*x = GetReviewsRequest{}
	mi := &file_review_methods_get_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsRequest) ProtoMessage() {}

func (x *GetReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_methods_get_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetReviewsRequest) Descriptor() ([]byte, []int) {
	return file_review_methods_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetReviewsRequest) GetStatus() ReviewStatus {
	if x != nil {
		return x.Status
	}
	return ReviewStatus_PENDING
}

func (x *GetReviewsRequest) GetType() TargetType {
	if x != nil {
		return x.Type
	}
	return TargetType_CREATION
}

type GetReviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg     *common.ApiResponse `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Reviews []*NewReview        `protobuf:"bytes,2,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *GetReviewsResponse) Reset() {
	*x = GetReviewsResponse{}
	mi := &file_review_methods_get_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsResponse) ProtoMessage() {}

func (x *GetReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_methods_get_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsResponse.ProtoReflect.Descriptor instead.
func (*GetReviewsResponse) Descriptor() ([]byte, []int) {
	return file_review_methods_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetReviewsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *GetReviewsResponse) GetReviews() []*NewReview {
	if x != nil {
		return x.Reviews
	}
	return nil
}

var File_review_methods_get_proto protoreflect.FileDescriptor

var file_review_methods_get_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x1c, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x71,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x34, 0x0a, 0x07, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4e,
	0x65, 0x77, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_review_methods_get_proto_rawDescOnce sync.Once
	file_review_methods_get_proto_rawDescData = file_review_methods_get_proto_rawDesc
)

func file_review_methods_get_proto_rawDescGZIP() []byte {
	file_review_methods_get_proto_rawDescOnce.Do(func() {
		file_review_methods_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_methods_get_proto_rawDescData)
	})
	return file_review_methods_get_proto_rawDescData
}

var file_review_methods_get_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_review_methods_get_proto_goTypes = []any{
	(*GetReviewsRequest)(nil),  // 0: review.methods.GetReviewsRequest
	(*GetReviewsResponse)(nil), // 1: review.methods.GetReviewsResponse
	(ReviewStatus)(0),          // 2: review.messages.ReviewStatus
	(TargetType)(0),            // 3: review.messages.TargetType
	(*common.ApiResponse)(nil), // 4: common.ApiResponse
	(*NewReview)(nil),          // 5: review.messages.NewReview
}
var file_review_methods_get_proto_depIdxs = []int32{
	2, // 0: review.methods.GetReviewsRequest.status:type_name -> review.messages.ReviewStatus
	3, // 1: review.methods.GetReviewsRequest.type:type_name -> review.messages.TargetType
	4, // 2: review.methods.GetReviewsResponse.msg:type_name -> common.ApiResponse
	5, // 3: review.methods.GetReviewsResponse.reviews:type_name -> review.messages.NewReview
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_review_methods_get_proto_init() }
func file_review_methods_get_proto_init() {
	if File_review_methods_get_proto != nil {
		return
	}
	file_review_messages_status_proto_init()
	file_review_messages_type_proto_init()
	file_review_messages_new_review_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_review_methods_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_review_methods_get_proto_goTypes,
		DependencyIndexes: file_review_methods_get_proto_depIdxs,
		MessageInfos:      file_review_methods_get_proto_msgTypes,
	}.Build()
	File_review_methods_get_proto = out.File
	file_review_methods_get_proto_rawDesc = nil
	file_review_methods_get_proto_goTypes = nil
	file_review_methods_get_proto_depIdxs = nil
}
