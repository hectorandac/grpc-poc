// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protos/billing.proto

package protos

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Recipient   string  `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Body        string  `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Sender      string  `protobuf:"bytes,4,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Type        string  `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	CreatedOn   int64   `protobuf:"varint,6,opt,name=CreatedOn,proto3" json:"CreatedOn,omitempty"`
	ReceivedOn  int64   `protobuf:"varint,7,opt,name=ReceivedOn,proto3" json:"ReceivedOn,omitempty"`
	ProcessedOn int64   `protobuf:"varint,8,opt,name=ProcessedOn,proto3" json:"ProcessedOn,omitempty"`
	Rate        float32 `protobuf:"fixed32,9,opt,name=Rate,proto3" json:"Rate,omitempty"`
	Refunded    bool    `protobuf:"varint,10,opt,name=Refunded,proto3" json:"Refunded,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protos_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_protos_billing_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Message) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetCreatedOn() int64 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

func (x *Message) GetReceivedOn() int64 {
	if x != nil {
		return x.ReceivedOn
	}
	return 0
}

func (x *Message) GetProcessedOn() int64 {
	if x != nil {
		return x.ProcessedOn
	}
	return 0
}

func (x *Message) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Message) GetRefunded() bool {
	if x != nil {
		return x.Refunded
	}
	return false
}

type BillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID    string   `protobuf:"bytes,1,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	Deliverable *Message `protobuf:"bytes,2,opt,name=deliverable,proto3" json:"deliverable,omitempty"`
}

func (x *BillRequest) Reset() {
	*x = BillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillRequest) ProtoMessage() {}

func (x *BillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillRequest.ProtoReflect.Descriptor instead.
func (*BillRequest) Descriptor() ([]byte, []int) {
	return file_protos_billing_proto_rawDescGZIP(), []int{1}
}

func (x *BillRequest) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *BillRequest) GetDeliverable() *Message {
	if x != nil {
		return x.Deliverable
	}
	return nil
}

type BillableEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillableEntityID string  `protobuf:"bytes,1,opt,name=billableEntityID,proto3" json:"billableEntityID,omitempty"`
	UserUUID         string  `protobuf:"bytes,2,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	Balance          float32 `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
	FullName         string  `protobuf:"bytes,4,opt,name=fullName,proto3" json:"fullName,omitempty"`
}

func (x *BillableEntity) Reset() {
	*x = BillableEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_billing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillableEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillableEntity) ProtoMessage() {}

func (x *BillableEntity) ProtoReflect() protoreflect.Message {
	mi := &file_protos_billing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillableEntity.ProtoReflect.Descriptor instead.
func (*BillableEntity) Descriptor() ([]byte, []int) {
	return file_protos_billing_proto_rawDescGZIP(), []int{2}
}

func (x *BillableEntity) GetBillableEntityID() string {
	if x != nil {
		return x.BillableEntityID
	}
	return ""
}

func (x *BillableEntity) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *BillableEntity) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *BillableEntity) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

type BillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Body    string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *BillResponse) Reset() {
	*x = BillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_billing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillResponse) ProtoMessage() {}

