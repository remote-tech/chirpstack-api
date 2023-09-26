// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nc/nc.proto

package nc

import (
	context "context"
	fmt "fmt"
	gw "github.com/remote-tech/chirpstack-api/go/v3/gw"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MType int32

const (
	MType_UNKNOWN               MType = 0
	MType_JOIN_REQUEST          MType = 1
	MType_JOIN_ACCEPT           MType = 2
	MType_UNCONFIRMED_DATA_UP   MType = 3
	MType_UNCONFIRMED_DATA_DOWN MType = 4
	MType_CONFIRMED_DATA_UP     MType = 5
	MType_CONFIRMED_DATA_DOWN   MType = 6
	MType_REJOIN_REQUEST        MType = 7
)

var MType_name = map[int32]string{
	0: "UNKNOWN",
	1: "JOIN_REQUEST",
	2: "JOIN_ACCEPT",
	3: "UNCONFIRMED_DATA_UP",
	4: "UNCONFIRMED_DATA_DOWN",
	5: "CONFIRMED_DATA_UP",
	6: "CONFIRMED_DATA_DOWN",
	7: "REJOIN_REQUEST",
}

var MType_value = map[string]int32{
	"UNKNOWN":               0,
	"JOIN_REQUEST":          1,
	"JOIN_ACCEPT":           2,
	"UNCONFIRMED_DATA_UP":   3,
	"UNCONFIRMED_DATA_DOWN": 4,
	"CONFIRMED_DATA_UP":     5,
	"CONFIRMED_DATA_DOWN":   6,
	"REJOIN_REQUEST":        7,
}

func (x MType) String() string {
	return proto.EnumName(MType_name, int32(x))
}

func (MType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7dd10f8ea2e14d19, []int{0}
}

type HandleUplinkMetaDataRequest struct {
	// Device EUI (8 bytes).
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// TX meta-data.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,2,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RX meta-data.
	RxInfo []*gw.UplinkRXInfo `protobuf:"bytes,3,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	// PHYPayload byte count.
	PhyPayloadByteCount uint32 `protobuf:"varint,4,opt,name=phy_payload_byte_count,json=phyPayloadByteCount,proto3" json:"phy_payload_byte_count,omitempty"`
	// MAC-Command byte count.
	MacCommandByteCount uint32 `protobuf:"varint,5,opt,name=mac_command_byte_count,json=macCommandByteCount,proto3" json:"mac_command_byte_count,omitempty"`
	// Application payload byte count.
	ApplicationPayloadByteCount uint32 `protobuf:"varint,6,opt,name=application_payload_byte_count,json=applicationPayloadByteCount,proto3" json:"application_payload_byte_count,omitempty"`
	// Message type.
	MessageType          MType    `protobuf:"varint,7,opt,name=message_type,json=messageType,proto3,enum=nc.MType" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleUplinkMetaDataRequest) Reset()         { *m = HandleUplinkMetaDataRequest{} }
func (m *HandleUplinkMetaDataRequest) String() string { return proto.CompactTextString(m) }
func (*HandleUplinkMetaDataRequest) ProtoMessage()    {}
func (*HandleUplinkMetaDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dd10f8ea2e14d19, []int{0}
}

func (m *HandleUplinkMetaDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleUplinkMetaDataRequest.Unmarshal(m, b)
}
func (m *HandleUplinkMetaDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleUplinkMetaDataRequest.Marshal(b, m, deterministic)
}
func (m *HandleUplinkMetaDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleUplinkMetaDataRequest.Merge(m, src)
}
func (m *HandleUplinkMetaDataRequest) XXX_Size() int {
	return xxx_messageInfo_HandleUplinkMetaDataRequest.Size(m)
}
func (m *HandleUplinkMetaDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleUplinkMetaDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleUplinkMetaDataRequest proto.InternalMessageInfo

func (m *HandleUplinkMetaDataRequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *HandleUplinkMetaDataRequest) GetTxInfo() *gw.UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *HandleUplinkMetaDataRequest) GetRxInfo() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

func (m *HandleUplinkMetaDataRequest) GetPhyPayloadByteCount() uint32 {
	if m != nil {
		return m.PhyPayloadByteCount
	}
	return 0
}

func (m *HandleUplinkMetaDataRequest) GetMacCommandByteCount() uint32 {
	if m != nil {
		return m.MacCommandByteCount
	}
	return 0
}

func (m *HandleUplinkMetaDataRequest) GetApplicationPayloadByteCount() uint32 {
	if m != nil {
		return m.ApplicationPayloadByteCount
	}
	return 0
}

func (m *HandleUplinkMetaDataRequest) GetMessageType() MType {
	if m != nil {
		return m.MessageType
	}
	return MType_UNKNOWN
}

type HandleDownlinkMetaDataRequest struct {
	// Device EUI (8 bytes).
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Multicast Group ID (UUID).
	MulticastGroupId []byte `protobuf:"bytes,2,opt,name=multicast_group_id,json=multicastGroupId,proto3" json:"multicast_group_id,omitempty"`
	// TX meta-data.
	TxInfo *gw.DownlinkTXInfo `protobuf:"bytes,3,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// PHYPayload byte count.
	PhyPayloadByteCount uint32 `protobuf:"varint,4,opt,name=phy_payload_byte_count,json=phyPayloadByteCount,proto3" json:"phy_payload_byte_count,omitempty"`
	// MAC-Command byte count.
	MacCommandByteCount uint32 `protobuf:"varint,5,opt,name=mac_command_byte_count,json=macCommandByteCount,proto3" json:"mac_command_byte_count,omitempty"`
	// Application payload byte count.
	ApplicationPayloadByteCount uint32 `protobuf:"varint,6,opt,name=application_payload_byte_count,json=applicationPayloadByteCount,proto3" json:"application_payload_byte_count,omitempty"`
	// Message type.
	MessageType MType `protobuf:"varint,7,opt,name=message_type,json=messageType,proto3,enum=nc.MType" json:"message_type,omitempty"`
	// Gateway ID.
	GatewayId            []byte   `protobuf:"bytes,8,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleDownlinkMetaDataRequest) Reset()         { *m = HandleDownlinkMetaDataRequest{} }
func (m *HandleDownlinkMetaDataRequest) String() string { return proto.CompactTextString(m) }
func (*HandleDownlinkMetaDataRequest) ProtoMessage()    {}
func (*HandleDownlinkMetaDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dd10f8ea2e14d19, []int{1}
}

func (m *HandleDownlinkMetaDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleDownlinkMetaDataRequest.Unmarshal(m, b)
}
func (m *HandleDownlinkMetaDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleDownlinkMetaDataRequest.Marshal(b, m, deterministic)
}
func (m *HandleDownlinkMetaDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleDownlinkMetaDataRequest.Merge(m, src)
}
func (m *HandleDownlinkMetaDataRequest) XXX_Size() int {
	return xxx_messageInfo_HandleDownlinkMetaDataRequest.Size(m)
}
func (m *HandleDownlinkMetaDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleDownlinkMetaDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleDownlinkMetaDataRequest proto.InternalMessageInfo

func (m *HandleDownlinkMetaDataRequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *HandleDownlinkMetaDataRequest) GetMulticastGroupId() []byte {
	if m != nil {
		return m.MulticastGroupId
	}
	return nil
}

func (m *HandleDownlinkMetaDataRequest) GetTxInfo() *gw.DownlinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *HandleDownlinkMetaDataRequest) GetPhyPayloadByteCount() uint32 {
	if m != nil {
		return m.PhyPayloadByteCount
	}
	return 0
}

func (m *HandleDownlinkMetaDataRequest) GetMacCommandByteCount() uint32 {
	if m != nil {
		return m.MacCommandByteCount
	}
	return 0
}

func (m *HandleDownlinkMetaDataRequest) GetApplicationPayloadByteCount() uint32 {
	if m != nil {
		return m.ApplicationPayloadByteCount
	}
	return 0
}

func (m *HandleDownlinkMetaDataRequest) GetMessageType() MType {
	if m != nil {
		return m.MessageType
	}
	return MType_UNKNOWN
}

func (m *HandleDownlinkMetaDataRequest) GetGatewayId() []byte {
	if m != nil {
		return m.GatewayId
	}
	return nil
}

type HandleUplinkMACCommandRequest struct {
	// Device EUI (8 bytes).
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Command identifier (specified by the LoRaWAN specs).
	Cid uint32 `protobuf:"varint,2,opt,name=cid,proto3" json:"cid,omitempty"`
	// MAC-command payload(s).
	Commands             [][]byte `protobuf:"bytes,6,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleUplinkMACCommandRequest) Reset()         { *m = HandleUplinkMACCommandRequest{} }
func (m *HandleUplinkMACCommandRequest) String() string { return proto.CompactTextString(m) }
func (*HandleUplinkMACCommandRequest) ProtoMessage()    {}
func (*HandleUplinkMACCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dd10f8ea2e14d19, []int{2}
}

func (m *HandleUplinkMACCommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleUplinkMACCommandRequest.Unmarshal(m, b)
}
func (m *HandleUplinkMACCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleUplinkMACCommandRequest.Marshal(b, m, deterministic)
}
func (m *HandleUplinkMACCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleUplinkMACCommandRequest.Merge(m, src)
}
func (m *HandleUplinkMACCommandRequest) XXX_Size() int {
	return xxx_messageInfo_HandleUplinkMACCommandRequest.Size(m)
}
func (m *HandleUplinkMACCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleUplinkMACCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleUplinkMACCommandRequest proto.InternalMessageInfo

func (m *HandleUplinkMACCommandRequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *HandleUplinkMACCommandRequest) GetCid() uint32 {
	if m != nil {
		return m.Cid
	}
	return 0
}

func (m *HandleUplinkMACCommandRequest) GetCommands() [][]byte {
	if m != nil {
		return m.Commands
	}
	return nil
}

type HandleRejectedUplinkFrameSetRequest struct {
	FrameSet             *gw.UplinkFrameSet `protobuf:"bytes,1,opt,name=frame_set,json=frameSet,proto3" json:"frame_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HandleRejectedUplinkFrameSetRequest) Reset()         { *m = HandleRejectedUplinkFrameSetRequest{} }
func (m *HandleRejectedUplinkFrameSetRequest) String() string { return proto.CompactTextString(m) }
func (*HandleRejectedUplinkFrameSetRequest) ProtoMessage()    {}
func (*HandleRejectedUplinkFrameSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dd10f8ea2e14d19, []int{3}
}

func (m *HandleRejectedUplinkFrameSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleRejectedUplinkFrameSetRequest.Unmarshal(m, b)
}
func (m *HandleRejectedUplinkFrameSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleRejectedUplinkFrameSetRequest.Marshal(b, m, deterministic)
}
func (m *HandleRejectedUplinkFrameSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleRejectedUplinkFrameSetRequest.Merge(m, src)
}
func (m *HandleRejectedUplinkFrameSetRequest) XXX_Size() int {
	return xxx_messageInfo_HandleRejectedUplinkFrameSetRequest.Size(m)
}
func (m *HandleRejectedUplinkFrameSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleRejectedUplinkFrameSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleRejectedUplinkFrameSetRequest proto.InternalMessageInfo

func (m *HandleRejectedUplinkFrameSetRequest) GetFrameSet() *gw.UplinkFrameSet {
	if m != nil {
		return m.FrameSet
	}
	return nil
}

func init() {
	proto.RegisterEnum("nc.MType", MType_name, MType_value)
	proto.RegisterType((*HandleUplinkMetaDataRequest)(nil), "nc.HandleUplinkMetaDataRequest")
	proto.RegisterType((*HandleDownlinkMetaDataRequest)(nil), "nc.HandleDownlinkMetaDataRequest")
	proto.RegisterType((*HandleUplinkMACCommandRequest)(nil), "nc.HandleUplinkMACCommandRequest")
	proto.RegisterType((*HandleRejectedUplinkFrameSetRequest)(nil), "nc.HandleRejectedUplinkFrameSetRequest")
}

func init() {
	proto.RegisterFile("nc/nc.proto", fileDescriptor_7dd10f8ea2e14d19)
}

var fileDescriptor_7dd10f8ea2e14d19 = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xd1, 0x72, 0xd2, 0x40,
	0x14, 0x6d, 0xa0, 0x85, 0x76, 0xa1, 0x35, 0x6e, 0x5b, 0x8a, 0xd4, 0x2a, 0xe2, 0x83, 0x58, 0x6b,
	0x32, 0xd3, 0x7e, 0x01, 0x0d, 0x54, 0xd1, 0x29, 0xa5, 0x29, 0xa8, 0xe3, 0x38, 0x93, 0x59, 0x36,
	0x4b, 0x88, 0x4d, 0x76, 0x63, 0xd8, 0x40, 0xf3, 0x15, 0x7e, 0x8a, 0xdf, 0xe5, 0x4f, 0x38, 0xce,
	0x26, 0xa1, 0xb4, 0x85, 0xa2, 0xbe, 0xfa, 0x44, 0x76, 0xef, 0xb9, 0xe7, 0x84, 0x73, 0x6e, 0x2e,
	0xc8, 0x51, 0xac, 0x52, 0xac, 0x78, 0x3e, 0xe3, 0x0c, 0xa6, 0x28, 0x2e, 0xed, 0x5a, 0x8c, 0x59,
	0x0e, 0x51, 0xa3, 0x9b, 0x5e, 0xd0, 0x57, 0x89, 0xeb, 0xf1, 0x30, 0x06, 0x94, 0x72, 0xd6, 0x58,
	0xb5, 0xc6, 0xf1, 0xa1, 0xf2, 0x33, 0x05, 0x76, 0xdf, 0x22, 0x6a, 0x3a, 0xa4, 0xeb, 0x39, 0x36,
	0xbd, 0x3c, 0x25, 0x1c, 0xd5, 0x11, 0x47, 0x3a, 0xf9, 0x16, 0x90, 0x21, 0x87, 0x3b, 0x20, 0x6b,
	0x92, 0x91, 0x41, 0x02, 0xbb, 0x28, 0x95, 0xa5, 0x6a, 0x5e, 0xcf, 0x98, 0x64, 0xd4, 0x08, 0x6c,
	0xf8, 0x12, 0x64, 0xf9, 0x95, 0x61, 0xd3, 0x3e, 0x2b, 0xa6, 0xca, 0x52, 0x35, 0x77, 0x28, 0x2b,
	0xd6, 0x58, 0x89, 0x49, 0x3a, 0x9f, 0x9a, 0xb4, 0xcf, 0xf4, 0x0c, 0xbf, 0x12, 0xbf, 0x02, 0xea,
	0x27, 0xd0, 0x74, 0x39, 0x7d, 0x1b, 0xaa, 0x27, 0x50, 0x3f, 0x86, 0x1e, 0x81, 0x82, 0x37, 0x08,
	0x0d, 0x0f, 0x85, 0x0e, 0x43, 0xa6, 0xd1, 0x0b, 0x39, 0x31, 0x30, 0x0b, 0x28, 0x2f, 0x2e, 0x97,
	0xa5, 0xea, 0xba, 0xbe, 0xe9, 0x0d, 0xc2, 0x76, 0x5c, 0x3c, 0x0e, 0x39, 0xd1, 0x44, 0x49, 0x34,
	0xb9, 0x08, 0x1b, 0x98, 0xb9, 0x2e, 0xa2, 0xb7, 0x9a, 0x56, 0xe2, 0x26, 0x17, 0x61, 0x2d, 0x2e,
	0x4e, 0x9b, 0x34, 0xf0, 0x04, 0x79, 0x9e, 0x63, 0x63, 0xc4, 0x6d, 0x46, 0xe7, 0x29, 0x66, 0xa2,
	0xe6, 0xdd, 0x1b, 0xa8, 0x19, 0xe5, 0x03, 0x90, 0x77, 0xc9, 0x70, 0x88, 0x2c, 0x62, 0xf0, 0xd0,
	0x23, 0xc5, 0x6c, 0x59, 0xaa, 0x6e, 0x1c, 0xae, 0x29, 0x14, 0x2b, 0xa7, 0x9d, 0xd0, 0x23, 0x7a,
	0x2e, 0x29, 0x8b, 0x43, 0xe5, 0x7b, 0x1a, 0xec, 0xc5, 0x5e, 0xd7, 0xd9, 0x98, 0xfe, 0x93, 0xdb,
	0x07, 0x00, 0xba, 0x81, 0xc3, 0x6d, 0x8c, 0x86, 0xdc, 0xb0, 0x7c, 0x16, 0x78, 0x86, 0x6d, 0x46,
	0xc6, 0xe7, 0x75, 0xf9, 0xba, 0xf2, 0x46, 0x14, 0x9a, 0x26, 0x7c, 0x35, 0xcd, 0x26, 0x1d, 0x65,
	0x03, 0x85, 0xe1, 0x13, 0xd1, 0x3b, 0xe9, 0xfc, 0xcf, 0x96, 0xc3, 0x3d, 0x00, 0x2c, 0xc4, 0xc9,
	0x18, 0x85, 0xc2, 0xaf, 0xd5, 0xc8, 0xaf, 0xb5, 0xe4, 0xa6, 0x69, 0x56, 0xfa, 0x93, 0x40, 0x92,
	0xe1, 0xaf, 0x69, 0xc9, 0x4b, 0xff, 0x31, 0x10, 0x19, 0xa4, 0x71, 0x92, 0xc0, 0xba, 0x2e, 0x1e,
	0x61, 0x09, 0xac, 0x26, 0x76, 0x0c, 0x8b, 0x99, 0x72, 0xba, 0x9a, 0xd7, 0xaf, 0xcf, 0x95, 0x0f,
	0xe0, 0x79, 0xac, 0xa3, 0x93, 0xaf, 0x04, 0x73, 0x62, 0xc6, 0x7a, 0x27, 0x3e, 0x72, 0xc9, 0x05,
	0xe1, 0x13, 0x35, 0x15, 0xac, 0xf5, 0xc5, 0x95, 0x31, 0x24, 0x3c, 0xd2, 0x4b, 0x92, 0xbb, 0x83,
	0x5e, 0xed, 0x27, 0x4f, 0xfb, 0x3f, 0x24, 0xb0, 0x12, 0xfd, 0x6b, 0x98, 0x03, 0xd9, 0x6e, 0xeb,
	0x7d, 0xeb, 0xec, 0x63, 0x4b, 0x5e, 0x82, 0x32, 0xc8, 0xbf, 0x3b, 0x6b, 0xb6, 0x0c, 0xbd, 0x71,
	0xde, 0x6d, 0x5c, 0x74, 0x64, 0x09, 0x3e, 0x00, 0xb9, 0xe8, 0xa6, 0xa6, 0x69, 0x8d, 0x76, 0x47,
	0x4e, 0xc1, 0x1d, 0xb0, 0xd9, 0x6d, 0x69, 0x67, 0xad, 0x93, 0xa6, 0x7e, 0xda, 0xa8, 0x1b, 0xf5,
	0x5a, 0xa7, 0x66, 0x74, 0xdb, 0x72, 0x1a, 0x3e, 0x02, 0xdb, 0x33, 0x85, 0xba, 0xa0, 0x5d, 0x86,
	0xdb, 0xe0, 0xe1, 0x6c, 0xc7, 0x8a, 0xa0, 0x9a, 0x87, 0xcf, 0x40, 0x08, 0x36, 0xf4, 0xc6, 0xad,
	0x17, 0xc9, 0x1e, 0xfe, 0x4a, 0x81, 0x62, 0x8b, 0xf0, 0x31, 0xf3, 0x2f, 0x35, 0x46, 0xb9, 0xcf,
	0x1c, 0x87, 0xf8, 0x17, 0xc4, 0x1f, 0xd9, 0x98, 0xc0, 0x73, 0xb0, 0x35, 0x6f, 0x17, 0xc1, 0xa7,
	0x22, 0xdd, 0x05, 0x5b, 0xaa, 0x54, 0x50, 0xe2, 0x85, 0xa7, 0x4c, 0x16, 0x9e, 0xd2, 0x10, 0x0b,
	0xaf, 0xb2, 0x04, 0xbb, 0xa0, 0x30, 0xff, 0x93, 0x83, 0xcf, 0xa6, 0xa4, 0xf7, 0x7c, 0x8e, 0x7f,
	0x43, 0x7b, 0x77, 0x70, 0x6e, 0xd2, 0xde, 0x33, 0x54, 0x0b, 0x68, 0x11, 0x78, 0xbc, 0x68, 0x4e,
	0xe0, 0x8b, 0x29, 0xf9, 0xc2, 0x49, 0xba, 0x5f, 0xe2, 0xf8, 0x0b, 0xd8, 0xb2, 0x99, 0x82, 0x07,
	0xb6, 0xef, 0x0d, 0x39, 0xc2, 0x97, 0x0a, 0xf2, 0x6c, 0x85, 0xe2, 0xe3, 0xc2, 0x4c, 0x2a, 0x6d,
	0xd1, 0xda, 0x96, 0x3e, 0xef, 0x5b, 0x36, 0x1f, 0x04, 0x3d, 0x05, 0x33, 0x57, 0xed, 0xf9, 0x0c,
	0x23, 0xe4, 0xab, 0xd3, 0xfe, 0xd7, 0xc8, 0xb3, 0x55, 0x8b, 0xa9, 0xa3, 0x23, 0x95, 0xe2, 0x5e,
	0x26, 0xd2, 0x3b, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x67, 0x19, 0x6a, 0x61, 0x92, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NetworkControllerServiceClient is the client API for NetworkControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkControllerServiceClient interface {
	// HandleUplinkMetaData handles uplink meta-rata.
	HandleUplinkMetaData(ctx context.Context, in *HandleUplinkMetaDataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// HandleDownlinkMetaData handles downlink meta-data.
	HandleDownlinkMetaData(ctx context.Context, in *HandleDownlinkMetaDataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// HandleUplinkMACCommand handles an uplink mac-command.
	// This method will only be called in case the mac-command request was
	// enqueued throught the API or when the CID is >= 0x80 (proprietary
	// mac-command range).
	HandleUplinkMACCommand(ctx context.Context, in *HandleUplinkMACCommandRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// HandleRejectedUplinkFrameSet handles a rejected uplink.
	// And uplink can be rejected in the case the device has not (yet) been
	// provisioned, because of invalid frame-counter, MIC, ...
	HandleRejectedUplinkFrameSet(ctx context.Context, in *HandleRejectedUplinkFrameSetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkControllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkControllerServiceClient(cc grpc.ClientConnInterface) NetworkControllerServiceClient {
	return &networkControllerServiceClient{cc}
}

func (c *networkControllerServiceClient) HandleUplinkMetaData(ctx context.Context, in *HandleUplinkMetaDataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nc.NetworkControllerService/HandleUplinkMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkControllerServiceClient) HandleDownlinkMetaData(ctx context.Context, in *HandleDownlinkMetaDataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nc.NetworkControllerService/HandleDownlinkMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkControllerServiceClient) HandleUplinkMACCommand(ctx context.Context, in *HandleUplinkMACCommandRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nc.NetworkControllerService/HandleUplinkMACCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkControllerServiceClient) HandleRejectedUplinkFrameSet(ctx context.Context, in *HandleRejectedUplinkFrameSetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nc.NetworkControllerService/HandleRejectedUplinkFrameSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkControllerServiceServer is the server API for NetworkControllerService service.
type NetworkControllerServiceServer interface {
	// HandleUplinkMetaData handles uplink meta-rata.
	HandleUplinkMetaData(context.Context, *HandleUplinkMetaDataRequest) (*empty.Empty, error)
	// HandleDownlinkMetaData handles downlink meta-data.
	HandleDownlinkMetaData(context.Context, *HandleDownlinkMetaDataRequest) (*empty.Empty, error)
	// HandleUplinkMACCommand handles an uplink mac-command.
	// This method will only be called in case the mac-command request was
	// enqueued throught the API or when the CID is >= 0x80 (proprietary
	// mac-command range).
	HandleUplinkMACCommand(context.Context, *HandleUplinkMACCommandRequest) (*empty.Empty, error)
	// HandleRejectedUplinkFrameSet handles a rejected uplink.
	// And uplink can be rejected in the case the device has not (yet) been
	// provisioned, because of invalid frame-counter, MIC, ...
	HandleRejectedUplinkFrameSet(context.Context, *HandleRejectedUplinkFrameSetRequest) (*empty.Empty, error)
}

// UnimplementedNetworkControllerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkControllerServiceServer struct {
}

func (*UnimplementedNetworkControllerServiceServer) HandleUplinkMetaData(ctx context.Context, req *HandleUplinkMetaDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUplinkMetaData not implemented")
}
func (*UnimplementedNetworkControllerServiceServer) HandleDownlinkMetaData(ctx context.Context, req *HandleDownlinkMetaDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleDownlinkMetaData not implemented")
}
func (*UnimplementedNetworkControllerServiceServer) HandleUplinkMACCommand(ctx context.Context, req *HandleUplinkMACCommandRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUplinkMACCommand not implemented")
}
func (*UnimplementedNetworkControllerServiceServer) HandleRejectedUplinkFrameSet(ctx context.Context, req *HandleRejectedUplinkFrameSetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleRejectedUplinkFrameSet not implemented")
}

func RegisterNetworkControllerServiceServer(s *grpc.Server, srv NetworkControllerServiceServer) {
	s.RegisterService(&_NetworkControllerService_serviceDesc, srv)
}

func _NetworkControllerService_HandleUplinkMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleUplinkMetaDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkControllerServiceServer).HandleUplinkMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nc.NetworkControllerService/HandleUplinkMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkControllerServiceServer).HandleUplinkMetaData(ctx, req.(*HandleUplinkMetaDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkControllerService_HandleDownlinkMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleDownlinkMetaDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkControllerServiceServer).HandleDownlinkMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nc.NetworkControllerService/HandleDownlinkMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkControllerServiceServer).HandleDownlinkMetaData(ctx, req.(*HandleDownlinkMetaDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkControllerService_HandleUplinkMACCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleUplinkMACCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkControllerServiceServer).HandleUplinkMACCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nc.NetworkControllerService/HandleUplinkMACCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkControllerServiceServer).HandleUplinkMACCommand(ctx, req.(*HandleUplinkMACCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkControllerService_HandleRejectedUplinkFrameSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleRejectedUplinkFrameSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkControllerServiceServer).HandleRejectedUplinkFrameSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nc.NetworkControllerService/HandleRejectedUplinkFrameSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkControllerServiceServer).HandleRejectedUplinkFrameSet(ctx, req.(*HandleRejectedUplinkFrameSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkControllerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nc.NetworkControllerService",
	HandlerType: (*NetworkControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleUplinkMetaData",
			Handler:    _NetworkControllerService_HandleUplinkMetaData_Handler,
		},
		{
			MethodName: "HandleDownlinkMetaData",
			Handler:    _NetworkControllerService_HandleDownlinkMetaData_Handler,
		},
		{
			MethodName: "HandleUplinkMACCommand",
			Handler:    _NetworkControllerService_HandleUplinkMACCommand_Handler,
		},
		{
			MethodName: "HandleRejectedUplinkFrameSet",
			Handler:    _NetworkControllerService_HandleRejectedUplinkFrameSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nc/nc.proto",
}
