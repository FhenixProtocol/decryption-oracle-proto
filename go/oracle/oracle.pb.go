// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: oracle/oracle.proto

package oracle

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type EncryptedType int32

const (
	EncryptedType_Uint8   EncryptedType = 0
	EncryptedType_Uint16  EncryptedType = 1
	EncryptedType_Uint32  EncryptedType = 2
	EncryptedType_Uint64  EncryptedType = 3
	EncryptedType_Uint128 EncryptedType = 4
	EncryptedType_Uint256 EncryptedType = 5
)

// Enum value maps for EncryptedType.
var (
	EncryptedType_name = map[int32]string{
		0: "Uint8",
		1: "Uint16",
		2: "Uint32",
		3: "Uint64",
		4: "Uint128",
		5: "Uint256",
	}
	EncryptedType_value = map[string]int32{
		"Uint8":   0,
		"Uint16":  1,
		"Uint32":  2,
		"Uint64":  3,
		"Uint128": 4,
		"Uint256": 5,
	}
)

func (x EncryptedType) Enum() *EncryptedType {
	p := new(EncryptedType)
	*p = x
	return p
}

func (x EncryptedType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncryptedType) Descriptor() protoreflect.EnumDescriptor {
	return file_oracle_oracle_proto_enumTypes[0].Descriptor()
}

func (EncryptedType) Type() protoreflect.EnumType {
	return &file_oracle_oracle_proto_enumTypes[0]
}

func (x EncryptedType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncryptedType.Descriptor instead.
func (EncryptedType) EnumDescriptor() ([]byte, []int) {
	return file_oracle_oracle_proto_rawDescGZIP(), []int{0}
}

type FheEncrypted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte        `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type EncryptedType `protobuf:"varint,2,opt,name=type,proto3,enum=oracle.EncryptedType" json:"type,omitempty"`
}

func (x *FheEncrypted) Reset() {
	*x = FheEncrypted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_oracle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FheEncrypted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FheEncrypted) ProtoMessage() {}

func (x *FheEncrypted) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_oracle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FheEncrypted.ProtoReflect.Descriptor instead.
func (*FheEncrypted) Descriptor() ([]byte, []int) {
	return file_oracle_oracle_proto_rawDescGZIP(), []int{0}
}

func (x *FheEncrypted) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FheEncrypted) GetType() EncryptedType {
	if x != nil {
		return x.Type
	}
	return EncryptedType_Uint8
}

// The request message containing hex encoded encrypted number
// and a currently used field with some proof (for future use)
type IsNilRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encrypted *FheEncrypted `protobuf:"bytes,1,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	Proof     string        `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *IsNilRequest) Reset() {
	*x = IsNilRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_oracle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsNilRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsNilRequest) ProtoMessage() {}

func (x *IsNilRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_oracle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsNilRequest.ProtoReflect.Descriptor instead.
func (*IsNilRequest) Descriptor() ([]byte, []int) {
	return file_oracle_oracle_proto_rawDescGZIP(), []int{1}
}

func (x *IsNilRequest) GetEncrypted() *FheEncrypted {
	if x != nil {
		return x.Encrypted
	}
	return nil
}

func (x *IsNilRequest) GetProof() string {
	if x != nil {
		return x.Proof
	}
	return ""
}

// The request message containing hex encoded encrypted number
// and the public key of the requesting user (also hex encoded)
// and a currently used field with some proof (for future use)
type ReencryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encrypted     *FheEncrypted `protobuf:"bytes,1,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	UserPublicKey string        `protobuf:"bytes,2,opt,name=user_public_key,json=userPublicKey,proto3" json:"user_public_key,omitempty"`
	Proof         string        `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *ReencryptRequest) Reset() {
	*x = ReencryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_oracle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReencryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReencryptRequest) ProtoMessage() {}

