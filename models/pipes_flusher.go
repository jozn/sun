package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
)

func flusher_flushPushMsgs(uid int) {
	logPipes.Println("flusher_flushPushMsgs() uid: ", uid)
	uids, err := x.NewMsgPush_Selector().
		Select_MsgUid().
		ToUser_Eq(uid).
		OrderBy_Id_Desc().
		GetIntSlice(base.DB)
	if err != nil {
		return
	}

	msgs := Message_GetMsgsByUids(uids)

	logPipes.Println("flusher_flushPushMsgs() len: ", len(msgs), " ", len(uids))

	MessageModel_PushToPipeMsgsToUser(uid, msgs)
}

func flusher_flushPushEvents(uid int) {
	logPipes.Println("flusher_flushPushEvents() uid: ", uid)
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
	logPipes.Println("flusher_flushPushNotifications() uid: ", uid)

}
