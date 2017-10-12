package models

import (
	"ms/sun/helper"
	"ms/sun/models/x"
    "fmt"
)

func ChatPush_PushChnageMessageId(newMsg *x.DirectMessage, clientMsgId int, chatKey string) {
    fmt.Println("ChatPush_PushChnageMessageId ",newMsg,clientMsgId,chatKey)
	row := x.DirectUpdate{
		DirectUpdateId: helper.NextRowsSeqId(),
		ToUserId:       newMsg.UserId,
		MessageId:      newMsg.MessageId,
		MessageFileId:  0,
		OtherId:        clientMsgId,
		ChatKey:        chatKey,
		PeerUserId:     0,
		//EventType: 0,
		RoomLogTypeId: int(Push_MESSAGE_ID_CHANGE),
		FromSeq:       0,
		ToSeq:         0,
		ExtraJson:     "",
        ExtraPB:[]byte{},
        AtTimeMs:      helper.TimeNowMs(),
	}
	//row.Save(base.DB)

	LiveUpdateFramer.HereDirectDelayer <- UpdateDelayer{directUpdate: row}
}

func ChatPush_PushChnageMessageFileId(newMsg *x.MessageFile, UserId, clientMsgFileId int, chatKey string) {
    fmt.Println("ChatPush_Push  ChnageMessageId ",newMsg,clientMsgFileId,chatKey)
	row := x.DirectUpdate{
		DirectUpdateId: helper.NextRowsSeqId(),
		ToUserId:       UserId,
		MessageId:      0,
		MessageFileId:  newMsg.MessageFileId,
		OtherId:        clientMsgFileId,
		ChatKey:        chatKey,
		PeerUserId:     0,
		//EventType: 0,
		RoomLogTypeId: int(Push_MESSAGE_FILE_ID_CHANGE),
		FromSeq:       0,
		ToSeq:         0,
		ExtraJson:     "",
		ExtraPB:[]byte{},
		AtTimeMs:      helper.TimeNowMs(),
	}
	//row.Save(base.DB)

	LiveUpdateFramer.HereDirectDelayer <- UpdateDelayer{directUpdate: row}
}
