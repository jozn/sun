package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"strings"
)

func KeyNewMessageKey(UserId int) string {
	return fmt.Sprintf("%d_%d_%s", UserId, helper.TimeNow(), helper.RandString(4)) //todo extrac this to client
}

//"d25_56" "d56:24" ////not yet "g6_15646548_cdc" "g6:6_15646548_cdc"
func RoomKeyToOtherUser(roomKey string, userId int) int {
	key := strings.Replace(roomKey, "d", "", -1)
	key = strings.Replace(key, "g", "", -1)
	key = strings.Replace(key, ":", "_", -1)

	parts := strings.Split(key, "_")
	if len(parts) != 2 {
		return 0
	}

	i1 := helper.StrToInt(parts[0], 0)
	i2 := helper.StrToInt(parts[1], 0)

	if i1 == userId {
		return i2
	}

	if i2 == userId {
		return i1
	}

	return 0
}

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
			ChatKey:              UsersToChatKey(me, peer),
			RoomKey:              UsersToRoomKey(me, peer),
			RoomTypeEnumId:       int(x.RoomTypeEnum_DIRECT),
			UserId:               me,
			PeerUserId:           peer,
			GroupId:              0,
			CreatedSe:            helper.TimeNow(),
			StartMessageIdFrom:   helper.NextRowsSeqIdBase(),
			LastDeletedMessageId: 0,
			LastSeenMessageId:    0,
			LastMessageId:        0,
			UpdatedMs:            helper.TimeNowMs(),
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
	//c.CurrentSeq += 1 //todo dep - remove this
	c.UpdatedMs = helper.TimeNowMs()
}

func Chat_setLastMsg(c *x.Chat, msg *x.DirectMessage) {
	if msg.MessageId > 0 { //must alwasye be true
		c.LastMessageId = msg.MessageId
		c.UpdatedMs = helper.TimeNowMs()
	}
}

/*

func Chat_GetChatByIdAndUserId(chatId, userId int) (*x.Chat, error) {
	return x.NewChat_Selector().ChatId_Eq(chatId).UserId_Eq(userId).GetRow(base.DB)
}
*/
