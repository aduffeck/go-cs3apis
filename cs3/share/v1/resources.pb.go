// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/share/v1/resources.proto

package sharev1pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TargetType int32

const (
	TargetType_TARGET_TYPE_INVALID TargetType = 0
	TargetType_TARGET_TYPE_USER    TargetType = 1
	TargetType_TARGET_TYPE_GROUP   TargetType = 2
)

var TargetType_name = map[int32]string{
	0: "TARGET_TYPE_INVALID",
	1: "TARGET_TYPE_USER",
	2: "TARGET_TYPE_GROUP",
}

var TargetType_value = map[string]int32{
	"TARGET_TYPE_INVALID": 0,
	"TARGET_TYPE_USER":    1,
	"TARGET_TYPE_GROUP":   2,
}

func (x TargetType) String() string {
	return proto.EnumName(TargetType_name, int32(x))
}

func (TargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7eb1d0b8aa21ab13, []int{0}
}

type ReceivedShareState int32

const (
	ReceivedShareState_RECEIVED_SHARE_STATE_INVALID  ReceivedShareState = 0
	ReceivedShareState_RECEIVED_SHARE_STATE_ACCEPTED ReceivedShareState = 1
	ReceivedShareState_RECEIVED_SHARE_STATE_REJECTED ReceivedShareState = 2
	ReceivedShareState_RECEIVED_SHARE_STATE_PENDING  ReceivedShareState = 3
)

var ReceivedShareState_name = map[int32]string{
	0: "RECEIVED_SHARE_STATE_INVALID",
	1: "RECEIVED_SHARE_STATE_ACCEPTED",
	2: "RECEIVED_SHARE_STATE_REJECTED",
	3: "RECEIVED_SHARE_STATE_PENDING",
}

var ReceivedShareState_value = map[string]int32{
	"RECEIVED_SHARE_STATE_INVALID":  0,
	"RECEIVED_SHARE_STATE_ACCEPTED": 1,
	"RECEIVED_SHARE_STATE_REJECTED": 2,
	"RECEIVED_SHARE_STATE_PENDING":  3,
}

func (x ReceivedShareState) String() string {
	return proto.EnumName(ReceivedShareState_name, int32(x))
}

func (ReceivedShareState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7eb1d0b8aa21ab13, []int{1}
}

