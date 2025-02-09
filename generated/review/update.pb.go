// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: review/methods/update.proto

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

type DelReviewerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewerId  int64               `protobuf:"varint,1,opt,name=reviewer_id,json=reviewerId,proto3" json:"reviewer_id,omitempty"`
	AccessToken *common.AccessToken `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *DelReviewerRequest) Reset() {
	*x = DelReviewerRequest{}
	mi := &file_review_methods_update_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelReviewerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelReviewerRequest) ProtoMessage() {}

func (x *DelReviewerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_methods_update_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelReviewerRequest.ProtoReflect.Descriptor instead.
func (*DelReviewerRequest) Descriptor() ([]byte, []int) {
	return file_review_methods_update_proto_rawDescGZIP(), []int{0}
}

func (x *DelReviewerRequest) GetReviewerId() int64 {
	if x != nil {
		return x.ReviewerId
	}
	return 0
}

func (x *DelReviewerRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type DelReviewerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *common.ApiResponse `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DelReviewerResponse) Reset() {
	*x = DelReviewerResponse{}
	mi := &file_review_methods_update_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelReviewerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelReviewerResponse) ProtoMessage() {}

func (x *DelReviewerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_methods_update_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelReviewerResponse.ProtoReflect.Descriptor instead.
func (*DelReviewerResponse) Descriptor() ([]byte, []int) {
	return file_review_methods_update_proto_rawDescGZIP(), []int{1}
}

func (x *DelReviewerResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

type UpdateReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Review      *Review             `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	AccessToken *common.AccessToken `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *UpdateReviewRequest) Reset() {
	*x = UpdateReviewRequest{}
	mi := &file_review_methods_update_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewRequest) ProtoMessage() {}

func (x *UpdateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_methods_update_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewRequest.ProtoReflect.Descriptor instead.
func (*UpdateReviewRequest) Descriptor() ([]byte, []int) {
	return file_review_methods_update_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateReviewRequest) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

func (x *UpdateReviewRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type UpdateReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *common.ApiResponse `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateReviewResponse) Reset() {
	*x = UpdateReviewResponse{}
	mi := &file_review_methods_update_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewResponse) ProtoMessage() {}

func (x *UpdateReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_methods_update_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewResponse.ProtoReflect.Descriptor instead.
func (*UpdateReviewResponse) Descriptor() ([]byte, []int) {
	return file_review_methods_update_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReviewResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_review_methods_update_proto protoreflect.FileDescriptor

var file_review_methods_update_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x1c, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x3c, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x7e,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37,
	0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_methods_update_proto_rawDescOnce sync.Once
	file_review_methods_update_proto_rawDescData = file_review_methods_update_proto_rawDesc
)

func file_review_methods_update_proto_rawDescGZIP() []byte {
	file_review_methods_update_proto_rawDescOnce.Do(func() {
		file_review_methods_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_methods_update_proto_rawDescData)
	})
	return file_review_methods_update_proto_rawDescData
}

var file_review_methods_update_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_review_methods_update_proto_goTypes = []any{
	(*DelReviewerRequest)(nil),   // 0: review.methods.DelReviewerRequest
	(*DelReviewerResponse)(nil),  // 1: review.methods.DelReviewerResponse
	(*UpdateReviewRequest)(nil),  // 2: review.methods.UpdateReviewRequest
	(*UpdateReviewResponse)(nil), // 3: review.methods.UpdateReviewResponse
	(*common.AccessToken)(nil),   // 4: common.AccessToken
	(*common.ApiResponse)(nil),   // 5: common.ApiResponse
	(*Review)(nil),               // 6: review.messages.Review
}
var file_review_methods_update_proto_depIdxs = []int32{
	4, // 0: review.methods.DelReviewerRequest.access_token:type_name -> common.AccessToken
	5, // 1: review.methods.DelReviewerResponse.msg:type_name -> common.ApiResponse
	6, // 2: review.methods.UpdateReviewRequest.review:type_name -> review.messages.Review
	4, // 3: review.methods.UpdateReviewRequest.access_token:type_name -> common.AccessToken
	5, // 4: review.methods.UpdateReviewResponse.msg:type_name -> common.ApiResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_review_methods_update_proto_init() }
func file_review_methods_update_proto_init() {
	if File_review_methods_update_proto != nil {
		return
	}
	file_review_messages_review_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_review_methods_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_review_methods_update_proto_goTypes,
		DependencyIndexes: file_review_methods_update_proto_depIdxs,
		MessageInfos:      file_review_methods_update_proto_msgTypes,
	}.Build()
	File_review_methods_update_proto = out.File
	file_review_methods_update_proto_rawDesc = nil
	file_review_methods_update_proto_goTypes = nil
	file_review_methods_update_proto_depIdxs = nil
}
