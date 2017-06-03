package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
)

func Message_GetMsgsByUids(uids []int) (res []*x.Message) {
	res, _ = x.NewMessage_Selector().Uid_In(uids).GetRows(base.DB)
	return
}
