// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: decider/decider.proto

package decider

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShouldSendMessageRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	NetworkId           uint32                 `protobuf:"varint,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	SourceChainId       []byte                 `protobuf:"bytes,2,opt,name=source_chain_id,json=sourceChainId,proto3" json:"source_chain_id,omitempty"`
	Payload             []byte                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	BytesRepresentation []byte                 `protobuf:"bytes,4,opt,name=bytes_representation,json=bytesRepresentation,proto3" json:"bytes_representation,omitempty"`
	Id                  []byte                 `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ShouldSendMessageRequest) Reset() {
	*x = ShouldSendMessageRequest{}
	mi := &file_decider_decider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShouldSendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShouldSendMessageRequest) ProtoMessage() {}

func (x *ShouldSendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_decider_decider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShouldSendMessageRequest.ProtoReflect.Descriptor instead.
func (*ShouldSendMessageRequest) Descriptor() ([]byte, []int) {
	return file_decider_decider_proto_rawDescGZIP(), []int{0}
}

func (x *ShouldSendMessageRequest) GetNetworkId() uint32 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *ShouldSendMessageRequest) GetSourceChainId() []byte {
	if x != nil {
		return x.SourceChainId
	}
	return nil
}

func (x *ShouldSendMessageRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ShouldSendMessageRequest) GetBytesRepresentation() []byte {
	if x != nil {
		return x.BytesRepresentation
	}
	return nil
}

func (x *ShouldSendMessageRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type ShouldSendMessageResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ShouldSendMessage bool                   `protobuf:"varint,1,opt,name=should_send_message,json=shouldSendMessage,proto3" json:"should_send_message,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ShouldSendMessageResponse) Reset() {
	*x = ShouldSendMessageResponse{}
	mi := &file_decider_decider_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShouldSendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShouldSendMessageResponse) ProtoMessage() {}

func (x *ShouldSendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_decider_decider_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShouldSendMessageResponse.ProtoReflect.Descriptor instead.
func (*ShouldSendMessageResponse) Descriptor() ([]byte, []int) {
	return file_decider_decider_proto_rawDescGZIP(), []int{1}
}

func (x *ShouldSendMessageResponse) GetShouldSendMessage() bool {
	if x != nil {
		return x.ShouldSendMessage
	}
	return false
}

var File_decider_decider_proto protoreflect.FileDescriptor

var file_decider_decider_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x72,
	0x22, 0xbe, 0x01, 0x0a, 0x18, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x31,
	0x0a, 0x14, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4b, 0x0a, 0x19, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x68, 0x6f,
	0x75, 0x6c, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x6c,
	0x0a, 0x0e, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5a, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x61, 0x2d, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x69, 0x63, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_decider_decider_proto_rawDescOnce sync.Once
	file_decider_decider_proto_rawDescData []byte
)

func file_decider_decider_proto_rawDescGZIP() []byte {
	file_decider_decider_proto_rawDescOnce.Do(func() {
		file_decider_decider_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_decider_decider_proto_rawDesc), len(file_decider_decider_proto_rawDesc)))
	})
	return file_decider_decider_proto_rawDescData
}

var file_decider_decider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_decider_decider_proto_goTypes = []any{
	(*ShouldSendMessageRequest)(nil),  // 0: decider.ShouldSendMessageRequest
	(*ShouldSendMessageResponse)(nil), // 1: decider.ShouldSendMessageResponse
}
var file_decider_decider_proto_depIdxs = []int32{
	0, // 0: decider.DeciderService.ShouldSendMessage:input_type -> decider.ShouldSendMessageRequest
	1, // 1: decider.DeciderService.ShouldSendMessage:output_type -> decider.ShouldSendMessageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_decider_decider_proto_init() }
func file_decider_decider_proto_init() {
	if File_decider_decider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_decider_decider_proto_rawDesc), len(file_decider_decider_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_decider_decider_proto_goTypes,
		DependencyIndexes: file_decider_decider_proto_depIdxs,
		MessageInfos:      file_decider_decider_proto_msgTypes,
	}.Build()
	File_decider_decider_proto = out.File
	file_decider_decider_proto_goTypes = nil
	file_decider_decider_proto_depIdxs = nil
}
