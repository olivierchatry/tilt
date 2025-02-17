// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.18.1
// source: pkg/webview/log.proto

package webview

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LogLevel int32

const (
	// For backwards-compatibility, the int value doesn't say
	// anything about relative severity.
	LogLevel_NONE    LogLevel = 0
	LogLevel_INFO    LogLevel = 1
	LogLevel_VERBOSE LogLevel = 2
	LogLevel_DEBUG   LogLevel = 3
	LogLevel_WARN    LogLevel = 4
	LogLevel_ERROR   LogLevel = 5
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "NONE",
		1: "INFO",
		2: "VERBOSE",
		3: "DEBUG",
		4: "WARN",
		5: "ERROR",
	}
	LogLevel_value = map[string]int32{
		"NONE":    0,
		"INFO":    1,
		"VERBOSE": 2,
		"DEBUG":   3,
		"WARN":    4,
		"ERROR":   5,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_webview_log_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_pkg_webview_log_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_pkg_webview_log_proto_rawDescGZIP(), []int{0}
}

type LogSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpanId string                 `protobuf:"bytes,1,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Time   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Text   string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Level  LogLevel               `protobuf:"varint,4,opt,name=level,proto3,enum=webview.LogLevel" json:"level,omitempty"`
	// When we store warnings in the LogStore, we break them up into lines and
	// store them as a series of line segments. 'anchor' marks the beginning of a
	// series of logs that should be kept together.
	//
	// Anchor warning1, line1
	//        warning1, line2
	// Anchor warning2, line1
	Anchor bool `protobuf:"varint,5,opt,name=anchor,proto3" json:"anchor,omitempty"`
	// Context-specific optional fields for a log segment.
	// Used for experimenting with new types of log metadata.
	Fields map[string]string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LogSegment) Reset() {
	*x = LogSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_webview_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSegment) ProtoMessage() {}

func (x *LogSegment) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_webview_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSegment.ProtoReflect.Descriptor instead.
func (*LogSegment) Descriptor() ([]byte, []int) {
	return file_pkg_webview_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogSegment) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *LogSegment) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *LogSegment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *LogSegment) GetLevel() LogLevel {
	if x != nil {
		return x.Level
	}
	return LogLevel_NONE
}

func (x *LogSegment) GetAnchor() bool {
	if x != nil {
		return x.Anchor
	}
	return false
}

func (x *LogSegment) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type LogSpan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManifestName string `protobuf:"bytes,1,opt,name=manifest_name,json=manifestName,proto3" json:"manifest_name,omitempty"`
}

func (x *LogSpan) Reset() {
	*x = LogSpan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_webview_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSpan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSpan) ProtoMessage() {}

func (x *LogSpan) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_webview_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSpan.ProtoReflect.Descriptor instead.
func (*LogSpan) Descriptor() ([]byte, []int) {
	return file_pkg_webview_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogSpan) GetManifestName() string {
	if x != nil {
		return x.ManifestName
	}
	return ""
}

type LogList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spans    map[string]*LogSpan `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Segments []*LogSegment       `protobuf:"bytes,2,rep,name=segments,proto3" json:"segments,omitempty"`
	// from_checkpoint and to_checkpoint express an interval on the
	// central log-store, with an inclusive start and an exclusive end
	//
	// [from_checkpoint, to_checkpoint)
	//
	// An interval of [0, 0) means that the server isn't using
	// the incremental load protocol.
	//
	// An interval of [-1, -1) means that the server doesn't have new logs
	// to send down.
	FromCheckpoint int32 `protobuf:"varint,3,opt,name=from_checkpoint,json=fromCheckpoint,proto3" json:"from_checkpoint,omitempty"`
	ToCheckpoint   int32 `protobuf:"varint,4,opt,name=to_checkpoint,json=toCheckpoint,proto3" json:"to_checkpoint,omitempty"`
}

func (x *LogList) Reset() {
	*x = LogList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_webview_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogList) ProtoMessage() {}

func (x *LogList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_webview_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogList.ProtoReflect.Descriptor instead.
func (*LogList) Descriptor() ([]byte, []int) {
	return file_pkg_webview_log_proto_rawDescGZIP(), []int{2}
}

func (x *LogList) GetSpans() map[string]*LogSpan {
	if x != nil {
		return x.Spans
	}
	return nil
}

func (x *LogList) GetSegments() []*LogSegment {
	if x != nil {
		return x.Segments
	}
	return nil
}

func (x *LogList) GetFromCheckpoint() int32 {
	if x != nil {
		return x.FromCheckpoint
	}
	return 0
}

func (x *LogList) GetToCheckpoint() int32 {
	if x != nil {
		return x.ToCheckpoint
	}
	return 0
}

var File_pkg_webview_log_proto protoreflect.FileDescriptor

var file_pkg_webview_log_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x77,
	0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x37,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x77, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x2e, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x87, 0x02, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x77, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x53, 0x70, 0x61, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x73, 0x70, 0x61, 0x6e,
	0x73, 0x12, 0x2f, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x72, 0x6f,
	0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x6f, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x1a, 0x4a, 0x0a, 0x0a, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x77, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x70, 0x61,
	0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x4b, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42,
	0x55, 0x47, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x04, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6c, 0x74, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x74, 0x69, 0x6c, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x65, 0x62, 0x76, 0x69, 0x65,
	0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_webview_log_proto_rawDescOnce sync.Once
	file_pkg_webview_log_proto_rawDescData = file_pkg_webview_log_proto_rawDesc
)

func file_pkg_webview_log_proto_rawDescGZIP() []byte {
	file_pkg_webview_log_proto_rawDescOnce.Do(func() {
		file_pkg_webview_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_webview_log_proto_rawDescData)
	})
	return file_pkg_webview_log_proto_rawDescData
}

var file_pkg_webview_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_webview_log_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_webview_log_proto_goTypes = []interface{}{
	(LogLevel)(0),                 // 0: webview.LogLevel
	(*LogSegment)(nil),            // 1: webview.LogSegment
	(*LogSpan)(nil),               // 2: webview.LogSpan
	(*LogList)(nil),               // 3: webview.LogList
	nil,                           // 4: webview.LogSegment.FieldsEntry
	nil,                           // 5: webview.LogList.SpansEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_pkg_webview_log_proto_depIdxs = []int32{
	6, // 0: webview.LogSegment.time:type_name -> google.protobuf.Timestamp
	0, // 1: webview.LogSegment.level:type_name -> webview.LogLevel
	4, // 2: webview.LogSegment.fields:type_name -> webview.LogSegment.FieldsEntry
	5, // 3: webview.LogList.spans:type_name -> webview.LogList.SpansEntry
	1, // 4: webview.LogList.segments:type_name -> webview.LogSegment
	2, // 5: webview.LogList.SpansEntry.value:type_name -> webview.LogSpan
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_webview_log_proto_init() }
func file_pkg_webview_log_proto_init() {
	if File_pkg_webview_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_webview_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSegment); i {
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
		file_pkg_webview_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSpan); i {
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
		file_pkg_webview_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogList); i {
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
			RawDescriptor: file_pkg_webview_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_webview_log_proto_goTypes,
		DependencyIndexes: file_pkg_webview_log_proto_depIdxs,
		EnumInfos:         file_pkg_webview_log_proto_enumTypes,
		MessageInfos:      file_pkg_webview_log_proto_msgTypes,
	}.Build()
	File_pkg_webview_log_proto = out.File
	file_pkg_webview_log_proto_rawDesc = nil
	file_pkg_webview_log_proto_goTypes = nil
	file_pkg_webview_log_proto_depIdxs = nil
}
