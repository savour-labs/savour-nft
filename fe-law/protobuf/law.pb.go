// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: law.proto

package protobuf

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

type ReturnCode int32

const (
	ReturnCode_SUCCESS ReturnCode = 0
	ReturnCode_ERROR   ReturnCode = 1
)

// Enum value maps for ReturnCode.
var (
	ReturnCode_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
	}
	ReturnCode_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   1,
	}
)

func (x ReturnCode) Enum() *ReturnCode {
	p := new(ReturnCode)
	*p = x
	return p
}

func (x ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_law_proto_enumTypes[0].Descriptor()
}

func (ReturnCode) Type() protoreflect.EnumType {
	return &file_law_proto_enumTypes[0]
}

func (x ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReturnCode.Descriptor instead.
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_law_proto_rawDescGZIP(), []int{0}
}

type GiantWhaleWalletAddressLawReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
}

func (x *GiantWhaleWalletAddressLawReq) Reset() {
	*x = GiantWhaleWalletAddressLawReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_law_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiantWhaleWalletAddressLawReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiantWhaleWalletAddressLawReq) ProtoMessage() {}

func (x *GiantWhaleWalletAddressLawReq) ProtoReflect() protoreflect.Message {
	mi := &file_law_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiantWhaleWalletAddressLawReq.ProtoReflect.Descriptor instead.
func (*GiantWhaleWalletAddressLawReq) Descriptor() ([]byte, []int) {
	return file_law_proto_rawDescGZIP(), []int{0}
}

func (x *GiantWhaleWalletAddressLawReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type GiantWhaleWalletAddressLawRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code          ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=fe_law.protobuf.ReturnCode" json:"code,omitempty"`
	Msg           string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	TokenValue    string     `protobuf:"bytes,3,opt,name=token_value,json=tokenValue,proto3" json:"token_value,omitempty"`
	NftValue      string     `protobuf:"bytes,4,opt,name=nft_value,json=nftValue,proto3" json:"nft_value,omitempty"`
	TokenActivity uint64     `protobuf:"varint,5,opt,name=token_activity,json=tokenActivity,proto3" json:"token_activity,omitempty"`
	NftActivity   uint64     `protobuf:"varint,6,opt,name=nft_activity,json=nftActivity,proto3" json:"nft_activity,omitempty"`
	TotalToken    uint64     `protobuf:"varint,7,opt,name=total_token,json=totalToken,proto3" json:"total_token,omitempty"`
	TotalNft      uint64     `protobuf:"varint,8,opt,name=total_nft,json=totalNft,proto3" json:"total_nft,omitempty"`
}

func (x *GiantWhaleWalletAddressLawRep) Reset() {
	*x = GiantWhaleWalletAddressLawRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_law_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiantWhaleWalletAddressLawRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiantWhaleWalletAddressLawRep) ProtoMessage() {}

func (x *GiantWhaleWalletAddressLawRep) ProtoReflect() protoreflect.Message {
	mi := &file_law_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiantWhaleWalletAddressLawRep.ProtoReflect.Descriptor instead.
func (*GiantWhaleWalletAddressLawRep) Descriptor() ([]byte, []int) {
	return file_law_proto_rawDescGZIP(), []int{1}
}

func (x *GiantWhaleWalletAddressLawRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *GiantWhaleWalletAddressLawRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GiantWhaleWalletAddressLawRep) GetTokenValue() string {
	if x != nil {
		return x.TokenValue
	}
	return ""
}

func (x *GiantWhaleWalletAddressLawRep) GetNftValue() string {
	if x != nil {
		return x.NftValue
	}
	return ""
}

func (x *GiantWhaleWalletAddressLawRep) GetTokenActivity() uint64 {
	if x != nil {
		return x.TokenActivity
	}
	return 0
}

func (x *GiantWhaleWalletAddressLawRep) GetNftActivity() uint64 {
	if x != nil {
		return x.NftActivity
	}
	return 0
}

func (x *GiantWhaleWalletAddressLawRep) GetTotalToken() uint64 {
	if x != nil {
		return x.TotalToken
	}
	return 0
}

func (x *GiantWhaleWalletAddressLawRep) GetTotalNft() uint64 {
	if x != nil {
		return x.TotalNft
	}
	return 0
}

type NftCollectionsLawReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
}

func (x *NftCollectionsLawReq) Reset() {
	*x = NftCollectionsLawReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_law_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftCollectionsLawReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftCollectionsLawReq) ProtoMessage() {}

func (x *NftCollectionsLawReq) ProtoReflect() protoreflect.Message {
	mi := &file_law_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftCollectionsLawReq.ProtoReflect.Descriptor instead.
func (*NftCollectionsLawReq) Descriptor() ([]byte, []int) {
	return file_law_proto_rawDescGZIP(), []int{2}
}

func (x *NftCollectionsLawReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type NftCollectionsLawRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code                              ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=fe_law.protobuf.ReturnCode" json:"code,omitempty"`
	Msg                               string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	CollectionAverageValue            string     `protobuf:"bytes,3,opt,name=collection_average_value,json=collectionAverageValue,proto3" json:"collection_average_value,omitempty"`
	CollectionAverageTransactions     uint64     `protobuf:"varint,4,opt,name=collection_average_transactions,json=collectionAverageTransactions,proto3" json:"collection_average_transactions,omitempty"`
	CollectionDailyTransactions       uint64     `protobuf:"varint,5,opt,name=collection_daily_transactions,json=collectionDailyTransactions,proto3" json:"collection_daily_transactions,omitempty"`
	CollectionHolderAddress           uint64     `protobuf:"varint,6,opt,name=collection_holder_address,json=collectionHolderAddress,proto3" json:"collection_holder_address,omitempty"`
	CollectionAverageTransactionPrice string     `protobuf:"bytes,7,opt,name=collection_average_transactionPrice,json=collectionAverageTransactionPrice,proto3" json:"collection_average_transactionPrice,omitempty"`
	CollectionDailyTransactionPrice   string     `protobuf:"bytes,8,opt,name=collection_daily_transactionPrice,json=collectionDailyTransactionPrice,proto3" json:"collection_daily_transactionPrice,omitempty"`
}

func (x *NftCollectionsLawRep) Reset() {
	*x = NftCollectionsLawRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_law_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftCollectionsLawRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftCollectionsLawRep) ProtoMessage() {}

func (x *NftCollectionsLawRep) ProtoReflect() protoreflect.Message {
	mi := &file_law_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftCollectionsLawRep.ProtoReflect.Descriptor instead.
func (*NftCollectionsLawRep) Descriptor() ([]byte, []int) {
	return file_law_proto_rawDescGZIP(), []int{3}
}

func (x *NftCollectionsLawRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *NftCollectionsLawRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *NftCollectionsLawRep) GetCollectionAverageValue() string {
	if x != nil {
		return x.CollectionAverageValue
	}
	return ""
}

func (x *NftCollectionsLawRep) GetCollectionAverageTransactions() uint64 {
	if x != nil {
		return x.CollectionAverageTransactions
	}
	return 0
}

func (x *NftCollectionsLawRep) GetCollectionDailyTransactions() uint64 {
	if x != nil {
		return x.CollectionDailyTransactions
	}
	return 0
}

func (x *NftCollectionsLawRep) GetCollectionHolderAddress() uint64 {
	if x != nil {
		return x.CollectionHolderAddress
	}
	return 0
}

func (x *NftCollectionsLawRep) GetCollectionAverageTransactionPrice() string {
	if x != nil {
		return x.CollectionAverageTransactionPrice
	}
	return ""
}

func (x *NftCollectionsLawRep) GetCollectionDailyTransactionPrice() string {
	if x != nil {
		return x.CollectionDailyTransactionPrice
	}
	return ""
}

type SingleNftLawReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
}

func (x *SingleNftLawReq) Reset() {
	*x = SingleNftLawReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_law_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleNftLawReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleNftLawReq) ProtoMessage() {}

func (x *SingleNftLawReq) ProtoReflect() protoreflect.Message {
	mi := &file_law_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleNftLawReq.ProtoReflect.Descriptor instead.
func (*SingleNftLawReq) Descriptor() ([]byte, []int) {
	return file_law_proto_rawDescGZIP(), []int{4}
}

func (x *SingleNftLawReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type SingleNftLawRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code                       ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=fe_law.protobuf.ReturnCode" json:"code,omitempty"`
	Msg                        string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	NftTotalTransactions       uint64     `protobuf:"varint,3,opt,name=nft_total_transactions,json=nftTotalTransactions,proto3" json:"nft_total_transactions,omitempty"`
	NftDailyTransactions       uint64     `protobuf:"varint,4,opt,name=nft_daily_transactions,json=nftDailyTransactions,proto3" json:"nft_daily_transactions,omitempty"`
	NftLatestPrice             string     `protobuf:"bytes,5,opt,name=nft_latest_price,json=nftLatestPrice,proto3" json:"nft_latest_price,omitempty"`
	NftAverageTransactionPrice string     `protobuf:"bytes,6,opt,name=nft_average_transactionPrice,json=nftAverageTransactionPrice,proto3" json:"nft_average_transactionPrice,omitempty"`
	NftDailyTransactionPrice   string     `protobuf:"bytes,7,opt,name=nft_daily_transactionPrice,json=nftDailyTransactionPrice,proto3" json:"nft_daily_transactionPrice,omitempty"`
}

func (x *SingleNftLawRep) Reset() {
	*x = SingleNftLawRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_law_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleNftLawRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleNftLawRep) ProtoMessage() {}

func (x *SingleNftLawRep) ProtoReflect() protoreflect.Message {
	mi := &file_law_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleNftLawRep.ProtoReflect.Descriptor instead.
func (*SingleNftLawRep) Descriptor() ([]byte, []int) {
	return file_law_proto_rawDescGZIP(), []int{5}
}

func (x *SingleNftLawRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *SingleNftLawRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SingleNftLawRep) GetNftTotalTransactions() uint64 {
	if x != nil {
		return x.NftTotalTransactions
	}
	return 0
}

func (x *SingleNftLawRep) GetNftDailyTransactions() uint64 {
	if x != nil {
		return x.NftDailyTransactions
	}
	return 0
}

func (x *SingleNftLawRep) GetNftLatestPrice() string {
	if x != nil {
		return x.NftLatestPrice
	}
	return ""
}

func (x *SingleNftLawRep) GetNftAverageTransactionPrice() string {
	if x != nil {
		return x.NftAverageTransactionPrice
	}
	return ""
}

func (x *SingleNftLawRep) GetNftDailyTransactionPrice() string {
	if x != nil {
		return x.NftDailyTransactionPrice
	}
	return ""
}

var File_law_proto protoreflect.FileDescriptor

var file_law_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x65, 0x5f,
	0x6c, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x46, 0x0a, 0x1d,
	0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa8, 0x02, 0x0a, 0x1d, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68,
	0x61, 0x6c, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4c, 0x61, 0x77, 0x52, 0x65, 0x70, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x65, 0x5f, 0x6c, 0x61, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x66,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x66, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x66, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6e, 0x66, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x66, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x66, 0x74, 0x22,
	0x3d, 0x0a, 0x14, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf7,
	0x03, 0x0a, 0x14, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x70, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x65, 0x5f, 0x6c, 0x61, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x18, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x1f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1d, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x1d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x1b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3a, 0x0a, 0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4e, 0x0a, 0x23,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x21, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x21,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x0f, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x4e, 0x66, 0x74, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xea, 0x02, 0x0a, 0x0f, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4e, 0x66, 0x74,
	0x4c, 0x61, 0x77, 0x52, 0x65, 0x70, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x65, 0x5f, 0x6c, 0x61, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x34, 0x0a, 0x16, 0x6e, 0x66, 0x74,
	0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x6e, 0x66, 0x74, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x34, 0x0a, 0x16, 0x6e, 0x66, 0x74, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x14, 0x6e, 0x66, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x66, 0x74, 0x5f, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6e, 0x66, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x1c, 0x6e, 0x66, 0x74, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x6e, 0x66, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x1a, 0x6e, 0x66, 0x74, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6e, 0x66, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x2a,
	0x24, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0xd4, 0x02, 0x0a, 0x0d, 0x4c, 0x61, 0x77, 0x52, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x47,
	0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x61, 0x77, 0x12, 0x2e, 0x2e, 0x66, 0x65, 0x5f, 0x6c,
	0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x69, 0x61, 0x6e,
	0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x66, 0x65, 0x5f, 0x6c,
	0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x69, 0x61, 0x6e,
	0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x4c, 0x61, 0x77, 0x12, 0x25, 0x2e, 0x66, 0x65, 0x5f, 0x6c, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x66, 0x65, 0x5f,
	0x6c, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x66, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x61, 0x77, 0x52, 0x65,
	0x70, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x4e, 0x66, 0x74, 0x4c, 0x61, 0x77, 0x12, 0x20, 0x2e, 0x66, 0x65, 0x5f, 0x6c, 0x61, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4e,
	0x66, 0x74, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x66, 0x65, 0x5f, 0x6c, 0x61,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x4e, 0x66, 0x74, 0x4c, 0x61, 0x77, 0x52, 0x65, 0x70, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12,
	0x2e, 0x2e, 0x2f, 0x66, 0x65, 0x2d, 0x6c, 0x61, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_law_proto_rawDescOnce sync.Once
	file_law_proto_rawDescData = file_law_proto_rawDesc
)

func file_law_proto_rawDescGZIP() []byte {
	file_law_proto_rawDescOnce.Do(func() {
		file_law_proto_rawDescData = protoimpl.X.CompressGZIP(file_law_proto_rawDescData)
	})
	return file_law_proto_rawDescData
}

var file_law_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_law_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_law_proto_goTypes = []interface{}{
	(ReturnCode)(0),                       // 0: fe_law.protobuf.ReturnCode
	(*GiantWhaleWalletAddressLawReq)(nil), // 1: fe_law.protobuf.GiantWhaleWalletAddressLawReq
	(*GiantWhaleWalletAddressLawRep)(nil), // 2: fe_law.protobuf.GiantWhaleWalletAddressLawRep
	(*NftCollectionsLawReq)(nil),          // 3: fe_law.protobuf.NftCollectionsLawReq
	(*NftCollectionsLawRep)(nil),          // 4: fe_law.protobuf.NftCollectionsLawRep
	(*SingleNftLawReq)(nil),               // 5: fe_law.protobuf.SingleNftLawReq
	(*SingleNftLawRep)(nil),               // 6: fe_law.protobuf.SingleNftLawRep
}
var file_law_proto_depIdxs = []int32{
	0, // 0: fe_law.protobuf.GiantWhaleWalletAddressLawRep.code:type_name -> fe_law.protobuf.ReturnCode
	0, // 1: fe_law.protobuf.NftCollectionsLawRep.code:type_name -> fe_law.protobuf.ReturnCode
	0, // 2: fe_law.protobuf.SingleNftLawRep.code:type_name -> fe_law.protobuf.ReturnCode
	1, // 3: fe_law.protobuf.LawRpcService.GetGiantWhaleWalletAddressLaw:input_type -> fe_law.protobuf.GiantWhaleWalletAddressLawReq
	3, // 4: fe_law.protobuf.LawRpcService.GetNftCollectionsLaw:input_type -> fe_law.protobuf.NftCollectionsLawReq
	5, // 5: fe_law.protobuf.LawRpcService.GetSingleNftLaw:input_type -> fe_law.protobuf.SingleNftLawReq
	2, // 6: fe_law.protobuf.LawRpcService.GetGiantWhaleWalletAddressLaw:output_type -> fe_law.protobuf.GiantWhaleWalletAddressLawRep
	4, // 7: fe_law.protobuf.LawRpcService.GetNftCollectionsLaw:output_type -> fe_law.protobuf.NftCollectionsLawRep
	6, // 8: fe_law.protobuf.LawRpcService.GetSingleNftLaw:output_type -> fe_law.protobuf.SingleNftLawRep
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_law_proto_init() }
func file_law_proto_init() {
	if File_law_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_law_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiantWhaleWalletAddressLawReq); i {
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
		file_law_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiantWhaleWalletAddressLawRep); i {
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
		file_law_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftCollectionsLawReq); i {
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
		file_law_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftCollectionsLawRep); i {
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
		file_law_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleNftLawReq); i {
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
		file_law_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleNftLawRep); i {
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
			RawDescriptor: file_law_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_law_proto_goTypes,
		DependencyIndexes: file_law_proto_depIdxs,
		EnumInfos:         file_law_proto_enumTypes,
		MessageInfos:      file_law_proto_msgTypes,
	}.Build()
	File_law_proto = out.File
	file_law_proto_rawDesc = nil
	file_law_proto_goTypes = nil
	file_law_proto_depIdxs = nil
}