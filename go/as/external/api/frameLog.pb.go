// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.13.0
// source: as/external/api/frameLog.proto

package api

import (
	_ "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/remote-tech/chirpstack-api/go/v3/common"
	gw "github.com/remote-tech/chirpstack-api/go/v3/gw"
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

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

// Enum value maps for RXWindow.
var (
	RXWindow_name = map[int32]string{
		0: "RX1",
		1: "RX2",
	}
	RXWindow_value = map[string]int32{
		"RX1": 0,
		"RX2": 1,
	}
)

func (x RXWindow) Enum() *RXWindow {
	p := new(RXWindow)
	*p = x
	return p
}

func (x RXWindow) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RXWindow) Descriptor() protoreflect.EnumDescriptor {
	return file_as_external_api_frameLog_proto_enumTypes[0].Descriptor()
}

func (RXWindow) Type() protoreflect.EnumType {
	return &file_as_external_api_frameLog_proto_enumTypes[0]
}

func (x RXWindow) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RXWindow.Descriptor instead.
func (RXWindow) EnumDescriptor() ([]byte, []int) {
	return file_as_external_api_frameLog_proto_rawDescGZIP(), []int{0}
}

type UplinkFrameLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TX information of the uplink.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,1,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RX information of the uplink.
	RxInfo []*gw.UplinkRXInfo `protobuf:"bytes,2,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	// LoRaWAN PHYPayload.
	PhyPayloadJson string `protobuf:"bytes,3,opt,name=phy_payload_json,json=phyPayloadJSON,proto3" json:"phy_payload_json,omitempty"`
	// Published at timestamp.
	PublishedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
}

func (x *UplinkFrameLog) Reset() {
	*x = UplinkFrameLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_frameLog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UplinkFrameLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UplinkFrameLog) ProtoMessage() {}

func (x *UplinkFrameLog) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_frameLog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UplinkFrameLog.ProtoReflect.Descriptor instead.
func (*UplinkFrameLog) Descriptor() ([]byte, []int) {
	return file_as_external_api_frameLog_proto_rawDescGZIP(), []int{0}
}

func (x *UplinkFrameLog) GetTxInfo() *gw.UplinkTXInfo {
	if x != nil {
		return x.TxInfo
	}
	return nil
}

func (x *UplinkFrameLog) GetRxInfo() []*gw.UplinkRXInfo {
	if x != nil {
		return x.RxInfo
	}
	return nil
}

func (x *UplinkFrameLog) GetPhyPayloadJson() string {
	if x != nil {
		return x.PhyPayloadJson
	}
	return ""
}

func (x *UplinkFrameLog) GetPublishedAt() *timestamp.Timestamp {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

type DownlinkFrameLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TX information of the downlink.
	TxInfo *gw.DownlinkTXInfo `protobuf:"bytes,1,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// LoRaWAN PHYPayload.
	PhyPayloadJson string `protobuf:"bytes,2,opt,name=phy_payload_json,json=phyPayloadJSON,proto3" json:"phy_payload_json,omitempty"`
	// Gateway ID.
	GatewayId string `protobuf:"bytes,3,opt,name=gateway_id,json=gatewayID,proto3" json:"gateway_id,omitempty"`
	// Published at timestamp.
	PublishedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
}

func (x *DownlinkFrameLog) Reset() {
	*x = DownlinkFrameLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_frameLog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownlinkFrameLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownlinkFrameLog) ProtoMessage() {}

func (x *DownlinkFrameLog) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_frameLog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownlinkFrameLog.ProtoReflect.Descriptor instead.
func (*DownlinkFrameLog) Descriptor() ([]byte, []int) {
	return file_as_external_api_frameLog_proto_rawDescGZIP(), []int{1}
}

func (x *DownlinkFrameLog) GetTxInfo() *gw.DownlinkTXInfo {
	if x != nil {
		return x.TxInfo
	}
	return nil
}

