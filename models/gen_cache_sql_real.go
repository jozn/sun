
package models

import (
	"ms/sun/base"
	"strconv"
)

//Me Modified names

func (c _StoreImpl) GetCommentById(Id int) (*Comment, bool) {
    obj, err := Cacher.Get("Comment:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = CommentById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetFollowingListByUserId(UserId int) (*FollowingList, bool) {
    obj, err := Cacher.Get("FollowingList:" + strconv.Itoa(UserId))
    if err == nil {
        return obj, true
    }
    obj, err = FollowingListByUserId(base.DB, UserId)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetFollowingListMemberById(Id int) (*FollowingListMember, bool) {
    obj, err := Cacher.Get("FollowingListMember:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = FollowingListMemberById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetFollowingListMemberHistoryById(Id int) (*FollowingListMemberHistory, bool) {
    obj, err := Cacher.Get("FollowingListMemberHistory:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = FollowingListMemberHistoryById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetLikeById(Id int) (*Like, bool) {
    obj, err := Cacher.Get("Like:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = LikeById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetMediaById(Id int) (*Media, bool) {
    obj, err := Cacher.Get("Media:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = MediaById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetMessageById(Id int) (*Message, bool) {
    obj, err := Cacher.Get("Message:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = MessageById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetMsgDeletedFromServerById(Id int) (*MsgDeletedFromServer, bool) {
    obj, err := Cacher.Get("MsgDeletedFromServer:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = MsgDeletedFromServerById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetMsgReceivedToPeerById(Id int) (*MsgReceivedToPeer, bool) {
    obj, err := Cacher.Get("MsgReceivedToPeer:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = MsgReceivedToPeerById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetMsgSeenByPeerById(Id int) (*MsgSeenByPeer, bool) {
    obj, err := Cacher.Get("MsgSeenByPeer:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = MsgSeenByPeerById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetNotificationById(Id int) (*Notification, bool) {
    obj, err := Cacher.Get("Notification:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = NotificationById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetNotificationRemovedByNotificationId(NotificationId int) (*NotificationRemoved, bool) {
    obj, err := Cacher.Get("NotificationRemoved:" + strconv.Itoa(NotificationId))
    if err == nil {
        return obj, true
    }
    obj, err = NotificationRemovedByNotificationId(base.DB, NotificationId)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetPhoneContactById(Id int) (*PhoneContact, bool) {
    obj, err := Cacher.Get("PhoneContact:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = PhoneContactById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetPostById(Id int) (*Post, bool) {
    obj, err := Cacher.Get("Post:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = PostById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetRecommendUserById(Id int) (*RecommendUser, bool) {
    obj, err := Cacher.Get("RecommendUser:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = RecommendUserById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetSearchClickedById(Id int) (*SearchClicked, bool) {
    obj, err := Cacher.Get("SearchClicked:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = SearchClickedById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetSessionBySessionUuid(SessionUuid int) (*Session, bool) {
    obj, err := Cacher.Get("Session:" + strconv.Itoa(SessionUuid))
    if err == nil {
        return obj, true
    }
    obj, err = SessionBySessionUuid(base.DB, SessionUuid)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetTagById(Id int) (*Tag, bool) {
    obj, err := Cacher.Get("Tag:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = TagById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetTagsPostById(Id int) (*TagsPost, bool) {
    obj, err := Cacher.Get("TagsPost:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = TagsPostById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetUserById(Id int) (*User, bool) {
    obj, err := Cacher.Get("User:" + strconv.Itoa(Id))
    if err == nil {
        return obj, true
    }
    obj, err = UserById(base.DB, Id)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetUserMetaInfoByUserId(UserId int) (*UserMetaInfo, bool) {
    obj, err := Cacher.Get("UserMetaInfo:" + strconv.Itoa(UserId))
    if err == nil {
        return obj, true
    }
    obj, err = UserMetaInfoByUserId(base.DB, UserId)
    if err == nil {
        return obj, true
    }
    return nil, false
}

func (c _StoreImpl) GetUserPasswordByUserId(UserId int) (*UserPassword, bool) {
    obj, err := Cacher.Get("UserPassword:" + strconv.Itoa(UserId))
    if err == nil {
        return obj, true
    }
    obj, err = UserPasswordByUserId(base.DB, UserId)
    if err == nil {
        return obj, true
    }
    return nil, false
}
