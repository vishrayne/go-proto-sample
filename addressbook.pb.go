// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

package prototest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_addressbook_e0bcdd12bdbe6e61, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressbook_e0bcdd12bdbe6e61, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=prototest.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressbook_e0bcdd12bdbe6e61, []int{0, 0}
}
func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (dst *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(dst, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressbook_e0bcdd12bdbe6e61, []int{1}
}
func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (dst *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(dst, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "prototest.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "prototest.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "prototest.AddressBook")
	proto.RegisterEnum("prototest.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor_addressbook_e0bcdd12bdbe6e61) }

var fileDescriptor_addressbook_e0bcdd12bdbe6e61 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0x34, 0x5d, 0xcc, 0x04, 0x4a, 0x3a, 0x88, 0x04, 0x45, 0x08, 0x39, 0x45, 0x84,
	0x08, 0x15, 0xc1, 0xab, 0x85, 0x82, 0xa2, 0x35, 0x65, 0x11, 0x3d, 0x27, 0x64, 0xc0, 0xd0, 0x26,
	0xb3, 0x64, 0xe3, 0xa1, 0x7f, 0x5e, 0x24, 0x9b, 0x50, 0x04, 0xe9, 0x69, 0xdf, 0xcc, 0x7c, 0xcc,
	0xbc, 0x7d, 0x30, 0xcf, 0xcb, 0xb2, 0x25, 0x63, 0x0a, 0xe6, 0x6d, 0xaa, 0x5b, 0xee, 0x18, 0x3d,
	0xfb, 0x74, 0x64, 0xba, 0xf8, 0x47, 0x80, 0xdc, 0x50, 0x6b, 0xb8, 0x41, 0x04, 0xb7, 0xc9, 0x6b,
	0x0a, 0x45, 0x24, 0x12, 0x4f, 0x59, 0x8d, 0x33, 0x70, 0xaa, 0x32, 0x74, 0x22, 0x91, 0x4c, 0x95,
	0x53, 0x95, 0x78, 0x06, 0x53, 0xaa, 0xf3, 0x6a, 0x17, 0x4e, 0x2c, 0x34, 0x14, 0x78, 0x0f, 0x52,
	0x7f, 0x71, 0x43, 0x26, 0x74, 0xa3, 0x49, 0xe2, 0x2f, 0xae, 0xd2, 0xc3, 0x81, 0x74, 0x58, 0x9e,
	0x6e, 0xfa, 0xf9, 0xdb, 0x77, 0x5d, 0x50, 0xab, 0x46, 0xf8, 0xe2, 0x03, 0xfc, 0x3f, 0x6d, 0x3c,
	0x07, 0xd9, 0x58, 0x35, 0x3a, 0x18, 0x2b, 0xbc, 0x05, 0xb7, 0xdb, 0x6b, 0xb2, 0x2e, 0x66, 0x8b,
	0xcb, 0x23, 0xbb, 0xdf, 0xf7, 0x9a, 0x94, 0x05, 0xe3, 0x1b, 0xf0, 0x0e, 0x2d, 0x04, 0x90, 0xeb,
	0x6c, 0xf9, 0xfc, 0xba, 0x0a, 0x4e, 0xf0, 0x14, 0xdc, 0xa7, 0x6c, 0xbd, 0x0a, 0x44, 0xaf, 0x3e,
	0x33, 0xf5, 0x12, 0x38, 0xf1, 0x03, 0xf8, 0x8f, 0x43, 0x40, 0x4b, 0xe6, 0x2d, 0x5e, 0x83, 0xd4,
	0xc4, 0x7a, 0xd7, 0xc7, 0xd0, 0x7f, 0x65, 0xfe, 0xef, 0x9c, 0x1a, 0x81, 0x42, 0xda, 0xc9, 0xdd,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xec, 0xf7, 0x43, 0x61, 0x01, 0x00, 0x00,
}
