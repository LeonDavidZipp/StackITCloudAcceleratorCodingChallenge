// Code generated with goa v3.19.1, DO NOT EDIT.
//
// ForwarderService protocol buffer definition
//
// Command:
// $ goa gen
// github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/design
// --output ./src/api

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: goagen_api_forwarder_service.proto

package forwarder_servicepb

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

type ForwardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of notification
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The name of the event
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the event
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ForwardRequest) Reset() {
	*x = ForwardRequest{}
	mi := &file_goagen_api_forwarder_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForwardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardRequest) ProtoMessage() {}

func (x *ForwardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_api_forwarder_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardRequest.ProtoReflect.Descriptor instead.
func (*ForwardRequest) Descriptor() ([]byte, []int) {
	return file_goagen_api_forwarder_service_proto_rawDescGZIP(), []int{0}
}

func (x *ForwardRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ForwardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ForwardRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ForwardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForwardResponse) Reset() {
	*x = ForwardResponse{}
	mi := &file_goagen_api_forwarder_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForwardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardResponse) ProtoMessage() {}

func (x *ForwardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_api_forwarder_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardResponse.ProtoReflect.Descriptor instead.
func (*ForwardResponse) Descriptor() ([]byte, []int) {
	return file_goagen_api_forwarder_service_proto_rawDescGZIP(), []int{1}
}

var File_goagen_api_forwarder_service_proto protoreflect.FileDescriptor

var file_goagen_api_forwarder_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x61, 0x67, 0x65, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5a, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x64, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14,
	0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goagen_api_forwarder_service_proto_rawDescOnce sync.Once
	file_goagen_api_forwarder_service_proto_rawDescData = file_goagen_api_forwarder_service_proto_rawDesc
)

func file_goagen_api_forwarder_service_proto_rawDescGZIP() []byte {
	file_goagen_api_forwarder_service_proto_rawDescOnce.Do(func() {
		file_goagen_api_forwarder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_goagen_api_forwarder_service_proto_rawDescData)
	})
	return file_goagen_api_forwarder_service_proto_rawDescData
}

var file_goagen_api_forwarder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goagen_api_forwarder_service_proto_goTypes = []any{
	(*ForwardRequest)(nil),  // 0: forwarder_service.ForwardRequest
	(*ForwardResponse)(nil), // 1: forwarder_service.ForwardResponse
}
var file_goagen_api_forwarder_service_proto_depIdxs = []int32{
	0, // 0: forwarder_service.ForwarderService.Forward:input_type -> forwarder_service.ForwardRequest
	1, // 1: forwarder_service.ForwarderService.Forward:output_type -> forwarder_service.ForwardResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goagen_api_forwarder_service_proto_init() }
func file_goagen_api_forwarder_service_proto_init() {
	if File_goagen_api_forwarder_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goagen_api_forwarder_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_api_forwarder_service_proto_goTypes,
		DependencyIndexes: file_goagen_api_forwarder_service_proto_depIdxs,
		MessageInfos:      file_goagen_api_forwarder_service_proto_msgTypes,
	}.Build()
	File_goagen_api_forwarder_service_proto = out.File
	file_goagen_api_forwarder_service_proto_rawDesc = nil
	file_goagen_api_forwarder_service_proto_goTypes = nil
	file_goagen_api_forwarder_service_proto_depIdxs = nil
}
