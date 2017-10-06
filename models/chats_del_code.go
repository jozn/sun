package models

import "ms/sun/models/x"

//dep
func PushView_directLogsTo_PB_ChangesHolderView(meId int, logs []*x.DirectUpdate) *x.PB_PushHolderView {

    //preload in here
    usersToLoad := make(map[int]bool)
    chatIdsToLoad := make(map[int]bool)
    msgIdsToLoad := []int{}
    msgFileIdsToLoad := []int{}
    for _, log := range logs { //each user
        if log.RoomLogTypeId == int(Push_NEW_DIRECT_MESSAGE) {
            usersToLoad[log.PeerUserId] = true
            //chatIdsToLoad[log.ChatId] = true
            msgIdsToLoad = append(msgIdsToLoad, log.MessageId)
            if log.MessageFileId > 0 {
                msgFileIdsToLoad = append(msgFileIdsToLoad, log.MessageFileId)
            }
        }
    }

    //
    res := &x.PB_PushHolderView{}

    for _, logRow := range logs { //each user
        switch Push(logRow.RoomLogTypeId) {
        case Push_NEW_DIRECT_MESSAGE:
            if v, ok := _pushView_newDirectMessage(logRow); ok {
                res.NewMessages = append(res.NewMessages, v)
            }

            //metas
        case Push_MESSAGE_RECIVED_TO_SERVER:
            if v, ok := _pushView_messageMeta(logRow); ok {
                res.MessagesDelivierdToServer = append(res.MessagesDelivierdToServer, v)
            }

        case Push_MESSAGE_DELIVIERD_TO_PEER:
            if v, ok := _pushView_messageMeta(logRow); ok {
                res.MessagesDelivierdToPeer = append(res.MessagesDelivierdToPeer, v)
            }

        case Push_MESSAGE_SEEN_BY_PEER:
            if v, ok := _pushView_messageMeta(logRow); ok {
                res.MessagesSeenByPeer = append(res.MessagesSeenByPeer, v)
            }

        case Push_MESSAGE_DELETED_FROM_SERVER:
            if v, ok := _pushView_messageMeta(logRow); ok {
                res.MessagesDeletedFromServer = append(res.MessagesDeletedFromServer, v)
            }

        }
    }

    if len(usersToLoad) > 0 {
        res.Users = ViewUser_GetUserViewList(meId, usersToLoad)
    }

    if len(chatIdsToLoad) > 0 {
        //res.Chats = ViewChat_GetChatViewList_map(meId, chatIdsToLoad)
    }

    //TODO: add messages files id

    return res
}

