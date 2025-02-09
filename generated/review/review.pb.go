// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: review/messages/review.proto

package review

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

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	New        *NewReview             `protobuf:"bytes,1,opt,name=new,proto3" json:"new,omitempty"`
	ReviewerId int64                  `protobuf:"varint,2,opt,name=reviewer_id,json=reviewerId,proto3" json:"reviewer_id,omitempty"`
	Status     ReviewStatus           `protobuf:"varint,3,opt,name=status,proto3,enum=review.messages.ReviewStatus" json:"status,omitempty"` // 对象类型
	Remark     string                 `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	mi := &file_review_messages_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_review_messages_review_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_review_messages_review_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetNew() *NewReview {
	if x != nil {
		return x.New
	}
	return nil
}

func (x *Review) GetReviewerId() int64 {
	if x != nil {
		return x.ReviewerId
	}
	return 0
}

func (x *Review) GetStatus() ReviewStatus {
	if x != nil {
		return x.Status
	}
	return ReviewStatus_PENDING
}

func (x *Review) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Review) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_review_messages_review_proto protoreflect.FileDescriptor

var file_review_messages_review_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a,
	0x20, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe1, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x2c, 0x0a, 0x03, 0x6e,
	0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_messages_review_proto_rawDescOnce sync.Once
	file_review_messages_review_proto_rawDescData = file_review_messages_review_proto_rawDesc
)

func file_review_messages_review_proto_rawDescGZIP() []byte {
	file_review_messages_review_proto_rawDescOnce.Do(func() {
		file_review_messages_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_messages_review_proto_rawDescData)
	})
	return file_review_messages_review_proto_rawDescData
}

var file_review_messages_review_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_review_messages_review_proto_goTypes = []any{
	(*Review)(nil),                // 0: review.messages.Review
	(*NewReview)(nil),             // 1: review.messages.NewReview
	(ReviewStatus)(0),             // 2: review.messages.ReviewStatus
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_review_messages_review_proto_depIdxs = []int32{
	1, // 0: review.messages.Review.new:type_name -> review.messages.NewReview
	2, // 1: review.messages.Review.status:type_name -> review.messages.ReviewStatus
	3, // 2: review.messages.Review.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_review_messages_review_proto_init() }
func file_review_messages_review_proto_init() {
	if File_review_messages_review_proto != nil {
		return
	}
	file_review_messages_new_review_proto_init()
	file_review_messages_status_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_review_messages_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_review_messages_review_proto_goTypes,
		DependencyIndexes: file_review_messages_review_proto_depIdxs,
		MessageInfos:      file_review_messages_review_proto_msgTypes,
	}.Build()
	File_review_messages_review_proto = out.File
	file_review_messages_review_proto_rawDesc = nil
	file_review_messages_review_proto_goTypes = nil
	file_review_messages_review_proto_depIdxs = nil
}
