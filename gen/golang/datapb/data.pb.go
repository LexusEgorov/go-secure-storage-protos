// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/data/data.proto

package datapb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Перечисление типов
type Category int32

const (
	Category_PASSWORD Category = 0
	Category_TEXT     Category = 1
	Category_BINARY   Category = 2
	Category_CARD     Category = 3
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "PASSWORD",
		1: "TEXT",
		2: "BINARY",
		3: "CARD",
	}
	Category_value = map[string]int32{
		"PASSWORD": 0,
		"TEXT":     1,
		"BINARY":   2,
		"CARD":     3,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_data_data_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_proto_data_data_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{0}
}

// Сохранение
type AddRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Category Category               `protobuf:"varint,1,opt,name=category,proto3,enum=data.Category" json:"category,omitempty"`
	// Types that are valid to be assigned to Request:
	//
	//	*AddRequest_Password
	//	*AddRequest_Text
	//	*AddRequest_Binary
	//	*AddRequest_Card
	Request       isAddRequest_Request `protobuf_oneof:"request"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_proto_data_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_PASSWORD
}

func (x *AddRequest) GetRequest() isAddRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *AddRequest) GetPassword() *Password {
	if x != nil {
		if x, ok := x.Request.(*AddRequest_Password); ok {
			return x.Password
		}
	}
	return nil
}

func (x *AddRequest) GetText() *Text {
	if x != nil {
		if x, ok := x.Request.(*AddRequest_Text); ok {
			return x.Text
		}
	}
	return nil
}

func (x *AddRequest) GetBinary() *Binary {
	if x != nil {
		if x, ok := x.Request.(*AddRequest_Binary); ok {
			return x.Binary
		}
	}
	return nil
}

func (x *AddRequest) GetCard() *Card {
	if x != nil {
		if x, ok := x.Request.(*AddRequest_Card); ok {
			return x.Card
		}
	}
	return nil
}

type isAddRequest_Request interface {
	isAddRequest_Request()
}

type AddRequest_Password struct {
	Password *Password `protobuf:"bytes,2,opt,name=password,proto3,oneof"`
}

type AddRequest_Text struct {
	Text *Text `protobuf:"bytes,3,opt,name=text,proto3,oneof"`
}

type AddRequest_Binary struct {
	Binary *Binary `protobuf:"bytes,4,opt,name=binary,proto3,oneof"`
}

type AddRequest_Card struct {
	Card *Card `protobuf:"bytes,5,opt,name=card,proto3,oneof"`
}

func (*AddRequest_Password) isAddRequest_Request() {}

func (*AddRequest_Text) isAddRequest_Request() {}

func (*AddRequest_Binary) isAddRequest_Request() {}

func (*AddRequest_Card) isAddRequest_Request() {}

// Получение
type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      Category               `protobuf:"varint,1,opt,name=category,proto3,enum=data.Category" json:"category,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_proto_data_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_PASSWORD
}

func (x *GetRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

// Сохранение
type AddResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Ok    bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	// Types that are valid to be assigned to Response:
	//
	//	*AddResponse_Success
	//	*AddResponse_Bad
	Response      isAddResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	mi := &file_proto_data_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{2}
}

func (x *AddResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *AddResponse) GetResponse() isAddResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *AddResponse) GetSuccess() *SuccessAddResponse {
	if x != nil {
		if x, ok := x.Response.(*AddResponse_Success); ok {
			return x.Success
		}
	}
	return nil
}

func (x *AddResponse) GetBad() *BadResponse {
	if x != nil {
		if x, ok := x.Response.(*AddResponse_Bad); ok {
			return x.Bad
		}
	}
	return nil
}

type isAddResponse_Response interface {
	isAddResponse_Response()
}

type AddResponse_Success struct {
	Success *SuccessAddResponse `protobuf:"bytes,2,opt,name=success,proto3,oneof"`
}

type AddResponse_Bad struct {
	Bad *BadResponse `protobuf:"bytes,3,opt,name=bad,proto3,oneof"`
}

func (*AddResponse_Success) isAddResponse_Response() {}

func (*AddResponse_Bad) isAddResponse_Response() {}

// Получение
type GetResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Ok    bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	// Types that are valid to be assigned to Response:
	//
	//	*GetResponse_Bad
	//	*GetResponse_Success
	Response      isGetResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_proto_data_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *GetResponse) GetResponse() isGetResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *GetResponse) GetBad() *BadResponse {
	if x != nil {
		if x, ok := x.Response.(*GetResponse_Bad); ok {
			return x.Bad
		}
	}
	return nil
}

func (x *GetResponse) GetSuccess() *SuccessGetResponse {
	if x != nil {
		if x, ok := x.Response.(*GetResponse_Success); ok {
			return x.Success
		}
	}
	return nil
}

