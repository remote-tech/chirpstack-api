# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/as_pb.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from chirpstack_api.common import common_pb2 as chirpstack__api_dot_common_dot_common__pb2
from chirpstack_api.gw import gw_pb2 as chirpstack__api_dot_gw_dot_gw__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='chirpstack-api/as_pb/as_pb.proto',
  package='as',
  syntax='proto3',
  serialized_options=b'Z*github.com/brocaar/chirpstack-api/go/v3/as',
  serialized_pb=b'\n chirpstack-api/as_pb/as_pb.proto\x12\x02\x61s\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\"chirpstack-api/common/common.proto\x1a\x1a\x63hirpstack-api/gw/gw.proto\"S\n\x17\x44\x65viceActivationContext\x12\x10\n\x08\x64\x65v_addr\x18\x01 \x01(\x0c\x12&\n\tapp_s_key\x18\x02 \x01(\x0b\x32\x13.common.KeyEnvelope\"\xa2\x02\n\x17HandleUplinkDataRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\x10\n\x08join_eui\x18\x02 \x01(\x0c\x12\r\n\x05\x66_cnt\x18\x03 \x01(\r\x12\x0e\n\x06\x66_port\x18\x04 \x01(\r\x12\x0b\n\x03\x61\x64r\x18\x05 \x01(\x08\x12\n\n\x02\x64r\x18\x06 \x01(\r\x12!\n\x07tx_info\x18\x07 \x01(\x0b\x32\x10.gw.UplinkTXInfo\x12!\n\x07rx_info\x18\x08 \x03(\x0b\x32\x10.gw.UplinkRXInfo\x12\x0c\n\x04\x64\x61ta\x18\t \x01(\x0c\x12>\n\x19\x64\x65vice_activation_context\x18\n \x01(\x0b\x32\x1b.as.DeviceActivationContext\x12\x18\n\x10\x63onfirmed_uplink\x18\x0b \x01(\x08\"\x88\x01\n\x1eHandleProprietaryUplinkRequest\x12\x13\n\x0bmac_payload\x18\x01 \x01(\x0c\x12\x0b\n\x03mic\x18\x02 \x01(\x0c\x12!\n\x07tx_info\x18\x03 \x01(\x0b\x32\x10.gw.UplinkTXInfo\x12!\n\x07rx_info\x18\x04 \x03(\x0b\x32\x10.gw.UplinkRXInfo\"`\n\x12HandleErrorRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\x1b\n\x04type\x18\x03 \x01(\x0e\x32\r.as.ErrorType\x12\r\n\x05\x65rror\x18\x04 \x01(\t\x12\r\n\x05\x66_cnt\x18\x05 \x01(\r\"P\n\x18HandleDownlinkACKRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\r\n\x05\x66_cnt\x18\x02 \x01(\r\x12\x14\n\x0c\x61\x63knowledged\x18\x03 \x01(\x08\"\xa3\x01\n\x16SetDeviceStatusRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\x0f\n\x07\x62\x61ttery\x18\x02 \x01(\r\x12\x0e\n\x06margin\x18\x03 \x01(\x05\x12\x1d\n\x15\x65xternal_power_source\x18\x04 \x01(\x08\x12!\n\x19\x62\x61ttery_level_unavailable\x18\x05 \x01(\x08\x12\x15\n\rbattery_level\x18\x06 \x01(\x02\"c\n\x18SetDeviceLocationRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\"\n\x08location\x18\x02 \x01(\x0b\x32\x10.common.Location\x12\x12\n\nuplink_ids\x18\x03 \x03(\x0c\"\xf5\x02\n\x19HandleGatewayStatsRequest\x12\x12\n\ngateway_id\x18\x01 \x01(\x0c\x12\x10\n\x08stats_id\x18\x02 \x01(\x0c\x12(\n\x04time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\"\n\x08location\x18\x04 \x01(\x0b\x32\x10.common.Location\x12\x1b\n\x13rx_packets_received\x18\x05 \x01(\r\x12\x1e\n\x16rx_packets_received_ok\x18\x06 \x01(\r\x12\x1b\n\x13tx_packets_received\x18\x07 \x01(\r\x12\x1a\n\x12tx_packets_emitted\x18\x08 \x01(\r\x12=\n\x08metadata\x18\t \x03(\x0b\x32+.as.HandleGatewayStatsRequest.MetadataEntry\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"4\n\x12HandleTxAckRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\r\n\x05\x66_cnt\x18\x02 \x01(\r*\x1c\n\x08RXWindow\x12\x07\n\x03RX1\x10\x00\x12\x07\n\x03RX2\x10\x01*\xbb\x01\n\tErrorType\x12\x0b\n\x07GENERIC\x10\x00\x12\x08\n\x04OTAA\x10\x01\x12\x16\n\x12\x44\x41TA_UP_FCNT_RESET\x10\x02\x12\x0f\n\x0b\x44\x41TA_UP_MIC\x10\x03\x12\x1a\n\x16\x44\x45VICE_QUEUE_ITEM_SIZE\x10\x04\x12\x1a\n\x16\x44\x45VICE_QUEUE_ITEM_FCNT\x10\x05\x12\x1f\n\x1b\x44\x41TA_UP_FCNT_RETRANSMISSION\x10\x06\x12\x15\n\x11\x44\x41TA_DOWN_GATEWAY\x10\x07\x32\xf2\x04\n\x18\x41pplicationServerService\x12I\n\x10HandleUplinkData\x12\x1b.as.HandleUplinkDataRequest\x1a\x16.google.protobuf.Empty\"\x00\x12W\n\x17HandleProprietaryUplink\x12\".as.HandleProprietaryUplinkRequest\x1a\x16.google.protobuf.Empty\"\x00\x12?\n\x0bHandleError\x12\x16.as.HandleErrorRequest\x1a\x16.google.protobuf.Empty\"\x00\x12K\n\x11HandleDownlinkACK\x12\x1c.as.HandleDownlinkACKRequest\x1a\x16.google.protobuf.Empty\"\x00\x12M\n\x12HandleGatewayStats\x12\x1d.as.HandleGatewayStatsRequest\x1a\x16.google.protobuf.Empty\"\x00\x12?\n\x0bHandleTxAck\x12\x16.as.HandleTxAckRequest\x1a\x16.google.protobuf.Empty\"\x00\x12G\n\x0fSetDeviceStatus\x12\x1a.as.SetDeviceStatusRequest\x1a\x16.google.protobuf.Empty\"\x00\x12K\n\x11SetDeviceLocation\x12\x1c.as.SetDeviceLocationRequest\x1a\x16.google.protobuf.Empty\"\x00\x42,Z*github.com/brocaar/chirpstack-api/go/v3/asb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,chirpstack__api_dot_common_dot_common__pb2.DESCRIPTOR,chirpstack__api_dot_gw_dot_gw__pb2.DESCRIPTOR,])

