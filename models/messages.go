package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
)

func Message_GetMsgsByUids(uids []int) ([]*x.Message) {
	res, err := x.NewMessage_Selector().Uid_In(uids).GetRows(base.DB)
	if err != nil {
		logPipes.Println("Message_GetMsgsByUids error: ", err)
        return []*x.Message{}
	}

	return res
}