type isGetResponse_Response interface {
	isGetResponse_Response()
}

type GetResponse_Bad struct {
	Bad *BadResponse `protobuf:"bytes,2,opt,name=bad,proto3,oneof"`
}

type GetResponse_Success struct {
	Success *SuccessGetResponse `protobuf:"bytes,3,opt,name=success,proto3,oneof"`
}

func (*GetResponse_Bad) isGetResponse_Response() {}

func (*GetResponse_Success) isGetResponse_Response() {}

// Сохранение
type SuccessAddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      Category               `protobuf:"varint,1,opt,name=category,proto3,enum=data.Category" json:"category,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuccessAddResponse) Reset() {
	*x = SuccessAddResponse{}
	mi := &file_proto_data_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuccessAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessAddResponse) ProtoMessage() {}

func (x *SuccessAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessAddResponse.ProtoReflect.Descriptor instead.
func (*SuccessAddResponse) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{4}
}

func (x *SuccessAddResponse) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_PASSWORD
}

func (x *SuccessAddResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

// Получение
type SuccessGetResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*SuccessGetResponse_Password
	//	*SuccessGetResponse_Text
	//	*SuccessGetResponse_Binary
	//	*SuccessGetResponse_Card
	Data          isSuccessGetResponse_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuccessGetResponse) Reset() {
	*x = SuccessGetResponse{}
	mi := &file_proto_data_data_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuccessGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessGetResponse) ProtoMessage() {}

func (x *SuccessGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessGetResponse.ProtoReflect.Descriptor instead.
func (*SuccessGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{5}
}

func (x *SuccessGetResponse) GetData() isSuccessGetResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SuccessGetResponse) GetPassword() *Password {
	if x != nil {
		if x, ok := x.Data.(*SuccessGetResponse_Password); ok {
			return x.Password
		}
	}
	return nil
}

func (x *SuccessGetResponse) GetText() *Text {
	if x != nil {
		if x, ok := x.Data.(*SuccessGetResponse_Text); ok {
			return x.Text
		}
	}
	return nil
}

func (x *SuccessGetResponse) GetBinary() *Binary {
	if x != nil {
		if x, ok := x.Data.(*SuccessGetResponse_Binary); ok {
			return x.Binary
		}
	}
	return nil
}

func (x *SuccessGetResponse) GetCard() *Card {
	if x != nil {
		if x, ok := x.Data.(*SuccessGetResponse_Card); ok {
			return x.Card
		}
	}
	return nil
}

type isSuccessGetResponse_Data interface {
	isSuccessGetResponse_Data()
}

type SuccessGetResponse_Password struct {
	Password *Password `protobuf:"bytes,1,opt,name=password,proto3,oneof"`
}

type SuccessGetResponse_Text struct {
	Text *Text `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

type SuccessGetResponse_Binary struct {
	Binary *Binary `protobuf:"bytes,3,opt,name=binary,proto3,oneof"`
}

type SuccessGetResponse_Card struct {
	Card *Card `protobuf:"bytes,4,opt,name=card,proto3,oneof"`
}

func (*SuccessGetResponse_Password) isSuccessGetResponse_Data() {}

func (*SuccessGetResponse_Text) isSuccessGetResponse_Data() {}

func (*SuccessGetResponse_Binary) isSuccessGetResponse_Data() {}

func (*SuccessGetResponse_Card) isSuccessGetResponse_Data() {}

type BadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BadResponse) Reset() {
	*x = BadResponse{}
	mi := &file_proto_data_data_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadResponse) ProtoMessage() {}

func (x *BadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadResponse.ProtoReflect.Descriptor instead.
func (*BadResponse) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{6}
}

func (x *BadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Пароль
type Password struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Password) Reset() {
	*x = Password{}
	mi := &file_proto_data_data_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Password) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Password) ProtoMessage() {}

func (x *Password) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Password.ProtoReflect.Descriptor instead.
func (*Password) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{7}
}

func (x *Password) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Password) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Текст
type Text struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Text) Reset() {
	*x = Text{}
	mi := &file_proto_data_data_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{8}
}

