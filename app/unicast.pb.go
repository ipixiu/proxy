// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: unicast.proto

/*
	Package main is a generated protocol buffer package.

	It is generated from these files:
		unicast.proto

	It has these top-level messages:
		UnicastMsg
		Notify
		ReuqestClientRecivePacketRq
		RequestClientRecivePacketRs
*/
package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// 使用者发送的单播协议
type UnicastMsg struct {
	UserId      uint64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id"`
	MsgType     int32  `protobuf:"varint,2,opt,name=msg_type,json=msgType" json:"msg_type"`
	MsgContent  []byte `protobuf:"bytes,3,opt,name=msg_content,json=msgContent" json:"msg_content"`
	MinProtocol int32  `protobuf:"varint,4,opt,name=min_protocol,json=minProtocol" json:"min_protocol"`
	AdaptMsg    []byte `protobuf:"bytes,5,opt,name=adapt_msg,json=adaptMsg" json:"adapt_msg"`
	// optional int64 adapt_sender = 6 [default = 10000];
	AdaptSender int64  `protobuf:"varint,6,opt,name=adapt_sender,json=adaptSender" json:"adapt_sender"`
	Ttl         int32  `protobuf:"varint,7,opt,name=ttl" json:"ttl"`
	UniqMark    []byte `protobuf:"bytes,8,opt,name=uniq_mark,json=uniqMark" json:"uniq_mark"`
	AllEs       int32  `protobuf:"varint,9,opt,name=all_es,json=allEs" json:"all_es"`
	Persistent  bool   `protobuf:"varint,10,opt,name=persistent" json:"persistent"`
	MsgId       uint64 `protobuf:"varint,11,opt,name=msg_id,json=msgId" json:"msg_id"`
}

func (m *UnicastMsg) Reset()                    { *m = UnicastMsg{} }
func (*UnicastMsg) ProtoMessage()               {}
func (*UnicastMsg) Descriptor() ([]byte, []int) { return fileDescriptorUnicast, []int{0} }

// msg_saver处理消息后通知broker协议
type Notify struct {
	UserId uint64 `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id"`
}

func (m *Notify) Reset()                    { *m = Notify{} }
func (*Notify) ProtoMessage()               {}
func (*Notify) Descriptor() ([]byte, []int) { return fileDescriptorUnicast, []int{1} }

// 发送给es转到客户端的包
type ReuqestClientRecivePacketRq struct {
	PacketId      int64  `protobuf:"varint,1,req,name=packet_id,json=packetId" json:"packet_id"`
	PacketType    int32  `protobuf:"varint,2,req,name=packet_type,json=packetType" json:"packet_type"`
	PacketContent []byte `protobuf:"bytes,3,req,name=packet_content,json=packetContent" json:"packet_content"`
}

func (m *ReuqestClientRecivePacketRq) Reset()      { *m = ReuqestClientRecivePacketRq{} }
func (*ReuqestClientRecivePacketRq) ProtoMessage() {}
func (*ReuqestClientRecivePacketRq) Descriptor() ([]byte, []int) {
	return fileDescriptorUnicast, []int{2}
}

// 应答
type RequestClientRecivePacketRs struct {
	PacketId int64 `protobuf:"varint,1,req,name=packet_id,json=packetId" json:"packet_id"`
}

func (m *RequestClientRecivePacketRs) Reset()      { *m = RequestClientRecivePacketRs{} }
func (*RequestClientRecivePacketRs) ProtoMessage() {}
func (*RequestClientRecivePacketRs) Descriptor() ([]byte, []int) {
	return fileDescriptorUnicast, []int{3}
}

