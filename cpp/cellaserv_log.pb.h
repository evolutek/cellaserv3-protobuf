// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cellaserv_log.proto

#ifndef PROTOBUF_INCLUDED_cellaserv_5flog_2eproto
#define PROTOBUF_INCLUDED_cellaserv_5flog_2eproto

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3006001
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3006001 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#define PROTOBUF_INTERNAL_EXPORT_protobuf_cellaserv_5flog_2eproto 

namespace protobuf_cellaserv_5flog_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[1];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors();
}  // namespace protobuf_cellaserv_5flog_2eproto
namespace cellaserv {
class LogMessage;
class LogMessageDefaultTypeInternal;
extern LogMessageDefaultTypeInternal _LogMessage_default_instance_;
}  // namespace cellaserv
namespace google {
namespace protobuf {
template<> ::cellaserv::LogMessage* Arena::CreateMaybeMessage<::cellaserv::LogMessage>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace cellaserv {

// ===================================================================

class LogMessage : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:cellaserv.LogMessage) */ {
 public:
  LogMessage();
  virtual ~LogMessage();

  LogMessage(const LogMessage& from);

  inline LogMessage& operator=(const LogMessage& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  LogMessage(LogMessage&& from) noexcept
    : LogMessage() {
    *this = ::std::move(from);
  }

  inline LogMessage& operator=(LogMessage&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _internal_metadata_.unknown_fields();
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return _internal_metadata_.mutable_unknown_fields();
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const LogMessage& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const LogMessage* internal_default_instance() {
    return reinterpret_cast<const LogMessage*>(
               &_LogMessage_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(LogMessage* other);
  friend void swap(LogMessage& a, LogMessage& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline LogMessage* New() const final {
    return CreateMaybeMessage<LogMessage>(NULL);
  }

  LogMessage* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<LogMessage>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const LogMessage& from);
  void MergeFrom(const LogMessage& from);
  void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(LogMessage* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // required string sender = 1;
  bool has_sender() const;
  void clear_sender();
  static const int kSenderFieldNumber = 1;
  const ::std::string& sender() const;
  void set_sender(const ::std::string& value);
  #if LANG_CXX11
  void set_sender(::std::string&& value);
  #endif
  void set_sender(const char* value);
  void set_sender(const char* value, size_t size);
  ::std::string* mutable_sender();
  ::std::string* release_sender();
  void set_allocated_sender(::std::string* sender);

  // required bytes content = 2;
  bool has_content() const;
  void clear_content();
  static const int kContentFieldNumber = 2;
  const ::std::string& content() const;
  void set_content(const ::std::string& value);
  #if LANG_CXX11
  void set_content(::std::string&& value);
  #endif
  void set_content(const char* value);
  void set_content(const void* value, size_t size);
  ::std::string* mutable_content();
  ::std::string* release_content();
  void set_allocated_content(::std::string* content);

  // optional string destination = 3;
  bool has_destination() const;
  void clear_destination();
  static const int kDestinationFieldNumber = 3;
  const ::std::string& destination() const;
  void set_destination(const ::std::string& value);
  #if LANG_CXX11
  void set_destination(::std::string&& value);
  #endif
  void set_destination(const char* value);
  void set_destination(const char* value, size_t size);
  ::std::string* mutable_destination();
  ::std::string* release_destination();
  void set_allocated_destination(::std::string* destination);

  // @@protoc_insertion_point(class_scope:cellaserv.LogMessage)
 private:
  void set_has_sender();
  void clear_has_sender();
  void set_has_destination();
  void clear_has_destination();
  void set_has_content();
  void clear_has_content();

  // helper for ByteSizeLong()
  size_t RequiredFieldsByteSizeFallback() const;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::HasBits<1> _has_bits_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  ::google::protobuf::internal::ArenaStringPtr sender_;
  ::google::protobuf::internal::ArenaStringPtr content_;
  ::google::protobuf::internal::ArenaStringPtr destination_;
  friend struct ::protobuf_cellaserv_5flog_2eproto::TableStruct;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// LogMessage

// required string sender = 1;
inline bool LogMessage::has_sender() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void LogMessage::set_has_sender() {
  _has_bits_[0] |= 0x00000001u;
}
inline void LogMessage::clear_has_sender() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void LogMessage::clear_sender() {
  sender_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_sender();
}
inline const ::std::string& LogMessage::sender() const {
  // @@protoc_insertion_point(field_get:cellaserv.LogMessage.sender)
  return sender_.GetNoArena();
}
inline void LogMessage::set_sender(const ::std::string& value) {
  set_has_sender();
  sender_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cellaserv.LogMessage.sender)
}
#if LANG_CXX11
inline void LogMessage::set_sender(::std::string&& value) {
  set_has_sender();
  sender_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cellaserv.LogMessage.sender)
}
#endif
inline void LogMessage::set_sender(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  set_has_sender();
  sender_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cellaserv.LogMessage.sender)
}
inline void LogMessage::set_sender(const char* value, size_t size) {
  set_has_sender();
  sender_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cellaserv.LogMessage.sender)
}
inline ::std::string* LogMessage::mutable_sender() {
  set_has_sender();
  // @@protoc_insertion_point(field_mutable:cellaserv.LogMessage.sender)
  return sender_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* LogMessage::release_sender() {
  // @@protoc_insertion_point(field_release:cellaserv.LogMessage.sender)
  if (!has_sender()) {
    return NULL;
  }
  clear_has_sender();
  return sender_.ReleaseNonDefaultNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void LogMessage::set_allocated_sender(::std::string* sender) {
  if (sender != NULL) {
    set_has_sender();
  } else {
    clear_has_sender();
  }
  sender_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), sender);
  // @@protoc_insertion_point(field_set_allocated:cellaserv.LogMessage.sender)
}

// optional string destination = 3;
inline bool LogMessage::has_destination() const {
  return (_has_bits_[0] & 0x00000004u) != 0;
}
inline void LogMessage::set_has_destination() {
  _has_bits_[0] |= 0x00000004u;
}
inline void LogMessage::clear_has_destination() {
  _has_bits_[0] &= ~0x00000004u;
}
inline void LogMessage::clear_destination() {
  destination_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_destination();
}
inline const ::std::string& LogMessage::destination() const {
  // @@protoc_insertion_point(field_get:cellaserv.LogMessage.destination)
  return destination_.GetNoArena();
}
inline void LogMessage::set_destination(const ::std::string& value) {
  set_has_destination();
  destination_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cellaserv.LogMessage.destination)
}
#if LANG_CXX11
inline void LogMessage::set_destination(::std::string&& value) {
  set_has_destination();
  destination_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cellaserv.LogMessage.destination)
}
#endif
inline void LogMessage::set_destination(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  set_has_destination();
  destination_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cellaserv.LogMessage.destination)
}
inline void LogMessage::set_destination(const char* value, size_t size) {
  set_has_destination();
  destination_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cellaserv.LogMessage.destination)
}
inline ::std::string* LogMessage::mutable_destination() {
  set_has_destination();
  // @@protoc_insertion_point(field_mutable:cellaserv.LogMessage.destination)
  return destination_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* LogMessage::release_destination() {
  // @@protoc_insertion_point(field_release:cellaserv.LogMessage.destination)
  if (!has_destination()) {
    return NULL;
  }
  clear_has_destination();
  return destination_.ReleaseNonDefaultNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void LogMessage::set_allocated_destination(::std::string* destination) {
  if (destination != NULL) {
    set_has_destination();
  } else {
    clear_has_destination();
  }
  destination_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), destination);
  // @@protoc_insertion_point(field_set_allocated:cellaserv.LogMessage.destination)
}

