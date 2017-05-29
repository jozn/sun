package models

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/tools/go/gcimporter15/testdata"
	"ms/sun/helper"
	"ms/sun/models/x"
)

//just return the by value

func PBConv_PB_Message_toNew_Message(pb *x.PB_Message) x.Message {
	msg := &x.Message{
		Id:            0,
		UserId:        pb.UserId,
		MessageKey:    pb.MessageKey,
		RoomKey:       pb.RoomKey,
		MessageType:   pb.MessageTypeId,
		RoomType:      pb.RoomTypeId,
		DataPB:        nil,
		DataJson:      helper.ToJson(pb),
		CreatedTimeMs: helper.TimeNowMs(),
	}
	bs, err := proto.Marshal(pb)
	if err == nil {
		msg.DataPB = bs
	}

	return msg
}

/*func PBConv_Message_toNew_PB_Message(pb *x.Message) x.PB_Message {
    msg := &x.PB_Message{
        Id:            0,
        UserId:        pb.UserId,
        MessageKey:    pb.MessageKey,
        RoomKey:       pb.RoomKey,
        MessageType:   pb.MessageTypeId,
        RoomType:      pb.RoomTypeId,
        DataPB:        nil,
        DataJson:      helper.ToJson(pb),
        CreatedTimeMs: helper.TimeNowMs(),
    }
    bs, err := proto.Marshal(pb)
    if err == nil {
        msg.DataPB = bs
    }

    return msg
}*/

func PBConv_PB_MsgSeen_toNew_MsgPushEvent(pb *x.PB_MsgSeen) x.MsgPushEvent {
	msg := &x.MsgPushEvent{
		Id:         0,
		ToUserId:   -1, //must set
		MsgKey:     pb.MessageKey,
		RoomKey:    pb.RoomKey,
		PeerUserId: -1, //must set
		EventType:  MESSAGE_PUSH_EVENT_SEEN_BY_PEER,
		AtTime:     pb.AtTime,
	}

	return msg
}

func PBConv_User_toNew_PB_UserWithMe(p *x.User, meId int) x.PB_UserWithMe {
	u := &x.PB_UserWithMe{
		UserId:         p.Id,
		UserName:       p.UserName, //must set
		FirstName:      p.FirstName,
		LastName:       p.LastName,
		About:          p.About, //must set
		FullName:       p.FirstName + " " + p.LastName,
		AvatarUrl:      p.AvatarUrl,
		PrivacyProfile: p.PrivacyProfile,
		Phone:          p.Phone,
		UpdatedTime:    p.UpdatedTime,
		AppVersion:     p.AppVersion,
		FollowingType:  MemoryStore.UserFollowingList_GetFollowingTypeForUsers(meId, p.Id),
	}

	return u
}

func PBNew_PB_UserWithMe(UserId, meId int) x.PB_UserWithMe {

	p, ok := MemoryStore_User.GetUser(UserId)
	if ok {
		return PBConv_User_toNew_PB_UserWithMe(p, meId)
	}

	msg := &x.PB_UserWithMe{
		UserId: UserId,
	}

	return msg
}
