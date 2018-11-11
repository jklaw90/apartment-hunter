# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: apthunter.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='apthunter.proto',
  package='pb',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0f\x61pthunter.proto\x12\x02pb\"\xc1\x02\n\x15\x43reateOrUpdateRequest\x12\n\n\x02id\x18\n \x01(\t\x12\x0f\n\x07\x61\x64\x64ress\x18\x14 \x01(\t\x12\x0b\n\x03url\x18\x1e \x01(\t\x12\r\n\x05title\x18( \x01(\t\x12\r\n\x05price\x18\x32 \x01(\x02\x12\x10\n\x08\x62\x65\x64rooms\x18< \x01(\x04\x12\x11\n\tbathrooms\x18\x46 \x01(\x02\x12\x0c\n\x04sqft\x18P \x01(\x02\x12\x16\n\x0e\x61vailable_date\x18Z \x01(\t\x12\x0c\n\x04\x63\x61ts\x18\x64 \x01(\x08\x12\x0c\n\x04\x64ogs\x18n \x01(\x08\x12\x14\n\x0chousing_type\x18x \x01(\t\x12\x10\n\x07wd_type\x18\x82\x01 \x01(\t\x12\x15\n\x0cparking_type\x18\x8c\x01 \x01(\t\x12\x0f\n\x06images\x18\x96\x01 \x03(\t\x12\r\n\x04\x62ody\x18\xa0\x01 \x01(\t\x12\x0c\n\x03lng\x18\xaa\x01 \x01(\x02\x12\x0c\n\x03lat\x18\xb4\x01 \x01(\x02\"$\n\x16\x43reateOrUpdateResponse\x12\n\n\x02id\x18\n \x01(\t\"-\n\rStreamRequest\x12\n\n\x02id\x18\n \x01(\t\x12\x10\n\x08lastSeen\x18\x14 \x01(\t\"w\n\x0e\x41ggregateEvent\x12\n\n\x02id\x18\n \x01(\t\x12\x13\n\x0b\x61ggregateId\x18\x14 \x01(\t\x12\x11\n\ttimestamp\x18\x1e \x01(\x03\x12\x0f\n\x07version\x18( \x01(\r\x12\x11\n\teventType\x18\x32 \x01(\t\x12\r\n\x05\x65vent\x18< \x01(\x0c\"\x0e\n\x0c\x45ndedRequest\"\x0f\n\rEndedResponse2\x88\x01\n\x06Writer\x12I\n\x0e\x43reateOrUpdate\x12\x19.pb.CreateOrUpdateRequest\x1a\x1a.pb.CreateOrUpdateResponse\"\x00\x12\x33\n\x06Stream\x12\x11.pb.StreamRequest\x1a\x12.pb.AggregateEvent\"\x00\x30\x01\x62\x06proto3')
)




_CREATEORUPDATEREQUEST = _descriptor.Descriptor(
  name='CreateOrUpdateRequest',
  full_name='pb.CreateOrUpdateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.CreateOrUpdateRequest.id', index=0,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address', full_name='pb.CreateOrUpdateRequest.address', index=1,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='url', full_name='pb.CreateOrUpdateRequest.url', index=2,
      number=30, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title', full_name='pb.CreateOrUpdateRequest.title', index=3,
      number=40, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='pb.CreateOrUpdateRequest.price', index=4,
      number=50, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bedrooms', full_name='pb.CreateOrUpdateRequest.bedrooms', index=5,
      number=60, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bathrooms', full_name='pb.CreateOrUpdateRequest.bathrooms', index=6,
      number=70, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sqft', full_name='pb.CreateOrUpdateRequest.sqft', index=7,
      number=80, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='available_date', full_name='pb.CreateOrUpdateRequest.available_date', index=8,
      number=90, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cats', full_name='pb.CreateOrUpdateRequest.cats', index=9,
      number=100, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dogs', full_name='pb.CreateOrUpdateRequest.dogs', index=10,
      number=110, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='housing_type', full_name='pb.CreateOrUpdateRequest.housing_type', index=11,
      number=120, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='wd_type', full_name='pb.CreateOrUpdateRequest.wd_type', index=12,
      number=130, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parking_type', full_name='pb.CreateOrUpdateRequest.parking_type', index=13,
      number=140, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='images', full_name='pb.CreateOrUpdateRequest.images', index=14,
      number=150, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='body', full_name='pb.CreateOrUpdateRequest.body', index=15,
      number=160, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lng', full_name='pb.CreateOrUpdateRequest.lng', index=16,
      number=170, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lat', full_name='pb.CreateOrUpdateRequest.lat', index=17,
      number=180, type=2, cpp_type=6, label=1,
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
  serialized_start=24,
  serialized_end=345,
)


_CREATEORUPDATERESPONSE = _descriptor.Descriptor(
  name='CreateOrUpdateResponse',
  full_name='pb.CreateOrUpdateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.CreateOrUpdateResponse.id', index=0,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=347,
  serialized_end=383,
)


