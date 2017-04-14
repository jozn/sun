package models

import "ms/sun/helper"

/*
type Message struct {
	Id         int    `json:"Id"`         // Id -
	ToUserId   int    `json:"ToUserId"`   // ToUserId -
	RoomKey    string `json:"RoomKey"`    // RoomKey -
	MessageKey string `json:"MessageKey"` // MessageKey -
	FromUserID int    `json:"FromUserID"` // FromUserID -
	Data       string `json:"Data"`       // Data -
	TimeMs     int    `json:"TimeMs"`     // TimeMs -

	// xo fields
	_exists, _deleted bool
}
*/

func (m *Message) FromClientMessage(toUser int, msg MessagesTableFromClient) {
	m.FromUserID = msg.UserId
	m.ToUserId = toUser
	m.MessageKey = msg.MessageKey
	m.RoomKey = msg.RoomKey
	m.Data = helper.ToJson(msg)
	m.TimeMs = msg.CreatedMs
}

func (m *Message) FromClientMessageOptimized(toUser int, msg MessagesTableFromClient) {
	m.FromUserID = msg.UserId
	m.ToUserId = toUser
	m.MessageKey = msg.MessageKey
	m.RoomKey = msg.RoomKey
	//m.Data=       helper.ToJson(msg)
	m.TimeMs = msg.CreatedMs
}

/*
type MsgDeletedFromServer struct {
	Id         int    `json:"Id"`         // Id -
	ToUserId   int    `json:"ToUserId"`   // ToUserId -
	MsgKey     string `json:"MsgKey"`     // MsgKey -
	PeerUserId int    `json:"PeerUserId"` // PeerUserId -
	RoomKey    string `json:"RoomKey"`    // RoomKey -
	AtTime     int    `json:"AtTime"`     // AtTime -

	// xo fields
	_exists, _deleted bool
}

type MsgReceivedToPeer struct {
	Id         int    `json:"Id"`         // Id -
	ToUserId   int    `json:"ToUserId"`   // ToUserId -
	MsgKey     string `json:"MsgKey"`     // MsgKey -
	RoomKey    string `json:"RoomKey"`    // RoomKey -
	PeerUserId int    `json:"PeerUserId"` // PeerUserId -
	AtTime     int    `json:"AtTime"`     // AtTime -

	// xo fields
	_exists, _deleted bool
}

type MsgSeenByPeer struct {
	Id         int    `json:"Id"`         // Id -
	ToUserId   int    `json:"ToUserId"`   // ToUserId -
	MsgKey     string `json:"MsgKey"`     // MsgKey -
	RoomKey    string `json:"RoomKey"`    // RoomKey -
	PeerUserId int    `json:"PeerUserId"` // PeerUserId -
	AtTime     int    `json:"AtTime"`     // AtTime -

	// xo fields
	_exists, _deleted bool
}
*/

/////////////////////// Others ////////////////
type MessagesTableFromClient struct {
	MessageKey    string
	RoomKey       string
	UserId        int
	RoomTypeId    int
	MessageTypeId int
	Text          string
	ExtraJson     string
	CreatedMs     int

	MsgFile_LocalSrc string
	MsgFile_Status   int
	MsgFile          *MsgFile //not in sqlite just json
}

type MsgFile struct {
	LocalSrc  string
	Hash      string
	ServerSrc string
	FileType  int
	Status    int
	Origin    int
	Thumb64   string
	Name      string
	Size      int
	Duration  int
	Height    int
	Width     int
	Extension string
	CreatedMs int
}

/////// Communication ///////

type MessageSyncMeta struct {
	MessageKey string
	RoomKey    string
	ByUserId   int
	AtTimeMs   int // this is client time
	ExtraData  interface{}
}

/////////////////////////
// Orma
/*type PhoneContact struct {
	Id                    int    `json:"Id"`                    // Id -
	PhoneDisplayName      string `json:"PhoneDisplayName"`      // PhoneDisplayName -
	PhoneFamilyName       string `json:"PhoneFamilyName"`       // PhoneFamilyName -
	PhoneNumber           string `json:"PhoneNumber"`           // PhoneNumber -
	PhoneNormalizedNumber string `json:"PhoneNormalizedNumber"` // PhoneNormalizedNumber -
	PhoneContactRowId     int    `json:"PhoneContactRowId"`     // PhoneContactRowId -
	UserId                int    `json:"UserId"`                // UserId -
	DeviceUuidId          int    `json:"DeviceUuidId"`          // DeviceUuidId -
	CreatedTime           int    `json:"CreatedTime"`           // CreatedTime -
	UpdatedTime           int    `json:"UpdatedTime"`           // UpdatedTime -

	// xo fields
	_exists, _deleted bool
}*/
