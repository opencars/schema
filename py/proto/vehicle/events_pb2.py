# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/vehicle/events.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1aproto/vehicle/events.proto\x12\x07vehicle\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc0\x01\n\x14RegistrationSearched\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08token_id\x18\x07 \x01(\t\x12\x12\n\ntoken_name\x18\x08 \x01(\t\x12\x0c\n\x04\x63ode\x18\x02 \x01(\t\x12\x0e\n\x06number\x18\x03 \x01(\t\x12\x0b\n\x03vin\x18\x04 \x01(\t\x12\x15\n\rresult_amount\x18\x05 \x01(\r\x12/\n\x0bsearched_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xbd\x01\n\x11OperationSearched\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08token_id\x18\x07 \x01(\t\x12\x12\n\ntoken_name\x18\x08 \x01(\t\x12\x0c\n\x04\x63ode\x18\x02 \x01(\t\x12\x0e\n\x06number\x18\x03 \x01(\t\x12\x0b\n\x03vin\x18\x04 \x01(\t\x12\x15\n\rresult_amount\x18\x05 \x01(\r\x12/\n\x0bsearched_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x81\x01\n\nVINDecoded\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08token_id\x18\x02 \x01(\t\x12\x12\n\ntoken_name\x18\x03 \x01(\t\x12\x0b\n\x03vin\x18\x04 \x01(\t\x12/\n\x0bsearched_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampB$Z\"github.com/opencars/schema/vehicleb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.vehicle.events_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\"github.com/opencars/schema/vehicle'
  _REGISTRATIONSEARCHED._serialized_start=73
  _REGISTRATIONSEARCHED._serialized_end=265
  _OPERATIONSEARCHED._serialized_start=268
  _OPERATIONSEARCHED._serialized_end=457
  _VINDECODED._serialized_start=460
  _VINDECODED._serialized_end=589
# @@protoc_insertion_point(module_scope)