_RXWINDOW = _descriptor.EnumDescriptor(
  name='RXWindow',
  full_name='as.RXWindow',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='RX1', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RX2', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1560,
  serialized_end=1588,
)
_sym_db.RegisterEnumDescriptor(_RXWINDOW)

RXWindow = enum_type_wrapper.EnumTypeWrapper(_RXWINDOW)
_ERRORTYPE = _descriptor.EnumDescriptor(
  name='ErrorType',
  full_name='as.ErrorType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='GENERIC', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OTAA', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DATA_UP_FCNT_RESET', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DATA_UP_MIC', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEVICE_QUEUE_ITEM_SIZE', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEVICE_QUEUE_ITEM_FCNT', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DATA_UP_FCNT_RETRANSMISSION', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DATA_DOWN_GATEWAY', index=7, number=7,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1591,
  serialized_end=1778,
)
_sym_db.RegisterEnumDescriptor(_ERRORTYPE)

ErrorType = enum_type_wrapper.EnumTypeWrapper(_ERRORTYPE)
RX1 = 0
RX2 = 1
GENERIC = 0
OTAA = 1
DATA_UP_FCNT_RESET = 2
DATA_UP_MIC = 3
DEVICE_QUEUE_ITEM_SIZE = 4
DEVICE_QUEUE_ITEM_FCNT = 5
DATA_UP_FCNT_RETRANSMISSION = 6
DATA_DOWN_GATEWAY = 7



_DEVICEACTIVATIONCONTEXT = _descriptor.Descriptor(
  name='DeviceActivationContext',
  full_name='as.DeviceActivationContext',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_addr', full_name='as.DeviceActivationContext.dev_addr', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='app_s_key', full_name='as.DeviceActivationContext.app_s_key', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=166,
  serialized_end=249,
)


