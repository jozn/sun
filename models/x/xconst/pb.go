package xconst

const (
    GeoLocation = "GeoLocation"
    RoomMessageLog = "RoomMessageLog"
    RoomMessageForwardFrom = "RoomMessageForwardFrom"
    RoomDraft = "RoomDraft"
    ChatRoom = "ChatRoom"
    Pagination = "Pagination"
    PB_CommandToServer = "PB_CommandToServer"
    PB_CommandToClient = "PB_CommandToClient"
    PB_CommandReachedToServer = "PB_CommandReachedToServer"
    PB_CommandReachedToClient = "PB_CommandReachedToClient"
    PB_ResponseToClient = "PB_ResponseToClient"
    PB_Offline_NewDirectMessage = "PB_Offline_NewDirectMessage"
    PB_Offline_MessagesReachedServer = "PB_Offline_MessagesReachedServer"
    PB_Offline_MessagesDeliveredToUser = "PB_Offline_MessagesDeliveredToUser"
    PB_Offline_MessagesSeenByPeer = "PB_Offline_MessagesSeenByPeer"
    PB_Offline_MessagesDeletedFromServer = "PB_Offline_MessagesDeletedFromServer"
    PB_Offline_ChangeMessageId = "PB_Offline_ChangeMessageId"
    PB_Offline_ChangeMessageFileId = "PB_Offline_ChangeMessageFileId"
    PB_Offline_MessageToEdit = "PB_Offline_MessageToEdit"
    PB_Offline_MessageToDelete = "PB_Offline_MessageToDelete"
    PB_Online_RoomActionDoing = "PB_Online_RoomActionDoing"
    PB_Offline_Messagings = "PB_Offline_Messagings"
    PB_UserParam_CheckUserName2 = "PB_UserParam_CheckUserName2"
    PB_UserResponse_CheckUserName2 = "PB_UserResponse_CheckUserName2"
    PB_ChatParam_AddNewMessage = "PB_ChatParam_AddNewMessage"
    PB_ChatResponse_AddNewMessage = "PB_ChatResponse_AddNewMessage"
    PB_ChatParam_SetRoomActionDoing = "PB_ChatParam_SetRoomActionDoing"
    PB_ChatResponse_SetRoomActionDoing = "PB_ChatResponse_SetRoomActionDoing"
    PB_ChatParam_SetChatMessagesRangeAsSeen = "PB_ChatParam_SetChatMessagesRangeAsSeen"
    PB_ChatResponse_SetChatMessagesRangeAsSeen = "PB_ChatResponse_SetChatMessagesRangeAsSeen"
    PB_ChatParam_DeleteChatHistory = "PB_ChatParam_DeleteChatHistory"
    PB_ChatResponse_DeleteChatHistory = "PB_ChatResponse_DeleteChatHistory"
    PB_ChatParam_DeleteMessagesByIds = "PB_ChatParam_DeleteMessagesByIds"
    PB_ChatResponse_DeleteMessagesByIds = "PB_ChatResponse_DeleteMessagesByIds"
    PB_ChatParam_SetMessagesAsReceived = "PB_ChatParam_SetMessagesAsReceived"
    PB_ChatResponse_SetMessagesAsReceived = "PB_ChatResponse_SetMessagesAsReceived"
    PB_ChatParam_EditMessage = "PB_ChatParam_EditMessage"
    PB_ChatResponse_EditMessage = "PB_ChatResponse_EditMessage"
    PB_ChatParam_GetChatList = "PB_ChatParam_GetChatList"
    PB_ChatResponse_GetChatList = "PB_ChatResponse_GetChatList"
    PB_ChatParam_GetChatHistoryToOlder = "PB_ChatParam_GetChatHistoryToOlder"
    PB_ChatResponse_GetChatHistoryToOlder = "PB_ChatResponse_GetChatHistoryToOlder"
    PB_ChatParam_GetFreshAllDirectMessagesList = "PB_ChatParam_GetFreshAllDirectMessagesList"
    PB_ChatResponse_GetFreshAllDirectMessagesList = "PB_ChatResponse_GetFreshAllDirectMessagesList"
    PB_MsgParam_AddNewTextMessage = "PB_MsgParam_AddNewTextMessage"
    PB_MsgResponse_AddNewTextMessage = "PB_MsgResponse_AddNewTextMessage"
    PB_MsgParam_AddNewMessage = "PB_MsgParam_AddNewMessage"
    PB_MsgResponse_AddNewMessage = "PB_MsgResponse_AddNewMessage"
    PB_MsgParam_SetRoomActionDoing = "PB_MsgParam_SetRoomActionDoing"
    PB_MsgResponse_SetRoomActionDoing = "PB_MsgResponse_SetRoomActionDoing"
    PB_MsgParam_GetMessagesByIds = "PB_MsgParam_GetMessagesByIds"
    PB_MsgResponse_GetMessagesByIds = "PB_MsgResponse_GetMessagesByIds"
    PB_MsgParam_GetMessagesHistory = "PB_MsgParam_GetMessagesHistory"
    PB_MsgResponse_GetMessagesHistory = "PB_MsgResponse_GetMessagesHistory"
    PB_MsgParam_SetChatMessagesRangeAsSeen = "PB_MsgParam_SetChatMessagesRangeAsSeen"
    PB_MsgResponse_SetChatMessagesRangeAsSeen = "PB_MsgResponse_SetChatMessagesRangeAsSeen"
    PB_MsgParam_DeleteChatHistory = "PB_MsgParam_DeleteChatHistory"
    PB_MsgResponse_DeleteChatHistory = "PB_MsgResponse_DeleteChatHistory"
    PB_MsgParam_DeleteMessagesByIds = "PB_MsgParam_DeleteMessagesByIds"
    PB_MsgResponse_DeleteMessagesByIds = "PB_MsgResponse_DeleteMessagesByIds"
    PB_MsgParam_SetMessagesAsReceived = "PB_MsgParam_SetMessagesAsReceived"
    PB_MsgResponse_SetMessagesAsReceived = "PB_MsgResponse_SetMessagesAsReceived"
    PB_MsgParam_ForwardMessages = "PB_MsgParam_ForwardMessages"
    PB_MsgResponse_ForwardMessages = "PB_MsgResponse_ForwardMessages"
    PB_MsgParam_EditMessage = "PB_MsgParam_EditMessage"
    PB_MsgResponse_EditMessage = "PB_MsgResponse_EditMessage"
    PB_MsgParam_BroadcastNewMessage = "PB_MsgParam_BroadcastNewMessage"
    PB_MsgResponse_BroadcastNewMessage = "PB_MsgResponse_BroadcastNewMessage"
    PB_MsgParam_GetFreshChatList = "PB_MsgParam_GetFreshChatList"
    PB_MsgResponse_GetFreshChatList = "PB_MsgResponse_GetFreshChatList"
    PB_MsgParam_GetFreshRoomMessagesList = "PB_MsgParam_GetFreshRoomMessagesList"
    PB_MsgResponse_GetFreshRoomMessagesList = "PB_MsgResponse_GetFreshRoomMessagesList"
    PB_MsgParam_GetFreshAllDirectMessagesList = "PB_MsgParam_GetFreshAllDirectMessagesList"
    PB_MsgResponse_GetFreshAllDirectMessagesList = "PB_MsgResponse_GetFreshAllDirectMessagesList"
    PB_MessageForwardedFrom = "PB_MessageForwardedFrom"
    PB_MessagesCollections = "PB_MessagesCollections"
    PB_MsgParam_Echo = "PB_MsgParam_Echo"
    PB_MsgResponse_PB_MsgParam_Echo = "PB_MsgResponse_PB_MsgParam_Echo"
    PB_SyncParam_GetDirectUpdates = "PB_SyncParam_GetDirectUpdates"
    PB_SyncResponse_GetDirectUpdates = "PB_SyncResponse_GetDirectUpdates"
    PB_SyncParam_GetGeneralUpdates = "PB_SyncParam_GetGeneralUpdates"
    PB_SyncResponse_GetGeneralUpdates = "PB_SyncResponse_GetGeneralUpdates"
    PB_SyncParam_GetNotifyUpdates = "PB_SyncParam_GetNotifyUpdates"
    PB_SyncResponse_GetNotifyUpdates = "PB_SyncResponse_GetNotifyUpdates"
    PB_SyncParam_SetLastSyncDirectUpdateId = "PB_SyncParam_SetLastSyncDirectUpdateId"
    PB_SyncResponse_SetLastSyncDirectUpdateId = "PB_SyncResponse_SetLastSyncDirectUpdateId"
    PB_SyncParam_SetLastSyncGeneralUpdateId = "PB_SyncParam_SetLastSyncGeneralUpdateId"
    PB_SyncResponse_SetLastSyncGeneralUpdateId = "PB_SyncResponse_SetLastSyncGeneralUpdateId"
    PB_SyncParam_SetLastSyncNotifyUpdateId = "PB_SyncParam_SetLastSyncNotifyUpdateId"
    PB_SyncResponse_SetLastSyncNotifyUpdateId = "PB_SyncResponse_SetLastSyncNotifyUpdateId"
    PB_NotifyUpdatesView = "PB_NotifyUpdatesView"
    PB_AllLivePushes = "PB_AllLivePushes"
    PB_UserParam_BlockUser = "PB_UserParam_BlockUser"
    PB_UserOfflineResponse_BlockUser = "PB_UserOfflineResponse_BlockUser"
    PB_UserParam_UnBlockUser = "PB_UserParam_UnBlockUser"
    PB_UserOfflineResponse_UnBlockUser = "PB_UserOfflineResponse_UnBlockUser"
    PB_UserParam_BlockedList = "PB_UserParam_BlockedList"
    PB_UserResponse_BlockedList = "PB_UserResponse_BlockedList"
    PB_UserParam_UpdateAbout = "PB_UserParam_UpdateAbout"
    PB_UserOfflineResponse_UpdateAbout = "PB_UserOfflineResponse_UpdateAbout"
    PB_UserParam_UpdateUserName = "PB_UserParam_UpdateUserName"
    PB_UserOfflineResponse_UpdateUserName = "PB_UserOfflineResponse_UpdateUserName"
    PB_UserParam_ChangeAvatar = "PB_UserParam_ChangeAvatar"
    PB_UserOfflineResponse_ChangeAvatar = "PB_UserOfflineResponse_ChangeAvatar"
    PB_UserParam_ChangePrivacy = "PB_UserParam_ChangePrivacy"
    PB_UserResponseOffline_ChangePrivacy = "PB_UserResponseOffline_ChangePrivacy"
    PB_UserParam_CheckUserName = "PB_UserParam_CheckUserName"
    PB_UserResponse_CheckUserName = "PB_UserResponse_CheckUserName"
    UserView = "UserView"
    PB_Activity = "PB_Activity"
    PB_Bucket = "PB_Bucket"
    PB_Chat = "PB_Chat"
    PB_Comment = "PB_Comment"
    PB_DirectMessage = "PB_DirectMessage"
    PB_DirectOffline = "PB_DirectOffline"
    PB_DirectToMessage = "PB_DirectToMessage"
    PB_DirectUpdate = "PB_DirectUpdate"
    PB_FollowingList = "PB_FollowingList"
    PB_FollowingListMember = "PB_FollowingListMember"
    PB_FollowingListMemberHistory = "PB_FollowingListMemberHistory"
    PB_GeneralLog = "PB_GeneralLog"
    PB_Group = "PB_Group"
    PB_GroupMember = "PB_GroupMember"
    PB_GroupMessage = "PB_GroupMessage"
    PB_GroupToMessage = "PB_GroupToMessage"
    PB_Like = "PB_Like"
    PB_LogChange = "PB_LogChange"
    PB_Media = "PB_Media"
    PB_MessageFile = "PB_MessageFile"
    PB_Notification = "PB_Notification"
    PB_NotificationRemoved = "PB_NotificationRemoved"
    PB_Offline = "PB_Offline"
    PB_OldMessage = "PB_OldMessage"
    PB_OldMsgFile = "PB_OldMsgFile"
    PB_OldMsgPush = "PB_OldMsgPush"
    PB_OldMsgPushEvent = "PB_OldMsgPushEvent"
    PB_PhoneContact = "PB_PhoneContact"
    PB_Photo = "PB_Photo"
    PB_Post = "PB_Post"
    PB_PushEvent = "PB_PushEvent"
    PB_PushMessage = "PB_PushMessage"
    PB_RecommendUser = "PB_RecommendUser"
    PB_Room = "PB_Room"
    PB_SearchClicked = "PB_SearchClicked"
    PB_Session = "PB_Session"
    PB_SettingClient = "PB_SettingClient"
    PB_SettingNotification = "PB_SettingNotification"
    PB_Tag = "PB_Tag"
    PB_TagsPost = "PB_TagsPost"
    PB_TestChat = "PB_TestChat"
    PB_TriggerLog = "PB_TriggerLog"
    PB_User = "PB_User"
    PB_UserMetaInfo = "PB_UserMetaInfo"
    PB_UserPassword = "PB_UserPassword"
    PB_UpdateGroupParticipants = "PB_UpdateGroupParticipants"
    PB_UpdateNotifySettings = "PB_UpdateNotifySettings"
    PB_UpdateServiceNotification = "PB_UpdateServiceNotification"
    PB_UpdateMessageMeta = "PB_UpdateMessageMeta"
    PB_UpdateMessageId = "PB_UpdateMessageId"
    PB_UpdateMessageToEdit = "PB_UpdateMessageToEdit"
    PB_UpdateMessageToDelete = "PB_UpdateMessageToDelete"
    PB_UpdateRoomActionDoing = "PB_UpdateRoomActionDoing"
    PB_UpdateUserBlocked = "PB_UpdateUserBlocked"
    PB_ChatView = "PB_ChatView"
    PB_MessageView = "PB_MessageView"
    PB_MessageFileView = "PB_MessageFileView"
    PB_UserView = "PB_UserView"
)
