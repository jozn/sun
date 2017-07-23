package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

func UsersToRoomKey32(me, peer int32) string {
	return fmt.Sprintf("%d_%d", me, peer)
}

func UsersToRoomKey(me int, peer int) string {
	return fmt.Sprintf("%d_%d", me, peer)
}

func GetOrCreateDirectChatForPeers(me int, peer int) (*x.Chat, error) {
	chatMe, err := x.NewChat_Selector().UserId_Eq(me).PeerUserId_Eq(peer).GetRow(base.DB)
	if err != nil {
		chatMe = &x.Chat{
			ChatId:         0,
			ChatKey:        UsersToRoomKey(me, peer),
			RoomTypeEnumId: int(x.RoomTypeEnum_DIRECT),
			UserId:         me,
			LastSeqSeen:    0,
			LastSeqDelete:  0,
			PeerUserId:     peer,
			GroupId:        0,
			CreatedTime:    helper.TimeNow(),
			CurrentSeq:     0,
		}
		err = chatMe.Save(base.DB)
		if err != nil {
			return nil, err
		}
	}
	return chatMe, nil
}

func Chat_IncermentForNewMessage(c *x.Chat) {
	c.CurrentSeq += 1
	c.UpdatedMs = helper.TimeNowMs()
}
