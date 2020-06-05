// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: user/proto/user.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User          *User          `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	ValidationErr *ValidationErr `protobuf:"bytes,2,opt,name=validationErr,proto3" json:"validationErr,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Response) GetValidationErr() *ValidationErr {
	if x != nil {
		return x.ValidationErr
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname  string               `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname   string               `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	ValidFrom  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=validFrom,proto3" json:"validFrom,omitempty"`
	ValidThru  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=validThru,proto3" json:"validThru,omitempty"`
	Active     bool                 `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	Pwd        string               `protobuf:"bytes,7,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Name       string               `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Email      string               `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	Company    string               `protobuf:"bytes,10,opt,name=company,proto3" json:"company,omitempty"`
	Createdate *timestamp.Timestamp `protobuf:"bytes,11,opt,name=createdate,proto3" json:"createdate,omitempty"`
	Updatedate *timestamp.Timestamp `protobuf:"bytes,12,opt,name=updatedate,proto3" json:"updatedate,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetValidFrom() *timestamp.Timestamp {
	if x != nil {
		return x.ValidFrom
	}
	return nil
}

func (x *User) GetValidThru() *timestamp.Timestamp {
	if x != nil {
		return x.ValidThru
	}
	return nil
}

func (x *User) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *User) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *User) GetCreatedate() *timestamp.Timestamp {
	if x != nil {
		return x.Createdate
	}
	return nil
}

func (x *User) GetUpdatedate() *timestamp.Timestamp {
	if x != nil {
		return x.Updatedate
	}
	return nil
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *Users) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

type SearchParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fisrtname string               `protobuf:"bytes,2,opt,name=fisrtname,proto3" json:"fisrtname,omitempty"`
	Lastname  string               `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	ValidDate *timestamp.Timestamp `protobuf:"bytes,5,opt,name=validDate,proto3" json:"validDate,omitempty"`
	Email     string               `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Company   string               `protobuf:"bytes,7,opt,name=company,proto3" json:"company,omitempty"`
	Pwd       string               `protobuf:"bytes,8,opt,name=pwd,proto3" json:"pwd,omitempty"`
}

func (x *SearchParams) Reset() {
	*x = SearchParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchParams) ProtoMessage() {}

func (x *SearchParams) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchParams.ProtoReflect.Descriptor instead.
func (*SearchParams) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *SearchParams) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchParams) GetFisrtname() string {
	if x != nil {
		return x.Fisrtname
	}
	return ""
}

func (x *SearchParams) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *SearchParams) GetValidDate() *timestamp.Timestamp {
	if x != nil {
		return x.ValidDate
	}
	return nil
}

func (x *SearchParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SearchParams) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *SearchParams) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type SearchString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SearchString) Reset() {
	*x = SearchString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchString) ProtoMessage() {}

func (x *SearchString) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchString.ProtoReflect.Descriptor instead.
func (*SearchString) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *SearchString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SearchId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SearchId) Reset() {
	*x = SearchId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchId) ProtoMessage() {}

func (x *SearchId) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchId.ProtoReflect.Descriptor instead.
func (*SearchId) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{5}
}

func (x *SearchId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AffectedCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AffectedCount) Reset() {
	*x = AffectedCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AffectedCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AffectedCount) ProtoMessage() {}

func (x *AffectedCount) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AffectedCount.ProtoReflect.Descriptor instead.
func (*AffectedCount) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{6}
}

func (x *AffectedCount) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ValidationErr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailureDesc []string `protobuf:"bytes,1,rep,name=failureDesc,proto3" json:"failureDesc,omitempty"`
}

func (x *ValidationErr) Reset() {
	*x = ValidationErr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationErr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationErr) ProtoMessage() {}

func (x *ValidationErr) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationErr.ProtoReflect.Descriptor instead.
func (*ValidationErr) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{7}
}

func (x *ValidationErr) GetFailureDesc() []string {
	if x != nil {
		return x.FailureDesc
	}
	return nil
}

type AfterFuncErr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailureDesc []string `protobuf:"bytes,1,rep,name=failureDesc,proto3" json:"failureDesc,omitempty"`
}

func (x *AfterFuncErr) Reset() {
	*x = AfterFuncErr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AfterFuncErr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AfterFuncErr) ProtoMessage() {}

func (x *AfterFuncErr) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AfterFuncErr.ProtoReflect.Descriptor instead.
func (*AfterFuncErr) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{8}
}

func (x *AfterFuncErr) GetFailureDesc() []string {
	if x != nil {
		return x.FailureDesc
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid bool   `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_user_proto_user_proto_rawDescGZIP(), []int{9}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_user_proto_user_proto protoreflect.FileDescriptor

var file_user_proto_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x22, 0xaa, 0x03, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x38, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x68, 0x72, 0x75, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x68, 0x72, 0x75, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x70, 0x77, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x61,
	0x74, 0x65, 0x22, 0x27, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xd4, 0x01, 0x0a, 0x0c,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x73, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x73, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x77, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x0d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x22, 0x30,
	0x0a, 0x0c, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x45, 0x72, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x22, 0x33, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0xbc, 0x05, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53, 0x72,
	0x76, 0x12, 0x2b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x64,
	0x1a, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2d,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x10, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x10, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x10, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x13,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x41, 0x66, 0x74, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x46, 0x75, 0x6e, 0x63, 0x45, 0x72, 0x72, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x45, 0x72, 0x72, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x0f, 0x41, 0x66, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x12,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x45,
	0x72, 0x72, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x6f, 0x54, 0x65, 0x6d, 0x70, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_proto_user_proto_rawDescOnce sync.Once
	file_user_proto_user_proto_rawDescData = file_user_proto_user_proto_rawDesc
)

