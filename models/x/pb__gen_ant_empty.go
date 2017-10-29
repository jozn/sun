package x

/////////////// Empty Sample RPC - mainly for mocking ////////////////

/////////////////// RPC_Auth  -  EmptyRPC_RPC_Auth /////////////////////
type EmptyRPC_RPC_Auth int

var Empty_RPC_RPC_Auth_Sample = EmptyRPC_RPC_Auth(0)

func (EmptyRPC_RPC_Auth) CheckPhone(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SendCode(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SendCodeToSms(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SendCodeToTelgram(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SingUp(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SingIn(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) LogOut(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}

/////////////////// RPC_Chat  -  EmptyRPC_RPC_Chat /////////////////////
type EmptyRPC_RPC_Chat int

var Empty_RPC_RPC_Chat_Sample = EmptyRPC_RPC_Chat(0)

func (EmptyRPC_RPC_Chat) AddNewMessage(i *PB_ChatParam_AddNewMessage, p RPC_UserParam) (*PB_ChatResponse_AddNewMessage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) SetRoomActionDoing(i *PB_ChatParam_SetRoomActionDoing, p RPC_UserParam) (*PB_ChatResponse_SetRoomActionDoing, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) SetMessagesRangeAsSeen(i *PB_ChatParam_SetChatMessagesRangeAsSeen, p RPC_UserParam) (*PB_ChatResponse_SetChatMessagesRangeAsSeen, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) DeleteChatHistory(i *PB_ChatParam_DeleteChatHistory, p RPC_UserParam) (*PB_ChatResponse_DeleteChatHistory, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) DeleteMessagesByIds(i *PB_ChatParam_DeleteMessagesByIds, p RPC_UserParam) (*PB_ChatResponse_DeleteMessagesByIds, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) SetMessagesAsReceived(i *PB_ChatParam_SetMessagesAsReceived, p RPC_UserParam) (*PB_ChatResponse_SetMessagesAsReceived, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) EditMessage(i *PB_ChatParam_EditMessage, p RPC_UserParam) (*PB_ChatResponse_EditMessage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetChatList(i *PB_ChatParam_GetChatList, p RPC_UserParam) (*PB_ChatResponse_GetChatList, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetChatHistoryToOlder(i *PB_ChatParam_GetChatHistoryToOlder, p RPC_UserParam) (*PB_ChatResponse_GetChatHistoryToOlder, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetFreshAllDirectMessagesList(i *PB_ChatParam_GetFreshAllDirectMessagesList, p RPC_UserParam) (*PB_ChatResponse_GetFreshAllDirectMessagesList, error) {
	return nil, nil
}

/////////////////// RPC_Msg  -  EmptyRPC_RPC_Msg /////////////////////
type EmptyRPC_RPC_Msg int

var Empty_RPC_RPC_Msg_Sample = EmptyRPC_RPC_Msg(0)

func (EmptyRPC_RPC_Msg) AddNewTextMessage(i *PB_MsgParam_AddNewTextMessage, p RPC_UserParam) (*PB_MsgResponse_AddNewTextMessage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) AddNewMessage(i *PB_MsgParam_AddNewMessage, p RPC_UserParam) (*PB_MsgResponse_AddNewMessage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) SetRoomActionDoing(i *PB_MsgParam_SetRoomActionDoing, p RPC_UserParam) (*PB_MsgResponse_SetRoomActionDoing, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) GetMessagesByIds(i *PB_MsgParam_GetMessagesByIds, p RPC_UserParam) (*PB_MsgResponse_GetMessagesByIds, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) GetMessagesHistory(i *PB_MsgParam_GetMessagesHistory, p RPC_UserParam) (*PB_MsgResponse_GetMessagesHistory, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) SetMessagesRangeAsSeen(i *PB_MsgParam_SetChatMessagesRangeAsSeen, p RPC_UserParam) (*PB_MsgResponse_SetChatMessagesRangeAsSeen, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) DeleteChatHistory(i *PB_MsgParam_DeleteChatHistory, p RPC_UserParam) (*PB_MsgResponse_DeleteChatHistory, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) DeleteMessagesByIds(i *PB_MsgParam_DeleteMessagesByIds, p RPC_UserParam) (*PB_MsgResponse_DeleteMessagesByIds, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) SetMessagesAsReceived(i *PB_MsgParam_SetMessagesAsReceived, p RPC_UserParam) (*PB_MsgResponse_SetMessagesAsReceived, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) ForwardMessages(i *PB_MsgParam_ForwardMessages, p RPC_UserParam) (*PB_MsgResponse_ForwardMessages, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) EditMessage(i *PB_MsgParam_EditMessage, p RPC_UserParam) (*PB_MsgResponse_EditMessage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) BroadcastNewMessage(i *PB_MsgParam_BroadcastNewMessage, p RPC_UserParam) (*PB_MsgResponse_BroadcastNewMessage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) GetFreshChatList(i *PB_MsgParam_GetFreshChatList, p RPC_UserParam) (*PB_MsgResponse_GetFreshChatList, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) GetFreshRoomMessagesList(i *PB_MsgParam_GetFreshRoomMessagesList, p RPC_UserParam) (*PB_MsgResponse_GetFreshRoomMessagesList, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) GetFreshAllDirectMessagesList(i *PB_MsgParam_GetFreshAllDirectMessagesList, p RPC_UserParam) (*PB_MsgResponse_GetFreshAllDirectMessagesList, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Msg) Echo(i *PB_MsgParam_Echo, p RPC_UserParam) (*PB_MsgResponse_PB_MsgParam_Echo, error) {
	return nil, nil
}

/////////////////// RPC_Sync  -  EmptyRPC_RPC_Sync /////////////////////
type EmptyRPC_RPC_Sync int

var Empty_RPC_RPC_Sync_Sample = EmptyRPC_RPC_Sync(0)

func (EmptyRPC_RPC_Sync) GetGeneralUpdates(i *PB_SyncParam_GetGeneralUpdates, p RPC_UserParam) (*PB_SyncResponse_GetGeneralUpdates, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Sync) GetNotifyUpdates(i *PB_SyncParam_GetNotifyUpdates, p RPC_UserParam) (*PB_SyncResponse_GetNotifyUpdates, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Sync) SetLastSyncDirectUpdateId(i *PB_SyncParam_SetLastSyncDirectUpdateId, p RPC_UserParam) (*PB_SyncResponse_SetLastSyncDirectUpdateId, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Sync) SetLastSyncGeneralUpdateId(i *PB_SyncParam_SetLastSyncGeneralUpdateId, p RPC_UserParam) (*PB_SyncResponse_SetLastSyncGeneralUpdateId, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Sync) SetLastSyncNotifyUpdateId(i *PB_SyncParam_SetLastSyncNotifyUpdateId, p RPC_UserParam) (*PB_SyncResponse_SetLastSyncNotifyUpdateId, error) {
	return nil, nil
}

/////////////////// RPC_UserOffline  -  EmptyRPC_RPC_UserOffline /////////////////////
type EmptyRPC_RPC_UserOffline int

var Empty_RPC_RPC_UserOffline_Sample = EmptyRPC_RPC_UserOffline(0)

func (EmptyRPC_RPC_UserOffline) BlockUser(i *PB_UserParam_BlockUser, p RPC_UserParam) (*PB_UserOfflineResponse_BlockUser, error) {
	return nil, nil
}
func (EmptyRPC_RPC_UserOffline) UnBlockUser(i *PB_UserParam_UnBlockUser, p RPC_UserParam) (*PB_UserOfflineResponse_UnBlockUser, error) {
	return nil, nil
}
func (EmptyRPC_RPC_UserOffline) UpdateAbout(i *PB_UserParam_UpdateAbout, p RPC_UserParam) (*PB_UserOfflineResponse_UpdateAbout, error) {
	return nil, nil
}
func (EmptyRPC_RPC_UserOffline) UpdateUserName(i *PB_UserParam_UpdateUserName, p RPC_UserParam) (*PB_UserOfflineResponse_UpdateUserName, error) {
	return nil, nil
}
func (EmptyRPC_RPC_UserOffline) ChangePrivacy(i *PB_UserParam_ChangePrivacy, p RPC_UserParam) (*PB_UserResponseOffline_ChangePrivacy, error) {
	return nil, nil
}
func (EmptyRPC_RPC_UserOffline) ChangeAvatar(i *PB_UserParam_ChangeAvatar, p RPC_UserParam) (*PB_UserOfflineResponse_ChangeAvatar, error) {
	return nil, nil
}

/////////////////// RPC_User  -  EmptyRPC_RPC_User /////////////////////
type EmptyRPC_RPC_User int

var Empty_RPC_RPC_User_Sample = EmptyRPC_RPC_User(0)

func (EmptyRPC_RPC_User) CheckUserName(i *PB_UserParam_CheckUserName, p RPC_UserParam) (*PB_UserResponse_CheckUserName, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) GetBlockedList(i *PB_UserParam_BlockedList, p RPC_UserParam) (*PB_UserResponse_BlockedList, error) {
	return nil, nil
}
