// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: label.proto

package catalog_service

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

type CreateLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request    *common.Request          `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Parameters *LabelParametrs          `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	Content    map[string]*LabelContent `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateLabelRequest) Reset() {
	*x = CreateLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLabelRequest) ProtoMessage() {}

func (x *CreateLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLabelRequest.ProtoReflect.Descriptor instead.
func (*CreateLabelRequest) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLabelRequest) GetRequest() *common.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *CreateLabelRequest) GetParameters() *LabelParametrs {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *CreateLabelRequest) GetContent() map[string]*LabelContent {
	if x != nil {
		return x.Content
	}
	return nil
}

type LabelParametrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Width  int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *LabelParametrs) Reset() {
	*x = LabelParametrs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelParametrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelParametrs) ProtoMessage() {}

func (x *LabelParametrs) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelParametrs.ProtoReflect.Descriptor instead.
func (*LabelParametrs) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{1}
}

func (x *LabelParametrs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelParametrs) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *LabelParametrs) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type TextFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FontFamily string `protobuf:"bytes,1,opt,name=font_family,json=fontFamily,proto3" json:"font_family,omitempty"`
	FontStyle  string `protobuf:"bytes,2,opt,name=font_style,json=fontStyle,proto3" json:"font_style,omitempty"` // [itelic, normal]
	FontSize   int32  `protobuf:"varint,3,opt,name=font_size,json=fontSize,proto3" json:"font_size,omitempty"`
	FontWeight int32  `protobuf:"varint,4,opt,name=font_weight,json=fontWeight,proto3" json:"font_weight,omitempty"`
	TextAlign  string `protobuf:"bytes,5,opt,name=text_align,json=textAlign,proto3" json:"text_align,omitempty"` // [left, right, center]
}

func (x *TextFormat) Reset() {
	*x = TextFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextFormat) ProtoMessage() {}

func (x *TextFormat) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextFormat.ProtoReflect.Descriptor instead.
func (*TextFormat) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{2}
}

func (x *TextFormat) GetFontFamily() string {
	if x != nil {
		return x.FontFamily
	}
	return ""
}

func (x *TextFormat) GetFontStyle() string {
	if x != nil {
		return x.FontStyle
	}
	return ""
}

func (x *TextFormat) GetFontSize() int32 {
	if x != nil {
		return x.FontSize
	}
	return 0
}

func (x *TextFormat) GetFontWeight() int32 {
	if x != nil {
		return x.FontWeight
	}
	return 0
}

func (x *TextFormat) GetTextAlign() string {
	if x != nil {
		return x.TextAlign
	}
	return ""
}

type LabelPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *LabelPosition) Reset() {
	*x = LabelPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelPosition) ProtoMessage() {}

func (x *LabelPosition) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelPosition.ProtoReflect.Descriptor instead.
func (*LabelPosition) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{3}
}

func (x *LabelPosition) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *LabelPosition) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type LabelContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Position      *LabelPosition `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Width         int32          `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height        int32          `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Type          string         `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"` // enum  [image, text, barcode]
	ProductImage  string         `protobuf:"bytes,6,opt,name=product_image,json=productImage,proto3" json:"product_image,omitempty"`
	FieldName     string         `protobuf:"bytes,7,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Format        *TextFormat    `protobuf:"bytes,8,opt,name=format,proto3" json:"format,omitempty"`
	IsCustomField bool           `protobuf:"varint,9,opt,name=is_custom_field,json=isCustomField,proto3" json:"is_custom_field,omitempty"`
}

func (x *LabelContent) Reset() {
	*x = LabelContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelContent) ProtoMessage() {}

func (x *LabelContent) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelContent.ProtoReflect.Descriptor instead.
func (*LabelContent) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{4}
}

func (x *LabelContent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LabelContent) GetPosition() *LabelPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *LabelContent) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *LabelContent) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *LabelContent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LabelContent) GetProductImage() string {
	if x != nil {
		return x.ProductImage
	}
	return ""
}

func (x *LabelContent) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *LabelContent) GetFormat() *TextFormat {
	if x != nil {
		return x.Format
	}
	return nil
}

func (x *LabelContent) GetIsCustomField() bool {
	if x != nil {
		return x.IsCustomField
	}
	return false
}

type GetProductFieldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsCustomField bool   `protobuf:"varint,2,opt,name=is_custom_field,json=isCustomField,proto3" json:"is_custom_field,omitempty"`
}

func (x *GetProductFieldResponse) Reset() {
	*x = GetProductFieldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductFieldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductFieldResponse) ProtoMessage() {}

func (x *GetProductFieldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductFieldResponse.ProtoReflect.Descriptor instead.
func (*GetProductFieldResponse) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductFieldResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductFieldResponse) GetIsCustomField() bool {
	if x != nil {
		return x.IsCustomField
	}
	return false
}

type GetProductFieldsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*GetProductFieldResponse `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *GetProductFieldsResponse) Reset() {
	*x = GetProductFieldsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductFieldsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductFieldsResponse) ProtoMessage() {}

func (x *GetProductFieldsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductFieldsResponse.ProtoReflect.Descriptor instead.
func (*GetProductFieldsResponse) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{6}
}

func (x *GetProductFieldsResponse) GetFields() []*GetProductFieldResponse {
	if x != nil {
		return x.Fields
	}
	return nil
}

type GetLabelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parameters *LabelParametrs          `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	Content    map[string]*LabelContent `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map<string, string> layers = 4;
}

func (x *GetLabelResponse) Reset() {
	*x = GetLabelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLabelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLabelResponse) ProtoMessage() {}

func (x *GetLabelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLabelResponse.ProtoReflect.Descriptor instead.
func (*GetLabelResponse) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{7}
}

