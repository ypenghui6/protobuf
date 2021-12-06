// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue630.proto

package issue630

import (
	fmt "fmt"
	_ "github.com/ypenghui6/protobuf/gogoproto"
	proto "github.com/ypenghui6/protobuf/proto"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Foo struct {
	Bar1                 []Bar    `protobuf:"bytes,1,rep,name=Bar1" json:"Bar1"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_60374e06e51d301c, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetBar1() []Bar {
	if m != nil {
		return m.Bar1
	}
	return nil
}

type Qux struct {
	Bar1                 []*Bar   `protobuf:"bytes,1,rep,name=Bar1" json:"Bar1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Qux) Reset()         { *m = Qux{} }
func (m *Qux) String() string { return proto.CompactTextString(m) }
func (*Qux) ProtoMessage()    {}
func (*Qux) Descriptor() ([]byte, []int) {
	return fileDescriptor_60374e06e51d301c, []int{1}
}
func (m *Qux) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Qux.Unmarshal(m, b)
}
func (m *Qux) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Qux.Marshal(b, m, deterministic)
}
func (m *Qux) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Qux.Merge(m, src)
}
func (m *Qux) XXX_Size() int {
	return xxx_messageInfo_Qux.Size(m)
}
func (m *Qux) XXX_DiscardUnknown() {
	xxx_messageInfo_Qux.DiscardUnknown(m)
}

var xxx_messageInfo_Qux proto.InternalMessageInfo

func (m *Qux) GetBar1() []*Bar {
	if m != nil {
		return m.Bar1
	}
	return nil
}

type Bar struct {
	Baz                  string   `protobuf:"bytes,1,opt,name=Baz" json:"Baz"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_60374e06e51d301c, []int{2}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}
func (m *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(m, src)
}
func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetBaz() string {
	if m != nil {
		return m.Baz
	}
	return ""
}

func init() {
	proto.RegisterType((*Foo)(nil), "issue630.Foo")
	proto.RegisterType((*Qux)(nil), "issue630.Qux")
	proto.RegisterType((*Bar)(nil), "issue630.Bar")
}

func init() { proto.RegisterFile("issue630.proto", fileDescriptor_60374e06e51d301c) }

var fileDescriptor_60374e06e51d301c = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x33, 0x36, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54, 0xd2, 0xe3, 0x62,
	0x76, 0xcb, 0xcf, 0x17, 0x52, 0xe7, 0x62, 0x71, 0x4a, 0x2c, 0x32, 0x94, 0x60, 0x54, 0x60, 0xd6,
	0xe0, 0x36, 0xe2, 0xd5, 0x83, 0x1b, 0xef, 0x94, 0x58, 0xe4, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43,
	0x10, 0x58, 0x01, 0x48, 0x7d, 0x60, 0x69, 0x05, 0x61, 0xf5, 0x8c, 0x50, 0xf5, 0xb2, 0x5c, 0xcc,
	0x4e, 0x89, 0x45, 0x42, 0x62, 0x20, 0xaa, 0x4a, 0x82, 0x51, 0x81, 0x51, 0x83, 0x13, 0x6a, 0x1e,
	0x48, 0xc0, 0x89, 0xe5, 0xc3, 0x43, 0x39, 0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x5d,
	0xa4, 0x14, 0xce, 0x00, 0x00, 0x00,
}

func (this *Foo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&issue630.Foo{")
	if this.Bar1 != nil {
		vs := make([]Bar, len(this.Bar1))
		for i := range vs {
			vs[i] = this.Bar1[i]
		}
		s = append(s, "Bar1: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Qux) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&issue630.Qux{")
	if this.Bar1 != nil {
		s = append(s, "Bar1: "+fmt.Sprintf("%#v", this.Bar1)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Bar) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&issue630.Bar{")
	s = append(s, "Baz: "+fmt.Sprintf("%#v", this.Baz)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringIssue630(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
