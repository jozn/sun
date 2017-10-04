package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

//d(lower userId)_(HigherUserId)
func UsersToRoomKey(me int, peer int) string {
	if me > peer {
		me, peer = peer, me
	}
	return fmt.Sprintf("d%d_%d", me, peer)
}

func UsersToChatKey(me int, peer int) string {
    return fmt.Sprintf("d%d:%d", me, peer)
}

func GetOrCreateDirectChatForPeers(me int, peer int) (*x.Chat, error) {
	keyChach := fmt.Sprintf("chat_%d_%d", me, peer)
	ch, ok := RowCache.GetChatByChatId(keyChach)
	if ok {
		return ch, nil
	}

	chatMe, err := x.NewChat_Selector().UserId_Eq(me).PeerUserId_Eq(peer).GetRow(base.DB)
	if err != nil {
		chatMe = &x.Chat{
			//ChatId:         helper.NextRowsSeqId(),
			ChatKey:        UsersToChatKey(me, peer),
			RoomKey:        UsersToRoomKey(me, peer),
			RoomTypeEnumId: int(x.RoomTypeEnum_DIRECT),
			UserId:         me,
			LastSeqSeen:    0,
			LastSeqDelete:  0,
			PeerUserId:     peer,
			GroupId:        0,
			CreatedSe:    helper.TimeNow(),
			CurrentSeq:     0,
		}
		err = chatMe.Save(base.DB)
		if err != nil {
			return nil, err
		}
	}
	RowCache.SetChatByKey(keyChach, chatMe)
	return chatMe, nil
}

func Chat_IncermentForNewMessage(c *x.Chat) {
	c.CurrentSeq += 1 //todo dep - remove this
	c.UpdatedMs = helper.TimeNowMs()
}
/*

func Chat_GetChatByIdAndUserId(chatId, userId int) (*x.Chat, error) {
	return x.NewChat_Selector().ChatId_Eq(chatId).UserId_Eq(userId).GetRow(base.DB)
}
*/
