package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

const (
	LogType_FOLLOW       int = 1
	LogType_UNFOLLOW     int = 2
	LogType_YOU_BOCKED   int = 3
	LogType_YOU_UNBOCKED int = 4
)

func GeneralLog_Follow(me, peer int) {
	l := &x.GeneralLog{
		Id:        helper.NextRowsSeqId(),
		ToUserId:  me,
		TargetId:  peer,
		LogTypeId: LogType_FOLLOW,
		ExtraPB:   []byte{},
		ExtraJson: "",
		CreatedMs: helper.TimeNowMs(),
	}
	l.Save(base.DB)
}

func GeneralLog_UnFollow(me, peer int) {
	l := &x.GeneralLog{
		Id:        helper.NextRowsSeqId(),
		ToUserId:  me,
		TargetId:  peer,
		LogTypeId: LogType_UNFOLLOW,
		ExtraPB:   []byte{},
		ExtraJson: "",
		CreatedMs: helper.TimeNowMs(),
	}
	l.Save(base.DB)
}

//todo do we need aggertion like: just the last one is important: followed or unfollowed
func GeneralLog_GetLastView(me, lastId int) (*x.PB_SyncResponse_GetGeneralUpdates, error) {
	rows, err := x.NewGeneralLog_Selector().Id_GT(lastId).ToUserId_Eq(me).GetRows(base.DB)
	if err != nil {
		return nil, err
	}

	v := &x.PB_SyncResponse_GetGeneralUpdates{}

	for _, r := range rows {
		switch r.LogTypeId {
		case LogType_FOLLOW:
			o := &x.PB_UpdateUserBlocked{
				UserId:  int32(r.TargetId),
				Blocked: true,
			}
			v.UserBlockedByMe = append(v.UserBlockedByMe, o)

		case LogType_UNFOLLOW:
			o := &x.PB_UpdateUserBlocked{
				UserId:  int32(r.TargetId),
				Blocked: false,
			}
			v.UserBlockedByMe = append(v.UserBlockedByMe, o)
		}
	}

	return v, nil
}
