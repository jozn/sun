package models

//TODO delete this?
type tMsgRecivedToUser struct {
	userId int
	msgUid int
}

var chanMsgRecivedToUserBuffer = make(chan tMsgRecivedToUser, 20000)
