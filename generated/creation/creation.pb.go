// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: creation/messages/creation.proto

package creation

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

type CreationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creation           *Creation           `protobuf:"bytes,1,opt,name=creation,proto3" json:"creation,omitempty"`
	CreationEngagement *CreationEngagement `protobuf:"bytes,2,opt,name=creation_engagement,json=creationEngagement,proto3" json:"creation_engagement,omitempty"`
	Category           *Category           `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CreationInfo) Reset() {
	*x = CreationInfo{}
	mi := &file_creation_messages_creation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreationInfo) ProtoMessage() {}

func (x *CreationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_creation_messages_creation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreationInfo.ProtoReflect.Descriptor instead.
func (*CreationInfo) Descriptor() ([]byte, []int) {
	return file_creation_messages_creation_proto_rawDescGZIP(), []int{0}
}

func (x *CreationInfo) GetCreation() *Creation {
	if x != nil {
		return x.Creation
	}
	return nil
}

func (x *CreationInfo) GetCreationEngagement() *CreationEngagement {
	if x != nil {
		return x.CreationEngagement
	}
	return nil
}

func (x *CreationInfo) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type Creation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId int64                  `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
	BaseInfo   *CreationUpload        `protobuf:"bytes,2,opt,name=base_info,json=baseInfo,proto3" json:"base_info,omitempty"`
	UploadTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=upload_time,json=uploadTime,proto3" json:"upload_time,omitempty"`
}

func (x *Creation) Reset() {
	*x = Creation{}
	mi := &file_creation_messages_creation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Creation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Creation) ProtoMessage() {}

func (x *Creation) ProtoReflect() protoreflect.Message {
	mi := &file_creation_messages_creation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Creation.ProtoReflect.Descriptor instead.
func (*Creation) Descriptor() ([]byte, []int) {
	return file_creation_messages_creation_proto_rawDescGZIP(), []int{1}
}

func (x *Creation) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

func (x *Creation) GetBaseInfo() *CreationUpload {
	if x != nil {
		return x.BaseInfo
	}
	return nil
}

func (x *Creation) GetUploadTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UploadTime
	}
	return nil
}

