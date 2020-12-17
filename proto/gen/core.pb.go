// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: core.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StateBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]*StateBag_Value `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StateBag) Reset() {
	*x = StateBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateBag) ProtoMessage() {}

func (x *StateBag) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateBag.ProtoReflect.Descriptor instead.
func (*StateBag) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

func (x *StateBag) GetData() map[string]*StateBag_Value {
	if x != nil {
		return x.Data
	}
	return nil
}

type Machine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Machine) Reset() {
	*x = Machine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

type StateBag_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*StateBag_Value_Text
	//	*StateBag_Value_Map
	Value isStateBag_Value_Value `protobuf_oneof:"value"`
}

func (x *StateBag_Value) Reset() {
	*x = StateBag_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateBag_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateBag_Value) ProtoMessage() {}

func (x *StateBag_Value) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateBag_Value.ProtoReflect.Descriptor instead.
func (*StateBag_Value) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0, 1}
}

func (m *StateBag_Value) GetValue() isStateBag_Value_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *StateBag_Value) GetText() string {
	if x, ok := x.GetValue().(*StateBag_Value_Text); ok {
		return x.Text
	}
	return ""
}

func (x *StateBag_Value) GetMap() *any.Any {
	if x, ok := x.GetValue().(*StateBag_Value_Map); ok {
		return x.Map
	}
	return nil
}

type isStateBag_Value_Value interface {
	isStateBag_Value_Value()
}

type StateBag_Value_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type StateBag_Value_Map struct {
	Map *any.Any `protobuf:"bytes,2,opt,name=map,proto3,oneof"`
}

func (*StateBag_Value_Text) isStateBag_Value_Value() {}

func (*StateBag_Value_Map) isStateBag_Value_Value() {}

type Machine_DataDirResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Machine_DataDirResp) Reset() {
	*x = Machine_DataDirResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine_DataDirResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine_DataDirResp) ProtoMessage() {}

func (x *Machine_DataDirResp) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine_DataDirResp.ProtoReflect.Descriptor instead.
func (*Machine_DataDirResp) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Machine_DataDirResp) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Machine_IDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Machine_IDResp) Reset() {
	*x = Machine_IDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine_IDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine_IDResp) ProtoMessage() {}

func (x *Machine_IDResp) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine_IDResp.ProtoReflect.Descriptor instead.
func (*Machine_IDResp) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Machine_IDResp) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x73, 0x64, 0x6b, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x08, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x67, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x5e, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x46, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x1a, 0x21, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x18, 0x0a, 0x06, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32,
	0xb4, 0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x42, 0x6f, 0x78, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x2e, 0x42,
	0x6f, 0x78, 0x12, 0x4d, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x72, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x4e, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x41, 0x72, 0x67, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x43, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x25, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_core_proto_goTypes = []interface{}{
	(*StateBag)(nil),            // 0: hashicorp.vagrant.sdk.StateBag
	(*Machine)(nil),             // 1: hashicorp.vagrant.sdk.Machine
	nil,                         // 2: hashicorp.vagrant.sdk.StateBag.DataEntry
	(*StateBag_Value)(nil),      // 3: hashicorp.vagrant.sdk.StateBag.Value
	(*Machine_DataDirResp)(nil), // 4: hashicorp.vagrant.sdk.Machine.DataDirResp
	(*Machine_IDResp)(nil),      // 5: hashicorp.vagrant.sdk.Machine.IDResp
	(*any.Any)(nil),             // 6: google.protobuf.Any
	(*empty.Empty)(nil),         // 7: google.protobuf.Empty
	(*Args_Box)(nil),            // 8: hashicorp.vagrant.sdk.Args.Box
	(*Args_Environment)(nil),    // 9: hashicorp.vagrant.sdk.Args.Environment
}
var file_core_proto_depIdxs = []int32{
	2, // 0: hashicorp.vagrant.sdk.StateBag.data:type_name -> hashicorp.vagrant.sdk.StateBag.DataEntry
	3, // 1: hashicorp.vagrant.sdk.StateBag.DataEntry.value:type_name -> hashicorp.vagrant.sdk.StateBag.Value
	6, // 2: hashicorp.vagrant.sdk.StateBag.Value.map:type_name -> google.protobuf.Any
	7, // 3: hashicorp.vagrant.sdk.MachineService.Box:input_type -> google.protobuf.Empty
	7, // 4: hashicorp.vagrant.sdk.MachineService.DataDir:input_type -> google.protobuf.Empty
	7, // 5: hashicorp.vagrant.sdk.MachineService.Environment:input_type -> google.protobuf.Empty
	7, // 6: hashicorp.vagrant.sdk.MachineService.ID:input_type -> google.protobuf.Empty
	8, // 7: hashicorp.vagrant.sdk.MachineService.Box:output_type -> hashicorp.vagrant.sdk.Args.Box
	4, // 8: hashicorp.vagrant.sdk.MachineService.DataDir:output_type -> hashicorp.vagrant.sdk.Machine.DataDirResp
	9, // 9: hashicorp.vagrant.sdk.MachineService.Environment:output_type -> hashicorp.vagrant.sdk.Args.Environment
	5, // 10: hashicorp.vagrant.sdk.MachineService.ID:output_type -> hashicorp.vagrant.sdk.Machine.IDResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	file_plugin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateBag); i {
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
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine); i {
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
		file_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateBag_Value); i {
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
		file_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine_DataDirResp); i {
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
		file_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine_IDResp); i {
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
	file_core_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*StateBag_Value_Text)(nil),
		(*StateBag_Value_Map)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MachineServiceClient is the client API for MachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MachineServiceClient interface {
	Box(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Args_Box, error)
	DataDir(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Machine_DataDirResp, error)
	Environment(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Args_Environment, error)
	ID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Machine_IDResp, error)
}

type machineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineServiceClient(cc grpc.ClientConnInterface) MachineServiceClient {
	return &machineServiceClient{cc}
}

func (c *machineServiceClient) Box(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Args_Box, error) {
	out := new(Args_Box)
	err := c.cc.Invoke(ctx, "/hashicorp.vagrant.sdk.MachineService/Box", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) DataDir(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Machine_DataDirResp, error) {
	out := new(Machine_DataDirResp)
	err := c.cc.Invoke(ctx, "/hashicorp.vagrant.sdk.MachineService/DataDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) Environment(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Args_Environment, error) {
	out := new(Args_Environment)
	err := c.cc.Invoke(ctx, "/hashicorp.vagrant.sdk.MachineService/Environment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) ID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Machine_IDResp, error) {
	out := new(Machine_IDResp)
	err := c.cc.Invoke(ctx, "/hashicorp.vagrant.sdk.MachineService/ID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServiceServer is the server API for MachineService service.
type MachineServiceServer interface {
	Box(context.Context, *empty.Empty) (*Args_Box, error)
	DataDir(context.Context, *empty.Empty) (*Machine_DataDirResp, error)
	Environment(context.Context, *empty.Empty) (*Args_Environment, error)
	ID(context.Context, *empty.Empty) (*Machine_IDResp, error)
}

// UnimplementedMachineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMachineServiceServer struct {
}

func (*UnimplementedMachineServiceServer) Box(context.Context, *empty.Empty) (*Args_Box, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Box not implemented")
}
func (*UnimplementedMachineServiceServer) DataDir(context.Context, *empty.Empty) (*Machine_DataDirResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataDir not implemented")
}
func (*UnimplementedMachineServiceServer) Environment(context.Context, *empty.Empty) (*Args_Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Environment not implemented")
}
func (*UnimplementedMachineServiceServer) ID(context.Context, *empty.Empty) (*Machine_IDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ID not implemented")
}

func RegisterMachineServiceServer(s *grpc.Server, srv MachineServiceServer) {
	s.RegisterService(&_MachineService_serviceDesc, srv)
}

func _MachineService_Box_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).Box(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.vagrant.sdk.MachineService/Box",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).Box(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_DataDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).DataDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.vagrant.sdk.MachineService/DataDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).DataDir(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_Environment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).Environment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.vagrant.sdk.MachineService/Environment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).Environment(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_ID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).ID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.vagrant.sdk.MachineService/ID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).ID(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MachineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.vagrant.sdk.MachineService",
	HandlerType: (*MachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Box",
			Handler:    _MachineService_Box_Handler,
		},
		{
			MethodName: "DataDir",
			Handler:    _MachineService_DataDir_Handler,
		},
		{
			MethodName: "Environment",
			Handler:    _MachineService_Environment_Handler,
		},
		{
			MethodName: "ID",
			Handler:    _MachineService_ID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}