func init() {
	proto.RegisterType((*UnicastMsg)(nil), "main.UnicastMsg")
	proto.RegisterType((*Notify)(nil), "main.Notify")
	proto.RegisterType((*ReuqestClientRecivePacketRq)(nil), "main.ReuqestClientRecivePacketRq")
	proto.RegisterType((*RequestClientRecivePacketRs)(nil), "main.RequestClientRecivePacketRs")
}
func (this *UnicastMsg) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*UnicastMsg)
	if !ok {
		that2, ok := that.(UnicastMsg)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *UnicastMsg")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *UnicastMsg but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *UnicastMsg but is not nil && this == nil")
	}
	if this.UserId != that1.UserId {
		return fmt.Errorf("UserId this(%v) Not Equal that(%v)", this.UserId, that1.UserId)
	}
	if this.MsgType != that1.MsgType {
		return fmt.Errorf("MsgType this(%v) Not Equal that(%v)", this.MsgType, that1.MsgType)
	}
	if !bytes.Equal(this.MsgContent, that1.MsgContent) {
		return fmt.Errorf("MsgContent this(%v) Not Equal that(%v)", this.MsgContent, that1.MsgContent)
	}
	if this.MinProtocol != that1.MinProtocol {
		return fmt.Errorf("MinProtocol this(%v) Not Equal that(%v)", this.MinProtocol, that1.MinProtocol)
	}
	if !bytes.Equal(this.AdaptMsg, that1.AdaptMsg) {
		return fmt.Errorf("AdaptMsg this(%v) Not Equal that(%v)", this.AdaptMsg, that1.AdaptMsg)
	}
	if this.AdaptSender != that1.AdaptSender {
		return fmt.Errorf("AdaptSender this(%v) Not Equal that(%v)", this.AdaptSender, that1.AdaptSender)
	}
	if this.Ttl != that1.Ttl {
		return fmt.Errorf("Ttl this(%v) Not Equal that(%v)", this.Ttl, that1.Ttl)
	}
	if !bytes.Equal(this.UniqMark, that1.UniqMark) {
		return fmt.Errorf("UniqMark this(%v) Not Equal that(%v)", this.UniqMark, that1.UniqMark)
	}
	if this.AllEs != that1.AllEs {
		return fmt.Errorf("AllEs this(%v) Not Equal that(%v)", this.AllEs, that1.AllEs)
	}
	if this.Persistent != that1.Persistent {
		return fmt.Errorf("Persistent this(%v) Not Equal that(%v)", this.Persistent, that1.Persistent)
	}
	if this.MsgId != that1.MsgId {
		return fmt.Errorf("MsgId this(%v) Not Equal that(%v)", this.MsgId, that1.MsgId)
	}
	return nil
}
func (this *UnicastMsg) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*UnicastMsg)
	if !ok {
		that2, ok := that.(UnicastMsg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.UserId != that1.UserId {
		return false
	}
	if this.MsgType != that1.MsgType {
		return false
	}
	if !bytes.Equal(this.MsgContent, that1.MsgContent) {
		return false
	}
	if this.MinProtocol != that1.MinProtocol {
		return false
	}
	if !bytes.Equal(this.AdaptMsg, that1.AdaptMsg) {
		return false
	}
	if this.AdaptSender != that1.AdaptSender {
		return false
	}
	if this.Ttl != that1.Ttl {
		return false
	}
	if !bytes.Equal(this.UniqMark, that1.UniqMark) {
		return false
	}
	if this.AllEs != that1.AllEs {
		return false
	}
	if this.Persistent != that1.Persistent {
		return false
	}
	if this.MsgId != that1.MsgId {
		return false
	}
	return true
}
func (this *Notify) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Notify)
	if !ok {
		that2, ok := that.(Notify)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Notify")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Notify but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Notify but is not nil && this == nil")
	}
	if this.UserId != that1.UserId {
		return fmt.Errorf("UserId this(%v) Not Equal that(%v)", this.UserId, that1.UserId)
	}
	return nil
}
func (this *Notify) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Notify)
	if !ok {
		that2, ok := that.(Notify)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.UserId != that1.UserId {
		return false
	}
	return true
}
func (this *ReuqestClientRecivePacketRq) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ReuqestClientRecivePacketRq)
	if !ok {
		that2, ok := that.(ReuqestClientRecivePacketRq)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ReuqestClientRecivePacketRq")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ReuqestClientRecivePacketRq but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ReuqestClientRecivePacketRq but is not nil && this == nil")
	}
	if this.PacketId != that1.PacketId {
		return fmt.Errorf("PacketId this(%v) Not Equal that(%v)", this.PacketId, that1.PacketId)
	}
	if this.PacketType != that1.PacketType {
		return fmt.Errorf("PacketType this(%v) Not Equal that(%v)", this.PacketType, that1.PacketType)
	}
	if !bytes.Equal(this.PacketContent, that1.PacketContent) {
		return fmt.Errorf("PacketContent this(%v) Not Equal that(%v)", this.PacketContent, that1.PacketContent)
	}
	return nil
}
func (this *ReuqestClientRecivePacketRq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ReuqestClientRecivePacketRq)
	if !ok {
		that2, ok := that.(ReuqestClientRecivePacketRq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.PacketId != that1.PacketId {
		return false
	}
	if this.PacketType != that1.PacketType {
		return false
	}
	if !bytes.Equal(this.PacketContent, that1.PacketContent) {
		return false
	}
	return true
}
func (this *RequestClientRecivePacketRs) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*RequestClientRecivePacketRs)
	if !ok {
		that2, ok := that.(RequestClientRecivePacketRs)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *RequestClientRecivePacketRs")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *RequestClientRecivePacketRs but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *RequestClientRecivePacketRs but is not nil && this == nil")
	}
	if this.PacketId != that1.PacketId {
		return fmt.Errorf("PacketId this(%v) Not Equal that(%v)", this.PacketId, that1.PacketId)
	}
	return nil
}
func (this *RequestClientRecivePacketRs) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RequestClientRecivePacketRs)
	if !ok {
		that2, ok := that.(RequestClientRecivePacketRs)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.PacketId != that1.PacketId {
		return false
	}
	return true
}
func (this *UnicastMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 15)
	s = append(s, "&main.UnicastMsg{")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",\n")
	s = append(s, "MsgType: "+fmt.Sprintf("%#v", this.MsgType)+",\n")
	s = append(s, "MsgContent: "+fmt.Sprintf("%#v", this.MsgContent)+",\n")
	s = append(s, "MinProtocol: "+fmt.Sprintf("%#v", this.MinProtocol)+",\n")
	s = append(s, "AdaptMsg: "+fmt.Sprintf("%#v", this.AdaptMsg)+",\n")
	s = append(s, "AdaptSender: "+fmt.Sprintf("%#v", this.AdaptSender)+",\n")
	s = append(s, "Ttl: "+fmt.Sprintf("%#v", this.Ttl)+",\n")
	s = append(s, "UniqMark: "+fmt.Sprintf("%#v", this.UniqMark)+",\n")
	s = append(s, "AllEs: "+fmt.Sprintf("%#v", this.AllEs)+",\n")
	s = append(s, "Persistent: "+fmt.Sprintf("%#v", this.Persistent)+",\n")
	s = append(s, "MsgId: "+fmt.Sprintf("%#v", this.MsgId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Notify) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.Notify{")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReuqestClientRecivePacketRq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&main.ReuqestClientRecivePacketRq{")
	s = append(s, "PacketId: "+fmt.Sprintf("%#v", this.PacketId)+",\n")
	s = append(s, "PacketType: "+fmt.Sprintf("%#v", this.PacketType)+",\n")
	s = append(s, "PacketContent: "+fmt.Sprintf("%#v", this.PacketContent)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RequestClientRecivePacketRs) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&main.RequestClientRecivePacketRs{")
	s = append(s, "PacketId: "+fmt.Sprintf("%#v", this.PacketId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUnicast(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UnicastMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnicastMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.UserId))
	dAtA[i] = 0x10
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.MsgType))
	if m.MsgContent != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUnicast(dAtA, i, uint64(len(m.MsgContent)))
		i += copy(dAtA[i:], m.MsgContent)
	}
	dAtA[i] = 0x20
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.MinProtocol))
	if m.AdaptMsg != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintUnicast(dAtA, i, uint64(len(m.AdaptMsg)))
		i += copy(dAtA[i:], m.AdaptMsg)
	}
	dAtA[i] = 0x30
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.AdaptSender))
	dAtA[i] = 0x38
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.Ttl))
	if m.UniqMark != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintUnicast(dAtA, i, uint64(len(m.UniqMark)))
		i += copy(dAtA[i:], m.UniqMark)
	}
	dAtA[i] = 0x48
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.AllEs))
	dAtA[i] = 0x50
	i++
	if m.Persistent {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x58
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.MsgId))
	return i, nil
}

