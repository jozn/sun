
package x


/*
func PBConvPB__Activity_To_Activity( o *PB_Activity) *Activity {
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

func PBConvPB_Activity_To_Activity ( o *Activity) *PB_Activity {
     n := &PB_Activity{
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
func PBConvPB__Bucket_To_Bucket( o *PB_Bucket) *Bucket {
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

func PBConvPB_Bucket_To_Bucket ( o *Bucket) *PB_Bucket {
     n := &PB_Bucket{
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
func PBConvPB__Comment_To_Comment( o *PB_Comment) *Comment {
     n := &Comment{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      PostId: int ( o.PostId ),
      Text: string ( o.Text ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Comment_To_Comment ( o *Comment) *PB_Comment {
     n := &PB_Comment{
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
func PBConvPB__FollowingList_To_FollowingList( o *PB_FollowingList) *FollowingList {
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

func PBConvPB_FollowingList_To_FollowingList ( o *FollowingList) *PB_FollowingList {
     n := &PB_FollowingList{
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
func PBConvPB__FollowingListMember_To_FollowingListMember( o *PB_FollowingListMember) *FollowingListMember {
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

func PBConvPB_FollowingListMember_To_FollowingListMember ( o *FollowingListMember) *PB_FollowingListMember {
     n := &PB_FollowingListMember{
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
func PBConvPB__FollowingListMemberHistory_To_FollowingListMemberHistory( o *PB_FollowingListMemberHistory) *FollowingListMemberHistory {
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

func PBConvPB_FollowingListMemberHistory_To_FollowingListMemberHistory ( o *FollowingListMemberHistory) *PB_FollowingListMemberHistory {
     n := &PB_FollowingListMemberHistory{
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
func PBConvPB__Like_To_Like( o *PB_Like) *Like {
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

func PBConvPB_Like_To_Like ( o *Like) *PB_Like {
     n := &PB_Like{
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
func PBConvPB__Media_To_Media( o *PB_Media) *Media {
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

func PBConvPB_Media_To_Media ( o *Media) *PB_Media {
     n := &PB_Media{
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
func PBConvPB__Message_To_Message( o *PB_Message) *Message {
     n := &Message{
      Id: int ( o.Id ),
      Uid: int ( o.Uid ),
      UserId: int ( o.UserId ),
      MessageKey: string ( o.MessageKey ),
      RoomKey: string ( o.RoomKey ),
      MessageType: int ( o.MessageType ),
      RoomType: int ( o.RoomType ),
      MsgFileId: int ( o.MsgFileId ),
      DataPB: []byte ( o.DataPB ),
      Data64: string ( o.Data64 ),
      DataJson: string ( o.DataJson ),
      CreatedTimeMs: int ( o.CreatedTimeMs ),
    }
    return n
}

func PBConvPB_Message_To_Message ( o *Message) *PB_Message {
     n := &PB_Message{
      Id: int64 ( o.Id ),
      Uid: int64 ( o.Uid ),
      UserId: int64 ( o.UserId ),
      MessageKey: string ( o.MessageKey ),
      RoomKey: string ( o.RoomKey ),
      MessageType: int32 ( o.MessageType ),
      RoomType: int32 ( o.RoomType ),
      MsgFileId: int64 ( o.MsgFileId ),
      DataPB: []byte ( o.DataPB ),
      Data64: string ( o.Data64 ),
      DataJson: string ( o.DataJson ),
      CreatedTimeMs: int64 ( o.CreatedTimeMs ),
    }
    return n
}
*/
/*
func PBConvPB__MsgFile_To_MsgFile( o *PB_MsgFile) *MsgFile {
     n := &MsgFile{
      Id: int ( o.Id ),
      Name: string ( o.Name ),
      Size: int ( o.Size ),
      FileType: int ( o.FileType ),
      MimeType: string ( o.MimeType ),
      Width: int ( o.Width ),
      Height: int ( o.Height ),
      Duration: int ( o.Duration ),
      Extension: string ( o.Extension ),
      ThumbData: []byte ( o.ThumbData ),
      ThumbData64: string ( o.ThumbData64 ),
      ServerSrc: string ( o.ServerSrc ),
      ServerPath: string ( o.ServerPath ),
      ServerId: int ( o.ServerId ),
      CanDel: int ( o.CanDel ),
    }
    return n
}

func PBConvPB_MsgFile_To_MsgFile ( o *MsgFile) *PB_MsgFile {
     n := &PB_MsgFile{
      Id: int64 ( o.Id ),
      Name: string ( o.Name ),
      Size: int32 ( o.Size ),
      FileType: int32 ( o.FileType ),
      MimeType: string ( o.MimeType ),
      Width: int32 ( o.Width ),
      Height: int32 ( o.Height ),
      Duration: int32 ( o.Duration ),
      Extension: string ( o.Extension ),
      ThumbData: []byte ( o.ThumbData ),
      ThumbData64: string ( o.ThumbData64 ),
      ServerSrc: string ( o.ServerSrc ),
      ServerPath: string ( o.ServerPath ),
      ServerId: int32 ( o.ServerId ),
      CanDel: int32 ( o.CanDel ),
    }
    return n
}
*/
/*
func PBConvPB__MsgPush_To_MsgPush( o *PB_MsgPush) *MsgPush {
     n := &MsgPush{
      Id: int ( o.Id ),
      Uid: int ( o.Uid ),
      ToUser: int ( o.ToUser ),
      MsgUid: int ( o.MsgUid ),
      CreatedTimeMs: int ( o.CreatedTimeMs ),
    }
    return n
}

func PBConvPB_MsgPush_To_MsgPush ( o *MsgPush) *PB_MsgPush {
     n := &PB_MsgPush{
      Id: int64 ( o.Id ),
      Uid: int64 ( o.Uid ),
      ToUser: int64 ( o.ToUser ),
      MsgUid: int64 ( o.MsgUid ),
      CreatedTimeMs: int64 ( o.CreatedTimeMs ),
    }
    return n
}
*/
/*
func PBConvPB__MsgPushEvent_To_MsgPushEvent( o *PB_MsgPushEvent) *MsgPushEvent {
     n := &MsgPushEvent{
      Id: int ( o.Id ),
      Uid: int ( o.Uid ),
      ToUserId: int ( o.ToUserId ),
      MsgUid: int ( o.MsgUid ),
      MsgKey: string ( o.MsgKey ),
      RoomKey: string ( o.RoomKey ),
      PeerUserId: int ( o.PeerUserId ),
      EventType: int ( o.EventType ),
      AtTime: int ( o.AtTime ),
    }
    return n
}

func PBConvPB_MsgPushEvent_To_MsgPushEvent ( o *MsgPushEvent) *PB_MsgPushEvent {
     n := &PB_MsgPushEvent{
      Id: int64 ( o.Id ),
      Uid: int64 ( o.Uid ),
      ToUserId: int32 ( o.ToUserId ),
      MsgUid: int64 ( o.MsgUid ),
      MsgKey: string ( o.MsgKey ),
      RoomKey: string ( o.RoomKey ),
      PeerUserId: int32 ( o.PeerUserId ),
      EventType: int32 ( o.EventType ),
      AtTime: int32 ( o.AtTime ),
    }
    return n
}
*/
/*
func PBConvPB__Notification_To_Notification( o *PB_Notification) *Notification {
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

func PBConvPB_Notification_To_Notification ( o *Notification) *PB_Notification {
     n := &PB_Notification{
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
func PBConvPB__NotificationRemoved_To_NotificationRemoved( o *PB_NotificationRemoved) *NotificationRemoved {
     n := &NotificationRemoved{
      NotificationId: int ( o.NotificationId ),
      ForUserId: int ( o.ForUserId ),
    }
    return n
}

func PBConvPB_NotificationRemoved_To_NotificationRemoved ( o *NotificationRemoved) *PB_NotificationRemoved {
     n := &PB_NotificationRemoved{
      NotificationId: int32 ( o.NotificationId ),
      ForUserId: int32 ( o.ForUserId ),
    }
    return n
}
*/
/*
func PBConvPB__PhoneContact_To_PhoneContact( o *PB_PhoneContact) *PhoneContact {
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

func PBConvPB_PhoneContact_To_PhoneContact ( o *PhoneContact) *PB_PhoneContact {
     n := &PB_PhoneContact{
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
func PBConvPB__Photo_To_Photo( o *PB_Photo) *Photo {
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

func PBConvPB_Photo_To_Photo ( o *Photo) *PB_Photo {
     n := &PB_Photo{
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
func PBConvPB__Post_To_Post( o *PB_Post) *Post {
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

func PBConvPB_Post_To_Post ( o *Post) *PB_Post {
     n := &PB_Post{
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
func PBConvPB__RecommendUser_To_RecommendUser( o *PB_RecommendUser) *RecommendUser {
     n := &RecommendUser{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      TargetId: int ( o.TargetId ),
      Weight: float32 ( o.Weight ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_RecommendUser_To_RecommendUser ( o *RecommendUser) *PB_RecommendUser {
     n := &PB_RecommendUser{
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
func PBConvPB__SearchClicked_To_SearchClicked( o *PB_SearchClicked) *SearchClicked {
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

func PBConvPB_SearchClicked_To_SearchClicked ( o *SearchClicked) *PB_SearchClicked {
     n := &PB_SearchClicked{
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
func PBConvPB__Session_To_Session( o *PB_Session) *Session {
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

func PBConvPB_Session_To_Session ( o *Session) *PB_Session {
     n := &PB_Session{
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
func PBConvPB__Tag_To_Tag( o *PB_Tag) *Tag {
     n := &Tag{
      Id: int ( o.Id ),
      Name: string ( o.Name ),
      Count: int ( o.Count ),
      IsBlocked: int ( o.IsBlocked ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Tag_To_Tag ( o *Tag) *PB_Tag {
     n := &PB_Tag{
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
func PBConvPB__TagsPost_To_TagsPost( o *PB_TagsPost) *TagsPost {
     n := &TagsPost{
      Id: int ( o.Id ),
      TagId: int ( o.TagId ),
      PostId: int ( o.PostId ),
      TypeId: int ( o.TypeId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_TagsPost_To_TagsPost ( o *TagsPost) *PB_TagsPost {
     n := &PB_TagsPost{
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
func PBConvPB__User_To_User( o *PB_User) *User {n := &User{}
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

func PBConvPB_User_To_User ( o *User) *PB_User {n := &PB_User{}
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
func PBConvPB__UserMetaInfo_To_UserMetaInfo( o *PB_UserMetaInfo) *UserMetaInfo {
     n := &UserMetaInfo{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      IsNotificationDirty: int ( o.IsNotificationDirty ),
      LastUserRecGen: int ( o.LastUserRecGen ),
    }
    return n
}

func PBConvPB_UserMetaInfo_To_UserMetaInfo ( o *UserMetaInfo) *PB_UserMetaInfo {
     n := &PB_UserMetaInfo{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      IsNotificationDirty: int32 ( o.IsNotificationDirty ),
      LastUserRecGen: int32 ( o.LastUserRecGen ),
    }
    return n
}
*/
/*
func PBConvPB__UserPassword_To_UserPassword( o *PB_UserPassword) *UserPassword {
     n := &UserPassword{
      UserId: int ( o.UserId ),
      Password: string ( o.Password ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_UserPassword_To_UserPassword ( o *UserPassword) *PB_UserPassword {
     n := &PB_UserPassword{
      UserId: int32 ( o.UserId ),
      Password: string ( o.Password ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