_HANDLEUPLINKDATAREQUEST = _descriptor.Descriptor(
  name='HandleUplinkDataRequest',
  full_name='as.HandleUplinkDataRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='as.HandleUplinkDataRequest.dev_eui', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='join_eui', full_name='as.HandleUplinkDataRequest.join_eui', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='f_cnt', full_name='as.HandleUplinkDataRequest.f_cnt', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='f_port', full_name='as.HandleUplinkDataRequest.f_port', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='adr', full_name='as.HandleUplinkDataRequest.adr', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dr', full_name='as.HandleUplinkDataRequest.dr', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tx_info', full_name='as.HandleUplinkDataRequest.tx_info', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rx_info', full_name='as.HandleUplinkDataRequest.rx_info', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='as.HandleUplinkDataRequest.data', index=8,
      number=9, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='device_activation_context', full_name='as.HandleUplinkDataRequest.device_activation_context', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='confirmed_uplink', full_name='as.HandleUplinkDataRequest.confirmed_uplink', index=10,
      number=11, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=252,
  serialized_end=542,
)


_HANDLEPROPRIETARYUPLINKREQUEST = _descriptor.Descriptor(
  name='HandleProprietaryUplinkRequest',
  full_name='as.HandleProprietaryUplinkRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mac_payload', full_name='as.HandleProprietaryUplinkRequest.mac_payload', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mic', full_name='as.HandleProprietaryUplinkRequest.mic', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tx_info', full_name='as.HandleProprietaryUplinkRequest.tx_info', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rx_info', full_name='as.HandleProprietaryUplinkRequest.rx_info', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=545,
  serialized_end=681,
)


_HANDLEERRORREQUEST = _descriptor.Descriptor(
  name='HandleErrorRequest',
  full_name='as.HandleErrorRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='as.HandleErrorRequest.dev_eui', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='as.HandleErrorRequest.type', index=1,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='as.HandleErrorRequest.error', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='f_cnt', full_name='as.HandleErrorRequest.f_cnt', index=3,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=683,
  serialized_end=779,
)


_HANDLEDOWNLINKACKREQUEST = _descriptor.Descriptor(
  name='HandleDownlinkACKRequest',
  full_name='as.HandleDownlinkACKRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='as.HandleDownlinkACKRequest.dev_eui', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='f_cnt', full_name='as.HandleDownlinkACKRequest.f_cnt', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='acknowledged', full_name='as.HandleDownlinkACKRequest.acknowledged', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=781,
  serialized_end=861,
)


_SETDEVICESTATUSREQUEST = _descriptor.Descriptor(
  name='SetDeviceStatusRequest',
  full_name='as.SetDeviceStatusRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='as.SetDeviceStatusRequest.dev_eui', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='battery', full_name='as.SetDeviceStatusRequest.battery', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='margin', full_name='as.SetDeviceStatusRequest.margin', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='external_power_source', full_name='as.SetDeviceStatusRequest.external_power_source', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='battery_level_unavailable', full_name='as.SetDeviceStatusRequest.battery_level_unavailable', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='battery_level', full_name='as.SetDeviceStatusRequest.battery_level', index=5,
      number=6, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=864,
  serialized_end=1027,
)


_SETDEVICELOCATIONREQUEST = _descriptor.Descriptor(
  name='SetDeviceLocationRequest',
  full_name='as.SetDeviceLocationRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='as.SetDeviceLocationRequest.dev_eui', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location', full_name='as.SetDeviceLocationRequest.location', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uplink_ids', full_name='as.SetDeviceLocationRequest.uplink_ids', index=2,
      number=3, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1029,
  serialized_end=1128,
)


_HANDLEGATEWAYSTATSREQUEST_METADATAENTRY = _descriptor.Descriptor(
  name='MetadataEntry',
  full_name='as.HandleGatewayStatsRequest.MetadataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='as.HandleGatewayStatsRequest.MetadataEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='as.HandleGatewayStatsRequest.MetadataEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1457,
  serialized_end=1504,
)

_HANDLEGATEWAYSTATSREQUEST = _descriptor.Descriptor(
  name='HandleGatewayStatsRequest',
  full_name='as.HandleGatewayStatsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='gateway_id', full_name='as.HandleGatewayStatsRequest.gateway_id', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='stats_id', full_name='as.HandleGatewayStatsRequest.stats_id', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='time', full_name='as.HandleGatewayStatsRequest.time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location', full_name='as.HandleGatewayStatsRequest.location', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rx_packets_received', full_name='as.HandleGatewayStatsRequest.rx_packets_received', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rx_packets_received_ok', full_name='as.HandleGatewayStatsRequest.rx_packets_received_ok', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tx_packets_received', full_name='as.HandleGatewayStatsRequest.tx_packets_received', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tx_packets_emitted', full_name='as.HandleGatewayStatsRequest.tx_packets_emitted', index=7,
      number=8, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='as.HandleGatewayStatsRequest.metadata', index=8,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_HANDLEGATEWAYSTATSREQUEST_METADATAENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1131,
  serialized_end=1504,
)


