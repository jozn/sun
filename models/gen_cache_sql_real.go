
package models

import (
    "strconv"
    "ms/sun/base"
)

//Me Modified names

func (c _StoreImpl) GetActivityById(Id int) (*Activity, bool) {
    o, ok := RowCache.Get("Activity:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Activity); ok {
            return obj, true
        }
    }
    obj2, err := ActivityById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetCommentById(Id int) (*Comment, bool) {
    o, ok := RowCache.Get("Comment:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Comment); ok {
            return obj, true
        }
    }
    obj2, err := CommentById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetFollowingListByUserId(UserId int) (*FollowingList, bool) {
    o, ok := RowCache.Get("FollowingList:" + strconv.Itoa(UserId))
    if ok {
        if obj, ok := o.(*FollowingList); ok {
            return obj, true
        }
    }
    obj2, err := FollowingListByUserId(base.DB, UserId)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetFollowingListMemberById(Id int) (*FollowingListMember, bool) {
    o, ok := RowCache.Get("FollowingListMember:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*FollowingListMember); ok {
            return obj, true
        }
    }
    obj2, err := FollowingListMemberById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetFollowingListMemberHistoryById(Id int) (*FollowingListMemberHistory, bool) {
    o, ok := RowCache.Get("FollowingListMemberHistory:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*FollowingListMemberHistory); ok {
            return obj, true
        }
    }
    obj2, err := FollowingListMemberHistoryById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetLikeById(Id int) (*Like, bool) {
    o, ok := RowCache.Get("Like:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Like); ok {
            return obj, true
        }
    }
    obj2, err := LikeById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetMediaById(Id int) (*Media, bool) {
    o, ok := RowCache.Get("Media:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Media); ok {
            return obj, true
        }
    }
    obj2, err := MediaById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetMessageById(Id int) (*Message, bool) {
    o, ok := RowCache.Get("Message:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Message); ok {
            return obj, true
        }
    }
    obj2, err := MessageById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetMsgDeletedFromServerById(Id int) (*MsgDeletedFromServer, bool) {
    o, ok := RowCache.Get("MsgDeletedFromServer:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*MsgDeletedFromServer); ok {
            return obj, true
        }
    }
    obj2, err := MsgDeletedFromServerById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetMsgReceivedToPeerById(Id int) (*MsgReceivedToPeer, bool) {
    o, ok := RowCache.Get("MsgReceivedToPeer:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*MsgReceivedToPeer); ok {
            return obj, true
        }
    }
    obj2, err := MsgReceivedToPeerById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetMsgSeenByPeerById(Id int) (*MsgSeenByPeer, bool) {
    o, ok := RowCache.Get("MsgSeenByPeer:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*MsgSeenByPeer); ok {
            return obj, true
        }
    }
    obj2, err := MsgSeenByPeerById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetNotificationById(Id int) (*Notification, bool) {
    o, ok := RowCache.Get("Notification:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Notification); ok {
            return obj, true
        }
    }
    obj2, err := NotificationById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetNotificationRemovedByNotificationId(NotificationId int) (*NotificationRemoved, bool) {
    o, ok := RowCache.Get("NotificationRemoved:" + strconv.Itoa(NotificationId))
    if ok {
        if obj, ok := o.(*NotificationRemoved); ok {
            return obj, true
        }
    }
    obj2, err := NotificationRemovedByNotificationId(base.DB, NotificationId)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetPhoneContactById(Id int) (*PhoneContact, bool) {
    o, ok := RowCache.Get("PhoneContact:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*PhoneContact); ok {
            return obj, true
        }
    }
    obj2, err := PhoneContactById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetPostById(Id int) (*Post, bool) {
    o, ok := RowCache.Get("Post:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Post); ok {
            return obj, true
        }
    }
    obj2, err := PostById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetRecommendUserById(Id int) (*RecommendUser, bool) {
    o, ok := RowCache.Get("RecommendUser:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*RecommendUser); ok {
            return obj, true
        }
    }
    obj2, err := RecommendUserById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetSearchClickedById(Id int) (*SearchClicked, bool) {
    o, ok := RowCache.Get("SearchClicked:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*SearchClicked); ok {
            return obj, true
        }
    }
    obj2, err := SearchClickedById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetSessionById(Id int) (*Session, bool) {
    o, ok := RowCache.Get("Session:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Session); ok {
            return obj, true
        }
    }
    obj2, err := SessionById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetTagById(Id int) (*Tag, bool) {
    o, ok := RowCache.Get("Tag:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*Tag); ok {
            return obj, true
        }
    }
    obj2, err := TagById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetTagsPostById(Id int) (*TagsPost, bool) {
    o, ok := RowCache.Get("TagsPost:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*TagsPost); ok {
            return obj, true
        }
    }
    obj2, err := TagsPostById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetUserById(Id int) (*User, bool) {
    o, ok := RowCache.Get("User:" + strconv.Itoa(Id))
    if ok {
        if obj, ok := o.(*User); ok {
            return obj, true
        }
    }
    obj2, err := UserById(base.DB, Id)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetUserMetaInfoByUserId(UserId int) (*UserMetaInfo, bool) {
    o, ok := RowCache.Get("UserMetaInfo:" + strconv.Itoa(UserId))
    if ok {
        if obj, ok := o.(*UserMetaInfo); ok {
            return obj, true
        }
    }
    obj2, err := UserMetaInfoByUserId(base.DB, UserId)
    if err == nil {
        return obj2, true
    }
    return nil, false
}

func (c _StoreImpl) GetUserPasswordByUserId(UserId int) (*UserPassword, bool) {
    o, ok := RowCache.Get("UserPassword:" + strconv.Itoa(UserId))
    if ok {
        if obj, ok := o.(*UserPassword); ok {
            return obj, true
        }
    }
    obj2, err := UserPasswordByUserId(base.DB, UserId)
    if err == nil {
        return obj2, true
    }
    return nil, false
}
