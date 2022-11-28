// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: setlist.proto

package buffs

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

type SetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BandId string `protobuf:"bytes,1,opt,name=band_id,json=bandId,proto3" json:"band_id,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetList) Reset() {
	*x = SetList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetList) ProtoMessage() {}

func (x *SetList) ProtoReflect() protoreflect.Message {
	mi := &file_setlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetList.ProtoReflect.Descriptor instead.
func (*SetList) Descriptor() ([]byte, []int) {
	return file_setlist_proto_rawDescGZIP(), []int{0}
}

func (x *SetList) GetBandId() string {
	if x != nil {
		return x.BandId
	}
	return ""
}

func (x *SetList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateSetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BandId string `protobuf:"bytes,1,opt,name=band_id,json=bandId,proto3" json:"band_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateSetListRequest) Reset() {
	*x = CreateSetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSetListRequest) ProtoMessage() {}

func (x *CreateSetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_setlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSetListRequest.ProtoReflect.Descriptor instead.
func (*CreateSetListRequest) Descriptor() ([]byte, []int) {
	return file_setlist_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSetListRequest) GetBandId() string {
	if x != nil {
		return x.BandId
	}
	return ""
}

func (x *CreateSetListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ReadSetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadSetListRequest) Reset() {
	*x = ReadSetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSetListRequest) ProtoMessage() {}

func (x *ReadSetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_setlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSetListRequest.ProtoReflect.Descriptor instead.
func (*ReadSetListRequest) Descriptor() ([]byte, []int) {
	return file_setlist_proto_rawDescGZIP(), []int{2}
}

func (x *ReadSetListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateSetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateSetListRequest) Reset() {
	*x = UpdateSetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setlist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSetListRequest) ProtoMessage() {}

func (x *UpdateSetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_setlist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSetListRequest.ProtoReflect.Descriptor instead.
func (*UpdateSetListRequest) Descriptor() ([]byte, []int) {
	return file_setlist_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSetListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSetListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteSetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSetListRequest) Reset() {
	*x = DeleteSetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setlist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSetListRequest) ProtoMessage() {}

func (x *DeleteSetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_setlist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSetListRequest.ProtoReflect.Descriptor instead.
func (*DeleteSetListRequest) Descriptor() ([]byte, []int) {
	return file_setlist_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSetListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_setlist_proto protoreflect.FileDescriptor

var file_setlist_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x73, 0x65, 0x74, 0x6c, 0x69, 0x73, 0x74,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46,
	0x0a, 0x07, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x64,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x52,
	0x65, 0x61, 0x64, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3a, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd9, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x62, 0x61, 0x72, 0x62,
	0x61, 0x6e, 0x64, 0x2e, 0x73, 0x65, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x73, 0x65, 0x74, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x62, 0x61,
	0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x73, 0x65, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x73, 0x65, 0x74, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e,
	0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x73, 0x65, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x73,
	0x65, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x25, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x73, 0x65, 0x74, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61,
	0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x34, 0x6e, 0x67, 0x35, 0x79, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_setlist_proto_rawDescOnce sync.Once
	file_setlist_proto_rawDescData = file_setlist_proto_rawDesc
)

func file_setlist_proto_rawDescGZIP() []byte {
	file_setlist_proto_rawDescOnce.Do(func() {
		file_setlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_setlist_proto_rawDescData)
	})
	return file_setlist_proto_rawDescData
}

var file_setlist_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_setlist_proto_goTypes = []interface{}{
	(*SetList)(nil),              // 0: barband.setlist.SetList
	(*CreateSetListRequest)(nil), // 1: barband.setlist.CreateSetListRequest
	(*ReadSetListRequest)(nil),   // 2: barband.setlist.ReadSetListRequest
	(*UpdateSetListRequest)(nil), // 3: barband.setlist.UpdateSetListRequest
	(*DeleteSetListRequest)(nil), // 4: barband.setlist.DeleteSetListRequest
	(*Empty)(nil),                // 5: barband.common.Empty
}
var file_setlist_proto_depIdxs = []int32{
	1, // 0: barband.setlist.SetListService.CreateSetList:input_type -> barband.setlist.CreateSetListRequest
	2, // 1: barband.setlist.SetListService.ReadSetList:input_type -> barband.setlist.ReadSetListRequest
	3, // 2: barband.setlist.SetListService.UpdateSetList:input_type -> barband.setlist.UpdateSetListRequest
	4, // 3: barband.setlist.SetListService.DeleteSetList:input_type -> barband.setlist.DeleteSetListRequest
	0, // 4: barband.setlist.SetListService.CreateSetList:output_type -> barband.setlist.SetList
	0, // 5: barband.setlist.SetListService.ReadSetList:output_type -> barband.setlist.SetList
	0, // 6: barband.setlist.SetListService.UpdateSetList:output_type -> barband.setlist.SetList
	5, // 7: barband.setlist.SetListService.DeleteSetList:output_type -> barband.common.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_setlist_proto_init() }
func file_setlist_proto_init() {
	if File_setlist_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_setlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetList); i {
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
		file_setlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSetListRequest); i {
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
		file_setlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSetListRequest); i {
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
		file_setlist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSetListRequest); i {
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
		file_setlist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSetListRequest); i {
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
			RawDescriptor: file_setlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_setlist_proto_goTypes,
		DependencyIndexes: file_setlist_proto_depIdxs,
		MessageInfos:      file_setlist_proto_msgTypes,
	}.Build()
	File_setlist_proto = out.File
	file_setlist_proto_rawDesc = nil
	file_setlist_proto_goTypes = nil
	file_setlist_proto_depIdxs = nil
}
