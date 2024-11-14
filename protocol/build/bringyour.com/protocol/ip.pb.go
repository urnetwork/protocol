// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.23.4
// source: ip.proto

package protocol

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

// a raw packet
// the provider is expected to do its own parsing to confirm the data verbatim
type IpPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketBytes []byte `protobuf:"bytes,1,opt,name=packet_bytes,json=packetBytes,proto3" json:"packet_bytes,omitempty"`
}

func (x *IpPacket) Reset() {
	*x = IpPacket{}
	mi := &file_ip_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IpPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPacket) ProtoMessage() {}

func (x *IpPacket) ProtoReflect() protoreflect.Message {
	mi := &file_ip_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpPacket.ProtoReflect.Descriptor instead.
func (*IpPacket) Descriptor() ([]byte, []int) {
	return file_ip_proto_rawDescGZIP(), []int{0}
}

func (x *IpPacket) GetPacketBytes() []byte {
	if x != nil {
		return x.PacketBytes
	}
	return nil
}

type IpPacketToProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpPacket *IpPacket `protobuf:"bytes,1,opt,name=ip_packet,json=ipPacket,proto3" json:"ip_packet,omitempty"`
}

func (x *IpPacketToProvider) Reset() {
	*x = IpPacketToProvider{}
	mi := &file_ip_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IpPacketToProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPacketToProvider) ProtoMessage() {}

func (x *IpPacketToProvider) ProtoReflect() protoreflect.Message {
	mi := &file_ip_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpPacketToProvider.ProtoReflect.Descriptor instead.
func (*IpPacketToProvider) Descriptor() ([]byte, []int) {
	return file_ip_proto_rawDescGZIP(), []int{1}
}

func (x *IpPacketToProvider) GetIpPacket() *IpPacket {
	if x != nil {
		return x.IpPacket
	}
	return nil
}

type IpPacketFromProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpPacket *IpPacket `protobuf:"bytes,1,opt,name=ip_packet,json=ipPacket,proto3" json:"ip_packet,omitempty"`
}

func (x *IpPacketFromProvider) Reset() {
	*x = IpPacketFromProvider{}
	mi := &file_ip_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IpPacketFromProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPacketFromProvider) ProtoMessage() {}

func (x *IpPacketFromProvider) ProtoReflect() protoreflect.Message {
	mi := &file_ip_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpPacketFromProvider.ProtoReflect.Descriptor instead.
func (*IpPacketFromProvider) Descriptor() ([]byte, []int) {
	return file_ip_proto_rawDescGZIP(), []int{2}
}

func (x *IpPacketFromProvider) GetIpPacket() *IpPacket {
	if x != nil {
		return x.IpPacket
	}
	return nil
}

type IpPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IpPing) Reset() {
	*x = IpPing{}
	mi := &file_ip_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IpPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPing) ProtoMessage() {}

func (x *IpPing) ProtoReflect() protoreflect.Message {
	mi := &file_ip_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpPing.ProtoReflect.Descriptor instead.
func (*IpPing) Descriptor() ([]byte, []int) {
	return file_ip_proto_rawDescGZIP(), []int{3}
}

var File_ip_proto protoreflect.FileDescriptor

var file_ip_proto_rawDesc = []byte{
	0x0a, 0x08, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x72, 0x69, 0x6e,
	0x67, 0x79, 0x6f, 0x75, 0x72, 0x22, 0x2d, 0x0a, 0x08, 0x49, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x49, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x54, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x69, 0x70,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x62, 0x72, 0x69, 0x6e, 0x67, 0x79, 0x6f, 0x75, 0x72, 0x2e, 0x49, 0x70, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x08, 0x69, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x48, 0x0a, 0x14,
	0x49, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x69, 0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x72, 0x69, 0x6e, 0x67, 0x79,
	0x6f, 0x75, 0x72, 0x2e, 0x49, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x08, 0x69, 0x70,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x08, 0x0a, 0x06, 0x49, 0x70, 0x50, 0x69, 0x6e, 0x67,
	0x42, 0x18, 0x5a, 0x16, 0x62, 0x72, 0x69, 0x6e, 0x67, 0x79, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ip_proto_rawDescOnce sync.Once
	file_ip_proto_rawDescData = file_ip_proto_rawDesc
)

func file_ip_proto_rawDescGZIP() []byte {
	file_ip_proto_rawDescOnce.Do(func() {
		file_ip_proto_rawDescData = protoimpl.X.CompressGZIP(file_ip_proto_rawDescData)
	})
	return file_ip_proto_rawDescData
}

var file_ip_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ip_proto_goTypes = []any{
	(*IpPacket)(nil),             // 0: bringyour.IpPacket
	(*IpPacketToProvider)(nil),   // 1: bringyour.IpPacketToProvider
	(*IpPacketFromProvider)(nil), // 2: bringyour.IpPacketFromProvider
	(*IpPing)(nil),               // 3: bringyour.IpPing
}
var file_ip_proto_depIdxs = []int32{
	0, // 0: bringyour.IpPacketToProvider.ip_packet:type_name -> bringyour.IpPacket
	0, // 1: bringyour.IpPacketFromProvider.ip_packet:type_name -> bringyour.IpPacket
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ip_proto_init() }
func file_ip_proto_init() {
	if File_ip_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ip_proto_goTypes,
		DependencyIndexes: file_ip_proto_depIdxs,
		MessageInfos:      file_ip_proto_msgTypes,
	}.Build()
	File_ip_proto = out.File
	file_ip_proto_rawDesc = nil
	file_ip_proto_goTypes = nil
	file_ip_proto_depIdxs = nil
}
