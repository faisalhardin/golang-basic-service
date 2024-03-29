// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: summary.proto

package summary_proto

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

type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock         string `protobuf:"bytes,1,opt,name=stock,proto3" json:"stock,omitempty"`
	Previousprice int64  `protobuf:"varint,2,opt,name=previousprice,proto3" json:"previousprice,omitempty"`
	Openprice     int64  `protobuf:"varint,3,opt,name=openprice,proto3" json:"openprice,omitempty"`
	Highestprice  int64  `protobuf:"varint,4,opt,name=highestprice,proto3" json:"highestprice,omitempty"`
	Lowestprice   int64  `protobuf:"varint,5,opt,name=lowestprice,proto3" json:"lowestprice,omitempty"`
	Closeprice    int64  `protobuf:"varint,6,opt,name=closeprice,proto3" json:"closeprice,omitempty"`
	Volume        int64  `protobuf:"varint,7,opt,name=volume,proto3" json:"volume,omitempty"`
	Value         int64  `protobuf:"varint,8,opt,name=value,proto3" json:"value,omitempty"`
	Average       int64  `protobuf:"varint,9,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{0}
}

func (x *Summary) GetStock() string {
	if x != nil {
		return x.Stock
	}
	return ""
}

func (x *Summary) GetPreviousprice() int64 {
	if x != nil {
		return x.Previousprice
	}
	return 0
}

func (x *Summary) GetOpenprice() int64 {
	if x != nil {
		return x.Openprice
	}
	return 0
}

func (x *Summary) GetHighestprice() int64 {
	if x != nil {
		return x.Highestprice
	}
	return 0
}

func (x *Summary) GetLowestprice() int64 {
	if x != nil {
		return x.Lowestprice
	}
	return 0
}

func (x *Summary) GetCloseprice() int64 {
	if x != nil {
		return x.Closeprice
	}
	return 0
}

func (x *Summary) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Summary) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Summary) GetAverage() int64 {
	if x != nil {
		return x.Average
	}
	return 0
}

type GetStockSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock string `protobuf:"bytes,1,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *GetStockSummaryRequest) Reset() {
	*x = GetStockSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockSummaryRequest) ProtoMessage() {}

func (x *GetStockSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockSummaryRequest.ProtoReflect.Descriptor instead.
func (*GetStockSummaryRequest) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{1}
}

func (x *GetStockSummaryRequest) GetStock() string {
	if x != nil {
		return x.Stock
	}
	return ""
}

type GetStockSummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary *Summary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *GetStockSummaryResponse) Reset() {
	*x = GetStockSummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockSummaryResponse) ProtoMessage() {}

func (x *GetStockSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockSummaryResponse.ProtoReflect.Descriptor instead.
func (*GetStockSummaryResponse) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{2}
}

func (x *GetStockSummaryResponse) GetSummary() *Summary {
	if x != nil {
		return x.Summary
	}
	return nil
}

var File_summary_proto protoreflect.FileDescriptor

var file_summary_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91,
	0x02, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x68, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c,
	0x6f, 0x77, 0x65, 0x73, 0x74, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x22, 0x2e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x22, 0x4b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x32,
	0x6d, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x25, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f,
	0x5a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_summary_proto_rawDescOnce sync.Once
	file_summary_proto_rawDescData = file_summary_proto_rawDesc
)

func file_summary_proto_rawDescGZIP() []byte {
	file_summary_proto_rawDescOnce.Do(func() {
		file_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_summary_proto_rawDescData)
	})
	return file_summary_proto_rawDescData
}

var file_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_summary_proto_goTypes = []interface{}{
	(*Summary)(nil),                 // 0: summary_proto.Summary
	(*GetStockSummaryRequest)(nil),  // 1: summary_proto.GetStockSummaryRequest
	(*GetStockSummaryResponse)(nil), // 2: summary_proto.GetStockSummaryResponse
}
var file_summary_proto_depIdxs = []int32{
	0, // 0: summary_proto.GetStockSummaryResponse.summary:type_name -> summary_proto.Summary
	1, // 1: summary_proto.SummaryService.GetSummary:input_type -> summary_proto.GetStockSummaryRequest
	2, // 2: summary_proto.SummaryService.GetSummary:output_type -> summary_proto.GetStockSummaryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_summary_proto_init() }
func file_summary_proto_init() {
	if File_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
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
		file_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockSummaryRequest); i {
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
		file_summary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockSummaryResponse); i {
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
			RawDescriptor: file_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_summary_proto_goTypes,
		DependencyIndexes: file_summary_proto_depIdxs,
		MessageInfos:      file_summary_proto_msgTypes,
	}.Build()
	File_summary_proto = out.File
	file_summary_proto_rawDesc = nil
	file_summary_proto_goTypes = nil
	file_summary_proto_depIdxs = nil
}