func file_user_proto_user_proto_rawDescGZIP() []byte {
	file_user_proto_user_proto_rawDescOnce.Do(func() {
		file_user_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_user_proto_rawDescData)
	})
	return file_user_proto_user_proto_rawDescData
}

var file_user_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_proto_user_proto_goTypes = []interface{}{
	(*Response)(nil),            // 0: user.response
	(*User)(nil),                // 1: user.user
	(*Users)(nil),               // 2: user.users
	(*SearchParams)(nil),        // 3: user.SearchParams
	(*SearchString)(nil),        // 4: user.SearchString
	(*SearchId)(nil),            // 5: user.SearchId
	(*AffectedCount)(nil),       // 6: user.affectedCount
	(*ValidationErr)(nil),       // 7: user.validationErr
	(*AfterFuncErr)(nil),        // 8: user.AfterFuncErr
	(*Token)(nil),               // 9: user.Token
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_user_proto_user_proto_depIdxs = []int32{
	1,  // 0: user.response.user:type_name -> user.user
	7,  // 1: user.response.validationErr:type_name -> user.validationErr
	10, // 2: user.user.validFrom:type_name -> google.protobuf.Timestamp
	10, // 3: user.user.validThru:type_name -> google.protobuf.Timestamp
	10, // 4: user.user.createdate:type_name -> google.protobuf.Timestamp
	10, // 5: user.user.updatedate:type_name -> google.protobuf.Timestamp
	1,  // 6: user.users.user:type_name -> user.user
	10, // 7: user.SearchParams.validDate:type_name -> google.protobuf.Timestamp
	5,  // 8: user.userSrv.GetUserById:input_type -> user.SearchId
	3,  // 9: user.userSrv.GetUsers:input_type -> user.SearchParams
	4,  // 10: user.userSrv.GetUsersByEmail:input_type -> user.SearchString
	1,  // 11: user.userSrv.CreateUser:input_type -> user.user
	1,  // 12: user.userSrv.UpdateUser:input_type -> user.user
	5,  // 13: user.userSrv.DeleteUser:input_type -> user.SearchId
	1,  // 14: user.userSrv.BeforeCreateUser:input_type -> user.user
	1,  // 15: user.userSrv.BeforeUpdateUser:input_type -> user.user
	1,  // 16: user.userSrv.BeforeDeleteUser:input_type -> user.user
	1,  // 17: user.userSrv.AfterCreateUser:input_type -> user.user
	1,  // 18: user.userSrv.AfterUpdateUser:input_type -> user.user
	1,  // 19: user.userSrv.AfterDeleteUser:input_type -> user.user
	1,  // 20: user.userSrv.Auth:input_type -> user.user
	9,  // 21: user.userSrv.ValidateToken:input_type -> user.Token
	1,  // 22: user.userSrv.GetUserById:output_type -> user.user
	2,  // 23: user.userSrv.GetUsers:output_type -> user.users
	2,  // 24: user.userSrv.GetUsersByEmail:output_type -> user.users
	0,  // 25: user.userSrv.CreateUser:output_type -> user.response
	0,  // 26: user.userSrv.UpdateUser:output_type -> user.response
	6,  // 27: user.userSrv.DeleteUser:output_type -> user.affectedCount
	7,  // 28: user.userSrv.BeforeCreateUser:output_type -> user.validationErr
	7,  // 29: user.userSrv.BeforeUpdateUser:output_type -> user.validationErr
	7,  // 30: user.userSrv.BeforeDeleteUser:output_type -> user.validationErr
	8,  // 31: user.userSrv.AfterCreateUser:output_type -> user.AfterFuncErr
	8,  // 32: user.userSrv.AfterUpdateUser:output_type -> user.AfterFuncErr
	8,  // 33: user.userSrv.AfterDeleteUser:output_type -> user.AfterFuncErr
	9,  // 34: user.userSrv.Auth:output_type -> user.Token
	9,  // 35: user.userSrv.ValidateToken:output_type -> user.Token
	22, // [22:36] is the sub-list for method output_type
	8,  // [8:22] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_user_proto_user_proto_init() }
func file_user_proto_user_proto_init() {
	if File_user_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_user_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_user_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchParams); i {
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
		file_user_proto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchString); i {
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
		file_user_proto_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchId); i {
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
		file_user_proto_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AffectedCount); i {
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
		file_user_proto_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationErr); i {
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
		file_user_proto_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AfterFuncErr); i {
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
		file_user_proto_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_user_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_user_proto_goTypes,
		DependencyIndexes: file_user_proto_user_proto_depIdxs,
		MessageInfos:      file_user_proto_user_proto_msgTypes,
	}.Build()
	File_user_proto_user_proto = out.File
	file_user_proto_user_proto_rawDesc = nil
	file_user_proto_user_proto_goTypes = nil
	file_user_proto_user_proto_depIdxs = nil
}