func (x *BillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_billing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillResponse.ProtoReflect.Descriptor instead.
func (*BillResponse) Descriptor() ([]byte, []int) {
	return file_protos_billing_proto_rawDescGZIP(), []int{3}
}

func (x *BillResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BillResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type RefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message        *Message        `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	BillableEntity *BillableEntity `protobuf:"bytes,2,opt,name=billableEntity,proto3" json:"billableEntity,omitempty"`
}

func (x *RefundRequest) Reset() {
	*x = RefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_billing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRequest) ProtoMessage() {}

func (x *RefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_billing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRequest.ProtoReflect.Descriptor instead.
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return file_protos_billing_proto_rawDescGZIP(), []int{4}
}

func (x *RefundRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *RefundRequest) GetBillableEntity() *BillableEntity {
	if x != nil {
		return x.BillableEntity
	}
	return nil
}

var File_protos_billing_proto protoreflect.FileDescriptor

var file_protos_billing_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x22,
	0x87, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x22, 0x5d, 0x0a, 0x0b, 0x42, 0x69, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x55, 0x55, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x70, 0x6f, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x42, 0x69, 0x6c,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x62,
	0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x42, 0x69, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x7c, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x70, 0x6f, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xda, 0x02, 0x0a, 0x07, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x12, 0x39, 0x0a, 0x08, 0x42, 0x69, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a,
	0x19, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x14,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x70, 0x6f, 0x63, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x2e, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x70, 0x6f, 0x63, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x4d, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x70, 0x6f, 0x63, 0x42, 0x07,
	0x47, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x63, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x61,
	0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x6f, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_billing_proto_rawDescOnce sync.Once
	file_protos_billing_proto_rawDescData = file_protos_billing_proto_rawDesc
)

func file_protos_billing_proto_rawDescGZIP() []byte {
	file_protos_billing_proto_rawDescOnce.Do(func() {
		file_protos_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_billing_proto_rawDescData)
	})
	return file_protos_billing_proto_rawDescData
}

var file_protos_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_billing_proto_goTypes = []interface{}{
	(*Message)(nil),        // 0: grpcpoc.Message
	(*BillRequest)(nil),    // 1: grpcpoc.BillRequest
	(*BillableEntity)(nil), // 2: grpcpoc.BillableEntity
	(*BillResponse)(nil),   // 3: grpcpoc.BillResponse
	(*RefundRequest)(nil),  // 4: grpcpoc.RefundRequest
}
var file_protos_billing_proto_depIdxs = []int32{
	0, // 0: grpcpoc.BillRequest.deliverable:type_name -> grpcpoc.Message
	0, // 1: grpcpoc.RefundRequest.message:type_name -> grpcpoc.Message
	2, // 2: grpcpoc.RefundRequest.billableEntity:type_name -> grpcpoc.BillableEntity
	1, // 3: grpcpoc.Billing.BillUser:input_type -> grpcpoc.BillRequest
	2, // 4: grpcpoc.Billing.UserBalance:input_type -> grpcpoc.BillableEntity
	2, // 5: grpcpoc.Billing.FindOrCreateBillingEntity:input_type -> grpcpoc.BillableEntity
	1, // 6: grpcpoc.Billing.CalculateMessageRate:input_type -> grpcpoc.BillRequest
	4, // 7: grpcpoc.Billing.RefundUser:input_type -> grpcpoc.RefundRequest
	3, // 8: grpcpoc.Billing.BillUser:output_type -> grpcpoc.BillResponse
	3, // 9: grpcpoc.Billing.UserBalance:output_type -> grpcpoc.BillResponse
	3, // 10: grpcpoc.Billing.FindOrCreateBillingEntity:output_type -> grpcpoc.BillResponse
	3, // 11: grpcpoc.Billing.CalculateMessageRate:output_type -> grpcpoc.BillResponse
	3, // 12: grpcpoc.Billing.RefundUser:output_type -> grpcpoc.BillResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_billing_proto_init() }
func file_protos_billing_proto_init() {
	if File_protos_billing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_protos_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillRequest); i {
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
		file_protos_billing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillableEntity); i {
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
		file_protos_billing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillResponse); i {
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
		file_protos_billing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundRequest); i {
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
			RawDescriptor: file_protos_billing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_billing_proto_goTypes,
		DependencyIndexes: file_protos_billing_proto_depIdxs,
		MessageInfos:      file_protos_billing_proto_msgTypes,
	}.Build()
	File_protos_billing_proto = out.File
	file_protos_billing_proto_rawDesc = nil
	file_protos_billing_proto_goTypes = nil
	file_protos_billing_proto_depIdxs = nil
}
