// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: common_product.proto

package common

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

type CreateProductCopyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sku                   string                        `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Name                  string                        `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Image                 string                        `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	BrandId               string                        `protobuf:"bytes,6,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	MxikCode              string                        `protobuf:"bytes,7,opt,name=mxik_code,json=mxikCode,proto3" json:"mxik_code,omitempty"`
	ParentId              string                        `protobuf:"bytes,8,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Description           string                        `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	ProductTypeId         string                        `protobuf:"bytes,10,opt,name=product_type_id,json=productTypeId,proto3" json:"product_type_id,omitempty"`
	MeasurementUnitId     string                        `protobuf:"bytes,11,opt,name=measurement_unit_id,json=measurementUnitId,proto3" json:"measurement_unit_id,omitempty"`
	IsMarking             bool                          `protobuf:"varint,12,opt,name=is_marking,json=isMarking,proto3" json:"is_marking,omitempty"`
	Barcode               []string                      `protobuf:"bytes,13,rep,name=barcode,proto3" json:"barcode,omitempty"`
	ShopMeasurementValues []*CommonShopMeasurementValue `protobuf:"bytes,14,rep,name=shop_measurement_values,json=shopMeasurementValues,proto3" json:"shop_measurement_values,omitempty"`
	Request               *Request                      `protobuf:"bytes,15,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *CreateProductCopyRequest) Reset() {
	*x = CreateProductCopyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductCopyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductCopyRequest) ProtoMessage() {}

func (x *CreateProductCopyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductCopyRequest.ProtoReflect.Descriptor instead.
func (*CreateProductCopyRequest) Descriptor() ([]byte, []int) {
	return file_common_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductCopyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProductCopyRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *CreateProductCopyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductCopyRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateProductCopyRequest) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *CreateProductCopyRequest) GetMxikCode() string {
	if x != nil {
		return x.MxikCode
	}
	return ""
}

func (x *CreateProductCopyRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *CreateProductCopyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductCopyRequest) GetProductTypeId() string {
	if x != nil {
		return x.ProductTypeId
	}
	return ""
}

func (x *CreateProductCopyRequest) GetMeasurementUnitId() string {
	if x != nil {
		return x.MeasurementUnitId
	}
	return ""
}

func (x *CreateProductCopyRequest) GetIsMarking() bool {
	if x != nil {
		return x.IsMarking
	}
	return false
}

func (x *CreateProductCopyRequest) GetBarcode() []string {
	if x != nil {
		return x.Barcode
	}
	return nil
}

func (x *CreateProductCopyRequest) GetShopMeasurementValues() []*CommonShopMeasurementValue {
	if x != nil {
		return x.ShopMeasurementValues
	}
	return nil
}

func (x *CreateProductCopyRequest) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

type CommonShopMeasurementValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId         string  `protobuf:"bytes,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	IsAviable      bool    `protobuf:"varint,2,opt,name=is_aviable,json=isAviable,proto3" json:"is_aviable,omitempty"`
	InStock        float32 `protobuf:"fixed32,3,opt,name=in_stock,json=inStock,proto3" json:"in_stock,omitempty"`
	RetailPrice    float32 `protobuf:"fixed32,4,opt,name=retail_price,json=retailPrice,proto3" json:"retail_price,omitempty"`
	SupplyPrice    float32 `protobuf:"fixed32,5,opt,name=supply_price,json=supplyPrice,proto3" json:"supply_price,omitempty"`
	MinPrice       float32 `protobuf:"fixed32,6,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice       float32 `protobuf:"fixed32,7,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	WholeSalePrice float32 `protobuf:"fixed32,8,opt,name=whole_sale_price,json=wholeSalePrice,proto3" json:"whole_sale_price,omitempty"`
}

func (x *CommonShopMeasurementValue) Reset() {
	*x = CommonShopMeasurementValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonShopMeasurementValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonShopMeasurementValue) ProtoMessage() {}

func (x *CommonShopMeasurementValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonShopMeasurementValue.ProtoReflect.Descriptor instead.
func (*CommonShopMeasurementValue) Descriptor() ([]byte, []int) {
	return file_common_product_proto_rawDescGZIP(), []int{1}
}

func (x *CommonShopMeasurementValue) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *CommonShopMeasurementValue) GetIsAviable() bool {
	if x != nil {
		return x.IsAviable
	}
	return false
}

func (x *CommonShopMeasurementValue) GetInStock() float32 {
	if x != nil {
		return x.InStock
	}
	return 0
}

func (x *CommonShopMeasurementValue) GetRetailPrice() float32 {
	if x != nil {
		return x.RetailPrice
	}
	return 0
}

func (x *CommonShopMeasurementValue) GetSupplyPrice() float32 {
	if x != nil {
		return x.SupplyPrice
	}
	return 0
}

func (x *CommonShopMeasurementValue) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *CommonShopMeasurementValue) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *CommonShopMeasurementValue) GetWholeSalePrice() float32 {
	if x != nil {
		return x.WholeSalePrice
	}
	return 0
}

var File_common_product_proto protoreflect.FileDescriptor

var file_common_product_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x03, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x78, 0x69,
	0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x78,
	0x69, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x13, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x53, 0x0a, 0x17, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x53, 0x68, 0x6f, 0x70, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x73, 0x68, 0x6f, 0x70, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x99, 0x02, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x68, 0x6f, 0x70, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x61, 0x76,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41,
	0x76, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x68, 0x6f, 0x6c, 0x65, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x77, 0x68, 0x6f,
	0x6c, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_product_proto_rawDescOnce sync.Once
	file_common_product_proto_rawDescData = file_common_product_proto_rawDesc
)

func file_common_product_proto_rawDescGZIP() []byte {
	file_common_product_proto_rawDescOnce.Do(func() {
		file_common_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_product_proto_rawDescData)
	})
	return file_common_product_proto_rawDescData
}

var file_common_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_product_proto_goTypes = []interface{}{
	(*CreateProductCopyRequest)(nil),   // 0: CreateProductCopyRequest
	(*CommonShopMeasurementValue)(nil), // 1: CommonShopMeasurementValue
	(*Request)(nil),                    // 2: Request
}
var file_common_product_proto_depIdxs = []int32{
	1, // 0: CreateProductCopyRequest.shop_measurement_values:type_name -> CommonShopMeasurementValue
	2, // 1: CreateProductCopyRequest.request:type_name -> Request
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_product_proto_init() }
func file_common_product_proto_init() {
	if File_common_product_proto != nil {
		return
	}
	file_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductCopyRequest); i {
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
		file_common_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonShopMeasurementValue); i {
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
			RawDescriptor: file_common_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_product_proto_goTypes,
		DependencyIndexes: file_common_product_proto_depIdxs,
		MessageInfos:      file_common_product_proto_msgTypes,
	}.Build()
	File_common_product_proto = out.File
	file_common_product_proto_rawDesc = nil
	file_common_product_proto_goTypes = nil
	file_common_product_proto_depIdxs = nil
}