_HANDLETXACKREQUEST = _descriptor.Descriptor(
  name='HandleTxAckRequest',
  full_name='as.HandleTxAckRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='as.HandleTxAckRequest.dev_eui', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='f_cnt', full_name='as.HandleTxAckRequest.f_cnt', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1506,
  serialized_end=1558,
)

_DEVICEACTIVATIONCONTEXT.fields_by_name['app_s_key'].message_type = chirpstack__api_dot_common_dot_common__pb2._KEYENVELOPE
_HANDLEUPLINKDATAREQUEST.fields_by_name['tx_info'].message_type = chirpstack__api_dot_gw_dot_gw__pb2._UPLINKTXINFO
_HANDLEUPLINKDATAREQUEST.fields_by_name['rx_info'].message_type = chirpstack__api_dot_gw_dot_gw__pb2._UPLINKRXINFO
_HANDLEUPLINKDATAREQUEST.fields_by_name['device_activation_context'].message_type = _DEVICEACTIVATIONCONTEXT
_HANDLEPROPRIETARYUPLINKREQUEST.fields_by_name['tx_info'].message_type = chirpstack__api_dot_gw_dot_gw__pb2._UPLINKTXINFO
_HANDLEPROPRIETARYUPLINKREQUEST.fields_by_name['rx_info'].message_type = chirpstack__api_dot_gw_dot_gw__pb2._UPLINKRXINFO
_HANDLEERRORREQUEST.fields_by_name['type'].enum_type = _ERRORTYPE
_SETDEVICELOCATIONREQUEST.fields_by_name['location'].message_type = chirpstack__api_dot_common_dot_common__pb2._LOCATION
_HANDLEGATEWAYSTATSREQUEST_METADATAENTRY.containing_type = _HANDLEGATEWAYSTATSREQUEST
_HANDLEGATEWAYSTATSREQUEST.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_HANDLEGATEWAYSTATSREQUEST.fields_by_name['location'].message_type = chirpstack__api_dot_common_dot_common__pb2._LOCATION
_HANDLEGATEWAYSTATSREQUEST.fields_by_name['metadata'].message_type = _HANDLEGATEWAYSTATSREQUEST_METADATAENTRY
DESCRIPTOR.message_types_by_name['DeviceActivationContext'] = _DEVICEACTIVATIONCONTEXT
DESCRIPTOR.message_types_by_name['HandleUplinkDataRequest'] = _HANDLEUPLINKDATAREQUEST
DESCRIPTOR.message_types_by_name['HandleProprietaryUplinkRequest'] = _HANDLEPROPRIETARYUPLINKREQUEST
DESCRIPTOR.message_types_by_name['HandleErrorRequest'] = _HANDLEERRORREQUEST
DESCRIPTOR.message_types_by_name['HandleDownlinkACKRequest'] = _HANDLEDOWNLINKACKREQUEST
DESCRIPTOR.message_types_by_name['SetDeviceStatusRequest'] = _SETDEVICESTATUSREQUEST
DESCRIPTOR.message_types_by_name['SetDeviceLocationRequest'] = _SETDEVICELOCATIONREQUEST
DESCRIPTOR.message_types_by_name['HandleGatewayStatsRequest'] = _HANDLEGATEWAYSTATSREQUEST
DESCRIPTOR.message_types_by_name['HandleTxAckRequest'] = _HANDLETXACKREQUEST
DESCRIPTOR.enum_types_by_name['RXWindow'] = _RXWINDOW
DESCRIPTOR.enum_types_by_name['ErrorType'] = _ERRORTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DeviceActivationContext = _reflection.GeneratedProtocolMessageType('DeviceActivationContext', (_message.Message,), {
  'DESCRIPTOR' : _DEVICEACTIVATIONCONTEXT,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.DeviceActivationContext)
  })
_sym_db.RegisterMessage(DeviceActivationContext)

HandleUplinkDataRequest = _reflection.GeneratedProtocolMessageType('HandleUplinkDataRequest', (_message.Message,), {
  'DESCRIPTOR' : _HANDLEUPLINKDATAREQUEST,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.HandleUplinkDataRequest)
  })
_sym_db.RegisterMessage(HandleUplinkDataRequest)

HandleProprietaryUplinkRequest = _reflection.GeneratedProtocolMessageType('HandleProprietaryUplinkRequest', (_message.Message,), {
  'DESCRIPTOR' : _HANDLEPROPRIETARYUPLINKREQUEST,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.HandleProprietaryUplinkRequest)
  })
