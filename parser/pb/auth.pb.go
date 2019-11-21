// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package pb

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

// 区服状态
type ServerStatus int32

const (
	ServerStatus_Good  ServerStatus = 0
	ServerStatus_Hot   ServerStatus = 1
	ServerStatus_Fix   ServerStatus = 2
	ServerStatus_Close ServerStatus = 3
)

var ServerStatus_name = map[int32]string{
	0: "Good",
	1: "Hot",
	2: "Fix",
	3: "Close",
}
var ServerStatus_value = map[string]int32{
	"Good":  0,
	"Hot":   1,
	"Fix":   2,
	"Close": 3,
}

func (x ServerStatus) String() string {
	return proto.EnumName(ServerStatus_name, int32(x))
}
func (ServerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{0}
}

type SDKChannel int32

const (
	SDKChannel_GM   SDKChannel = 0
	SDKChannel_Mine SDKChannel = 1
	SDKChannel_DeNA SDKChannel = 2
)

var SDKChannel_name = map[int32]string{
	0: "GM",
	1: "Mine",
	2: "DeNA",
}
var SDKChannel_value = map[string]int32{
	"GM":   0,
	"Mine": 1,
	"DeNA": 2,
}

func (x SDKChannel) String() string {
	return proto.EnumName(SDKChannel_name, int32(x))
}
func (SDKChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{1}
}

type OrderStatus int32

const (
	OrderStatus_Notify OrderStatus = 0
	OrderStatus_Finish OrderStatus = 1
)

var OrderStatus_name = map[int32]string{
	0: "Notify",
	1: "Finish",
}
var OrderStatus_value = map[string]int32{
	"Notify": 0,
	"Finish": 1,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{2}
}

type UserInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{0}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *UserInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserRoleInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Server               int32    `protobuf:"varint,3,opt,name=server" json:"server,omitempty"`
	Ctime                int64    `protobuf:"varint,4,opt,name=ctime" json:"ctime,omitempty"`
	Uid                  string   `protobuf:"bytes,5,opt,name=uid" json:"uid,omitempty"`
	TypeId               int32    `protobuf:"varint,6,opt,name=typeId" json:"typeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRoleInfo) Reset()         { *m = UserRoleInfo{} }
func (m *UserRoleInfo) String() string { return proto.CompactTextString(m) }
func (*UserRoleInfo) ProtoMessage()    {}
func (*UserRoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{1}
}
func (m *UserRoleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRoleInfo.Unmarshal(m, b)
}
func (m *UserRoleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRoleInfo.Marshal(b, m, deterministic)
}
func (dst *UserRoleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRoleInfo.Merge(dst, src)
}
func (m *UserRoleInfo) XXX_Size() int {
	return xxx_messageInfo_UserRoleInfo.Size(m)
}
func (m *UserRoleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRoleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserRoleInfo proto.InternalMessageInfo

func (m *UserRoleInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRoleInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRoleInfo) GetServer() int32 {
	if m != nil {
		return m.Server
	}
	return 0
}

func (m *UserRoleInfo) GetCtime() int64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *UserRoleInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserRoleInfo) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type LogicServerInfo struct {
	Id                   int32        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Start                int64        `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
	End                  int64        `protobuf:"varint,4,opt,name=end" json:"end,omitempty"`
	Time                 int64        `protobuf:"varint,5,opt,name=time" json:"time,omitempty"`
	State                ServerStatus `protobuf:"varint,6,opt,name=state,enum=ServerStatus" json:"state,omitempty"`
	CycleStart           int64        `protobuf:"varint,7,opt,name=cycleStart" json:"cycleStart,omitempty"`
	CycleEnd             int64        `protobuf:"varint,8,opt,name=cycleEnd" json:"cycleEnd,omitempty"`
	IsCycle              bool         `protobuf:"varint,9,opt,name=isCycle" json:"isCycle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LogicServerInfo) Reset()         { *m = LogicServerInfo{} }
func (m *LogicServerInfo) String() string { return proto.CompactTextString(m) }
func (*LogicServerInfo) ProtoMessage()    {}
func (*LogicServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{2}
}
func (m *LogicServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicServerInfo.Unmarshal(m, b)
}
func (m *LogicServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicServerInfo.Marshal(b, m, deterministic)
}
func (dst *LogicServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicServerInfo.Merge(dst, src)
}
func (m *LogicServerInfo) XXX_Size() int {
	return xxx_messageInfo_LogicServerInfo.Size(m)
}
func (m *LogicServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogicServerInfo proto.InternalMessageInfo

func (m *LogicServerInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LogicServerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogicServerInfo) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *LogicServerInfo) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *LogicServerInfo) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *LogicServerInfo) GetState() ServerStatus {
	if m != nil {
		return m.State
	}
	return ServerStatus_Good
}

func (m *LogicServerInfo) GetCycleStart() int64 {
	if m != nil {
		return m.CycleStart
	}
	return 0
}

func (m *LogicServerInfo) GetCycleEnd() int64 {
	if m != nil {
		return m.CycleEnd
	}
	return 0
}

func (m *LogicServerInfo) GetIsCycle() bool {
	if m != nil {
		return m.IsCycle
	}
	return false
}

type SuperAccount struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Gid                  int32    `protobuf:"varint,2,opt,name=gid" json:"gid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuperAccount) Reset()         { *m = SuperAccount{} }
func (m *SuperAccount) String() string { return proto.CompactTextString(m) }
func (*SuperAccount) ProtoMessage()    {}
func (*SuperAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{3}
}
func (m *SuperAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuperAccount.Unmarshal(m, b)
}
func (m *SuperAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuperAccount.Marshal(b, m, deterministic)
}
func (dst *SuperAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperAccount.Merge(dst, src)
}
func (m *SuperAccount) XXX_Size() int {
	return xxx_messageInfo_SuperAccount.Size(m)
}
func (m *SuperAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SuperAccount proto.InternalMessageInfo

func (m *SuperAccount) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SuperAccount) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

