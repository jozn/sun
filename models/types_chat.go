package models

import "ms/sun/helper"

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
