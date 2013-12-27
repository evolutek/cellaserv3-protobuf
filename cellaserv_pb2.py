# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cellaserv.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import metaclass as _metaclass
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)




DESCRIPTOR = _descriptor.FileDescriptor(
  name='cellaserv.proto',
  package='cellaserv',
  serialized_pb=b'\n\x0f\x63\x65llaserv.proto\x12\tcellaserv\"\x99\x01\n\x07Message\x12,\n\x04type\x18\x01 \x02(\x0e\x32\x1e.cellaserv.Message.MessageType\x12\x0f\n\x07\x63ontent\x18\x02 \x02(\x0c\"O\n\x0bMessageType\x12\x0c\n\x08Register\x10\x00\x12\x0b\n\x07Request\x10\x01\x12\t\n\x05Reply\x10\x02\x12\r\n\tSubscribe\x10\x03\x12\x0b\n\x07Publish\x10\x04\"0\n\x08Register\x12\x0c\n\x04name\x18\x01 \x02(\t\x12\x16\n\x0eidentification\x18\x02 \x01(\t\"i\n\x07Request\x12\x14\n\x0cservice_name\x18\x01 \x02(\t\x12\x1e\n\x16service_identification\x18\x02 \x01(\t\x12\x0e\n\x06method\x18\x03 \x02(\t\x12\x0c\n\x04\x64\x61ta\x18\x04 \x01(\x0c\x12\n\n\x02id\x18\x63 \x02(\x04\"!\n\x05Reply\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\x0c\x12\n\n\x02id\x18\x63 \x02(\x04\"\x1a\n\tSubscribe\x12\r\n\x05\x65vent\x18\x01 \x02(\t\"&\n\x07Publish\x12\r\n\x05\x65vent\x18\x01 \x02(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c')



_MESSAGE_MESSAGETYPE = _descriptor.EnumDescriptor(
  name='MessageType',
  full_name='cellaserv.Message.MessageType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Register', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Request', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Reply', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Subscribe', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Publish', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=105,
  serialized_end=184,
)


_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='cellaserv.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='cellaserv.Message.type', index=0,
      number=1, type=14, cpp_type=8, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='content', full_name='cellaserv.Message.content', index=1,
      number=2, type=12, cpp_type=9, label=2,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _MESSAGE_MESSAGETYPE,
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=31,
  serialized_end=184,
)


_REGISTER = _descriptor.Descriptor(
  name='Register',
  full_name='cellaserv.Register',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='cellaserv.Register.name', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='identification', full_name='cellaserv.Register.identification', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=186,
  serialized_end=234,
)


_REQUEST = _descriptor.Descriptor(
  name='Request',
  full_name='cellaserv.Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='service_name', full_name='cellaserv.Request.service_name', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='service_identification', full_name='cellaserv.Request.service_identification', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='method', full_name='cellaserv.Request.method', index=2,
      number=3, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='data', full_name='cellaserv.Request.data', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='cellaserv.Request.id', index=4,
      number=99, type=4, cpp_type=4, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=236,
  serialized_end=341,
)


_REPLY = _descriptor.Descriptor(
  name='Reply',
  full_name='cellaserv.Reply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='cellaserv.Reply.data', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='cellaserv.Reply.id', index=1,
      number=99, type=4, cpp_type=4, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=343,
  serialized_end=376,
)


_SUBSCRIBE = _descriptor.Descriptor(
  name='Subscribe',
  full_name='cellaserv.Subscribe',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='event', full_name='cellaserv.Subscribe.event', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=378,
  serialized_end=404,
)


_PUBLISH = _descriptor.Descriptor(
  name='Publish',
  full_name='cellaserv.Publish',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='event', full_name='cellaserv.Publish.event', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='data', full_name='cellaserv.Publish.data', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=406,
  serialized_end=444,
)

_MESSAGE.fields_by_name['type'].enum_type = _MESSAGE_MESSAGETYPE
_MESSAGE_MESSAGETYPE.containing_type = _MESSAGE;
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
DESCRIPTOR.message_types_by_name['Register'] = _REGISTER
DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
DESCRIPTOR.message_types_by_name['Reply'] = _REPLY
DESCRIPTOR.message_types_by_name['Subscribe'] = _SUBSCRIBE
DESCRIPTOR.message_types_by_name['Publish'] = _PUBLISH

@_metaclass.decorator
class Message(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _MESSAGE

  # @@protoc_insertion_point(class_scope:cellaserv.Message)

@_metaclass.decorator
class Register(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _REGISTER

  # @@protoc_insertion_point(class_scope:cellaserv.Register)

@_metaclass.decorator
class Request(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _REQUEST

  # @@protoc_insertion_point(class_scope:cellaserv.Request)

@_metaclass.decorator
class Reply(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _REPLY

  # @@protoc_insertion_point(class_scope:cellaserv.Reply)

@_metaclass.decorator
class Subscribe(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _SUBSCRIBE

  # @@protoc_insertion_point(class_scope:cellaserv.Subscribe)

@_metaclass.decorator
class Publish(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _PUBLISH

  # @@protoc_insertion_point(class_scope:cellaserv.Publish)


# @@protoc_insertion_point(module_scope)
