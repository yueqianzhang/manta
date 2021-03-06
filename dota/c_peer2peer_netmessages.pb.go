// Code generated by protoc-gen-go. DO NOT EDIT.
// source: c_peer2peer_netmessages.proto

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

type P2P_Messages int32

const (
	P2P_Messages_p2p_TextMessage          P2P_Messages = 256
	P2P_Messages_p2p_Voice                P2P_Messages = 257
	P2P_Messages_p2p_Ping                 P2P_Messages = 258
	P2P_Messages_p2p_VRAvatarPosition     P2P_Messages = 259
	P2P_Messages_p2p_WatchSynchronization P2P_Messages = 260
)

var P2P_Messages_name = map[int32]string{
	256: "p2p_TextMessage",
	257: "p2p_Voice",
	258: "p2p_Ping",
	259: "p2p_VRAvatarPosition",
	260: "p2p_WatchSynchronization",
}

var P2P_Messages_value = map[string]int32{
	"p2p_TextMessage":          256,
	"p2p_Voice":                257,
	"p2p_Ping":                 258,
	"p2p_VRAvatarPosition":     259,
	"p2p_WatchSynchronization": 260,
}

func (x P2P_Messages) Enum() *P2P_Messages {
	p := new(P2P_Messages)
	*p = x
	return p
}

func (x P2P_Messages) String() string {
	return proto.EnumName(P2P_Messages_name, int32(x))
}

func (x *P2P_Messages) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(P2P_Messages_value, data, "P2P_Messages")
	if err != nil {
		return err
	}
	*x = P2P_Messages(value)
	return nil
}

func (P2P_Messages) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{0}
}

type CP2P_Voice_Handler_Flags int32

const (
	CP2P_Voice_Played_Audio CP2P_Voice_Handler_Flags = 1
)

var CP2P_Voice_Handler_Flags_name = map[int32]string{
	1: "Played_Audio",
}

var CP2P_Voice_Handler_Flags_value = map[string]int32{
	"Played_Audio": 1,
}

func (x CP2P_Voice_Handler_Flags) Enum() *CP2P_Voice_Handler_Flags {
	p := new(CP2P_Voice_Handler_Flags)
	*p = x
	return p
}

func (x CP2P_Voice_Handler_Flags) String() string {
	return proto.EnumName(CP2P_Voice_Handler_Flags_name, int32(x))
}

func (x *CP2P_Voice_Handler_Flags) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CP2P_Voice_Handler_Flags_value, data, "CP2P_Voice_Handler_Flags")
	if err != nil {
		return err
	}
	*x = CP2P_Voice_Handler_Flags(value)
	return nil
}

func (CP2P_Voice_Handler_Flags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{2, 0}
}

type CP2P_TextMessage struct {
	Text                 []byte   `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CP2P_TextMessage) Reset()         { *m = CP2P_TextMessage{} }
func (m *CP2P_TextMessage) String() string { return proto.CompactTextString(m) }
func (*CP2P_TextMessage) ProtoMessage()    {}
func (*CP2P_TextMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{0}
}

func (m *CP2P_TextMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CP2P_TextMessage.Unmarshal(m, b)
}
func (m *CP2P_TextMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CP2P_TextMessage.Marshal(b, m, deterministic)
}
func (m *CP2P_TextMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CP2P_TextMessage.Merge(m, src)
}
func (m *CP2P_TextMessage) XXX_Size() int {
	return xxx_messageInfo_CP2P_TextMessage.Size(m)
}
func (m *CP2P_TextMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CP2P_TextMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CP2P_TextMessage proto.InternalMessageInfo

func (m *CP2P_TextMessage) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

type CSteam_Voice_Encoding struct {
	VoiceData            []byte   `protobuf:"bytes,1,opt,name=voice_data,json=voiceData" json:"voice_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CSteam_Voice_Encoding) Reset()         { *m = CSteam_Voice_Encoding{} }
func (m *CSteam_Voice_Encoding) String() string { return proto.CompactTextString(m) }
func (*CSteam_Voice_Encoding) ProtoMessage()    {}
func (*CSteam_Voice_Encoding) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{1}
}

