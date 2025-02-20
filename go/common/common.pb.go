// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.13.0
// source: common/common.proto

package common

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

type Modulation int32

const (
	// LoRa
	Modulation_LORA Modulation = 0
	// FSK
	Modulation_FSK Modulation = 1
	// LR-FHSS
	Modulation_LR_FHSS Modulation = 2
)

// Enum value maps for Modulation.
var (
	Modulation_name = map[int32]string{
		0: "LORA",
		1: "FSK",
		2: "LR_FHSS",
	}
	Modulation_value = map[string]int32{
		"LORA":    0,
		"FSK":     1,
		"LR_FHSS": 2,
	}
)

func (x Modulation) Enum() *Modulation {
	p := new(Modulation)
	*p = x
	return p
}

func (x Modulation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Modulation) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[0].Descriptor()
}

func (Modulation) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[0]
}

func (x Modulation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Modulation.Descriptor instead.
func (Modulation) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

type Region int32

const (
	// EU868
	Region_EU868 Region = 0
	// US915
	Region_US915 Region = 2
	// CN779
	Region_CN779 Region = 3
	// EU433
	Region_EU433 Region = 4
	// AU915
	Region_AU915 Region = 5
	// CN470
	Region_CN470 Region = 6
	// AS923
	Region_AS923 Region = 7
	// AS923 with -1.80 MHz frequency offset
	Region_AS923_2 Region = 12
	// AS923 with -6.60 MHz frequency offset
	Region_AS923_3 Region = 13
	// AS923 with -5.90 MHz frequency offset
	Region_AS923_4 Region = 14
	// KR920
	Region_KR920 Region = 8
	// IN865
	Region_IN865 Region = 9
	// RU864
	Region_RU864 Region = 10
	// ISM2400 (LoRaWAN 2.4 GHz)
	Region_ISM2400 Region = 11
)

// Enum value maps for Region.
var (
	Region_name = map[int32]string{
		0:  "EU868",
		2:  "US915",
		3:  "CN779",
		4:  "EU433",
		5:  "AU915",
		6:  "CN470",
		7:  "AS923",
		12: "AS923_2",
		13: "AS923_3",
		14: "AS923_4",
		8:  "KR920",
		9:  "IN865",
		10: "RU864",
		11: "ISM2400",
	}
	Region_value = map[string]int32{
		"EU868":   0,
		"US915":   2,
		"CN779":   3,
		"EU433":   4,
		"AU915":   5,
		"CN470":   6,
		"AS923":   7,
		"AS923_2": 12,
		"AS923_3": 13,
		"AS923_4": 14,
		"KR920":   8,
		"IN865":   9,
		"RU864":   10,
		"ISM2400": 11,
	}
)

func (x Region) Enum() *Region {
	p := new(Region)
	*p = x
	return p
}

func (x Region) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Region) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[1].Descriptor()
}

func (Region) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[1]
}

func (x Region) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Region.Descriptor instead.
func (Region) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

type MType int32

const (
	// JoinRequest.
	MType_JoinRequest MType = 0
	// JoinAccept.
	MType_JoinAccept MType = 1
	// UnconfirmedDataUp.
	MType_UnconfirmedDataUp MType = 2
	// UnconfirmedDataDown.
	MType_UnconfirmedDataDown MType = 3
	// ConfirmedDataUp.
	MType_ConfirmedDataUp MType = 4
	// ConfirmedDataDown.
	MType_ConfirmedDataDown MType = 5
	// RejoinRequest.
	MType_RejoinRequest MType = 6
	// Proprietary.
	MType_Proprietary MType = 7
)

// Enum value maps for MType.
var (
	MType_name = map[int32]string{
		0: "JoinRequest",
		1: "JoinAccept",
		2: "UnconfirmedDataUp",
		3: "UnconfirmedDataDown",
		4: "ConfirmedDataUp",
		5: "ConfirmedDataDown",
		6: "RejoinRequest",
		7: "Proprietary",
	}
	MType_value = map[string]int32{
		"JoinRequest":         0,
		"JoinAccept":          1,
		"UnconfirmedDataUp":   2,
		"UnconfirmedDataDown": 3,
		"ConfirmedDataUp":     4,
		"ConfirmedDataDown":   5,
		"RejoinRequest":       6,
		"Proprietary":         7,
	}
)

func (x MType) Enum() *MType {
	p := new(MType)
	*p = x
	return p
}

func (x MType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[2].Descriptor()
}

func (MType) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[2]
}

