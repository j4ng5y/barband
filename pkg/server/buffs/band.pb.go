// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: band.proto

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

type Band struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Band) Reset() {
	*x = Band{}
	if protoimpl.UnsafeEnabled {
		mi := &file_band_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Band) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Band) ProtoMessage() {}

func (x *Band) ProtoReflect() protoreflect.Message {
	mi := &file_band_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Band.ProtoReflect.Descriptor instead.
func (*Band) Descriptor() ([]byte, []int) {
	return file_band_proto_rawDescGZIP(), []int{0}
}

func (x *Band) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Band) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateBandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBandRequest) Reset() {
	*x = CreateBandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_band_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBandRequest) ProtoMessage() {}

func (x *CreateBandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_band_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBandRequest.ProtoReflect.Descriptor instead.
func (*CreateBandRequest) Descriptor() ([]byte, []int) {
	return file_band_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ReadBandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadBandRequest) Reset() {
	*x = ReadBandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_band_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBandRequest) ProtoMessage() {}

func (x *ReadBandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_band_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBandRequest.ProtoReflect.Descriptor instead.
func (*ReadBandRequest) Descriptor() ([]byte, []int) {
	return file_band_proto_rawDescGZIP(), []int{2}
}

func (x *ReadBandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateBandRequest) Reset() {
	*x = UpdateBandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_band_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBandRequest) ProtoMessage() {}

func (x *UpdateBandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_band_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBandRequest.ProtoReflect.Descriptor instead.
func (*UpdateBandRequest) Descriptor() ([]byte, []int) {
	return file_band_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteBandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBandRequest) Reset() {
	*x = DeleteBandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_band_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBandRequest) ProtoMessage() {}

func (x *DeleteBandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_band_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBandRequest.ProtoReflect.Descriptor instead.
func (*DeleteBandRequest) Descriptor() ([]byte, []int) {
	return file_band_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_band_proto protoreflect.FileDescriptor

var file_band_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61,
	0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a,
	0x0f, 0x52, 0x65, 0x61, 0x64, 0x42, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x37, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa0,
	0x02, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x62,
	0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x42, 0x61, 0x6e,
	0x64, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x42, 0x61, 0x6e, 0x64, 0x12,
	0x1d, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x42, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x42, 0x61,
	0x6e, 0x64, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x62, 0x61, 0x6e,
	0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x62, 0x61,
	0x6e, 0x64, 0x2e, 0x42, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e,
	0x64, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x61, 0x72, 0x62, 0x61,
	0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x34, 0x6e, 0x67, 0x35, 0x79, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x61, 0x6e, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_band_proto_rawDescOnce sync.Once
	file_band_proto_rawDescData = file_band_proto_rawDesc
)

func file_band_proto_rawDescGZIP() []byte {
	file_band_proto_rawDescOnce.Do(func() {
		file_band_proto_rawDescData = protoimpl.X.CompressGZIP(file_band_proto_rawDescData)
	})
	return file_band_proto_rawDescData
}

var file_band_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_band_proto_goTypes = []interface{}{
	(*Band)(nil),              // 0: barband.band.Band
	(*CreateBandRequest)(nil), // 1: barband.band.CreateBandRequest
	(*ReadBandRequest)(nil),   // 2: barband.band.ReadBandRequest
	(*UpdateBandRequest)(nil), // 3: barband.band.UpdateBandRequest
	(*DeleteBandRequest)(nil), // 4: barband.band.DeleteBandRequest
	(*Empty)(nil),             // 5: barband.common.Empty
}
var file_band_proto_depIdxs = []int32{
	1, // 0: barband.band.BandService.CreateBand:input_type -> barband.band.CreateBandRequest
	2, // 1: barband.band.BandService.ReadBand:input_type -> barband.band.ReadBandRequest
	3, // 2: barband.band.BandService.UpdateBand:input_type -> barband.band.UpdateBandRequest
	4, // 3: barband.band.BandService.DeleteBand:input_type -> barband.band.DeleteBandRequest
	0, // 4: barband.band.BandService.CreateBand:output_type -> barband.band.Band
	0, // 5: barband.band.BandService.ReadBand:output_type -> barband.band.Band
	0, // 6: barband.band.BandService.UpdateBand:output_type -> barband.band.Band
	5, // 7: barband.band.BandService.DeleteBand:output_type -> barband.common.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_band_proto_init() }
func file_band_proto_init() {
	if File_band_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_band_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Band); i {
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
		file_band_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBandRequest); i {
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
		file_band_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBandRequest); i {
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
		file_band_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBandRequest); i {
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
		file_band_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBandRequest); i {
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
			RawDescriptor: file_band_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_band_proto_goTypes,
		DependencyIndexes: file_band_proto_depIdxs,
		MessageInfos:      file_band_proto_msgTypes,
	}.Build()
	File_band_proto = out.File
	file_band_proto_rawDesc = nil
	file_band_proto_goTypes = nil
	file_band_proto_depIdxs = nil
}
