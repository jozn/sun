
package x



func PBCon_Activity_PB_To_Activity(* o Activity_PB) Activity {
    n := &Activity
    n.Id = n.Id
    n.ActorUserId = n.ActorUserId
    n.ActionTypeId = n.ActionTypeId
    n.RowId = n.RowId
    n.RootId = n.RootId
    n.RefId = n.RefId
    n.CreatedAt = n.CreatedAt

    return n
}

func PBCon_Bucket_PB_To_Bucket(* o Bucket_PB) Bucket {
    n := &Bucket
    n.BucketId = n.BucketId
    n.BucketName = n.BucketName
    n.Server1Id = n.Server1Id
    n.Server2Id = n.Server2Id
    n.BackupServerId = n.BackupServerId
    n.ContentObjectTypeId = n.ContentObjectTypeId
    n.ContentObjectCount = n.ContentObjectCount
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_Comment_PB_To_Comment(* o Comment_PB) Comment {
    n := &Comment
    n.Id = n.Id
    n.UserId = n.UserId
    n.PostId = n.PostId
    n.Text = n.Text
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_FollowingList_PB_To_FollowingList(* o FollowingList_PB) FollowingList {
    n := &FollowingList
    n.Id = n.Id
    n.UserId = n.UserId
    n.ListType = n.ListType
    n.Name = n.Name
    n.Count = n.Count
    n.IsAuto = n.IsAuto
    n.IsPimiry = n.IsPimiry
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_FollowingListMember_PB_To_FollowingListMember(* o FollowingListMember_PB) FollowingListMember {
    n := &FollowingListMember
    n.Id = n.Id
    n.ListId = n.ListId
    n.UserId = n.UserId
    n.FollowedUserId = n.FollowedUserId
    n.FollowType = n.FollowType
    n.UpdatedTimeMs = n.UpdatedTimeMs

    return n
}

func PBCon_FollowingListMemberHistory_PB_To_FollowingListMemberHistory(* o FollowingListMemberHistory_PB) FollowingListMemberHistory {
    n := &FollowingListMemberHistory
    n.Id = n.Id
    n.ListId = n.ListId
    n.UserId = n.UserId
    n.FollowedUserId = n.FollowedUserId
    n.FollowType = n.FollowType
    n.UpdatedTimeMs = n.UpdatedTimeMs
    n.FollowId = n.FollowId

    return n
}

func PBCon_Like_PB_To_Like(* o Like_PB) Like {
    n := &Like
    n.Id = n.Id
    n.PostId = n.PostId
    n.PostTypeId = n.PostTypeId
    n.UserId = n.UserId
    n.TypeId = n.TypeId
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_Media_PB_To_Media(* o Media_PB) Media {
    n := &Media
    n.Id = n.Id
    n.UserId = n.UserId
    n.PostId = n.PostId
    n.AlbumId = n.AlbumId
    n.TypeId = n.TypeId
    n.CreatedTime = n.CreatedTime
    n.Src = n.Src

    return n
}

func PBCon_Message_PB_To_Message(* o Message_PB) Message {
    n := &Message
    n.Id = n.Id
    n.ToUserId = n.ToUserId
    n.RoomKey = n.RoomKey
    n.MessageKey = n.MessageKey
    n.FromUserID = n.FromUserID
    n.Data = n.Data
    n.TimeMs = n.TimeMs

    return n
}

func PBCon_MsgDeletedFromServer_PB_To_MsgDeletedFromServer(* o MsgDeletedFromServer_PB) MsgDeletedFromServer {
    n := &MsgDeletedFromServer
    n.Id = n.Id
    n.ToUserId = n.ToUserId
    n.MsgKey = n.MsgKey
    n.PeerUserId = n.PeerUserId
    n.RoomKey = n.RoomKey
    n.AtTime = n.AtTime

    return n
}

func PBCon_MsgReceivedToPeer_PB_To_MsgReceivedToPeer(* o MsgReceivedToPeer_PB) MsgReceivedToPeer {
    n := &MsgReceivedToPeer
    n.Id = n.Id
    n.ToUserId = n.ToUserId
    n.MsgKey = n.MsgKey
    n.RoomKey = n.RoomKey
    n.PeerUserId = n.PeerUserId
    n.AtTime = n.AtTime

    return n
}

func PBCon_MsgSeenByPeer_PB_To_MsgSeenByPeer(* o MsgSeenByPeer_PB) MsgSeenByPeer {
    n := &MsgSeenByPeer
    n.Id = n.Id
    n.ToUserId = n.ToUserId
    n.MsgKey = n.MsgKey
    n.RoomKey = n.RoomKey
    n.PeerUserId = n.PeerUserId
    n.AtTime = n.AtTime

    return n
}

func PBCon_Notification_PB_To_Notification(* o Notification_PB) Notification {
    n := &Notification
    n.Id = n.Id
    n.ForUserId = n.ForUserId
    n.ActorUserId = n.ActorUserId
    n.ActionTypeId = n.ActionTypeId
    n.ObjectTypeId = n.ObjectTypeId
    n.RowId = n.RowId
    n.RootId = n.RootId
    n.RefId = n.RefId
    n.SeenStatus = n.SeenStatus
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_NotificationRemoved_PB_To_NotificationRemoved(* o NotificationRemoved_PB) NotificationRemoved {
    n := &NotificationRemoved
    n.NotificationId = n.NotificationId
    n.ForUserId = n.ForUserId

    return n
}

func PBCon_PhoneContact_PB_To_PhoneContact(* o PhoneContact_PB) PhoneContact {
    n := &PhoneContact
    n.Id = n.Id
    n.PhoneDisplayName = n.PhoneDisplayName
    n.PhoneFamilyName = n.PhoneFamilyName
    n.PhoneNumber = n.PhoneNumber
    n.PhoneNormalizedNumber = n.PhoneNormalizedNumber
    n.PhoneContactRowId = n.PhoneContactRowId
    n.UserId = n.UserId
    n.DeviceUuidId = n.DeviceUuidId
    n.CreatedTime = n.CreatedTime
    n.UpdatedTime = n.UpdatedTime

    return n
}

func PBCon_Photo_PB_To_Photo(* o Photo_PB) Photo {
    n := &Photo
    n.PhotoId = n.PhotoId
    n.UserId = n.UserId
    n.PostId = n.PostId
    n.AlbumId = n.AlbumId
    n.ImageTypeId = n.ImageTypeId
    n.Title = n.Title
    n.Src = n.Src
    n.PathSrc = n.PathSrc
    n.BucketId = n.BucketId
    n.Width = n.Width
    n.Height = n.Height
    n.Ratio = n.Ratio
    n.HashMd5 = n.HashMd5
    n.CreatedTime = n.CreatedTime
    n.W1080 = n.W1080
    n.W720 = n.W720
    n.W480 = n.W480
    n.W320 = n.W320
    n.W160 = n.W160
    n.W80 = n.W80

    return n
}

func PBCon_Post_PB_To_Post(* o Post_PB) Post {
    n := &Post
    n.Id = n.Id
    n.UserId = n.UserId
    n.TypeId = n.TypeId
    n.Text = n.Text
    n.FormatedText = n.FormatedText
    n.MediaCount = n.MediaCount
    n.SharedTo = n.SharedTo
    n.DisableComment = n.DisableComment
    n.HasTag = n.HasTag
    n.LikesCount = n.LikesCount
    n.CommentsCount = n.CommentsCount
    n.EditedTime = n.EditedTime
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_RecommendUser_PB_To_RecommendUser(* o RecommendUser_PB) RecommendUser {
    n := &RecommendUser
    n.Id = n.Id
    n.UserId = n.UserId
    n.TargetId = n.TargetId
    n.Weight = n.Weight
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_SearchClicked_PB_To_SearchClicked(* o SearchClicked_PB) SearchClicked {
    n := &SearchClicked
    n.Id = n.Id
    n.Query = n.Query
    n.ClickType = n.ClickType
    n.TargetId = n.TargetId
    n.UserId = n.UserId
    n.CreatedAt = n.CreatedAt

    return n
}

func PBCon_Session_PB_To_Session(* o Session_PB) Session {
    n := &Session
    n.Id = n.Id
    n.UserId = n.UserId
    n.SessionUuid = n.SessionUuid
    n.ClientUuid = n.ClientUuid
    n.DeviceUuid = n.DeviceUuid
    n.LastActivityTime = n.LastActivityTime
    n.LastIpAddress = n.LastIpAddress
    n.LastWifiMacAddress = n.LastWifiMacAddress
    n.LastNetworkType = n.LastNetworkType
    n.LastNetworkTypeId = n.LastNetworkTypeId
    n.AppVersion = n.AppVersion
    n.UpdatedTime = n.UpdatedTime
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_Tag_PB_To_Tag(* o Tag_PB) Tag {
    n := &Tag
    n.Id = n.Id
    n.Name = n.Name
    n.Count = n.Count
    n.IsBlocked = n.IsBlocked
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_TagsPost_PB_To_TagsPost(* o TagsPost_PB) TagsPost {
    n := &TagsPost
    n.Id = n.Id
    n.TagId = n.TagId
    n.PostId = n.PostId
    n.TypeId = n.TypeId
    n.CreatedTime = n.CreatedTime

    return n
}

func PBCon_User_PB_To_User(* o User_PB) User {
    n := &User
    n.Id = n.Id
    n.UserName = n.UserName
    n.UserNameLower = n.UserNameLower
    n.FirstName = n.FirstName
    n.LastName = n.LastName
    n.About = n.About
    n.FullName = n.FullName
    n.AvatarUrl = n.AvatarUrl
    n.PrivacyProfile = n.PrivacyProfile
    n.Phone = n.Phone
    n.Email = n.Email
    n.IsDeleted = n.IsDeleted
    n.PasswordHash = n.PasswordHash
    n.PasswordSalt = n.PasswordSalt
    n.FollowersCount = n.FollowersCount
    n.FollowingCount = n.FollowingCount
    n.PostsCount = n.PostsCount
    n.MediaCount = n.MediaCount
    n.LikesCount = n.LikesCount
    n.ResharedCount = n.ResharedCount
    n.LastActionTime = n.LastActionTime
    n.LastPostTime = n.LastPostTime
    n.PrimaryFollowingList = n.PrimaryFollowingList
    n.CreatedTime = n.CreatedTime
    n.UpdatedTime = n.UpdatedTime
    n.SessionUuid = n.SessionUuid
    n.DeviceUuid = n.DeviceUuid
    n.LastWifiMacAddress = n.LastWifiMacAddress
    n.LastNetworkType = n.LastNetworkType
    n.AppVersion = n.AppVersion
    n.LastActivityTime = n.LastActivityTime
    n.LastLoginTime = n.LastLoginTime
    n.LastIpAddress = n.LastIpAddress

    return n
}

func PBCon_UserMetaInfo_PB_To_UserMetaInfo(* o UserMetaInfo_PB) UserMetaInfo {
    n := &UserMetaInfo
    n.Id = n.Id
    n.UserId = n.UserId
    n.IsNotificationDirty = n.IsNotificationDirty
    n.LastUserRecGen = n.LastUserRecGen

    return n
}

func PBCon_UserPassword_PB_To_UserPassword(* o UserPassword_PB) UserPassword {
    n := &UserPassword
    n.UserId = n.UserId
    n.Password = n.Password
    n.CreatedTime = n.CreatedTime

    return n
}

