// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: validation_main.proto

package validation_service

import (
	common "genproto/common"
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

type GetAllImportItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit    int32           `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	Page     int32           `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	ImportId string          `protobuf:"bytes,1,opt,name=import_id,json=importId,proto3" json:"import_id,omitempty"`
	Search   string          `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	SortBy   string          `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	SortType string          `protobuf:"bytes,4,opt,name=sort_type,json=sortType,proto3" json:"sort_type,omitempty"`
	Request  *common.Request `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetAllImportItemsRequest) Reset() {
	*x = GetAllImportItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllImportItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllImportItemsRequest) ProtoMessage() {}

func (x *GetAllImportItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_validation_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllImportItemsRequest.ProtoReflect.Descriptor instead.
func (*GetAllImportItemsRequest) Descriptor() ([]byte, []int) {
	return file_validation_main_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllImportItemsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllImportItemsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllImportItemsRequest) GetImportId() string {
	if x != nil {
		return x.ImportId
	}
	return ""
}

func (x *GetAllImportItemsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllImportItemsRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *GetAllImportItemsRequest) GetSortType() string {
	if x != nil {
		return x.SortType
	}
	return ""
}

func (x *GetAllImportItemsRequest) GetRequest() *common.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

type ImmportItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ImportId    string  `protobuf:"bytes,2,opt,name=import_id,json=importId,proto3" json:"import_id,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ExternalId  string  `protobuf:"bytes,4,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Category    string  `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Barcode     string  `protobuf:"bytes,6,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Quantity    int32   `protobuf:"varint,7,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Sku         string  `protobuf:"bytes,8,opt,name=sku,proto3" json:"sku,omitempty"`
	SupplyPrice float32 `protobuf:"fixed32,9,opt,name=supply_price,json=supplyPrice,proto3" json:"supply_price,omitempty"`
	SalePrice   float32 `protobuf:"fixed32,10,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`
	BrandName   string  `protobuf:"bytes,11,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
}

func (x *ImmportItem) Reset() {
	*x = ImmportItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImmportItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImmportItem) ProtoMessage() {}

func (x *ImmportItem) ProtoReflect() protoreflect.Message {
	mi := &file_validation_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImmportItem.ProtoReflect.Descriptor instead.
func (*ImmportItem) Descriptor() ([]byte, []int) {
	return file_validation_main_proto_rawDescGZIP(), []int{1}
}

func (x *ImmportItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ImmportItem) GetImportId() string {
	if x != nil {
		return x.ImportId
	}
	return ""
}

func (x *ImmportItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImmportItem) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *ImmportItem) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ImmportItem) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *ImmportItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ImmportItem) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *ImmportItem) GetSupplyPrice() float32 {
	if x != nil {
		return x.SupplyPrice
	}
	return 0
}

func (x *ImmportItem) GetSalePrice() float32 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

func (x *ImmportItem) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

type ConfirmImportItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImportId string `protobuf:"bytes,1,opt,name=import_id,json=importId,proto3" json:"import_id,omitempty"`
}

func (x *ConfirmImportItemReq) Reset() {
	*x = ConfirmImportItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmImportItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmImportItemReq) ProtoMessage() {}

func (x *ConfirmImportItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_validation_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmImportItemReq.ProtoReflect.Descriptor instead.
func (*ConfirmImportItemReq) Descriptor() ([]byte, []int) {
	return file_validation_main_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmImportItemReq) GetImportId() string {
	if x != nil {
		return x.ImportId
	}
	return ""
}

type GetAllImportItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*ImmportItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalCount int32          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetAllImportItemsResponse) Reset() {
	*x = GetAllImportItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllImportItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllImportItemsResponse) ProtoMessage() {}

func (x *GetAllImportItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_validation_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllImportItemsResponse.ProtoReflect.Descriptor instead.
func (*GetAllImportItemsResponse) Descriptor() ([]byte, []int) {
	return file_validation_main_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllImportItemsResponse) GetData() []*ImmportItem {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetAllImportItemsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_validation_main_proto protoreflect.FileDescriptor

var file_validation_main_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd3, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x72, 0x74, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb4, 0x02, 0x0a, 0x0b, 0x49, 0x6d, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x14,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x22, 0x5e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49,
	0x6d, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0x98, 0x02, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x42, 0x79, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x4a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_validation_main_proto_rawDescOnce sync.Once
	file_validation_main_proto_rawDescData = file_validation_main_proto_rawDesc
)

func file_validation_main_proto_rawDescGZIP() []byte {
	file_validation_main_proto_rawDescOnce.Do(func() {
		file_validation_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_validation_main_proto_rawDescData)
	})
	return file_validation_main_proto_rawDescData
}

var file_validation_main_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_validation_main_proto_goTypes = []interface{}{
	(*GetAllImportItemsRequest)(nil),        // 0: GetAllImportItemsRequest
	(*ImmportItem)(nil),                     // 1: ImmportItem
	(*ConfirmImportItemReq)(nil),            // 2: ConfirmImportItemReq
	(*GetAllImportItemsResponse)(nil),       // 3: GetAllImportItemsResponse
	(*common.Request)(nil),                  // 4: Request
	(*CreateImportErrorReq)(nil),            // 5: CreateImportErrorReq
	(*CreateImportReq)(nil),                 // 6: CreateImportReq
	(*ConfirmImportReq)(nil),                // 7: ConfirmImportReq
	(*GetImportValidationLogsResponse)(nil), // 8: GetImportValidationLogsResponse
	(*CreateImportRes)(nil),                 // 9: CreateImportRes
}
var file_validation_main_proto_depIdxs = []int32{
	4, // 0: GetAllImportItemsRequest.request:type_name -> Request
	1, // 1: GetAllImportItemsResponse.data:type_name -> ImmportItem
	5, // 2: ValidationService.GetLogsByImportId:input_type -> CreateImportErrorReq
	6, // 3: ValidationService.CreateImport:input_type -> CreateImportReq
	7, // 4: ValidationService.ConfirmImport:input_type -> ConfirmImportReq
	0, // 5: ValidationService.GetAllImportItems:input_type -> GetAllImportItemsRequest
	8, // 6: ValidationService.GetLogsByImportId:output_type -> GetImportValidationLogsResponse
	9, // 7: ValidationService.CreateImport:output_type -> CreateImportRes
	7, // 8: ValidationService.ConfirmImport:output_type -> ConfirmImportReq
	3, // 9: ValidationService.GetAllImportItems:output_type -> GetAllImportItemsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_validation_main_proto_init() }
func file_validation_main_proto_init() {
	if File_validation_main_proto != nil {
		return
	}
	file_validation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_validation_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllImportItemsRequest); i {
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
		file_validation_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImmportItem); i {
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
		file_validation_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmImportItemReq); i {
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
		file_validation_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllImportItemsResponse); i {
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
			RawDescriptor: file_validation_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_validation_main_proto_goTypes,
		DependencyIndexes: file_validation_main_proto_depIdxs,
		MessageInfos:      file_validation_main_proto_msgTypes,
	}.Build()
	File_validation_main_proto = out.File
	file_validation_main_proto_rawDesc = nil
	file_validation_main_proto_goTypes = nil
	file_validation_main_proto_depIdxs = nil
}
