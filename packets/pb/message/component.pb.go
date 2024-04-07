// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/message/component.proto

package message

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type Attr struct {
	CodePage       int32  `protobuf:"varint,1,opt"`
	Time           int32  `protobuf:"varint,2,opt"`
	Random         int32  `protobuf:"varint,3,opt"`
	Color          int32  `protobuf:"varint,4,opt"`
	Size           int32  `protobuf:"varint,5,opt"`
	Effect         int32  `protobuf:"varint,6,opt"`
	CharSet        int32  `protobuf:"varint,7,opt"`
	PitchAndFamily int32  `protobuf:"varint,8,opt"`
	FontName       string `protobuf:"bytes,9,opt"`
	ReserveData    []byte `protobuf:"bytes,10,opt"`
}

type GroupInfo struct {
	GroupCode     uint64 `protobuf:"varint,1,opt"`
	GroupType     int32  `protobuf:"varint,2,opt"`
	GroupInfoSeq  int64  `protobuf:"varint,3,opt"`
	GroupCard     string `protobuf:"bytes,4,opt"`
	GroupLevel    int32  `protobuf:"varint,5,opt"`
	GroupCardType int32  `protobuf:"varint,6,opt"`
	GroupName     []byte `protobuf:"bytes,7,opt"`
}

type MutilTransHead struct {
	Status int32 `protobuf:"varint,1,opt"`
	MsgId  int32 `protobuf:"varint,2,opt"`
	_      [0]func()
}

type NotOnlineFile struct {
	FileType      proto.Option[int32]  `protobuf:"varint,1,opt"`
	Sig           []byte               `protobuf:"bytes,2,opt"`
	FileUuid      proto.Option[string] `protobuf:"bytes,3,opt"`
	FileMd5       []byte               `protobuf:"bytes,4,opt"`
	FileName      proto.Option[string] `protobuf:"bytes,5,opt"`
	FileSize      proto.Option[int64]  `protobuf:"varint,6,opt"`
	Note          []byte               `protobuf:"bytes,7,opt"`
	Reserved      proto.Option[int32]  `protobuf:"varint,8,opt"`
	Subcmd        proto.Option[int32]  `protobuf:"varint,9,opt"`
	MicroCloud    proto.Option[int32]  `protobuf:"varint,10,opt"`
	BytesFileUrls [][]byte             `protobuf:"bytes,11,rep"`
	DownloadFlag  proto.Option[int32]  `protobuf:"varint,12,opt"`
	DangerEvel    proto.Option[int32]  `protobuf:"varint,50,opt"`
	LifeTime      proto.Option[int32]  `protobuf:"varint,51,opt"`
	UploadTime    proto.Option[int32]  `protobuf:"varint,52,opt"`
	AbsFileType   proto.Option[int32]  `protobuf:"varint,53,opt"`
	ClientType    proto.Option[int32]  `protobuf:"varint,54,opt"`
	ExpireTime    proto.Option[int32]  `protobuf:"varint,55,opt"`
	PbReserve     []byte               `protobuf:"bytes,56,opt"`
	FileHash      proto.Option[string] `protobuf:"bytes,57,opt"`
}

type Ptt struct {
	FileType      int32    `protobuf:"varint,1,opt"`
	SrcUin        uint64   `protobuf:"varint,2,opt"`
	FileUuid      string   `protobuf:"bytes,3,opt"`
	FileMd5       []byte   `protobuf:"bytes,4,opt"`
	FileName      string   `protobuf:"bytes,5,opt"`
	FileSize      int32    `protobuf:"varint,6,opt"`
	Reserve       []byte   `protobuf:"bytes,7,opt"`
	FileId        int32    `protobuf:"varint,8,opt"`
	ServerIp      int32    `protobuf:"varint,9,opt"`
	ServerPort    int32    `protobuf:"varint,10,opt"`
	BoolValid     bool     `protobuf:"varint,11,opt"`
	Signature     []byte   `protobuf:"bytes,12,opt"`
	Shortcut      []byte   `protobuf:"bytes,13,opt"`
	FileKey       []byte   `protobuf:"bytes,14,opt"`
	MagicPttIndex int32    `protobuf:"varint,15,opt"`
	VoiceSwitch   int32    `protobuf:"varint,16,opt"`
	PttUrl        []byte   `protobuf:"bytes,17,opt"`
	GroupFileKey  string   `protobuf:"bytes,18,opt"`
	Time          int32    `protobuf:"varint,19,opt"`
	DownPara      []byte   `protobuf:"bytes,20,opt"`
	Format        int32    `protobuf:"varint,29,opt"`
	PbReserve     []byte   `protobuf:"bytes,30,opt"`
	BytesPttUrls  [][]byte `protobuf:"bytes,31,rep"`
	DownloadFlag  int32    `protobuf:"varint,32,opt"`
}