func (x *Text) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Бинарник
type Binary struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Binary        []byte                 `protobuf:"bytes,1,opt,name=binary,proto3" json:"binary,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Binary) Reset() {
	*x = Binary{}
	mi := &file_proto_data_data_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Binary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Binary) ProtoMessage() {}

func (x *Binary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Binary.ProtoReflect.Descriptor instead.
func (*Binary) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{9}
}

func (x *Binary) GetBinary() []byte {
	if x != nil {
		return x.Binary
	}
	return nil
}

// Банковская карта
type Card struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        string                 `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Holder        string                 `protobuf:"bytes,2,opt,name=holder,proto3" json:"holder,omitempty"`
	Exp           string                 `protobuf:"bytes,3,opt,name=exp,proto3" json:"exp,omitempty"`
	Cvv           string                 `protobuf:"bytes,4,opt,name=cvv,proto3" json:"cvv,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Card) Reset() {
	*x = Card{}
	mi := &file_proto_data_data_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_data_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_proto_data_data_proto_rawDescGZIP(), []int{10}
}

func (x *Card) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Card) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

func (x *Card) GetExp() string {
	if x != nil {
		return x.Exp
	}
	return ""
}

func (x *Card) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

var File_proto_data_data_proto protoreflect.FileDescriptor

const file_proto_data_data_proto_rawDesc = "" +
	"\n" +
	"\x15proto/data/data.proto\x12\x04data\"\xdd\x01\n" +
	"\n" +
	"AddRequest\x12*\n" +
	"\bcategory\x18\x01 \x01(\x0e2\x0e.data.CategoryR\bcategory\x12,\n" +
	"\bpassword\x18\x02 \x01(\v2\x0e.data.PasswordH\x00R\bpassword\x12 \n" +
	"\x04text\x18\x03 \x01(\v2\n" +
	".data.TextH\x00R\x04text\x12&\n" +
	"\x06binary\x18\x04 \x01(\v2\f.data.BinaryH\x00R\x06binary\x12 \n" +
	"\x04card\x18\x05 \x01(\v2\n" +
	".data.CardH\x00R\x04cardB\t\n" +
	"\arequest\"T\n" +
	"\n" +
	"GetRequest\x12*\n" +
	"\bcategory\x18\x01 \x01(\x0e2\x0e.data.CategoryR\bcategory\x12\x1a\n" +
	"\bfilename\x18\x02 \x01(\tR\bfilename\"\x86\x01\n" +
	"\vAddResponse\x12\x0e\n" +
	"\x02ok\x18\x01 \x01(\bR\x02ok\x124\n" +
	"\asuccess\x18\x02 \x01(\v2\x18.data.SuccessAddResponseH\x00R\asuccess\x12%\n" +
	"\x03bad\x18\x03 \x01(\v2\x11.data.BadResponseH\x00R\x03badB\n" +
	"\n" +
	"\bresponse\"\x86\x01\n" +
	"\vGetResponse\x12\x0e\n" +
	"\x02ok\x18\x01 \x01(\bR\x02ok\x12%\n" +
	"\x03bad\x18\x02 \x01(\v2\x11.data.BadResponseH\x00R\x03bad\x124\n" +
	"\asuccess\x18\x03 \x01(\v2\x18.data.SuccessGetResponseH\x00R\asuccessB\n" +
	"\n" +
	"\bresponse\"\\\n" +
	"\x12SuccessAddResponse\x12*\n" +
	"\bcategory\x18\x01 \x01(\x0e2\x0e.data.CategoryR\bcategory\x12\x1a\n" +
	"\bfilename\x18\x02 \x01(\tR\bfilename\"\xb6\x01\n" +
	"\x12SuccessGetResponse\x12,\n" +
	"\bpassword\x18\x01 \x01(\v2\x0e.data.PasswordH\x00R\bpassword\x12 \n" +
	"\x04text\x18\x02 \x01(\v2\n" +
	".data.TextH\x00R\x04text\x12&\n" +
	"\x06binary\x18\x03 \x01(\v2\f.data.BinaryH\x00R\x06binary\x12 \n" +
	"\x04card\x18\x04 \x01(\v2\n" +
	".data.CardH\x00R\x04cardB\x06\n" +
	"\x04data\"'\n" +
	"\vBadResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"<\n" +
	"\bPassword\x12\x14\n" +
	"\x05login\x18\x01 \x01(\tR\x05login\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"\x1a\n" +
	"\x04Text\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\" \n" +
	"\x06Binary\x12\x16\n" +
	"\x06binary\x18\x01 \x01(\fR\x06binary\"Z\n" +
	"\x04Card\x12\x16\n" +
	"\x06number\x18\x01 \x01(\tR\x06number\x12\x16\n" +
	"\x06holder\x18\x02 \x01(\tR\x06holder\x12\x10\n" +
	"\x03exp\x18\x03 \x01(\tR\x03exp\x12\x10\n" +
	"\x03cvv\x18\x04 \x01(\tR\x03cvv*8\n" +
	"\bCategory\x12\f\n" +
	"\bPASSWORD\x10\x00\x12\b\n" +
	"\x04TEXT\x10\x01\x12\n" +
	"\n" +
	"\x06BINARY\x10\x02\x12\b\n" +
	"\x04CARD\x10\x032^\n" +
	"\x04Data\x12*\n" +
	"\x03Add\x12\x10.data.AddRequest\x1a\x11.data.AddResponse\x12*\n" +
	"\x03Get\x12\x10.data.GetRequest\x1a\x11.data.GetResponseB\tZ\a/datapbb\x06proto3"

var (
	file_proto_data_data_proto_rawDescOnce sync.Once
	file_proto_data_data_proto_rawDescData []byte
)

func file_proto_data_data_proto_rawDescGZIP() []byte {
	file_proto_data_data_proto_rawDescOnce.Do(func() {
		file_proto_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_data_data_proto_rawDesc), len(file_proto_data_data_proto_rawDesc)))
	})
	return file_proto_data_data_proto_rawDescData
}

var file_proto_data_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_data_data_proto_goTypes = []any{
	(Category)(0),              // 0: data.Category
	(*AddRequest)(nil),         // 1: data.AddRequest
	(*GetRequest)(nil),         // 2: data.GetRequest
	(*AddResponse)(nil),        // 3: data.AddResponse
	(*GetResponse)(nil),        // 4: data.GetResponse
	(*SuccessAddResponse)(nil), // 5: data.SuccessAddResponse
	(*SuccessGetResponse)(nil), // 6: data.SuccessGetResponse
	(*BadResponse)(nil),        // 7: data.BadResponse
	(*Password)(nil),           // 8: data.Password
	(*Text)(nil),               // 9: data.Text
	(*Binary)(nil),             // 10: data.Binary
	(*Card)(nil),               // 11: data.Card
}
var file_proto_data_data_proto_depIdxs = []int32{
	0,  // 0: data.AddRequest.category:type_name -> data.Category
	8,  // 1: data.AddRequest.password:type_name -> data.Password
	9,  // 2: data.AddRequest.text:type_name -> data.Text
	10, // 3: data.AddRequest.binary:type_name -> data.Binary
	11, // 4: data.AddRequest.card:type_name -> data.Card
	0,  // 5: data.GetRequest.category:type_name -> data.Category
	5,  // 6: data.AddResponse.success:type_name -> data.SuccessAddResponse
	7,  // 7: data.AddResponse.bad:type_name -> data.BadResponse
	7,  // 8: data.GetResponse.bad:type_name -> data.BadResponse
	6,  // 9: data.GetResponse.success:type_name -> data.SuccessGetResponse
	0,  // 10: data.SuccessAddResponse.category:type_name -> data.Category
	8,  // 11: data.SuccessGetResponse.password:type_name -> data.Password
	9,  // 12: data.SuccessGetResponse.text:type_name -> data.Text
	10, // 13: data.SuccessGetResponse.binary:type_name -> data.Binary
	11, // 14: data.SuccessGetResponse.card:type_name -> data.Card
	1,  // 15: data.Data.Add:input_type -> data.AddRequest
	2,  // 16: data.Data.Get:input_type -> data.GetRequest
	3,  // 17: data.Data.Add:output_type -> data.AddResponse
	4,  // 18: data.Data.Get:output_type -> data.GetResponse
	17, // [17:19] is the sub-list for method output_type
	15, // [15:17] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_proto_data_data_proto_init() }
func file_proto_data_data_proto_init() {
	if File_proto_data_data_proto != nil {
		return
	}
	file_proto_data_data_proto_msgTypes[0].OneofWrappers = []any{
		(*AddRequest_Password)(nil),
		(*AddRequest_Text)(nil),
		(*AddRequest_Binary)(nil),
		(*AddRequest_Card)(nil),
	}
	file_proto_data_data_proto_msgTypes[2].OneofWrappers = []any{
		(*AddResponse_Success)(nil),
		(*AddResponse_Bad)(nil),
	}
	file_proto_data_data_proto_msgTypes[3].OneofWrappers = []any{
		(*GetResponse_Bad)(nil),
		(*GetResponse_Success)(nil),
	}
	file_proto_data_data_proto_msgTypes[5].OneofWrappers = []any{
		(*SuccessGetResponse_Password)(nil),
		(*SuccessGetResponse_Text)(nil),
		(*SuccessGetResponse_Binary)(nil),
		(*SuccessGetResponse_Card)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_data_data_proto_rawDesc), len(file_proto_data_data_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_data_data_proto_goTypes,
		DependencyIndexes: file_proto_data_data_proto_depIdxs,
		EnumInfos:         file_proto_data_data_proto_enumTypes,
		MessageInfos:      file_proto_data_data_proto_msgTypes,
	}.Build()
	File_proto_data_data_proto = out.File
	file_proto_data_data_proto_goTypes = nil
	file_proto_data_data_proto_depIdxs = nil
}
