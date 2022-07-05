// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: emails.proto

package grpcControllers

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

type OrderStatus int32

const (
	OrderStatus_ACCEPTED OrderStatus = 0
	OrderStatus_REJECTED OrderStatus = 1
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ACCEPTED",
		1: "REJECTED",
	}
	OrderStatus_value = map[string]int32{
		"ACCEPTED": 0,
		"REJECTED": 1,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_emails_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_emails_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_emails_proto_rawDescGZIP(), []int{0}
}

type OrderEmailParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          OrderStatus `protobuf:"varint,1,opt,name=Status,proto3,enum=emails.OrderStatus" json:"Status,omitempty"`
	Message         string      `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	CustomerEmail   string      `protobuf:"bytes,3,opt,name=CustomerEmail,proto3" json:"CustomerEmail,omitempty"`
	BackgroundImage []byte      `protobuf:"bytes,4,opt,name=BackgroundImage,proto3,oneof" json:"BackgroundImage,omitempty"`
}

func (x *OrderEmailParams) Reset() {
	*x = OrderEmailParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderEmailParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderEmailParams) ProtoMessage() {}

func (x *OrderEmailParams) ProtoReflect() protoreflect.Message {
	mi := &file_emails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderEmailParams.ProtoReflect.Descriptor instead.
func (*OrderEmailParams) Descriptor() ([]byte, []int) {
	return file_emails_proto_rawDescGZIP(), []int{0}
}

func (x *OrderEmailParams) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ACCEPTED
}

func (x *OrderEmailParams) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderEmailParams) GetCustomerEmail() string {
	if x != nil {
		return x.CustomerEmail
	}
	return ""
}

func (x *OrderEmailParams) GetBackgroundImage() []byte {
	if x != nil {
		return x.BackgroundImage
	}
	return nil
}

type DefaultEmailParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params for sending Default Email not related to any theme
	EmailMessage    string `protobuf:"bytes,1,opt,name=EmailMessage,proto3" json:"EmailMessage,omitempty"`
	CustomerEmail   string `protobuf:"bytes,2,opt,name=customerEmail,proto3" json:"customerEmail,omitempty"`
	BackgroundImage []byte `protobuf:"bytes,3,opt,name=BackgroundImage,proto3,oneof" json:"BackgroundImage,omitempty"`
}

func (x *DefaultEmailParams) Reset() {
	*x = DefaultEmailParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultEmailParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultEmailParams) ProtoMessage() {}

func (x *DefaultEmailParams) ProtoReflect() protoreflect.Message {
	mi := &file_emails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultEmailParams.ProtoReflect.Descriptor instead.
func (*DefaultEmailParams) Descriptor() ([]byte, []int) {
	return file_emails_proto_rawDescGZIP(), []int{1}
}

func (x *DefaultEmailParams) GetEmailMessage() string {
	if x != nil {
		return x.EmailMessage
	}
	return ""
}

func (x *DefaultEmailParams) GetCustomerEmail() string {
	if x != nil {
		return x.CustomerEmail
	}
	return ""
}

func (x *DefaultEmailParams) GetBackgroundImage() []byte {
	if x != nil {
		return x.BackgroundImage
	}
	return nil
}

type EmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delivered bool `protobuf:"varint,1,opt,name=Delivered,proto3" json:"Delivered,omitempty"`
}

func (x *EmailResponse) Reset() {
	*x = EmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailResponse) ProtoMessage() {}

func (x *EmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailResponse.ProtoReflect.Descriptor instead.
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return file_emails_proto_rawDescGZIP(), []int{2}
}

func (x *EmailResponse) GetDelivered() bool {
	if x != nil {
		return x.Delivered
	}
	return false
}

var File_emails_proto protoreflect.FileDescriptor

var file_emails_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x42, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x12,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x0f,
	0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x2d, 0x0a, 0x0d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x2a, 0x29,
	0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x32, 0x90, 0x01, 0x0a, 0x0b, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x53, 0x65, 0x6e,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10,
	0x67, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emails_proto_rawDescOnce sync.Once
	file_emails_proto_rawDescData = file_emails_proto_rawDesc
)

func file_emails_proto_rawDescGZIP() []byte {
	file_emails_proto_rawDescOnce.Do(func() {
		file_emails_proto_rawDescData = protoimpl.X.CompressGZIP(file_emails_proto_rawDescData)
	})
	return file_emails_proto_rawDescData
}

var file_emails_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_emails_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_emails_proto_goTypes = []interface{}{
	(OrderStatus)(0),           // 0: emails.OrderStatus
	(*OrderEmailParams)(nil),   // 1: emails.OrderEmailParams
	(*DefaultEmailParams)(nil), // 2: emails.DefaultEmailParams
	(*EmailResponse)(nil),      // 3: emails.EmailResponse
}
var file_emails_proto_depIdxs = []int32{
	0, // 0: emails.OrderEmailParams.Status:type_name -> emails.OrderStatus
	2, // 1: emails.EmailSender.SendEmail:input_type -> emails.DefaultEmailParams
	1, // 2: emails.EmailSender.SendOrderEmail:input_type -> emails.OrderEmailParams
	3, // 3: emails.EmailSender.SendEmail:output_type -> emails.EmailResponse
	3, // 4: emails.EmailSender.SendOrderEmail:output_type -> emails.EmailResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_emails_proto_init() }
func file_emails_proto_init() {
	if File_emails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderEmailParams); i {
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
		file_emails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultEmailParams); i {
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
		file_emails_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailResponse); i {
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
	file_emails_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_emails_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_emails_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_emails_proto_goTypes,
		DependencyIndexes: file_emails_proto_depIdxs,
		EnumInfos:         file_emails_proto_enumTypes,
		MessageInfos:      file_emails_proto_msgTypes,
	}.Build()
	File_emails_proto = out.File
	file_emails_proto_rawDesc = nil
	file_emails_proto_goTypes = nil
	file_emails_proto_depIdxs = nil
}