type RichText struct {
	Attr          *Attr          `protobuf:"bytes,1,opt"`
	Elems         []*Elem        `protobuf:"bytes,2,rep"`
	NotOnlineFile *NotOnlineFile `protobuf:"bytes,3,opt"`
	Ptt           *Ptt           `protobuf:"bytes,4,opt"`
}

type ButtonExtra struct {
	Data *KeyboardData `protobuf:"bytes,1,opt"`
	_    [0]func()
}

type KeyboardData struct {
	Rows []*Row `protobuf:"bytes,1,rep"`
}

type Row struct {
	Buttons []*Button `protobuf:"bytes,1,rep"`
}

type Button struct {
	ID         string      `protobuf:"bytes,1,opt"`
	RenderData *RenderData `protobuf:"bytes,2,opt"`
	Action     *Action     `protobuf:"bytes,3,opt"`
	_          [0]func()
}

type RenderData struct {
	Label        string `protobuf:"bytes,1,opt"`
	VisitedLabel string `protobuf:"bytes,2,opt"`
	Style        int32  `protobuf:"varint,3,opt"`
	_            [0]func()
}

type Action struct {
	Type       int32       `protobuf:"varint,1,opt"`
	Permission *Permission `protobuf:"bytes,2,opt"`
	// uint64 ClickLimit = 3;
	UnsupportTips string `protobuf:"bytes,4,opt"`
	Data          string `protobuf:"bytes,5,opt"`
	Reply         bool   `protobuf:"varint,7,opt"`
	Enter         bool   `protobuf:"varint,8,opt"`
	_             [0]func()
}

type Permission struct {
	Type           int32    `protobuf:"varint,1,opt"`
	SpecifyRoleIds []string `protobuf:"bytes,2,rep"`
	SpecifyUserIds []string `protobuf:"bytes,3,rep"`
}

type FileExtra struct {
	File *NotOnlineFile `protobuf:"bytes,1,opt"`
	_    [0]func()
}

type GroupFileExtra struct {
	Field1   uint32               `protobuf:"varint,1,opt"`
	FileName string               `protobuf:"bytes,2,opt"`
	Display  string               `protobuf:"bytes,3,opt"`
	Inner    *GroupFileExtraInner `protobuf:"bytes,7,opt"`
	_        [0]func()
}

type GroupFileExtraInner struct {
	Info *GroupFileExtraInfo `protobuf:"bytes,2,opt"`
	_    [0]func()
}

type GroupFileExtraInfo struct {
	BusId    uint32 `protobuf:"varint,1,opt"`
	FileId   string `protobuf:"bytes,2,opt"`
	FileSize uint64 `protobuf:"varint,3,opt"`
	FileName string `protobuf:"bytes,4,opt"`
	Field5   uint32 `protobuf:"varint,5,opt"`
	Field7   string `protobuf:"bytes,7,opt"`
	FileMd5  string `protobuf:"bytes,8,opt"`
	_        [0]func()
}

type ImageExtraUrl struct {
	OrigUrl string `protobuf:"bytes,30,opt"`
	_       [0]func()
}

type PokeExtra struct {
	Type   uint32 `protobuf:"varint,1,opt"`
	Field7 uint32 `protobuf:"varint,7,opt"`
	Field8 uint32 `protobuf:"varint,8,opt"`
	_      [0]func()
}