func (m *Notify) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Notify) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.UserId))
	return i, nil
}

func (m *ReuqestClientRecivePacketRq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReuqestClientRecivePacketRq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.PacketId))
	dAtA[i] = 0x10
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.PacketType))
	if m.PacketContent != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUnicast(dAtA, i, uint64(len(m.PacketContent)))
		i += copy(dAtA[i:], m.PacketContent)
	}
	return i, nil
}

func (m *RequestClientRecivePacketRs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestClientRecivePacketRs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintUnicast(dAtA, i, uint64(m.PacketId))
	return i, nil
}

func encodeFixed64Unicast(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Unicast(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintUnicast(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UnicastMsg) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovUnicast(uint64(m.UserId))
	n += 1 + sovUnicast(uint64(m.MsgType))
	if m.MsgContent != nil {
		l = len(m.MsgContent)
		n += 1 + l + sovUnicast(uint64(l))
	}
	n += 1 + sovUnicast(uint64(m.MinProtocol))
	if m.AdaptMsg != nil {
		l = len(m.AdaptMsg)
		n += 1 + l + sovUnicast(uint64(l))
	}
	n += 1 + sovUnicast(uint64(m.AdaptSender))
	n += 1 + sovUnicast(uint64(m.Ttl))
	if m.UniqMark != nil {
		l = len(m.UniqMark)
		n += 1 + l + sovUnicast(uint64(l))
	}
	n += 1 + sovUnicast(uint64(m.AllEs))
	n += 2
	n += 1 + sovUnicast(uint64(m.MsgId))
	return n
}

func (m *Notify) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovUnicast(uint64(m.UserId))
	return n
}