type Permissions struct {
	Read                 bool     `protobuf:"varint,1,opt,name=read,proto3" json:"read,omitempty"`
	Modify               bool     `protobuf:"varint,2,opt,name=modify,proto3" json:"modify,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb1d0b8aa21ab13, []int{0}
}

func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *Permissions) GetModify() bool {
	if m != nil {
		return m.Modify
	}
	return false
}

type Share struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename             string       `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Target               string       `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	TargetType           TargetType   `protobuf:"varint,4,opt,name=target_type,json=targetType,proto3,enum=cs3.sharev1.TargetType" json:"target_type,omitempty"`
	Permissions          *Permissions `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Ctime                uint64       `protobuf:"varint,6,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                uint64       `protobuf:"varint,7,opt,name=mtime,proto3" json:"mtime,omitempty"`
	DisplayName          string       `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Owner                string       `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Share) Reset()         { *m = Share{} }
func (m *Share) String() string { return proto.CompactTextString(m) }
func (*Share) ProtoMessage()    {}
func (*Share) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb1d0b8aa21ab13, []int{1}
}

func (m *Share) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Share.Unmarshal(m, b)
}
func (m *Share) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Share.Marshal(b, m, deterministic)
}
func (m *Share) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Share.Merge(m, src)
}
func (m *Share) XXX_Size() int {
	return xxx_messageInfo_Share.Size(m)
}
func (m *Share) XXX_DiscardUnknown() {
	xxx_messageInfo_Share.DiscardUnknown(m)
}

var xxx_messageInfo_Share proto.InternalMessageInfo

func (m *Share) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Share) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Share) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Share) GetTargetType() TargetType {
	if m != nil {
		return m.TargetType
	}
	return TargetType_TARGET_TYPE_INVALID
}

func (m *Share) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Share) GetCtime() uint64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *Share) GetMtime() uint64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *Share) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Share) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type UpdatePolicy struct {
	UpdatePermissions    bool     `protobuf:"varint,1,opt,name=update_permissions,json=updatePermissions,proto3" json:"update_permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePolicy) Reset()         { *m = UpdatePolicy{} }
func (m *UpdatePolicy) String() string { return proto.CompactTextString(m) }
func (*UpdatePolicy) ProtoMessage()    {}
func (*UpdatePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb1d0b8aa21ab13, []int{2}
}

func (m *UpdatePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePolicy.Unmarshal(m, b)
}
func (m *UpdatePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePolicy.Marshal(b, m, deterministic)
}
func (m *UpdatePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePolicy.Merge(m, src)
}
func (m *UpdatePolicy) XXX_Size() int {
	return xxx_messageInfo_UpdatePolicy.Size(m)
}
func (m *UpdatePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePolicy proto.InternalMessageInfo

func (m *UpdatePolicy) GetUpdatePermissions() bool {
	if m != nil {
		return m.UpdatePermissions
	}
	return false
}

type ReceivedShare struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename             string             `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Target               string             `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	TargetType           TargetType         `protobuf:"varint,4,opt,name=target_type,json=targetType,proto3,enum=cs3.sharev1.TargetType" json:"target_type,omitempty"`
	Permissions          *Permissions       `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Ctime                uint64             `protobuf:"varint,6,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                uint64             `protobuf:"varint,7,opt,name=mtime,proto3" json:"mtime,omitempty"`
	DisplayName          string             `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Owner                string             `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	State                ReceivedShareState `protobuf:"varint,10,opt,name=state,proto3,enum=cs3.sharev1.ReceivedShareState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReceivedShare) Reset()         { *m = ReceivedShare{} }
func (m *ReceivedShare) String() string { return proto.CompactTextString(m) }
func (*ReceivedShare) ProtoMessage()    {}
func (*ReceivedShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb1d0b8aa21ab13, []int{3}
}

func (m *ReceivedShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedShare.Unmarshal(m, b)
}
func (m *ReceivedShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedShare.Marshal(b, m, deterministic)
}
func (m *ReceivedShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedShare.Merge(m, src)
}
func (m *ReceivedShare) XXX_Size() int {
	return xxx_messageInfo_ReceivedShare.Size(m)
}
func (m *ReceivedShare) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedShare.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedShare proto.InternalMessageInfo

func (m *ReceivedShare) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReceivedShare) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *ReceivedShare) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReceivedShare) GetTargetType() TargetType {
	if m != nil {
		return m.TargetType
	}
	return TargetType_TARGET_TYPE_INVALID
}

func (m *ReceivedShare) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *ReceivedShare) GetCtime() uint64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *ReceivedShare) GetMtime() uint64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *ReceivedShare) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ReceivedShare) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReceivedShare) GetState() ReceivedShareState {
	if m != nil {
		return m.State
	}
	return ReceivedShareState_RECEIVED_SHARE_STATE_INVALID
}

func init() {
	proto.RegisterEnum("cs3.sharev1.TargetType", TargetType_name, TargetType_value)
	proto.RegisterEnum("cs3.sharev1.ReceivedShareState", ReceivedShareState_name, ReceivedShareState_value)
	proto.RegisterType((*Permissions)(nil), "cs3.sharev1.Permissions")
	proto.RegisterType((*Share)(nil), "cs3.sharev1.Share")
	proto.RegisterType((*UpdatePolicy)(nil), "cs3.sharev1.UpdatePolicy")
	proto.RegisterType((*ReceivedShare)(nil), "cs3.sharev1.ReceivedShare")
}

func init() { proto.RegisterFile("cs3/share/v1/resources.proto", fileDescriptor_7eb1d0b8aa21ab13) }

var fileDescriptor_7eb1d0b8aa21ab13 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x36, 0xd9, 0xb6, 0xb6, 0x27, 0x6b, 0xcd, 0x8e, 0xab, 0x1b, 0x64, 0xc5, 0x6c, 0xaf, 0xca,
	0x82, 0x29, 0xdd, 0x22, 0xa8, 0xe0, 0x45, 0x9a, 0x0e, 0xb5, 0x22, 0xdd, 0x30, 0x49, 0x8b, 0x8a,
	0x10, 0xb2, 0xc9, 0xac, 0x0e, 0x34, 0x4d, 0xc8, 0x64, 0x2b, 0x79, 0x19, 0x2f, 0xbc, 0xf0, 0xc2,
	0x47, 0xf1, 0x3d, 0x7c, 0x0f, 0xc9, 0x24, 0xee, 0xa6, 0x68, 0xf1, 0x05, 0xbc, 0x9b, 0xef, 0xe7,
	0xcc, 0x39, 0xdf, 0xb9, 0x38, 0x70, 0x1c, 0xf0, 0xd1, 0x80, 0x7f, 0xf2, 0x53, 0x3a, 0xd8, 0x0c,
	0x07, 0x29, 0xe5, 0xf1, 0x55, 0x1a, 0x50, 0x6e, 0x24, 0x69, 0x9c, 0xc5, 0x48, 0x09, 0xf8, 0xc8,
	0x10, 0xea, 0x66, 0xd8, 0x7b, 0x0e, 0x8a, 0x4d, 0xd3, 0x88, 0x71, 0xce, 0xe2, 0x35, 0x47, 0x08,
	0x1a, 0x29, 0xf5, 0x43, 0x4d, 0xd2, 0xa5, 0x7e, 0x9b, 0x88, 0x37, 0x7a, 0x00, 0xad, 0x28, 0x0e,
	0xd9, 0x65, 0xae, 0xc9, 0x82, 0xad, 0x50, 0xef, 0x9b, 0x0c, 0x4d, 0xa7, 0xf8, 0x06, 0x75, 0x41,
	0x66, 0x65, 0x4d, 0x87, 0xc8, 0x2c, 0x44, 0x0f, 0xa1, 0x7d, 0xc9, 0x56, 0x74, 0xed, 0x47, 0x54,
	0xd4, 0x74, 0xc8, 0x35, 0x2e, 0x7e, 0xcb, 0xfc, 0xf4, 0x23, 0xcd, 0xb4, 0x3d, 0xa1, 0x54, 0x08,
	0x3d, 0x03, 0xa5, 0x7c, 0x79, 0x59, 0x9e, 0x50, 0xad, 0xa1, 0x4b, 0xfd, 0xee, 0xd9, 0x91, 0x51,
	0x9b, 0xd5, 0x70, 0x85, 0xee, 0xe6, 0x09, 0x25, 0x90, 0x5d, 0xbf, 0xd1, 0x0b, 0x50, 0x92, 0x9b,
	0x08, 0x5a, 0x53, 0x97, 0xfa, 0xca, 0x99, 0xb6, 0x55, 0x59, 0x8b, 0x48, 0xea, 0x66, 0x74, 0x08,
	0xcd, 0x20, 0x63, 0x11, 0xd5, 0x5a, 0xba, 0xd4, 0x6f, 0x90, 0x12, 0x14, 0x6c, 0x24, 0xd8, 0xdb,
	0x25, 0x2b, 0x00, 0x3a, 0x81, 0xfd, 0x90, 0xf1, 0x64, 0xe5, 0xe7, 0x9e, 0x48, 0xd6, 0x16, 0xf3,
	0x2b, 0x15, 0x37, 0xf7, 0xcb, 0xc2, 0xf8, 0xf3, 0x9a, 0xa6, 0x5a, 0x47, 0x68, 0x25, 0xe8, 0xbd,
	0x84, 0xfd, 0x45, 0x12, 0xfa, 0x19, 0xb5, 0xe3, 0x15, 0x0b, 0x72, 0xf4, 0x04, 0xd0, 0x95, 0xc0,
	0x5e, 0x7d, 0xee, 0x72, 0xe5, 0x07, 0xa5, 0x52, 0x1b, 0xb8, 0xf7, 0x53, 0x86, 0x3b, 0x84, 0x06,
	0x94, 0x6d, 0x68, 0xf8, 0x7f, 0xdf, 0xff, 0xde, 0x37, 0x7a, 0x0a, 0x4d, 0x9e, 0xf9, 0x19, 0xd5,
	0x40, 0x84, 0x7a, 0xbc, 0x35, 0xda, 0xd6, 0x26, 0x9d, 0xc2, 0x46, 0x4a, 0xf7, 0x29, 0x01, 0xb8,
	0x49, 0x8c, 0x8e, 0xe0, 0x9e, 0x6b, 0x92, 0x29, 0x76, 0x3d, 0xf7, 0x9d, 0x8d, 0xbd, 0xd9, 0x7c,
	0x69, 0xbe, 0x99, 0x4d, 0xd4, 0x5b, 0xe8, 0x10, 0xd4, 0xba, 0xb0, 0x70, 0x30, 0x51, 0x25, 0x74,
	0x1f, 0x0e, 0xea, 0xec, 0x94, 0x9c, 0x2f, 0x6c, 0x55, 0x3e, 0xfd, 0x22, 0x01, 0xfa, 0xb3, 0x23,
	0xd2, 0xe1, 0x98, 0x60, 0x0b, 0xcf, 0x96, 0x78, 0xe2, 0x39, 0xaf, 0x4c, 0x82, 0x3d, 0xc7, 0x35,
	0xdd, 0x7a, 0x97, 0x13, 0x78, 0xf4, 0x57, 0x87, 0x69, 0x59, 0xd8, 0x76, 0xf1, 0x44, 0x95, 0x76,
	0x5a, 0x08, 0x7e, 0x8d, 0xad, 0xc2, 0x22, 0xef, 0xec, 0x63, 0xe3, 0xf9, 0x64, 0x36, 0x9f, 0xaa,
	0x7b, 0x63, 0x07, 0xee, 0x06, 0x71, 0x54, 0xdf, 0xd0, 0xb8, 0x4b, 0x7e, 0x1f, 0x0c, 0xbb, 0xb8,
	0x17, 0xb6, 0xf4, 0xbe, 0x53, 0x49, 0xc9, 0xc5, 0x57, 0xb9, 0x65, 0x8d, 0xcf, 0xdf, 0x9a, 0xe3,
	0xef, 0xb2, 0x62, 0x39, 0x23, 0x43, 0x64, 0x5a, 0x0e, 0x7f, 0x08, 0xf4, 0xa1, 0x42, 0x17, 0x2d,
	0x71, 0x68, 0x46, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xf3, 0x20, 0x64, 0x88, 0x04, 0x00,
	0x00,
}