func (x *DownlinkFrameLog) GetPhyPayloadJson() string {
	if x != nil {
		return x.PhyPayloadJson
	}
	return ""
}

func (x *DownlinkFrameLog) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *DownlinkFrameLog) GetPublishedAt() *timestamp.Timestamp {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

var File_as_external_api_frameLog_proto protoreflect.FileDescriptor

var file_as_external_api_frameLog_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x67, 0x77, 0x2f,
	0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x29, 0x0a, 0x07, 0x74,
	0x78, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67,
	0x77, 0x2e, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x54, 0x58, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06,
	0x74, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x78, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x77, 0x2e, 0x55, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x52, 0x58, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x72, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x68, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x68, 0x79,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x3d, 0x0a, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc7, 0x01, 0x0a, 0x10, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x12,
	0x2b, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x77, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x54, 0x58,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x74, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10,
	0x70, 0x68, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6a, 0x73, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x68, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x41, 0x74, 0x2a, 0x1c, 0x0a, 0x08, 0x52, 0x58, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x12, 0x07, 0x0a, 0x03, 0x52, 0x58, 0x31, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x58, 0x32,
	0x10, 0x01, 0x42, 0x71, 0x0a, 0x21, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x73, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0d, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4c, 0x6f,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x74, 0x65, 0x63, 0x68,
	0x2f, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_as_external_api_frameLog_proto_rawDescOnce sync.Once
	file_as_external_api_frameLog_proto_rawDescData = file_as_external_api_frameLog_proto_rawDesc
)

func file_as_external_api_frameLog_proto_rawDescGZIP() []byte {
	file_as_external_api_frameLog_proto_rawDescOnce.Do(func() {
		file_as_external_api_frameLog_proto_rawDescData = protoimpl.X.CompressGZIP(file_as_external_api_frameLog_proto_rawDescData)
	})
	return file_as_external_api_frameLog_proto_rawDescData
}

var file_as_external_api_frameLog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_as_external_api_frameLog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_as_external_api_frameLog_proto_goTypes = []interface{}{
	(RXWindow)(0),               // 0: api.RXWindow
	(*UplinkFrameLog)(nil),      // 1: api.UplinkFrameLog
	(*DownlinkFrameLog)(nil),    // 2: api.DownlinkFrameLog
	(*gw.UplinkTXInfo)(nil),     // 3: gw.UplinkTXInfo
	(*gw.UplinkRXInfo)(nil),     // 4: gw.UplinkRXInfo
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*gw.DownlinkTXInfo)(nil),   // 6: gw.DownlinkTXInfo
}
var file_as_external_api_frameLog_proto_depIdxs = []int32{
	3, // 0: api.UplinkFrameLog.tx_info:type_name -> gw.UplinkTXInfo
	4, // 1: api.UplinkFrameLog.rx_info:type_name -> gw.UplinkRXInfo
	5, // 2: api.UplinkFrameLog.published_at:type_name -> google.protobuf.Timestamp
	6, // 3: api.DownlinkFrameLog.tx_info:type_name -> gw.DownlinkTXInfo
	5, // 4: api.DownlinkFrameLog.published_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_as_external_api_frameLog_proto_init() }
func file_as_external_api_frameLog_proto_init() {
	if File_as_external_api_frameLog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_as_external_api_frameLog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UplinkFrameLog); i {
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
		file_as_external_api_frameLog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownlinkFrameLog); i {
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
			RawDescriptor: file_as_external_api_frameLog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_as_external_api_frameLog_proto_goTypes,
		DependencyIndexes: file_as_external_api_frameLog_proto_depIdxs,
		EnumInfos:         file_as_external_api_frameLog_proto_enumTypes,
		MessageInfos:      file_as_external_api_frameLog_proto_msgTypes,
	}.Build()
	File_as_external_api_frameLog_proto = out.File
	file_as_external_api_frameLog_proto_rawDesc = nil
	file_as_external_api_frameLog_proto_goTypes = nil
	file_as_external_api_frameLog_proto_depIdxs = nil
}