// 排队
type Line struct {
	Lsid                 int32    `protobuf:"varint,1,opt,name=lsid" json:"lsid,omitempty"`
	On                   bool     `protobuf:"varint,2,opt,name=on" json:"on,omitempty"`
	Min                  int32    `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
	Max                  int32    `protobuf:"varint,4,opt,name=max" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Line) Reset()         { *m = Line{} }
func (m *Line) String() string { return proto.CompactTextString(m) }
func (*Line) ProtoMessage()    {}
func (*Line) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{4}
}
func (m *Line) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Line.Unmarshal(m, b)
}
func (m *Line) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Line.Marshal(b, m, deterministic)
}
func (dst *Line) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Line.Merge(dst, src)
}
func (m *Line) XXX_Size() int {
	return xxx_messageInfo_Line.Size(m)
}
func (m *Line) XXX_DiscardUnknown() {
	xxx_messageInfo_Line.DiscardUnknown(m)
}

var xxx_messageInfo_Line proto.InternalMessageInfo

func (m *Line) GetLsid() int32 {
	if m != nil {
		return m.Lsid
	}
	return 0
}

func (m *Line) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *Line) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Line) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

// 内部订单通知信息 根据SDK推送的订单信息生成
type OrderNotify struct {
	Gid                  int32       `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	Oid                  string      `protobuf:"bytes,2,opt,name=oid" json:"oid,omitempty"`
	Channel              SDKChannel  `protobuf:"varint,3,opt,name=channel,enum=SDKChannel" json:"channel,omitempty"`
	Money                int32       `protobuf:"varint,4,opt,name=money" json:"money,omitempty"`
	Time                 int64       `protobuf:"varint,5,opt,name=time" json:"time,omitempty"`
	Status               OrderStatus `protobuf:"varint,6,opt,name=status,enum=OrderStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OrderNotify) Reset()         { *m = OrderNotify{} }
func (m *OrderNotify) String() string { return proto.CompactTextString(m) }
func (*OrderNotify) ProtoMessage()    {}
func (*OrderNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{5}
}
func (m *OrderNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNotify.Unmarshal(m, b)
}
func (m *OrderNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNotify.Marshal(b, m, deterministic)
}
func (dst *OrderNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNotify.Merge(dst, src)
}
func (m *OrderNotify) XXX_Size() int {
	return xxx_messageInfo_OrderNotify.Size(m)
}
func (m *OrderNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNotify.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNotify proto.InternalMessageInfo

func (m *OrderNotify) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *OrderNotify) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func (m *OrderNotify) GetChannel() SDKChannel {
	if m != nil {
		return m.Channel
	}
	return SDKChannel_GM
}

func (m *OrderNotify) GetMoney() int32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *OrderNotify) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *OrderNotify) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_Notify
}