func (x *ReencryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_oracle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReencryptRequest.ProtoReflect.Descriptor instead.
func (*ReencryptRequest) Descriptor() ([]byte, []int) {
	return file_oracle_oracle_proto_rawDescGZIP(), []int{2}
}

func (x *ReencryptRequest) GetEncrypted() *FheEncrypted {
	if x != nil {
		return x.Encrypted
	}
	return nil
}

func (x *ReencryptRequest) GetUserPublicKey() string {
	if x != nil {
		return x.UserPublicKey
	}
	return ""
}

func (x *ReencryptRequest) GetProof() string {
	if x != nil {
		return x.Proof
	}
	return ""
}

// The request message containing hex encoded encrypted number
// and a currently used field with some proof (for future use)
type DecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encrypted *FheEncrypted `protobuf:"bytes,1,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	Proof     string        `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *DecryptRequest) Reset() {
	*x = DecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_oracle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptRequest) ProtoMessage() {}

func (x *DecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_oracle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptRequest.ProtoReflect.Descriptor instead.
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return file_oracle_oracle_proto_rawDescGZIP(), []int{3}
}

func (x *DecryptRequest) GetEncrypted() *FheEncrypted {
	if x != nil {
		return x.Encrypted
	}
	return nil
}

func (x *DecryptRequest) GetProof() string {
	if x != nil {
		return x.Proof
	}
	return ""
}

// The response message containing the greetings
type DecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decrypted string `protobuf:"bytes,1,opt,name=decrypted,proto3" json:"decrypted,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *DecryptResponse) Reset() {
	*x = DecryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_oracle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptResponse) ProtoMessage() {}

func (x *DecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_oracle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptResponse.ProtoReflect.Descriptor instead.
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return file_oracle_oracle_proto_rawDescGZIP(), []int{4}
}

func (x *DecryptResponse) GetDecrypted() string {
	if x != nil {
		return x.Decrypted
	}
	return ""
}

func (x *DecryptResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// The response message containing the result whether or not the
// assertion requested was nil
type IsNilResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsNil     bool   `protobuf:"varint,1,opt,name=is_nil,json=isNil,proto3" json:"is_nil,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *IsNilResponse) Reset() {
	*x = IsNilResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_oracle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsNilResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsNilResponse) ProtoMessage() {}

