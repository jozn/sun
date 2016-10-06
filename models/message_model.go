package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
)

type _messageModelImple int

var MessageModel _messageModelImple

func (e _messageModelImple) SendMessage(ToUserId int, msg MessagesTableFromClient) {
	if msg.CreatedMs == 0 {
		t := helper.TimeNowMs()
		msg.CreatedMs = t
	}

	msgT := Message{
		FromUserID: msg.UserId,
		ToUserId:   ToUserId,
		MessageKey: msg.MessageKey,
		Data:       helper.ToJson(msg),
		TimeMs:     msg.CreatedMs,
	}

	err := msgT.Insert(base.DB)
	helper.DebugPrintln(err)
}

///////////// Utils //////////////////

//format "p142_1569"
func UserIdsToRoomKey(UserId1, UserId2 int) string {
	if UserId2 < UserId1 {
		UserId1, UserId2 = UserId2, UserId1
	}
	return fmt.Sprintf("p%d_%d", UserId1, UserId2)
}
