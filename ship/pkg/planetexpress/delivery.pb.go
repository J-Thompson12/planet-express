// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: delivery.proto

package planetexpress

import (
	proto "github.com/golang/protobuf/proto"
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

type DeliveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item      string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Delivered bool   `protobuf:"varint,2,opt,name=delivered,proto3" json:"delivered,omitempty"`
}

func (x *DeliveryResponse) Reset() {
	*x = DeliveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryResponse) ProtoMessage() {}

func (x *DeliveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryResponse.ProtoReflect.Descriptor instead.
func (*DeliveryResponse) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryResponse) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

func (x *DeliveryResponse) GetDelivered() bool {
	if x != nil {
		return x.Delivered
	}
	return false
}

type DeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item     string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *DeliveryRequest) Reset() {
	*x = DeliveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryRequest) ProtoMessage() {}

func (x *DeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryRequest.ProtoReflect.Descriptor instead.
func (*DeliveryRequest) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryRequest) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

func (x *DeliveryRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *DeliveryRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type GetDeliveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delivery *DeliveryResponse `protobuf:"bytes,1,opt,name=delivery,proto3" json:"delivery,omitempty"`
}

func (x *GetDeliveryResponse) Reset() {
	*x = GetDeliveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeliveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeliveryResponse) ProtoMessage() {}

func (x *GetDeliveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeliveryResponse.ProtoReflect.Descriptor instead.
func (*GetDeliveryResponse) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeliveryResponse) GetDelivery() *DeliveryResponse {
	if x != nil {
		return x.Delivery
	}
	return nil
}

type GetDeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delivery *DeliveryRequest `protobuf:"bytes,1,opt,name=delivery,proto3" json:"delivery,omitempty"`
}

func (x *GetDeliveryRequest) Reset() {
	*x = GetDeliveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeliveryRequest) ProtoMessage() {}

func (x *GetDeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeliveryRequest.ProtoReflect.Descriptor instead.
func (*GetDeliveryRequest) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeliveryRequest) GetDelivery() *DeliveryRequest {
	if x != nil {
		return x.Delivery
	}
	return nil
}

type ListDeliveriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*DeliveryResponse `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListDeliveriesResponse) Reset() {
	*x = ListDeliveriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeliveriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeliveriesResponse) ProtoMessage() {}

func (x *ListDeliveriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeliveriesResponse.ProtoReflect.Descriptor instead.
func (*ListDeliveriesResponse) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{4}
}

func (x *ListDeliveriesResponse) GetList() []*DeliveryResponse {
	if x != nil {
		return x.List
	}
	return nil
}

var File_delivery_proto protoreflect.FileDescriptor

var file_delivery_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x44, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x65, 0x64, 0x22, 0x5d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x22, 0x50, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a,
	0x0a, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x22, 0x4d, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_delivery_proto_rawDescOnce sync.Once
	file_delivery_proto_rawDescData = file_delivery_proto_rawDesc
)

func file_delivery_proto_rawDescGZIP() []byte {
	file_delivery_proto_rawDescOnce.Do(func() {
		file_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_delivery_proto_rawDescData)
	})
	return file_delivery_proto_rawDescData
}

var file_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_delivery_proto_goTypes = []interface{}{
	(*DeliveryResponse)(nil),       // 0: planetexpress.DeliveryResponse
	(*DeliveryRequest)(nil),        // 1: planetexpress.DeliveryRequest
	(*GetDeliveryResponse)(nil),    // 2: planetexpress.GetDeliveryResponse
	(*GetDeliveryRequest)(nil),     // 3: planetexpress.GetDeliveryRequest
	(*ListDeliveriesResponse)(nil), // 4: planetexpress.ListDeliveriesResponse
}
var file_delivery_proto_depIdxs = []int32{
	0, // 0: planetexpress.GetDeliveryResponse.delivery:type_name -> planetexpress.DeliveryResponse
	1, // 1: planetexpress.GetDeliveryRequest.delivery:type_name -> planetexpress.DeliveryRequest
	0, // 2: planetexpress.ListDeliveriesResponse.list:type_name -> planetexpress.DeliveryResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_delivery_proto_init() }
func file_delivery_proto_init() {
	if File_delivery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_delivery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryResponse); i {
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
		file_delivery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryRequest); i {
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
		file_delivery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeliveryResponse); i {
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
		file_delivery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeliveryRequest); i {
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
		file_delivery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeliveriesResponse); i {
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
			RawDescriptor: file_delivery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_delivery_proto_goTypes,
		DependencyIndexes: file_delivery_proto_depIdxs,
		MessageInfos:      file_delivery_proto_msgTypes,
	}.Build()
	File_delivery_proto = out.File
	file_delivery_proto_rawDesc = nil
	file_delivery_proto_goTypes = nil
	file_delivery_proto_depIdxs = nil
}