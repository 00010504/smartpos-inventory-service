// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: inventory_main.proto

package inventory_service

import (
	common "genproto/common"
	order_service "genproto/order_service"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_inventory_main_proto protoreflect.FileDescriptor

var file_inventory_main_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf2, 0x0b, 0x0a, 0x10,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x31, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x37, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x47, 0x0a, 0x17, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x1d, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x19,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x44, 0x12, 0x3f, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x44, 0x12, 0x4c, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x22, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x69, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49,
	0x44, 0x12, 0x2e, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49,
	0x44, 0x12, 0x31, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x41,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44,
	0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_inventory_main_proto_goTypes = []interface{}{
	(*CreateImportRequest)(nil),                  // 0: CreateImportRequest
	(*order_service.CreateOrderCopyRequest)(nil), // 1: CreateOrderCopyRequest
	(*FinishImportReq)(nil),                      // 2: FinishImportReq
	(*common.SearchRequest)(nil),                 // 3: SearchRequest
	(*common.RequestID)(nil),                     // 4: RequestID
	(*CreateSupplierRequest)(nil),                // 5: CreateSupplierRequest
	(*UpdateSupplierRequest)(nil),                // 6: UpdateSupplierRequest
	(*CreateSupplierOrderRequest)(nil),           // 7: CreateSupplierOrderRequest
	(*CreateSupplierOrderItemRequest)(nil),       // 8: CreateSupplierOrderItemRequest
	(*GetAllSupplierOrderItemsRequest)(nil),      // 9: GetAllSupplierOrderItemsRequest
	(*FinishSupplierOrderRequest)(nil),           // 10: FinishSupplierOrderRequest
	(*UpdateSupplierOrderRecivedRequest)(nil),    // 11: UpdateSupplierOrderRecivedRequest
	(*GetProductHistoryReq)(nil),                 // 12: GetProductHistoryReq
	(*CreateTransferRequest)(nil),                // 13: CreateTransferRequest
	(*GetAllTransferRequest)(nil),                // 14: GetAllTransferRequest
	(*GetAllTransferItemsRequest)(nil),           // 15: GetAllTransferItemsRequest
	(*AddItemToTransferRequest)(nil),             // 16: AddItemToTransferRequest
	(*common.ResponseID)(nil),                    // 17: ResponseID
	(*GetAllImportRes)(nil),                      // 18: GetAllImportRes
	(*ImportById)(nil),                           // 19: ImportById
	(*GetAllSuppliersResponse)(nil),              // 20: GetAllSuppliersResponse
	(*GetSupplierByIdResponse)(nil),              // 21: GetSupplierByIdResponse
	(*GetAllSupplierOrderResponse)(nil),          // 22: GetAllSupplierOrderResponse
	(*GetAllSupplierOrderItemsResponse)(nil),     // 23: GetAllSupplierOrderItemsResponse
	(*GetSupplierOrderByIdResponse)(nil),         // 24: GetSupplierOrderByIdResponse
	(*common.Empty)(nil),                         // 25: Empty
	(*GetProductHistoryRes)(nil),                 // 26: GetProductHistoryRes
	(*GetAllTransferResponse)(nil),               // 27: GetAllTransferResponse
	(*GetAllTransferItemsResponse)(nil),          // 28: GetAllTransferItemsResponse
}
var file_inventory_main_proto_depIdxs = []int32{
	0,  // 0: InventoryService.CreateImport:input_type -> CreateImportRequest
	1,  // 1: InventoryService.CreateOrder:input_type -> CreateOrderCopyRequest
	2,  // 2: InventoryService.FinishImport:input_type -> FinishImportReq
	3,  // 3: InventoryService.GetAllImport:input_type -> SearchRequest
	4,  // 4: InventoryService.GetImportByID:input_type -> RequestID
	5,  // 5: InventoryService.CreateSupplier:input_type -> CreateSupplierRequest
	3,  // 6: InventoryService.GetAllSupplier:input_type -> SearchRequest
	6,  // 7: InventoryService.UpdateSupplier:input_type -> UpdateSupplierRequest
	4,  // 8: InventoryService.GetSupplierById:input_type -> RequestID
	4,  // 9: InventoryService.Delete:input_type -> RequestID
	7,  // 10: InventoryService.CreateSupplierOrder:input_type -> CreateSupplierOrderRequest
	8,  // 11: InventoryService.UpsertSupplierOrderItem:input_type -> CreateSupplierOrderItemRequest
	3,  // 12: InventoryService.GetAllSupplierOrder:input_type -> SearchRequest
	9,  // 13: InventoryService.GetAllSupplierOrderItems:input_type -> GetAllSupplierOrderItemsRequest
	4,  // 14: InventoryService.GetSupplierOrderById:input_type -> RequestID
	4,  // 15: InventoryService.UpdateSupplierOrderStatus:input_type -> RequestID
	10, // 16: InventoryService.FinishSupplierOrder:input_type -> FinishSupplierOrderRequest
	11, // 17: InventoryService.UpdateSupplierOrderAmount:input_type -> UpdateSupplierOrderRecivedRequest
	4,  // 18: InventoryService.DeleteSupplierOrder:input_type -> RequestID
	4,  // 19: InventoryService.DeleteSupplierOrderItemById:input_type -> RequestID
	12, // 20: InventoryService.GetProductHistory:input_type -> GetProductHistoryReq
	13, // 21: InventoryService.CreateTransfer:input_type -> CreateTransferRequest
	14, // 22: InventoryService.GetAllTransfer:input_type -> GetAllTransferRequest
	15, // 23: InventoryService.GetAllTransferItems:input_type -> GetAllTransferItemsRequest
	16, // 24: InventoryService.AddItemToTransfer:input_type -> AddItemToTransferRequest
	17, // 25: InventoryService.CreateImport:output_type -> ResponseID
	17, // 26: InventoryService.CreateOrder:output_type -> ResponseID
	17, // 27: InventoryService.FinishImport:output_type -> ResponseID
	18, // 28: InventoryService.GetAllImport:output_type -> GetAllImportRes
	19, // 29: InventoryService.GetImportByID:output_type -> ImportById
	17, // 30: InventoryService.CreateSupplier:output_type -> ResponseID
	20, // 31: InventoryService.GetAllSupplier:output_type -> GetAllSuppliersResponse
	17, // 32: InventoryService.UpdateSupplier:output_type -> ResponseID
	21, // 33: InventoryService.GetSupplierById:output_type -> GetSupplierByIdResponse
	17, // 34: InventoryService.Delete:output_type -> ResponseID
	17, // 35: InventoryService.CreateSupplierOrder:output_type -> ResponseID
	17, // 36: InventoryService.UpsertSupplierOrderItem:output_type -> ResponseID
	22, // 37: InventoryService.GetAllSupplierOrder:output_type -> GetAllSupplierOrderResponse
	23, // 38: InventoryService.GetAllSupplierOrderItems:output_type -> GetAllSupplierOrderItemsResponse
	24, // 39: InventoryService.GetSupplierOrderById:output_type -> GetSupplierOrderByIdResponse
	17, // 40: InventoryService.UpdateSupplierOrderStatus:output_type -> ResponseID
	17, // 41: InventoryService.FinishSupplierOrder:output_type -> ResponseID
	17, // 42: InventoryService.UpdateSupplierOrderAmount:output_type -> ResponseID
	17, // 43: InventoryService.DeleteSupplierOrder:output_type -> ResponseID
	25, // 44: InventoryService.DeleteSupplierOrderItemById:output_type -> Empty
	26, // 45: InventoryService.GetProductHistory:output_type -> GetProductHistoryRes
	17, // 46: InventoryService.CreateTransfer:output_type -> ResponseID
	27, // 47: InventoryService.GetAllTransfer:output_type -> GetAllTransferResponse
	28, // 48: InventoryService.GetAllTransferItems:output_type -> GetAllTransferItemsResponse
	17, // 49: InventoryService.AddItemToTransfer:output_type -> ResponseID
	25, // [25:50] is the sub-list for method output_type
	0,  // [0:25] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_inventory_main_proto_init() }
func file_inventory_main_proto_init() {
	if File_inventory_main_proto != nil {
		return
	}
	file_import_proto_init()
	file_supplier_proto_init()
	file_supplier_order_proto_init()
	file_transfer_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventory_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_main_proto_goTypes,
		DependencyIndexes: file_inventory_main_proto_depIdxs,
	}.Build()
	File_inventory_main_proto = out.File
	file_inventory_main_proto_rawDesc = nil
	file_inventory_main_proto_goTypes = nil
	file_inventory_main_proto_depIdxs = nil
}
