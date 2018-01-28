// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sortkeys/sortable.proto

/*
Package sortkeys is a generated protocol buffer package.

It is generated from these files:
	sortkeys/sortable.proto

It has these top-level messages:
	Object
*/
package sortkeys

import proto "github.com/Cyinx/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Cyinx/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Object struct {
}

func (m *Object) Reset()                    { *m = Object{} }
func (m *Object) String() string            { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptorSortable, []int{0} }

func init() {
	proto.RegisterType((*Object)(nil), "sortkeys.Object")
}
func (this *Object) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Object)
	if !ok {
		that2, ok := that.(Object)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *Object) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&sortkeys.Object{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSortable(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedObject(r randySortable, easy bool) *Object {
	this := &Object{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randySortable interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSortable(r randySortable) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSortable(r randySortable) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneSortable(r)
	}
	return string(tmps)
}
func randUnrecognizedSortable(r randySortable, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSortable(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSortable(dAtA []byte, r randySortable, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSortable(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateSortable(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateSortable(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSortable(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSortable(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSortable(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSortable(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("sortkeys/sortable.proto", fileDescriptorSortable) }

var fileDescriptorSortable = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xce, 0x2f, 0x2a,
	0xc9, 0x4e, 0xad, 0x2c, 0xd6, 0x07, 0x31, 0x12, 0x93, 0x72, 0x52, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x38, 0x60, 0x12, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0x05, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03,
	0x66, 0x41, 0x34, 0x2a, 0x71, 0x70, 0xb1, 0xf9, 0x27, 0x65, 0xa5, 0x26, 0x97, 0x38, 0x09, 0x7c,
	0x78, 0x28, 0xc7, 0xf8, 0xe3, 0xa1, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e, 0xc9, 0x31,
	0x26, 0xb1, 0x81, 0x95, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x1b, 0xba, 0xd6, 0x76,
	0x00, 0x00, 0x00,
}
