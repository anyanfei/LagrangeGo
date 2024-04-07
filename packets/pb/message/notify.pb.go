// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/message/notify.proto

package message

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type FriendRecall struct {
	Info            *FriendRecallInfo `protobuf:"bytes,1,opt"`
	InstId          uint32            `protobuf:"varint,2,opt"`
	AppId           uint32            `protobuf:"varint,3,opt"`
	LongMessageFlag uint32            `protobuf:"varint,4,opt"`
	Reserved        []byte            `protobuf:"bytes,5,opt"`
}

type FriendRecallInfo struct {
	FromUid  string `protobuf:"bytes,1,opt"`
	ToUid    string `protobuf:"bytes,2,opt"`
	Sequence uint32 `protobuf:"varint,3,opt"`
	NewId    uint64 `protobuf:"varint,4,opt"`
	Time     uint32 `protobuf:"varint,5,opt"`
	Random   uint32 `protobuf:"varint,6,opt"`
	PkgNum   uint32 `protobuf:"varint,7,opt"`
	PkgIndex uint32 `protobuf:"varint,8,opt"`
	DivSeq   uint32 `protobuf:"varint,9,opt"`
	_        [0]func()
}

type FriendRequest struct {
	Info *FriendRequestInfo `protobuf:"bytes,1,opt"`
	_    [0]func()
}

type FriendRequestInfo struct {
	TargetUid string `protobuf:"bytes,1,opt"`
	SourceUid string `protobuf:"bytes,2,opt"`
	Message   string `protobuf:"bytes,10,opt"`
	Source    string `protobuf:"bytes,11,opt"`
	_         [0]func()
}

type GroupAdmin struct {
	GroupUin  uint32          `protobuf:"varint,1,opt"`
	Flag      uint32          `protobuf:"varint,2,opt"`
	IsPromote bool            `protobuf:"varint,3,opt"`
	Body      *GroupAdminBody `protobuf:"bytes,4,opt"`
	_         [0]func()
}

type GroupAdminBody struct {
	ExtraDisable *GroupAdminExtra `protobuf:"bytes,1,opt"`
	ExtraEnable  *GroupAdminExtra `protobuf:"bytes,2,opt"`
	_            [0]func()
}

type GroupAdminExtra struct {
	AdminUid  string `protobuf:"bytes,1,opt"`
	IsPromote bool   `protobuf:"varint,2,opt"`
	_         [0]func()
}

type GroupChange struct {
	GroupUin     uint32               `protobuf:"varint,1,opt"`
	Flag         uint32               `protobuf:"varint,2,opt"`
	MemberUid    string               `protobuf:"bytes,3,opt"`
	DecreaseType uint32               `protobuf:"varint,4,opt"`
	OperatorUid  proto.Option[string] `protobuf:"bytes,5,opt"`
	IncreaseType uint32               `protobuf:"varint,6,opt"`
	Field7       []byte               `protobuf:"bytes,7,opt"`
}

type GroupInvitation struct {
	Cmd  int32           `protobuf:"varint,1,opt"`
	Info *InvitationInfo `protobuf:"bytes,2,opt"`
	_    [0]func()
}

type InvitationInfo struct {
	Inner *InvitationInner `protobuf:"bytes,1,opt"`
	_     [0]func()
}

type InvitationInner struct {
	GroupUin   uint32 `protobuf:"varint,1,opt"`
	Field2     uint32 `protobuf:"varint,2,opt"`
	Field3     uint32 `protobuf:"varint,3,opt"`
	Field4     uint32 `protobuf:"varint,4,opt"`
	TargetUid  string `protobuf:"bytes,5,opt"`
	InvitorUid string `protobuf:"bytes,6,opt"`
	Field7     uint32 `protobuf:"varint,7,opt"`
	Field9     uint32 `protobuf:"varint,9,opt"`
	Field10    []byte `protobuf:"bytes,10,opt"`
	Field11    uint32 `protobuf:"varint,11,opt"`
	Field12    string `protobuf:"bytes,12,opt"`
}

type GroupInvite struct {
	GroupInvite uint32 `protobuf:"varint,1,opt"`
	Field2      uint32 `protobuf:"varint,2,opt"` // 1
	Field3      uint32 `protobuf:"varint,3,opt"` // 4
	Field4      uint32 `protobuf:"varint,4,opt"` // 0
	InvitorUid  string `protobuf:"bytes,5,opt"`
	Hashes      []byte `protobuf:"bytes,6,opt"`
}

type GroupJoin struct {
	GroupUin  uint32 `protobuf:"varint,1,opt"`
	Field2    uint32 `protobuf:"varint,2,opt"`
	TargetUid string `protobuf:"bytes,3,opt"`
	Field4    uint32 `protobuf:"varint,4,opt"`
	Field6    uint32 `protobuf:"varint,6,opt"`
	Field7    string `protobuf:"bytes,7,opt"`
	Field8    uint32 `protobuf:"varint,8,opt"`
	Field9    []byte `protobuf:"bytes,9,opt"`
}

type GroupMute struct {
	GroupUin    uint32               `protobuf:"varint,1,opt"`
	SubType     uint32               `protobuf:"varint,2,opt"`
	Field3      uint32               `protobuf:"varint,3,opt"`
	OperatorUid proto.Option[string] `protobuf:"bytes,4,opt"`
	Data        *GroupMuteData       `protobuf:"bytes,5,opt"`
	_           [0]func()
}

type GroupMuteData struct {
	Timestamp uint32          `protobuf:"varint,1,opt"`
	Type      uint32          `protobuf:"varint,2,opt"`
	State     *GroupMuteState `protobuf:"bytes,3,opt"`
	_         [0]func()
}

type GroupMuteState struct {
	TargetUid proto.Option[string] `protobuf:"bytes,1,opt"`
	Duration  uint32               `protobuf:"varint,2,opt"` // uint.MaxValue = Mute else Lift
	_         [0]func()
}

type GroupRecall struct {
	OperatorUid    proto.Option[string] `protobuf:"bytes,1,opt"`
	RecallMessages []*RecallMessage     `protobuf:"bytes,3,rep"`
	UserDef        []byte               `protobuf:"bytes,5,opt"`
	GroupType      int32                `protobuf:"varint,6,opt"`
	OpType         int32                `protobuf:"varint,7,opt"`
}

type RecallMessage struct {
	Sequence  uint64 `protobuf:"varint,1,opt"`
	Time      uint32 `protobuf:"varint,2,opt"`
	Random    uint32 `protobuf:"varint,3,opt"`
	Type      uint32 `protobuf:"varint,4,opt"`
	Flag      uint32 `protobuf:"varint,5,opt"`
	AuthorUid string `protobuf:"bytes,6,opt"`
	_         [0]func()
}

type NotifyMessageBody struct {
	Type     uint32       `protobuf:"varint,1,opt"`
	GroupUin uint32       `protobuf:"varint,4,opt"`
	Recall   *GroupRecall `protobuf:"bytes,11,opt"`
	_        [0]func()
}
