// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clientmessages.proto

package dota

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

type EBaseClientMessages int32

const (
	EBaseClientMessages_CM_CustomGameEvent                  EBaseClientMessages = 280
	EBaseClientMessages_CM_CustomGameEventBounce            EBaseClientMessages = 281
	EBaseClientMessages_CM_ClientUIEvent                    EBaseClientMessages = 282
	EBaseClientMessages_CM_DevPaletteVisibilityChanged      EBaseClientMessages = 283
	EBaseClientMessages_CM_WorldUIControllerHasPanelChanged EBaseClientMessages = 284
	EBaseClientMessages_CM_RotateAnchor                     EBaseClientMessages = 285
	EBaseClientMessages_CM_MAX_BASE                         EBaseClientMessages = 300
)

var EBaseClientMessages_name = map[int32]string{
	280: "CM_CustomGameEvent",
	281: "CM_CustomGameEventBounce",
	282: "CM_ClientUIEvent",
	283: "CM_DevPaletteVisibilityChanged",
	284: "CM_WorldUIControllerHasPanelChanged",
	285: "CM_RotateAnchor",
	300: "CM_MAX_BASE",
}

var EBaseClientMessages_value = map[string]int32{
	"CM_CustomGameEvent":                  280,
	"CM_CustomGameEventBounce":            281,
	"CM_ClientUIEvent":                    282,
	"CM_DevPaletteVisibilityChanged":      283,
	"CM_WorldUIControllerHasPanelChanged": 284,
	"CM_RotateAnchor":                     285,
	"CM_MAX_BASE":                         300,
}

func (x EBaseClientMessages) Enum() *EBaseClientMessages {
	p := new(EBaseClientMessages)
	*p = x
	return p
}

func (x EBaseClientMessages) String() string {
	return proto.EnumName(EBaseClientMessages_name, int32(x))
}

func (x *EBaseClientMessages) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EBaseClientMessages_value, data, "EBaseClientMessages")
	if err != nil {
		return err
	}
	*x = EBaseClientMessages(value)
	return nil
}

func (EBaseClientMessages) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2ecb9f24e0641136, []int{0}
}

type EClientUIEvent int32

const (
	EClientUIEvent_EClientUIEvent_Invalid        EClientUIEvent = 0
	EClientUIEvent_EClientUIEvent_DialogFinished EClientUIEvent = 1
	EClientUIEvent_EClientUIEvent_FireOutput     EClientUIEvent = 2
)

var EClientUIEvent_name = map[int32]string{
	0: "EClientUIEvent_Invalid",
	1: "EClientUIEvent_DialogFinished",
	2: "EClientUIEvent_FireOutput",
}

var EClientUIEvent_value = map[string]int32{
	"EClientUIEvent_Invalid":        0,
	"EClientUIEvent_DialogFinished": 1,
	"EClientUIEvent_FireOutput":     2,
}

func (x EClientUIEvent) Enum() *EClientUIEvent {
	p := new(EClientUIEvent)
	*p = x
	return p
}

func (x EClientUIEvent) String() string {
	return proto.EnumName(EClientUIEvent_name, int32(x))
}

func (x *EClientUIEvent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EClientUIEvent_value, data, "EClientUIEvent")
	if err != nil {
		return err
	}
	*x = EClientUIEvent(value)
	return nil
}

func (EClientUIEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2ecb9f24e0641136, []int{1}
}

type CClientMsg_CustomGameEvent struct {
	EventName            *string  `protobuf:"bytes,1,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_CustomGameEvent) Reset()         { *m = CClientMsg_CustomGameEvent{} }
func (m *CClientMsg_CustomGameEvent) String() string { return proto.CompactTextString(m) }
func (*CClientMsg_CustomGameEvent) ProtoMessage()    {}
func (*CClientMsg_CustomGameEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ecb9f24e0641136, []int{0}
}

func (m *CClientMsg_CustomGameEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_CustomGameEvent.Unmarshal(m, b)
}
func (m *CClientMsg_CustomGameEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_CustomGameEvent.Marshal(b, m, deterministic)
}
func (m *CClientMsg_CustomGameEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_CustomGameEvent.Merge(m, src)
}
func (m *CClientMsg_CustomGameEvent) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_CustomGameEvent.Size(m)
}
func (m *CClientMsg_CustomGameEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_CustomGameEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_CustomGameEvent proto.InternalMessageInfo

func (m *CClientMsg_CustomGameEvent) GetEventName() string {
	if m != nil && m.EventName != nil {
		return *m.EventName
	}
	return ""
}

func (m *CClientMsg_CustomGameEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CClientMsg_CustomGameEventBounce struct {
	EventName            *string  `protobuf:"bytes,1,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	PlayerIndex          *int32   `protobuf:"varint,3,opt,name=player_index,json=playerIndex" json:"player_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_CustomGameEventBounce) Reset()         { *m = CClientMsg_CustomGameEventBounce{} }