_sym_db.RegisterMessage(HandleProprietaryUplinkRequest)

HandleErrorRequest = _reflection.GeneratedProtocolMessageType('HandleErrorRequest', (_message.Message,), {
  'DESCRIPTOR' : _HANDLEERRORREQUEST,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.HandleErrorRequest)
  })
_sym_db.RegisterMessage(HandleErrorRequest)

HandleDownlinkACKRequest = _reflection.GeneratedProtocolMessageType('HandleDownlinkACKRequest', (_message.Message,), {
  'DESCRIPTOR' : _HANDLEDOWNLINKACKREQUEST,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.HandleDownlinkACKRequest)
  })
_sym_db.RegisterMessage(HandleDownlinkACKRequest)

SetDeviceStatusRequest = _reflection.GeneratedProtocolMessageType('SetDeviceStatusRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETDEVICESTATUSREQUEST,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.SetDeviceStatusRequest)
  })
_sym_db.RegisterMessage(SetDeviceStatusRequest)

SetDeviceLocationRequest = _reflection.GeneratedProtocolMessageType('SetDeviceLocationRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETDEVICELOCATIONREQUEST,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.SetDeviceLocationRequest)
  })
_sym_db.RegisterMessage(SetDeviceLocationRequest)

HandleGatewayStatsRequest = _reflection.GeneratedProtocolMessageType('HandleGatewayStatsRequest', (_message.Message,), {

  'MetadataEntry' : _reflection.GeneratedProtocolMessageType('MetadataEntry', (_message.Message,), {
    'DESCRIPTOR' : _HANDLEGATEWAYSTATSREQUEST_METADATAENTRY,
    '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
    # @@protoc_insertion_point(class_scope:as.HandleGatewayStatsRequest.MetadataEntry)
    })
  ,
  'DESCRIPTOR' : _HANDLEGATEWAYSTATSREQUEST,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.HandleGatewayStatsRequest)
  })
_sym_db.RegisterMessage(HandleGatewayStatsRequest)
_sym_db.RegisterMessage(HandleGatewayStatsRequest.MetadataEntry)

HandleTxAckRequest = _reflection.GeneratedProtocolMessageType('HandleTxAckRequest', (_message.Message,), {
  'DESCRIPTOR' : _HANDLETXACKREQUEST,
  '__module__' : 'chirpstack_api.as_pb.as_pb_pb2'
  # @@protoc_insertion_point(class_scope:as.HandleTxAckRequest)
  })
_sym_db.RegisterMessage(HandleTxAckRequest)


DESCRIPTOR._options = None
_HANDLEGATEWAYSTATSREQUEST_METADATAENTRY._options = None

_APPLICATIONSERVERSERVICE = _descriptor.ServiceDescriptor(
  name='ApplicationServerService',
  full_name='as.ApplicationServerService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1781,
  serialized_end=2407,
  methods=[
  _descriptor.MethodDescriptor(
    name='HandleUplinkData',
    full_name='as.ApplicationServerService.HandleUplinkData',
    index=0,
    containing_service=None,
    input_type=_HANDLEUPLINKDATAREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='HandleProprietaryUplink',
    full_name='as.ApplicationServerService.HandleProprietaryUplink',
    index=1,
    containing_service=None,
    input_type=_HANDLEPROPRIETARYUPLINKREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='HandleError',
    full_name='as.ApplicationServerService.HandleError',
    index=2,
    containing_service=None,
    input_type=_HANDLEERRORREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='HandleDownlinkACK',
    full_name='as.ApplicationServerService.HandleDownlinkACK',
    index=3,
    containing_service=None,
    input_type=_HANDLEDOWNLINKACKREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='HandleGatewayStats',
    full_name='as.ApplicationServerService.HandleGatewayStats',
    index=4,
    containing_service=None,
    input_type=_HANDLEGATEWAYSTATSREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='HandleTxAck',
    full_name='as.ApplicationServerService.HandleTxAck',
    index=5,
    containing_service=None,
    input_type=_HANDLETXACKREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetDeviceStatus',
    full_name='as.ApplicationServerService.SetDeviceStatus',
    index=6,
    containing_service=None,
    input_type=_SETDEVICESTATUSREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetDeviceLocation',
    full_name='as.ApplicationServerService.SetDeviceLocation',
    index=7,
    containing_service=None,
    input_type=_SETDEVICELOCATIONREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_APPLICATIONSERVERSERVICE)

DESCRIPTOR.services_by_name['ApplicationServerService'] = _APPLICATIONSERVERSERVICE

# @@protoc_insertion_point(module_scope)