func (x MType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MType.Descriptor instead.
func (MType) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

type LocationSource int32

const (
	// Unknown.
	LocationSource_UNKNOWN LocationSource = 0
	// GPS.
	LocationSource_GPS LocationSource = 1
	// Manually configured.
	LocationSource_CONFIG LocationSource = 2
	// Geo resolver (TDOA).
	LocationSource_GEO_RESOLVER_TDOA LocationSource = 3
	// Geo resolver (RSSI).
	LocationSource_GEO_RESOLVER_RSSI LocationSource = 4
	// Geo resolver (GNSS).
	LocationSource_GEO_RESOLVER_GNSS LocationSource = 5
	// Geo resolver (WIFI).
	LocationSource_GEO_RESOLVER_WIFI LocationSource = 6
)

// Enum value maps for LocationSource.
var (
	LocationSource_name = map[int32]string{
		0: "UNKNOWN",
		1: "GPS",
		2: "CONFIG",
		3: "GEO_RESOLVER_TDOA",
		4: "GEO_RESOLVER_RSSI",
		5: "GEO_RESOLVER_GNSS",
		6: "GEO_RESOLVER_WIFI",
	}
	LocationSource_value = map[string]int32{
		"UNKNOWN":           0,
		"GPS":               1,
		"CONFIG":            2,
		"GEO_RESOLVER_TDOA": 3,
		"GEO_RESOLVER_RSSI": 4,
		"GEO_RESOLVER_GNSS": 5,
		"GEO_RESOLVER_WIFI": 6,
	}
)

func (x LocationSource) Enum() *LocationSource {
	p := new(LocationSource)
	*p = x
	return p
}

func (x LocationSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocationSource) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[3].Descriptor()
}

func (LocationSource) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[3]
}

func (x LocationSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocationSource.Descriptor instead.
func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

type KeyEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// KEK label.
	KekLabel string `protobuf:"bytes,1,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	// AES key (when the kek_label is set, this key is encrypted using a key
	// known to the join-server and application-server.
	// For more information please refer to the LoRaWAN Backend Interface
	// 'Key Transport Security' section.
	AesKey []byte `protobuf:"bytes,2,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
}

func (x *KeyEnvelope) Reset() {
	*x = KeyEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyEnvelope) ProtoMessage() {}

func (x *KeyEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyEnvelope.ProtoReflect.Descriptor instead.
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *KeyEnvelope) GetKekLabel() string {
	if x != nil {
		return x.KekLabel
	}
	return ""
}

func (x *KeyEnvelope) GetAesKey() []byte {
	if x != nil {
		return x.AesKey
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Location source.
	Source LocationSource `protobuf:"varint,4,opt,name=source,proto3,enum=common.LocationSource" json:"source,omitempty"`
	// Accuracy (in meters).
	Accuracy uint32 `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Location) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *Location) GetSource() LocationSource {
	if x != nil {
		return x.Source
	}
	return LocationSource_UNKNOWN
}

func (x *Location) GetAccuracy() uint32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x43, 0x0a,
	0x0b, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6b, 0x65, 0x6b, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x6b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x65, 0x73,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x65, 0x73, 0x4b,
	0x65, 0x79, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6c, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63,
	0x79, 0x2a, 0x2c, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x52, 0x41, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x53, 0x4b,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x52, 0x5f, 0x46, 0x48, 0x53, 0x53, 0x10, 0x02, 0x2a,
	0xaa, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x55,
	0x38, 0x36, 0x38, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x53, 0x39, 0x31, 0x35, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x43, 0x4e, 0x37, 0x37, 0x39, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x55, 0x34, 0x33, 0x33, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x39, 0x31, 0x35, 0x10,
	0x05, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4e, 0x34, 0x37, 0x30, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x53, 0x39, 0x32, 0x33, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x53, 0x39, 0x32, 0x33,
	0x5f, 0x32, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x33, 0x10,
	0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x34, 0x10, 0x0e, 0x12, 0x09,
	0x0a, 0x05, 0x4b, 0x52, 0x39, 0x32, 0x30, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x38,
	0x36, 0x35, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x55, 0x38, 0x36, 0x34, 0x10, 0x0a, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x53, 0x4d, 0x32, 0x34, 0x30, 0x30, 0x10, 0x0b, 0x2a, 0xa8, 0x01, 0x0a,
	0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x6f, 0x69, 0x6e, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x6e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x44, 0x6f, 0x77, 0x6e, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x44, 0x6f, 0x77,
	0x6e, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x6a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x70, 0x72, 0x69,
	0x65, 0x74, 0x61, 0x72, 0x79, 0x10, 0x07, 0x2a, 0x8e, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x50, 0x53, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x47, 0x45, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x54, 0x44, 0x4f,
	0x41, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x45, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c,
	0x56, 0x45, 0x52, 0x5f, 0x52, 0x53, 0x53, 0x49, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x45,
	0x4f, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x47, 0x4e, 0x53, 0x53, 0x10,
	0x05, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x45, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45,
	0x52, 0x5f, 0x57, 0x49, 0x46, 0x49, 0x10, 0x06, 0x42, 0x5d, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x63,
	0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x63, 0x68, 0x69, 0x72,
	0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_common_proto_goTypes = []interface{}{
	(Modulation)(0),     // 0: common.Modulation
	(Region)(0),         // 1: common.Region
	(MType)(0),          // 2: common.MType
	(LocationSource)(0), // 3: common.LocationSource
	(*KeyEnvelope)(nil), // 4: common.KeyEnvelope
	(*Location)(nil),    // 5: common.Location
}
var file_common_common_proto_depIdxs = []int32{
	3, // 0: common.Location.source:type_name -> common.LocationSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyEnvelope); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		EnumInfos:         file_common_common_proto_enumTypes,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