func (m *CClientMsg_CustomGameEventBounce) String() string { return proto.CompactTextString(m) }
func (*CClientMsg_CustomGameEventBounce) ProtoMessage()    {}
func (*CClientMsg_CustomGameEventBounce) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ecb9f24e0641136, []int{1}
}

func (m *CClientMsg_CustomGameEventBounce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_CustomGameEventBounce.Unmarshal(m, b)
}
func (m *CClientMsg_CustomGameEventBounce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_CustomGameEventBounce.Marshal(b, m, deterministic)
}
func (m *CClientMsg_CustomGameEventBounce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_CustomGameEventBounce.Merge(m, src)
}
func (m *CClientMsg_CustomGameEventBounce) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_CustomGameEventBounce.Size(m)
}
func (m *CClientMsg_CustomGameEventBounce) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_CustomGameEventBounce.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_CustomGameEventBounce proto.InternalMessageInfo

func (m *CClientMsg_CustomGameEventBounce) GetEventName() string {
	if m != nil && m.EventName != nil {
		return *m.EventName
	}
	return ""
}

func (m *CClientMsg_CustomGameEventBounce) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CClientMsg_CustomGameEventBounce) GetPlayerIndex() int32 {
	if m != nil && m.PlayerIndex != nil {
		return *m.PlayerIndex
	}
	return 0
}

type CClientMsg_ClientUIEvent struct {
	Event                *EClientUIEvent `protobuf:"varint,1,opt,name=event,enum=dota.EClientUIEvent,def=0" json:"event,omitempty"`
	EntEhandle           *uint32         `protobuf:"varint,2,opt,name=ent_ehandle,json=entEhandle" json:"ent_ehandle,omitempty"`
	ClientEhandle        *uint32         `protobuf:"varint,3,opt,name=client_ehandle,json=clientEhandle" json:"client_ehandle,omitempty"`
	Data1                *string         `protobuf:"bytes,4,opt,name=data1" json:"data1,omitempty"`
	Data2                *string         `protobuf:"bytes,5,opt,name=data2" json:"data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CClientMsg_ClientUIEvent) Reset()         { *m = CClientMsg_ClientUIEvent{} }
func (m *CClientMsg_ClientUIEvent) String() string { return proto.CompactTextString(m) }
func (*CClientMsg_ClientUIEvent) ProtoMessage()    {}
func (*CClientMsg_ClientUIEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ecb9f24e0641136, []int{2}
}

func (m *CClientMsg_ClientUIEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_ClientUIEvent.Unmarshal(m, b)
}
func (m *CClientMsg_ClientUIEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_ClientUIEvent.Marshal(b, m, deterministic)
}
func (m *CClientMsg_ClientUIEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_ClientUIEvent.Merge(m, src)
}
func (m *CClientMsg_ClientUIEvent) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_ClientUIEvent.Size(m)
}
func (m *CClientMsg_ClientUIEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_ClientUIEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_ClientUIEvent proto.InternalMessageInfo

const Default_CClientMsg_ClientUIEvent_Event EClientUIEvent = EClientUIEvent_EClientUIEvent_Invalid

func (m *CClientMsg_ClientUIEvent) GetEvent() EClientUIEvent {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return Default_CClientMsg_ClientUIEvent_Event
}

func (m *CClientMsg_ClientUIEvent) GetEntEhandle() uint32 {
	if m != nil && m.EntEhandle != nil {
		return *m.EntEhandle
	}
	return 0
}

func (m *CClientMsg_ClientUIEvent) GetClientEhandle() uint32 {
	if m != nil && m.ClientEhandle != nil {
		return *m.ClientEhandle
	}
	return 0
}

func (m *CClientMsg_ClientUIEvent) GetData1() string {
	if m != nil && m.Data1 != nil {
		return *m.Data1
	}
	return ""
}

func (m *CClientMsg_ClientUIEvent) GetData2() string {
	if m != nil && m.Data2 != nil {
		return *m.Data2
	}
	return ""
}

type CClientMsg_DevPaletteVisibilityChangedEvent struct {
	Visible              *bool    `protobuf:"varint,1,opt,name=visible" json:"visible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_DevPaletteVisibilityChangedEvent) Reset() {
	*m = CClientMsg_DevPaletteVisibilityChangedEvent{}
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) String() string {
	return proto.CompactTextString(m)
}
func (*CClientMsg_DevPaletteVisibilityChangedEvent) ProtoMessage() {}
func (*CClientMsg_DevPaletteVisibilityChangedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ecb9f24e0641136, []int{3}
}

