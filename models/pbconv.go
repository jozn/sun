package models

import (
	"github.com/golang/protobuf/proto"
	"ms/Console2/dev5/m"
	"ms/sun/helper"
	"ms/sun/models/x"
)

//just return the by value

func PBConv_PB_Message_toNew_Message(pb *x.PB_Message) x.Message {
	bytes, _ := proto.Marshal(pb)
	json := helper.ToJson(pb)
	msg := x.Message{
		Id:            0,
		Uid:           helper.RandomSeqUid(),
		UserId:        int(pb.UserId),
		MessageKey:    pb.MessageKey,
		RoomKey:       pb.RoomKey,
		MessageType:   int(pb.MessageTypeId),
		RoomType:      int(pb.RoomTypeId),
		DataPB:        []byte(""),                //bytes,
		Data64:        helper.ToBase64Bin(bytes), //bytes,
		DataJson:      json,
		CreatedTimeMs: int(helper.TimeNowMs()),
	}
	/*bs, err := proto.Marshal(pb)
	if err == nil {
		msg.DataPB = bs
	}*/

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
	msg := x.MsgPushEvent{
		Id:         0,
		Uid:        helper.RandomSeqUid(),
		ToUserId:   -1, //must set
		MsgKey:     pb.MessageKey,
		RoomKey:    pb.RoomKey,
		PeerUserId: -1, //must set
		EventType:  MESSAGE_PUSH_EVENT_SEEN_BY_PEER,
		AtTime:     int(pb.AtTime),
	}

	return msg
}

func PBConv_User_toNew_PB_UserWithMe(p *x.User, meId int) x.PB_UserWithMe {
	u := x.PB_UserWithMe{
		//UserId:         int32(p.Id),
		UserId:         int32(p.Id),
		UserName:       p.UserName, //must set
		FirstName:      p.FirstName,
		LastName:       p.LastName,
		About:          p.About, //must set
		FullName:       p.FirstName + " " + p.LastName,
		AvatarUrl:      p.AvatarUrl,
		PrivacyProfile: int32(p.PrivacyProfile),
		Phone:          p.Phone,
		UpdatedTime:    int32(p.UpdatedTime),
		AppVersion:     int32(p.AppVersion),
		FollowingType:  int32(MemoryStore.UserFollowingList_GetFollowingTypeForUsers(meId, p.Id)),
	}

	return u
}

func PBNew_PB_UserWithMe(UserId, meId int) *x.PB_UserWithMe {

	p, ok := MemoryStore_User.GetUser(UserId)
	if ok {
		nn := PBConv_User_toNew_PB_UserWithMe(&p, meId)
		return &nn
	}

	msg := &x.PB_UserWithMe{
		UserId: int32(UserId),
	}

	return msg
}

func PBConv_MsgPushEvent_toNew_PB_MsgEvent(m *x.MsgPushEvent) x.PB_MsgEvent {
	pbEv := x.PB_MsgEvent{
		MessageKey: m.MsgKey,
		RoomKey:    m.RoomKey,
		PeerUserId: int64(m.PeerUserId),
		EventType:  int32(m.EventType),
		AtTime:     int64(m.AtTime),
	}

	return pbEv
}

func PBConv_PB_MsgFile_toNew_MsgFile(f *x.PB_MsgFile) x.MsgFile {
	row := x.MsgFile{
		Id:          0,
		Name:        f.Name,
		Size:        int(f.Size),
		FileType:    int(f.FileType),
		MimeType:    f.MimeType,
		Width:       int(f.Width),
		Height:      int(f.Height),
		Duration:    int(f.Duration),
		Extension:   f.Extension,
		ThumbData:   []byte{},
		ThumbData64: helper.ToBase64Bin(f.ThumbData),
		ServerSrc:   "", //must set with hand
		ServerPath:  "", //must set with hand
		ServerId:    0,  //must set with hand
	}

	return row
}

func PBConv_MsgFile_toNew_PB_MsgFile(f *x.MsgFile) x.PB_MsgFile {
	data, _ := helper.FromBase64ToBin(f.ThumbData64)
	pb := x.PB_MsgFile{
		Name:      f.Name,
		Size:      int64(f.Size),
		FileType:  int32(f.FileType),
		MimeType:  f.MimeType,
		Width:     int32(f.Width),
		Height:    int32(f.Height),
		Duration:  int32(f.Duration),
		Extension: f.Extension,
		ThumbData: data,
		ServerSrc: f.ServerSrc, //must set with hand
	}

	return pb
}