// 封停结构
type Black struct {
	Gid                  int32    `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	Lsid                 int32    `protobuf:"varint,2,opt,name=lsid" json:"lsid,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=startTime" json:"startTime,omitempty"`
	Duration             int32    `protobuf:"varint,4,opt,name=duration" json:"duration,omitempty"`
	Info                 string   `protobuf:"bytes,5,opt,name=info" json:"info,omitempty"`
	Reason               string   `protobuf:"bytes,6,opt,name=reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Black) Reset()         { *m = Black{} }
func (m *Black) String() string { return proto.CompactTextString(m) }
func (*Black) ProtoMessage()    {}
func (*Black) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_41367a7fa062874c, []int{6}
}
func (m *Black) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Black.Unmarshal(m, b)
}
func (m *Black) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Black.Marshal(b, m, deterministic)
}
func (dst *Black) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Black.Merge(dst, src)
}
func (m *Black) XXX_Size() int {
	return xxx_messageInfo_Black.Size(m)
}
func (m *Black) XXX_DiscardUnknown() {
	xxx_messageInfo_Black.DiscardUnknown(m)
}

var xxx_messageInfo_Black proto.InternalMessageInfo

func (m *Black) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *Black) GetLsid() int32 {
	if m != nil {
		return m.Lsid
	}
	return 0
}

func (m *Black) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Black) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Black) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Black) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "UserInfo")
	proto.RegisterType((*UserRoleInfo)(nil), "UserRoleInfo")
	proto.RegisterType((*LogicServerInfo)(nil), "LogicServerInfo")
	proto.RegisterType((*SuperAccount)(nil), "SuperAccount")
	proto.RegisterType((*Line)(nil), "Line")
	proto.RegisterType((*OrderNotify)(nil), "OrderNotify")
	proto.RegisterType((*Black)(nil), "Black")
	proto.RegisterEnum("ServerStatus", ServerStatus_name, ServerStatus_value)
	proto.RegisterEnum("SDKChannel", SDKChannel_name, SDKChannel_value)
	proto.RegisterEnum("OrderStatus", OrderStatus_name, OrderStatus_value)
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_41367a7fa062874c) }

