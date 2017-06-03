package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
)

func flusher_flushPushMsgs(uid int) {
	uids, err := x.NewMsgPush_Selector().
		Select_Uid().
		ToUser_Eq(uid).
		OrderBy_Id_Desc().
		GetIntSlice(base.DB)
	if err != nil {
		return
	}

	msgs := Message_GetMsgsByUids(uids)
	MessageModel_PushToPipeMsgsToUser(uid, msgs)
}

func flusher_flushPushEvents(uid int) {
	eventsRows, err := x.NewMsgPushEvent_Selector().
		Select_Uid().
		ToUserId_Eq(uid).
		OrderBy_Id_Desc().
		GetRows2(base.DB)
	if err != nil {
		return
	}

	MessageModel_PushToPipeMsgEventsToUser(uid, eventsRows)
}

func flusher_flushPushNotifications(uid int) {

}