// required bytes content = 2;
inline bool LogMessage::has_content() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void LogMessage::set_has_content() {
  _has_bits_[0] |= 0x00000002u;
}
inline void LogMessage::clear_has_content() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void LogMessage::clear_content() {
  content_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_content();
}
inline const ::std::string& LogMessage::content() const {
  // @@protoc_insertion_point(field_get:cellaserv.LogMessage.content)
  return content_.GetNoArena();
}
inline void LogMessage::set_content(const ::std::string& value) {
  set_has_content();
  content_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cellaserv.LogMessage.content)
}
#if LANG_CXX11
inline void LogMessage::set_content(::std::string&& value) {
  set_has_content();
  content_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cellaserv.LogMessage.content)
}
#endif
inline void LogMessage::set_content(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  set_has_content();
  content_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cellaserv.LogMessage.content)
}
inline void LogMessage::set_content(const void* value, size_t size) {
  set_has_content();
  content_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cellaserv.LogMessage.content)
}
inline ::std::string* LogMessage::mutable_content() {
  set_has_content();
  // @@protoc_insertion_point(field_mutable:cellaserv.LogMessage.content)
  return content_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* LogMessage::release_content() {
  // @@protoc_insertion_point(field_release:cellaserv.LogMessage.content)
  if (!has_content()) {
    return NULL;
  }
  clear_has_content();
  return content_.ReleaseNonDefaultNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void LogMessage::set_allocated_content(::std::string* content) {
  if (content != NULL) {
    set_has_content();
  } else {
    clear_has_content();
  }
  content_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), content);
  // @@protoc_insertion_point(field_set_allocated:cellaserv.LogMessage.content)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace cellaserv

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_INCLUDED_cellaserv_5flog_2eproto