var fileDescriptor_auth_41367a7fa062874c = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x6e, 0xd3, 0x40,
	0x10, 0xed, 0xda, 0xb1, 0xe3, 0x4c, 0x43, 0xb1, 0x56, 0x15, 0xb2, 0x10, 0x42, 0x95, 0xa1, 0x52,
	0xd4, 0x8f, 0x7c, 0x14, 0x71, 0x80, 0xd2, 0xd2, 0x52, 0xd1, 0x16, 0xc9, 0x81, 0x03, 0x18, 0x7b,
	0xdb, 0xac, 0x70, 0x76, 0x23, 0xef, 0x1a, 0x92, 0x2b, 0xf0, 0xcb, 0x25, 0x38, 0x1c, 0x87, 0x40,
	0x33, 0xbb, 0x4e, 0x82, 0xd4, 0x0f, 0xfe, 0xde, 0x9b, 0xf5, 0xbc, 0x99, 0xbc, 0x37, 0x0a, 0x40,
	0xd9, 0xd9, 0xf9, 0x74, 0xd9, 0x6a, 0xab, 0xf3, 0x4b, 0x48, 0xbe, 0x18, 0xd1, 0x5e, 0xab, 0x7b,
	0xcd, 0x39, 0x0c, 0x54, 0xb9, 0x10, 0x19, 0x3b, 0x62, 0x93, 0x51, 0x41, 0x98, 0x3f, 0x83, 0x78,
	0x59, 0x1a, 0xf3, 0xa3, 0xce, 0x02, 0xaa, 0x7a, 0xc6, 0x0f, 0x20, 0x90, 0x75, 0x16, 0x52, 0x2d,
	0x90, 0x75, 0xfe, 0x93, 0xc1, 0x18, 0x85, 0x0a, 0xdd, 0x08, 0x12, 0x73, 0x1f, 0xa0, 0x54, 0x84,
	0x1f, 0x6c, 0xc4, 0x83, 0x7f, 0xc5, 0x8d, 0x68, 0xbf, 0x8b, 0x96, 0x84, 0xa2, 0xc2, 0x33, 0x7e,
	0x08, 0x51, 0x65, 0xe5, 0x42, 0x64, 0x83, 0x23, 0x36, 0x09, 0x0b, 0x47, 0x78, 0x0a, 0x61, 0x27,
	0xeb, 0x2c, 0x22, 0x01, 0x84, 0xd8, 0x6f, 0xd7, 0x4b, 0x71, 0x5d, 0x67, 0xb1, 0xeb, 0x77, 0x2c,
	0xff, 0xc3, 0xe0, 0xe9, 0x8d, 0x7e, 0x90, 0xd5, 0x8c, 0xf4, 0xfe, 0x7b, 0x9f, 0x43, 0x88, 0x8c,
	0x2d, 0x5b, 0x4b, 0xeb, 0x84, 0x85, 0x23, 0x38, 0x57, 0xa8, 0xda, 0xef, 0x82, 0x10, 0x7b, 0x69,
	0xbd, 0x88, 0x4a, 0x84, 0xf9, 0x2b, 0xea, 0xb5, 0x82, 0x56, 0x39, 0x38, 0x7d, 0x32, 0x75, 0xb3,
	0x67, 0xb6, 0xb4, 0x9d, 0x29, 0xdc, 0x1b, 0x7f, 0x09, 0x50, 0xad, 0xab, 0x46, 0xcc, 0x68, 0xca,
	0x90, 0xda, 0x77, 0x2a, 0xfc, 0x39, 0x24, 0xc4, 0xde, 0xab, 0x3a, 0x4b, 0xe8, 0x75, 0xc3, 0x79,
	0x06, 0x43, 0x69, 0xce, 0x91, 0x65, 0xa3, 0x23, 0x36, 0x49, 0x8a, 0x9e, 0xe6, 0xa7, 0x30, 0x9e,
	0x75, 0x4b, 0xd1, 0x9e, 0x55, 0x95, 0xee, 0x94, 0xed, 0x8d, 0x62, 0x5b, 0xa3, 0x52, 0x08, 0x1f,
	0xa4, 0x8b, 0x30, 0x2a, 0x10, 0xe6, 0x77, 0x30, 0xb8, 0x91, 0x4a, 0xe0, 0x4f, 0x69, 0xcc, 0xc6,
	0x18, 0xc2, 0x68, 0x95, 0x56, 0xf4, 0x71, 0x52, 0x04, 0x5a, 0x61, 0xf7, 0x42, 0x2a, 0x9f, 0x11,
	0x42, 0xaa, 0x94, 0x2b, 0xb2, 0x04, 0x2b, 0xe5, 0x2a, 0xff, 0xcd, 0x60, 0xff, 0x53, 0x5b, 0x8b,
	0xf6, 0x4e, 0x5b, 0x79, 0xbf, 0xee, 0x27, 0xb2, 0xcd, 0x44, 0xac, 0x68, 0xd9, 0x9f, 0x11, 0x42,
	0x7e, 0x0c, 0xc3, 0x6a, 0x5e, 0x2a, 0x25, 0x1a, 0xd2, 0x3e, 0x38, 0xdd, 0x9f, 0xce, 0x2e, 0x3e,
	0x9e, 0xbb, 0x52, 0xd1, 0xbf, 0x61, 0x2a, 0x0b, 0xad, 0xc4, 0xda, 0x8f, 0x73, 0xe4, 0xd1, 0x0c,
	0x5e, 0x43, 0x6c, 0xc8, 0x6f, 0x1f, 0xc2, 0x78, 0x4a, 0x2b, 0xf9, 0x0c, 0xfc, 0x5b, 0xfe, 0x8b,
	0x41, 0xf4, 0xae, 0x29, 0xab, 0x6f, 0x8f, 0x2c, 0xd9, 0xdb, 0x11, 0xec, 0xd8, 0xf1, 0x02, 0x46,
	0x74, 0x08, 0x9f, 0x71, 0x9c, 0xbb, 0x8c, 0x6d, 0x01, 0x23, 0xab, 0xbb, 0xb6, 0xb4, 0x52, 0x2b,
	0xbf, 0xe0, 0x86, 0xa3, 0x9a, 0x54, 0xf7, 0xda, 0x9f, 0x2c, 0x61, 0xbc, 0xd9, 0x56, 0x94, 0x46,
	0x2b, 0xda, 0x71, 0x54, 0x78, 0x76, 0xf2, 0x16, 0xc6, 0xbb, 0x17, 0xc3, 0x13, 0x18, 0x5c, 0x69,
	0x5d, 0xa7, 0x7b, 0x7c, 0x08, 0xe1, 0x07, 0x6d, 0x53, 0x86, 0xe0, 0x52, 0xae, 0xd2, 0x80, 0x8f,
	0x20, 0x3a, 0x6f, 0xb4, 0x11, 0x69, 0x78, 0x32, 0x01, 0xd8, 0x7a, 0xc6, 0x63, 0x08, 0xae, 0x6e,
	0xd3, 0x3d, 0x6c, 0xbe, 0x95, 0x4a, 0xa4, 0x0c, 0xd1, 0x85, 0xb8, 0x3b, 0x4b, 0x83, 0x93, 0x63,
	0x1f, 0x90, 0xd7, 0x07, 0x88, 0x5d, 0x54, 0xe9, 0x1e, 0xe2, 0x4b, 0xa9, 0xa4, 0x99, 0xa7, 0xec,
	0x6b, 0x4c, 0xff, 0x0b, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xca, 0xe9, 0x91, 0x25,
	0x04, 0x00, 0x00,
}
