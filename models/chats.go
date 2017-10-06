package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
)

func Chat_GetChatListForUser(me int) (res []*x.PB_ChatView, err error) {
	rows, err := x.NewChat_Selector().UserId_Eq(me).GetRows(base.DB)
	if err != nil {
		return
	}

	chatIds := make(map[string]bool)
    for _, r := range rows {
        chatIds[r.ChatKey] = true
    }

    res = ViewChat_GetChatViewList_map(me, chatIds)

	/*for _, r := range rows {
		chat := PBConv_Chat_To_PB_ChatView(r)
		res = append(res, chat)
	}*/

	return
}

//fixme :1 this must first load from direct-to_message 2: fix chatId
func Chat_GetMessageListForRoom(roomkey string, lastId int) (res []*x.PB_MessageView, err error) {
	rows, err := x.NewDirectMessage_Selector().RoomKey_Eq(roomkey).OrderBy_MessageId_Desc().GetRows(base.DB)
	if err != nil {
		return
	}

	for _, r := range rows {
		chat := PBConv_DirectMessage_to_PB_MessageView(r,"")
		res = append(res, chat)
	}

	return
}
