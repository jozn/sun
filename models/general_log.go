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