_STREAMREQUEST = _descriptor.Descriptor(
  name='StreamRequest',
  full_name='pb.StreamRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.StreamRequest.id', index=0,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lastSeen', full_name='pb.StreamRequest.lastSeen', index=1,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=385,
  serialized_end=430,
)


_AGGREGATEEVENT = _descriptor.Descriptor(
  name='AggregateEvent',
  full_name='pb.AggregateEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.AggregateEvent.id', index=0,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='aggregateId', full_name='pb.AggregateEvent.aggregateId', index=1,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='pb.AggregateEvent.timestamp', index=2,
      number=30, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='pb.AggregateEvent.version', index=3,
      number=40, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='eventType', full_name='pb.AggregateEvent.eventType', index=4,
      number=50, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='event', full_name='pb.AggregateEvent.event', index=5,
      number=60, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=432,
  serialized_end=551,
)


_ENDEDREQUEST = _descriptor.Descriptor(
  name='EndedRequest',
  full_name='pb.EndedRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
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
  serialized_start=553,
  serialized_end=567,
)


_ENDEDRESPONSE = _descriptor.Descriptor(
  name='EndedResponse',
  full_name='pb.EndedResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
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
  serialized_start=569,
  serialized_end=584,
)

DESCRIPTOR.message_types_by_name['CreateOrUpdateRequest'] = _CREATEORUPDATEREQUEST
DESCRIPTOR.message_types_by_name['CreateOrUpdateResponse'] = _CREATEORUPDATERESPONSE
DESCRIPTOR.message_types_by_name['StreamRequest'] = _STREAMREQUEST
DESCRIPTOR.message_types_by_name['AggregateEvent'] = _AGGREGATEEVENT
DESCRIPTOR.message_types_by_name['EndedRequest'] = _ENDEDREQUEST
DESCRIPTOR.message_types_by_name['EndedResponse'] = _ENDEDRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateOrUpdateRequest = _reflection.GeneratedProtocolMessageType('CreateOrUpdateRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEORUPDATEREQUEST,
  __module__ = 'apthunter_pb2'
  # @@protoc_insertion_point(class_scope:pb.CreateOrUpdateRequest)
  ))
_sym_db.RegisterMessage(CreateOrUpdateRequest)

CreateOrUpdateResponse = _reflection.GeneratedProtocolMessageType('CreateOrUpdateResponse', (_message.Message,), dict(
  DESCRIPTOR = _CREATEORUPDATERESPONSE,
  __module__ = 'apthunter_pb2'
  # @@protoc_insertion_point(class_scope:pb.CreateOrUpdateResponse)
  ))
_sym_db.RegisterMessage(CreateOrUpdateResponse)

StreamRequest = _reflection.GeneratedProtocolMessageType('StreamRequest', (_message.Message,), dict(
  DESCRIPTOR = _STREAMREQUEST,
  __module__ = 'apthunter_pb2'
  # @@protoc_insertion_point(class_scope:pb.StreamRequest)
  ))
_sym_db.RegisterMessage(StreamRequest)

AggregateEvent = _reflection.GeneratedProtocolMessageType('AggregateEvent', (_message.Message,), dict(
  DESCRIPTOR = _AGGREGATEEVENT,
  __module__ = 'apthunter_pb2'
  # @@protoc_insertion_point(class_scope:pb.AggregateEvent)
  ))
_sym_db.RegisterMessage(AggregateEvent)

EndedRequest = _reflection.GeneratedProtocolMessageType('EndedRequest', (_message.Message,), dict(
  DESCRIPTOR = _ENDEDREQUEST,
  __module__ = 'apthunter_pb2'
  # @@protoc_insertion_point(class_scope:pb.EndedRequest)
  ))
_sym_db.RegisterMessage(EndedRequest)

EndedResponse = _reflection.GeneratedProtocolMessageType('EndedResponse', (_message.Message,), dict(
  DESCRIPTOR = _ENDEDRESPONSE,
  __module__ = 'apthunter_pb2'
  # @@protoc_insertion_point(class_scope:pb.EndedResponse)
  ))
_sym_db.RegisterMessage(EndedResponse)



_WRITER = _descriptor.ServiceDescriptor(
  name='Writer',
  full_name='pb.Writer',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=587,
  serialized_end=723,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateOrUpdate',
    full_name='pb.Writer.CreateOrUpdate',
    index=0,
    containing_service=None,
    input_type=_CREATEORUPDATEREQUEST,
    output_type=_CREATEORUPDATERESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Stream',
    full_name='pb.Writer.Stream',
    index=1,
    containing_service=None,
    input_type=_STREAMREQUEST,
    output_type=_AGGREGATEEVENT,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_WRITER)

DESCRIPTOR.services_by_name['Writer'] = _WRITER

# @@protoc_insertion_point(module_scope)
