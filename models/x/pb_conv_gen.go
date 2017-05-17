
package x


/*
func PBConv_Activity_PB_To_Activity( o *Activity_PB) *Activity {
     n := &Activity{
      Id: int ( o.Id ),
      ActorUserId: int ( o.ActorUserId ),
      ActionTypeId: int ( o.ActionTypeId ),
      RowId: int ( o.RowId ),
      RootId: int ( o.RootId ),
      RefId: int ( o.RefId ),
      CreatedAt: int ( o.CreatedAt ),
    }
    return n
}

func PBConv_Activity_To_Activity_PB ( o *Activity) *Activity_PB {
     n := &Activity_PB{
      Id: int64 ( o.Id ),
      ActorUserId: int32 ( o.ActorUserId ),
      ActionTypeId: int32 ( o.ActionTypeId ),
      RowId: int32 ( o.RowId ),
      RootId: int32 ( o.RootId ),
      RefId: int64 ( o.RefId ),
      CreatedAt: int32 ( o.CreatedAt ),
    }
    return n
}
*/
/*
func PBConv_Bucket_PB_To_Bucket( o *Bucket_PB) *Bucket {
     n := &Bucket{
      BucketId: int ( o.BucketId ),
      BucketName: string ( o.BucketName ),
      Server1Id: int ( o.Server1Id ),
      Server2Id: int ( o.Server2Id ),
      BackupServerId: int ( o.BackupServerId ),
      ContentObjectTypeId: int ( o.ContentObjectTypeId ),
      ContentObjectCount: int ( o.ContentObjectCount ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_Bucket_To_Bucket_PB ( o *Bucket) *Bucket_PB {
     n := &Bucket_PB{
      BucketId: int32 ( o.BucketId ),
      BucketName: string ( o.BucketName ),
      Server1Id: int32 ( o.Server1Id ),
      Server2Id: int32 ( o.Server2Id ),
      BackupServerId: int32 ( o.BackupServerId ),
      ContentObjectTypeId: int32 ( o.ContentObjectTypeId ),
      ContentObjectCount: int32 ( o.ContentObjectCount ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_Comment_PB_To_Comment( o *Comment_PB) *Comment {
     n := &Comment{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      PostId: int ( o.PostId ),
      Text: string ( o.Text ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_Comment_To_Comment_PB ( o *Comment) *Comment_PB {
     n := &Comment_PB{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      PostId: int32 ( o.PostId ),
      Text: string ( o.Text ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_FollowingList_PB_To_FollowingList( o *FollowingList_PB) *FollowingList {
     n := &FollowingList{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      ListType: int ( o.ListType ),
      Name: string ( o.Name ),
      Count: int ( o.Count ),
      IsAuto: int ( o.IsAuto ),
      IsPimiry: int ( o.IsPimiry ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_FollowingList_To_FollowingList_PB ( o *FollowingList) *FollowingList_PB {
     n := &FollowingList_PB{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      ListType: int32 ( o.ListType ),
      Name: string ( o.Name ),
      Count: int32 ( o.Count ),
      IsAuto: int32 ( o.IsAuto ),
      IsPimiry: int32 ( o.IsPimiry ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_FollowingListMember_PB_To_FollowingListMember( o *FollowingListMember_PB) *FollowingListMember {
     n := &FollowingListMember{
      Id: int ( o.Id ),
      ListId: int ( o.ListId ),
      UserId: int ( o.UserId ),
      FollowedUserId: int ( o.FollowedUserId ),
      FollowType: int ( o.FollowType ),
      UpdatedTimeMs: int ( o.UpdatedTimeMs ),
    }
    return n
}

func PBConv_FollowingListMember_To_FollowingListMember_PB ( o *FollowingListMember) *FollowingListMember_PB {
     n := &FollowingListMember_PB{
      Id: int64 ( o.Id ),
      ListId: int32 ( o.ListId ),
      UserId: int32 ( o.UserId ),
      FollowedUserId: int32 ( o.FollowedUserId ),
      FollowType: int32 ( o.FollowType ),
      UpdatedTimeMs: int64 ( o.UpdatedTimeMs ),
    }
    return n
}
*/
/*
func PBConv_FollowingListMemberHistory_PB_To_FollowingListMemberHistory( o *FollowingListMemberHistory_PB) *FollowingListMemberHistory {
     n := &FollowingListMemberHistory{
      Id: int ( o.Id ),
      ListId: int ( o.ListId ),
      UserId: int ( o.UserId ),
      FollowedUserId: int ( o.FollowedUserId ),
      FollowType: int ( o.FollowType ),
      UpdatedTimeMs: int ( o.UpdatedTimeMs ),
      FollowId: int ( o.FollowId ),
    }
    return n
}

func PBConv_FollowingListMemberHistory_To_FollowingListMemberHistory_PB ( o *FollowingListMemberHistory) *FollowingListMemberHistory_PB {
     n := &FollowingListMemberHistory_PB{
      Id: int64 ( o.Id ),
      ListId: int32 ( o.ListId ),
      UserId: int32 ( o.UserId ),
      FollowedUserId: int32 ( o.FollowedUserId ),
      FollowType: int32 ( o.FollowType ),
      UpdatedTimeMs: int64 ( o.UpdatedTimeMs ),
      FollowId: int32 ( o.FollowId ),
    }
    return n
}
*/
/*
func PBConv_Like_PB_To_Like( o *Like_PB) *Like {
     n := &Like{
      Id: int ( o.Id ),
      PostId: int ( o.PostId ),
      PostTypeId: int ( o.PostTypeId ),
      UserId: int ( o.UserId ),
      TypeId: int ( o.TypeId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_Like_To_Like_PB ( o *Like) *Like_PB {
     n := &Like_PB{
      Id: int32 ( o.Id ),
      PostId: int32 ( o.PostId ),
      PostTypeId: int32 ( o.PostTypeId ),
      UserId: int32 ( o.UserId ),
      TypeId: int32 ( o.TypeId ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_Media_PB_To_Media( o *Media_PB) *Media {
     n := &Media{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      PostId: int ( o.PostId ),
      AlbumId: int ( o.AlbumId ),
      TypeId: int ( o.TypeId ),
      CreatedTime: int ( o.CreatedTime ),
      Src: string ( o.Src ),
    }
    return n
}

func PBConv_Media_To_Media_PB ( o *Media) *Media_PB {
     n := &Media_PB{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      PostId: int32 ( o.PostId ),
      AlbumId: int32 ( o.AlbumId ),
      TypeId: int32 ( o.TypeId ),
      CreatedTime: int32 ( o.CreatedTime ),
      Src: string ( o.Src ),
    }
    return n
}
*/
/*
func PBConv_Message_PB_To_Message( o *Message_PB) *Message {
     n := &Message{
      Id: int ( o.Id ),
      ToUserId: int ( o.ToUserId ),
      RoomKey: string ( o.RoomKey ),
      MessageKey: string ( o.MessageKey ),
      FromUserID: int ( o.FromUserID ),
      Data: string ( o.Data ),
      TimeMs: int ( o.TimeMs ),
    }
    return n
}

func PBConv_Message_To_Message_PB ( o *Message) *Message_PB {
     n := &Message_PB{
      Id: int64 ( o.Id ),
      ToUserId: int32 ( o.ToUserId ),
      RoomKey: string ( o.RoomKey ),
      MessageKey: string ( o.MessageKey ),
      FromUserID: int32 ( o.FromUserID ),
      Data: string ( o.Data ),
      TimeMs: int64 ( o.TimeMs ),
    }
    return n
}
*/
/*
func PBConv_MsgDeletedFromServer_PB_To_MsgDeletedFromServer( o *MsgDeletedFromServer_PB) *MsgDeletedFromServer {
     n := &MsgDeletedFromServer{
      Id: int ( o.Id ),
      ToUserId: int ( o.ToUserId ),
      MsgKey: string ( o.MsgKey ),
      PeerUserId: int ( o.PeerUserId ),
      RoomKey: string ( o.RoomKey ),
      AtTime: int ( o.AtTime ),
    }
    return n
}

func PBConv_MsgDeletedFromServer_To_MsgDeletedFromServer_PB ( o *MsgDeletedFromServer) *MsgDeletedFromServer_PB {
     n := &MsgDeletedFromServer_PB{
      Id: int64 ( o.Id ),
      ToUserId: int32 ( o.ToUserId ),
      MsgKey: string ( o.MsgKey ),
      PeerUserId: int32 ( o.PeerUserId ),
      RoomKey: string ( o.RoomKey ),
      AtTime: int32 ( o.AtTime ),
    }
    return n
}
*/
/*
func PBConv_MsgReceivedToPeer_PB_To_MsgReceivedToPeer( o *MsgReceivedToPeer_PB) *MsgReceivedToPeer {
     n := &MsgReceivedToPeer{
      Id: int ( o.Id ),
      ToUserId: int ( o.ToUserId ),
      MsgKey: string ( o.MsgKey ),
      RoomKey: string ( o.RoomKey ),
      PeerUserId: int ( o.PeerUserId ),
      AtTime: int ( o.AtTime ),
    }
    return n
}

func PBConv_MsgReceivedToPeer_To_MsgReceivedToPeer_PB ( o *MsgReceivedToPeer) *MsgReceivedToPeer_PB {
     n := &MsgReceivedToPeer_PB{
      Id: int64 ( o.Id ),
      ToUserId: int32 ( o.ToUserId ),
      MsgKey: string ( o.MsgKey ),
      RoomKey: string ( o.RoomKey ),
      PeerUserId: int32 ( o.PeerUserId ),
      AtTime: int32 ( o.AtTime ),
    }
    return n
}
*/
/*
func PBConv_MsgSeenByPeer_PB_To_MsgSeenByPeer( o *MsgSeenByPeer_PB) *MsgSeenByPeer {
     n := &MsgSeenByPeer{
      Id: int ( o.Id ),
      ToUserId: int ( o.ToUserId ),
      MsgKey: string ( o.MsgKey ),
      RoomKey: string ( o.RoomKey ),
      PeerUserId: int ( o.PeerUserId ),
      AtTime: int ( o.AtTime ),
    }
    return n
}

func PBConv_MsgSeenByPeer_To_MsgSeenByPeer_PB ( o *MsgSeenByPeer) *MsgSeenByPeer_PB {
     n := &MsgSeenByPeer_PB{
      Id: int64 ( o.Id ),
      ToUserId: int32 ( o.ToUserId ),
      MsgKey: string ( o.MsgKey ),
      RoomKey: string ( o.RoomKey ),
      PeerUserId: int32 ( o.PeerUserId ),
      AtTime: int32 ( o.AtTime ),
    }
    return n
}
*/
/*
func PBConv_Notification_PB_To_Notification( o *Notification_PB) *Notification {
     n := &Notification{
      Id: int ( o.Id ),
      ForUserId: int ( o.ForUserId ),
      ActorUserId: int ( o.ActorUserId ),
      ActionTypeId: int ( o.ActionTypeId ),
      ObjectTypeId: int ( o.ObjectTypeId ),
      RowId: int ( o.RowId ),
      RootId: int ( o.RootId ),
      RefId: int ( o.RefId ),
      SeenStatus: int ( o.SeenStatus ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_Notification_To_Notification_PB ( o *Notification) *Notification_PB {
     n := &Notification_PB{
      Id: int64 ( o.Id ),
      ForUserId: int32 ( o.ForUserId ),
      ActorUserId: int32 ( o.ActorUserId ),
      ActionTypeId: int32 ( o.ActionTypeId ),
      ObjectTypeId: int32 ( o.ObjectTypeId ),
      RowId: int32 ( o.RowId ),
      RootId: int32 ( o.RootId ),
      RefId: int32 ( o.RefId ),
      SeenStatus: int32 ( o.SeenStatus ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_NotificationRemoved_PB_To_NotificationRemoved( o *NotificationRemoved_PB) *NotificationRemoved {
     n := &NotificationRemoved{
      NotificationId: int ( o.NotificationId ),
      ForUserId: int ( o.ForUserId ),
    }
    return n
}

func PBConv_NotificationRemoved_To_NotificationRemoved_PB ( o *NotificationRemoved) *NotificationRemoved_PB {
     n := &NotificationRemoved_PB{
      NotificationId: int32 ( o.NotificationId ),
      ForUserId: int32 ( o.ForUserId ),
    }
    return n
}
*/
/*
func PBConv_PhoneContact_PB_To_PhoneContact( o *PhoneContact_PB) *PhoneContact {
     n := &PhoneContact{
      Id: int ( o.Id ),
      PhoneDisplayName: string ( o.PhoneDisplayName ),
      PhoneFamilyName: string ( o.PhoneFamilyName ),
      PhoneNumber: string ( o.PhoneNumber ),
      PhoneNormalizedNumber: string ( o.PhoneNormalizedNumber ),
      PhoneContactRowId: int ( o.PhoneContactRowId ),
      UserId: int ( o.UserId ),
      DeviceUuidId: int ( o.DeviceUuidId ),
      CreatedTime: int ( o.CreatedTime ),
      UpdatedTime: int ( o.UpdatedTime ),
    }
    return n
}

func PBConv_PhoneContact_To_PhoneContact_PB ( o *PhoneContact) *PhoneContact_PB {
     n := &PhoneContact_PB{
      Id: int32 ( o.Id ),
      PhoneDisplayName: string ( o.PhoneDisplayName ),
      PhoneFamilyName: string ( o.PhoneFamilyName ),
      PhoneNumber: string ( o.PhoneNumber ),
      PhoneNormalizedNumber: string ( o.PhoneNormalizedNumber ),
      PhoneContactRowId: int32 ( o.PhoneContactRowId ),
      UserId: int32 ( o.UserId ),
      DeviceUuidId: int32 ( o.DeviceUuidId ),
      CreatedTime: int32 ( o.CreatedTime ),
      UpdatedTime: int32 ( o.UpdatedTime ),
    }
    return n
}
*/
/*
func PBConv_Photo_PB_To_Photo( o *Photo_PB) *Photo {
     n := &Photo{
      PhotoId: int ( o.PhotoId ),
      UserId: int ( o.UserId ),
      PostId: int ( o.PostId ),
      AlbumId: int ( o.AlbumId ),
      ImageTypeId: int ( o.ImageTypeId ),
      Title: string ( o.Title ),
      Src: string ( o.Src ),
      PathSrc: string ( o.PathSrc ),
      BucketId: int ( o.BucketId ),
      Width: int ( o.Width ),
      Height: int ( o.Height ),
      Ratio: float32 ( o.Ratio ),
      HashMd5: string ( o.HashMd5 ),
      CreatedTime: int ( o.CreatedTime ),
      W1080: int ( o.W1080 ),
      W720: int ( o.W720 ),
      W480: int ( o.W480 ),
      W320: int ( o.W320 ),
      W160: int ( o.W160 ),
      W80: int ( o.W80 ),
    }
    return n
}

func PBConv_Photo_To_Photo_PB ( o *Photo) *Photo_PB {
     n := &Photo_PB{
      PhotoId: int32 ( o.PhotoId ),
      UserId: int32 ( o.UserId ),
      PostId: int32 ( o.PostId ),
      AlbumId: int32 ( o.AlbumId ),
      ImageTypeId: int32 ( o.ImageTypeId ),
      Title: string ( o.Title ),
      Src: string ( o.Src ),
      PathSrc: string ( o.PathSrc ),
      BucketId: int32 ( o.BucketId ),
      Width: int32 ( o.Width ),
      Height: int32 ( o.Height ),
      Ratio: float32 ( o.Ratio ),
      HashMd5: string ( o.HashMd5 ),
      CreatedTime: int32 ( o.CreatedTime ),
      W1080: int32 ( o.W1080 ),
      W720: int32 ( o.W720 ),
      W480: int32 ( o.W480 ),
      W320: int32 ( o.W320 ),
      W160: int32 ( o.W160 ),
      W80: int32 ( o.W80 ),
    }
    return n
}
*/
/*
func PBConv_Post_PB_To_Post( o *Post_PB) *Post {
     n := &Post{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      TypeId: int ( o.TypeId ),
      Text: string ( o.Text ),
      FormatedText: string ( o.FormatedText ),
      MediaCount: int ( o.MediaCount ),
      SharedTo: int ( o.SharedTo ),
      DisableComment: int ( o.DisableComment ),
      HasTag: int ( o.HasTag ),
      LikesCount: int ( o.LikesCount ),
      CommentsCount: int ( o.CommentsCount ),
      EditedTime: int ( o.EditedTime ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_Post_To_Post_PB ( o *Post) *Post_PB {
     n := &Post_PB{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      TypeId: int32 ( o.TypeId ),
      Text: string ( o.Text ),
      FormatedText: string ( o.FormatedText ),
      MediaCount: int32 ( o.MediaCount ),
      SharedTo: int32 ( o.SharedTo ),
      DisableComment: int32 ( o.DisableComment ),
      HasTag: int32 ( o.HasTag ),
      LikesCount: int32 ( o.LikesCount ),
      CommentsCount: int32 ( o.CommentsCount ),
      EditedTime: int32 ( o.EditedTime ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_RecommendUser_PB_To_RecommendUser( o *RecommendUser_PB) *RecommendUser {
     n := &RecommendUser{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      TargetId: int ( o.TargetId ),
      Weight: float32 ( o.Weight ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_RecommendUser_To_RecommendUser_PB ( o *RecommendUser) *RecommendUser_PB {
     n := &RecommendUser_PB{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      TargetId: int32 ( o.TargetId ),
      Weight: float32 ( o.Weight ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_SearchClicked_PB_To_SearchClicked( o *SearchClicked_PB) *SearchClicked {
     n := &SearchClicked{
      Id: int ( o.Id ),
      Query: string ( o.Query ),
      ClickType: int ( o.ClickType ),
      TargetId: int ( o.TargetId ),
      UserId: int ( o.UserId ),
      CreatedAt: int ( o.CreatedAt ),
    }
    return n
}

func PBConv_SearchClicked_To_SearchClicked_PB ( o *SearchClicked) *SearchClicked_PB {
     n := &SearchClicked_PB{
      Id: int64 ( o.Id ),
      Query: string ( o.Query ),
      ClickType: int32 ( o.ClickType ),
      TargetId: int32 ( o.TargetId ),
      UserId: int32 ( o.UserId ),
      CreatedAt: int32 ( o.CreatedAt ),
    }
    return n
}
*/
/*
func PBConv_Session_PB_To_Session( o *Session_PB) *Session {
     n := &Session{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      SessionUuid: string ( o.SessionUuid ),
      ClientUuid: string ( o.ClientUuid ),
      DeviceUuid: string ( o.DeviceUuid ),
      LastActivityTime: int ( o.LastActivityTime ),
      LastIpAddress: string ( o.LastIpAddress ),
      LastWifiMacAddress: string ( o.LastWifiMacAddress ),
      LastNetworkType: string ( o.LastNetworkType ),
      LastNetworkTypeId: int ( o.LastNetworkTypeId ),
      AppVersion: int ( o.AppVersion ),
      UpdatedTime: int ( o.UpdatedTime ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_Session_To_Session_PB ( o *Session) *Session_PB {
     n := &Session_PB{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      SessionUuid: string ( o.SessionUuid ),
      ClientUuid: string ( o.ClientUuid ),
      DeviceUuid: string ( o.DeviceUuid ),
      LastActivityTime: int32 ( o.LastActivityTime ),
      LastIpAddress: string ( o.LastIpAddress ),
      LastWifiMacAddress: string ( o.LastWifiMacAddress ),
      LastNetworkType: string ( o.LastNetworkType ),
      LastNetworkTypeId: int32 ( o.LastNetworkTypeId ),
      AppVersion: int32 ( o.AppVersion ),
      UpdatedTime: int32 ( o.UpdatedTime ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_Tag_PB_To_Tag( o *Tag_PB) *Tag {
     n := &Tag{
      Id: int ( o.Id ),
      Name: string ( o.Name ),
      Count: int ( o.Count ),
      IsBlocked: int ( o.IsBlocked ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_Tag_To_Tag_PB ( o *Tag) *Tag_PB {
     n := &Tag_PB{
      Id: int32 ( o.Id ),
      Name: string ( o.Name ),
      Count: int32 ( o.Count ),
      IsBlocked: int32 ( o.IsBlocked ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_TagsPost_PB_To_TagsPost( o *TagsPost_PB) *TagsPost {
     n := &TagsPost{
      Id: int ( o.Id ),
      TagId: int ( o.TagId ),
      PostId: int ( o.PostId ),
      TypeId: int ( o.TypeId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_TagsPost_To_TagsPost_PB ( o *TagsPost) *TagsPost_PB {
     n := &TagsPost_PB{
      Id: int32 ( o.Id ),
      TagId: int32 ( o.TagId ),
      PostId: int32 ( o.PostId ),
      TypeId: int32 ( o.TypeId ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConv_User_PB_To_User( o *User_PB) *User {n := &User{}
   n.Id = int ( o.Id )
   n.UserName = string ( o.UserName )
   n.UserNameLower = string ( o.UserNameLower )
   n.FirstName = string ( o.FirstName )
   n.LastName = string ( o.LastName )
   n.About = string ( o.About )
   n.FullName = string ( o.FullName )
   n.AvatarUrl = string ( o.AvatarUrl )
   n.PrivacyProfile = int ( o.PrivacyProfile )
   n.Phone = string ( o.Phone )
   n.Email = string ( o.Email )
   n.IsDeleted = int ( o.IsDeleted )
   n.PasswordHash = string ( o.PasswordHash )
   n.PasswordSalt = string ( o.PasswordSalt )
   n.FollowersCount = int ( o.FollowersCount )
   n.FollowingCount = int ( o.FollowingCount )
   n.PostsCount = int ( o.PostsCount )
   n.MediaCount = int ( o.MediaCount )
   n.LikesCount = int ( o.LikesCount )
   n.ResharedCount = int ( o.ResharedCount )
   n.LastActionTime = int ( o.LastActionTime )
   n.LastPostTime = int ( o.LastPostTime )
   n.PrimaryFollowingList = int ( o.PrimaryFollowingList )
   n.CreatedTime = int ( o.CreatedTime )
   n.UpdatedTime = int ( o.UpdatedTime )
   n.SessionUuid = string ( o.SessionUuid )
   n.DeviceUuid = string ( o.DeviceUuid )
   n.LastWifiMacAddress = string ( o.LastWifiMacAddress )
   n.LastNetworkType = string ( o.LastNetworkType )
   n.AppVersion = int ( o.AppVersion )
   n.LastActivityTime = int ( o.LastActivityTime )
   n.LastLoginTime = int ( o.LastLoginTime )
   n.LastIpAddress = string ( o.LastIpAddress )
    return n
}

func PBConv_User_To_User_PB ( o *User) *User_PB {n := &User_PB{}
   n.Id = int32 ( o.Id )
   n.UserName = string ( o.UserName )
   n.UserNameLower = string ( o.UserNameLower )
   n.FirstName = string ( o.FirstName )
   n.LastName = string ( o.LastName )
   n.About = string ( o.About )
   n.FullName = string ( o.FullName )
   n.AvatarUrl = string ( o.AvatarUrl )
   n.PrivacyProfile = int32 ( o.PrivacyProfile )
   n.Phone = string ( o.Phone )
   n.Email = string ( o.Email )
   n.IsDeleted = int32 ( o.IsDeleted )
   n.PasswordHash = string ( o.PasswordHash )
   n.PasswordSalt = string ( o.PasswordSalt )
   n.FollowersCount = int32 ( o.FollowersCount )
   n.FollowingCount = int32 ( o.FollowingCount )
   n.PostsCount = int32 ( o.PostsCount )
   n.MediaCount = int32 ( o.MediaCount )
   n.LikesCount = int32 ( o.LikesCount )
   n.ResharedCount = int32 ( o.ResharedCount )
   n.LastActionTime = int32 ( o.LastActionTime )
   n.LastPostTime = int32 ( o.LastPostTime )
   n.PrimaryFollowingList = int32 ( o.PrimaryFollowingList )
   n.CreatedTime = int32 ( o.CreatedTime )
   n.UpdatedTime = int32 ( o.UpdatedTime )
   n.SessionUuid = string ( o.SessionUuid )
   n.DeviceUuid = string ( o.DeviceUuid )
   n.LastWifiMacAddress = string ( o.LastWifiMacAddress )
   n.LastNetworkType = string ( o.LastNetworkType )
   n.AppVersion = int32 ( o.AppVersion )
   n.LastActivityTime = int32 ( o.LastActivityTime )
   n.LastLoginTime = int32 ( o.LastLoginTime )
   n.LastIpAddress = string ( o.LastIpAddress )
    return n
}
*/
/*
func PBConv_UserMetaInfo_PB_To_UserMetaInfo( o *UserMetaInfo_PB) *UserMetaInfo {
     n := &UserMetaInfo{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      IsNotificationDirty: int ( o.IsNotificationDirty ),
      LastUserRecGen: int ( o.LastUserRecGen ),
    }
    return n
}

func PBConv_UserMetaInfo_To_UserMetaInfo_PB ( o *UserMetaInfo) *UserMetaInfo_PB {
     n := &UserMetaInfo_PB{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      IsNotificationDirty: int32 ( o.IsNotificationDirty ),
      LastUserRecGen: int32 ( o.LastUserRecGen ),
    }
    return n
}
*/
/*
func PBConv_UserPassword_PB_To_UserPassword( o *UserPassword_PB) *UserPassword {
     n := &UserPassword{
      UserId: int ( o.UserId ),
      Password: string ( o.Password ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConv_UserPassword_To_UserPassword_PB ( o *UserPassword) *UserPassword_PB {
     n := &UserPassword_PB{
      UserId: int32 ( o.UserId ),
      Password: string ( o.Password ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