func (m *CSteam_Voice_Encoding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSteam_Voice_Encoding.Unmarshal(m, b)
}
func (m *CSteam_Voice_Encoding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSteam_Voice_Encoding.Marshal(b, m, deterministic)
}
func (m *CSteam_Voice_Encoding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSteam_Voice_Encoding.Merge(m, src)
}
func (m *CSteam_Voice_Encoding) XXX_Size() int {
	return xxx_messageInfo_CSteam_Voice_Encoding.Size(m)
}
func (m *CSteam_Voice_Encoding) XXX_DiscardUnknown() {
	xxx_messageInfo_CSteam_Voice_Encoding.DiscardUnknown(m)
}

var xxx_messageInfo_CSteam_Voice_Encoding proto.InternalMessageInfo

func (m *CSteam_Voice_Encoding) GetVoiceData() []byte {
	if m != nil {
		return m.VoiceData
	}
	return nil
}

type CP2P_Voice struct {
	Audio                *CMsgVoiceAudio `protobuf:"bytes,1,opt,name=audio" json:"audio,omitempty"`
	BroadcastGroup       *uint32         `protobuf:"varint,2,opt,name=broadcast_group,json=broadcastGroup" json:"broadcast_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CP2P_Voice) Reset()         { *m = CP2P_Voice{} }
func (m *CP2P_Voice) String() string { return proto.CompactTextString(m) }
func (*CP2P_Voice) ProtoMessage()    {}
func (*CP2P_Voice) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{2}
}

func (m *CP2P_Voice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CP2P_Voice.Unmarshal(m, b)
}
func (m *CP2P_Voice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CP2P_Voice.Marshal(b, m, deterministic)
}
func (m *CP2P_Voice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CP2P_Voice.Merge(m, src)
}
func (m *CP2P_Voice) XXX_Size() int {
	return xxx_messageInfo_CP2P_Voice.Size(m)
}
func (m *CP2P_Voice) XXX_DiscardUnknown() {
	xxx_messageInfo_CP2P_Voice.DiscardUnknown(m)
}

var xxx_messageInfo_CP2P_Voice proto.InternalMessageInfo

func (m *CP2P_Voice) GetAudio() *CMsgVoiceAudio {
	if m != nil {
		return m.Audio
	}
	return nil
}

func (m *CP2P_Voice) GetBroadcastGroup() uint32 {
	if m != nil && m.BroadcastGroup != nil {
		return *m.BroadcastGroup
	}
	return 0
}

type CP2P_Ping struct {
	SendTime             *uint64  `protobuf:"varint,1,req,name=send_time,json=sendTime" json:"send_time,omitempty"`
	IsReply              *bool    `protobuf:"varint,2,req,name=is_reply,json=isReply" json:"is_reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CP2P_Ping) Reset()         { *m = CP2P_Ping{} }
func (m *CP2P_Ping) String() string { return proto.CompactTextString(m) }
func (*CP2P_Ping) ProtoMessage()    {}
func (*CP2P_Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{3}
}

func (m *CP2P_Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CP2P_Ping.Unmarshal(m, b)
}
func (m *CP2P_Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CP2P_Ping.Marshal(b, m, deterministic)
}
func (m *CP2P_Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CP2P_Ping.Merge(m, src)
}
func (m *CP2P_Ping) XXX_Size() int {
	return xxx_messageInfo_CP2P_Ping.Size(m)
}
func (m *CP2P_Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_CP2P_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_CP2P_Ping proto.InternalMessageInfo

func (m *CP2P_Ping) GetSendTime() uint64 {
	if m != nil && m.SendTime != nil {
		return *m.SendTime
	}
	return 0
}

func (m *CP2P_Ping) GetIsReply() bool {
	if m != nil && m.IsReply != nil {
		return *m.IsReply
	}
	return false
}

type CP2P_VRAvatarPosition struct {
	BodyParts            []*CP2P_VRAvatarPosition_COrientation `protobuf:"bytes,1,rep,name=body_parts,json=bodyParts" json:"body_parts,omitempty"`
	HatId                *int32                                `protobuf:"varint,2,opt,name=hat_id,json=hatId" json:"hat_id,omitempty"`
	SceneId              *int32                                `protobuf:"varint,3,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	WorldScale           *int32                                `protobuf:"varint,4,opt,name=world_scale,json=worldScale" json:"world_scale,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *CP2P_VRAvatarPosition) Reset()         { *m = CP2P_VRAvatarPosition{} }
func (m *CP2P_VRAvatarPosition) String() string { return proto.CompactTextString(m) }
func (*CP2P_VRAvatarPosition) ProtoMessage()    {}
func (*CP2P_VRAvatarPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{4}
}

func (m *CP2P_VRAvatarPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CP2P_VRAvatarPosition.Unmarshal(m, b)
}
func (m *CP2P_VRAvatarPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CP2P_VRAvatarPosition.Marshal(b, m, deterministic)
}
func (m *CP2P_VRAvatarPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CP2P_VRAvatarPosition.Merge(m, src)
}
func (m *CP2P_VRAvatarPosition) XXX_Size() int {
	return xxx_messageInfo_CP2P_VRAvatarPosition.Size(m)
}
func (m *CP2P_VRAvatarPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_CP2P_VRAvatarPosition.DiscardUnknown(m)
}

var xxx_messageInfo_CP2P_VRAvatarPosition proto.InternalMessageInfo

func (m *CP2P_VRAvatarPosition) GetBodyParts() []*CP2P_VRAvatarPosition_COrientation {
	if m != nil {
		return m.BodyParts
	}
	return nil
}

func (m *CP2P_VRAvatarPosition) GetHatId() int32 {
	if m != nil && m.HatId != nil {
		return *m.HatId
	}
	return 0
}

func (m *CP2P_VRAvatarPosition) GetSceneId() int32 {
	if m != nil && m.SceneId != nil {
		return *m.SceneId
	}
	return 0
}

func (m *CP2P_VRAvatarPosition) GetWorldScale() int32 {
	if m != nil && m.WorldScale != nil {
		return *m.WorldScale
	}
	return 0
}

type CP2P_VRAvatarPosition_COrientation struct {
	Pos                  *CMsgVector `protobuf:"bytes,1,opt,name=pos" json:"pos,omitempty"`
	Ang                  *CMsgQAngle `protobuf:"bytes,2,opt,name=ang" json:"ang,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CP2P_VRAvatarPosition_COrientation) Reset()         { *m = CP2P_VRAvatarPosition_COrientation{} }
func (m *CP2P_VRAvatarPosition_COrientation) String() string { return proto.CompactTextString(m) }
func (*CP2P_VRAvatarPosition_COrientation) ProtoMessage()    {}
func (*CP2P_VRAvatarPosition_COrientation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{4, 0}
}

func (m *CP2P_VRAvatarPosition_COrientation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CP2P_VRAvatarPosition_COrientation.Unmarshal(m, b)
}
func (m *CP2P_VRAvatarPosition_COrientation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CP2P_VRAvatarPosition_COrientation.Marshal(b, m, deterministic)
}
func (m *CP2P_VRAvatarPosition_COrientation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CP2P_VRAvatarPosition_COrientation.Merge(m, src)
}
func (m *CP2P_VRAvatarPosition_COrientation) XXX_Size() int {
	return xxx_messageInfo_CP2P_VRAvatarPosition_COrientation.Size(m)
}
func (m *CP2P_VRAvatarPosition_COrientation) XXX_DiscardUnknown() {
	xxx_messageInfo_CP2P_VRAvatarPosition_COrientation.DiscardUnknown(m)
}

var xxx_messageInfo_CP2P_VRAvatarPosition_COrientation proto.InternalMessageInfo

func (m *CP2P_VRAvatarPosition_COrientation) GetPos() *CMsgVector {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *CP2P_VRAvatarPosition_COrientation) GetAng() *CMsgQAngle {
	if m != nil {
		return m.Ang
	}
	return nil
}

type CP2P_WatchSynchronization struct {
	DemoTick                         *int32   `protobuf:"varint,1,opt,name=demo_tick,json=demoTick" json:"demo_tick,omitempty"`
	Paused                           *bool    `protobuf:"varint,2,opt,name=paused" json:"paused,omitempty"`
	TvListenVoiceIndices             *int32   `protobuf:"varint,3,opt,name=tv_listen_voice_indices,json=tvListenVoiceIndices" json:"tv_listen_voice_indices,omitempty"`
	DotaSpectatorMode                *int32   `protobuf:"varint,4,opt,name=dota_spectator_mode,json=dotaSpectatorMode" json:"dota_spectator_mode,omitempty"`
	DotaSpectatorWatchingBroadcaster *int32   `protobuf:"varint,5,opt,name=dota_spectator_watching_broadcaster,json=dotaSpectatorWatchingBroadcaster" json:"dota_spectator_watching_broadcaster,omitempty"`
	DotaSpectatorHeroIndex           *int32   `protobuf:"varint,6,opt,name=dota_spectator_hero_index,json=dotaSpectatorHeroIndex" json:"dota_spectator_hero_index,omitempty"`
	DotaSpectatorAutospeed           *int32   `protobuf:"varint,7,opt,name=dota_spectator_autospeed,json=dotaSpectatorAutospeed" json:"dota_spectator_autospeed,omitempty"`
	DotaReplaySpeed                  *int32   `protobuf:"varint,8,opt,name=dota_replay_speed,json=dotaReplaySpeed" json:"dota_replay_speed,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *CP2P_WatchSynchronization) Reset()         { *m = CP2P_WatchSynchronization{} }
func (m *CP2P_WatchSynchronization) String() string { return proto.CompactTextString(m) }
func (*CP2P_WatchSynchronization) ProtoMessage()    {}
func (*CP2P_WatchSynchronization) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fec9daf05519245, []int{5}
}

func (m *CP2P_WatchSynchronization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CP2P_WatchSynchronization.Unmarshal(m, b)
}
func (m *CP2P_WatchSynchronization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CP2P_WatchSynchronization.Marshal(b, m, deterministic)
}
func (m *CP2P_WatchSynchronization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CP2P_WatchSynchronization.Merge(m, src)
}
func (m *CP2P_WatchSynchronization) XXX_Size() int {
	return xxx_messageInfo_CP2P_WatchSynchronization.Size(m)
}
func (m *CP2P_WatchSynchronization) XXX_DiscardUnknown() {
	xxx_messageInfo_CP2P_WatchSynchronization.DiscardUnknown(m)
}

var xxx_messageInfo_CP2P_WatchSynchronization proto.InternalMessageInfo

func (m *CP2P_WatchSynchronization) GetDemoTick() int32 {
	if m != nil && m.DemoTick != nil {
		return *m.DemoTick
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetPaused() bool {
	if m != nil && m.Paused != nil {
		return *m.Paused
	}
	return false
}

func (m *CP2P_WatchSynchronization) GetTvListenVoiceIndices() int32 {
	if m != nil && m.TvListenVoiceIndices != nil {
		return *m.TvListenVoiceIndices
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaSpectatorMode() int32 {
	if m != nil && m.DotaSpectatorMode != nil {
		return *m.DotaSpectatorMode
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaSpectatorWatchingBroadcaster() int32 {
	if m != nil && m.DotaSpectatorWatchingBroadcaster != nil {
		return *m.DotaSpectatorWatchingBroadcaster
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaSpectatorHeroIndex() int32 {
	if m != nil && m.DotaSpectatorHeroIndex != nil {
		return *m.DotaSpectatorHeroIndex
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaSpectatorAutospeed() int32 {
	if m != nil && m.DotaSpectatorAutospeed != nil {
		return *m.DotaSpectatorAutospeed
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaReplaySpeed() int32 {
	if m != nil && m.DotaReplaySpeed != nil {
		return *m.DotaReplaySpeed
	}
	return 0
}

func init() {
	proto.RegisterEnum("dota.P2P_Messages", P2P_Messages_name, P2P_Messages_value)
	proto.RegisterEnum("dota.CP2P_Voice_Handler_Flags", CP2P_Voice_Handler_Flags_name, CP2P_Voice_Handler_Flags_value)
	proto.RegisterType((*CP2P_TextMessage)(nil), "dota.CP2P_TextMessage")
	proto.RegisterType((*CSteam_Voice_Encoding)(nil), "dota.CSteam_Voice_Encoding")
	proto.RegisterType((*CP2P_Voice)(nil), "dota.CP2P_Voice")
	proto.RegisterType((*CP2P_Ping)(nil), "dota.CP2P_Ping")
	proto.RegisterType((*CP2P_VRAvatarPosition)(nil), "dota.CP2P_VRAvatarPosition")
	proto.RegisterType((*CP2P_VRAvatarPosition_COrientation)(nil), "dota.CP2P_VRAvatarPosition.COrientation")
	proto.RegisterType((*CP2P_WatchSynchronization)(nil), "dota.CP2P_WatchSynchronization")
}

func init() { proto.RegisterFile("c_peer2peer_netmessages.proto", fileDescriptor_1fec9daf05519245) }

var fileDescriptor_1fec9daf05519245 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x51, 0x6f, 0x23, 0x35,
	0x10, 0xc7, 0xd9, 0x4d, 0xd2, 0x26, 0xd3, 0xf4, 0xba, 0x67, 0xd2, 0xb2, 0xe9, 0xa9, 0x22, 0x2c,
	0x12, 0x44, 0x7d, 0xc8, 0x43, 0x24, 0x10, 0x3c, 0xe6, 0x02, 0xdc, 0x45, 0x22, 0x22, 0x6c, 0xaa,
	0xbb, 0x47, 0xcb, 0x5d, 0x8f, 0x36, 0x56, 0x37, 0xf6, 0xca, 0x76, 0xd2, 0x06, 0x09, 0xa9, 0x40,
	0xdf, 0xf8, 0x20, 0x7c, 0x4d, 0x64, 0x6f, 0x52, 0xda, 0xd0, 0x97, 0x28, 0xf3, 0xff, 0xcd, 0xdf,
	0x33, 0xb3, 0x1e, 0xc3, 0x45, 0x46, 0x4b, 0x44, 0x3d, 0x74, 0x3f, 0x54, 0xa2, 0x5d, 0xa2, 0x31,
	0x2c, 0x47, 0x33, 0x28, 0xb5, 0xb2, 0x8a, 0xd4, 0xb9, 0xb2, 0xec, 0xfc, 0xf5, 0xff, 0xc0, 0xf9,
	0x99, 0x44, 0x7b, 0xab, 0xf4, 0xcd, 0x35, 0x33, 0x68, 0x37, 0xe5, 0x4e, 0x4f, 0xbe, 0x82, 0x68,
	0x3c, 0x1b, 0xce, 0xe8, 0x15, 0xde, 0xd9, 0x69, 0x65, 0x21, 0x04, 0xea, 0x16, 0xef, 0x6c, 0x1c,
	0xf4, 0x82, 0x7e, 0x3b, 0xf5, 0xff, 0x93, 0x6f, 0xe1, 0x74, 0x3c, 0xb7, 0xc8, 0x96, 0xf4, 0x83,
	0x12, 0x19, 0xd2, 0x1f, 0x65, 0xa6, 0xb8, 0x90, 0x39, 0xb9, 0x00, 0x58, 0x7b, 0x85, 0x33, 0xcb,
	0xb6, 0x96, 0x96, 0x57, 0x7e, 0x60, 0x96, 0x25, 0x0f, 0x01, 0x80, 0x2f, 0xe0, 0x6d, 0xe4, 0x12,
	0x1a, 0x6c, 0xc5, 0x85, 0xf2, 0x89, 0x47, 0xc3, 0xce, 0xc0, 0xf5, 0x3b, 0x18, 0x4f, 0x4d, 0xee,
	0xf9, 0xc8, 0xb1, 0xb4, 0x4a, 0x21, 0x5f, 0xc3, 0xc9, 0xb5, 0x56, 0x8c, 0x67, 0xcc, 0x58, 0x9a,
	0x6b, 0xb5, 0x2a, 0xe3, 0xb0, 0x17, 0xf4, 0x8f, 0xd3, 0x57, 0x8f, 0xf2, 0x3b, 0xa7, 0x26, 0x5f,
	0xc0, 0xf1, 0x7b, 0x26, 0x79, 0x81, 0x9a, 0xfe, 0x54, 0xb0, 0xdc, 0x90, 0x08, 0xda, 0xb3, 0x82,
	0x6d, 0x90, 0x53, 0x7f, 0x60, 0x14, 0x24, 0x63, 0x68, 0xf9, 0x2e, 0x66, 0xae, 0xe5, 0x37, 0xd0,
	0x32, 0x28, 0x39, 0xb5, 0x62, 0x89, 0x71, 0xd0, 0x0b, 0xfb, 0xf5, 0xb4, 0xe9, 0x84, 0x2b, 0xb1,
	0x44, 0xd2, 0x85, 0xa6, 0x30, 0x54, 0x63, 0x59, 0x6c, 0xe2, 0xb0, 0x17, 0xf6, 0x9b, 0xe9, 0xa1,
	0x30, 0xa9, 0x0b, 0x93, 0xbf, 0x43, 0x38, 0xad, 0x66, 0x49, 0x47, 0x6b, 0x66, 0x99, 0x9e, 0x29,
	0x23, 0xac, 0x50, 0x92, 0xbc, 0x03, 0xb8, 0x56, 0x7c, 0x43, 0x4b, 0xa6, 0xad, 0x89, 0x83, 0x5e,
	0xad, 0x7f, 0x34, 0xec, 0x6f, 0x67, 0x7b, 0xc9, 0x30, 0x18, 0xff, 0xa2, 0x05, 0x4a, 0xcb, 0x5c,
	0x90, 0xb6, 0x9c, 0x77, 0xe6, 0xac, 0xe4, 0x14, 0x0e, 0x16, 0xcc, 0x52, 0xc1, 0xfd, 0xa8, 0x8d,
	0xb4, 0xb1, 0x60, 0x76, 0xc2, 0x5d, 0x53, 0x26, 0x43, 0x89, 0x0e, 0xd4, 0x3c, 0x38, 0xf4, 0xf1,
	0x84, 0x93, 0xcf, 0xe1, 0xe8, 0x56, 0xe9, 0x82, 0x53, 0x93, 0xb1, 0x02, 0xe3, 0xba, 0xa7, 0xe0,
	0xa5, 0xb9, 0x53, 0xce, 0x3f, 0x40, 0xfb, 0x69, 0x35, 0x92, 0x40, 0xad, 0x54, 0x66, 0x7b, 0x01,
	0xd1, 0x93, 0x0b, 0xc0, 0xcc, 0x2a, 0x9d, 0x3a, 0xe8, 0x72, 0x98, 0xcc, 0x7d, 0x0f, 0xcf, 0x72,
	0x7e, 0x1d, 0xc9, 0xbc, 0xc0, 0xd4, 0xc1, 0xe4, 0x9f, 0x1a, 0x74, 0xfd, 0x70, 0x1f, 0x99, 0xcd,
	0x16, 0xf3, 0x8d, 0xcc, 0x16, 0x5a, 0x49, 0xf1, 0x5b, 0x55, 0xe5, 0x0d, 0xb4, 0x38, 0x2e, 0x15,
	0xb5, 0x22, 0xbb, 0xf1, 0xb5, 0x1a, 0x69, 0xd3, 0x09, 0x57, 0x22, 0xbb, 0x21, 0x67, 0x70, 0x50,
	0xb2, 0x95, 0xc1, 0x6a, 0xca, 0x66, 0xba, 0x8d, 0xc8, 0x37, 0xf0, 0x99, 0x5d, 0xd3, 0x42, 0x18,
	0x8b, 0x92, 0x56, 0x5b, 0x25, 0x24, 0x17, 0x19, 0x9a, 0xed, 0xd4, 0x1d, 0xbb, 0xfe, 0xd9, 0x53,
	0xbf, 0x2d, 0x93, 0x8a, 0x91, 0x01, 0x7c, 0xea, 0x3a, 0xa4, 0xa6, 0xc4, 0xcc, 0x32, 0xab, 0x34,
	0x5d, 0x2a, 0xbe, 0xfb, 0x14, 0xaf, 0x1d, 0x9a, 0xef, 0xc8, 0x54, 0x71, 0x24, 0x53, 0xf8, 0x72,
	0x2f, 0xff, 0xd6, 0x8d, 0x20, 0x64, 0x4e, 0x1f, 0x37, 0x0b, 0x75, 0xdc, 0xf0, 0xfe, 0xde, 0x33,
	0xff, 0xc7, 0x6d, 0xe2, 0xdb, 0xff, 0xf2, 0xc8, 0xf7, 0xd0, 0xdd, 0x3b, 0x6e, 0x81, 0x5a, 0xb9,
	0xce, 0xf1, 0x2e, 0x3e, 0xf0, 0x87, 0x9c, 0x3d, 0x3b, 0xe4, 0x3d, 0x6a, 0x35, 0x71, 0x94, 0x7c,
	0x07, 0xf1, 0x9e, 0x95, 0xad, 0xac, 0x32, 0x25, 0x22, 0x8f, 0x0f, 0x5f, 0x70, 0x8e, 0x76, 0x94,
	0x5c, 0x82, 0x1f, 0xcc, 0x2f, 0x2a, 0xdb, 0xd0, 0xca, 0xd2, 0xf4, 0x96, 0x13, 0x07, 0x52, 0xaf,
	0xcf, 0x9d, 0x7c, 0xf9, 0x3b, 0xb4, 0xdd, 0x3d, 0x6d, 0x9f, 0xb7, 0x21, 0x1d, 0x38, 0x29, 0x87,
	0xe5, 0xd3, 0x27, 0x1f, 0xdd, 0x87, 0xe4, 0x15, 0xb4, 0x9c, 0xea, 0xbf, 0x6c, 0xf4, 0x47, 0x48,
	0x8e, 0xa1, 0xe9, 0x62, 0xf7, 0x62, 0xa2, 0x3f, 0x43, 0xd2, 0x85, 0x8e, 0xc7, 0x7b, 0x9b, 0x1c,
	0xfd, 0x15, 0x92, 0x0b, 0x88, 0x1d, 0x7a, 0x69, 0x0f, 0xa2, 0x87, 0xf0, 0x6d, 0xed, 0x3e, 0xf8,
	0xe4, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x36, 0xaa, 0x22, 0xb8, 0x04, 0x00, 0x00,
}
