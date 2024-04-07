// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/message/message.proto

package message

import (
	system "github.com/Redmomn/LagrangeGo/packets/pb/system"
	proto "github.com/RomiChan/protobuf/proto"
)

type ContentHead struct {
	Type      uint32               `protobuf:"varint,1,opt"`
	SubType   proto.Option[uint32] `protobuf:"varint,2,opt"`
	DivSeq    proto.Option[uint32] `protobuf:"varint,3,opt"`
	MsgId     proto.Option[uint64] `protobuf:"varint,4,opt"`
	Sequence  proto.Option[uint32] `protobuf:"varint,5,opt"`
	TimeStamp proto.Option[uint64] `protobuf:"varint,6,opt"`
	Field7    proto.Option[uint64] `protobuf:"varint,7,opt"`
	Field8    proto.Option[uint32] `protobuf:"varint,8,opt"`
	Field9    proto.Option[uint32] `protobuf:"varint,9,opt"`
	NewId     proto.Option[uint64] `protobuf:"varint,12,opt"`
	Foward    *ForwardHead         `protobuf:"bytes,15,opt"`
	_         [0]func()
}

type MessageBody struct {
	RichText          *RichText `protobuf:"bytes,1,opt"`
	MsgContent        []byte    `protobuf:"bytes,2,opt"`
	MsgEncryptContent []byte    `protobuf:"bytes,3,opt"`
}

type GroupRecallMsg struct {
	Type     uint32                `protobuf:"varint,1,opt"`
	GroupUin uint32                `protobuf:"varint,2,opt"`
	Field3   *GroupRecallMsgField3 `protobuf:"bytes,3,opt"`
	Field4   *GroupRecallMsgField4 `protobuf:"bytes,4,opt"`
	_        [0]func()
}

type GroupRecallMsgField3 struct {
	Sequence uint32 `protobuf:"varint,1,opt"`
	Random   uint32 `protobuf:"varint,2,opt"`
	Field3   uint32 `protobuf:"varint,3,opt"`
	_        [0]func()
}

type GroupRecallMsgField4 struct {
	Field1 uint32 `protobuf:"varint,1,opt"`
	_      [0]func()
}

type Message struct {
	RoutingHead *RoutingHead         `protobuf:"bytes,1,opt"`
	ContentHead *ContentHead         `protobuf:"bytes,2,opt"`
	Body        *MessageBody         `protobuf:"bytes,3,opt"`
	Seq         proto.Option[uint32] `protobuf:"varint,4,opt"`
	Rand        proto.Option[uint32] `protobuf:"varint,5,opt"`
	SyncCookie  []byte               `protobuf:"bytes,6,opt"`
	// optional AppShareInfo AppShare = 7;
	Via         proto.Option[uint32] `protobuf:"varint,8,opt"`
	DataStatist proto.Option[uint32] `protobuf:"varint,9,opt"`
	// optional MultiMsgAssist MultiMsgAssist = 10;
	// optional InputNotifyInfo InputNotifyInfo = 11;
	Ctrl *MessageControl `protobuf:"bytes,12,opt"`
	// optional ReceiptReq ReceiptReq = 13;
	MultiSendSeq uint32 `protobuf:"varint,14,opt"`
}

type MessageControl struct {
	MsgFlag int32 `protobuf:"varint,1,opt"`
	_       [0]func()
}

type PushMsg struct {
	Message     *PushMsgBody        `protobuf:"bytes,1,opt"`
	Status      proto.Option[int32] `protobuf:"varint,3,opt"`
	NtEvent     *system.NTSysEvent  `protobuf:"bytes,4,opt"`
	PingFLag    proto.Option[int32] `protobuf:"varint,5,opt"`
	GeneralFlag proto.Option[int32] `protobuf:"varint,9,opt"`
	_           [0]func()
}

type PushMsgBody struct {
	ResponseHead *ResponseHead `protobuf:"bytes,1,opt"`
	ContentHead  *ContentHead  `protobuf:"bytes,2,opt"`
	Body         *MessageBody  `protobuf:"bytes,3,opt"`
	_            [0]func()
}

type ResponseHead struct {
	FromUin uint32               `protobuf:"varint,1,opt"`
	FromUid proto.Option[string] `protobuf:"bytes,2,opt"`
	Type    uint32               `protobuf:"varint,3,opt"`
	SigMap  uint32               `protobuf:"varint,4,opt"` // 鬼知道是啥
	ToUin   uint32               `protobuf:"varint,5,opt"`
	ToUid   proto.Option[string] `protobuf:"bytes,6,opt"`
	Forward *ResponseForward     `protobuf:"bytes,7,opt"`
	Grp     *ResponseGrp         `protobuf:"bytes,8,opt"`
	_       [0]func()
}

type RoutingHead struct {
	C2C        *C2C        `protobuf:"bytes,1,opt"`
	Grp        *Grp        `protobuf:"bytes,2,opt"`
	GrpTmp     *GrpTmp     `protobuf:"bytes,3,opt"`
	WpaTmp     *WPATmp     `protobuf:"bytes,6,opt"`
	Trans0X211 *Trans0X211 `protobuf:"bytes,15,opt"`
	_          [0]func()
}

type SsoReadedReport struct {
	Group *SsoReadedReportGroup `protobuf:"bytes,1,opt"`
	C2C   *SsoReadedReportC2C   `protobuf:"bytes,2,opt"`
	_     [0]func()
}

type SsoReadedReportC2C struct {
	TargetUid     proto.Option[string] `protobuf:"bytes,2,opt"`
	Time          uint32               `protobuf:"varint,3,opt"`
	StartSequence uint32               `protobuf:"varint,4,opt"`
	_             [0]func()
}

type SsoReadedReportGroup struct {
	GroupUin      uint32 `protobuf:"varint,1,opt"`
	StartSequence uint32 `protobuf:"varint,2,opt"`
	_             [0]func()
}
