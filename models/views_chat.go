package models

import "ms/sun/models/x"

type _chatViewImple int

var ChatViews = _chatViewImple(1)

func (_chatViewImple) MessageIdToMessageView(id int) *x.PB_MessageView {

	return nil
}

func (_chatViewImple) MessageIdsToMessageViews(ids []int) []*x.PB_MessageView {

	x.Store.PreLoadDirectMessageByMessageIds(ids)

	return nil
}