func (m *ReuqestClientRecivePacketRq) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovUnicast(uint64(m.PacketId))
	n += 1 + sovUnicast(uint64(m.PacketType))
	if m.PacketContent != nil {
		l = len(m.PacketContent)
		n += 1 + l + sovUnicast(uint64(l))
	}
	return n
}

func (m *RequestClientRecivePacketRs) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovUnicast(uint64(m.PacketId))
	return n
}

func sovUnicast(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUnicast(x uint64) (n int) {
	return sovUnicast(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UnicastMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnicastMsg{`,
		`UserId:` + fmt.Sprintf("%v", this.UserId) + `,`,
		`MsgType:` + fmt.Sprintf("%v", this.MsgType) + `,`,
		`MsgContent:` + fmt.Sprintf("%v", this.MsgContent) + `,`,
		`MinProtocol:` + fmt.Sprintf("%v", this.MinProtocol) + `,`,
		`AdaptMsg:` + fmt.Sprintf("%v", this.AdaptMsg) + `,`,
		`AdaptSender:` + fmt.Sprintf("%v", this.AdaptSender) + `,`,
		`Ttl:` + fmt.Sprintf("%v", this.Ttl) + `,`,
		`UniqMark:` + fmt.Sprintf("%v", this.UniqMark) + `,`,
		`AllEs:` + fmt.Sprintf("%v", this.AllEs) + `,`,
		`Persistent:` + fmt.Sprintf("%v", this.Persistent) + `,`,
		`MsgId:` + fmt.Sprintf("%v", this.MsgId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Notify) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Notify{`,
		`UserId:` + fmt.Sprintf("%v", this.UserId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReuqestClientRecivePacketRq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReuqestClientRecivePacketRq{`,
		`PacketId:` + fmt.Sprintf("%v", this.PacketId) + `,`,
		`PacketType:` + fmt.Sprintf("%v", this.PacketType) + `,`,
		`PacketContent:` + fmt.Sprintf("%v", this.PacketContent) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RequestClientRecivePacketRs) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RequestClientRecivePacketRs{`,
		`PacketId:` + fmt.Sprintf("%v", this.PacketId) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUnicast(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UnicastMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnicast
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnicastMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnicastMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			m.MsgType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgContent", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUnicast
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgContent = append(m.MsgContent[:0], dAtA[iNdEx:postIndex]...)
			if m.MsgContent == nil {
				m.MsgContent = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinProtocol", wireType)
			}
			m.MinProtocol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinProtocol |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdaptMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUnicast
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdaptMsg = append(m.AdaptMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.AdaptMsg == nil {
				m.AdaptMsg = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdaptSender", wireType)
			}
			m.AdaptSender = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdaptSender |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniqMark", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUnicast
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniqMark = append(m.UniqMark[:0], dAtA[iNdEx:postIndex]...)
			if m.UniqMark == nil {
				m.UniqMark = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllEs", wireType)
			}
			m.AllEs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllEs |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Persistent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Persistent = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgId", wireType)
			}
			m.MsgId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnicast(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnicast
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Notify) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnicast
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Notify: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Notify: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipUnicast(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnicast
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("user_id")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReuqestClientRecivePacketRq) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnicast
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReuqestClientRecivePacketRq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReuqestClientRecivePacketRq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			m.PacketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketType", wireType)
			}
			m.PacketType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketContent", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUnicast
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketContent = append(m.PacketContent[:0], dAtA[iNdEx:postIndex]...)
			if m.PacketContent == nil {
				m.PacketContent = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipUnicast(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnicast
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("packet_id")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("packet_type")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("packet_content")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestClientRecivePacketRs) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnicast
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestClientRecivePacketRs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestClientRecivePacketRs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			m.PacketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipUnicast(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnicast
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("packet_id")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUnicast(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUnicast
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnicast
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUnicast
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUnicast
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUnicast(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUnicast = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUnicast   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("unicast.proto", fileDescriptorUnicast) }

var fileDescriptorUnicast = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x3f, 0x6f, 0xd4, 0x30,
	0x1c, 0x8d, 0xef, 0x6f, 0xfa, 0xbb, 0x2b, 0x43, 0x06, 0x64, 0x51, 0xe1, 0x86, 0x13, 0xe8, 0x22,
	0x21, 0xda, 0x85, 0x0f, 0x80, 0x5a, 0x31, 0xdc, 0x50, 0x54, 0x1d, 0x30, 0x47, 0x69, 0xe2, 0x1a,
	0xeb, 0x62, 0x27, 0x17, 0x3b, 0x48, 0xb7, 0x21, 0xf1, 0x05, 0x18, 0xf8, 0x10, 0x7c, 0x94, 0x1b,
	0x3b, 0x32, 0xa1, 0x26, 0x2c, 0x8c, 0xfd, 0x08, 0xc8, 0x4e, 0xae, 0x17, 0x10, 0x03, 0x9b, 0xfd,
	0xde, 0xf3, 0xf3, 0xfb, 0xfd, 0x1e, 0x1c, 0x96, 0x92, 0xc7, 0x91, 0xd2, 0x27, 0x79, 0x91, 0xe9,
	0xcc, 0x1b, 0x88, 0x88, 0xcb, 0x47, 0x2f, 0x18, 0xd7, 0x1f, 0xca, 0xab, 0x93, 0x38, 0x13, 0xa7,
	0x2c, 0x63, 0xd9, 0xa9, 0x25, 0xaf, 0xca, 0x6b, 0x7b, 0xb3, 0x17, 0x7b, 0x6a, 0x1e, 0xcd, 0x3e,
	0xf7, 0x01, 0xde, 0x37, 0x36, 0x17, 0x8a, 0x79, 0x8f, 0x61, 0x5c, 0x2a, 0x5a, 0x84, 0x3c, 0xc1,
	0xc8, 0x47, 0xc1, 0xe0, 0x6c, 0xb0, 0xfd, 0x71, 0xec, 0x2c, 0x47, 0x06, 0x5c, 0x24, 0xde, 0x31,
	0xb8, 0x42, 0xb1, 0x50, 0x6f, 0x72, 0x8a, 0x7b, 0x3e, 0x0a, 0x86, 0x2d, 0x3f, 0x16, 0x8a, 0xbd,
	0xdb, 0xe4, 0xd4, 0x7b, 0x06, 0x13, 0x23, 0x88, 0x33, 0xa9, 0xa9, 0xd4, 0xb8, 0xef, 0xa3, 0x60,
	0xda, 0x6a, 0x40, 0x28, 0x76, 0xde, 0xe0, 0xde, 0x1c, 0xa6, 0x82, 0xcb, 0xd0, 0x46, 0x88, 0xb3,
	0x14, 0x0f, 0x3a, 0x5e, 0x13, 0xc1, 0xe5, 0x65, 0x4b, 0x78, 0x4f, 0xe0, 0x20, 0x4a, 0xa2, 0x5c,
	0x87, 0x42, 0x31, 0x3c, 0xec, 0xb8, 0xb9, 0x16, 0x36, 0x91, 0xe7, 0x30, 0x6d, 0x24, 0x8a, 0xca,
	0x84, 0x16, 0x78, 0xe4, 0xa3, 0xa0, 0xbf, 0xf3, 0xb2, 0xcc, 0x5b, 0x4b, 0x78, 0x0f, 0xa1, 0xaf,
	0x75, 0x8a, 0xc7, 0x9d, 0xbf, 0x0c, 0x60, 0xfe, 0x28, 0x25, 0x5f, 0x87, 0x22, 0x2a, 0x56, 0xd8,
	0xed, 0xfe, 0x61, 0xe0, 0x8b, 0xa8, 0x58, 0x79, 0x47, 0x30, 0x8a, 0xd2, 0x34, 0xa4, 0x0a, 0x1f,
	0x74, 0x5e, 0x0f, 0xa3, 0x34, 0x7d, 0xad, 0xbc, 0xa7, 0x00, 0x39, 0x2d, 0x14, 0x57, 0x76, 0x64,
	0xf0, 0x51, 0xe0, 0xee, 0x46, 0xde, 0xe3, 0xc6, 0xc2, 0x6c, 0x86, 0x27, 0x78, 0xd2, 0x59, 0xec,
	0x50, 0x28, 0xb6, 0x48, 0x66, 0x73, 0x18, 0xbd, 0xc9, 0x34, 0xbf, 0xde, 0xfc, 0x59, 0x40, 0xef,
	0xef, 0x02, 0x66, 0x5f, 0x11, 0x1c, 0x2d, 0x69, 0xb9, 0xa6, 0x4a, 0x9f, 0xa7, 0x9c, 0x4a, 0xbd,
	0xa4, 0x31, 0xff, 0x48, 0x2f, 0xa3, 0x78, 0x45, 0xf5, 0x72, 0x6d, 0x66, 0xc9, 0xed, 0x79, 0x67,
	0xb0, 0xdb, 0x84, 0xdb, 0xc0, 0x8b, 0xc4, 0x54, 0xd4, 0x4a, 0xda, 0x1a, 0x7b, 0xf7, 0x03, 0x41,
	0x43, 0xd8, 0x26, 0x9f, 0xc3, 0x83, 0x56, 0xb6, 0x2f, 0xb3, 0x77, 0xbf, 0x9a, 0xc3, 0x86, 0x6b,
	0xfb, 0x9c, 0xbd, 0x32, 0xa9, 0xd6, 0xe5, 0xbf, 0x53, 0xa9, 0xff, 0x48, 0x75, 0xf6, 0x72, 0x5b,
	0x11, 0xe7, 0xa6, 0x22, 0xce, 0xf7, 0x8a, 0x38, 0xb7, 0x15, 0x41, 0x77, 0x15, 0x41, 0x9f, 0x6a,
	0x82, 0xbe, 0xd5, 0x04, 0x6d, 0x6b, 0x82, 0x6e, 0x6a, 0x82, 0x6e, 0x6b, 0x82, 0x7e, 0xd5, 0xc4,
	0xb9, 0xab, 0x09, 0xfa, 0xf2, 0x93, 0x38, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x33, 0xaf, 0x8e,
	0x2d, 0x02, 0x03, 0x00, 0x00,
}