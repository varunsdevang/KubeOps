# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ratelimit.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'ratelimit.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0fratelimit.proto\x12\x0cratelimitapi\"!\n\x0eServiceMessage\x12\x0f\n\x07service\x18\x01 \x01(\t\"2\n\x12ServiceRateMessage\x12\x0f\n\x07service\x18\x01 \x01(\t\x12\x0b\n\x03rpm\x18\x02 \x01(\x05\"/\n\x10RateLimitRequest\x12\n\n\x02ip\x18\x01 \x01(\t\x12\x0f\n\x07service\x18\x02 \x01(\t\"&\n\x11RateLimitResponse\x12\x11\n\tisAllowed\x18\x01 \x01(\x08\"\x07\n\x05\x45mpty2\x9a\x01\n\x0bRateService\x12@\n\x07SetRate\x12 .ratelimitapi.ServiceRateMessage\x1a\x13.ratelimitapi.Empty\x12I\n\x07GetRate\x12\x1c.ratelimitapi.ServiceMessage\x1a .ratelimitapi.ServiceRateMessage2g\n\x10RateLimitService\x12S\n\x10IsRequestAllowed\x12\x1e.ratelimitapi.RateLimitRequest\x1a\x1f.ratelimitapi.RateLimitResponseB\x0eZ\x0c./proto/grpcb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ratelimit_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\014./proto/grpc'
  _globals['_SERVICEMESSAGE']._serialized_start=33
  _globals['_SERVICEMESSAGE']._serialized_end=66
  _globals['_SERVICERATEMESSAGE']._serialized_start=68
  _globals['_SERVICERATEMESSAGE']._serialized_end=118
  _globals['_RATELIMITREQUEST']._serialized_start=120
  _globals['_RATELIMITREQUEST']._serialized_end=167
  _globals['_RATELIMITRESPONSE']._serialized_start=169
  _globals['_RATELIMITRESPONSE']._serialized_end=207
  _globals['_EMPTY']._serialized_start=209
  _globals['_EMPTY']._serialized_end=216
  _globals['_RATESERVICE']._serialized_start=219
  _globals['_RATESERVICE']._serialized_end=373
  _globals['_RATELIMITSERVICE']._serialized_start=375
  _globals['_RATELIMITSERVICE']._serialized_end=478
# @@protoc_insertion_point(module_scope)