type CreationEngagement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId  int64                  `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
	Views       int32                  `protobuf:"varint,2,opt,name=views,proto3" json:"views,omitempty"`
	Likes       int32                  `protobuf:"varint,3,opt,name=likes,proto3" json:"likes,omitempty"`
	Saves       int32                  `protobuf:"varint,4,opt,name=saves,proto3" json:"saves,omitempty"`
	PublishTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
}

func (x *CreationEngagement) Reset() {
	*x = CreationEngagement{}
	mi := &file_creation_messages_creation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreationEngagement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreationEngagement) ProtoMessage() {}

func (x *CreationEngagement) ProtoReflect() protoreflect.Message {
	mi := &file_creation_messages_creation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreationEngagement.ProtoReflect.Descriptor instead.
func (*CreationEngagement) Descriptor() ([]byte, []int) {
	return file_creation_messages_creation_proto_rawDescGZIP(), []int{2}
}

func (x *CreationEngagement) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

func (x *CreationEngagement) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *CreationEngagement) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *CreationEngagement) GetSaves() int32 {
	if x != nil {
		return x.Saves
	}
	return 0
}

func (x *CreationEngagement) GetPublishTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishTime
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId  int32  `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Parent      int32  `protobuf:"varint,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	mi := &file_creation_messages_creation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_creation_messages_creation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_creation_messages_creation_proto_rawDescGZIP(), []int{3}
}

func (x *Category) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Category) GetParent() int32 {
	if x != nil {
		return x.Parent
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AnyCreation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyCreation []*Creation `protobuf:"bytes,1,rep,name=any_creation,json=anyCreation,proto3" json:"any_creation,omitempty"`
}

func (x *AnyCreation) Reset() {
	*x = AnyCreation{}
	mi := &file_creation_messages_creation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnyCreation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyCreation) ProtoMessage() {}

func (x *AnyCreation) ProtoReflect() protoreflect.Message {
	mi := &file_creation_messages_creation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyCreation.ProtoReflect.Descriptor instead.
func (*AnyCreation) Descriptor() ([]byte, []int) {
	return file_creation_messages_creation_proto_rawDescGZIP(), []int{4}
}

func (x *AnyCreation) GetAnyCreation() []*Creation {
	if x != nil {
		return x.AnyCreation
	}
	return nil
}

type AnyCreationEngagement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyCount []*CreationEngagement `protobuf:"bytes,1,rep,name=any_count,json=anyCount,proto3" json:"any_count,omitempty"`
}

func (x *AnyCreationEngagement) Reset() {
	*x = AnyCreationEngagement{}
	mi := &file_creation_messages_creation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnyCreationEngagement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyCreationEngagement) ProtoMessage() {}

func (x *AnyCreationEngagement) ProtoReflect() protoreflect.Message {
	mi := &file_creation_messages_creation_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyCreationEngagement.ProtoReflect.Descriptor instead.
func (*AnyCreationEngagement) Descriptor() ([]byte, []int) {
	return file_creation_messages_creation_proto_rawDescGZIP(), []int{5}
}

func (x *AnyCreationEngagement) GetAnyCount() []*CreationEngagement {
	if x != nil {
		return x.AnyCount
	}
	return nil
}

var File_creation_messages_creation_proto protoreflect.FileDescriptor

var file_creation_messages_creation_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd8, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x37, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x13, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x12, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0xa8, 0x01, 0x0a, 0x08, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x76,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x61, 0x76, 0x65, 0x73, 0x12,
	0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x79,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4d, 0x0a, 0x0b, 0x41, 0x6e, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x61, 0x6e, 0x79, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x6e, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x15, 0x41, 0x6e, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x42, 0x0a, 0x09, 0x61, 0x6e, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x6e, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creation_messages_creation_proto_rawDescOnce sync.Once
	file_creation_messages_creation_proto_rawDescData = file_creation_messages_creation_proto_rawDesc
)

func file_creation_messages_creation_proto_rawDescGZIP() []byte {
	file_creation_messages_creation_proto_rawDescOnce.Do(func() {
		file_creation_messages_creation_proto_rawDescData = protoimpl.X.CompressGZIP(file_creation_messages_creation_proto_rawDescData)
	})
	return file_creation_messages_creation_proto_rawDescData
}

var file_creation_messages_creation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_creation_messages_creation_proto_goTypes = []any{
	(*CreationInfo)(nil),          // 0: creation.messages.CreationInfo
	(*Creation)(nil),              // 1: creation.messages.Creation
	(*CreationEngagement)(nil),    // 2: creation.messages.CreationEngagement
	(*Category)(nil),              // 3: creation.messages.Category
	(*AnyCreation)(nil),           // 4: creation.messages.AnyCreation
	(*AnyCreationEngagement)(nil), // 5: creation.messages.AnyCreationEngagement
	(*CreationUpload)(nil),        // 6: creation.messages.CreationUpload
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_creation_messages_creation_proto_depIdxs = []int32{
	1, // 0: creation.messages.CreationInfo.creation:type_name -> creation.messages.Creation
	2, // 1: creation.messages.CreationInfo.creation_engagement:type_name -> creation.messages.CreationEngagement
	3, // 2: creation.messages.CreationInfo.category:type_name -> creation.messages.Category
	6, // 3: creation.messages.Creation.base_info:type_name -> creation.messages.CreationUpload
	7, // 4: creation.messages.Creation.upload_time:type_name -> google.protobuf.Timestamp
	7, // 5: creation.messages.CreationEngagement.publish_time:type_name -> google.protobuf.Timestamp
	1, // 6: creation.messages.AnyCreation.any_creation:type_name -> creation.messages.Creation
	2, // 7: creation.messages.AnyCreationEngagement.any_count:type_name -> creation.messages.CreationEngagement
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_creation_messages_creation_proto_init() }
func file_creation_messages_creation_proto_init() {
	if File_creation_messages_creation_proto != nil {
		return
	}
	file_creation_messages_creation_upload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_creation_messages_creation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_creation_messages_creation_proto_goTypes,
		DependencyIndexes: file_creation_messages_creation_proto_depIdxs,
		MessageInfos:      file_creation_messages_creation_proto_msgTypes,
	}.Build()
	File_creation_messages_creation_proto = out.File
	file_creation_messages_creation_proto_rawDesc = nil
	file_creation_messages_creation_proto_goTypes = nil
	file_creation_messages_creation_proto_depIdxs = nil
}
