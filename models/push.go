package models

import "ms/sun/models/x"

func Push_PushMsgsAddMsg(ToUserId int, msgRow x.Message, msgPb x.PB_Message) {
	AllPipesMap.SendToUserWithCallBack()
}

func Push_PushMsgsEvents(events []x.MsgPushEvent) {
	mpGroupByuser := make(map[int][]x.MsgSeenByPeer)
	for _, seen := range events {
		mpGroupByuser[seen.ToUserId] = append(mpGroupByuser[seen.ToUserId], seen)
	}

}
