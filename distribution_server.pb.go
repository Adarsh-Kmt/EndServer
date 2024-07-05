// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: distribution_server.proto

package main

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

type DistributionServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId   string `protobuf:"bytes,1,opt,name=senderId,proto3" json:"senderId,omitempty"`
	ReceiverId string `protobuf:"bytes,2,opt,name=receiverId,proto3" json:"receiverId,omitempty"`
	Body       string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DistributionServerMessage) Reset() {
	*x = DistributionServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribution_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionServerMessage) ProtoMessage() {}

func (x *DistributionServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_distribution_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionServerMessage.ProtoReflect.Descriptor instead.
func (*DistributionServerMessage) Descriptor() ([]byte, []int) {
	return file_distribution_server_proto_rawDescGZIP(), []int{0}
}

func (x *DistributionServerMessage) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *DistributionServerMessage) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

func (x *DistributionServerMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type DistributionServerConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	EndServerAddress string `protobuf:"bytes,2,opt,name=EndServerAddress,proto3" json:"EndServerAddress,omitempty"`
}

func (x *DistributionServerConnectionRequest) Reset() {
	*x = DistributionServerConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribution_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionServerConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionServerConnectionRequest) ProtoMessage() {}

func (x *DistributionServerConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_distribution_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionServerConnectionRequest.ProtoReflect.Descriptor instead.
func (*DistributionServerConnectionRequest) Descriptor() ([]byte, []int) {
	return file_distribution_server_proto_rawDescGZIP(), []int{1}
}

func (x *DistributionServerConnectionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DistributionServerConnectionRequest) GetEndServerAddress() string {
	if x != nil {
		return x.EndServerAddress
	}
	return ""
}

type DistributionServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseStatus int32 `protobuf:"varint,1,opt,name=responseStatus,proto3" json:"responseStatus,omitempty"`
}

func (x *DistributionServerResponse) Reset() {
	*x = DistributionServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribution_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionServerResponse) ProtoMessage() {}

func (x *DistributionServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_distribution_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionServerResponse.ProtoReflect.Descriptor instead.
func (*DistributionServerResponse) Descriptor() ([]byte, []int) {
	return file_distribution_server_proto_rawDescGZIP(), []int{2}
}

func (x *DistributionServerResponse) GetResponseStatus() int32 {
	if x != nil {
		return x.ResponseStatus
	}
	return 0
}

var File_distribution_server_proto protoreflect.FileDescriptor

var file_distribution_server_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x19, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x69, 0x0a, 0x23, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x6e, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x1a, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9b, 0x02, 0x0a, 0x20, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x2e, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x24, 0x2e, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x12, 0x24, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x61, 0x72, 0x73, 0x68, 0x2d, 0x4b, 0x6d, 0x74,
	0x2f, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_distribution_server_proto_rawDescOnce sync.Once
	file_distribution_server_proto_rawDescData = file_distribution_server_proto_rawDesc
)

func file_distribution_server_proto_rawDescGZIP() []byte {
	file_distribution_server_proto_rawDescOnce.Do(func() {
		file_distribution_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_distribution_server_proto_rawDescData)
	})
	return file_distribution_server_proto_rawDescData
}

var file_distribution_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_distribution_server_proto_goTypes = []interface{}{
	(*DistributionServerMessage)(nil),           // 0: DistributionServerMessage
	(*DistributionServerConnectionRequest)(nil), // 1: DistributionServerConnectionRequest
	(*DistributionServerResponse)(nil),          // 2: DistributionServerResponse
}
var file_distribution_server_proto_depIdxs = []int32{
	0, // 0: DistributionServerMessageService.SendMessage:input_type -> DistributionServerMessage
	1, // 1: DistributionServerMessageService.UserConnected:input_type -> DistributionServerConnectionRequest
	1, // 2: DistributionServerMessageService.UserDisconnected:input_type -> DistributionServerConnectionRequest
	2, // 3: DistributionServerMessageService.SendMessage:output_type -> DistributionServerResponse
	2, // 4: DistributionServerMessageService.UserConnected:output_type -> DistributionServerResponse
	2, // 5: DistributionServerMessageService.UserDisconnected:output_type -> DistributionServerResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_distribution_server_proto_init() }
func file_distribution_server_proto_init() {
	if File_distribution_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_distribution_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionServerMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_distribution_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionServerConnectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_distribution_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionServerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_distribution_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_distribution_server_proto_goTypes,
		DependencyIndexes: file_distribution_server_proto_depIdxs,
		MessageInfos:      file_distribution_server_proto_msgTypes,
	}.Build()
	File_distribution_server_proto = out.File
	file_distribution_server_proto_rawDesc = nil
	file_distribution_server_proto_goTypes = nil
	file_distribution_server_proto_depIdxs = nil
}
