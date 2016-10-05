package models

type Message struct {
	Id         int    `json:"Id"`         // Id -
	MessageKey string `json:"MessageKey"` // MessageKey -
	ToUserId   int    `json:"ToUserId"`   // ToUserId -
	FromUserID int    `json:"FromUserID"` // FromUserID -
	Data       string `json:"Data"`       // Text -
	TimeMs     int    `json:"TimeMs"`     // TimeMs -

	// xo fields
	_exists, _deleted bool
}

type MsgDeletedFromServer struct {
	Id         int    `json:"Id"`         // Id -
	ToUserId   int    `json:"ToUserId"`   // ToUserId -
	MsgKey     string `json:"MsgKey"`     // MsgKey -
	PeerUserId int    `json:"PeerUserId"` // PeerUserId -
	Time       int    `json:"Time"`       // Time -

	// xo fields
	_exists, _deleted bool
}

type MsgReceivedToPeer struct {
	Id       int    `json:"Id"`       // Id -
	ToUserId int    `json:"ToUserId"` // ToUserId -
	MsgKey   string `json:"MsgKey"`   // MsgKey -
	Time     int    `json:"Time"`     // Time -

	// xo fields
	_exists, _deleted bool
}

type MsgSeenByPeer struct {
	Id         int    `json:"Id"`         // Id -
	ToUserId   int    `json:"ToUserId"`   // ToUserId -
	MsgKey     string `json:"MsgKey"`     // MsgKey -
	PeerUserId int    `json:"PeerUserId"` // PeerUserId -
	Time       int    `json:"Time"`       // Time -

	// xo fields
	_exists, _deleted bool
}