func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.Unmarshal(m, b)
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.Marshal(b, m, deterministic)
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.Merge(m, src)
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.Size(m)
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent proto.InternalMessageInfo

func (m *CClientMsg_DevPaletteVisibilityChangedEvent) GetVisible() bool {
	if m != nil && m.Visible != nil {
		return *m.Visible
	}
	return false
}

type CClientMsg_WorldUIControllerHasPanelChangedEvent struct {
	HasPanel             *bool    `protobuf:"varint,1,opt,name=has_panel,json=hasPanel" json:"has_panel,omitempty"`
	ClientEhandle        *uint32  `protobuf:"varint,2,opt,name=client_ehandle,json=clientEhandle" json:"client_ehandle,omitempty"`
	LiteralHandType      *uint32  `protobuf:"varint,3,opt,name=literal_hand_type,json=literalHandType" json:"literal_hand_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) Reset() {
	*m = CClientMsg_WorldUIControllerHasPanelChangedEvent{}
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) String() string {
	return proto.CompactTextString(m)
}
func (*CClientMsg_WorldUIControllerHasPanelChangedEvent) ProtoMessage() {}
func (*CClientMsg_WorldUIControllerHasPanelChangedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ecb9f24e0641136, []int{4}
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.Unmarshal(m, b)
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.Marshal(b, m, deterministic)
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.Merge(m, src)
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.Size(m)
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent proto.InternalMessageInfo

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetHasPanel() bool {
	if m != nil && m.HasPanel != nil {
		return *m.HasPanel
	}
	return false
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetClientEhandle() uint32 {
	if m != nil && m.ClientEhandle != nil {
		return *m.ClientEhandle
	}
	return 0
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetLiteralHandType() uint32 {
	if m != nil && m.LiteralHandType != nil {
		return *m.LiteralHandType
	}
	return 0
}

type CClientMsg_RotateAnchor struct {
	Angle                *float32 `protobuf:"fixed32,1,opt,name=angle" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_RotateAnchor) Reset()         { *m = CClientMsg_RotateAnchor{} }
func (m *CClientMsg_RotateAnchor) String() string { return proto.CompactTextString(m) }
func (*CClientMsg_RotateAnchor) ProtoMessage()    {}
func (*CClientMsg_RotateAnchor) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ecb9f24e0641136, []int{5}
}

func (m *CClientMsg_RotateAnchor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_RotateAnchor.Unmarshal(m, b)
}
func (m *CClientMsg_RotateAnchor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_RotateAnchor.Marshal(b, m, deterministic)
}
func (m *CClientMsg_RotateAnchor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_RotateAnchor.Merge(m, src)
}
func (m *CClientMsg_RotateAnchor) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_RotateAnchor.Size(m)
}
func (m *CClientMsg_RotateAnchor) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_RotateAnchor.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_RotateAnchor proto.InternalMessageInfo

func (m *CClientMsg_RotateAnchor) GetAngle() float32 {
	if m != nil && m.Angle != nil {
		return *m.Angle
	}
	return 0
}