func (x *IsNilResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_oracle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsNilResponse.ProtoReflect.Descriptor instead.
func (*IsNilResponse) Descriptor() ([]byte, []int) {
	return file_oracle_oracle_proto_rawDescGZIP(), []int{5}
}

func (x *IsNilResponse) GetIsNil() bool {
	if x != nil {
		return x.IsNil
	}
	return false
}

func (x *IsNilResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// The response message containing a hex encoded reencrypted number
type ReencryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reencrypted string `protobuf:"bytes,1,opt,name=reencrypted,proto3" json:"reencrypted,omitempty"`
	Signature   string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ReencryptResponse) Reset() {
	*x = ReencryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_oracle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReencryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReencryptResponse) ProtoMessage() {}

func (x *ReencryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_oracle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReencryptResponse.ProtoReflect.Descriptor instead.
func (*ReencryptResponse) Descriptor() ([]byte, []int) {
	return file_oracle_oracle_proto_rawDescGZIP(), []int{6}
}

func (x *ReencryptResponse) GetReencrypted() string {
	if x != nil {
		return x.Reencrypted
	}
	return ""
}

func (x *ReencryptResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_oracle_oracle_proto protoreflect.FileDescriptor

var file_oracle_oracle_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52,
	0x0a, 0x0c, 0x46, 0x68, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x17,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x5d, 0x0a, 0x0c, 0x49, 0x73, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x46,
	0x68, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x2e, 0x46, 0x68, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x5f, 0x0a,
	0x0e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x37, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x46, 0x68, 0x65, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x4d,
	0x0a, 0x0f, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x44, 0x0a,
	0x0d, 0x49, 0x73, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6e, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x69, 0x73, 0x4e, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x53, 0x0a, 0x11, 0x52, 0x65, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2a, 0x58, 0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x69, 0x6e,
	0x74, 0x38, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x69, 0x6e, 0x74, 0x31, 0x36, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x69, 0x6e, 0x74,
	0x31, 0x32, 0x38, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x69, 0x6e, 0x74, 0x32, 0x35, 0x36,
	0x10, 0x05, 0x32, 0xd2, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x12, 0x16, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x52, 0x65, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x12, 0x18, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x41, 0x73, 0x73,
	0x65, 0x72, 0x74, 0x49, 0x73, 0x4e, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x2e, 0x49, 0x73, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x49, 0x73, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x6f, 0x2f, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oracle_oracle_proto_rawDescOnce sync.Once
	file_oracle_oracle_proto_rawDescData = file_oracle_oracle_proto_rawDesc
)

func file_oracle_oracle_proto_rawDescGZIP() []byte {
	file_oracle_oracle_proto_rawDescOnce.Do(func() {
		file_oracle_oracle_proto_rawDescData = protoimpl.X.CompressGZIP(file_oracle_oracle_proto_rawDescData)
	})
	return file_oracle_oracle_proto_rawDescData
}

var file_oracle_oracle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_oracle_oracle_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_oracle_oracle_proto_goTypes = []interface{}{
	(EncryptedType)(0),        // 0: oracle.EncryptedType
	(*FheEncrypted)(nil),      // 1: oracle.FheEncrypted
	(*IsNilRequest)(nil),      // 2: oracle.IsNilRequest
	(*ReencryptRequest)(nil),  // 3: oracle.ReencryptRequest
	(*DecryptRequest)(nil),    // 4: oracle.DecryptRequest
	(*DecryptResponse)(nil),   // 5: oracle.DecryptResponse
	(*IsNilResponse)(nil),     // 6: oracle.IsNilResponse
	(*ReencryptResponse)(nil), // 7: oracle.ReencryptResponse
}
var file_oracle_oracle_proto_depIdxs = []int32{
	0, // 0: oracle.FheEncrypted.type:type_name -> oracle.EncryptedType
	1, // 1: oracle.IsNilRequest.encrypted:type_name -> oracle.FheEncrypted
	1, // 2: oracle.ReencryptRequest.encrypted:type_name -> oracle.FheEncrypted
	1, // 3: oracle.DecryptRequest.encrypted:type_name -> oracle.FheEncrypted
	4, // 4: oracle.DecryptionOracle.Decrypt:input_type -> oracle.DecryptRequest
	3, // 5: oracle.DecryptionOracle.Reencrypt:input_type -> oracle.ReencryptRequest
	2, // 6: oracle.DecryptionOracle.AssertIsNil:input_type -> oracle.IsNilRequest
	5, // 7: oracle.DecryptionOracle.Decrypt:output_type -> oracle.DecryptResponse
	7, // 8: oracle.DecryptionOracle.Reencrypt:output_type -> oracle.ReencryptResponse
	6, // 9: oracle.DecryptionOracle.AssertIsNil:output_type -> oracle.IsNilResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_oracle_oracle_proto_init() }
func file_oracle_oracle_proto_init() {
	if File_oracle_oracle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oracle_oracle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FheEncrypted); i {
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
		file_oracle_oracle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsNilRequest); i {
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
		file_oracle_oracle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReencryptRequest); i {
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
		file_oracle_oracle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptRequest); i {
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
		file_oracle_oracle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptResponse); i {
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
		file_oracle_oracle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsNilResponse); i {
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
		file_oracle_oracle_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReencryptResponse); i {
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
			RawDescriptor: file_oracle_oracle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oracle_oracle_proto_goTypes,
		DependencyIndexes: file_oracle_oracle_proto_depIdxs,
		EnumInfos:         file_oracle_oracle_proto_enumTypes,
		MessageInfos:      file_oracle_oracle_proto_msgTypes,
	}.Build()
	File_oracle_oracle_proto = out.File
	file_oracle_oracle_proto_rawDesc = nil
	file_oracle_oracle_proto_goTypes = nil
	file_oracle_oracle_proto_depIdxs = nil
}