func (x *GetLabelResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLabelResponse) GetParameters() *LabelParametrs {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *GetLabelResponse) GetContent() map[string]*LabelContent {
	if x != nil {
		return x.Content
	}
	return nil
}

type UpdateLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request    *common.Request          `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Id         string                   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Parameters *LabelParametrs          `protobuf:"bytes,3,opt,name=parameters,proto3" json:"parameters,omitempty"`
	Content    map[string]*LabelContent `protobuf:"bytes,4,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateLabelRequest) Reset() {
	*x = UpdateLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLabelRequest) ProtoMessage() {}

func (x *UpdateLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLabelRequest.ProtoReflect.Descriptor instead.
func (*UpdateLabelRequest) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateLabelRequest) GetRequest() *common.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *UpdateLabelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateLabelRequest) GetParameters() *LabelParametrs {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *UpdateLabelRequest) GetContent() map[string]*LabelContent {
	if x != nil {
		return x.Content
	}
	return nil
}

type GetAllLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*GetLabelResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetAllLabelsResponse) Reset() {
	*x = GetAllLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLabelsResponse) ProtoMessage() {}

func (x *GetAllLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLabelsResponse.ProtoReflect.Descriptor instead.
func (*GetAllLabelsResponse) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllLabelsResponse) GetData() []*GetLabelResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetAllLabelsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_label_proto protoreflect.FileDescriptor

var file_label_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x49, 0x0a, 0x0c, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x0e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x0a, 0x54,
	0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6e,
	0x74, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x6f, 0x6e, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f,
	0x6e, 0x74, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x6f, 0x6e, 0x74, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6e,
	0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x6f,
	0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6e, 0x74, 0x5f, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x6f, 0x6e,
	0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x61, 0x6c, 0x69, 0x67, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x78,
	0x74, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x22, 0x2b, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x22, 0x9d, 0x02, 0x0a, 0x0c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x69,
	0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x22, 0x55, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x4c, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xd8, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x38,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x49, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x80, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f,
	0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x49, 0x0a, 0x0c, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x53, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x1a, 0x5a, 0x18, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_label_proto_rawDescOnce sync.Once
	file_label_proto_rawDescData = file_label_proto_rawDesc
)

func file_label_proto_rawDescGZIP() []byte {
	file_label_proto_rawDescOnce.Do(func() {
		file_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_label_proto_rawDescData)
	})
	return file_label_proto_rawDescData
}

var file_label_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_label_proto_goTypes = []interface{}{
	(*CreateLabelRequest)(nil),       // 0: CreateLabelRequest
	(*LabelParametrs)(nil),           // 1: LabelParametrs
	(*TextFormat)(nil),               // 2: TextFormat
	(*LabelPosition)(nil),            // 3: LabelPosition
	(*LabelContent)(nil),             // 4: LabelContent
	(*GetProductFieldResponse)(nil),  // 5: GetProductFieldResponse
	(*GetProductFieldsResponse)(nil), // 6: GetProductFieldsResponse
	(*GetLabelResponse)(nil),         // 7: GetLabelResponse
	(*UpdateLabelRequest)(nil),       // 8: UpdateLabelRequest
	(*GetAllLabelsResponse)(nil),     // 9: GetAllLabelsResponse
	nil,                              // 10: CreateLabelRequest.ContentEntry
	nil,                              // 11: GetLabelResponse.ContentEntry
	nil,                              // 12: UpdateLabelRequest.ContentEntry
	(*common.Request)(nil),           // 13: Request
}
var file_label_proto_depIdxs = []int32{
	13, // 0: CreateLabelRequest.request:type_name -> Request
	1,  // 1: CreateLabelRequest.parameters:type_name -> LabelParametrs
	10, // 2: CreateLabelRequest.content:type_name -> CreateLabelRequest.ContentEntry
	3,  // 3: LabelContent.position:type_name -> LabelPosition
	2,  // 4: LabelContent.format:type_name -> TextFormat
	5,  // 5: GetProductFieldsResponse.fields:type_name -> GetProductFieldResponse
	1,  // 6: GetLabelResponse.parameters:type_name -> LabelParametrs
	11, // 7: GetLabelResponse.content:type_name -> GetLabelResponse.ContentEntry
	13, // 8: UpdateLabelRequest.request:type_name -> Request
	1,  // 9: UpdateLabelRequest.parameters:type_name -> LabelParametrs
	12, // 10: UpdateLabelRequest.content:type_name -> UpdateLabelRequest.ContentEntry
	7,  // 11: GetAllLabelsResponse.data:type_name -> GetLabelResponse
	4,  // 12: CreateLabelRequest.ContentEntry.value:type_name -> LabelContent
	4,  // 13: GetLabelResponse.ContentEntry.value:type_name -> LabelContent
	4,  // 14: UpdateLabelRequest.ContentEntry.value:type_name -> LabelContent
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_label_proto_init() }
func file_label_proto_init() {
	if File_label_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLabelRequest); i {
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
		file_label_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelParametrs); i {
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
		file_label_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextFormat); i {
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
		file_label_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelPosition); i {
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
		file_label_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelContent); i {
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
		file_label_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductFieldResponse); i {
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
		file_label_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductFieldsResponse); i {
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
		file_label_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLabelResponse); i {
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
		file_label_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLabelRequest); i {
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
		file_label_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllLabelsResponse); i {
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
			RawDescriptor: file_label_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_label_proto_goTypes,
		DependencyIndexes: file_label_proto_depIdxs,
		MessageInfos:      file_label_proto_msgTypes,
	}.Build()
	File_label_proto = out.File
	file_label_proto_rawDesc = nil
	file_label_proto_goTypes = nil
	file_label_proto_depIdxs = nil
}