# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/vehicle/operation.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.common import date_pb2 as proto_dot_common_dot_date__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dproto/vehicle/operation.proto\x12\x07vehicle\x1a\x17proto/common/date.proto\"(\n\nDepartment\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xab\x01\n\x05Owner\x12%\n\x06\x65ntity\x18\x01 \x01(\x0e\x32\x15.vehicle.Owner.Entity\x12.\n\x0cregistration\x18\x02 \x01(\x0b\x32\x18.vehicle.Owner.Territory\x1a\x19\n\tTerritory\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\"0\n\x06\x45ntity\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0e\n\nINDIVIDUAL\x10\x01\x12\t\n\x05LEGAL\x10\x02\"\xbe\x02\n\tOperation\x12\x0e\n\x06number\x18\x01 \x01(\t\x12\x0b\n\x03vin\x18\x10 \x01(\t\x12\r\n\x05\x62rand\x18\x02 \x01(\t\x12\r\n\x05model\x18\x03 \x01(\t\x12\x0c\n\x04year\x18\x04 \x01(\x05\x12\x10\n\x08\x63\x61pacity\x18\x05 \x01(\x05\x12\r\n\x05\x63olor\x18\x06 \x01(\t\x12\x0c\n\x04\x66uel\x18\x07 \x01(\t\x12\x0c\n\x04kind\x18\x08 \x01(\t\x12\x0c\n\x04\x62ody\x18\t \x01(\t\x12\x0f\n\x07purpose\x18\n \x01(\t\x12\x12\n\nown_weight\x18\x0b \x01(\x05\x12\x14\n\x0ctotal_weight\x18\x0c \x01(\x05\x12\x1a\n\x04\x64\x61te\x18\r \x01(\x0b\x32\x0c.common.Date\x12\'\n\ndepartment\x18\x0e \x01(\x0b\x32\x13.vehicle.Department\x12\x1d\n\x05owner\x18\x0f \x01(\x0b\x32\x0e.vehicle.OwnerB$Z\"github.com/opencars/schema/vehicleb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.vehicle.operation_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\"github.com/opencars/schema/vehicle'
  _DEPARTMENT._serialized_start=67
  _DEPARTMENT._serialized_end=107
  _OWNER._serialized_start=110
  _OWNER._serialized_end=281
  _OWNER_TERRITORY._serialized_start=206
  _OWNER_TERRITORY._serialized_end=231
  _OWNER_ENTITY._serialized_start=233
  _OWNER_ENTITY._serialized_end=281
  _OPERATION._serialized_start=284
  _OPERATION._serialized_end=602
# @@protoc_insertion_point(module_scope)