func init() {
	proto.RegisterEnum("dota.EBaseClientMessages", EBaseClientMessages_name, EBaseClientMessages_value)
	proto.RegisterEnum("dota.EClientUIEvent", EClientUIEvent_name, EClientUIEvent_value)
	proto.RegisterType((*CClientMsg_CustomGameEvent)(nil), "dota.CClientMsg_CustomGameEvent")
	proto.RegisterType((*CClientMsg_CustomGameEventBounce)(nil), "dota.CClientMsg_CustomGameEventBounce")
	proto.RegisterType((*CClientMsg_ClientUIEvent)(nil), "dota.CClientMsg_ClientUIEvent")
	proto.RegisterType((*CClientMsg_DevPaletteVisibilityChangedEvent)(nil), "dota.CClientMsg_DevPaletteVisibilityChangedEvent")
	proto.RegisterType((*CClientMsg_WorldUIControllerHasPanelChangedEvent)(nil), "dota.CClientMsg_WorldUIControllerHasPanelChangedEvent")
	proto.RegisterType((*CClientMsg_RotateAnchor)(nil), "dota.CClientMsg_RotateAnchor")
}

func init() { proto.RegisterFile("clientmessages.proto", fileDescriptor_2ecb9f24e0641136) }

var fileDescriptor_2ecb9f24e0641136 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0xac, 0xdd, 0x46, 0x34, 0x2f, 0x6d, 0xba, 0x2c, 0x81, 0x9a, 0xa2, 0x40, 0xea, 0x0a, 0x29,
	0x0a, 0x52, 0x80, 0x1c, 0xb9, 0x25, 0x6e, 0xda, 0xe4, 0x60, 0x5a, 0x19, 0x0a, 0xdc, 0xac, 0x25,
	0x7e, 0x8a, 0x57, 0xda, 0xac, 0x2d, 0x7b, 0x13, 0x35, 0x37, 0x7e, 0x06, 0x9f, 0x17, 0xce, 0xfc,
	0x1a, 0xae, 0xfc, 0x19, 0x64, 0xaf, 0x03, 0x49, 0x68, 0x41, 0xe2, 0xb6, 0x6f, 0x66, 0xf6, 0xed,
	0xce, 0x68, 0xa0, 0x36, 0x12, 0x1c, 0xa5, 0x9a, 0x60, 0x9a, 0xb2, 0x31, 0xa6, 0xed, 0x38, 0x89,
	0x54, 0x44, 0xb7, 0x82, 0x48, 0x31, 0xfb, 0x0c, 0x0e, 0x1c, 0x27, 0xa7, 0xdd, 0x74, 0xec, 0x3b,
	0xd3, 0x54, 0x45, 0x93, 0x53, 0x36, 0xc1, 0xfe, 0x0c, 0xa5, 0xa2, 0x75, 0x00, 0xcc, 0x0e, 0xbe,
	0x64, 0x13, 0xb4, 0x8c, 0x86, 0xd1, 0x2c, 0x7b, 0xe5, 0x1c, 0x79, 0xce, 0x26, 0x48, 0x29, 0x6c,
	0x05, 0x4c, 0x31, 0xcb, 0x6c, 0x18, 0xcd, 0x1d, 0x2f, 0x3f, 0xdb, 0x97, 0xd0, 0xb8, 0x7e, 0x61,
	0x2f, 0x9a, 0xca, 0x11, 0xfe, 0xc7, 0x5a, 0x7a, 0x08, 0x3b, 0xb1, 0x60, 0x73, 0x4c, 0x7c, 0x2e,
	0x03, 0xbc, 0xb4, 0x36, 0x1b, 0x46, 0xb3, 0xe4, 0x55, 0x34, 0x36, 0xcc, 0x20, 0xfb, 0xbb, 0x01,
	0xd6, 0xf2, 0xd3, 0xf9, 0xe9, 0x62, 0xa8, 0x9d, 0xf4, 0xa0, 0x94, 0x3f, 0x90, 0xbf, 0x56, 0xed,
	0xd4, 0xda, 0x99, 0xfb, 0x76, 0x7f, 0x45, 0xf4, 0xec, 0xce, 0xea, 0xec, 0x0f, 0xe5, 0x8c, 0x09,
	0x1e, 0x78, 0xfa, 0x2a, 0x7d, 0x00, 0x95, 0x0c, 0xc5, 0x90, 0xc9, 0x40, 0x60, 0xfe, 0xbd, 0x5d,
	0x0f, 0x50, 0xaa, 0xbe, 0x46, 0xe8, 0x43, 0xa8, 0xea, 0xa8, 0x7f, 0x69, 0x36, 0x73, 0xcd, 0xae,
	0x46, 0x17, 0xb2, 0x1a, 0x94, 0x32, 0x4f, 0x4f, 0xad, 0xad, 0xdc, 0xb9, 0x1e, 0x16, 0x68, 0xc7,
	0x2a, 0xfd, 0x46, 0x3b, 0xf6, 0x29, 0x3c, 0x5a, 0xf2, 0x74, 0x8c, 0xb3, 0x73, 0x26, 0x50, 0x29,
	0x7c, 0xc5, 0x53, 0xfe, 0x96, 0x0b, 0xae, 0xe6, 0x4e, 0xc8, 0xe4, 0x18, 0x03, 0x6d, 0xd3, 0x82,
	0x1b, 0xb3, 0x8c, 0x11, 0x3a, 0xd6, 0x6d, 0x6f, 0x31, 0xda, 0x5f, 0x0d, 0x78, 0xb2, 0xb4, 0xe9,
	0x75, 0x94, 0x88, 0xe0, 0x62, 0xe8, 0x44, 0x52, 0x25, 0x91, 0x10, 0x98, 0x0c, 0x58, 0x7a, 0xce,
	0x24, 0x8a, 0x95, 0x75, 0xf7, 0xa0, 0x1c, 0xb2, 0xd4, 0x8f, 0x33, 0xa2, 0x58, 0xb8, 0x1d, 0x16,
	0xc2, 0x2b, 0xdc, 0x9a, 0x57, 0xb9, 0x6d, 0xc1, 0x4d, 0xc1, 0x15, 0x26, 0x4c, 0xf8, 0x19, 0xe2,
	0xab, 0x79, 0xbc, 0xc8, 0x65, 0xaf, 0x20, 0x06, 0x4c, 0x06, 0x2f, 0xe7, 0x31, 0xda, 0x8f, 0x61,
	0x7f, 0xe9, 0x8f, 0x5e, 0xa4, 0x98, 0xc2, 0xae, 0x1c, 0x85, 0x51, 0x92, 0xc5, 0xc3, 0xe4, 0xb8,
	0xf0, 0x65, 0x7a, 0x7a, 0x68, 0xfd, 0x30, 0xe0, 0x56, 0xbf, 0xc7, 0x52, 0x2c, 0x6e, 0x15, 0x15,
	0xa7, 0xfb, 0x40, 0x1d, 0x77, 0xbd, 0x7d, 0xe4, 0xbd, 0x49, 0xeb, 0x60, 0xfd, 0x49, 0xe8, 0x5a,
	0x92, 0x0f, 0x26, 0xbd, 0x0d, 0x24, 0xa3, 0x97, 0x5b, 0x40, 0x3e, 0x9a, 0xf4, 0x08, 0xee, 0x3b,
	0xee, 0xdf, 0xd2, 0x27, 0x9f, 0x4c, 0xda, 0x84, 0x23, 0xc7, 0xfd, 0x67, 0xb0, 0xe4, 0xb3, 0x49,
	0x6b, 0xb0, 0xe7, 0xb8, 0x2b, 0xf6, 0xc8, 0x17, 0x93, 0x12, 0xa8, 0x38, 0xae, 0xef, 0x76, 0xdf,
	0xf8, 0xbd, 0xee, 0x8b, 0x3e, 0xf9, 0x66, 0xb6, 0x24, 0x54, 0x57, 0x1b, 0x49, 0x0f, 0xe0, 0x9a,
	0x8e, 0x92, 0x0d, 0x7a, 0x08, 0xf5, 0x35, 0xee, 0x98, 0x33, 0x11, 0x8d, 0x4f, 0xb8, 0xe4, 0x69,
	0x88, 0x01, 0x31, 0x68, 0x1d, 0xee, 0xae, 0x49, 0x4e, 0x78, 0x82, 0x67, 0x53, 0x15, 0x4f, 0x15,
	0x31, 0x7b, 0xa5, 0x81, 0xf1, 0xce, 0xd8, 0xf8, 0x19, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x8b, 0x63,
	0x64, 0x30, 0x04, 0x00, 0x00,
}
