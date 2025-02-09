// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: review/messages/status.proto

package review

import (
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

type ReviewStatus int32

const (
	ReviewStatus_PENDING  ReviewStatus = 0
	ReviewStatus_APPROVED ReviewStatus = 1
	ReviewStatus_REJECTED ReviewStatus = 2
	ReviewStatus_DELETE   ReviewStatus = 3
)

// Enum value maps for ReviewStatus.
var (
	ReviewStatus_name = map[int32]string{
		0: "PENDING",
		1: "APPROVED",
		2: "REJECTED",
		3: "DELETE",
	}
	ReviewStatus_value = map[string]int32{
		"PENDING":  0,
		"APPROVED": 1,
		"REJECTED": 2,
		"DELETE":   3,
	}
)

func (x ReviewStatus) Enum() *ReviewStatus {
	p := new(ReviewStatus)
	*p = x
	return p
}

func (x ReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_review_messages_status_proto_enumTypes[0].Descriptor()
}

func (ReviewStatus) Type() protoreflect.EnumType {
	return &file_review_messages_status_proto_enumTypes[0]
}

func (x ReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReviewStatus.Descriptor instead.
func (ReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_review_messages_status_proto_rawDescGZIP(), []int{0}
}

var File_review_messages_status_proto protoreflect.FileDescriptor

var file_review_messages_status_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2a,
	0x43, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x03, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_messages_status_proto_rawDescOnce sync.Once
	file_review_messages_status_proto_rawDescData = file_review_messages_status_proto_rawDesc
)

func file_review_messages_status_proto_rawDescGZIP() []byte {
	file_review_messages_status_proto_rawDescOnce.Do(func() {
		file_review_messages_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_messages_status_proto_rawDescData)
	})
	return file_review_messages_status_proto_rawDescData
}

var file_review_messages_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_review_messages_status_proto_goTypes = []any{
	(ReviewStatus)(0), // 0: review.messages.ReviewStatus
}
var file_review_messages_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_review_messages_status_proto_init() }
func file_review_messages_status_proto_init() {
	if File_review_messages_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_review_messages_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_review_messages_status_proto_goTypes,
		DependencyIndexes: file_review_messages_status_proto_depIdxs,
		EnumInfos:         file_review_messages_status_proto_enumTypes,
	}.Build()
	File_review_messages_status_proto = out.File
	file_review_messages_status_proto_rawDesc = nil
	file_review_messages_status_proto_goTypes = nil
	file_review_messages_status_proto_depIdxs = nil
}
