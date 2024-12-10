// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: user/messages/user_gender.proto

package user

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

type UserGender int32

const (
	UserGender_UNDEFINED UserGender = 0
	UserGender_MALE      UserGender = 1
	UserGender_FEMALE    UserGender = 2
)

// Enum value maps for UserGender.
var (
	UserGender_name = map[int32]string{
		0: "UNDEFINED",
		1: "MALE",
		2: "FEMALE",
	}
	UserGender_value = map[string]int32{
		"UNDEFINED": 0,
		"MALE":      1,
		"FEMALE":    2,
	}
)

func (x UserGender) Enum() *UserGender {
	p := new(UserGender)
	*p = x
	return p
}

func (x UserGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserGender) Descriptor() protoreflect.EnumDescriptor {
	return file_user_messages_user_gender_proto_enumTypes[0].Descriptor()
}

func (UserGender) Type() protoreflect.EnumType {
	return &file_user_messages_user_gender_proto_enumTypes[0]
}

func (x UserGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserGender.Descriptor instead.
func (UserGender) EnumDescriptor() ([]byte, []int) {
	return file_user_messages_user_gender_proto_rawDescGZIP(), []int{0}
}

var File_user_messages_user_gender_proto protoreflect.FileDescriptor

var file_user_messages_user_gender_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2a, 0x31, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0d,
	0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c,
	0x45, 0x10, 0x02, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_messages_user_gender_proto_rawDescOnce sync.Once
	file_user_messages_user_gender_proto_rawDescData = file_user_messages_user_gender_proto_rawDesc
)

func file_user_messages_user_gender_proto_rawDescGZIP() []byte {
	file_user_messages_user_gender_proto_rawDescOnce.Do(func() {
		file_user_messages_user_gender_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_messages_user_gender_proto_rawDescData)
	})
	return file_user_messages_user_gender_proto_rawDescData
}

var file_user_messages_user_gender_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_messages_user_gender_proto_goTypes = []any{
	(UserGender)(0), // 0: user.messages.UserGender
}
var file_user_messages_user_gender_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_messages_user_gender_proto_init() }
func file_user_messages_user_gender_proto_init() {
	if File_user_messages_user_gender_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_messages_user_gender_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_messages_user_gender_proto_goTypes,
		DependencyIndexes: file_user_messages_user_gender_proto_depIdxs,
		EnumInfos:         file_user_messages_user_gender_proto_enumTypes,
	}.Build()
	File_user_messages_user_gender_proto = out.File
	file_user_messages_user_gender_proto_rawDesc = nil
	file_user_messages_user_gender_proto_goTypes = nil
	file_user_messages_user_gender_proto_depIdxs = nil
}
