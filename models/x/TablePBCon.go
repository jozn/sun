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
func PBConvPB__Chat_To_Chat( o *PB_Chat) *Chat {
     n := &Chat{
      ChatKey: string ( o.ChatKey ),
      RoomKey: string ( o.RoomKey ),
      RoomTypeEnumId: int ( o.RoomTypeEnumId ),
      UserId: int ( o.UserId ),
      PeerUserId: int ( o.PeerUserId ),
      GroupId: int ( o.GroupId ),
      CreatedSe: int ( o.CreatedSe ),
      StartMessageIdFrom: int ( o.StartMessageIdFrom ),
      LastDeletedMessageId: int ( o.LastDeletedMessageId ),
      LastSeenMessageId: int ( o.LastSeenMessageId ),
      LastMessageId: int ( o.LastMessageId ),
      UpdatedMs: int ( o.UpdatedMs ),
    }
    return n
}

func PBConvPB_Chat_To_Chat ( o *Chat) *PB_Chat {
     n := &PB_Chat{
      ChatKey: string ( o.ChatKey ),
      RoomKey: string ( o.RoomKey ),
      RoomTypeEnumId: int32 ( o.RoomTypeEnumId ),
      UserId: int32 ( o.UserId ),
      PeerUserId: int32 ( o.PeerUserId ),
      GroupId: int64 ( o.GroupId ),
      CreatedSe: int32 ( o.CreatedSe ),
      StartMessageIdFrom: int64 ( o.StartMessageIdFrom ),
      LastDeletedMessageId: int64 ( o.LastDeletedMessageId ),
      LastSeenMessageId: int64 ( o.LastSeenMessageId ),
      LastMessageId: int64 ( o.LastMessageId ),
      UpdatedMs: int64 ( o.UpdatedMs ),
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
func PBConvPB__DirectMessage_To_DirectMessage( o *PB_DirectMessage) *DirectMessage {
     n := &DirectMessage{
      MessageId: int ( o.MessageId ),
      MessageKey: string ( o.MessageKey ),
      RoomKey: string ( o.RoomKey ),
      UserId: int ( o.UserId ),
      MessageFileId: int ( o.MessageFileId ),
      MessageTypeEnumId: int ( o.MessageTypeEnumId ),
      Text: string ( o.Text ),
      CreatedSe: int ( o.CreatedSe ),
      PeerReceivedTime: int ( o.PeerReceivedTime ),
      PeerSeenTime: int ( o.PeerSeenTime ),
      DeliviryStatusEnumId: int ( o.DeliviryStatusEnumId ),
    }
    return n
}

func PBConvPB_DirectMessage_To_DirectMessage ( o *DirectMessage) *PB_DirectMessage {
     n := &PB_DirectMessage{
      MessageId: int64 ( o.MessageId ),
      MessageKey: string ( o.MessageKey ),
      RoomKey: string ( o.RoomKey ),
      UserId: int32 ( o.UserId ),
      MessageFileId: int64 ( o.MessageFileId ),
      MessageTypeEnumId: int32 ( o.MessageTypeEnumId ),
      Text: string ( o.Text ),
      CreatedSe: int32 ( o.CreatedSe ),
      PeerReceivedTime: int32 ( o.PeerReceivedTime ),
      PeerSeenTime: int32 ( o.PeerSeenTime ),
      DeliviryStatusEnumId: int32 ( o.DeliviryStatusEnumId ),
    }
    return n
}
*/
/*
func PBConvPB__DirectOffline_To_DirectOffline( o *PB_DirectOffline) *DirectOffline {
     n := &DirectOffline{
      DirectOfflineId: int ( o.DirectOfflineId ),
      ToUserId: int ( o.ToUserId ),
      ChatKey: string ( o.ChatKey ),
      PBClass: string ( o.PBClass ),
      DataPB: []byte ( o.DataPB ),
      DataJson: string ( o.DataJson ),
      DataTemp: string ( o.DataTemp ),
      AtTimeMs: int ( o.AtTimeMs ),
    }
    return n
}

func PBConvPB_DirectOffline_To_DirectOffline ( o *DirectOffline) *PB_DirectOffline {
     n := &PB_DirectOffline{
      DirectOfflineId: int64 ( o.DirectOfflineId ),
      ToUserId: int32 ( o.ToUserId ),
      ChatKey: string ( o.ChatKey ),
      PBClass: string ( o.PBClass ),
      DataPB: []byte ( o.DataPB ),
      DataJson: string ( o.DataJson ),
      DataTemp: string ( o.DataTemp ),
      AtTimeMs: int64 ( o.AtTimeMs ),
    }
    return n
}
*/
/*
func PBConvPB__DirectOfflineDep_To_DirectOfflineDep( o *PB_DirectOfflineDep) *DirectOfflineDep {
     n := &DirectOfflineDep{
      DirectOfflineId: int ( o.DirectOfflineId ),
      ToUserId: int ( o.ToUserId ),
      MessageId: int ( o.MessageId ),
      MessageFileId: int ( o.MessageFileId ),
      OtherId: int ( o.OtherId ),
      ChatKey: string ( o.ChatKey ),
      PeerUserId: int ( o.PeerUserId ),
      RoomLogTypeId: int ( o.RoomLogTypeId ),
      PBClass: string ( o.PBClass ),
      DataPB: []byte ( o.DataPB ),
      DataJson: string ( o.DataJson ),
      DataTemp: string ( o.DataTemp ),
      AtTimeMs: int ( o.AtTimeMs ),
    }
    return n
}

func PBConvPB_DirectOfflineDep_To_DirectOfflineDep ( o *DirectOfflineDep) *PB_DirectOfflineDep {
     n := &PB_DirectOfflineDep{
      DirectOfflineId: int64 ( o.DirectOfflineId ),
      ToUserId: int32 ( o.ToUserId ),
      MessageId: int64 ( o.MessageId ),
      MessageFileId: int64 ( o.MessageFileId ),
      OtherId: int64 ( o.OtherId ),
      ChatKey: string ( o.ChatKey ),
      PeerUserId: int32 ( o.PeerUserId ),
      RoomLogTypeId: int32 ( o.RoomLogTypeId ),
      PBClass: string ( o.PBClass ),
      DataPB: []byte ( o.DataPB ),
      DataJson: string ( o.DataJson ),
      DataTemp: string ( o.DataTemp ),
      AtTimeMs: int64 ( o.AtTimeMs ),
    }
    return n
}
*/
/*
func PBConvPB__DirectToMessage_To_DirectToMessage( o *PB_DirectToMessage) *DirectToMessage {
     n := &DirectToMessage{
      Id: int ( o.Id ),
      ChatKey: string ( o.ChatKey ),
      MessageId: int ( o.MessageId ),
      SourceEnumId: int ( o.SourceEnumId ),
    }
    return n
}

func PBConvPB_DirectToMessage_To_DirectToMessage ( o *DirectToMessage) *PB_DirectToMessage {
     n := &PB_DirectToMessage{
      Id: int64 ( o.Id ),
      ChatKey: string ( o.ChatKey ),
      MessageId: int64 ( o.MessageId ),
      SourceEnumId: int32 ( o.SourceEnumId ),
    }
    return n
}
*/
/*
func PBConvPB__DirectUpdate_To_DirectUpdate( o *PB_DirectUpdate) *DirectUpdate {
     n := &DirectUpdate{
      DirectUpdateId: int ( o.DirectUpdateId ),
      ToUserId: int ( o.ToUserId ),
      MessageId: int ( o.MessageId ),
      MessageFileId: int ( o.MessageFileId ),
      OtherId: int ( o.OtherId ),
      ChatKey: string ( o.ChatKey ),
      PeerUserId: int ( o.PeerUserId ),
      EventType: int ( o.EventType ),
      RoomLogTypeId: int ( o.RoomLogTypeId ),
      FromSeq: int ( o.FromSeq ),
      ToSeq: int ( o.ToSeq ),
      ExtraPB: []byte ( o.ExtraPB ),
      ExtraJson: string ( o.ExtraJson ),
      AtTimeMs: int ( o.AtTimeMs ),
    }
    return n
}

func PBConvPB_DirectUpdate_To_DirectUpdate ( o *DirectUpdate) *PB_DirectUpdate {
     n := &PB_DirectUpdate{
      DirectUpdateId: int64 ( o.DirectUpdateId ),
      ToUserId: int32 ( o.ToUserId ),
      MessageId: int64 ( o.MessageId ),
      MessageFileId: int64 ( o.MessageFileId ),
      OtherId: int64 ( o.OtherId ),
      ChatKey: string ( o.ChatKey ),
      PeerUserId: int32 ( o.PeerUserId ),
      EventType: int32 ( o.EventType ),
      RoomLogTypeId: int32 ( o.RoomLogTypeId ),
      FromSeq: int32 ( o.FromSeq ),
      ToSeq: int32 ( o.ToSeq ),
      ExtraPB: []byte ( o.ExtraPB ),
      ExtraJson: string ( o.ExtraJson ),
      AtTimeMs: int64 ( o.AtTimeMs ),
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
func PBConvPB__GeneralLog_To_GeneralLog( o *PB_GeneralLog) *GeneralLog {
     n := &GeneralLog{
      Id: int ( o.Id ),
      ToUserId: int ( o.ToUserId ),
      TargetId: int ( o.TargetId ),
      LogTypeId: int ( o.LogTypeId ),
      ExtraPB: []byte ( o.ExtraPB ),
      ExtraJson: string ( o.ExtraJson ),
      CreatedMs: int ( o.CreatedMs ),
    }
    return n
}

func PBConvPB_GeneralLog_To_GeneralLog ( o *GeneralLog) *PB_GeneralLog {
     n := &PB_GeneralLog{
      Id: int64 ( o.Id ),
      ToUserId: int32 ( o.ToUserId ),
      TargetId: int32 ( o.TargetId ),
      LogTypeId: int32 ( o.LogTypeId ),
      ExtraPB: []byte ( o.ExtraPB ),
      ExtraJson: string ( o.ExtraJson ),
      CreatedMs: int64 ( o.CreatedMs ),
    }
    return n
}
*/
/*
func PBConvPB__Group_To_Group( o *PB_Group) *Group {
     n := &Group{
      GroupId: int ( o.GroupId ),
      GroupName: string ( o.GroupName ),
      MembersCount: int ( o.MembersCount ),
      GroupPrivacyEnum: int ( o.GroupPrivacyEnum ),
      CreatorUserId: int ( o.CreatorUserId ),
      CreatedTime: int ( o.CreatedTime ),
      UpdatedMs: int ( o.UpdatedMs ),
      CurrentSeq: int ( o.CurrentSeq ),
    }
    return n
}

func PBConvPB_Group_To_Group ( o *Group) *PB_Group {
     n := &PB_Group{
      GroupId: int64 ( o.GroupId ),
      GroupName: string ( o.GroupName ),
      MembersCount: int32 ( o.MembersCount ),
      GroupPrivacyEnum: int32 ( o.GroupPrivacyEnum ),
      CreatorUserId: int32 ( o.CreatorUserId ),
      CreatedTime: int32 ( o.CreatedTime ),
      UpdatedMs: int64 ( o.UpdatedMs ),
      CurrentSeq: int32 ( o.CurrentSeq ),
    }
    return n
}
*/
/*
func PBConvPB__GroupMember_To_GroupMember( o *PB_GroupMember) *GroupMember {
     n := &GroupMember{
      Id: int ( o.Id ),
      GroupId: int ( o.GroupId ),
      GroupKey: string ( o.GroupKey ),
      UserId: int ( o.UserId ),
      ByUserId: int ( o.ByUserId ),
      GroupRoleEnumId: int ( o.GroupRoleEnumId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_GroupMember_To_GroupMember ( o *GroupMember) *PB_GroupMember {
     n := &PB_GroupMember{
      Id: int64 ( o.Id ),
      GroupId: int64 ( o.GroupId ),
      GroupKey: string ( o.GroupKey ),
      UserId: int32 ( o.UserId ),
      ByUserId: int32 ( o.ByUserId ),
      GroupRoleEnumId: int32 ( o.GroupRoleEnumId ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__GroupMessage_To_GroupMessage( o *PB_GroupMessage) *GroupMessage {
     n := &GroupMessage{
      MessageId: int ( o.MessageId ),
      RoomKey: string ( o.RoomKey ),
      UserId: int ( o.UserId ),
      MessageFileId: int ( o.MessageFileId ),
      MessageTypeEnum: int ( o.MessageTypeEnum ),
      Text: string ( o.Text ),
      CreatedMs: int ( o.CreatedMs ),
      DeliveryStatusEnum: int ( o.DeliveryStatusEnum ),
    }
    return n
}

func PBConvPB_GroupMessage_To_GroupMessage ( o *GroupMessage) *PB_GroupMessage {
     n := &PB_GroupMessage{
      MessageId: int64 ( o.MessageId ),
      RoomKey: string ( o.RoomKey ),
      UserId: int32 ( o.UserId ),
      MessageFileId: int64 ( o.MessageFileId ),
      MessageTypeEnum: int32 ( o.MessageTypeEnum ),
      Text: string ( o.Text ),
      CreatedMs: int64 ( o.CreatedMs ),
      DeliveryStatusEnum: int32 ( o.DeliveryStatusEnum ),
    }
    return n
}
*/
/*
func PBConvPB__GroupToMessage_To_GroupToMessage( o *PB_GroupToMessage) *GroupToMessage {
     n := &GroupToMessage{
      Id: int ( o.Id ),
      GroupId: int ( o.GroupId ),
      MessageId: int ( o.MessageId ),
      CreatedMs: int ( o.CreatedMs ),
      StatusEnum: int ( o.StatusEnum ),
    }
    return n
}

func PBConvPB_GroupToMessage_To_GroupToMessage ( o *GroupToMessage) *PB_GroupToMessage {
     n := &PB_GroupToMessage{
      Id: int64 ( o.Id ),
      GroupId: int64 ( o.GroupId ),
      MessageId: int64 ( o.MessageId ),
      CreatedMs: int64 ( o.CreatedMs ),
      StatusEnum: int32 ( o.StatusEnum ),
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
func PBConvPB__LogChange_To_LogChange( o *PB_LogChange) *LogChange {
     n := &LogChange{
      Id: int ( o.Id ),
      T: string ( o.T ),
    }
    return n
}

func PBConvPB_LogChange_To_LogChange ( o *LogChange) *PB_LogChange {
     n := &PB_LogChange{
      Id: int32 ( o.Id ),
      T: string ( o.T ),
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
func PBConvPB__MessageFile_To_MessageFile( o *PB_MessageFile) *MessageFile {
     n := &MessageFile{
      MessageFileId: int ( o.MessageFileId ),
      MessageFileKey: string ( o.MessageFileKey ),
      OriginalUserId: int ( o.OriginalUserId ),
      Name: string ( o.Name ),
      Size: int ( o.Size ),
      FileTypeEnumId: int ( o.FileTypeEnumId ),
      Width: int ( o.Width ),
      Height: int ( o.Height ),
      Duration: int ( o.Duration ),
      Extension: string ( o.Extension ),
      HashMd5: string ( o.HashMd5 ),
      HashAccess: int ( o.HashAccess ),
      CreatedSe: int ( o.CreatedSe ),
      ServerSrc: string ( o.ServerSrc ),
      ServerPath: string ( o.ServerPath ),
      ServerThumbPath: string ( o.ServerThumbPath ),
      BucketId: string ( o.BucketId ),
      ServerId: int ( o.ServerId ),
      CanDel: int ( o.CanDel ),
    }
    return n
}

func PBConvPB_MessageFile_To_MessageFile ( o *MessageFile) *PB_MessageFile {
     n := &PB_MessageFile{
      MessageFileId: int64 ( o.MessageFileId ),
      MessageFileKey: string ( o.MessageFileKey ),
      OriginalUserId: int32 ( o.OriginalUserId ),
      Name: string ( o.Name ),
      Size: int32 ( o.Size ),
      FileTypeEnumId: int32 ( o.FileTypeEnumId ),
      Width: int32 ( o.Width ),
      Height: int32 ( o.Height ),
      Duration: int32 ( o.Duration ),
      Extension: string ( o.Extension ),
      HashMd5: string ( o.HashMd5 ),
      HashAccess: int64 ( o.HashAccess ),
      CreatedSe: int32 ( o.CreatedSe ),
      ServerSrc: string ( o.ServerSrc ),
      ServerPath: string ( o.ServerPath ),
      ServerThumbPath: string ( o.ServerThumbPath ),
      BucketId: string ( o.BucketId ),
      ServerId: int32 ( o.ServerId ),
      CanDel: int32 ( o.CanDel ),
    }
    return n
}
*/
/*
func PBConvPB__Msg_To_Msg( o *PB_Msg) *Msg {
     n := &Msg{
      Key: string ( o.Key ),
      Name: string ( o.Name ),
      Id: int ( o.Id ),
    }
    return n
}

func PBConvPB_Msg_To_Msg ( o *Msg) *PB_Msg {
     n := &PB_Msg{
      Key: string ( o.Key ),
      Name: string ( o.Name ),
      Id: int32 ( o.Id ),
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
func PBConvPB__Offline_To_Offline( o *PB_Offline) *Offline {
     n := &Offline{
      Id: int ( o.Id ),
      FromUserId: int ( o.FromUserId ),
      ToUserId: int ( o.ToUserId ),
      RpcName: string ( o.RpcName ),
      PBClassName: string ( o.PBClassName ),
      Key: string ( o.Key ),
      DataJson: string ( o.DataJson ),
      DataBlob: []byte ( o.DataBlob ),
      DataLength: int ( o.DataLength ),
      CreatedMs: int ( o.CreatedMs ),
    }
    return n
}

func PBConvPB_Offline_To_Offline ( o *Offline) *PB_Offline {
     n := &PB_Offline{
      Id: int64 ( o.Id ),
      FromUserId: int32 ( o.FromUserId ),
      ToUserId: int32 ( o.ToUserId ),
      RpcName: string ( o.RpcName ),
      PBClassName: string ( o.PBClassName ),
      Key: string ( o.Key ),
      DataJson: string ( o.DataJson ),
      DataBlob: []byte ( o.DataBlob ),
      DataLength: int32 ( o.DataLength ),
      CreatedMs: int64 ( o.CreatedMs ),
    }
    return n
}
*/
/*
func PBConvPB__OldMessage_To_OldMessage( o *PB_OldMessage) *OldMessage {
     n := &OldMessage{
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

func PBConvPB_OldMessage_To_OldMessage ( o *OldMessage) *PB_OldMessage {
     n := &PB_OldMessage{
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
func PBConvPB__OldMsgFile_To_OldMsgFile( o *PB_OldMsgFile) *OldMsgFile {
     n := &OldMsgFile{
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

func PBConvPB_OldMsgFile_To_OldMsgFile ( o *OldMsgFile) *PB_OldMsgFile {
     n := &PB_OldMsgFile{
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
func PBConvPB__OldMsgPush_To_OldMsgPush( o *PB_OldMsgPush) *OldMsgPush {
     n := &OldMsgPush{
      Id: int ( o.Id ),
      Uid: int ( o.Uid ),
      ToUser: int ( o.ToUser ),
      MsgUid: int ( o.MsgUid ),
      CreatedTimeMs: int ( o.CreatedTimeMs ),
    }
    return n
}

func PBConvPB_OldMsgPush_To_OldMsgPush ( o *OldMsgPush) *PB_OldMsgPush {
     n := &PB_OldMsgPush{
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
func PBConvPB__OldMsgPushEvent_To_OldMsgPushEvent( o *PB_OldMsgPushEvent) *OldMsgPushEvent {
     n := &OldMsgPushEvent{
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

func PBConvPB_OldMsgPushEvent_To_OldMsgPushEvent ( o *OldMsgPushEvent) *PB_OldMsgPushEvent {
     n := &PB_OldMsgPushEvent{
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
      Color: string ( o.Color ),
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
      Color: string ( o.Color ),
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
func PBConvPB__PushEvent_To_PushEvent( o *PB_PushEvent) *PushEvent {
     n := &PushEvent{
      PushEventId: int ( o.PushEventId ),
      ToUserId: int ( o.ToUserId ),
      ToDeviceId: int ( o.ToDeviceId ),
      MessageId: int ( o.MessageId ),
      RoomTypeEnum: int ( o.RoomTypeEnum ),
      RoomId: int ( o.RoomId ),
      PeerUserId: int ( o.PeerUserId ),
      PushEventTypeEnum: int ( o.PushEventTypeEnum ),
      AtTime: int ( o.AtTime ),
    }
    return n
}

func PBConvPB_PushEvent_To_PushEvent ( o *PushEvent) *PB_PushEvent {
     n := &PB_PushEvent{
      PushEventId: int64 ( o.PushEventId ),
      ToUserId: int32 ( o.ToUserId ),
      ToDeviceId: int64 ( o.ToDeviceId ),
      MessageId: int64 ( o.MessageId ),
      RoomTypeEnum: int32 ( o.RoomTypeEnum ),
      RoomId: int64 ( o.RoomId ),
      PeerUserId: int32 ( o.PeerUserId ),
      PushEventTypeEnum: int32 ( o.PushEventTypeEnum ),
      AtTime: int32 ( o.AtTime ),
    }
    return n
}
*/
/*
func PBConvPB__PushMessage_To_PushMessage( o *PB_PushMessage) *PushMessage {
     n := &PushMessage{
      PushMessageId: int ( o.PushMessageId ),
      ToUserId: int ( o.ToUserId ),
      ToDeviceId: int ( o.ToDeviceId ),
      MessageId: int ( o.MessageId ),
      RoomTypeEnum: int ( o.RoomTypeEnum ),
      CreatedMs: int ( o.CreatedMs ),
    }
    return n
}

func PBConvPB_PushMessage_To_PushMessage ( o *PushMessage) *PB_PushMessage {
     n := &PB_PushMessage{
      PushMessageId: int64 ( o.PushMessageId ),
      ToUserId: int32 ( o.ToUserId ),
      ToDeviceId: int64 ( o.ToDeviceId ),
      MessageId: int64 ( o.MessageId ),
      RoomTypeEnum: int32 ( o.RoomTypeEnum ),
      CreatedMs: int64 ( o.CreatedMs ),
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
func PBConvPB__Room_To_Room( o *PB_Room) *Room {
     n := &Room{
      RoomId: int ( o.RoomId ),
      RoomKey: string ( o.RoomKey ),
      RoomTypeEnum: int ( o.RoomTypeEnum ),
      UserId: int ( o.UserId ),
      LastSeqSeen: int ( o.LastSeqSeen ),
      LastSeqDelete: int ( o.LastSeqDelete ),
      PeerUserId: int ( o.PeerUserId ),
      GroupId: int ( o.GroupId ),
      CreatedTime: int ( o.CreatedTime ),
      CurrentSeq: int ( o.CurrentSeq ),
    }
    return n
}

func PBConvPB_Room_To_Room ( o *Room) *PB_Room {
     n := &PB_Room{
      RoomId: int64 ( o.RoomId ),
      RoomKey: string ( o.RoomKey ),
      RoomTypeEnum: int32 ( o.RoomTypeEnum ),
      UserId: int32 ( o.UserId ),
      LastSeqSeen: int32 ( o.LastSeqSeen ),
      LastSeqDelete: int32 ( o.LastSeqDelete ),
      PeerUserId: int32 ( o.PeerUserId ),
      GroupId: int64 ( o.GroupId ),
      CreatedTime: int32 ( o.CreatedTime ),
      CurrentSeq: int32 ( o.CurrentSeq ),
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
      LastSyncDirectUpdateId: int ( o.LastSyncDirectUpdateId ),
      LastSyncGeneralUpdateId: int ( o.LastSyncGeneralUpdateId ),
      LastSyncNotifyUpdateId: int ( o.LastSyncNotifyUpdateId ),
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
      LastSyncDirectUpdateId: int64 ( o.LastSyncDirectUpdateId ),
      LastSyncGeneralUpdateId: int64 ( o.LastSyncGeneralUpdateId ),
      LastSyncNotifyUpdateId: int64 ( o.LastSyncNotifyUpdateId ),
    }
    return n
}
*/
/*
func PBConvPB__SettingClient_To_SettingClient( o *PB_SettingClient) *SettingClient {
     n := &SettingClient{
      UserId: int ( o.UserId ),
      AutoDownloadWifiVoice: int ( o.AutoDownloadWifiVoice ),
      AutoDownloadWifiImage: int ( o.AutoDownloadWifiImage ),
      AutoDownloadWifiVideo: int ( o.AutoDownloadWifiVideo ),
      AutoDownloadWifiFile: int ( o.AutoDownloadWifiFile ),
      AutoDownloadWifiMusic: int ( o.AutoDownloadWifiMusic ),
      AutoDownloadWifiGif: int ( o.AutoDownloadWifiGif ),
      AutoDownloadCellularVoice: int ( o.AutoDownloadCellularVoice ),
      AutoDownloadCellularImage: int ( o.AutoDownloadCellularImage ),
      AutoDownloadCellularVideo: int ( o.AutoDownloadCellularVideo ),
      AutoDownloadCellularFile: int ( o.AutoDownloadCellularFile ),
      AutoDownloadCellularMusic: int ( o.AutoDownloadCellularMusic ),
      AutoDownloadCellularGif: int ( o.AutoDownloadCellularGif ),
      AutoDownloadRoamingVoice: int ( o.AutoDownloadRoamingVoice ),
      AutoDownloadRoamingImage: int ( o.AutoDownloadRoamingImage ),
      AutoDownloadRoamingVideo: int ( o.AutoDownloadRoamingVideo ),
      AutoDownloadRoamingFile: int ( o.AutoDownloadRoamingFile ),
      AutoDownloadRoamingMusic: int ( o.AutoDownloadRoamingMusic ),
      AutoDownloadRoamingGif: int ( o.AutoDownloadRoamingGif ),
      SaveToGallery: int ( o.SaveToGallery ),
    }
    return n
}

func PBConvPB_SettingClient_To_SettingClient ( o *SettingClient) *PB_SettingClient {
     n := &PB_SettingClient{
      UserId: int32 ( o.UserId ),
      AutoDownloadWifiVoice: int32 ( o.AutoDownloadWifiVoice ),
      AutoDownloadWifiImage: int32 ( o.AutoDownloadWifiImage ),
      AutoDownloadWifiVideo: int32 ( o.AutoDownloadWifiVideo ),
      AutoDownloadWifiFile: int32 ( o.AutoDownloadWifiFile ),
      AutoDownloadWifiMusic: int32 ( o.AutoDownloadWifiMusic ),
      AutoDownloadWifiGif: int32 ( o.AutoDownloadWifiGif ),
      AutoDownloadCellularVoice: int32 ( o.AutoDownloadCellularVoice ),
      AutoDownloadCellularImage: int32 ( o.AutoDownloadCellularImage ),
      AutoDownloadCellularVideo: int32 ( o.AutoDownloadCellularVideo ),
      AutoDownloadCellularFile: int32 ( o.AutoDownloadCellularFile ),
      AutoDownloadCellularMusic: int32 ( o.AutoDownloadCellularMusic ),
      AutoDownloadCellularGif: int32 ( o.AutoDownloadCellularGif ),
      AutoDownloadRoamingVoice: int32 ( o.AutoDownloadRoamingVoice ),
      AutoDownloadRoamingImage: int32 ( o.AutoDownloadRoamingImage ),
      AutoDownloadRoamingVideo: int32 ( o.AutoDownloadRoamingVideo ),
      AutoDownloadRoamingFile: int32 ( o.AutoDownloadRoamingFile ),
      AutoDownloadRoamingMusic: int32 ( o.AutoDownloadRoamingMusic ),
      AutoDownloadRoamingGif: int32 ( o.AutoDownloadRoamingGif ),
      SaveToGallery: int32 ( o.SaveToGallery ),
    }
    return n
}
*/
/*
func PBConvPB__SettingNotification_To_SettingNotification( o *PB_SettingNotification) *SettingNotification {
     n := &SettingNotification{
      UserId: int ( o.UserId ),
      SocialLedOn: int ( o.SocialLedOn ),
      SocialLedColor: string ( o.SocialLedColor ),
      ReqestToFollowYou: int ( o.ReqestToFollowYou ),
      FollowedYou: int ( o.FollowedYou ),
      AccptedYourFollowRequest: int ( o.AccptedYourFollowRequest ),
      YourPostLiked: int ( o.YourPostLiked ),
      YourPostCommented: int ( o.YourPostCommented ),
      MenthenedYouInPost: int ( o.MenthenedYouInPost ),
      MenthenedYouInComment: int ( o.MenthenedYouInComment ),
      YourContactsJoined: int ( o.YourContactsJoined ),
      DirectMessage: int ( o.DirectMessage ),
      DirectAlert: int ( o.DirectAlert ),
      DirectPerview: int ( o.DirectPerview ),
      DirectLedOn: int ( o.DirectLedOn ),
      DirectLedColor: int ( o.DirectLedColor ),
      DirectVibrate: int ( o.DirectVibrate ),
      DirectPopup: int ( o.DirectPopup ),
      DirectSound: int ( o.DirectSound ),
      DirectPriority: int ( o.DirectPriority ),
    }
    return n
}

func PBConvPB_SettingNotification_To_SettingNotification ( o *SettingNotification) *PB_SettingNotification {
     n := &PB_SettingNotification{
      UserId: int32 ( o.UserId ),
      SocialLedOn: int32 ( o.SocialLedOn ),
      SocialLedColor: string ( o.SocialLedColor ),
      ReqestToFollowYou: int32 ( o.ReqestToFollowYou ),
      FollowedYou: int32 ( o.FollowedYou ),
      AccptedYourFollowRequest: int32 ( o.AccptedYourFollowRequest ),
      YourPostLiked: int32 ( o.YourPostLiked ),
      YourPostCommented: int32 ( o.YourPostCommented ),
      MenthenedYouInPost: int32 ( o.MenthenedYouInPost ),
      MenthenedYouInComment: int32 ( o.MenthenedYouInComment ),
      YourContactsJoined: int32 ( o.YourContactsJoined ),
      DirectMessage: int32 ( o.DirectMessage ),
      DirectAlert: int32 ( o.DirectAlert ),
      DirectPerview: int32 ( o.DirectPerview ),
      DirectLedOn: int32 ( o.DirectLedOn ),
      DirectLedColor: int32 ( o.DirectLedColor ),
      DirectVibrate: int32 ( o.DirectVibrate ),
      DirectPopup: int32 ( o.DirectPopup ),
      DirectSound: int32 ( o.DirectSound ),
      DirectPriority: int32 ( o.DirectPriority ),
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
func PBConvPB__TestChat_To_TestChat( o *PB_TestChat) *TestChat {
     n := &TestChat{
      Id: int ( o.Id ),
      Id4: int ( o.Id4 ),
      TimeMs: int ( o.TimeMs ),
      Text: string ( o.Text ),
      Name: string ( o.Name ),
      UserId: int ( o.UserId ),
      C2: int ( o.C2 ),
      C3: int ( o.C3 ),
      C4: int ( o.C4 ),
      C5: int ( o.C5 ),
    }
    return n
}

func PBConvPB_TestChat_To_TestChat ( o *TestChat) *PB_TestChat {
     n := &PB_TestChat{
      Id: int64 ( o.Id ),
      Id4: int64 ( o.Id4 ),
      TimeMs: int64 ( o.TimeMs ),
      Text: string ( o.Text ),
      Name: string ( o.Name ),
      UserId: int64 ( o.UserId ),
      C2: int32 ( o.C2 ),
      C3: int32 ( o.C3 ),
      C4: int32 ( o.C4 ),
      C5: int32 ( o.C5 ),
    }
    return n
}
*/
/*
func PBConvPB__TriggerLog_To_TriggerLog( o *PB_TriggerLog) *TriggerLog {
     n := &TriggerLog{
      Id: int ( o.Id ),
      ModelName: string ( o.ModelName ),
      ChangeType: string ( o.ChangeType ),
      TargetId: int ( o.TargetId ),
      TargetStr: string ( o.TargetStr ),
      CreatedSe: int ( o.CreatedSe ),
    }
    return n
}

func PBConvPB_TriggerLog_To_TriggerLog ( o *TriggerLog) *PB_TriggerLog {
     n := &PB_TriggerLog{
      Id: int64 ( o.Id ),
      ModelName: string ( o.ModelName ),
      ChangeType: string ( o.ChangeType ),
      TargetId: int64 ( o.TargetId ),
      TargetStr: string ( o.TargetStr ),
      CreatedSe: int32 ( o.CreatedSe ),
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
