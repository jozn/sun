package models

import "ms/sun/models/x"

func PushProxy_PushSendMessageToUser(ToUserId int, msg x.PB_Message) {

}

func PushProxy_PushMsgsAddMsg(ToUserId int, msgRow x.Message, msgPb x.PB_Message) {

}

func PushProxy_PushMsgsEvents(events []x.MsgPushEvent) {
	/*mpGroupByuser := make(map[int][]x.MsgSeenByPeer)
	for _, seen := range events {
		mpGroupByuser[seen.ToUserId] = append(mpGroupByuser[seen.ToUserId],nil)// seen)
	}*/

}
