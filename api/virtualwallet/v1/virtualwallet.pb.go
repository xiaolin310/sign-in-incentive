// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: api/virtualwallet/v1/virtualwallet.proto

package v1

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

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User int64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_api_virtualwallet_v1_virtualwallet_proto_rawDescGZIP(), []int{0}
}

func (x *GetBalanceRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

type GetBalanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance float64 `protobuf:"fixed64,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetBalanceReply) Reset() {
	*x = GetBalanceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceReply) ProtoMessage() {}

func (x *GetBalanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceReply.ProtoReflect.Descriptor instead.
func (*GetBalanceReply) Descriptor() ([]byte, []int) {
	return file_api_virtualwallet_v1_virtualwallet_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceReply) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type DebitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   int64   `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DebitRequest) Reset() {
	*x = DebitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebitRequest) ProtoMessage() {}

func (x *DebitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebitRequest.ProtoReflect.Descriptor instead.
func (*DebitRequest) Descriptor() ([]byte, []int) {
	return file_api_virtualwallet_v1_virtualwallet_proto_rawDescGZIP(), []int{2}
}

func (x *DebitRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *DebitRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DebitReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *DebitReply) Reset() {
	*x = DebitReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebitReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebitReply) ProtoMessage() {}

func (x *DebitReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebitReply.ProtoReflect.Descriptor instead.
func (*DebitReply) Descriptor() ([]byte, []int) {
	return file_api_virtualwallet_v1_virtualwallet_proto_rawDescGZIP(), []int{3}
}

func (x *DebitReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type CreditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   int64   `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreditRequest) Reset() {
	*x = CreditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditRequest) ProtoMessage() {}

func (x *CreditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditRequest.ProtoReflect.Descriptor instead.
func (*CreditRequest) Descriptor() ([]byte, []int) {
	return file_api_virtualwallet_v1_virtualwallet_proto_rawDescGZIP(), []int{4}
}

func (x *CreditRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CreditRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreditReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CreditReply) Reset() {
	*x = CreditReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditReply) ProtoMessage() {}

func (x *CreditReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditReply.ProtoReflect.Descriptor instead.
func (*CreditReply) Descriptor() ([]byte, []int) {
	return file_api_virtualwallet_v1_virtualwallet_proto_rawDescGZIP(), []int{5}
}

func (x *CreditReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_api_virtualwallet_v1_virtualwallet_proto protoreflect.FileDescriptor

var file_api_virtualwallet_v1_virtualwallet_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x27, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x67, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x3a, 0x0a, 0x0c, 0x64, 0x65, 0x62, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x20, 0x0a, 0x0a, 0x64, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x3b, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x21, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x8e, 0x02, 0x0a, 0x0d, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x5c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x65, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x69, 0x74, 0x12, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x23, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x39, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x1f, 0x73, 0x69, 0x67, 0x6e, 0x2d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_virtualwallet_v1_virtualwallet_proto_rawDescOnce sync.Once
	file_api_virtualwallet_v1_virtualwallet_proto_rawDescData = file_api_virtualwallet_v1_virtualwallet_proto_rawDesc
)

func file_api_virtualwallet_v1_virtualwallet_proto_rawDescGZIP() []byte {
	file_api_virtualwallet_v1_virtualwallet_proto_rawDescOnce.Do(func() {
		file_api_virtualwallet_v1_virtualwallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_virtualwallet_v1_virtualwallet_proto_rawDescData)
	})
	return file_api_virtualwallet_v1_virtualwallet_proto_rawDescData
}

var file_api_virtualwallet_v1_virtualwallet_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_virtualwallet_v1_virtualwallet_proto_goTypes = []interface{}{
	(*GetBalanceRequest)(nil), // 0: api.virtualwallet.v1.getBalanceRequest
	(*GetBalanceReply)(nil),   // 1: api.virtualwallet.v1.getBalanceReply
	(*DebitRequest)(nil),      // 2: api.virtualwallet.v1.debitRequest
	(*DebitReply)(nil),        // 3: api.virtualwallet.v1.debitReply
	(*CreditRequest)(nil),     // 4: api.virtualwallet.v1.creditRequest
	(*CreditReply)(nil),       // 5: api.virtualwallet.v1.creditReply
}
var file_api_virtualwallet_v1_virtualwallet_proto_depIdxs = []int32{
	0, // 0: api.virtualwallet.v1.VirtualWallet.GetBalance:input_type -> api.virtualwallet.v1.getBalanceRequest
	2, // 1: api.virtualwallet.v1.VirtualWallet.Debit:input_type -> api.virtualwallet.v1.debitRequest
	4, // 2: api.virtualwallet.v1.VirtualWallet.Credit:input_type -> api.virtualwallet.v1.creditRequest
	1, // 3: api.virtualwallet.v1.VirtualWallet.GetBalance:output_type -> api.virtualwallet.v1.getBalanceReply
	3, // 4: api.virtualwallet.v1.VirtualWallet.Debit:output_type -> api.virtualwallet.v1.debitReply
	5, // 5: api.virtualwallet.v1.VirtualWallet.Credit:output_type -> api.virtualwallet.v1.creditReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_virtualwallet_v1_virtualwallet_proto_init() }
func file_api_virtualwallet_v1_virtualwallet_proto_init() {
	if File_api_virtualwallet_v1_virtualwallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceReply); i {
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
		file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebitRequest); i {
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
		file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebitReply); i {
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
		file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditRequest); i {
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
		file_api_virtualwallet_v1_virtualwallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditReply); i {
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
			RawDescriptor: file_api_virtualwallet_v1_virtualwallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_virtualwallet_v1_virtualwallet_proto_goTypes,
		DependencyIndexes: file_api_virtualwallet_v1_virtualwallet_proto_depIdxs,
		MessageInfos:      file_api_virtualwallet_v1_virtualwallet_proto_msgTypes,
	}.Build()
	File_api_virtualwallet_v1_virtualwallet_proto = out.File
	file_api_virtualwallet_v1_virtualwallet_proto_rawDesc = nil
	file_api_virtualwallet_v1_virtualwallet_proto_goTypes = nil
	file_api_virtualwallet_v1_virtualwallet_proto_depIdxs = nil
}
