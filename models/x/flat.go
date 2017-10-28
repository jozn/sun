package x

import "ms/sun/helper"


type GeoLocation_Flat struct {
    Lat float64
    Lon float64
}
//ToPB
func(m *GeoLocation)ToFlat() *GeoLocation_Flat {
r := &GeoLocation_Flat{
    Lat:  float64(m.Lat) ,
    Lon:  float64(m.Lon) ,
}
return r
}
//ToPB
func(m *GeoLocation_Flat)ToPB() *GeoLocation {
r := &GeoLocation{
    Lat:  m.Lat ,
    Lon:  m.Lon ,
}
return r
}
//folding
var GeoLocation__FOlD = &GeoLocation{
    Lat:  0.0 ,
    Lon:  0.0 ,
}


type RoomMessageLog_Flat struct {
    typ RoomMessageLogEnum
    TargetUserId int
    ByUserId int
}
//ToPB
func(m *RoomMessageLog)ToFlat() *RoomMessageLog_Flat {
r := &RoomMessageLog_Flat{
    
    TargetUserId:  int(m.TargetUserId) ,
    ByUserId:  int(m.ByUserId) ,
}
return r
}
//ToPB
func(m *RoomMessageLog_Flat)ToPB() *RoomMessageLog {
r := &RoomMessageLog{
    
    TargetUserId:  uint64(m.TargetUserId) ,
    ByUserId:  uint64(m.ByUserId) ,
}
return r
}
//folding
var RoomMessageLog__FOlD = &RoomMessageLog{
    
    TargetUserId:  0 ,
    ByUserId:  0 ,
}


type RoomMessageForwardFrom_Flat struct {
    RoomId int
    MessageId int
    RoomTypeEnum int
}
//ToPB
func(m *RoomMessageForwardFrom)ToFlat() *RoomMessageForwardFrom_Flat {
r := &RoomMessageForwardFrom_Flat{
    RoomId:  int(m.RoomId) ,
    MessageId:  int(m.MessageId) ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
}
return r
}
//ToPB
func(m *RoomMessageForwardFrom_Flat)ToPB() *RoomMessageForwardFrom {
r := &RoomMessageForwardFrom{
    RoomId:  uint64(m.RoomId) ,
    MessageId:  uint64(m.MessageId) ,
    RoomTypeEnum:  uint32(m.RoomTypeEnum) ,
}
return r
}
//folding
var RoomMessageForwardFrom__FOlD = &RoomMessageForwardFrom{
    RoomId:  0 ,
    MessageId:  0 ,
    RoomTypeEnum:  0 ,
}


type RoomDraft_Flat struct {
    Message string
    ReplyTo int
}
//ToPB
func(m *RoomDraft)ToFlat() *RoomDraft_Flat {
r := &RoomDraft_Flat{
    Message:  m.Message ,
    ReplyTo:  int(m.ReplyTo) ,
}
return r
}
//ToPB
func(m *RoomDraft_Flat)ToPB() *RoomDraft {
r := &RoomDraft{
    Message:  m.Message ,
    ReplyTo:  uint64(m.ReplyTo) ,
}
return r
}
//folding
var RoomDraft__FOlD = &RoomDraft{
    Message:  "" ,
    ReplyTo:  0 ,
}


type ChatRoom_Flat struct {
}
//ToPB
func(m *ChatRoom)ToFlat() *ChatRoom_Flat {
r := &ChatRoom_Flat{
}
return r
}
//ToPB
func(m *ChatRoom_Flat)ToPB() *ChatRoom {
r := &ChatRoom{
}
return r
}
//folding
var ChatRoom__FOlD = &ChatRoom{
}


type Pagination_Flat struct {
    Offset int
    Limit int
}
//ToPB
func(m *Pagination)ToFlat() *Pagination_Flat {
r := &Pagination_Flat{
    Offset:  int(m.Offset) ,
    Limit:  int(m.Limit) ,
}
return r
}
//ToPB
func(m *Pagination_Flat)ToPB() *Pagination {
r := &Pagination{
    Offset:  uint32(m.Offset) ,
    Limit:  uint32(m.Limit) ,
}
return r
}
//folding
var Pagination__FOlD = &Pagination{
    Offset:  0 ,
    Limit:  0 ,
}


type PB_CommandToServer_Flat struct {
    ClientCallId int
    Command string
    RespondReached bool
    Data []byte
}
//ToPB
func(m *PB_CommandToServer)ToFlat() *PB_CommandToServer_Flat {
r := &PB_CommandToServer_Flat{
    ClientCallId:  int(m.ClientCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  []byte(m.Data) ,
}
return r
}
//ToPB
func(m *PB_CommandToServer_Flat)ToPB() *PB_CommandToServer {
r := &PB_CommandToServer{
    ClientCallId:  int64(m.ClientCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  m.Data ,
}
return r
}
//folding
var PB_CommandToServer__FOlD = &PB_CommandToServer{
    ClientCallId:  0 ,
    Command:  "" ,
    RespondReached:  false ,
    Data:  []byte{} ,
}


type PB_CommandToClient_Flat struct {
    ServerCallId int
    Command string
    RespondReached bool
    Data []byte
}
//ToPB
func(m *PB_CommandToClient)ToFlat() *PB_CommandToClient_Flat {
r := &PB_CommandToClient_Flat{
    ServerCallId:  int(m.ServerCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  []byte(m.Data) ,
}
return r
}
//ToPB
func(m *PB_CommandToClient_Flat)ToPB() *PB_CommandToClient {
r := &PB_CommandToClient{
    ServerCallId:  int64(m.ServerCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  m.Data ,
}
return r
}
//folding
var PB_CommandToClient__FOlD = &PB_CommandToClient{
    ServerCallId:  0 ,
    Command:  "" ,
    RespondReached:  false ,
    Data:  []byte{} ,
}


type PB_CommandReachedToServer_Flat struct {
    ClientCallId int
}
//ToPB
func(m *PB_CommandReachedToServer)ToFlat() *PB_CommandReachedToServer_Flat {
r := &PB_CommandReachedToServer_Flat{
    ClientCallId:  int(m.ClientCallId) ,
}
return r
}
//ToPB
func(m *PB_CommandReachedToServer_Flat)ToPB() *PB_CommandReachedToServer {
r := &PB_CommandReachedToServer{
    ClientCallId:  int64(m.ClientCallId) ,
}
return r
}
//folding
var PB_CommandReachedToServer__FOlD = &PB_CommandReachedToServer{
    ClientCallId:  0 ,
}


type PB_CommandReachedToClient_Flat struct {
    ServerCallId int
}
//ToPB
func(m *PB_CommandReachedToClient)ToFlat() *PB_CommandReachedToClient_Flat {
r := &PB_CommandReachedToClient_Flat{
    ServerCallId:  int(m.ServerCallId) ,
}
return r
}
//ToPB
func(m *PB_CommandReachedToClient_Flat)ToPB() *PB_CommandReachedToClient {
r := &PB_CommandReachedToClient{
    ServerCallId:  int64(m.ServerCallId) ,
}
return r
}
//folding
var PB_CommandReachedToClient__FOlD = &PB_CommandReachedToClient{
    ServerCallId:  0 ,
}


type PB_ResponseToClient_Flat struct {
    ClientCallId int
    PBClass string
    RpcFullName string
    Data []byte
}
//ToPB
func(m *PB_ResponseToClient)ToFlat() *PB_ResponseToClient_Flat {
r := &PB_ResponseToClient_Flat{
    ClientCallId:  int(m.ClientCallId) ,
    PBClass:  m.PBClass ,
    RpcFullName:  m.RpcFullName ,
    Data:  []byte(m.Data) ,
}
return r
}
//ToPB
func(m *PB_ResponseToClient_Flat)ToPB() *PB_ResponseToClient {
r := &PB_ResponseToClient{
    ClientCallId:  int64(m.ClientCallId) ,
    PBClass:  m.PBClass ,
    RpcFullName:  m.RpcFullName ,
    Data:  m.Data ,
}
return r
}
//folding
var PB_ResponseToClient__FOlD = &PB_ResponseToClient{
    ClientCallId:  0 ,
    PBClass:  "" ,
    RpcFullName:  "" ,
    Data:  []byte{} ,
}


type PB_Offline_NewDirectMessage_Flat struct {
    ChatKey string
    FromMessageId int
    AtTime int
}
//ToPB
func(m *PB_Offline_NewDirectMessage)ToFlat() *PB_Offline_NewDirectMessage_Flat {
r := &PB_Offline_NewDirectMessage_Flat{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int(m.FromMessageId) ,
    AtTime:  int(m.AtTime) ,
}
return r
}
//ToPB
func(m *PB_Offline_NewDirectMessage_Flat)ToPB() *PB_Offline_NewDirectMessage {
r := &PB_Offline_NewDirectMessage{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int64(m.FromMessageId) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}
//folding
var PB_Offline_NewDirectMessage__FOlD = &PB_Offline_NewDirectMessage{
    ChatKey:  "" ,
    FromMessageId:  0 ,
    AtTime:  0 ,
}


type PB_Offline_MessagesReachedServer_Flat struct {
    MessageKeys []string
    AtTime int
}
//ToPB
func(m *PB_Offline_MessagesReachedServer)ToFlat() *PB_Offline_MessagesReachedServer_Flat {
r := &PB_Offline_MessagesReachedServer_Flat{
    MessageKeys:  m.MessageKeys ,
    AtTime:  int(m.AtTime) ,
}
return r
}
//ToPB
func(m *PB_Offline_MessagesReachedServer_Flat)ToPB() *PB_Offline_MessagesReachedServer {
r := &PB_Offline_MessagesReachedServer{
    MessageKeys:  m.MessageKeys ,
    AtTime:  int64(m.AtTime) ,
}
return r
}
//folding
var PB_Offline_MessagesReachedServer__FOlD = &PB_Offline_MessagesReachedServer{
    
    AtTime:  0 ,
}


type PB_Offline_MessagesDeliveredToUser_Flat struct {
    MessageKeys []string
    RoomKey string
    AtTime int
}
//ToPB
func(m *PB_Offline_MessagesDeliveredToUser)ToFlat() *PB_Offline_MessagesDeliveredToUser_Flat {
r := &PB_Offline_MessagesDeliveredToUser_Flat{
    MessageKeys:  m.MessageKeys ,
    RoomKey:  m.RoomKey ,
    AtTime:  int(m.AtTime) ,
}
return r
}
//ToPB
func(m *PB_Offline_MessagesDeliveredToUser_Flat)ToPB() *PB_Offline_MessagesDeliveredToUser {
r := &PB_Offline_MessagesDeliveredToUser{
    MessageKeys:  m.MessageKeys ,
    RoomKey:  m.RoomKey ,
    AtTime:  int64(m.AtTime) ,
}
return r
}
//folding
var PB_Offline_MessagesDeliveredToUser__FOlD = &PB_Offline_MessagesDeliveredToUser{
    
    RoomKey:  "" ,
    AtTime:  0 ,
}


type PB_Offline_MessagesSeenByPeer_Flat struct {
    MessageKeys []string
    RoomKey string
    AtTime int
}
//ToPB
func(m *PB_Offline_MessagesSeenByPeer)ToFlat() *PB_Offline_MessagesSeenByPeer_Flat {
r := &PB_Offline_MessagesSeenByPeer_Flat{
    MessageKeys:  m.MessageKeys ,
    RoomKey:  m.RoomKey ,
    AtTime:  int(m.AtTime) ,
}
return r
}
//ToPB
func(m *PB_Offline_MessagesSeenByPeer_Flat)ToPB() *PB_Offline_MessagesSeenByPeer {
r := &PB_Offline_MessagesSeenByPeer{
    MessageKeys:  m.MessageKeys ,
    RoomKey:  m.RoomKey ,
    AtTime:  int64(m.AtTime) ,
}
return r
}
//folding
var PB_Offline_MessagesSeenByPeer__FOlD = &PB_Offline_MessagesSeenByPeer{
    
    RoomKey:  "" ,
    AtTime:  0 ,
}


type PB_Offline_MessagesDeletedFromServer_Flat struct {
    MessageKeys []string
    AtTime int
}
//ToPB
func(m *PB_Offline_MessagesDeletedFromServer)ToFlat() *PB_Offline_MessagesDeletedFromServer_Flat {
r := &PB_Offline_MessagesDeletedFromServer_Flat{
    MessageKeys:  m.MessageKeys ,
    AtTime:  int(m.AtTime) ,
}
return r
}
//ToPB
func(m *PB_Offline_MessagesDeletedFromServer_Flat)ToPB() *PB_Offline_MessagesDeletedFromServer {
r := &PB_Offline_MessagesDeletedFromServer{
    MessageKeys:  m.MessageKeys ,
    AtTime:  int64(m.AtTime) ,
}
return r
}
//folding
var PB_Offline_MessagesDeletedFromServer__FOlD = &PB_Offline_MessagesDeletedFromServer{
    
    AtTime:  0 ,
}


type PB_Offline_ChangeMessageId_Flat struct {
    MessageKey int
    NewMessageId int
}
//ToPB
func(m *PB_Offline_ChangeMessageId)ToFlat() *PB_Offline_ChangeMessageId_Flat {
r := &PB_Offline_ChangeMessageId_Flat{
    MessageKey:  int(m.MessageKey) ,
    NewMessageId:  int(m.NewMessageId) ,
}
return r
}
//ToPB
func(m *PB_Offline_ChangeMessageId_Flat)ToPB() *PB_Offline_ChangeMessageId {
r := &PB_Offline_ChangeMessageId{
    MessageKey:  int64(m.MessageKey) ,
    NewMessageId:  int64(m.NewMessageId) ,
}
return r
}
//folding
var PB_Offline_ChangeMessageId__FOlD = &PB_Offline_ChangeMessageId{
    MessageKey:  0 ,
    NewMessageId:  0 ,
}


type PB_Offline_ChangeMessageFileId_Flat struct {
    MessageFileKey int
    NewMessageFileId int
}
//ToPB
func(m *PB_Offline_ChangeMessageFileId)ToFlat() *PB_Offline_ChangeMessageFileId_Flat {
r := &PB_Offline_ChangeMessageFileId_Flat{
    MessageFileKey:  int(m.MessageFileKey) ,
    NewMessageFileId:  int(m.NewMessageFileId) ,
}
return r
}
//ToPB
func(m *PB_Offline_ChangeMessageFileId_Flat)ToPB() *PB_Offline_ChangeMessageFileId {
r := &PB_Offline_ChangeMessageFileId{
    MessageFileKey:  int64(m.MessageFileKey) ,
    NewMessageFileId:  int64(m.NewMessageFileId) ,
}
return r
}
//folding
var PB_Offline_ChangeMessageFileId__FOlD = &PB_Offline_ChangeMessageFileId{
    MessageFileKey:  0 ,
    NewMessageFileId:  0 ,
}


type PB_Offline_MessageToEdit_Flat struct {
    MessageKey int
    NewText string
}
//ToPB
func(m *PB_Offline_MessageToEdit)ToFlat() *PB_Offline_MessageToEdit_Flat {
r := &PB_Offline_MessageToEdit_Flat{
    MessageKey:  int(m.MessageKey) ,
    NewText:  m.NewText ,
}
return r
}
//ToPB
func(m *PB_Offline_MessageToEdit_Flat)ToPB() *PB_Offline_MessageToEdit {
r := &PB_Offline_MessageToEdit{
    MessageKey:  int64(m.MessageKey) ,
    NewText:  m.NewText ,
}
return r
}
//folding
var PB_Offline_MessageToEdit__FOlD = &PB_Offline_MessageToEdit{
    MessageKey:  0 ,
    NewText:  "" ,
}


type PB_Offline_MessageToDelete_Flat struct {
    MessageKey int
}
//ToPB
func(m *PB_Offline_MessageToDelete)ToFlat() *PB_Offline_MessageToDelete_Flat {
r := &PB_Offline_MessageToDelete_Flat{
    MessageKey:  int(m.MessageKey) ,
}
return r
}
//ToPB
func(m *PB_Offline_MessageToDelete_Flat)ToPB() *PB_Offline_MessageToDelete {
r := &PB_Offline_MessageToDelete{
    MessageKey:  int64(m.MessageKey) ,
}
return r
}
//folding
var PB_Offline_MessageToDelete__FOlD = &PB_Offline_MessageToDelete{
    MessageKey:  0 ,
}


type PB_Online_RoomActionDoing_Flat struct {
    RoomKey string
    ActionType RoomActionDoingEnum
}
//ToPB
func(m *PB_Online_RoomActionDoing)ToFlat() *PB_Online_RoomActionDoing_Flat {
r := &PB_Online_RoomActionDoing_Flat{
    RoomKey:  m.RoomKey ,
    
}
return r
}
//ToPB
func(m *PB_Online_RoomActionDoing_Flat)ToPB() *PB_Online_RoomActionDoing {
r := &PB_Online_RoomActionDoing{
    RoomKey:  m.RoomKey ,
    
}
return r
}
//folding
var PB_Online_RoomActionDoing__FOlD = &PB_Online_RoomActionDoing{
    RoomKey:  "" ,
    
}


type PB_Offline_Messagings_Flat struct {
    NewMessages []PB_MessageView
    ChatFiles []PB_MessageFileView
    Chats []PB_ChatView
    Users []PB_UserView
    ChangeMessageIds []PB_Offline_ChangeMessageId
    ChangeMessageFileIds []PB_Offline_ChangeMessageFileId
    MessagesToEdit []PB_Offline_MessageToEdit
    MessagesToDelete []PB_Offline_MessageToDelete
    MessagesReachedServer []PB_Offline_MessagesReachedServer
    MessagesDeliveredToUser []PB_Offline_MessagesDeliveredToUser
    MessagesSeenByPeer []PB_Offline_MessagesSeenByPeer
    MessagesDeletedFromServer []PB_Offline_MessagesDeletedFromServer
    RoomActionDoing []PB_Online_RoomActionDoing
    LastId int
}
//ToPB
func(m *PB_Offline_Messagings)ToFlat() *PB_Offline_Messagings_Flat {
r := &PB_Offline_Messagings_Flat{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_Offline_Messagings_Flat)ToPB() *PB_Offline_Messagings {
r := &PB_Offline_Messagings{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_Offline_Messagings__FOlD = &PB_Offline_Messagings{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  0 ,
}


type PB_UserParam_CheckUserName2_Flat struct {
}
//ToPB
func(m *PB_UserParam_CheckUserName2)ToFlat() *PB_UserParam_CheckUserName2_Flat {
r := &PB_UserParam_CheckUserName2_Flat{
}
return r
}
//ToPB
func(m *PB_UserParam_CheckUserName2_Flat)ToPB() *PB_UserParam_CheckUserName2 {
r := &PB_UserParam_CheckUserName2{
}
return r
}
//folding
var PB_UserParam_CheckUserName2__FOlD = &PB_UserParam_CheckUserName2{
}


type PB_UserResponse_CheckUserName2_Flat struct {
}
//ToPB
func(m *PB_UserResponse_CheckUserName2)ToFlat() *PB_UserResponse_CheckUserName2_Flat {
r := &PB_UserResponse_CheckUserName2_Flat{
}
return r
}
//ToPB
func(m *PB_UserResponse_CheckUserName2_Flat)ToPB() *PB_UserResponse_CheckUserName2 {
r := &PB_UserResponse_CheckUserName2{
}
return r
}
//folding
var PB_UserResponse_CheckUserName2__FOlD = &PB_UserResponse_CheckUserName2{
}


type PB_ChatParam_AddNewMessage_Flat struct {
    ReplyToMessageId int
    Blob []byte
    MessageView PB_MessageView
}
//ToPB
func(m *PB_ChatParam_AddNewMessage)ToFlat() *PB_ChatParam_AddNewMessage_Flat {
r := &PB_ChatParam_AddNewMessage_Flat{
    ReplyToMessageId:  int(m.ReplyToMessageId) ,
    Blob:  []byte(m.Blob) ,
    
}
return r
}
//ToPB
func(m *PB_ChatParam_AddNewMessage_Flat)ToPB() *PB_ChatParam_AddNewMessage {
r := &PB_ChatParam_AddNewMessage{
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,
    Blob:  m.Blob ,
    
}
return r
}
//folding
var PB_ChatParam_AddNewMessage__FOlD = &PB_ChatParam_AddNewMessage{
    ReplyToMessageId:  0 ,
    Blob:  []byte{} ,
    
}


type PB_ChatResponse_AddNewMessage_Flat struct {
}
//ToPB
func(m *PB_ChatResponse_AddNewMessage)ToFlat() *PB_ChatResponse_AddNewMessage_Flat {
r := &PB_ChatResponse_AddNewMessage_Flat{
}
return r
}
//ToPB
func(m *PB_ChatResponse_AddNewMessage_Flat)ToPB() *PB_ChatResponse_AddNewMessage {
r := &PB_ChatResponse_AddNewMessage{
}
return r
}
//folding
var PB_ChatResponse_AddNewMessage__FOlD = &PB_ChatResponse_AddNewMessage{
}


type PB_ChatParam_SetRoomActionDoing_Flat struct {
    GroupId int
    DirectRoomKey string
    ActionType RoomActionDoingEnum
}
//ToPB
func(m *PB_ChatParam_SetRoomActionDoing)ToFlat() *PB_ChatParam_SetRoomActionDoing_Flat {
r := &PB_ChatParam_SetRoomActionDoing_Flat{
    GroupId:  int(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,
    
}
return r
}
//ToPB
func(m *PB_ChatParam_SetRoomActionDoing_Flat)ToPB() *PB_ChatParam_SetRoomActionDoing {
r := &PB_ChatParam_SetRoomActionDoing{
    GroupId:  int64(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,
    
}
return r
}
//folding
var PB_ChatParam_SetRoomActionDoing__FOlD = &PB_ChatParam_SetRoomActionDoing{
    GroupId:  0 ,
    DirectRoomKey:  "" ,
    
}


type PB_ChatResponse_SetRoomActionDoing_Flat struct {
}
//ToPB
func(m *PB_ChatResponse_SetRoomActionDoing)ToFlat() *PB_ChatResponse_SetRoomActionDoing_Flat {
r := &PB_ChatResponse_SetRoomActionDoing_Flat{
}
return r
}
//ToPB
func(m *PB_ChatResponse_SetRoomActionDoing_Flat)ToPB() *PB_ChatResponse_SetRoomActionDoing {
r := &PB_ChatResponse_SetRoomActionDoing{
}
return r
}
//folding
var PB_ChatResponse_SetRoomActionDoing__FOlD = &PB_ChatResponse_SetRoomActionDoing{
}


type PB_ChatParam_SetChatMessagesRangeAsSeen_Flat struct {
    ChatKey string
    BottomMessageId int
    TopMessageId int
    SeenTimeMs int
}
//ToPB
func(m *PB_ChatParam_SetChatMessagesRangeAsSeen)ToFlat() *PB_ChatParam_SetChatMessagesRangeAsSeen_Flat {
r := &PB_ChatParam_SetChatMessagesRangeAsSeen_Flat{
    ChatKey:  m.ChatKey ,
    BottomMessageId:  int(m.BottomMessageId) ,
    TopMessageId:  int(m.TopMessageId) ,
    SeenTimeMs:  int(m.SeenTimeMs) ,
}
return r
}
//ToPB
func(m *PB_ChatParam_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_ChatParam_SetChatMessagesRangeAsSeen {
r := &PB_ChatParam_SetChatMessagesRangeAsSeen{
    ChatKey:  m.ChatKey ,
    BottomMessageId:  int64(m.BottomMessageId) ,
    TopMessageId:  int64(m.TopMessageId) ,
    SeenTimeMs:  int64(m.SeenTimeMs) ,
}
return r
}
//folding
var PB_ChatParam_SetChatMessagesRangeAsSeen__FOlD = &PB_ChatParam_SetChatMessagesRangeAsSeen{
    ChatKey:  "" ,
    BottomMessageId:  0 ,
    TopMessageId:  0 ,
    SeenTimeMs:  0 ,
}


type PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat struct {
}
//ToPB
func(m *PB_ChatResponse_SetChatMessagesRangeAsSeen)ToFlat() *PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat {
r := &PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat{
}
return r
}
//ToPB
func(m *PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_ChatResponse_SetChatMessagesRangeAsSeen {
r := &PB_ChatResponse_SetChatMessagesRangeAsSeen{
}
return r
}
//folding
var PB_ChatResponse_SetChatMessagesRangeAsSeen__FOlD = &PB_ChatResponse_SetChatMessagesRangeAsSeen{
}


type PB_ChatParam_DeleteChatHistory_Flat struct {
    ChatKey string
    FromMessageId int
}
//ToPB
func(m *PB_ChatParam_DeleteChatHistory)ToFlat() *PB_ChatParam_DeleteChatHistory_Flat {
r := &PB_ChatParam_DeleteChatHistory_Flat{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int(m.FromMessageId) ,
}
return r
}
//ToPB
func(m *PB_ChatParam_DeleteChatHistory_Flat)ToPB() *PB_ChatParam_DeleteChatHistory {
r := &PB_ChatParam_DeleteChatHistory{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int64(m.FromMessageId) ,
}
return r
}
//folding
var PB_ChatParam_DeleteChatHistory__FOlD = &PB_ChatParam_DeleteChatHistory{
    ChatKey:  "" ,
    FromMessageId:  0 ,
}


type PB_ChatResponse_DeleteChatHistory_Flat struct {
}
//ToPB
func(m *PB_ChatResponse_DeleteChatHistory)ToFlat() *PB_ChatResponse_DeleteChatHistory_Flat {
r := &PB_ChatResponse_DeleteChatHistory_Flat{
}
return r
}
//ToPB
func(m *PB_ChatResponse_DeleteChatHistory_Flat)ToPB() *PB_ChatResponse_DeleteChatHistory {
r := &PB_ChatResponse_DeleteChatHistory{
}
return r
}
//folding
var PB_ChatResponse_DeleteChatHistory__FOlD = &PB_ChatResponse_DeleteChatHistory{
}


type PB_ChatParam_DeleteMessagesByIds_Flat struct {
    ChatKey string
    Both bool
    MessagesIds []int
}
//ToPB
func(m *PB_ChatParam_DeleteMessagesByIds)ToFlat() *PB_ChatParam_DeleteMessagesByIds_Flat {
r := &PB_ChatParam_DeleteMessagesByIds_Flat{
    ChatKey:  m.ChatKey ,
    Both:  m.Both ,
    MessagesIds:  helper.SliceInt64ToInt(m.MessagesIds) ,
}
return r
}
//ToPB
func(m *PB_ChatParam_DeleteMessagesByIds_Flat)ToPB() *PB_ChatParam_DeleteMessagesByIds {
r := &PB_ChatParam_DeleteMessagesByIds{
    ChatKey:  m.ChatKey ,
    Both:  m.Both ,
    MessagesIds:  helper.SliceIntToInt64(m.MessagesIds) ,
}
return r
}
//folding
var PB_ChatParam_DeleteMessagesByIds__FOlD = &PB_ChatParam_DeleteMessagesByIds{
    ChatKey:  "" ,
    Both:  false ,
    
}


type PB_ChatResponse_DeleteMessagesByIds_Flat struct {
}
//ToPB
func(m *PB_ChatResponse_DeleteMessagesByIds)ToFlat() *PB_ChatResponse_DeleteMessagesByIds_Flat {
r := &PB_ChatResponse_DeleteMessagesByIds_Flat{
}
return r
}
//ToPB
func(m *PB_ChatResponse_DeleteMessagesByIds_Flat)ToPB() *PB_ChatResponse_DeleteMessagesByIds {
r := &PB_ChatResponse_DeleteMessagesByIds{
}
return r
}
//folding
var PB_ChatResponse_DeleteMessagesByIds__FOlD = &PB_ChatResponse_DeleteMessagesByIds{
}


type PB_ChatParam_SetMessagesAsReceived_Flat struct {
    ChatRoom string
    MessageIds []int
}
//ToPB
func(m *PB_ChatParam_SetMessagesAsReceived)ToFlat() *PB_ChatParam_SetMessagesAsReceived_Flat {
r := &PB_ChatParam_SetMessagesAsReceived_Flat{
    ChatRoom:  m.ChatRoom ,
    MessageIds:  helper.SliceInt64ToInt(m.MessageIds) ,
}
return r
}
//ToPB
func(m *PB_ChatParam_SetMessagesAsReceived_Flat)ToPB() *PB_ChatParam_SetMessagesAsReceived {
r := &PB_ChatParam_SetMessagesAsReceived{
    ChatRoom:  m.ChatRoom ,
    MessageIds:  helper.SliceIntToInt64(m.MessageIds) ,
}
return r
}
//folding
var PB_ChatParam_SetMessagesAsReceived__FOlD = &PB_ChatParam_SetMessagesAsReceived{
    ChatRoom:  "" ,
    
}


type PB_ChatResponse_SetMessagesAsReceived_Flat struct {
}
//ToPB
func(m *PB_ChatResponse_SetMessagesAsReceived)ToFlat() *PB_ChatResponse_SetMessagesAsReceived_Flat {
r := &PB_ChatResponse_SetMessagesAsReceived_Flat{
}
return r
}
//ToPB
func(m *PB_ChatResponse_SetMessagesAsReceived_Flat)ToPB() *PB_ChatResponse_SetMessagesAsReceived {
r := &PB_ChatResponse_SetMessagesAsReceived{
}
return r
}
//folding
var PB_ChatResponse_SetMessagesAsReceived__FOlD = &PB_ChatResponse_SetMessagesAsReceived{
}


type PB_ChatParam_EditMessage_Flat struct {
    RoomKey string
    MessageId int
    NewText string
}
//ToPB
func(m *PB_ChatParam_EditMessage)ToFlat() *PB_ChatParam_EditMessage_Flat {
r := &PB_ChatParam_EditMessage_Flat{
    RoomKey:  m.RoomKey ,
    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}
//ToPB
func(m *PB_ChatParam_EditMessage_Flat)ToPB() *PB_ChatParam_EditMessage {
r := &PB_ChatParam_EditMessage{
    RoomKey:  m.RoomKey ,
    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}
//folding
var PB_ChatParam_EditMessage__FOlD = &PB_ChatParam_EditMessage{
    RoomKey:  "" ,
    MessageId:  0 ,
    NewText:  "" ,
}


type PB_ChatResponse_EditMessage_Flat struct {
}
//ToPB
func(m *PB_ChatResponse_EditMessage)ToFlat() *PB_ChatResponse_EditMessage_Flat {
r := &PB_ChatResponse_EditMessage_Flat{
}
return r
}
//ToPB
func(m *PB_ChatResponse_EditMessage_Flat)ToPB() *PB_ChatResponse_EditMessage {
r := &PB_ChatResponse_EditMessage{
}
return r
}
//folding
var PB_ChatResponse_EditMessage__FOlD = &PB_ChatResponse_EditMessage{
}


type PB_ChatParam_GetChatList_Flat struct {
}
//ToPB
func(m *PB_ChatParam_GetChatList)ToFlat() *PB_ChatParam_GetChatList_Flat {
r := &PB_ChatParam_GetChatList_Flat{
}
return r
}
//ToPB
func(m *PB_ChatParam_GetChatList_Flat)ToPB() *PB_ChatParam_GetChatList {
r := &PB_ChatParam_GetChatList{
}
return r
}
//folding
var PB_ChatParam_GetChatList__FOlD = &PB_ChatParam_GetChatList{
}


type PB_ChatResponse_GetChatList_Flat struct {
    Chats []PB_ChatView
    Users []PB_UserView
}
//ToPB
func(m *PB_ChatResponse_GetChatList)ToFlat() *PB_ChatResponse_GetChatList_Flat {
r := &PB_ChatResponse_GetChatList_Flat{
    
    
}
return r
}
//ToPB
func(m *PB_ChatResponse_GetChatList_Flat)ToPB() *PB_ChatResponse_GetChatList {
r := &PB_ChatResponse_GetChatList{
    
    
}
return r
}
//folding
var PB_ChatResponse_GetChatList__FOlD = &PB_ChatResponse_GetChatList{
    
    
}


type PB_ChatParam_GetChatHistoryToOlder_Flat struct {
    ChatKey string
    Limit int
    FromTopMessageId int
}
//ToPB
func(m *PB_ChatParam_GetChatHistoryToOlder)ToFlat() *PB_ChatParam_GetChatHistoryToOlder_Flat {
r := &PB_ChatParam_GetChatHistoryToOlder_Flat{
    ChatKey:  m.ChatKey ,
    Limit:  int(m.Limit) ,
    FromTopMessageId:  int(m.FromTopMessageId) ,
}
return r
}
//ToPB
func(m *PB_ChatParam_GetChatHistoryToOlder_Flat)ToPB() *PB_ChatParam_GetChatHistoryToOlder {
r := &PB_ChatParam_GetChatHistoryToOlder{
    ChatKey:  m.ChatKey ,
    Limit:  int32(m.Limit) ,
    FromTopMessageId:  int64(m.FromTopMessageId) ,
}
return r
}
//folding
var PB_ChatParam_GetChatHistoryToOlder__FOlD = &PB_ChatParam_GetChatHistoryToOlder{
    ChatKey:  "" ,
    Limit:  0 ,
    FromTopMessageId:  0 ,
}


type PB_ChatResponse_GetChatHistoryToOlder_Flat struct {
    MessagesViews []PB_MessageView
    HasMore bool
}
//ToPB
func(m *PB_ChatResponse_GetChatHistoryToOlder)ToFlat() *PB_ChatResponse_GetChatHistoryToOlder_Flat {
r := &PB_ChatResponse_GetChatHistoryToOlder_Flat{
    
    HasMore:  m.HasMore ,
}
return r
}
//ToPB
func(m *PB_ChatResponse_GetChatHistoryToOlder_Flat)ToPB() *PB_ChatResponse_GetChatHistoryToOlder {
r := &PB_ChatResponse_GetChatHistoryToOlder{
    
    HasMore:  m.HasMore ,
}
return r
}
//folding
var PB_ChatResponse_GetChatHistoryToOlder__FOlD = &PB_ChatResponse_GetChatHistoryToOlder{
    
    HasMore:  false ,
}


type PB_ChatParam_GetFreshAllDirectMessagesList_Flat struct {
    LowMessageId int
    HighMessageId int
}
//ToPB
func(m *PB_ChatParam_GetFreshAllDirectMessagesList)ToFlat() *PB_ChatParam_GetFreshAllDirectMessagesList_Flat {
r := &PB_ChatParam_GetFreshAllDirectMessagesList_Flat{
    LowMessageId:  int(m.LowMessageId) ,
    HighMessageId:  int(m.HighMessageId) ,
}
return r
}
//ToPB
func(m *PB_ChatParam_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_ChatParam_GetFreshAllDirectMessagesList {
r := &PB_ChatParam_GetFreshAllDirectMessagesList{
    LowMessageId:  int64(m.LowMessageId) ,
    HighMessageId:  int64(m.HighMessageId) ,
}
return r
}
//folding
var PB_ChatParam_GetFreshAllDirectMessagesList__FOlD = &PB_ChatParam_GetFreshAllDirectMessagesList{
    LowMessageId:  0 ,
    HighMessageId:  0 ,
}


type PB_ChatResponse_GetFreshAllDirectMessagesList_Flat struct {
    Messages []PB_MessageView
    HasMore bool
}
//ToPB
func(m *PB_ChatResponse_GetFreshAllDirectMessagesList)ToFlat() *PB_ChatResponse_GetFreshAllDirectMessagesList_Flat {
r := &PB_ChatResponse_GetFreshAllDirectMessagesList_Flat{
    
    HasMore:  m.HasMore ,
}
return r
}
//ToPB
func(m *PB_ChatResponse_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_ChatResponse_GetFreshAllDirectMessagesList {
r := &PB_ChatResponse_GetFreshAllDirectMessagesList{
    
    HasMore:  m.HasMore ,
}
return r
}
//folding
var PB_ChatResponse_GetFreshAllDirectMessagesList__FOlD = &PB_ChatResponse_GetFreshAllDirectMessagesList{
    
    HasMore:  false ,
}


type PB_MsgParam_AddNewTextMessage_Flat struct {
    Text string
    MessageKey string
    ToRoomKey string
    PeerId int
    Time int
    ReplyToMessageId int
    Forward PB_MessageForwardedFrom
}
//ToPB
func(m *PB_MsgParam_AddNewTextMessage)ToFlat() *PB_MsgParam_AddNewTextMessage_Flat {
r := &PB_MsgParam_AddNewTextMessage_Flat{
    Text:  m.Text ,
    MessageKey:  m.MessageKey ,
    ToRoomKey:  m.ToRoomKey ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,
    ReplyToMessageId:  int(m.ReplyToMessageId) ,
    
}
return r
}
//ToPB
func(m *PB_MsgParam_AddNewTextMessage_Flat)ToPB() *PB_MsgParam_AddNewTextMessage {
r := &PB_MsgParam_AddNewTextMessage{
    Text:  m.Text ,
    MessageKey:  m.MessageKey ,
    ToRoomKey:  m.ToRoomKey ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,
    
}
return r
}
//folding
var PB_MsgParam_AddNewTextMessage__FOlD = &PB_MsgParam_AddNewTextMessage{
    Text:  "" ,
    MessageKey:  "" ,
    ToRoomKey:  "" ,
    PeerId:  0 ,
    Time:  0 ,
    ReplyToMessageId:  0 ,
    
}


type PB_MsgResponse_AddNewTextMessage_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_AddNewTextMessage)ToFlat() *PB_MsgResponse_AddNewTextMessage_Flat {
r := &PB_MsgResponse_AddNewTextMessage_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_AddNewTextMessage_Flat)ToPB() *PB_MsgResponse_AddNewTextMessage {
r := &PB_MsgResponse_AddNewTextMessage{
}
return r
}
//folding
var PB_MsgResponse_AddNewTextMessage__FOlD = &PB_MsgResponse_AddNewTextMessage{
}


type PB_MsgParam_AddNewMessage_Flat struct {
    Text string
    MessageKey string
    ToRoomKey string
    PeerId int
    Time int
    ReplyToMessageId int
    Forward PB_MessageForwardedFrom
    RoomMessageType RoomMessageTypeEnum
    Blob []byte
    MessageView PB_MessageView
}
//ToPB
func(m *PB_MsgParam_AddNewMessage)ToFlat() *PB_MsgParam_AddNewMessage_Flat {
r := &PB_MsgParam_AddNewMessage_Flat{
    Text:  m.Text ,
    MessageKey:  m.MessageKey ,
    ToRoomKey:  m.ToRoomKey ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,
    ReplyToMessageId:  int(m.ReplyToMessageId) ,
    
    
    Blob:  []byte(m.Blob) ,
    
}
return r
}
//ToPB
func(m *PB_MsgParam_AddNewMessage_Flat)ToPB() *PB_MsgParam_AddNewMessage {
r := &PB_MsgParam_AddNewMessage{
    Text:  m.Text ,
    MessageKey:  m.MessageKey ,
    ToRoomKey:  m.ToRoomKey ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,
    
    
    Blob:  m.Blob ,
    
}
return r
}
//folding
var PB_MsgParam_AddNewMessage__FOlD = &PB_MsgParam_AddNewMessage{
    Text:  "" ,
    MessageKey:  "" ,
    ToRoomKey:  "" ,
    PeerId:  0 ,
    Time:  0 ,
    ReplyToMessageId:  0 ,
    
    
    Blob:  []byte{} ,
    
}


type PB_MsgResponse_AddNewMessage_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_AddNewMessage)ToFlat() *PB_MsgResponse_AddNewMessage_Flat {
r := &PB_MsgResponse_AddNewMessage_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_AddNewMessage_Flat)ToPB() *PB_MsgResponse_AddNewMessage {
r := &PB_MsgResponse_AddNewMessage{
}
return r
}
//folding
var PB_MsgResponse_AddNewMessage__FOlD = &PB_MsgResponse_AddNewMessage{
}


type PB_MsgParam_SetRoomActionDoing_Flat struct {
    GroupId int
    DirectRoomKey string
    ActionType RoomActionDoingEnum
}
//ToPB
func(m *PB_MsgParam_SetRoomActionDoing)ToFlat() *PB_MsgParam_SetRoomActionDoing_Flat {
r := &PB_MsgParam_SetRoomActionDoing_Flat{
    GroupId:  int(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,
    
}
return r
}
//ToPB
func(m *PB_MsgParam_SetRoomActionDoing_Flat)ToPB() *PB_MsgParam_SetRoomActionDoing {
r := &PB_MsgParam_SetRoomActionDoing{
    GroupId:  int64(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,
    
}
return r
}
//folding
var PB_MsgParam_SetRoomActionDoing__FOlD = &PB_MsgParam_SetRoomActionDoing{
    GroupId:  0 ,
    DirectRoomKey:  "" ,
    
}


type PB_MsgResponse_SetRoomActionDoing_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_SetRoomActionDoing)ToFlat() *PB_MsgResponse_SetRoomActionDoing_Flat {
r := &PB_MsgResponse_SetRoomActionDoing_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_SetRoomActionDoing_Flat)ToPB() *PB_MsgResponse_SetRoomActionDoing {
r := &PB_MsgResponse_SetRoomActionDoing{
}
return r
}
//folding
var PB_MsgResponse_SetRoomActionDoing__FOlD = &PB_MsgResponse_SetRoomActionDoing{
}


type PB_MsgParam_GetMessagesByIds_Flat struct {
    MessagesCollections PB_MessagesCollections
}
//ToPB
func(m *PB_MsgParam_GetMessagesByIds)ToFlat() *PB_MsgParam_GetMessagesByIds_Flat {
r := &PB_MsgParam_GetMessagesByIds_Flat{
    
}
return r
}
//ToPB
func(m *PB_MsgParam_GetMessagesByIds_Flat)ToPB() *PB_MsgParam_GetMessagesByIds {
r := &PB_MsgParam_GetMessagesByIds{
    
}
return r
}
//folding
var PB_MsgParam_GetMessagesByIds__FOlD = &PB_MsgParam_GetMessagesByIds{
    
}


type PB_MsgResponse_GetMessagesByIds_Flat struct {
    MessagesViews []PB_MessageView
}
//ToPB
func(m *PB_MsgResponse_GetMessagesByIds)ToFlat() *PB_MsgResponse_GetMessagesByIds_Flat {
r := &PB_MsgResponse_GetMessagesByIds_Flat{
    
}
return r
}
//ToPB
func(m *PB_MsgResponse_GetMessagesByIds_Flat)ToPB() *PB_MsgResponse_GetMessagesByIds {
r := &PB_MsgResponse_GetMessagesByIds{
    
}
return r
}
//folding
var PB_MsgResponse_GetMessagesByIds__FOlD = &PB_MsgResponse_GetMessagesByIds{
    
}


type PB_MsgParam_GetMessagesHistory_Flat struct {
    ChatKey string
    FromSeq int
    EndSeq int
}
//ToPB
func(m *PB_MsgParam_GetMessagesHistory)ToFlat() *PB_MsgParam_GetMessagesHistory_Flat {
r := &PB_MsgParam_GetMessagesHistory_Flat{
    ChatKey:  m.ChatKey ,
    FromSeq:  int(m.FromSeq) ,
    EndSeq:  int(m.EndSeq) ,
}
return r
}
//ToPB
func(m *PB_MsgParam_GetMessagesHistory_Flat)ToPB() *PB_MsgParam_GetMessagesHistory {
r := &PB_MsgParam_GetMessagesHistory{
    ChatKey:  m.ChatKey ,
    FromSeq:  int32(m.FromSeq) ,
    EndSeq:  int32(m.EndSeq) ,
}
return r
}
//folding
var PB_MsgParam_GetMessagesHistory__FOlD = &PB_MsgParam_GetMessagesHistory{
    ChatKey:  "" ,
    FromSeq:  0 ,
    EndSeq:  0 ,
}


type PB_MsgResponse_GetMessagesHistory_Flat struct {
    MessagesViews []PB_MessageView
}
//ToPB
func(m *PB_MsgResponse_GetMessagesHistory)ToFlat() *PB_MsgResponse_GetMessagesHistory_Flat {
r := &PB_MsgResponse_GetMessagesHistory_Flat{
    
}
return r
}
//ToPB
func(m *PB_MsgResponse_GetMessagesHistory_Flat)ToPB() *PB_MsgResponse_GetMessagesHistory {
r := &PB_MsgResponse_GetMessagesHistory{
    
}
return r
}
//folding
var PB_MsgResponse_GetMessagesHistory__FOlD = &PB_MsgResponse_GetMessagesHistory{
    
}


type PB_MsgParam_SetChatMessagesRangeAsSeen_Flat struct {
    ChatKey string
    FromSeq int
    EndSeq int
    SeenTimeMs int
}
//ToPB
func(m *PB_MsgParam_SetChatMessagesRangeAsSeen)ToFlat() *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat {
r := &PB_MsgParam_SetChatMessagesRangeAsSeen_Flat{
    ChatKey:  m.ChatKey ,
    FromSeq:  int(m.FromSeq) ,
    EndSeq:  int(m.EndSeq) ,
    SeenTimeMs:  int(m.SeenTimeMs) ,
}
return r
}
//ToPB
func(m *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_MsgParam_SetChatMessagesRangeAsSeen {
r := &PB_MsgParam_SetChatMessagesRangeAsSeen{
    ChatKey:  m.ChatKey ,
    FromSeq:  int32(m.FromSeq) ,
    EndSeq:  int32(m.EndSeq) ,
    SeenTimeMs:  int64(m.SeenTimeMs) ,
}
return r
}
//folding
var PB_MsgParam_SetChatMessagesRangeAsSeen__FOlD = &PB_MsgParam_SetChatMessagesRangeAsSeen{
    ChatKey:  "" ,
    FromSeq:  0 ,
    EndSeq:  0 ,
    SeenTimeMs:  0 ,
}


type PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_SetChatMessagesRangeAsSeen)ToFlat() *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat {
r := &PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_MsgResponse_SetChatMessagesRangeAsSeen {
r := &PB_MsgResponse_SetChatMessagesRangeAsSeen{
}
return r
}
//folding
var PB_MsgResponse_SetChatMessagesRangeAsSeen__FOlD = &PB_MsgResponse_SetChatMessagesRangeAsSeen{
}


type PB_MsgParam_DeleteChatHistory_Flat struct {
    ChatKey string
    ToSeq int
}
//ToPB
func(m *PB_MsgParam_DeleteChatHistory)ToFlat() *PB_MsgParam_DeleteChatHistory_Flat {
r := &PB_MsgParam_DeleteChatHistory_Flat{
    ChatKey:  m.ChatKey ,
    ToSeq:  int(m.ToSeq) ,
}
return r
}
//ToPB
func(m *PB_MsgParam_DeleteChatHistory_Flat)ToPB() *PB_MsgParam_DeleteChatHistory {
r := &PB_MsgParam_DeleteChatHistory{
    ChatKey:  m.ChatKey ,
    ToSeq:  int32(m.ToSeq) ,
}
return r
}
//folding
var PB_MsgParam_DeleteChatHistory__FOlD = &PB_MsgParam_DeleteChatHistory{
    ChatKey:  "" ,
    ToSeq:  0 ,
}


type PB_MsgResponse_DeleteChatHistory_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_DeleteChatHistory)ToFlat() *PB_MsgResponse_DeleteChatHistory_Flat {
r := &PB_MsgResponse_DeleteChatHistory_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_DeleteChatHistory_Flat)ToPB() *PB_MsgResponse_DeleteChatHistory {
r := &PB_MsgResponse_DeleteChatHistory{
}
return r
}
//folding
var PB_MsgResponse_DeleteChatHistory__FOlD = &PB_MsgResponse_DeleteChatHistory{
}


type PB_MsgParam_DeleteMessagesByIds_Flat struct {
    MessagesCollections PB_MessagesCollections
}
//ToPB
func(m *PB_MsgParam_DeleteMessagesByIds)ToFlat() *PB_MsgParam_DeleteMessagesByIds_Flat {
r := &PB_MsgParam_DeleteMessagesByIds_Flat{
    
}
return r
}
//ToPB
func(m *PB_MsgParam_DeleteMessagesByIds_Flat)ToPB() *PB_MsgParam_DeleteMessagesByIds {
r := &PB_MsgParam_DeleteMessagesByIds{
    
}
return r
}
//folding
var PB_MsgParam_DeleteMessagesByIds__FOlD = &PB_MsgParam_DeleteMessagesByIds{
    
}


type PB_MsgResponse_DeleteMessagesByIds_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_DeleteMessagesByIds)ToFlat() *PB_MsgResponse_DeleteMessagesByIds_Flat {
r := &PB_MsgResponse_DeleteMessagesByIds_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_DeleteMessagesByIds_Flat)ToPB() *PB_MsgResponse_DeleteMessagesByIds {
r := &PB_MsgResponse_DeleteMessagesByIds{
}
return r
}
//folding
var PB_MsgResponse_DeleteMessagesByIds__FOlD = &PB_MsgResponse_DeleteMessagesByIds{
}


type PB_MsgParam_SetMessagesAsReceived_Flat struct {
    MessagesCollections PB_MessagesCollections
}
//ToPB
func(m *PB_MsgParam_SetMessagesAsReceived)ToFlat() *PB_MsgParam_SetMessagesAsReceived_Flat {
r := &PB_MsgParam_SetMessagesAsReceived_Flat{
    
}
return r
}
//ToPB
func(m *PB_MsgParam_SetMessagesAsReceived_Flat)ToPB() *PB_MsgParam_SetMessagesAsReceived {
r := &PB_MsgParam_SetMessagesAsReceived{
    
}
return r
}
//folding
var PB_MsgParam_SetMessagesAsReceived__FOlD = &PB_MsgParam_SetMessagesAsReceived{
    
}


type PB_MsgResponse_SetMessagesAsReceived_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_SetMessagesAsReceived)ToFlat() *PB_MsgResponse_SetMessagesAsReceived_Flat {
r := &PB_MsgResponse_SetMessagesAsReceived_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_SetMessagesAsReceived_Flat)ToPB() *PB_MsgResponse_SetMessagesAsReceived {
r := &PB_MsgResponse_SetMessagesAsReceived{
}
return r
}
//folding
var PB_MsgResponse_SetMessagesAsReceived__FOlD = &PB_MsgResponse_SetMessagesAsReceived{
}


type PB_MsgParam_ForwardMessages_Flat struct {
    MessagesCollections PB_MessagesCollections
    ToDirectChatKeys []string
    ToGroupChatIds []int
}
//ToPB
func(m *PB_MsgParam_ForwardMessages)ToFlat() *PB_MsgParam_ForwardMessages_Flat {
r := &PB_MsgParam_ForwardMessages_Flat{
    
    ToDirectChatKeys:  m.ToDirectChatKeys ,
    ToGroupChatIds:  helper.SliceInt64ToInt(m.ToGroupChatIds) ,
}
return r
}
//ToPB
func(m *PB_MsgParam_ForwardMessages_Flat)ToPB() *PB_MsgParam_ForwardMessages {
r := &PB_MsgParam_ForwardMessages{
    
    ToDirectChatKeys:  m.ToDirectChatKeys ,
    ToGroupChatIds:  helper.SliceIntToInt64(m.ToGroupChatIds) ,
}
return r
}
//folding
var PB_MsgParam_ForwardMessages__FOlD = &PB_MsgParam_ForwardMessages{
    
    
    
}


type PB_MsgResponse_ForwardMessages_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_ForwardMessages)ToFlat() *PB_MsgResponse_ForwardMessages_Flat {
r := &PB_MsgResponse_ForwardMessages_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_ForwardMessages_Flat)ToPB() *PB_MsgResponse_ForwardMessages {
r := &PB_MsgResponse_ForwardMessages{
}
return r
}
//folding
var PB_MsgResponse_ForwardMessages__FOlD = &PB_MsgResponse_ForwardMessages{
}


type PB_MsgParam_EditMessage_Flat struct {
    ChatKey string
    RoomType RoomTypeEnum
    MessageId int
    NewText string
}
//ToPB
func(m *PB_MsgParam_EditMessage)ToFlat() *PB_MsgParam_EditMessage_Flat {
r := &PB_MsgParam_EditMessage_Flat{
    ChatKey:  m.ChatKey ,
    
    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}
//ToPB
func(m *PB_MsgParam_EditMessage_Flat)ToPB() *PB_MsgParam_EditMessage {
r := &PB_MsgParam_EditMessage{
    ChatKey:  m.ChatKey ,
    
    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}
//folding
var PB_MsgParam_EditMessage__FOlD = &PB_MsgParam_EditMessage{
    ChatKey:  "" ,
    
    MessageId:  0 ,
    NewText:  "" ,
}


type PB_MsgResponse_EditMessage_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_EditMessage)ToFlat() *PB_MsgResponse_EditMessage_Flat {
r := &PB_MsgResponse_EditMessage_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_EditMessage_Flat)ToPB() *PB_MsgResponse_EditMessage {
r := &PB_MsgResponse_EditMessage{
}
return r
}
//folding
var PB_MsgResponse_EditMessage__FOlD = &PB_MsgResponse_EditMessage{
}


type PB_MsgParam_BroadcastNewMessage_Flat struct {
    Text string
    PeerId int
    Time int
    Forward PB_MessageForwardedFrom
}
//ToPB
func(m *PB_MsgParam_BroadcastNewMessage)ToFlat() *PB_MsgParam_BroadcastNewMessage_Flat {
r := &PB_MsgParam_BroadcastNewMessage_Flat{
    Text:  m.Text ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,
    
}
return r
}
//ToPB
func(m *PB_MsgParam_BroadcastNewMessage_Flat)ToPB() *PB_MsgParam_BroadcastNewMessage {
r := &PB_MsgParam_BroadcastNewMessage{
    Text:  m.Text ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,
    
}
return r
}
//folding
var PB_MsgParam_BroadcastNewMessage__FOlD = &PB_MsgParam_BroadcastNewMessage{
    Text:  "" ,
    PeerId:  0 ,
    Time:  0 ,
    
}


type PB_MsgResponse_BroadcastNewMessage_Flat struct {
}
//ToPB
func(m *PB_MsgResponse_BroadcastNewMessage)ToFlat() *PB_MsgResponse_BroadcastNewMessage_Flat {
r := &PB_MsgResponse_BroadcastNewMessage_Flat{
}
return r
}
//ToPB
func(m *PB_MsgResponse_BroadcastNewMessage_Flat)ToPB() *PB_MsgResponse_BroadcastNewMessage {
r := &PB_MsgResponse_BroadcastNewMessage{
}
return r
}
//folding
var PB_MsgResponse_BroadcastNewMessage__FOlD = &PB_MsgResponse_BroadcastNewMessage{
}


type PB_MsgParam_GetFreshChatList_Flat struct {
}
//ToPB
func(m *PB_MsgParam_GetFreshChatList)ToFlat() *PB_MsgParam_GetFreshChatList_Flat {
r := &PB_MsgParam_GetFreshChatList_Flat{
}
return r
}
//ToPB
func(m *PB_MsgParam_GetFreshChatList_Flat)ToPB() *PB_MsgParam_GetFreshChatList {
r := &PB_MsgParam_GetFreshChatList{
}
return r
}
//folding
var PB_MsgParam_GetFreshChatList__FOlD = &PB_MsgParam_GetFreshChatList{
}


type PB_MsgResponse_GetFreshChatList_Flat struct {
    Chats []PB_ChatView
    Users []PB_UserView
}
//ToPB
func(m *PB_MsgResponse_GetFreshChatList)ToFlat() *PB_MsgResponse_GetFreshChatList_Flat {
r := &PB_MsgResponse_GetFreshChatList_Flat{
    
    
}
return r
}
//ToPB
func(m *PB_MsgResponse_GetFreshChatList_Flat)ToPB() *PB_MsgResponse_GetFreshChatList {
r := &PB_MsgResponse_GetFreshChatList{
    
    
}
return r
}
//folding
var PB_MsgResponse_GetFreshChatList__FOlD = &PB_MsgResponse_GetFreshChatList{
    
    
}


type PB_MsgParam_GetFreshRoomMessagesList_Flat struct {
    ChatKey string
    RoomKey string
    Last int
}
//ToPB
func(m *PB_MsgParam_GetFreshRoomMessagesList)ToFlat() *PB_MsgParam_GetFreshRoomMessagesList_Flat {
r := &PB_MsgParam_GetFreshRoomMessagesList_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    Last:  int(m.Last) ,
}
return r
}
//ToPB
func(m *PB_MsgParam_GetFreshRoomMessagesList_Flat)ToPB() *PB_MsgParam_GetFreshRoomMessagesList {
r := &PB_MsgParam_GetFreshRoomMessagesList{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    Last:  int64(m.Last) ,
}
return r
}
//folding
var PB_MsgParam_GetFreshRoomMessagesList__FOlD = &PB_MsgParam_GetFreshRoomMessagesList{
    ChatKey:  "" ,
    RoomKey:  "" ,
    Last:  0 ,
}


type PB_MsgResponse_GetFreshRoomMessagesList_Flat struct {
    Messages []PB_MessageView
    HasMore bool
}
//ToPB
func(m *PB_MsgResponse_GetFreshRoomMessagesList)ToFlat() *PB_MsgResponse_GetFreshRoomMessagesList_Flat {
r := &PB_MsgResponse_GetFreshRoomMessagesList_Flat{
    
    HasMore:  m.HasMore ,
}
return r
}
//ToPB
func(m *PB_MsgResponse_GetFreshRoomMessagesList_Flat)ToPB() *PB_MsgResponse_GetFreshRoomMessagesList {
r := &PB_MsgResponse_GetFreshRoomMessagesList{
    
    HasMore:  m.HasMore ,
}
return r
}
//folding
var PB_MsgResponse_GetFreshRoomMessagesList__FOlD = &PB_MsgResponse_GetFreshRoomMessagesList{
    
    HasMore:  false ,
}


type PB_MsgParam_GetFreshAllDirectMessagesList_Flat struct {
    LowMessageId int
    HighMessageId int
}
//ToPB
func(m *PB_MsgParam_GetFreshAllDirectMessagesList)ToFlat() *PB_MsgParam_GetFreshAllDirectMessagesList_Flat {
r := &PB_MsgParam_GetFreshAllDirectMessagesList_Flat{
    LowMessageId:  int(m.LowMessageId) ,
    HighMessageId:  int(m.HighMessageId) ,
}
return r
}
//ToPB
func(m *PB_MsgParam_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_MsgParam_GetFreshAllDirectMessagesList {
r := &PB_MsgParam_GetFreshAllDirectMessagesList{
    LowMessageId:  int64(m.LowMessageId) ,
    HighMessageId:  int64(m.HighMessageId) ,
}
return r
}
//folding
var PB_MsgParam_GetFreshAllDirectMessagesList__FOlD = &PB_MsgParam_GetFreshAllDirectMessagesList{
    LowMessageId:  0 ,
    HighMessageId:  0 ,
}


type PB_MsgResponse_GetFreshAllDirectMessagesList_Flat struct {
    Messages []PB_MessageView
    HasMore bool
}
//ToPB
func(m *PB_MsgResponse_GetFreshAllDirectMessagesList)ToFlat() *PB_MsgResponse_GetFreshAllDirectMessagesList_Flat {
r := &PB_MsgResponse_GetFreshAllDirectMessagesList_Flat{
    
    HasMore:  m.HasMore ,
}
return r
}
//ToPB
func(m *PB_MsgResponse_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_MsgResponse_GetFreshAllDirectMessagesList {
r := &PB_MsgResponse_GetFreshAllDirectMessagesList{
    
    HasMore:  m.HasMore ,
}
return r
}
//folding
var PB_MsgResponse_GetFreshAllDirectMessagesList__FOlD = &PB_MsgResponse_GetFreshAllDirectMessagesList{
    
    HasMore:  false ,
}


type PB_MessageForwardedFrom_Flat struct {
    RoomId int
    RoomTypeEnum RoomTypeEnum
    MessageId int
    MessageSeq int
}
//ToPB
func(m *PB_MessageForwardedFrom)ToFlat() *PB_MessageForwardedFrom_Flat {
r := &PB_MessageForwardedFrom_Flat{
    RoomId:  int(m.RoomId) ,
    
    MessageId:  int(m.MessageId) ,
    MessageSeq:  int(m.MessageSeq) ,
}
return r
}
//ToPB
func(m *PB_MessageForwardedFrom_Flat)ToPB() *PB_MessageForwardedFrom {
r := &PB_MessageForwardedFrom{
    RoomId:  int64(m.RoomId) ,
    
    MessageId:  int64(m.MessageId) ,
    MessageSeq:  int32(m.MessageSeq) ,
}
return r
}
//folding
var PB_MessageForwardedFrom__FOlD = &PB_MessageForwardedFrom{
    RoomId:  0 ,
    
    MessageId:  0 ,
    MessageSeq:  0 ,
}


type PB_MessagesCollections_Flat struct {
    DirectMessagesIds []int
    GroupMessagesIds []int
    BroadCatMessagesIds []int
}
//ToPB
func(m *PB_MessagesCollections)ToFlat() *PB_MessagesCollections_Flat {
r := &PB_MessagesCollections_Flat{
    DirectMessagesIds:  helper.SliceInt64ToInt(m.DirectMessagesIds) ,
    GroupMessagesIds:  helper.SliceInt64ToInt(m.GroupMessagesIds) ,
    BroadCatMessagesIds:  helper.SliceInt64ToInt(m.BroadCatMessagesIds) ,
}
return r
}
//ToPB
func(m *PB_MessagesCollections_Flat)ToPB() *PB_MessagesCollections {
r := &PB_MessagesCollections{
    DirectMessagesIds:  helper.SliceIntToInt64(m.DirectMessagesIds) ,
    GroupMessagesIds:  helper.SliceIntToInt64(m.GroupMessagesIds) ,
    BroadCatMessagesIds:  helper.SliceIntToInt64(m.BroadCatMessagesIds) ,
}
return r
}
//folding
var PB_MessagesCollections__FOlD = &PB_MessagesCollections{
    
    
    
}


type PB_MsgParam_Echo_Flat struct {
    Text string
}
//ToPB
func(m *PB_MsgParam_Echo)ToFlat() *PB_MsgParam_Echo_Flat {
r := &PB_MsgParam_Echo_Flat{
    Text:  m.Text ,
}
return r
}
//ToPB
func(m *PB_MsgParam_Echo_Flat)ToPB() *PB_MsgParam_Echo {
r := &PB_MsgParam_Echo{
    Text:  m.Text ,
}
return r
}
//folding
var PB_MsgParam_Echo__FOlD = &PB_MsgParam_Echo{
    Text:  "" ,
}


type PB_MsgResponse_PB_MsgParam_Echo_Flat struct {
    Text string
}
//ToPB
func(m *PB_MsgResponse_PB_MsgParam_Echo)ToFlat() *PB_MsgResponse_PB_MsgParam_Echo_Flat {
r := &PB_MsgResponse_PB_MsgParam_Echo_Flat{
    Text:  m.Text ,
}
return r
}
//ToPB
func(m *PB_MsgResponse_PB_MsgParam_Echo_Flat)ToPB() *PB_MsgResponse_PB_MsgParam_Echo {
r := &PB_MsgResponse_PB_MsgParam_Echo{
    Text:  m.Text ,
}
return r
}
//folding
var PB_MsgResponse_PB_MsgParam_Echo__FOlD = &PB_MsgResponse_PB_MsgParam_Echo{
    Text:  "" ,
}


type PB_SyncParam_GetDirectUpdates_Flat struct {
    LastId int
}
//ToPB
func(m *PB_SyncParam_GetDirectUpdates)ToFlat() *PB_SyncParam_GetDirectUpdates_Flat {
r := &PB_SyncParam_GetDirectUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncParam_GetDirectUpdates_Flat)ToPB() *PB_SyncParam_GetDirectUpdates {
r := &PB_SyncParam_GetDirectUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncParam_GetDirectUpdates__FOlD = &PB_SyncParam_GetDirectUpdates{
    LastId:  0 ,
}


type PB_SyncResponse_GetDirectUpdates_Flat struct {
    NewMessages []PB_MessageView
    ChatFiles []PB_MessageFileView
    Chats []PB_ChatView
    Users []PB_UserView
    MessagesChangeIds []PB_UpdateMessageId
    MessagesFileChangeIds []PB_UpdateMessageId
    MessagesToUpdate []PB_UpdateMessageToEdit
    MessagesToDelete []PB_UpdateMessageToDelete
    MessagesDelivierdToServer []PB_UpdateMessageMeta
    MessagesDelivierdToPeer []PB_UpdateMessageMeta
    MessagesSeenByPeer []PB_UpdateMessageMeta
    MessagesDeletedFromServer []PB_UpdateMessageMeta
    RoomActionDoing []PB_UpdateRoomActionDoing
    LastId int
}
//ToPB
func(m *PB_SyncResponse_GetDirectUpdates)ToFlat() *PB_SyncResponse_GetDirectUpdates_Flat {
r := &PB_SyncResponse_GetDirectUpdates_Flat{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncResponse_GetDirectUpdates_Flat)ToPB() *PB_SyncResponse_GetDirectUpdates {
r := &PB_SyncResponse_GetDirectUpdates{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncResponse_GetDirectUpdates__FOlD = &PB_SyncResponse_GetDirectUpdates{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  0 ,
}


type PB_SyncParam_GetGeneralUpdates_Flat struct {
    LastId int
}
//ToPB
func(m *PB_SyncParam_GetGeneralUpdates)ToFlat() *PB_SyncParam_GetGeneralUpdates_Flat {
r := &PB_SyncParam_GetGeneralUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncParam_GetGeneralUpdates_Flat)ToPB() *PB_SyncParam_GetGeneralUpdates {
r := &PB_SyncParam_GetGeneralUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncParam_GetGeneralUpdates__FOlD = &PB_SyncParam_GetGeneralUpdates{
    LastId:  0 ,
}


type PB_SyncResponse_GetGeneralUpdates_Flat struct {
    UserBlockedByMe []PB_UpdateUserBlocked
    UserBlockedMe []PB_UpdateUserBlocked
    LastId int
}
//ToPB
func(m *PB_SyncResponse_GetGeneralUpdates)ToFlat() *PB_SyncResponse_GetGeneralUpdates_Flat {
r := &PB_SyncResponse_GetGeneralUpdates_Flat{
    
    
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncResponse_GetGeneralUpdates_Flat)ToPB() *PB_SyncResponse_GetGeneralUpdates {
r := &PB_SyncResponse_GetGeneralUpdates{
    
    
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncResponse_GetGeneralUpdates__FOlD = &PB_SyncResponse_GetGeneralUpdates{
    
    
    LastId:  0 ,
}


type PB_SyncParam_GetNotifyUpdates_Flat struct {
    LastId int
}
//ToPB
func(m *PB_SyncParam_GetNotifyUpdates)ToFlat() *PB_SyncParam_GetNotifyUpdates_Flat {
r := &PB_SyncParam_GetNotifyUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncParam_GetNotifyUpdates_Flat)ToPB() *PB_SyncParam_GetNotifyUpdates {
r := &PB_SyncParam_GetNotifyUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncParam_GetNotifyUpdates__FOlD = &PB_SyncParam_GetNotifyUpdates{
    LastId:  0 ,
}


type PB_SyncResponse_GetNotifyUpdates_Flat struct {
    Updates PB_NotifyUpdatesView
    LastId int
}
//ToPB
func(m *PB_SyncResponse_GetNotifyUpdates)ToFlat() *PB_SyncResponse_GetNotifyUpdates_Flat {
r := &PB_SyncResponse_GetNotifyUpdates_Flat{
    
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncResponse_GetNotifyUpdates_Flat)ToPB() *PB_SyncResponse_GetNotifyUpdates {
r := &PB_SyncResponse_GetNotifyUpdates{
    
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncResponse_GetNotifyUpdates__FOlD = &PB_SyncResponse_GetNotifyUpdates{
    
    LastId:  0 ,
}


type PB_SyncParam_SetLastSyncDirectUpdateId_Flat struct {
    LastId int
}
//ToPB
func(m *PB_SyncParam_SetLastSyncDirectUpdateId)ToFlat() *PB_SyncParam_SetLastSyncDirectUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncDirectUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncParam_SetLastSyncDirectUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncDirectUpdateId {
r := &PB_SyncParam_SetLastSyncDirectUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncParam_SetLastSyncDirectUpdateId__FOlD = &PB_SyncParam_SetLastSyncDirectUpdateId{
    LastId:  0 ,
}


type PB_SyncResponse_SetLastSyncDirectUpdateId_Flat struct {
}
//ToPB
func(m *PB_SyncResponse_SetLastSyncDirectUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncDirectUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncDirectUpdateId_Flat{
}
return r
}
//ToPB
func(m *PB_SyncResponse_SetLastSyncDirectUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncDirectUpdateId {
r := &PB_SyncResponse_SetLastSyncDirectUpdateId{
}
return r
}
//folding
var PB_SyncResponse_SetLastSyncDirectUpdateId__FOlD = &PB_SyncResponse_SetLastSyncDirectUpdateId{
}


type PB_SyncParam_SetLastSyncGeneralUpdateId_Flat struct {
    LastId int
}
//ToPB
func(m *PB_SyncParam_SetLastSyncGeneralUpdateId)ToFlat() *PB_SyncParam_SetLastSyncGeneralUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncGeneralUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncParam_SetLastSyncGeneralUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncGeneralUpdateId {
r := &PB_SyncParam_SetLastSyncGeneralUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncParam_SetLastSyncGeneralUpdateId__FOlD = &PB_SyncParam_SetLastSyncGeneralUpdateId{
    LastId:  0 ,
}


type PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat struct {
}
//ToPB
func(m *PB_SyncResponse_SetLastSyncGeneralUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat{
}
return r
}
//ToPB
func(m *PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncGeneralUpdateId {
r := &PB_SyncResponse_SetLastSyncGeneralUpdateId{
}
return r
}
//folding
var PB_SyncResponse_SetLastSyncGeneralUpdateId__FOlD = &PB_SyncResponse_SetLastSyncGeneralUpdateId{
}


type PB_SyncParam_SetLastSyncNotifyUpdateId_Flat struct {
    LastId int
}
//ToPB
func(m *PB_SyncParam_SetLastSyncNotifyUpdateId)ToFlat() *PB_SyncParam_SetLastSyncNotifyUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncNotifyUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}
//ToPB
func(m *PB_SyncParam_SetLastSyncNotifyUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncNotifyUpdateId {
r := &PB_SyncParam_SetLastSyncNotifyUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}
//folding
var PB_SyncParam_SetLastSyncNotifyUpdateId__FOlD = &PB_SyncParam_SetLastSyncNotifyUpdateId{
    LastId:  0 ,
}


type PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat struct {
}
//ToPB
func(m *PB_SyncResponse_SetLastSyncNotifyUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat{
}
return r
}
//ToPB
func(m *PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncNotifyUpdateId {
r := &PB_SyncResponse_SetLastSyncNotifyUpdateId{
}
return r
}
//folding
var PB_SyncResponse_SetLastSyncNotifyUpdateId__FOlD = &PB_SyncResponse_SetLastSyncNotifyUpdateId{
}


type PB_NotifyUpdatesView_Flat struct {
    UserBlockedByMe []PB_UpdateUserBlocked
    UserBlockedMe []PB_UpdateUserBlocked
}
//ToPB
func(m *PB_NotifyUpdatesView)ToFlat() *PB_NotifyUpdatesView_Flat {
r := &PB_NotifyUpdatesView_Flat{
    
    
}
return r
}
//ToPB
func(m *PB_NotifyUpdatesView_Flat)ToPB() *PB_NotifyUpdatesView {
r := &PB_NotifyUpdatesView{
    
    
}
return r
}
//folding
var PB_NotifyUpdatesView__FOlD = &PB_NotifyUpdatesView{
    
    
}


type PB_AllLivePushes_Flat struct {
    DirectUpdates PB_SyncResponse_GetDirectUpdates
    GeneralUpdates PB_SyncResponse_GetGeneralUpdates
    Offline_Messagings PB_Offline_Messagings
}
//ToPB
func(m *PB_AllLivePushes)ToFlat() *PB_AllLivePushes_Flat {
r := &PB_AllLivePushes_Flat{
    
    
    
}
return r
}
//ToPB
func(m *PB_AllLivePushes_Flat)ToPB() *PB_AllLivePushes {
r := &PB_AllLivePushes{
    
    
    
}
return r
}
//folding
var PB_AllLivePushes__FOlD = &PB_AllLivePushes{
    
    
    
}


type PB_UserParam_BlockUser_Flat struct {
    UserId int
    UserName string
}
//ToPB
func(m *PB_UserParam_BlockUser)ToFlat() *PB_UserParam_BlockUser_Flat {
r := &PB_UserParam_BlockUser_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}
//ToPB
func(m *PB_UserParam_BlockUser_Flat)ToPB() *PB_UserParam_BlockUser {
r := &PB_UserParam_BlockUser{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}
//folding
var PB_UserParam_BlockUser__FOlD = &PB_UserParam_BlockUser{
    UserId:  0 ,
    UserName:  "" ,
}


type PB_UserOfflineResponse_BlockUser_Flat struct {
    ByUserId int
    TargetUserId int
    TargetUserName string
}
//ToPB
func(m *PB_UserOfflineResponse_BlockUser)ToFlat() *PB_UserOfflineResponse_BlockUser_Flat {
r := &PB_UserOfflineResponse_BlockUser_Flat{
    ByUserId:  int(m.ByUserId) ,
    TargetUserId:  int(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}
//ToPB
func(m *PB_UserOfflineResponse_BlockUser_Flat)ToPB() *PB_UserOfflineResponse_BlockUser {
r := &PB_UserOfflineResponse_BlockUser{
    ByUserId:  int32(m.ByUserId) ,
    TargetUserId:  int32(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}
//folding
var PB_UserOfflineResponse_BlockUser__FOlD = &PB_UserOfflineResponse_BlockUser{
    ByUserId:  0 ,
    TargetUserId:  0 ,
    TargetUserName:  "" ,
}


type PB_UserParam_UnBlockUser_Flat struct {
    Offset int
    Limit int
}
//ToPB
func(m *PB_UserParam_UnBlockUser)ToFlat() *PB_UserParam_UnBlockUser_Flat {
r := &PB_UserParam_UnBlockUser_Flat{
    Offset:  int(m.Offset) ,
    Limit:  int(m.Limit) ,
}
return r
}
//ToPB
func(m *PB_UserParam_UnBlockUser_Flat)ToPB() *PB_UserParam_UnBlockUser {
r := &PB_UserParam_UnBlockUser{
    Offset:  int32(m.Offset) ,
    Limit:  int32(m.Limit) ,
}
return r
}
//folding
var PB_UserParam_UnBlockUser__FOlD = &PB_UserParam_UnBlockUser{
    Offset:  0 ,
    Limit:  0 ,
}


type PB_UserOfflineResponse_UnBlockUser_Flat struct {
    Users []UserView
}
//ToPB
func(m *PB_UserOfflineResponse_UnBlockUser)ToFlat() *PB_UserOfflineResponse_UnBlockUser_Flat {
r := &PB_UserOfflineResponse_UnBlockUser_Flat{
    
}
return r
}
//ToPB
func(m *PB_UserOfflineResponse_UnBlockUser_Flat)ToPB() *PB_UserOfflineResponse_UnBlockUser {
r := &PB_UserOfflineResponse_UnBlockUser{
    
}
return r
}
//folding
var PB_UserOfflineResponse_UnBlockUser__FOlD = &PB_UserOfflineResponse_UnBlockUser{
    
}


type PB_UserParam_BlockedList_Flat struct {
    UserId int
    UserName string
}
//ToPB
func(m *PB_UserParam_BlockedList)ToFlat() *PB_UserParam_BlockedList_Flat {
r := &PB_UserParam_BlockedList_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}
//ToPB
func(m *PB_UserParam_BlockedList_Flat)ToPB() *PB_UserParam_BlockedList {
r := &PB_UserParam_BlockedList{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}
//folding
var PB_UserParam_BlockedList__FOlD = &PB_UserParam_BlockedList{
    UserId:  0 ,
    UserName:  "" ,
}


type PB_UserResponse_BlockedList_Flat struct {
    ByUserId int
    TargetUserId int
    TargetUserName string
}
//ToPB
func(m *PB_UserResponse_BlockedList)ToFlat() *PB_UserResponse_BlockedList_Flat {
r := &PB_UserResponse_BlockedList_Flat{
    ByUserId:  int(m.ByUserId) ,
    TargetUserId:  int(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}
//ToPB
func(m *PB_UserResponse_BlockedList_Flat)ToPB() *PB_UserResponse_BlockedList {
r := &PB_UserResponse_BlockedList{
    ByUserId:  int32(m.ByUserId) ,
    TargetUserId:  int32(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}
//folding
var PB_UserResponse_BlockedList__FOlD = &PB_UserResponse_BlockedList{
    ByUserId:  0 ,
    TargetUserId:  0 ,
    TargetUserName:  "" ,
}


type PB_UserParam_UpdateAbout_Flat struct {
    NewAbout string
}
//ToPB
func(m *PB_UserParam_UpdateAbout)ToFlat() *PB_UserParam_UpdateAbout_Flat {
r := &PB_UserParam_UpdateAbout_Flat{
    NewAbout:  m.NewAbout ,
}
return r
}
//ToPB
func(m *PB_UserParam_UpdateAbout_Flat)ToPB() *PB_UserParam_UpdateAbout {
r := &PB_UserParam_UpdateAbout{
    NewAbout:  m.NewAbout ,
}
return r
}
//folding
var PB_UserParam_UpdateAbout__FOlD = &PB_UserParam_UpdateAbout{
    NewAbout:  "" ,
}


type PB_UserOfflineResponse_UpdateAbout_Flat struct {
    UserId int
    NewAbout string
}
//ToPB
func(m *PB_UserOfflineResponse_UpdateAbout)ToFlat() *PB_UserOfflineResponse_UpdateAbout_Flat {
r := &PB_UserOfflineResponse_UpdateAbout_Flat{
    UserId:  int(m.UserId) ,
    NewAbout:  m.NewAbout ,
}
return r
}
//ToPB
func(m *PB_UserOfflineResponse_UpdateAbout_Flat)ToPB() *PB_UserOfflineResponse_UpdateAbout {
r := &PB_UserOfflineResponse_UpdateAbout{
    UserId:  int32(m.UserId) ,
    NewAbout:  m.NewAbout ,
}
return r
}
//folding
var PB_UserOfflineResponse_UpdateAbout__FOlD = &PB_UserOfflineResponse_UpdateAbout{
    UserId:  0 ,
    NewAbout:  "" ,
}


type PB_UserParam_UpdateUserName_Flat struct {
    NewUserName string
}
//ToPB
func(m *PB_UserParam_UpdateUserName)ToFlat() *PB_UserParam_UpdateUserName_Flat {
r := &PB_UserParam_UpdateUserName_Flat{
    NewUserName:  m.NewUserName ,
}
return r
}
//ToPB
func(m *PB_UserParam_UpdateUserName_Flat)ToPB() *PB_UserParam_UpdateUserName {
r := &PB_UserParam_UpdateUserName{
    NewUserName:  m.NewUserName ,
}
return r
}
//folding
var PB_UserParam_UpdateUserName__FOlD = &PB_UserParam_UpdateUserName{
    NewUserName:  "" ,
}


type PB_UserOfflineResponse_UpdateUserName_Flat struct {
    UserId int
    NewUserName string
}
//ToPB
func(m *PB_UserOfflineResponse_UpdateUserName)ToFlat() *PB_UserOfflineResponse_UpdateUserName_Flat {
r := &PB_UserOfflineResponse_UpdateUserName_Flat{
    UserId:  int(m.UserId) ,
    NewUserName:  m.NewUserName ,
}
return r
}
//ToPB
func(m *PB_UserOfflineResponse_UpdateUserName_Flat)ToPB() *PB_UserOfflineResponse_UpdateUserName {
r := &PB_UserOfflineResponse_UpdateUserName{
    UserId:  int32(m.UserId) ,
    NewUserName:  m.NewUserName ,
}
return r
}
//folding
var PB_UserOfflineResponse_UpdateUserName__FOlD = &PB_UserOfflineResponse_UpdateUserName{
    UserId:  0 ,
    NewUserName:  "" ,
}


type PB_UserParam_ChangeAvatar_Flat struct {
    None bool
    ImageData2 []byte
}
//ToPB
func(m *PB_UserParam_ChangeAvatar)ToFlat() *PB_UserParam_ChangeAvatar_Flat {
r := &PB_UserParam_ChangeAvatar_Flat{
    None:  m.None ,
    ImageData2:  []byte(m.ImageData2) ,
}
return r
}
//ToPB
func(m *PB_UserParam_ChangeAvatar_Flat)ToPB() *PB_UserParam_ChangeAvatar {
r := &PB_UserParam_ChangeAvatar{
    None:  m.None ,
    ImageData2:  m.ImageData2 ,
}
return r
}
//folding
var PB_UserParam_ChangeAvatar__FOlD = &PB_UserParam_ChangeAvatar{
    None:  false ,
    ImageData2:  []byte{} ,
}


type PB_UserOfflineResponse_ChangeAvatar_Flat struct {
}
//ToPB
func(m *PB_UserOfflineResponse_ChangeAvatar)ToFlat() *PB_UserOfflineResponse_ChangeAvatar_Flat {
r := &PB_UserOfflineResponse_ChangeAvatar_Flat{
}
return r
}
//ToPB
func(m *PB_UserOfflineResponse_ChangeAvatar_Flat)ToPB() *PB_UserOfflineResponse_ChangeAvatar {
r := &PB_UserOfflineResponse_ChangeAvatar{
}
return r
}
//folding
var PB_UserOfflineResponse_ChangeAvatar__FOlD = &PB_UserOfflineResponse_ChangeAvatar{
}


type PB_UserParam_ChangePrivacy_Flat struct {
    Level ProfilePrivacyLevelEnum
}
//ToPB
func(m *PB_UserParam_ChangePrivacy)ToFlat() *PB_UserParam_ChangePrivacy_Flat {
r := &PB_UserParam_ChangePrivacy_Flat{
    
}
return r
}
//ToPB
func(m *PB_UserParam_ChangePrivacy_Flat)ToPB() *PB_UserParam_ChangePrivacy {
r := &PB_UserParam_ChangePrivacy{
    
}
return r
}
//folding
var PB_UserParam_ChangePrivacy__FOlD = &PB_UserParam_ChangePrivacy{
    
}


type PB_UserResponseOffline_ChangePrivacy_Flat struct {
}
//ToPB
func(m *PB_UserResponseOffline_ChangePrivacy)ToFlat() *PB_UserResponseOffline_ChangePrivacy_Flat {
r := &PB_UserResponseOffline_ChangePrivacy_Flat{
}
return r
}
//ToPB
func(m *PB_UserResponseOffline_ChangePrivacy_Flat)ToPB() *PB_UserResponseOffline_ChangePrivacy {
r := &PB_UserResponseOffline_ChangePrivacy{
}
return r
}
//folding
var PB_UserResponseOffline_ChangePrivacy__FOlD = &PB_UserResponseOffline_ChangePrivacy{
}


type PB_UserParam_CheckUserName_Flat struct {
    Level ProfilePrivacyLevelEnum
}
//ToPB
func(m *PB_UserParam_CheckUserName)ToFlat() *PB_UserParam_CheckUserName_Flat {
r := &PB_UserParam_CheckUserName_Flat{
    
}
return r
}
//ToPB
func(m *PB_UserParam_CheckUserName_Flat)ToPB() *PB_UserParam_CheckUserName {
r := &PB_UserParam_CheckUserName{
    
}
return r
}
//folding
var PB_UserParam_CheckUserName__FOlD = &PB_UserParam_CheckUserName{
    
}


type PB_UserResponse_CheckUserName_Flat struct {
}
//ToPB
func(m *PB_UserResponse_CheckUserName)ToFlat() *PB_UserResponse_CheckUserName_Flat {
r := &PB_UserResponse_CheckUserName_Flat{
}
return r
}
//ToPB
func(m *PB_UserResponse_CheckUserName_Flat)ToPB() *PB_UserResponse_CheckUserName {
r := &PB_UserResponse_CheckUserName{
}
return r
}
//folding
var PB_UserResponse_CheckUserName__FOlD = &PB_UserResponse_CheckUserName{
}


type UserView_Flat struct {
}
//ToPB
func(m *UserView)ToFlat() *UserView_Flat {
r := &UserView_Flat{
}
return r
}
//ToPB
func(m *UserView_Flat)ToPB() *UserView {
r := &UserView{
}
return r
}
//folding
var UserView__FOlD = &UserView{
}


type PB_Activity_Flat struct {
    Id int
    ActorUserId int
    ActionTypeId int
    RowId int
    RootId int
    RefId int
    CreatedAt int
}
//ToPB
func(m *PB_Activity)ToFlat() *PB_Activity_Flat {
r := &PB_Activity_Flat{
    Id:  int(m.Id) ,
    ActorUserId:  int(m.ActorUserId) ,
    ActionTypeId:  int(m.ActionTypeId) ,
    RowId:  int(m.RowId) ,
    RootId:  int(m.RootId) ,
    RefId:  int(m.RefId) ,
    CreatedAt:  int(m.CreatedAt) ,
}
return r
}
//ToPB
func(m *PB_Activity_Flat)ToPB() *PB_Activity {
r := &PB_Activity{
    Id:  int64(m.Id) ,
    ActorUserId:  int32(m.ActorUserId) ,
    ActionTypeId:  int32(m.ActionTypeId) ,
    RowId:  int32(m.RowId) ,
    RootId:  int32(m.RootId) ,
    RefId:  int64(m.RefId) ,
    CreatedAt:  int32(m.CreatedAt) ,
}
return r
}
//folding
var PB_Activity__FOlD = &PB_Activity{
    Id:  0 ,
    ActorUserId:  0 ,
    ActionTypeId:  0 ,
    RowId:  0 ,
    RootId:  0 ,
    RefId:  0 ,
    CreatedAt:  0 ,
}


type PB_Bucket_Flat struct {
    BucketId int
    BucketName string
    Server1Id int
    Server2Id int
    BackupServerId int
    ContentObjectTypeId int
    ContentObjectCount int
    CreatedTime int
}
//ToPB
func(m *PB_Bucket)ToFlat() *PB_Bucket_Flat {
r := &PB_Bucket_Flat{
    BucketId:  int(m.BucketId) ,
    BucketName:  m.BucketName ,
    Server1Id:  int(m.Server1Id) ,
    Server2Id:  int(m.Server2Id) ,
    BackupServerId:  int(m.BackupServerId) ,
    ContentObjectTypeId:  int(m.ContentObjectTypeId) ,
    ContentObjectCount:  int(m.ContentObjectCount) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_Bucket_Flat)ToPB() *PB_Bucket {
r := &PB_Bucket{
    BucketId:  int32(m.BucketId) ,
    BucketName:  m.BucketName ,
    Server1Id:  int32(m.Server1Id) ,
    Server2Id:  int32(m.Server2Id) ,
    BackupServerId:  int32(m.BackupServerId) ,
    ContentObjectTypeId:  int32(m.ContentObjectTypeId) ,
    ContentObjectCount:  int32(m.ContentObjectCount) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_Bucket__FOlD = &PB_Bucket{
    BucketId:  0 ,
    BucketName:  "" ,
    Server1Id:  0 ,
    Server2Id:  0 ,
    BackupServerId:  0 ,
    ContentObjectTypeId:  0 ,
    ContentObjectCount:  0 ,
    CreatedTime:  0 ,
}


type PB_Chat_Flat struct {
    ChatKey string
    RoomKey string
    RoomTypeEnumId int
    UserId int
    PeerUserId int
    GroupId int
    CreatedSe int
    StartMessageIdFrom int
    LastDeletedMessageId int
    LastSeenMessageId int
    LastMessageId int
    UpdatedMs int
}
//ToPB
func(m *PB_Chat)ToFlat() *PB_Chat_Flat {
r := &PB_Chat_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedSe:  int(m.CreatedSe) ,
    StartMessageIdFrom:  int(m.StartMessageIdFrom) ,
    LastDeletedMessageId:  int(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int(m.LastSeenMessageId) ,
    LastMessageId:  int(m.LastMessageId) ,
    UpdatedMs:  int(m.UpdatedMs) ,
}
return r
}
//ToPB
func(m *PB_Chat_Flat)ToPB() *PB_Chat {
r := &PB_Chat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedSe:  int32(m.CreatedSe) ,
    StartMessageIdFrom:  int64(m.StartMessageIdFrom) ,
    LastDeletedMessageId:  int64(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int64(m.LastSeenMessageId) ,
    LastMessageId:  int64(m.LastMessageId) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
}
return r
}
//folding
var PB_Chat__FOlD = &PB_Chat{
    ChatKey:  "" ,
    RoomKey:  "" ,
    RoomTypeEnumId:  0 ,
    UserId:  0 ,
    PeerUserId:  0 ,
    GroupId:  0 ,
    CreatedSe:  0 ,
    StartMessageIdFrom:  0 ,
    LastDeletedMessageId:  0 ,
    LastSeenMessageId:  0 ,
    LastMessageId:  0 ,
    UpdatedMs:  0 ,
}


type PB_Comment_Flat struct {
    Id int
    UserId int
    PostId int
    Text string
    CreatedTime int
}
//ToPB
func(m *PB_Comment)ToFlat() *PB_Comment_Flat {
r := &PB_Comment_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    Text:  m.Text ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_Comment_Flat)ToPB() *PB_Comment {
r := &PB_Comment{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    Text:  m.Text ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_Comment__FOlD = &PB_Comment{
    Id:  0 ,
    UserId:  0 ,
    PostId:  0 ,
    Text:  "" ,
    CreatedTime:  0 ,
}


type PB_DirectMessage_Flat struct {
    MessageId int
    MessageKey string
    RoomKey string
    UserId int
    MessageFileId int
    MessageTypeEnumId int
    Text string
    CreatedSe int
    PeerReceivedTime int
    PeerSeenTime int
    DeliviryStatusEnumId int
}
//ToPB
func(m *PB_DirectMessage)ToFlat() *PB_DirectMessage_Flat {
r := &PB_DirectMessage_Flat{
    MessageId:  int(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnumId:  int(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int(m.CreatedSe) ,
    PeerReceivedTime:  int(m.PeerReceivedTime) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int(m.DeliviryStatusEnumId) ,
}
return r
}
//ToPB
func(m *PB_DirectMessage_Flat)ToPB() *PB_DirectMessage {
r := &PB_DirectMessage{
    MessageId:  int64(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnumId:  int32(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int32(m.CreatedSe) ,
    PeerReceivedTime:  int32(m.PeerReceivedTime) ,
    PeerSeenTime:  int32(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int32(m.DeliviryStatusEnumId) ,
}
return r
}
//folding
var PB_DirectMessage__FOlD = &PB_DirectMessage{
    MessageId:  0 ,
    MessageKey:  "" ,
    RoomKey:  "" ,
    UserId:  0 ,
    MessageFileId:  0 ,
    MessageTypeEnumId:  0 ,
    Text:  "" ,
    CreatedSe:  0 ,
    PeerReceivedTime:  0 ,
    PeerSeenTime:  0 ,
    DeliviryStatusEnumId:  0 ,
}


type PB_DirectOffline_Flat struct {
    DirectOfflineId int
    ToUserId int
    MessageId int
    MessageFileId int
    OtherId int
    ChatKey string
    PeerUserId int
    RoomLogTypeId int
    PBClass string
    DataPB []byte
    DataJson string
    DataTemp string
    AtTimeMs int
}
//ToPB
func(m *PB_DirectOffline)ToFlat() *PB_DirectOffline_Flat {
r := &PB_DirectOffline_Flat{
    DirectOfflineId:  int(m.DirectOfflineId) ,
    ToUserId:  int(m.ToUserId) ,
    MessageId:  int(m.MessageId) ,
    MessageFileId:  int(m.MessageFileId) ,
    OtherId:  int(m.OtherId) ,
    ChatKey:  m.ChatKey ,
    PeerUserId:  int(m.PeerUserId) ,
    RoomLogTypeId:  int(m.RoomLogTypeId) ,
    PBClass:  m.PBClass ,
    DataPB:  []byte(m.DataPB) ,
    DataJson:  m.DataJson ,
    DataTemp:  m.DataTemp ,
    AtTimeMs:  int(m.AtTimeMs) ,
}
return r
}
//ToPB
func(m *PB_DirectOffline_Flat)ToPB() *PB_DirectOffline {
r := &PB_DirectOffline{
    DirectOfflineId:  int64(m.DirectOfflineId) ,
    ToUserId:  int32(m.ToUserId) ,
    MessageId:  int64(m.MessageId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    OtherId:  int64(m.OtherId) ,
    ChatKey:  m.ChatKey ,
    PeerUserId:  int32(m.PeerUserId) ,
    RoomLogTypeId:  int32(m.RoomLogTypeId) ,
    PBClass:  m.PBClass ,
    DataPB:  m.DataPB ,
    DataJson:  m.DataJson ,
    DataTemp:  m.DataTemp ,
    AtTimeMs:  int64(m.AtTimeMs) ,
}
return r
}
//folding
var PB_DirectOffline__FOlD = &PB_DirectOffline{
    DirectOfflineId:  0 ,
    ToUserId:  0 ,
    MessageId:  0 ,
    MessageFileId:  0 ,
    OtherId:  0 ,
    ChatKey:  "" ,
    PeerUserId:  0 ,
    RoomLogTypeId:  0 ,
    PBClass:  "" ,
    DataPB:  []byte{} ,
    DataJson:  "" ,
    DataTemp:  "" ,
    AtTimeMs:  0 ,
}


type PB_DirectToMessage_Flat struct {
    Id int
    ChatKey string
    MessageId int
    SourceEnumId int
}
//ToPB
func(m *PB_DirectToMessage)ToFlat() *PB_DirectToMessage_Flat {
r := &PB_DirectToMessage_Flat{
    Id:  int(m.Id) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int(m.MessageId) ,
    SourceEnumId:  int(m.SourceEnumId) ,
}
return r
}
//ToPB
func(m *PB_DirectToMessage_Flat)ToPB() *PB_DirectToMessage {
r := &PB_DirectToMessage{
    Id:  int64(m.Id) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int64(m.MessageId) ,
    SourceEnumId:  int32(m.SourceEnumId) ,
}
return r
}
//folding
var PB_DirectToMessage__FOlD = &PB_DirectToMessage{
    Id:  0 ,
    ChatKey:  "" ,
    MessageId:  0 ,
    SourceEnumId:  0 ,
}


type PB_DirectUpdate_Flat struct {
    DirectUpdateId int
    ToUserId int
    MessageId int
    MessageFileId int
    OtherId int
    ChatKey string
    PeerUserId int
    EventType int
    RoomLogTypeId int
    FromSeq int
    ToSeq int
    ExtraPB []byte
    ExtraJson string
    AtTimeMs int
}
//ToPB
func(m *PB_DirectUpdate)ToFlat() *PB_DirectUpdate_Flat {
r := &PB_DirectUpdate_Flat{
    DirectUpdateId:  int(m.DirectUpdateId) ,
    ToUserId:  int(m.ToUserId) ,
    MessageId:  int(m.MessageId) ,
    MessageFileId:  int(m.MessageFileId) ,
    OtherId:  int(m.OtherId) ,
    ChatKey:  m.ChatKey ,
    PeerUserId:  int(m.PeerUserId) ,
    EventType:  int(m.EventType) ,
    RoomLogTypeId:  int(m.RoomLogTypeId) ,
    FromSeq:  int(m.FromSeq) ,
    ToSeq:  int(m.ToSeq) ,
    ExtraPB:  []byte(m.ExtraPB) ,
    ExtraJson:  m.ExtraJson ,
    AtTimeMs:  int(m.AtTimeMs) ,
}
return r
}
//ToPB
func(m *PB_DirectUpdate_Flat)ToPB() *PB_DirectUpdate {
r := &PB_DirectUpdate{
    DirectUpdateId:  int64(m.DirectUpdateId) ,
    ToUserId:  int32(m.ToUserId) ,
    MessageId:  int64(m.MessageId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    OtherId:  int64(m.OtherId) ,
    ChatKey:  m.ChatKey ,
    PeerUserId:  int32(m.PeerUserId) ,
    EventType:  int32(m.EventType) ,
    RoomLogTypeId:  int32(m.RoomLogTypeId) ,
    FromSeq:  int32(m.FromSeq) ,
    ToSeq:  int32(m.ToSeq) ,
    ExtraPB:  m.ExtraPB ,
    ExtraJson:  m.ExtraJson ,
    AtTimeMs:  int64(m.AtTimeMs) ,
}
return r
}
//folding
var PB_DirectUpdate__FOlD = &PB_DirectUpdate{
    DirectUpdateId:  0 ,
    ToUserId:  0 ,
    MessageId:  0 ,
    MessageFileId:  0 ,
    OtherId:  0 ,
    ChatKey:  "" ,
    PeerUserId:  0 ,
    EventType:  0 ,
    RoomLogTypeId:  0 ,
    FromSeq:  0 ,
    ToSeq:  0 ,
    ExtraPB:  []byte{} ,
    ExtraJson:  "" ,
    AtTimeMs:  0 ,
}


type PB_FollowingList_Flat struct {
    Id int
    UserId int
    ListType int
    Name string
    Count int
    IsAuto int
    IsPimiry int
    CreatedTime int
}
//ToPB
func(m *PB_FollowingList)ToFlat() *PB_FollowingList_Flat {
r := &PB_FollowingList_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    ListType:  int(m.ListType) ,
    Name:  m.Name ,
    Count:  int(m.Count) ,
    IsAuto:  int(m.IsAuto) ,
    IsPimiry:  int(m.IsPimiry) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_FollowingList_Flat)ToPB() *PB_FollowingList {
r := &PB_FollowingList{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    ListType:  int32(m.ListType) ,
    Name:  m.Name ,
    Count:  int32(m.Count) ,
    IsAuto:  int32(m.IsAuto) ,
    IsPimiry:  int32(m.IsPimiry) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_FollowingList__FOlD = &PB_FollowingList{
    Id:  0 ,
    UserId:  0 ,
    ListType:  0 ,
    Name:  "" ,
    Count:  0 ,
    IsAuto:  0 ,
    IsPimiry:  0 ,
    CreatedTime:  0 ,
}


type PB_FollowingListMember_Flat struct {
    Id int
    ListId int
    UserId int
    FollowedUserId int
    FollowType int
    UpdatedTimeMs int
}
//ToPB
func(m *PB_FollowingListMember)ToFlat() *PB_FollowingListMember_Flat {
r := &PB_FollowingListMember_Flat{
    Id:  int(m.Id) ,
    ListId:  int(m.ListId) ,
    UserId:  int(m.UserId) ,
    FollowedUserId:  int(m.FollowedUserId) ,
    FollowType:  int(m.FollowType) ,
    UpdatedTimeMs:  int(m.UpdatedTimeMs) ,
}
return r
}
//ToPB
func(m *PB_FollowingListMember_Flat)ToPB() *PB_FollowingListMember {
r := &PB_FollowingListMember{
    Id:  int64(m.Id) ,
    ListId:  int32(m.ListId) ,
    UserId:  int32(m.UserId) ,
    FollowedUserId:  int32(m.FollowedUserId) ,
    FollowType:  int32(m.FollowType) ,
    UpdatedTimeMs:  int64(m.UpdatedTimeMs) ,
}
return r
}
//folding
var PB_FollowingListMember__FOlD = &PB_FollowingListMember{
    Id:  0 ,
    ListId:  0 ,
    UserId:  0 ,
    FollowedUserId:  0 ,
    FollowType:  0 ,
    UpdatedTimeMs:  0 ,
}


type PB_FollowingListMemberHistory_Flat struct {
    Id int
    ListId int
    UserId int
    FollowedUserId int
    FollowType int
    UpdatedTimeMs int
    FollowId int
}
//ToPB
func(m *PB_FollowingListMemberHistory)ToFlat() *PB_FollowingListMemberHistory_Flat {
r := &PB_FollowingListMemberHistory_Flat{
    Id:  int(m.Id) ,
    ListId:  int(m.ListId) ,
    UserId:  int(m.UserId) ,
    FollowedUserId:  int(m.FollowedUserId) ,
    FollowType:  int(m.FollowType) ,
    UpdatedTimeMs:  int(m.UpdatedTimeMs) ,
    FollowId:  int(m.FollowId) ,
}
return r
}
//ToPB
func(m *PB_FollowingListMemberHistory_Flat)ToPB() *PB_FollowingListMemberHistory {
r := &PB_FollowingListMemberHistory{
    Id:  int64(m.Id) ,
    ListId:  int32(m.ListId) ,
    UserId:  int32(m.UserId) ,
    FollowedUserId:  int32(m.FollowedUserId) ,
    FollowType:  int32(m.FollowType) ,
    UpdatedTimeMs:  int64(m.UpdatedTimeMs) ,
    FollowId:  int32(m.FollowId) ,
}
return r
}
//folding
var PB_FollowingListMemberHistory__FOlD = &PB_FollowingListMemberHistory{
    Id:  0 ,
    ListId:  0 ,
    UserId:  0 ,
    FollowedUserId:  0 ,
    FollowType:  0 ,
    UpdatedTimeMs:  0 ,
    FollowId:  0 ,
}


type PB_GeneralLog_Flat struct {
    Id int
    ToUserId int
    TargetId int
    LogTypeId int
    ExtraPB []byte
    ExtraJson string
    CreatedMs int
}
//ToPB
func(m *PB_GeneralLog)ToFlat() *PB_GeneralLog_Flat {
r := &PB_GeneralLog_Flat{
    Id:  int(m.Id) ,
    ToUserId:  int(m.ToUserId) ,
    TargetId:  int(m.TargetId) ,
    LogTypeId:  int(m.LogTypeId) ,
    ExtraPB:  []byte(m.ExtraPB) ,
    ExtraJson:  m.ExtraJson ,
    CreatedMs:  int(m.CreatedMs) ,
}
return r
}
//ToPB
func(m *PB_GeneralLog_Flat)ToPB() *PB_GeneralLog {
r := &PB_GeneralLog{
    Id:  int64(m.Id) ,
    ToUserId:  int32(m.ToUserId) ,
    TargetId:  int32(m.TargetId) ,
    LogTypeId:  int32(m.LogTypeId) ,
    ExtraPB:  m.ExtraPB ,
    ExtraJson:  m.ExtraJson ,
    CreatedMs:  int64(m.CreatedMs) ,
}
return r
}
//folding
var PB_GeneralLog__FOlD = &PB_GeneralLog{
    Id:  0 ,
    ToUserId:  0 ,
    TargetId:  0 ,
    LogTypeId:  0 ,
    ExtraPB:  []byte{} ,
    ExtraJson:  "" ,
    CreatedMs:  0 ,
}


type PB_Group_Flat struct {
    GroupId int
    GroupName string
    MembersCount int
    GroupPrivacyEnum int
    CreatorUserId int
    CreatedTime int
    UpdatedMs int
    CurrentSeq int
}
//ToPB
func(m *PB_Group)ToFlat() *PB_Group_Flat {
r := &PB_Group_Flat{
    GroupId:  int(m.GroupId) ,
    GroupName:  m.GroupName ,
    MembersCount:  int(m.MembersCount) ,
    GroupPrivacyEnum:  int(m.GroupPrivacyEnum) ,
    CreatorUserId:  int(m.CreatorUserId) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    CurrentSeq:  int(m.CurrentSeq) ,
}
return r
}
//ToPB
func(m *PB_Group_Flat)ToPB() *PB_Group {
r := &PB_Group{
    GroupId:  int64(m.GroupId) ,
    GroupName:  m.GroupName ,
    MembersCount:  int32(m.MembersCount) ,
    GroupPrivacyEnum:  int32(m.GroupPrivacyEnum) ,
    CreatorUserId:  int32(m.CreatorUserId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
}
return r
}
//folding
var PB_Group__FOlD = &PB_Group{
    GroupId:  0 ,
    GroupName:  "" ,
    MembersCount:  0 ,
    GroupPrivacyEnum:  0 ,
    CreatorUserId:  0 ,
    CreatedTime:  0 ,
    UpdatedMs:  0 ,
    CurrentSeq:  0 ,
}


type PB_GroupMember_Flat struct {
    Id int
    GroupId int
    GroupKey string
    UserId int
    ByUserId int
    GroupRoleEnumId int
    CreatedTime int
}
//ToPB
func(m *PB_GroupMember)ToFlat() *PB_GroupMember_Flat {
r := &PB_GroupMember_Flat{
    Id:  int(m.Id) ,
    GroupId:  int(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    UserId:  int(m.UserId) ,
    ByUserId:  int(m.ByUserId) ,
    GroupRoleEnumId:  int(m.GroupRoleEnumId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_GroupMember_Flat)ToPB() *PB_GroupMember {
r := &PB_GroupMember{
    Id:  int64(m.Id) ,
    GroupId:  int64(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    UserId:  int32(m.UserId) ,
    ByUserId:  int32(m.ByUserId) ,
    GroupRoleEnumId:  int32(m.GroupRoleEnumId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_GroupMember__FOlD = &PB_GroupMember{
    Id:  0 ,
    GroupId:  0 ,
    GroupKey:  "" ,
    UserId:  0 ,
    ByUserId:  0 ,
    GroupRoleEnumId:  0 ,
    CreatedTime:  0 ,
}


type PB_GroupMessage_Flat struct {
    MessageId int
    RoomKey string
    UserId int
    MessageFileId int
    MessageTypeEnum int
    Text string
    CreatedMs int
    DeliveryStatusEnum int
}
//ToPB
func(m *PB_GroupMessage)ToFlat() *PB_GroupMessage_Flat {
r := &PB_GroupMessage_Flat{
    MessageId:  int(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnum:  int(m.MessageTypeEnum) ,
    Text:  m.Text ,
    CreatedMs:  int(m.CreatedMs) ,
    DeliveryStatusEnum:  int(m.DeliveryStatusEnum) ,
}
return r
}
//ToPB
func(m *PB_GroupMessage_Flat)ToPB() *PB_GroupMessage {
r := &PB_GroupMessage{
    MessageId:  int64(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnum:  int32(m.MessageTypeEnum) ,
    Text:  m.Text ,
    CreatedMs:  int64(m.CreatedMs) ,
    DeliveryStatusEnum:  int32(m.DeliveryStatusEnum) ,
}
return r
}
//folding
var PB_GroupMessage__FOlD = &PB_GroupMessage{
    MessageId:  0 ,
    RoomKey:  "" ,
    UserId:  0 ,
    MessageFileId:  0 ,
    MessageTypeEnum:  0 ,
    Text:  "" ,
    CreatedMs:  0 ,
    DeliveryStatusEnum:  0 ,
}


type PB_GroupToMessage_Flat struct {
    Id int
    GroupId int
    MessageId int
    CreatedMs int
    StatusEnum int
}
//ToPB
func(m *PB_GroupToMessage)ToFlat() *PB_GroupToMessage_Flat {
r := &PB_GroupToMessage_Flat{
    Id:  int(m.Id) ,
    GroupId:  int(m.GroupId) ,
    MessageId:  int(m.MessageId) ,
    CreatedMs:  int(m.CreatedMs) ,
    StatusEnum:  int(m.StatusEnum) ,
}
return r
}
//ToPB
func(m *PB_GroupToMessage_Flat)ToPB() *PB_GroupToMessage {
r := &PB_GroupToMessage{
    Id:  int64(m.Id) ,
    GroupId:  int64(m.GroupId) ,
    MessageId:  int64(m.MessageId) ,
    CreatedMs:  int64(m.CreatedMs) ,
    StatusEnum:  int32(m.StatusEnum) ,
}
return r
}
//folding
var PB_GroupToMessage__FOlD = &PB_GroupToMessage{
    Id:  0 ,
    GroupId:  0 ,
    MessageId:  0 ,
    CreatedMs:  0 ,
    StatusEnum:  0 ,
}


type PB_Like_Flat struct {
    Id int
    PostId int
    PostTypeId int
    UserId int
    TypeId int
    CreatedTime int
}
//ToPB
func(m *PB_Like)ToFlat() *PB_Like_Flat {
r := &PB_Like_Flat{
    Id:  int(m.Id) ,
    PostId:  int(m.PostId) ,
    PostTypeId:  int(m.PostTypeId) ,
    UserId:  int(m.UserId) ,
    TypeId:  int(m.TypeId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_Like_Flat)ToPB() *PB_Like {
r := &PB_Like{
    Id:  int32(m.Id) ,
    PostId:  int32(m.PostId) ,
    PostTypeId:  int32(m.PostTypeId) ,
    UserId:  int32(m.UserId) ,
    TypeId:  int32(m.TypeId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_Like__FOlD = &PB_Like{
    Id:  0 ,
    PostId:  0 ,
    PostTypeId:  0 ,
    UserId:  0 ,
    TypeId:  0 ,
    CreatedTime:  0 ,
}


type PB_LogChange_Flat struct {
    Id int
    T string
}
//ToPB
func(m *PB_LogChange)ToFlat() *PB_LogChange_Flat {
r := &PB_LogChange_Flat{
    Id:  int(m.Id) ,
    T:  m.T ,
}
return r
}
//ToPB
func(m *PB_LogChange_Flat)ToPB() *PB_LogChange {
r := &PB_LogChange{
    Id:  int32(m.Id) ,
    T:  m.T ,
}
return r
}
//folding
var PB_LogChange__FOlD = &PB_LogChange{
    Id:  0 ,
    T:  "" ,
}


type PB_Media_Flat struct {
    Id int
    UserId int
    PostId int
    AlbumId int
    TypeId int
    CreatedTime int
    Src string
}
//ToPB
func(m *PB_Media)ToFlat() *PB_Media_Flat {
r := &PB_Media_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    AlbumId:  int(m.AlbumId) ,
    TypeId:  int(m.TypeId) ,
    CreatedTime:  int(m.CreatedTime) ,
    Src:  m.Src ,
}
return r
}
//ToPB
func(m *PB_Media_Flat)ToPB() *PB_Media {
r := &PB_Media{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    AlbumId:  int32(m.AlbumId) ,
    TypeId:  int32(m.TypeId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    Src:  m.Src ,
}
return r
}
//folding
var PB_Media__FOlD = &PB_Media{
    Id:  0 ,
    UserId:  0 ,
    PostId:  0 ,
    AlbumId:  0 ,
    TypeId:  0 ,
    CreatedTime:  0 ,
    Src:  "" ,
}


type PB_MessageFile_Flat struct {
    MessageFileId int
    MessageFileKey string
    OriginalUserId int
    Name string
    Size int
    FileTypeEnumId int
    Width int
    Height int
    Duration int
    Extension string
    HashMd5 string
    HashAccess int
    CreatedSe int
    ServerSrc string
    ServerPath string
    ServerThumbPath string
    BucketId string
    ServerId int
    CanDel int
}
//ToPB
func(m *PB_MessageFile)ToFlat() *PB_MessageFile_Flat {
r := &PB_MessageFile_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileTypeEnumId:  int(m.FileTypeEnumId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int(m.HashAccess) ,
    CreatedSe:  int(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
}
return r
}
//ToPB
func(m *PB_MessageFile_Flat)ToPB() *PB_MessageFile {
r := &PB_MessageFile{
    MessageFileId:  int64(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int32(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileTypeEnumId:  int32(m.FileTypeEnumId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int64(m.HashAccess) ,
    CreatedSe:  int32(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
}
return r
}
//folding
var PB_MessageFile__FOlD = &PB_MessageFile{
    MessageFileId:  0 ,
    MessageFileKey:  "" ,
    OriginalUserId:  0 ,
    Name:  "" ,
    Size:  0 ,
    FileTypeEnumId:  0 ,
    Width:  0 ,
    Height:  0 ,
    Duration:  0 ,
    Extension:  "" ,
    HashMd5:  "" ,
    HashAccess:  0 ,
    CreatedSe:  0 ,
    ServerSrc:  "" ,
    ServerPath:  "" ,
    ServerThumbPath:  "" ,
    BucketId:  "" ,
    ServerId:  0 ,
    CanDel:  0 ,
}


type PB_Notification_Flat struct {
    Id int
    ForUserId int
    ActorUserId int
    ActionTypeId int
    ObjectTypeId int
    RowId int
    RootId int
    RefId int
    SeenStatus int
    CreatedTime int
}
//ToPB
func(m *PB_Notification)ToFlat() *PB_Notification_Flat {
r := &PB_Notification_Flat{
    Id:  int(m.Id) ,
    ForUserId:  int(m.ForUserId) ,
    ActorUserId:  int(m.ActorUserId) ,
    ActionTypeId:  int(m.ActionTypeId) ,
    ObjectTypeId:  int(m.ObjectTypeId) ,
    RowId:  int(m.RowId) ,
    RootId:  int(m.RootId) ,
    RefId:  int(m.RefId) ,
    SeenStatus:  int(m.SeenStatus) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_Notification_Flat)ToPB() *PB_Notification {
r := &PB_Notification{
    Id:  int64(m.Id) ,
    ForUserId:  int32(m.ForUserId) ,
    ActorUserId:  int32(m.ActorUserId) ,
    ActionTypeId:  int32(m.ActionTypeId) ,
    ObjectTypeId:  int32(m.ObjectTypeId) ,
    RowId:  int32(m.RowId) ,
    RootId:  int32(m.RootId) ,
    RefId:  int32(m.RefId) ,
    SeenStatus:  int32(m.SeenStatus) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_Notification__FOlD = &PB_Notification{
    Id:  0 ,
    ForUserId:  0 ,
    ActorUserId:  0 ,
    ActionTypeId:  0 ,
    ObjectTypeId:  0 ,
    RowId:  0 ,
    RootId:  0 ,
    RefId:  0 ,
    SeenStatus:  0 ,
    CreatedTime:  0 ,
}


type PB_NotificationRemoved_Flat struct {
    NotificationId int
    ForUserId int
}
//ToPB
func(m *PB_NotificationRemoved)ToFlat() *PB_NotificationRemoved_Flat {
r := &PB_NotificationRemoved_Flat{
    NotificationId:  int(m.NotificationId) ,
    ForUserId:  int(m.ForUserId) ,
}
return r
}
//ToPB
func(m *PB_NotificationRemoved_Flat)ToPB() *PB_NotificationRemoved {
r := &PB_NotificationRemoved{
    NotificationId:  int32(m.NotificationId) ,
    ForUserId:  int32(m.ForUserId) ,
}
return r
}
//folding
var PB_NotificationRemoved__FOlD = &PB_NotificationRemoved{
    NotificationId:  0 ,
    ForUserId:  0 ,
}


type PB_Offline_Flat struct {
    Id int
    FromUserId int
    ToUserId int
    RpcName string
    PBClassName string
    Key string
    DataJson string
    DataBlob []byte
    DataLength int
    CreatedMs int
}
//ToPB
func(m *PB_Offline)ToFlat() *PB_Offline_Flat {
r := &PB_Offline_Flat{
    Id:  int(m.Id) ,
    FromUserId:  int(m.FromUserId) ,
    ToUserId:  int(m.ToUserId) ,
    RpcName:  m.RpcName ,
    PBClassName:  m.PBClassName ,
    Key:  m.Key ,
    DataJson:  m.DataJson ,
    DataBlob:  []byte(m.DataBlob) ,
    DataLength:  int(m.DataLength) ,
    CreatedMs:  int(m.CreatedMs) ,
}
return r
}
//ToPB
func(m *PB_Offline_Flat)ToPB() *PB_Offline {
r := &PB_Offline{
    Id:  int64(m.Id) ,
    FromUserId:  int32(m.FromUserId) ,
    ToUserId:  int32(m.ToUserId) ,
    RpcName:  m.RpcName ,
    PBClassName:  m.PBClassName ,
    Key:  m.Key ,
    DataJson:  m.DataJson ,
    DataBlob:  m.DataBlob ,
    DataLength:  int32(m.DataLength) ,
    CreatedMs:  int64(m.CreatedMs) ,
}
return r
}
//folding
var PB_Offline__FOlD = &PB_Offline{
    Id:  0 ,
    FromUserId:  0 ,
    ToUserId:  0 ,
    RpcName:  "" ,
    PBClassName:  "" ,
    Key:  "" ,
    DataJson:  "" ,
    DataBlob:  []byte{} ,
    DataLength:  0 ,
    CreatedMs:  0 ,
}


type PB_OldMessage_Flat struct {
    Id int
    Uid int
    UserId int
    MessageKey string
    RoomKey string
    MessageType int
    RoomType int
    MsgFileId int
    DataPB []byte
    Data64 string
    DataJson string
    CreatedTimeMs int
}
//ToPB
func(m *PB_OldMessage)ToFlat() *PB_OldMessage_Flat {
r := &PB_OldMessage_Flat{
    Id:  int(m.Id) ,
    Uid:  int(m.Uid) ,
    UserId:  int(m.UserId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    MessageType:  int(m.MessageType) ,
    RoomType:  int(m.RoomType) ,
    MsgFileId:  int(m.MsgFileId) ,
    DataPB:  []byte(m.DataPB) ,
    Data64:  m.Data64 ,
    DataJson:  m.DataJson ,
    CreatedTimeMs:  int(m.CreatedTimeMs) ,
}
return r
}
//ToPB
func(m *PB_OldMessage_Flat)ToPB() *PB_OldMessage {
r := &PB_OldMessage{
    Id:  int64(m.Id) ,
    Uid:  int64(m.Uid) ,
    UserId:  int64(m.UserId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    MessageType:  int32(m.MessageType) ,
    RoomType:  int32(m.RoomType) ,
    MsgFileId:  int64(m.MsgFileId) ,
    DataPB:  m.DataPB ,
    Data64:  m.Data64 ,
    DataJson:  m.DataJson ,
    CreatedTimeMs:  int64(m.CreatedTimeMs) ,
}
return r
}
//folding
var PB_OldMessage__FOlD = &PB_OldMessage{
    Id:  0 ,
    Uid:  0 ,
    UserId:  0 ,
    MessageKey:  "" ,
    RoomKey:  "" ,
    MessageType:  0 ,
    RoomType:  0 ,
    MsgFileId:  0 ,
    DataPB:  []byte{} ,
    Data64:  "" ,
    DataJson:  "" ,
    CreatedTimeMs:  0 ,
}


type PB_OldMsgFile_Flat struct {
    Id int
    Name string
    Size int
    FileType int
    MimeType string
    Width int
    Height int
    Duration int
    Extension string
    ThumbData []byte
    ThumbData64 string
    ServerSrc string
    ServerPath string
    ServerId int
    CanDel int
}
//ToPB
func(m *PB_OldMsgFile)ToFlat() *PB_OldMsgFile_Flat {
r := &PB_OldMsgFile_Flat{
    Id:  int(m.Id) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileType:  int(m.FileType) ,
    MimeType:  m.MimeType ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    ThumbData:  []byte(m.ThumbData) ,
    ThumbData64:  m.ThumbData64 ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
}
return r
}
//ToPB
func(m *PB_OldMsgFile_Flat)ToPB() *PB_OldMsgFile {
r := &PB_OldMsgFile{
    Id:  int64(m.Id) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileType:  int32(m.FileType) ,
    MimeType:  m.MimeType ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    ThumbData:  m.ThumbData ,
    ThumbData64:  m.ThumbData64 ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
}
return r
}
//folding
var PB_OldMsgFile__FOlD = &PB_OldMsgFile{
    Id:  0 ,
    Name:  "" ,
    Size:  0 ,
    FileType:  0 ,
    MimeType:  "" ,
    Width:  0 ,
    Height:  0 ,
    Duration:  0 ,
    Extension:  "" ,
    ThumbData:  []byte{} ,
    ThumbData64:  "" ,
    ServerSrc:  "" ,
    ServerPath:  "" ,
    ServerId:  0 ,
    CanDel:  0 ,
}


type PB_OldMsgPush_Flat struct {
    Id int
    Uid int
    ToUser int
    MsgUid int
    CreatedTimeMs int
}
//ToPB
func(m *PB_OldMsgPush)ToFlat() *PB_OldMsgPush_Flat {
r := &PB_OldMsgPush_Flat{
    Id:  int(m.Id) ,
    Uid:  int(m.Uid) ,
    ToUser:  int(m.ToUser) ,
    MsgUid:  int(m.MsgUid) ,
    CreatedTimeMs:  int(m.CreatedTimeMs) ,
}
return r
}
//ToPB
func(m *PB_OldMsgPush_Flat)ToPB() *PB_OldMsgPush {
r := &PB_OldMsgPush{
    Id:  int64(m.Id) ,
    Uid:  int64(m.Uid) ,
    ToUser:  int64(m.ToUser) ,
    MsgUid:  int64(m.MsgUid) ,
    CreatedTimeMs:  int64(m.CreatedTimeMs) ,
}
return r
}
//folding
var PB_OldMsgPush__FOlD = &PB_OldMsgPush{
    Id:  0 ,
    Uid:  0 ,
    ToUser:  0 ,
    MsgUid:  0 ,
    CreatedTimeMs:  0 ,
}


type PB_OldMsgPushEvent_Flat struct {
    Id int
    Uid int
    ToUserId int
    MsgUid int
    MsgKey string
    RoomKey string
    PeerUserId int
    EventType int
    AtTime int
}
//ToPB
func(m *PB_OldMsgPushEvent)ToFlat() *PB_OldMsgPushEvent_Flat {
r := &PB_OldMsgPushEvent_Flat{
    Id:  int(m.Id) ,
    Uid:  int(m.Uid) ,
    ToUserId:  int(m.ToUserId) ,
    MsgUid:  int(m.MsgUid) ,
    MsgKey:  m.MsgKey ,
    RoomKey:  m.RoomKey ,
    PeerUserId:  int(m.PeerUserId) ,
    EventType:  int(m.EventType) ,
    AtTime:  int(m.AtTime) ,
}
return r
}
//ToPB
func(m *PB_OldMsgPushEvent_Flat)ToPB() *PB_OldMsgPushEvent {
r := &PB_OldMsgPushEvent{
    Id:  int64(m.Id) ,
    Uid:  int64(m.Uid) ,
    ToUserId:  int32(m.ToUserId) ,
    MsgUid:  int64(m.MsgUid) ,
    MsgKey:  m.MsgKey ,
    RoomKey:  m.RoomKey ,
    PeerUserId:  int32(m.PeerUserId) ,
    EventType:  int32(m.EventType) ,
    AtTime:  int32(m.AtTime) ,
}
return r
}
//folding
var PB_OldMsgPushEvent__FOlD = &PB_OldMsgPushEvent{
    Id:  0 ,
    Uid:  0 ,
    ToUserId:  0 ,
    MsgUid:  0 ,
    MsgKey:  "" ,
    RoomKey:  "" ,
    PeerUserId:  0 ,
    EventType:  0 ,
    AtTime:  0 ,
}


type PB_PhoneContact_Flat struct {
    Id int
    PhoneDisplayName string
    PhoneFamilyName string
    PhoneNumber string
    PhoneNormalizedNumber string
    PhoneContactRowId int
    UserId int
    DeviceUuidId int
    CreatedTime int
    UpdatedTime int
}
//ToPB
func(m *PB_PhoneContact)ToFlat() *PB_PhoneContact_Flat {
r := &PB_PhoneContact_Flat{
    Id:  int(m.Id) ,
    PhoneDisplayName:  m.PhoneDisplayName ,
    PhoneFamilyName:  m.PhoneFamilyName ,
    PhoneNumber:  m.PhoneNumber ,
    PhoneNormalizedNumber:  m.PhoneNormalizedNumber ,
    PhoneContactRowId:  int(m.PhoneContactRowId) ,
    UserId:  int(m.UserId) ,
    DeviceUuidId:  int(m.DeviceUuidId) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedTime:  int(m.UpdatedTime) ,
}
return r
}
//ToPB
func(m *PB_PhoneContact_Flat)ToPB() *PB_PhoneContact {
r := &PB_PhoneContact{
    Id:  int32(m.Id) ,
    PhoneDisplayName:  m.PhoneDisplayName ,
    PhoneFamilyName:  m.PhoneFamilyName ,
    PhoneNumber:  m.PhoneNumber ,
    PhoneNormalizedNumber:  m.PhoneNormalizedNumber ,
    PhoneContactRowId:  int32(m.PhoneContactRowId) ,
    UserId:  int32(m.UserId) ,
    DeviceUuidId:  int32(m.DeviceUuidId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
}
return r
}
//folding
var PB_PhoneContact__FOlD = &PB_PhoneContact{
    Id:  0 ,
    PhoneDisplayName:  "" ,
    PhoneFamilyName:  "" ,
    PhoneNumber:  "" ,
    PhoneNormalizedNumber:  "" ,
    PhoneContactRowId:  0 ,
    UserId:  0 ,
    DeviceUuidId:  0 ,
    CreatedTime:  0 ,
    UpdatedTime:  0 ,
}


type PB_Photo_Flat struct {
    PhotoId int
    UserId int
    PostId int
    AlbumId int
    ImageTypeId int
    Title string
    Src string
    PathSrc string
    BucketId int
    Width int
    Height int
    Ratio float32
    HashMd5 string
    Color string
    CreatedTime int
    W1080 int
    W720 int
    W480 int
    W320 int
    W160 int
    W80 int
}
//ToPB
func(m *PB_Photo)ToFlat() *PB_Photo_Flat {
r := &PB_Photo_Flat{
    PhotoId:  int(m.PhotoId) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    AlbumId:  int(m.AlbumId) ,
    ImageTypeId:  int(m.ImageTypeId) ,
    Title:  m.Title ,
    Src:  m.Src ,
    PathSrc:  m.PathSrc ,
    BucketId:  int(m.BucketId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Ratio:  float32(m.Ratio) ,
    HashMd5:  m.HashMd5 ,
    Color:  m.Color ,
    CreatedTime:  int(m.CreatedTime) ,
    W1080:  int(m.W1080) ,
    W720:  int(m.W720) ,
    W480:  int(m.W480) ,
    W320:  int(m.W320) ,
    W160:  int(m.W160) ,
    W80:  int(m.W80) ,
}
return r
}
//ToPB
func(m *PB_Photo_Flat)ToPB() *PB_Photo {
r := &PB_Photo{
    PhotoId:  int32(m.PhotoId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    AlbumId:  int32(m.AlbumId) ,
    ImageTypeId:  int32(m.ImageTypeId) ,
    Title:  m.Title ,
    Src:  m.Src ,
    PathSrc:  m.PathSrc ,
    BucketId:  int32(m.BucketId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Ratio:  m.Ratio ,
    HashMd5:  m.HashMd5 ,
    Color:  m.Color ,
    CreatedTime:  int32(m.CreatedTime) ,
    W1080:  int32(m.W1080) ,
    W720:  int32(m.W720) ,
    W480:  int32(m.W480) ,
    W320:  int32(m.W320) ,
    W160:  int32(m.W160) ,
    W80:  int32(m.W80) ,
}
return r
}
//folding
var PB_Photo__FOlD = &PB_Photo{
    PhotoId:  0 ,
    UserId:  0 ,
    PostId:  0 ,
    AlbumId:  0 ,
    ImageTypeId:  0 ,
    Title:  "" ,
    Src:  "" ,
    PathSrc:  "" ,
    BucketId:  0 ,
    Width:  0 ,
    Height:  0 ,
    Ratio:  0.0 ,
    HashMd5:  "" ,
    Color:  "" ,
    CreatedTime:  0 ,
    W1080:  0 ,
    W720:  0 ,
    W480:  0 ,
    W320:  0 ,
    W160:  0 ,
    W80:  0 ,
}


type PB_Post_Flat struct {
    Id int
    UserId int
    TypeId int
    Text string
    FormatedText string
    MediaCount int
    SharedTo int
    DisableComment int
    HasTag int
    LikesCount int
    CommentsCount int
    EditedTime int
    CreatedTime int
}
//ToPB
func(m *PB_Post)ToFlat() *PB_Post_Flat {
r := &PB_Post_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    TypeId:  int(m.TypeId) ,
    Text:  m.Text ,
    FormatedText:  m.FormatedText ,
    MediaCount:  int(m.MediaCount) ,
    SharedTo:  int(m.SharedTo) ,
    DisableComment:  int(m.DisableComment) ,
    HasTag:  int(m.HasTag) ,
    LikesCount:  int(m.LikesCount) ,
    CommentsCount:  int(m.CommentsCount) ,
    EditedTime:  int(m.EditedTime) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_Post_Flat)ToPB() *PB_Post {
r := &PB_Post{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    TypeId:  int32(m.TypeId) ,
    Text:  m.Text ,
    FormatedText:  m.FormatedText ,
    MediaCount:  int32(m.MediaCount) ,
    SharedTo:  int32(m.SharedTo) ,
    DisableComment:  int32(m.DisableComment) ,
    HasTag:  int32(m.HasTag) ,
    LikesCount:  int32(m.LikesCount) ,
    CommentsCount:  int32(m.CommentsCount) ,
    EditedTime:  int32(m.EditedTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_Post__FOlD = &PB_Post{
    Id:  0 ,
    UserId:  0 ,
    TypeId:  0 ,
    Text:  "" ,
    FormatedText:  "" ,
    MediaCount:  0 ,
    SharedTo:  0 ,
    DisableComment:  0 ,
    HasTag:  0 ,
    LikesCount:  0 ,
    CommentsCount:  0 ,
    EditedTime:  0 ,
    CreatedTime:  0 ,
}


type PB_PushEvent_Flat struct {
    PushEventId int
    ToUserId int
    ToDeviceId int
    MessageId int
    RoomTypeEnum int
    RoomId int
    PeerUserId int
    PushEventTypeEnum int
    AtTime int
}
//ToPB
func(m *PB_PushEvent)ToFlat() *PB_PushEvent_Flat {
r := &PB_PushEvent_Flat{
    PushEventId:  int(m.PushEventId) ,
    ToUserId:  int(m.ToUserId) ,
    ToDeviceId:  int(m.ToDeviceId) ,
    MessageId:  int(m.MessageId) ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
    RoomId:  int(m.RoomId) ,
    PeerUserId:  int(m.PeerUserId) ,
    PushEventTypeEnum:  int(m.PushEventTypeEnum) ,
    AtTime:  int(m.AtTime) ,
}
return r
}
//ToPB
func(m *PB_PushEvent_Flat)ToPB() *PB_PushEvent {
r := &PB_PushEvent{
    PushEventId:  int64(m.PushEventId) ,
    ToUserId:  int32(m.ToUserId) ,
    ToDeviceId:  int64(m.ToDeviceId) ,
    MessageId:  int64(m.MessageId) ,
    RoomTypeEnum:  int32(m.RoomTypeEnum) ,
    RoomId:  int64(m.RoomId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    PushEventTypeEnum:  int32(m.PushEventTypeEnum) ,
    AtTime:  int32(m.AtTime) ,
}
return r
}
//folding
var PB_PushEvent__FOlD = &PB_PushEvent{
    PushEventId:  0 ,
    ToUserId:  0 ,
    ToDeviceId:  0 ,
    MessageId:  0 ,
    RoomTypeEnum:  0 ,
    RoomId:  0 ,
    PeerUserId:  0 ,
    PushEventTypeEnum:  0 ,
    AtTime:  0 ,
}


type PB_PushMessage_Flat struct {
    PushMessageId int
    ToUserId int
    ToDeviceId int
    MessageId int
    RoomTypeEnum int
    CreatedMs int
}
//ToPB
func(m *PB_PushMessage)ToFlat() *PB_PushMessage_Flat {
r := &PB_PushMessage_Flat{
    PushMessageId:  int(m.PushMessageId) ,
    ToUserId:  int(m.ToUserId) ,
    ToDeviceId:  int(m.ToDeviceId) ,
    MessageId:  int(m.MessageId) ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
    CreatedMs:  int(m.CreatedMs) ,
}
return r
}
//ToPB
func(m *PB_PushMessage_Flat)ToPB() *PB_PushMessage {
r := &PB_PushMessage{
    PushMessageId:  int64(m.PushMessageId) ,
    ToUserId:  int32(m.ToUserId) ,
    ToDeviceId:  int64(m.ToDeviceId) ,
    MessageId:  int64(m.MessageId) ,
    RoomTypeEnum:  int32(m.RoomTypeEnum) ,
    CreatedMs:  int64(m.CreatedMs) ,
}
return r
}
//folding
var PB_PushMessage__FOlD = &PB_PushMessage{
    PushMessageId:  0 ,
    ToUserId:  0 ,
    ToDeviceId:  0 ,
    MessageId:  0 ,
    RoomTypeEnum:  0 ,
    CreatedMs:  0 ,
}


type PB_RecommendUser_Flat struct {
    Id int
    UserId int
    TargetId int
    Weight float32
    CreatedTime int
}
//ToPB
func(m *PB_RecommendUser)ToFlat() *PB_RecommendUser_Flat {
r := &PB_RecommendUser_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    TargetId:  int(m.TargetId) ,
    Weight:  float32(m.Weight) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_RecommendUser_Flat)ToPB() *PB_RecommendUser {
r := &PB_RecommendUser{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    TargetId:  int32(m.TargetId) ,
    Weight:  m.Weight ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_RecommendUser__FOlD = &PB_RecommendUser{
    Id:  0 ,
    UserId:  0 ,
    TargetId:  0 ,
    Weight:  0.0 ,
    CreatedTime:  0 ,
}


type PB_Room_Flat struct {
    RoomId int
    RoomKey string
    RoomTypeEnum int
    UserId int
    LastSeqSeen int
    LastSeqDelete int
    PeerUserId int
    GroupId int
    CreatedTime int
    CurrentSeq int
}
//ToPB
func(m *PB_Room)ToFlat() *PB_Room_Flat {
r := &PB_Room_Flat{
    RoomId:  int(m.RoomId) ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
    UserId:  int(m.UserId) ,
    LastSeqSeen:  int(m.LastSeqSeen) ,
    LastSeqDelete:  int(m.LastSeqDelete) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedTime:  int(m.CreatedTime) ,
    CurrentSeq:  int(m.CurrentSeq) ,
}
return r
}
//ToPB
func(m *PB_Room_Flat)ToPB() *PB_Room {
r := &PB_Room{
    RoomId:  int64(m.RoomId) ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnum:  int32(m.RoomTypeEnum) ,
    UserId:  int32(m.UserId) ,
    LastSeqSeen:  int32(m.LastSeqSeen) ,
    LastSeqDelete:  int32(m.LastSeqDelete) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
}
return r
}
//folding
var PB_Room__FOlD = &PB_Room{
    RoomId:  0 ,
    RoomKey:  "" ,
    RoomTypeEnum:  0 ,
    UserId:  0 ,
    LastSeqSeen:  0 ,
    LastSeqDelete:  0 ,
    PeerUserId:  0 ,
    GroupId:  0 ,
    CreatedTime:  0 ,
    CurrentSeq:  0 ,
}


type PB_SearchClicked_Flat struct {
    Id int
    Query string
    ClickType int
    TargetId int
    UserId int
    CreatedAt int
}
//ToPB
func(m *PB_SearchClicked)ToFlat() *PB_SearchClicked_Flat {
r := &PB_SearchClicked_Flat{
    Id:  int(m.Id) ,
    Query:  m.Query ,
    ClickType:  int(m.ClickType) ,
    TargetId:  int(m.TargetId) ,
    UserId:  int(m.UserId) ,
    CreatedAt:  int(m.CreatedAt) ,
}
return r
}
//ToPB
func(m *PB_SearchClicked_Flat)ToPB() *PB_SearchClicked {
r := &PB_SearchClicked{
    Id:  int64(m.Id) ,
    Query:  m.Query ,
    ClickType:  int32(m.ClickType) ,
    TargetId:  int32(m.TargetId) ,
    UserId:  int32(m.UserId) ,
    CreatedAt:  int32(m.CreatedAt) ,
}
return r
}
//folding
var PB_SearchClicked__FOlD = &PB_SearchClicked{
    Id:  0 ,
    Query:  "" ,
    ClickType:  0 ,
    TargetId:  0 ,
    UserId:  0 ,
    CreatedAt:  0 ,
}


type PB_Session_Flat struct {
    Id int
    UserId int
    SessionUuid string
    ClientUuid string
    DeviceUuid string
    LastActivityTime int
    LastIpAddress string
    LastWifiMacAddress string
    LastNetworkType string
    LastNetworkTypeId int
    AppVersion int
    UpdatedTime int
    CreatedTime int
    LastSyncDirectUpdateId int
    LastSyncGeneralUpdateId int
    LastSyncNotifyUpdateId int
}
//ToPB
func(m *PB_Session)ToFlat() *PB_Session_Flat {
r := &PB_Session_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    SessionUuid:  m.SessionUuid ,
    ClientUuid:  m.ClientUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastActivityTime:  int(m.LastActivityTime) ,
    LastIpAddress:  m.LastIpAddress ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    LastNetworkTypeId:  int(m.LastNetworkTypeId) ,
    AppVersion:  int(m.AppVersion) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    LastSyncDirectUpdateId:  int(m.LastSyncDirectUpdateId) ,
    LastSyncGeneralUpdateId:  int(m.LastSyncGeneralUpdateId) ,
    LastSyncNotifyUpdateId:  int(m.LastSyncNotifyUpdateId) ,
}
return r
}
//ToPB
func(m *PB_Session_Flat)ToPB() *PB_Session {
r := &PB_Session{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    SessionUuid:  m.SessionUuid ,
    ClientUuid:  m.ClientUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    LastIpAddress:  m.LastIpAddress ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    LastNetworkTypeId:  int32(m.LastNetworkTypeId) ,
    AppVersion:  int32(m.AppVersion) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    LastSyncDirectUpdateId:  int64(m.LastSyncDirectUpdateId) ,
    LastSyncGeneralUpdateId:  int64(m.LastSyncGeneralUpdateId) ,
    LastSyncNotifyUpdateId:  int64(m.LastSyncNotifyUpdateId) ,
}
return r
}
//folding
var PB_Session__FOlD = &PB_Session{
    Id:  0 ,
    UserId:  0 ,
    SessionUuid:  "" ,
    ClientUuid:  "" ,
    DeviceUuid:  "" ,
    LastActivityTime:  0 ,
    LastIpAddress:  "" ,
    LastWifiMacAddress:  "" ,
    LastNetworkType:  "" ,
    LastNetworkTypeId:  0 ,
    AppVersion:  0 ,
    UpdatedTime:  0 ,
    CreatedTime:  0 ,
    LastSyncDirectUpdateId:  0 ,
    LastSyncGeneralUpdateId:  0 ,
    LastSyncNotifyUpdateId:  0 ,
}


type PB_SettingClient_Flat struct {
    UserId int
    AutoDownloadWifiVoice int
    AutoDownloadWifiImage int
    AutoDownloadWifiVideo int
    AutoDownloadWifiFile int
    AutoDownloadWifiMusic int
    AutoDownloadWifiGif int
    AutoDownloadCellularVoice int
    AutoDownloadCellularImage int
    AutoDownloadCellularVideo int
    AutoDownloadCellularFile int
    AutoDownloadCellularMusic int
    AutoDownloadCellularGif int
    AutoDownloadRoamingVoice int
    AutoDownloadRoamingImage int
    AutoDownloadRoamingVideo int
    AutoDownloadRoamingFile int
    AutoDownloadRoamingMusic int
    AutoDownloadRoamingGif int
    SaveToGallery int
}
//ToPB
func(m *PB_SettingClient)ToFlat() *PB_SettingClient_Flat {
r := &PB_SettingClient_Flat{
    UserId:  int(m.UserId) ,
    AutoDownloadWifiVoice:  int(m.AutoDownloadWifiVoice) ,
    AutoDownloadWifiImage:  int(m.AutoDownloadWifiImage) ,
    AutoDownloadWifiVideo:  int(m.AutoDownloadWifiVideo) ,
    AutoDownloadWifiFile:  int(m.AutoDownloadWifiFile) ,
    AutoDownloadWifiMusic:  int(m.AutoDownloadWifiMusic) ,
    AutoDownloadWifiGif:  int(m.AutoDownloadWifiGif) ,
    AutoDownloadCellularVoice:  int(m.AutoDownloadCellularVoice) ,
    AutoDownloadCellularImage:  int(m.AutoDownloadCellularImage) ,
    AutoDownloadCellularVideo:  int(m.AutoDownloadCellularVideo) ,
    AutoDownloadCellularFile:  int(m.AutoDownloadCellularFile) ,
    AutoDownloadCellularMusic:  int(m.AutoDownloadCellularMusic) ,
    AutoDownloadCellularGif:  int(m.AutoDownloadCellularGif) ,
    AutoDownloadRoamingVoice:  int(m.AutoDownloadRoamingVoice) ,
    AutoDownloadRoamingImage:  int(m.AutoDownloadRoamingImage) ,
    AutoDownloadRoamingVideo:  int(m.AutoDownloadRoamingVideo) ,
    AutoDownloadRoamingFile:  int(m.AutoDownloadRoamingFile) ,
    AutoDownloadRoamingMusic:  int(m.AutoDownloadRoamingMusic) ,
    AutoDownloadRoamingGif:  int(m.AutoDownloadRoamingGif) ,
    SaveToGallery:  int(m.SaveToGallery) ,
}
return r
}
//ToPB
func(m *PB_SettingClient_Flat)ToPB() *PB_SettingClient {
r := &PB_SettingClient{
    UserId:  int32(m.UserId) ,
    AutoDownloadWifiVoice:  int32(m.AutoDownloadWifiVoice) ,
    AutoDownloadWifiImage:  int32(m.AutoDownloadWifiImage) ,
    AutoDownloadWifiVideo:  int32(m.AutoDownloadWifiVideo) ,
    AutoDownloadWifiFile:  int32(m.AutoDownloadWifiFile) ,
    AutoDownloadWifiMusic:  int32(m.AutoDownloadWifiMusic) ,
    AutoDownloadWifiGif:  int32(m.AutoDownloadWifiGif) ,
    AutoDownloadCellularVoice:  int32(m.AutoDownloadCellularVoice) ,
    AutoDownloadCellularImage:  int32(m.AutoDownloadCellularImage) ,
    AutoDownloadCellularVideo:  int32(m.AutoDownloadCellularVideo) ,
    AutoDownloadCellularFile:  int32(m.AutoDownloadCellularFile) ,
    AutoDownloadCellularMusic:  int32(m.AutoDownloadCellularMusic) ,
    AutoDownloadCellularGif:  int32(m.AutoDownloadCellularGif) ,
    AutoDownloadRoamingVoice:  int32(m.AutoDownloadRoamingVoice) ,
    AutoDownloadRoamingImage:  int32(m.AutoDownloadRoamingImage) ,
    AutoDownloadRoamingVideo:  int32(m.AutoDownloadRoamingVideo) ,
    AutoDownloadRoamingFile:  int32(m.AutoDownloadRoamingFile) ,
    AutoDownloadRoamingMusic:  int32(m.AutoDownloadRoamingMusic) ,
    AutoDownloadRoamingGif:  int32(m.AutoDownloadRoamingGif) ,
    SaveToGallery:  int32(m.SaveToGallery) ,
}
return r
}
//folding
var PB_SettingClient__FOlD = &PB_SettingClient{
    UserId:  0 ,
    AutoDownloadWifiVoice:  0 ,
    AutoDownloadWifiImage:  0 ,
    AutoDownloadWifiVideo:  0 ,
    AutoDownloadWifiFile:  0 ,
    AutoDownloadWifiMusic:  0 ,
    AutoDownloadWifiGif:  0 ,
    AutoDownloadCellularVoice:  0 ,
    AutoDownloadCellularImage:  0 ,
    AutoDownloadCellularVideo:  0 ,
    AutoDownloadCellularFile:  0 ,
    AutoDownloadCellularMusic:  0 ,
    AutoDownloadCellularGif:  0 ,
    AutoDownloadRoamingVoice:  0 ,
    AutoDownloadRoamingImage:  0 ,
    AutoDownloadRoamingVideo:  0 ,
    AutoDownloadRoamingFile:  0 ,
    AutoDownloadRoamingMusic:  0 ,
    AutoDownloadRoamingGif:  0 ,
    SaveToGallery:  0 ,
}


type PB_SettingNotification_Flat struct {
    UserId int
    SocialLedOn int
    SocialLedColor string
    ReqestToFollowYou int
    FollowedYou int
    AccptedYourFollowRequest int
    YourPostLiked int
    YourPostCommented int
    MenthenedYouInPost int
    MenthenedYouInComment int
    YourContactsJoined int
    DirectMessage int
    DirectAlert int
    DirectPerview int
    DirectLedOn int
    DirectLedColor int
    DirectVibrate int
    DirectPopup int
    DirectSound int
    DirectPriority int
}
//ToPB
func(m *PB_SettingNotification)ToFlat() *PB_SettingNotification_Flat {
r := &PB_SettingNotification_Flat{
    UserId:  int(m.UserId) ,
    SocialLedOn:  int(m.SocialLedOn) ,
    SocialLedColor:  m.SocialLedColor ,
    ReqestToFollowYou:  int(m.ReqestToFollowYou) ,
    FollowedYou:  int(m.FollowedYou) ,
    AccptedYourFollowRequest:  int(m.AccptedYourFollowRequest) ,
    YourPostLiked:  int(m.YourPostLiked) ,
    YourPostCommented:  int(m.YourPostCommented) ,
    MenthenedYouInPost:  int(m.MenthenedYouInPost) ,
    MenthenedYouInComment:  int(m.MenthenedYouInComment) ,
    YourContactsJoined:  int(m.YourContactsJoined) ,
    DirectMessage:  int(m.DirectMessage) ,
    DirectAlert:  int(m.DirectAlert) ,
    DirectPerview:  int(m.DirectPerview) ,
    DirectLedOn:  int(m.DirectLedOn) ,
    DirectLedColor:  int(m.DirectLedColor) ,
    DirectVibrate:  int(m.DirectVibrate) ,
    DirectPopup:  int(m.DirectPopup) ,
    DirectSound:  int(m.DirectSound) ,
    DirectPriority:  int(m.DirectPriority) ,
}
return r
}
//ToPB
func(m *PB_SettingNotification_Flat)ToPB() *PB_SettingNotification {
r := &PB_SettingNotification{
    UserId:  int32(m.UserId) ,
    SocialLedOn:  int32(m.SocialLedOn) ,
    SocialLedColor:  m.SocialLedColor ,
    ReqestToFollowYou:  int32(m.ReqestToFollowYou) ,
    FollowedYou:  int32(m.FollowedYou) ,
    AccptedYourFollowRequest:  int32(m.AccptedYourFollowRequest) ,
    YourPostLiked:  int32(m.YourPostLiked) ,
    YourPostCommented:  int32(m.YourPostCommented) ,
    MenthenedYouInPost:  int32(m.MenthenedYouInPost) ,
    MenthenedYouInComment:  int32(m.MenthenedYouInComment) ,
    YourContactsJoined:  int32(m.YourContactsJoined) ,
    DirectMessage:  int32(m.DirectMessage) ,
    DirectAlert:  int32(m.DirectAlert) ,
    DirectPerview:  int32(m.DirectPerview) ,
    DirectLedOn:  int32(m.DirectLedOn) ,
    DirectLedColor:  int32(m.DirectLedColor) ,
    DirectVibrate:  int32(m.DirectVibrate) ,
    DirectPopup:  int32(m.DirectPopup) ,
    DirectSound:  int32(m.DirectSound) ,
    DirectPriority:  int32(m.DirectPriority) ,
}
return r
}
//folding
var PB_SettingNotification__FOlD = &PB_SettingNotification{
    UserId:  0 ,
    SocialLedOn:  0 ,
    SocialLedColor:  "" ,
    ReqestToFollowYou:  0 ,
    FollowedYou:  0 ,
    AccptedYourFollowRequest:  0 ,
    YourPostLiked:  0 ,
    YourPostCommented:  0 ,
    MenthenedYouInPost:  0 ,
    MenthenedYouInComment:  0 ,
    YourContactsJoined:  0 ,
    DirectMessage:  0 ,
    DirectAlert:  0 ,
    DirectPerview:  0 ,
    DirectLedOn:  0 ,
    DirectLedColor:  0 ,
    DirectVibrate:  0 ,
    DirectPopup:  0 ,
    DirectSound:  0 ,
    DirectPriority:  0 ,
}


type PB_Tag_Flat struct {
    Id int
    Name string
    Count int
    IsBlocked int
    CreatedTime int
}
//ToPB
func(m *PB_Tag)ToFlat() *PB_Tag_Flat {
r := &PB_Tag_Flat{
    Id:  int(m.Id) ,
    Name:  m.Name ,
    Count:  int(m.Count) ,
    IsBlocked:  int(m.IsBlocked) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_Tag_Flat)ToPB() *PB_Tag {
r := &PB_Tag{
    Id:  int32(m.Id) ,
    Name:  m.Name ,
    Count:  int32(m.Count) ,
    IsBlocked:  int32(m.IsBlocked) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_Tag__FOlD = &PB_Tag{
    Id:  0 ,
    Name:  "" ,
    Count:  0 ,
    IsBlocked:  0 ,
    CreatedTime:  0 ,
}


type PB_TagsPost_Flat struct {
    Id int
    TagId int
    PostId int
    TypeId int
    CreatedTime int
}
//ToPB
func(m *PB_TagsPost)ToFlat() *PB_TagsPost_Flat {
r := &PB_TagsPost_Flat{
    Id:  int(m.Id) ,
    TagId:  int(m.TagId) ,
    PostId:  int(m.PostId) ,
    TypeId:  int(m.TypeId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_TagsPost_Flat)ToPB() *PB_TagsPost {
r := &PB_TagsPost{
    Id:  int32(m.Id) ,
    TagId:  int32(m.TagId) ,
    PostId:  int32(m.PostId) ,
    TypeId:  int32(m.TypeId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_TagsPost__FOlD = &PB_TagsPost{
    Id:  0 ,
    TagId:  0 ,
    PostId:  0 ,
    TypeId:  0 ,
    CreatedTime:  0 ,
}


type PB_TestChat_Flat struct {
    Id int
    Id4 int
    TimeMs int
    Text string
    Name string
    UserId int
    C2 int
    C3 int
    C4 int
    C5 int
}
//ToPB
func(m *PB_TestChat)ToFlat() *PB_TestChat_Flat {
r := &PB_TestChat_Flat{
    Id:  int(m.Id) ,
    Id4:  int(m.Id4) ,
    TimeMs:  int(m.TimeMs) ,
    Text:  m.Text ,
    Name:  m.Name ,
    UserId:  int(m.UserId) ,
    C2:  int(m.C2) ,
    C3:  int(m.C3) ,
    C4:  int(m.C4) ,
    C5:  int(m.C5) ,
}
return r
}
//ToPB
func(m *PB_TestChat_Flat)ToPB() *PB_TestChat {
r := &PB_TestChat{
    Id:  int64(m.Id) ,
    Id4:  int64(m.Id4) ,
    TimeMs:  int64(m.TimeMs) ,
    Text:  m.Text ,
    Name:  m.Name ,
    UserId:  int64(m.UserId) ,
    C2:  int32(m.C2) ,
    C3:  int32(m.C3) ,
    C4:  int32(m.C4) ,
    C5:  int32(m.C5) ,
}
return r
}
//folding
var PB_TestChat__FOlD = &PB_TestChat{
    Id:  0 ,
    Id4:  0 ,
    TimeMs:  0 ,
    Text:  "" ,
    Name:  "" ,
    UserId:  0 ,
    C2:  0 ,
    C3:  0 ,
    C4:  0 ,
    C5:  0 ,
}


type PB_TriggerLog_Flat struct {
    Id int
    ModelName string
    ChangeType string
    TargetId int
    TargetStr string
    CreatedSe int
}
//ToPB
func(m *PB_TriggerLog)ToFlat() *PB_TriggerLog_Flat {
r := &PB_TriggerLog_Flat{
    Id:  int(m.Id) ,
    ModelName:  m.ModelName ,
    ChangeType:  m.ChangeType ,
    TargetId:  int(m.TargetId) ,
    TargetStr:  m.TargetStr ,
    CreatedSe:  int(m.CreatedSe) ,
}
return r
}
//ToPB
func(m *PB_TriggerLog_Flat)ToPB() *PB_TriggerLog {
r := &PB_TriggerLog{
    Id:  int64(m.Id) ,
    ModelName:  m.ModelName ,
    ChangeType:  m.ChangeType ,
    TargetId:  int64(m.TargetId) ,
    TargetStr:  m.TargetStr ,
    CreatedSe:  int32(m.CreatedSe) ,
}
return r
}
//folding
var PB_TriggerLog__FOlD = &PB_TriggerLog{
    Id:  0 ,
    ModelName:  "" ,
    ChangeType:  "" ,
    TargetId:  0 ,
    TargetStr:  "" ,
    CreatedSe:  0 ,
}


type PB_User_Flat struct {
    Id int
    UserName string
    UserNameLower string
    FirstName string
    LastName string
    About string
    FullName string
    AvatarUrl string
    PrivacyProfile int
    Phone string
    Email string
    IsDeleted int
    PasswordHash string
    PasswordSalt string
    FollowersCount int
    FollowingCount int
    PostsCount int
    MediaCount int
    LikesCount int
    ResharedCount int
    LastActionTime int
    LastPostTime int
    PrimaryFollowingList int
    CreatedTime int
    UpdatedTime int
    SessionUuid string
    DeviceUuid string
    LastWifiMacAddress string
    LastNetworkType string
    AppVersion int
    LastActivityTime int
    LastLoginTime int
    LastIpAddress string
}
//ToPB
func(m *PB_User)ToFlat() *PB_User_Flat {
r := &PB_User_Flat{
    Id:  int(m.Id) ,
    UserName:  m.UserName ,
    UserNameLower:  m.UserNameLower ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int(m.PrivacyProfile) ,
    Phone:  m.Phone ,
    Email:  m.Email ,
    IsDeleted:  int(m.IsDeleted) ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    MediaCount:  int(m.MediaCount) ,
    LikesCount:  int(m.LikesCount) ,
    ResharedCount:  int(m.ResharedCount) ,
    LastActionTime:  int(m.LastActionTime) ,
    LastPostTime:  int(m.LastPostTime) ,
    PrimaryFollowingList:  int(m.PrimaryFollowingList) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    SessionUuid:  m.SessionUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    AppVersion:  int(m.AppVersion) ,
    LastActivityTime:  int(m.LastActivityTime) ,
    LastLoginTime:  int(m.LastLoginTime) ,
    LastIpAddress:  m.LastIpAddress ,
}
return r
}
//ToPB
func(m *PB_User_Flat)ToPB() *PB_User {
r := &PB_User{
    Id:  int32(m.Id) ,
    UserName:  m.UserName ,
    UserNameLower:  m.UserNameLower ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int32(m.PrivacyProfile) ,
    Phone:  m.Phone ,
    Email:  m.Email ,
    IsDeleted:  int32(m.IsDeleted) ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    MediaCount:  int32(m.MediaCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ResharedCount:  int32(m.ResharedCount) ,
    LastActionTime:  int32(m.LastActionTime) ,
    LastPostTime:  int32(m.LastPostTime) ,
    PrimaryFollowingList:  int32(m.PrimaryFollowingList) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    SessionUuid:  m.SessionUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    AppVersion:  int32(m.AppVersion) ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    LastLoginTime:  int32(m.LastLoginTime) ,
    LastIpAddress:  m.LastIpAddress ,
}
return r
}
//folding
var PB_User__FOlD = &PB_User{
    Id:  0 ,
    UserName:  "" ,
    UserNameLower:  "" ,
    FirstName:  "" ,
    LastName:  "" ,
    About:  "" ,
    FullName:  "" ,
    AvatarUrl:  "" ,
    PrivacyProfile:  0 ,
    Phone:  "" ,
    Email:  "" ,
    IsDeleted:  0 ,
    PasswordHash:  "" ,
    PasswordSalt:  "" ,
    FollowersCount:  0 ,
    FollowingCount:  0 ,
    PostsCount:  0 ,
    MediaCount:  0 ,
    LikesCount:  0 ,
    ResharedCount:  0 ,
    LastActionTime:  0 ,
    LastPostTime:  0 ,
    PrimaryFollowingList:  0 ,
    CreatedTime:  0 ,
    UpdatedTime:  0 ,
    SessionUuid:  "" ,
    DeviceUuid:  "" ,
    LastWifiMacAddress:  "" ,
    LastNetworkType:  "" ,
    AppVersion:  0 ,
    LastActivityTime:  0 ,
    LastLoginTime:  0 ,
    LastIpAddress:  "" ,
}


type PB_UserMetaInfo_Flat struct {
    Id int
    UserId int
    IsNotificationDirty int
    LastUserRecGen int
}
//ToPB
func(m *PB_UserMetaInfo)ToFlat() *PB_UserMetaInfo_Flat {
r := &PB_UserMetaInfo_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    IsNotificationDirty:  int(m.IsNotificationDirty) ,
    LastUserRecGen:  int(m.LastUserRecGen) ,
}
return r
}
//ToPB
func(m *PB_UserMetaInfo_Flat)ToPB() *PB_UserMetaInfo {
r := &PB_UserMetaInfo{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    IsNotificationDirty:  int32(m.IsNotificationDirty) ,
    LastUserRecGen:  int32(m.LastUserRecGen) ,
}
return r
}
//folding
var PB_UserMetaInfo__FOlD = &PB_UserMetaInfo{
    Id:  0 ,
    UserId:  0 ,
    IsNotificationDirty:  0 ,
    LastUserRecGen:  0 ,
}


type PB_UserPassword_Flat struct {
    UserId int
    Password string
    CreatedTime int
}
//ToPB
func(m *PB_UserPassword)ToFlat() *PB_UserPassword_Flat {
r := &PB_UserPassword_Flat{
    UserId:  int(m.UserId) ,
    Password:  m.Password ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}
//ToPB
func(m *PB_UserPassword_Flat)ToPB() *PB_UserPassword {
r := &PB_UserPassword{
    UserId:  int32(m.UserId) ,
    Password:  m.Password ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}
//folding
var PB_UserPassword__FOlD = &PB_UserPassword{
    UserId:  0 ,
    Password:  "" ,
    CreatedTime:  0 ,
}


type PB_UpdateGroupParticipants_Flat struct {
}
//ToPB
func(m *PB_UpdateGroupParticipants)ToFlat() *PB_UpdateGroupParticipants_Flat {
r := &PB_UpdateGroupParticipants_Flat{
}
return r
}
//ToPB
func(m *PB_UpdateGroupParticipants_Flat)ToPB() *PB_UpdateGroupParticipants {
r := &PB_UpdateGroupParticipants{
}
return r
}
//folding
var PB_UpdateGroupParticipants__FOlD = &PB_UpdateGroupParticipants{
}


type PB_UpdateNotifySettings_Flat struct {
}
//ToPB
func(m *PB_UpdateNotifySettings)ToFlat() *PB_UpdateNotifySettings_Flat {
r := &PB_UpdateNotifySettings_Flat{
}
return r
}
//ToPB
func(m *PB_UpdateNotifySettings_Flat)ToPB() *PB_UpdateNotifySettings {
r := &PB_UpdateNotifySettings{
}
return r
}
//folding
var PB_UpdateNotifySettings__FOlD = &PB_UpdateNotifySettings{
}


type PB_UpdateServiceNotification_Flat struct {
}
//ToPB
func(m *PB_UpdateServiceNotification)ToFlat() *PB_UpdateServiceNotification_Flat {
r := &PB_UpdateServiceNotification_Flat{
}
return r
}
//ToPB
func(m *PB_UpdateServiceNotification_Flat)ToPB() *PB_UpdateServiceNotification {
r := &PB_UpdateServiceNotification{
}
return r
}
//folding
var PB_UpdateServiceNotification__FOlD = &PB_UpdateServiceNotification{
}


type PB_UpdateMessageMeta_Flat struct {
    MessageId int
    AtTime int
}
//ToPB
func(m *PB_UpdateMessageMeta)ToFlat() *PB_UpdateMessageMeta_Flat {
r := &PB_UpdateMessageMeta_Flat{
    MessageId:  int(m.MessageId) ,
    AtTime:  int(m.AtTime) ,
}
return r
}
//ToPB
func(m *PB_UpdateMessageMeta_Flat)ToPB() *PB_UpdateMessageMeta {
r := &PB_UpdateMessageMeta{
    MessageId:  int64(m.MessageId) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}
//folding
var PB_UpdateMessageMeta__FOlD = &PB_UpdateMessageMeta{
    MessageId:  0 ,
    AtTime:  0 ,
}


type PB_UpdateMessageId_Flat struct {
    OldMessageId int
    NewMessageId int
}
//ToPB
func(m *PB_UpdateMessageId)ToFlat() *PB_UpdateMessageId_Flat {
r := &PB_UpdateMessageId_Flat{
    OldMessageId:  int(m.OldMessageId) ,
    NewMessageId:  int(m.NewMessageId) ,
}
return r
}
//ToPB
func(m *PB_UpdateMessageId_Flat)ToPB() *PB_UpdateMessageId {
r := &PB_UpdateMessageId{
    OldMessageId:  int64(m.OldMessageId) ,
    NewMessageId:  int64(m.NewMessageId) ,
}
return r
}
//folding
var PB_UpdateMessageId__FOlD = &PB_UpdateMessageId{
    OldMessageId:  0 ,
    NewMessageId:  0 ,
}


type PB_UpdateMessageToEdit_Flat struct {
    MessageId int
    NewText string
}
//ToPB
func(m *PB_UpdateMessageToEdit)ToFlat() *PB_UpdateMessageToEdit_Flat {
r := &PB_UpdateMessageToEdit_Flat{
    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}
//ToPB
func(m *PB_UpdateMessageToEdit_Flat)ToPB() *PB_UpdateMessageToEdit {
r := &PB_UpdateMessageToEdit{
    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}
//folding
var PB_UpdateMessageToEdit__FOlD = &PB_UpdateMessageToEdit{
    MessageId:  0 ,
    NewText:  "" ,
}


type PB_UpdateMessageToDelete_Flat struct {
    MessageId int
}
//ToPB
func(m *PB_UpdateMessageToDelete)ToFlat() *PB_UpdateMessageToDelete_Flat {
r := &PB_UpdateMessageToDelete_Flat{
    MessageId:  int(m.MessageId) ,
}
return r
}
//ToPB
func(m *PB_UpdateMessageToDelete_Flat)ToPB() *PB_UpdateMessageToDelete {
r := &PB_UpdateMessageToDelete{
    MessageId:  int64(m.MessageId) ,
}
return r
}
//folding
var PB_UpdateMessageToDelete__FOlD = &PB_UpdateMessageToDelete{
    MessageId:  0 ,
}


type PB_UpdateRoomActionDoing_Flat struct {
    RoomKey string
    ActionType RoomActionDoingEnum
}
//ToPB
func(m *PB_UpdateRoomActionDoing)ToFlat() *PB_UpdateRoomActionDoing_Flat {
r := &PB_UpdateRoomActionDoing_Flat{
    RoomKey:  m.RoomKey ,
    
}
return r
}
//ToPB
func(m *PB_UpdateRoomActionDoing_Flat)ToPB() *PB_UpdateRoomActionDoing {
r := &PB_UpdateRoomActionDoing{
    RoomKey:  m.RoomKey ,
    
}
return r
}
//folding
var PB_UpdateRoomActionDoing__FOlD = &PB_UpdateRoomActionDoing{
    RoomKey:  "" ,
    
}


type PB_UpdateUserBlocked_Flat struct {
    UserId int
    Blocked bool
}
//ToPB
func(m *PB_UpdateUserBlocked)ToFlat() *PB_UpdateUserBlocked_Flat {
r := &PB_UpdateUserBlocked_Flat{
    UserId:  int(m.UserId) ,
    Blocked:  m.Blocked ,
}
return r
}
//ToPB
func(m *PB_UpdateUserBlocked_Flat)ToPB() *PB_UpdateUserBlocked {
r := &PB_UpdateUserBlocked{
    UserId:  int32(m.UserId) ,
    Blocked:  m.Blocked ,
}
return r
}
//folding
var PB_UpdateUserBlocked__FOlD = &PB_UpdateUserBlocked{
    UserId:  0 ,
    Blocked:  false ,
}


type PB_ChatView_Flat struct {
    ChatKey string
    RoomKey string
    RoomTypeEnumId int
    UserId int
    PeerUserId int
    GroupId int
    CreatedSe int
    UpdatedMs int
    LastMessageId int
    LastDeletedMessageId int
    LastSeenMessageId int
    LastSeqSeen int
    LastSeqDelete int
    CurrentSeq int
    UserView PB_UserView
    SharedMediaCount int
    UnseenCount int
    FirstUnreadMessage PB_MessageView
    LastMessage PB_MessageView
}
//ToPB
func(m *PB_ChatView)ToFlat() *PB_ChatView_Flat {
r := &PB_ChatView_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedSe:  int(m.CreatedSe) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    LastMessageId:  int(m.LastMessageId) ,
    LastDeletedMessageId:  int(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int(m.LastSeenMessageId) ,
    LastSeqSeen:  int(m.LastSeqSeen) ,
    LastSeqDelete:  int(m.LastSeqDelete) ,
    CurrentSeq:  int(m.CurrentSeq) ,
    
    SharedMediaCount:  int(m.SharedMediaCount) ,
    UnseenCount:  int(m.UnseenCount) ,
    
    
}
return r
}
//ToPB
func(m *PB_ChatView_Flat)ToPB() *PB_ChatView {
r := &PB_ChatView{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedSe:  int32(m.CreatedSe) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    LastMessageId:  int64(m.LastMessageId) ,
    LastDeletedMessageId:  int64(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int64(m.LastSeenMessageId) ,
    LastSeqSeen:  int32(m.LastSeqSeen) ,
    LastSeqDelete:  int32(m.LastSeqDelete) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
    
    SharedMediaCount:  int32(m.SharedMediaCount) ,
    UnseenCount:  int32(m.UnseenCount) ,
    
    
}
return r
}
//folding
var PB_ChatView__FOlD = &PB_ChatView{
    ChatKey:  "" ,
    RoomKey:  "" ,
    RoomTypeEnumId:  0 ,
    UserId:  0 ,
    PeerUserId:  0 ,
    GroupId:  0 ,
    CreatedSe:  0 ,
    UpdatedMs:  0 ,
    LastMessageId:  0 ,
    LastDeletedMessageId:  0 ,
    LastSeenMessageId:  0 ,
    LastSeqSeen:  0 ,
    LastSeqDelete:  0 ,
    CurrentSeq:  0 ,
    
    SharedMediaCount:  0 ,
    UnseenCount:  0 ,
    
    
}


type PB_MessageView_Flat struct {
    MessageId int
    MessageKey string
    RoomKey string
    UserId int
    MessageFileId int
    MessageTypeEnumId int
    Text string
    CreatedSe int
    PeerReceivedTime int
    PeerSeenTime int
    DeliviryStatusEnumId int
    ChatKey string
    RoomTypeEnumId int
    IsByMe bool
    RemoteId int
    MessageFileView PB_MessageFileView
}
//ToPB
func(m *PB_MessageView)ToFlat() *PB_MessageView_Flat {
r := &PB_MessageView_Flat{
    MessageId:  int(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnumId:  int(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int(m.CreatedSe) ,
    PeerReceivedTime:  int(m.PeerReceivedTime) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int(m.DeliviryStatusEnumId) ,
    ChatKey:  m.ChatKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    IsByMe:  m.IsByMe ,
    RemoteId:  int(m.RemoteId) ,
    
}
return r
}
//ToPB
func(m *PB_MessageView_Flat)ToPB() *PB_MessageView {
r := &PB_MessageView{
    MessageId:  int64(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnumId:  int32(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int32(m.CreatedSe) ,
    PeerReceivedTime:  int32(m.PeerReceivedTime) ,
    PeerSeenTime:  int32(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int32(m.DeliviryStatusEnumId) ,
    ChatKey:  m.ChatKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    IsByMe:  m.IsByMe ,
    RemoteId:  int64(m.RemoteId) ,
    
}
return r
}
//folding
var PB_MessageView__FOlD = &PB_MessageView{
    MessageId:  0 ,
    MessageKey:  "" ,
    RoomKey:  "" ,
    UserId:  0 ,
    MessageFileId:  0 ,
    MessageTypeEnumId:  0 ,
    Text:  "" ,
    CreatedSe:  0 ,
    PeerReceivedTime:  0 ,
    PeerSeenTime:  0 ,
    DeliviryStatusEnumId:  0 ,
    ChatKey:  "" ,
    RoomTypeEnumId:  0 ,
    IsByMe:  false ,
    RemoteId:  0 ,
    
}


type PB_MessageFileView_Flat struct {
    MessageFileId int
    MessageFileKey string
    OriginalUserId int
    Name string
    Size int
    FileTypeEnumId int
    Width int
    Height int
    Duration int
    Extension string
    HashMd5 string
    HashAccess int
    CreatedSe int
    ServerSrc string
    ServerPath string
    ServerThumbPath string
    BucketId string
    ServerId int
    CanDel int
    ServerThumbLocalSrc string
    RemoteMessageFileId int
    LocalSrc string
    ThumbLocalSrc string
    MessageFileStatusId string
}
//ToPB
func(m *PB_MessageFileView)ToFlat() *PB_MessageFileView_Flat {
r := &PB_MessageFileView_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileTypeEnumId:  int(m.FileTypeEnumId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int(m.HashAccess) ,
    CreatedSe:  int(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
    ServerThumbLocalSrc:  m.ServerThumbLocalSrc ,
    RemoteMessageFileId:  int(m.RemoteMessageFileId) ,
    LocalSrc:  m.LocalSrc ,
    ThumbLocalSrc:  m.ThumbLocalSrc ,
    MessageFileStatusId:  m.MessageFileStatusId ,
}
return r
}
//ToPB
func(m *PB_MessageFileView_Flat)ToPB() *PB_MessageFileView {
r := &PB_MessageFileView{
    MessageFileId:  int64(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int32(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileTypeEnumId:  int32(m.FileTypeEnumId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int64(m.HashAccess) ,
    CreatedSe:  int32(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
    ServerThumbLocalSrc:  m.ServerThumbLocalSrc ,
    RemoteMessageFileId:  int64(m.RemoteMessageFileId) ,
    LocalSrc:  m.LocalSrc ,
    ThumbLocalSrc:  m.ThumbLocalSrc ,
    MessageFileStatusId:  m.MessageFileStatusId ,
}
return r
}
//folding
var PB_MessageFileView__FOlD = &PB_MessageFileView{
    MessageFileId:  0 ,
    MessageFileKey:  "" ,
    OriginalUserId:  0 ,
    Name:  "" ,
    Size:  0 ,
    FileTypeEnumId:  0 ,
    Width:  0 ,
    Height:  0 ,
    Duration:  0 ,
    Extension:  "" ,
    HashMd5:  "" ,
    HashAccess:  0 ,
    CreatedSe:  0 ,
    ServerSrc:  "" ,
    ServerPath:  "" ,
    ServerThumbPath:  "" ,
    BucketId:  "" ,
    ServerId:  0 ,
    CanDel:  0 ,
    ServerThumbLocalSrc:  "" ,
    RemoteMessageFileId:  0 ,
    LocalSrc:  "" ,
    ThumbLocalSrc:  "" ,
    MessageFileStatusId:  "" ,
}


type PB_UserView_Flat struct {
    UserId int
    UserName string
    FirstName string
    LastName string
    About string
    FullName string
    AvatarUrl string
    PrivacyProfile int
    IsDeleted int
    FollowersCount int
    FollowingCount int
    PostsCount int
    UpdatedTime int
    AppVersion int
    LastActivityTime int
    FollowingType int
}
//ToPB
func(m *PB_UserView)ToFlat() *PB_UserView_Flat {
r := &PB_UserView_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int(m.PrivacyProfile) ,
    IsDeleted:  int(m.IsDeleted) ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    AppVersion:  int(m.AppVersion) ,
    LastActivityTime:  int(m.LastActivityTime) ,
    FollowingType:  int(m.FollowingType) ,
}
return r
}
//ToPB
func(m *PB_UserView_Flat)ToPB() *PB_UserView {
r := &PB_UserView{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int32(m.PrivacyProfile) ,
    IsDeleted:  int32(m.IsDeleted) ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    AppVersion:  int32(m.AppVersion) ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    FollowingType:  int32(m.FollowingType) ,
}
return r
}
//folding
var PB_UserView__FOlD = &PB_UserView{
    UserId:  0 ,
    UserName:  "" ,
    FirstName:  "" ,
    LastName:  "" ,
    About:  "" ,
    FullName:  "" ,
    AvatarUrl:  "" ,
    PrivacyProfile:  0 ,
    IsDeleted:  0 ,
    FollowersCount:  0 ,
    FollowingCount:  0 ,
    PostsCount:  0 ,
    UpdatedTime:  0 ,
    AppVersion:  0 ,
    LastActivityTime:  0 ,
    FollowingType:  0 ,
}



/*
///// to_flat ///

func(m *GeoLocation)ToFlat() *GeoLocation_Flat {
r := &GeoLocation_Flat{
    Lat:  float64(m.Lat) ,
    Lon:  float64(m.Lon) ,
}
return r
}

func(m *RoomMessageLog)ToFlat() *RoomMessageLog_Flat {
r := &RoomMessageLog_Flat{
    
    TargetUserId:  int(m.TargetUserId) ,
    ByUserId:  int(m.ByUserId) ,
}
return r
}

func(m *RoomMessageForwardFrom)ToFlat() *RoomMessageForwardFrom_Flat {
r := &RoomMessageForwardFrom_Flat{
    RoomId:  int(m.RoomId) ,
    MessageId:  int(m.MessageId) ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
}
return r
}

func(m *RoomDraft)ToFlat() *RoomDraft_Flat {
r := &RoomDraft_Flat{
    Message:  m.Message ,
    ReplyTo:  int(m.ReplyTo) ,
}
return r
}

func(m *ChatRoom)ToFlat() *ChatRoom_Flat {
r := &ChatRoom_Flat{
}
return r
}

func(m *Pagination)ToFlat() *Pagination_Flat {
r := &Pagination_Flat{
    Offset:  int(m.Offset) ,
    Limit:  int(m.Limit) ,
}
return r
}

func(m *PB_CommandToServer)ToFlat() *PB_CommandToServer_Flat {
r := &PB_CommandToServer_Flat{
    ClientCallId:  int(m.ClientCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_CommandToClient)ToFlat() *PB_CommandToClient_Flat {
r := &PB_CommandToClient_Flat{
    ServerCallId:  int(m.ServerCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_CommandReachedToServer)ToFlat() *PB_CommandReachedToServer_Flat {
r := &PB_CommandReachedToServer_Flat{
    ClientCallId:  int(m.ClientCallId) ,
}
return r
}

func(m *PB_CommandReachedToClient)ToFlat() *PB_CommandReachedToClient_Flat {
r := &PB_CommandReachedToClient_Flat{
    ServerCallId:  int(m.ServerCallId) ,
}
return r
}

func(m *PB_ResponseToClient)ToFlat() *PB_ResponseToClient_Flat {
r := &PB_ResponseToClient_Flat{
    ClientCallId:  int(m.ClientCallId) ,
    PBClass:  m.PBClass ,
    RpcFullName:  m.RpcFullName ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_Offline_NewDirectMessage)ToFlat() *PB_Offline_NewDirectMessage_Flat {
r := &PB_Offline_NewDirectMessage_Flat{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int(m.FromMessageId) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_Offline_MessagesReachedServer)ToFlat() *PB_Offline_MessagesReachedServer_Flat {
r := &PB_Offline_MessagesReachedServer_Flat{
    MessageKeys:  m.MessageKeys ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_Offline_MessagesDeliveredToUser)ToFlat() *PB_Offline_MessagesDeliveredToUser_Flat {
r := &PB_Offline_MessagesDeliveredToUser_Flat{
    MessageKeys:  m.MessageKeys ,
    RoomKey:  m.RoomKey ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_Offline_MessagesSeenByPeer)ToFlat() *PB_Offline_MessagesSeenByPeer_Flat {
r := &PB_Offline_MessagesSeenByPeer_Flat{
    MessageKeys:  m.MessageKeys ,
    RoomKey:  m.RoomKey ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_Offline_MessagesDeletedFromServer)ToFlat() *PB_Offline_MessagesDeletedFromServer_Flat {
r := &PB_Offline_MessagesDeletedFromServer_Flat{
    MessageKeys:  m.MessageKeys ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_Offline_ChangeMessageId)ToFlat() *PB_Offline_ChangeMessageId_Flat {
r := &PB_Offline_ChangeMessageId_Flat{
    MessageKey:  int(m.MessageKey) ,
    NewMessageId:  int(m.NewMessageId) ,
}
return r
}

func(m *PB_Offline_ChangeMessageFileId)ToFlat() *PB_Offline_ChangeMessageFileId_Flat {
r := &PB_Offline_ChangeMessageFileId_Flat{
    MessageFileKey:  int(m.MessageFileKey) ,
    NewMessageFileId:  int(m.NewMessageFileId) ,
}
return r
}

func(m *PB_Offline_MessageToEdit)ToFlat() *PB_Offline_MessageToEdit_Flat {
r := &PB_Offline_MessageToEdit_Flat{
    MessageKey:  int(m.MessageKey) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_Offline_MessageToDelete)ToFlat() *PB_Offline_MessageToDelete_Flat {
r := &PB_Offline_MessageToDelete_Flat{
    MessageKey:  int(m.MessageKey) ,
}
return r
}

func(m *PB_Online_RoomActionDoing)ToFlat() *PB_Online_RoomActionDoing_Flat {
r := &PB_Online_RoomActionDoing_Flat{
    RoomKey:  m.RoomKey ,
    
}
return r
}

func(m *PB_Offline_Messagings)ToFlat() *PB_Offline_Messagings_Flat {
r := &PB_Offline_Messagings_Flat{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_UserParam_CheckUserName2)ToFlat() *PB_UserParam_CheckUserName2_Flat {
r := &PB_UserParam_CheckUserName2_Flat{
}
return r
}

func(m *PB_UserResponse_CheckUserName2)ToFlat() *PB_UserResponse_CheckUserName2_Flat {
r := &PB_UserResponse_CheckUserName2_Flat{
}
return r
}

func(m *PB_ChatParam_AddNewMessage)ToFlat() *PB_ChatParam_AddNewMessage_Flat {
r := &PB_ChatParam_AddNewMessage_Flat{
    ReplyToMessageId:  int(m.ReplyToMessageId) ,
    Blob:  []byte(m.Blob) ,
    
}
return r
}

func(m *PB_ChatResponse_AddNewMessage)ToFlat() *PB_ChatResponse_AddNewMessage_Flat {
r := &PB_ChatResponse_AddNewMessage_Flat{
}
return r
}

func(m *PB_ChatParam_SetRoomActionDoing)ToFlat() *PB_ChatParam_SetRoomActionDoing_Flat {
r := &PB_ChatParam_SetRoomActionDoing_Flat{
    GroupId:  int(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,
    
}
return r
}

func(m *PB_ChatResponse_SetRoomActionDoing)ToFlat() *PB_ChatResponse_SetRoomActionDoing_Flat {
r := &PB_ChatResponse_SetRoomActionDoing_Flat{
}
return r
}

func(m *PB_ChatParam_SetChatMessagesRangeAsSeen)ToFlat() *PB_ChatParam_SetChatMessagesRangeAsSeen_Flat {
r := &PB_ChatParam_SetChatMessagesRangeAsSeen_Flat{
    ChatKey:  m.ChatKey ,
    BottomMessageId:  int(m.BottomMessageId) ,
    TopMessageId:  int(m.TopMessageId) ,
    SeenTimeMs:  int(m.SeenTimeMs) ,
}
return r
}

func(m *PB_ChatResponse_SetChatMessagesRangeAsSeen)ToFlat() *PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat {
r := &PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat{
}
return r
}

func(m *PB_ChatParam_DeleteChatHistory)ToFlat() *PB_ChatParam_DeleteChatHistory_Flat {
r := &PB_ChatParam_DeleteChatHistory_Flat{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int(m.FromMessageId) ,
}
return r
}

func(m *PB_ChatResponse_DeleteChatHistory)ToFlat() *PB_ChatResponse_DeleteChatHistory_Flat {
r := &PB_ChatResponse_DeleteChatHistory_Flat{
}
return r
}

func(m *PB_ChatParam_DeleteMessagesByIds)ToFlat() *PB_ChatParam_DeleteMessagesByIds_Flat {
r := &PB_ChatParam_DeleteMessagesByIds_Flat{
    ChatKey:  m.ChatKey ,
    Both:  m.Both ,
    MessagesIds:  helper.SliceInt64ToInt(m.MessagesIds) ,
}
return r
}

func(m *PB_ChatResponse_DeleteMessagesByIds)ToFlat() *PB_ChatResponse_DeleteMessagesByIds_Flat {
r := &PB_ChatResponse_DeleteMessagesByIds_Flat{
}
return r
}

func(m *PB_ChatParam_SetMessagesAsReceived)ToFlat() *PB_ChatParam_SetMessagesAsReceived_Flat {
r := &PB_ChatParam_SetMessagesAsReceived_Flat{
    ChatRoom:  m.ChatRoom ,
    MessageIds:  helper.SliceInt64ToInt(m.MessageIds) ,
}
return r
}

func(m *PB_ChatResponse_SetMessagesAsReceived)ToFlat() *PB_ChatResponse_SetMessagesAsReceived_Flat {
r := &PB_ChatResponse_SetMessagesAsReceived_Flat{
}
return r
}

func(m *PB_ChatParam_EditMessage)ToFlat() *PB_ChatParam_EditMessage_Flat {
r := &PB_ChatParam_EditMessage_Flat{
    RoomKey:  m.RoomKey ,
    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_ChatResponse_EditMessage)ToFlat() *PB_ChatResponse_EditMessage_Flat {
r := &PB_ChatResponse_EditMessage_Flat{
}
return r
}

func(m *PB_ChatParam_GetChatList)ToFlat() *PB_ChatParam_GetChatList_Flat {
r := &PB_ChatParam_GetChatList_Flat{
}
return r
}

func(m *PB_ChatResponse_GetChatList)ToFlat() *PB_ChatResponse_GetChatList_Flat {
r := &PB_ChatResponse_GetChatList_Flat{
    
    
}
return r
}

func(m *PB_ChatParam_GetChatHistoryToOlder)ToFlat() *PB_ChatParam_GetChatHistoryToOlder_Flat {
r := &PB_ChatParam_GetChatHistoryToOlder_Flat{
    ChatKey:  m.ChatKey ,
    Limit:  int(m.Limit) ,
    FromTopMessageId:  int(m.FromTopMessageId) ,
}
return r
}

func(m *PB_ChatResponse_GetChatHistoryToOlder)ToFlat() *PB_ChatResponse_GetChatHistoryToOlder_Flat {
r := &PB_ChatResponse_GetChatHistoryToOlder_Flat{
    
    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_ChatParam_GetFreshAllDirectMessagesList)ToFlat() *PB_ChatParam_GetFreshAllDirectMessagesList_Flat {
r := &PB_ChatParam_GetFreshAllDirectMessagesList_Flat{
    LowMessageId:  int(m.LowMessageId) ,
    HighMessageId:  int(m.HighMessageId) ,
}
return r
}

func(m *PB_ChatResponse_GetFreshAllDirectMessagesList)ToFlat() *PB_ChatResponse_GetFreshAllDirectMessagesList_Flat {
r := &PB_ChatResponse_GetFreshAllDirectMessagesList_Flat{
    
    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_MsgParam_AddNewTextMessage)ToFlat() *PB_MsgParam_AddNewTextMessage_Flat {
r := &PB_MsgParam_AddNewTextMessage_Flat{
    Text:  m.Text ,
    MessageKey:  m.MessageKey ,
    ToRoomKey:  m.ToRoomKey ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,
    ReplyToMessageId:  int(m.ReplyToMessageId) ,
    
}
return r
}

func(m *PB_MsgResponse_AddNewTextMessage)ToFlat() *PB_MsgResponse_AddNewTextMessage_Flat {
r := &PB_MsgResponse_AddNewTextMessage_Flat{
}
return r
}

func(m *PB_MsgParam_AddNewMessage)ToFlat() *PB_MsgParam_AddNewMessage_Flat {
r := &PB_MsgParam_AddNewMessage_Flat{
    Text:  m.Text ,
    MessageKey:  m.MessageKey ,
    ToRoomKey:  m.ToRoomKey ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,
    ReplyToMessageId:  int(m.ReplyToMessageId) ,
    
    
    Blob:  []byte(m.Blob) ,
    
}
return r
}

func(m *PB_MsgResponse_AddNewMessage)ToFlat() *PB_MsgResponse_AddNewMessage_Flat {
r := &PB_MsgResponse_AddNewMessage_Flat{
}
return r
}

func(m *PB_MsgParam_SetRoomActionDoing)ToFlat() *PB_MsgParam_SetRoomActionDoing_Flat {
r := &PB_MsgParam_SetRoomActionDoing_Flat{
    GroupId:  int(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,
    
}
return r
}

func(m *PB_MsgResponse_SetRoomActionDoing)ToFlat() *PB_MsgResponse_SetRoomActionDoing_Flat {
r := &PB_MsgResponse_SetRoomActionDoing_Flat{
}
return r
}

func(m *PB_MsgParam_GetMessagesByIds)ToFlat() *PB_MsgParam_GetMessagesByIds_Flat {
r := &PB_MsgParam_GetMessagesByIds_Flat{
    
}
return r
}

func(m *PB_MsgResponse_GetMessagesByIds)ToFlat() *PB_MsgResponse_GetMessagesByIds_Flat {
r := &PB_MsgResponse_GetMessagesByIds_Flat{
    
}
return r
}

func(m *PB_MsgParam_GetMessagesHistory)ToFlat() *PB_MsgParam_GetMessagesHistory_Flat {
r := &PB_MsgParam_GetMessagesHistory_Flat{
    ChatKey:  m.ChatKey ,
    FromSeq:  int(m.FromSeq) ,
    EndSeq:  int(m.EndSeq) ,
}
return r
}

func(m *PB_MsgResponse_GetMessagesHistory)ToFlat() *PB_MsgResponse_GetMessagesHistory_Flat {
r := &PB_MsgResponse_GetMessagesHistory_Flat{
    
}
return r
}

func(m *PB_MsgParam_SetChatMessagesRangeAsSeen)ToFlat() *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat {
r := &PB_MsgParam_SetChatMessagesRangeAsSeen_Flat{
    ChatKey:  m.ChatKey ,
    FromSeq:  int(m.FromSeq) ,
    EndSeq:  int(m.EndSeq) ,
    SeenTimeMs:  int(m.SeenTimeMs) ,
}
return r
}

func(m *PB_MsgResponse_SetChatMessagesRangeAsSeen)ToFlat() *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat {
r := &PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat{
}
return r
}

func(m *PB_MsgParam_DeleteChatHistory)ToFlat() *PB_MsgParam_DeleteChatHistory_Flat {
r := &PB_MsgParam_DeleteChatHistory_Flat{
    ChatKey:  m.ChatKey ,
    ToSeq:  int(m.ToSeq) ,
}
return r
}

func(m *PB_MsgResponse_DeleteChatHistory)ToFlat() *PB_MsgResponse_DeleteChatHistory_Flat {
r := &PB_MsgResponse_DeleteChatHistory_Flat{
}
return r
}

func(m *PB_MsgParam_DeleteMessagesByIds)ToFlat() *PB_MsgParam_DeleteMessagesByIds_Flat {
r := &PB_MsgParam_DeleteMessagesByIds_Flat{
    
}
return r
}

func(m *PB_MsgResponse_DeleteMessagesByIds)ToFlat() *PB_MsgResponse_DeleteMessagesByIds_Flat {
r := &PB_MsgResponse_DeleteMessagesByIds_Flat{
}
return r
}

func(m *PB_MsgParam_SetMessagesAsReceived)ToFlat() *PB_MsgParam_SetMessagesAsReceived_Flat {
r := &PB_MsgParam_SetMessagesAsReceived_Flat{
    
}
return r
}

func(m *PB_MsgResponse_SetMessagesAsReceived)ToFlat() *PB_MsgResponse_SetMessagesAsReceived_Flat {
r := &PB_MsgResponse_SetMessagesAsReceived_Flat{
}
return r
}

func(m *PB_MsgParam_ForwardMessages)ToFlat() *PB_MsgParam_ForwardMessages_Flat {
r := &PB_MsgParam_ForwardMessages_Flat{
    
    ToDirectChatKeys:  m.ToDirectChatKeys ,
    ToGroupChatIds:  helper.SliceInt64ToInt(m.ToGroupChatIds) ,
}
return r
}

func(m *PB_MsgResponse_ForwardMessages)ToFlat() *PB_MsgResponse_ForwardMessages_Flat {
r := &PB_MsgResponse_ForwardMessages_Flat{
}
return r
}

func(m *PB_MsgParam_EditMessage)ToFlat() *PB_MsgParam_EditMessage_Flat {
r := &PB_MsgParam_EditMessage_Flat{
    ChatKey:  m.ChatKey ,
    
    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_MsgResponse_EditMessage)ToFlat() *PB_MsgResponse_EditMessage_Flat {
r := &PB_MsgResponse_EditMessage_Flat{
}
return r
}

func(m *PB_MsgParam_BroadcastNewMessage)ToFlat() *PB_MsgParam_BroadcastNewMessage_Flat {
r := &PB_MsgParam_BroadcastNewMessage_Flat{
    Text:  m.Text ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,
    
}
return r
}

func(m *PB_MsgResponse_BroadcastNewMessage)ToFlat() *PB_MsgResponse_BroadcastNewMessage_Flat {
r := &PB_MsgResponse_BroadcastNewMessage_Flat{
}
return r
}

func(m *PB_MsgParam_GetFreshChatList)ToFlat() *PB_MsgParam_GetFreshChatList_Flat {
r := &PB_MsgParam_GetFreshChatList_Flat{
}
return r
}

func(m *PB_MsgResponse_GetFreshChatList)ToFlat() *PB_MsgResponse_GetFreshChatList_Flat {
r := &PB_MsgResponse_GetFreshChatList_Flat{
    
    
}
return r
}

func(m *PB_MsgParam_GetFreshRoomMessagesList)ToFlat() *PB_MsgParam_GetFreshRoomMessagesList_Flat {
r := &PB_MsgParam_GetFreshRoomMessagesList_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    Last:  int(m.Last) ,
}
return r
}

func(m *PB_MsgResponse_GetFreshRoomMessagesList)ToFlat() *PB_MsgResponse_GetFreshRoomMessagesList_Flat {
r := &PB_MsgResponse_GetFreshRoomMessagesList_Flat{
    
    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_MsgParam_GetFreshAllDirectMessagesList)ToFlat() *PB_MsgParam_GetFreshAllDirectMessagesList_Flat {
r := &PB_MsgParam_GetFreshAllDirectMessagesList_Flat{
    LowMessageId:  int(m.LowMessageId) ,
    HighMessageId:  int(m.HighMessageId) ,
}
return r
}

func(m *PB_MsgResponse_GetFreshAllDirectMessagesList)ToFlat() *PB_MsgResponse_GetFreshAllDirectMessagesList_Flat {
r := &PB_MsgResponse_GetFreshAllDirectMessagesList_Flat{
    
    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_MessageForwardedFrom)ToFlat() *PB_MessageForwardedFrom_Flat {
r := &PB_MessageForwardedFrom_Flat{
    RoomId:  int(m.RoomId) ,
    
    MessageId:  int(m.MessageId) ,
    MessageSeq:  int(m.MessageSeq) ,
}
return r
}

func(m *PB_MessagesCollections)ToFlat() *PB_MessagesCollections_Flat {
r := &PB_MessagesCollections_Flat{
    DirectMessagesIds:  helper.SliceInt64ToInt(m.DirectMessagesIds) ,
    GroupMessagesIds:  helper.SliceInt64ToInt(m.GroupMessagesIds) ,
    BroadCatMessagesIds:  helper.SliceInt64ToInt(m.BroadCatMessagesIds) ,
}
return r
}

func(m *PB_MsgParam_Echo)ToFlat() *PB_MsgParam_Echo_Flat {
r := &PB_MsgParam_Echo_Flat{
    Text:  m.Text ,
}
return r
}

func(m *PB_MsgResponse_PB_MsgParam_Echo)ToFlat() *PB_MsgResponse_PB_MsgParam_Echo_Flat {
r := &PB_MsgResponse_PB_MsgParam_Echo_Flat{
    Text:  m.Text ,
}
return r
}

func(m *PB_SyncParam_GetDirectUpdates)ToFlat() *PB_SyncParam_GetDirectUpdates_Flat {
r := &PB_SyncParam_GetDirectUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetDirectUpdates)ToFlat() *PB_SyncResponse_GetDirectUpdates_Flat {
r := &PB_SyncResponse_GetDirectUpdates_Flat{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_GetGeneralUpdates)ToFlat() *PB_SyncParam_GetGeneralUpdates_Flat {
r := &PB_SyncParam_GetGeneralUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetGeneralUpdates)ToFlat() *PB_SyncResponse_GetGeneralUpdates_Flat {
r := &PB_SyncResponse_GetGeneralUpdates_Flat{
    
    
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_GetNotifyUpdates)ToFlat() *PB_SyncParam_GetNotifyUpdates_Flat {
r := &PB_SyncParam_GetNotifyUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetNotifyUpdates)ToFlat() *PB_SyncResponse_GetNotifyUpdates_Flat {
r := &PB_SyncResponse_GetNotifyUpdates_Flat{
    
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_SetLastSyncDirectUpdateId)ToFlat() *PB_SyncParam_SetLastSyncDirectUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncDirectUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncDirectUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncDirectUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncDirectUpdateId_Flat{
}
return r
}

func(m *PB_SyncParam_SetLastSyncGeneralUpdateId)ToFlat() *PB_SyncParam_SetLastSyncGeneralUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncGeneralUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncGeneralUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat{
}
return r
}

func(m *PB_SyncParam_SetLastSyncNotifyUpdateId)ToFlat() *PB_SyncParam_SetLastSyncNotifyUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncNotifyUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncNotifyUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat{
}
return r
}

func(m *PB_NotifyUpdatesView)ToFlat() *PB_NotifyUpdatesView_Flat {
r := &PB_NotifyUpdatesView_Flat{
    
    
}
return r
}

func(m *PB_AllLivePushes)ToFlat() *PB_AllLivePushes_Flat {
r := &PB_AllLivePushes_Flat{
    
    
    
}
return r
}

func(m *PB_UserParam_BlockUser)ToFlat() *PB_UserParam_BlockUser_Flat {
r := &PB_UserParam_BlockUser_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserOfflineResponse_BlockUser)ToFlat() *PB_UserOfflineResponse_BlockUser_Flat {
r := &PB_UserOfflineResponse_BlockUser_Flat{
    ByUserId:  int(m.ByUserId) ,
    TargetUserId:  int(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UnBlockUser)ToFlat() *PB_UserParam_UnBlockUser_Flat {
r := &PB_UserParam_UnBlockUser_Flat{
    Offset:  int(m.Offset) ,
    Limit:  int(m.Limit) ,
}
return r
}

func(m *PB_UserOfflineResponse_UnBlockUser)ToFlat() *PB_UserOfflineResponse_UnBlockUser_Flat {
r := &PB_UserOfflineResponse_UnBlockUser_Flat{
    
}
return r
}

func(m *PB_UserParam_BlockedList)ToFlat() *PB_UserParam_BlockedList_Flat {
r := &PB_UserParam_BlockedList_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserResponse_BlockedList)ToFlat() *PB_UserResponse_BlockedList_Flat {
r := &PB_UserResponse_BlockedList_Flat{
    ByUserId:  int(m.ByUserId) ,
    TargetUserId:  int(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UpdateAbout)ToFlat() *PB_UserParam_UpdateAbout_Flat {
r := &PB_UserParam_UpdateAbout_Flat{
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserOfflineResponse_UpdateAbout)ToFlat() *PB_UserOfflineResponse_UpdateAbout_Flat {
r := &PB_UserOfflineResponse_UpdateAbout_Flat{
    UserId:  int(m.UserId) ,
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserParam_UpdateUserName)ToFlat() *PB_UserParam_UpdateUserName_Flat {
r := &PB_UserParam_UpdateUserName_Flat{
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserOfflineResponse_UpdateUserName)ToFlat() *PB_UserOfflineResponse_UpdateUserName_Flat {
r := &PB_UserOfflineResponse_UpdateUserName_Flat{
    UserId:  int(m.UserId) ,
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserParam_ChangeAvatar)ToFlat() *PB_UserParam_ChangeAvatar_Flat {
r := &PB_UserParam_ChangeAvatar_Flat{
    None:  m.None ,
    ImageData2:  []byte(m.ImageData2) ,
}
return r
}

func(m *PB_UserOfflineResponse_ChangeAvatar)ToFlat() *PB_UserOfflineResponse_ChangeAvatar_Flat {
r := &PB_UserOfflineResponse_ChangeAvatar_Flat{
}
return r
}

func(m *PB_UserParam_ChangePrivacy)ToFlat() *PB_UserParam_ChangePrivacy_Flat {
r := &PB_UserParam_ChangePrivacy_Flat{
    
}
return r
}

func(m *PB_UserResponseOffline_ChangePrivacy)ToFlat() *PB_UserResponseOffline_ChangePrivacy_Flat {
r := &PB_UserResponseOffline_ChangePrivacy_Flat{
}
return r
}

func(m *PB_UserParam_CheckUserName)ToFlat() *PB_UserParam_CheckUserName_Flat {
r := &PB_UserParam_CheckUserName_Flat{
    
}
return r
}

func(m *PB_UserResponse_CheckUserName)ToFlat() *PB_UserResponse_CheckUserName_Flat {
r := &PB_UserResponse_CheckUserName_Flat{
}
return r
}

func(m *UserView)ToFlat() *UserView_Flat {
r := &UserView_Flat{
}
return r
}

func(m *PB_Activity)ToFlat() *PB_Activity_Flat {
r := &PB_Activity_Flat{
    Id:  int(m.Id) ,
    ActorUserId:  int(m.ActorUserId) ,
    ActionTypeId:  int(m.ActionTypeId) ,
    RowId:  int(m.RowId) ,
    RootId:  int(m.RootId) ,
    RefId:  int(m.RefId) ,
    CreatedAt:  int(m.CreatedAt) ,
}
return r
}

func(m *PB_Bucket)ToFlat() *PB_Bucket_Flat {
r := &PB_Bucket_Flat{
    BucketId:  int(m.BucketId) ,
    BucketName:  m.BucketName ,
    Server1Id:  int(m.Server1Id) ,
    Server2Id:  int(m.Server2Id) ,
    BackupServerId:  int(m.BackupServerId) ,
    ContentObjectTypeId:  int(m.ContentObjectTypeId) ,
    ContentObjectCount:  int(m.ContentObjectCount) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Chat)ToFlat() *PB_Chat_Flat {
r := &PB_Chat_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedSe:  int(m.CreatedSe) ,
    StartMessageIdFrom:  int(m.StartMessageIdFrom) ,
    LastDeletedMessageId:  int(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int(m.LastSeenMessageId) ,
    LastMessageId:  int(m.LastMessageId) ,
    UpdatedMs:  int(m.UpdatedMs) ,
}
return r
}

func(m *PB_Comment)ToFlat() *PB_Comment_Flat {
r := &PB_Comment_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    Text:  m.Text ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_DirectMessage)ToFlat() *PB_DirectMessage_Flat {
r := &PB_DirectMessage_Flat{
    MessageId:  int(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnumId:  int(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int(m.CreatedSe) ,
    PeerReceivedTime:  int(m.PeerReceivedTime) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int(m.DeliviryStatusEnumId) ,
}
return r
}

func(m *PB_DirectOffline)ToFlat() *PB_DirectOffline_Flat {
r := &PB_DirectOffline_Flat{
    DirectOfflineId:  int(m.DirectOfflineId) ,
    ToUserId:  int(m.ToUserId) ,
    MessageId:  int(m.MessageId) ,
    MessageFileId:  int(m.MessageFileId) ,
    OtherId:  int(m.OtherId) ,
    ChatKey:  m.ChatKey ,
    PeerUserId:  int(m.PeerUserId) ,
    RoomLogTypeId:  int(m.RoomLogTypeId) ,
    PBClass:  m.PBClass ,
    DataPB:  []byte(m.DataPB) ,
    DataJson:  m.DataJson ,
    DataTemp:  m.DataTemp ,
    AtTimeMs:  int(m.AtTimeMs) ,
}
return r
}

func(m *PB_DirectToMessage)ToFlat() *PB_DirectToMessage_Flat {
r := &PB_DirectToMessage_Flat{
    Id:  int(m.Id) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int(m.MessageId) ,
    SourceEnumId:  int(m.SourceEnumId) ,
}
return r
}

func(m *PB_DirectUpdate)ToFlat() *PB_DirectUpdate_Flat {
r := &PB_DirectUpdate_Flat{
    DirectUpdateId:  int(m.DirectUpdateId) ,
    ToUserId:  int(m.ToUserId) ,
    MessageId:  int(m.MessageId) ,
    MessageFileId:  int(m.MessageFileId) ,
    OtherId:  int(m.OtherId) ,
    ChatKey:  m.ChatKey ,
    PeerUserId:  int(m.PeerUserId) ,
    EventType:  int(m.EventType) ,
    RoomLogTypeId:  int(m.RoomLogTypeId) ,
    FromSeq:  int(m.FromSeq) ,
    ToSeq:  int(m.ToSeq) ,
    ExtraPB:  []byte(m.ExtraPB) ,
    ExtraJson:  m.ExtraJson ,
    AtTimeMs:  int(m.AtTimeMs) ,
}
return r
}

func(m *PB_FollowingList)ToFlat() *PB_FollowingList_Flat {
r := &PB_FollowingList_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    ListType:  int(m.ListType) ,
    Name:  m.Name ,
    Count:  int(m.Count) ,
    IsAuto:  int(m.IsAuto) ,
    IsPimiry:  int(m.IsPimiry) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_FollowingListMember)ToFlat() *PB_FollowingListMember_Flat {
r := &PB_FollowingListMember_Flat{
    Id:  int(m.Id) ,
    ListId:  int(m.ListId) ,
    UserId:  int(m.UserId) ,
    FollowedUserId:  int(m.FollowedUserId) ,
    FollowType:  int(m.FollowType) ,
    UpdatedTimeMs:  int(m.UpdatedTimeMs) ,
}
return r
}

func(m *PB_FollowingListMemberHistory)ToFlat() *PB_FollowingListMemberHistory_Flat {
r := &PB_FollowingListMemberHistory_Flat{
    Id:  int(m.Id) ,
    ListId:  int(m.ListId) ,
    UserId:  int(m.UserId) ,
    FollowedUserId:  int(m.FollowedUserId) ,
    FollowType:  int(m.FollowType) ,
    UpdatedTimeMs:  int(m.UpdatedTimeMs) ,
    FollowId:  int(m.FollowId) ,
}
return r
}

func(m *PB_GeneralLog)ToFlat() *PB_GeneralLog_Flat {
r := &PB_GeneralLog_Flat{
    Id:  int(m.Id) ,
    ToUserId:  int(m.ToUserId) ,
    TargetId:  int(m.TargetId) ,
    LogTypeId:  int(m.LogTypeId) ,
    ExtraPB:  []byte(m.ExtraPB) ,
    ExtraJson:  m.ExtraJson ,
    CreatedMs:  int(m.CreatedMs) ,
}
return r
}

func(m *PB_Group)ToFlat() *PB_Group_Flat {
r := &PB_Group_Flat{
    GroupId:  int(m.GroupId) ,
    GroupName:  m.GroupName ,
    MembersCount:  int(m.MembersCount) ,
    GroupPrivacyEnum:  int(m.GroupPrivacyEnum) ,
    CreatorUserId:  int(m.CreatorUserId) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    CurrentSeq:  int(m.CurrentSeq) ,
}
return r
}

func(m *PB_GroupMember)ToFlat() *PB_GroupMember_Flat {
r := &PB_GroupMember_Flat{
    Id:  int(m.Id) ,
    GroupId:  int(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    UserId:  int(m.UserId) ,
    ByUserId:  int(m.ByUserId) ,
    GroupRoleEnumId:  int(m.GroupRoleEnumId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_GroupMessage)ToFlat() *PB_GroupMessage_Flat {
r := &PB_GroupMessage_Flat{
    MessageId:  int(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnum:  int(m.MessageTypeEnum) ,
    Text:  m.Text ,
    CreatedMs:  int(m.CreatedMs) ,
    DeliveryStatusEnum:  int(m.DeliveryStatusEnum) ,
}
return r
}

func(m *PB_GroupToMessage)ToFlat() *PB_GroupToMessage_Flat {
r := &PB_GroupToMessage_Flat{
    Id:  int(m.Id) ,
    GroupId:  int(m.GroupId) ,
    MessageId:  int(m.MessageId) ,
    CreatedMs:  int(m.CreatedMs) ,
    StatusEnum:  int(m.StatusEnum) ,
}
return r
}

func(m *PB_Like)ToFlat() *PB_Like_Flat {
r := &PB_Like_Flat{
    Id:  int(m.Id) ,
    PostId:  int(m.PostId) ,
    PostTypeId:  int(m.PostTypeId) ,
    UserId:  int(m.UserId) ,
    TypeId:  int(m.TypeId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_LogChange)ToFlat() *PB_LogChange_Flat {
r := &PB_LogChange_Flat{
    Id:  int(m.Id) ,
    T:  m.T ,
}
return r
}

func(m *PB_Media)ToFlat() *PB_Media_Flat {
r := &PB_Media_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    AlbumId:  int(m.AlbumId) ,
    TypeId:  int(m.TypeId) ,
    CreatedTime:  int(m.CreatedTime) ,
    Src:  m.Src ,
}
return r
}

func(m *PB_MessageFile)ToFlat() *PB_MessageFile_Flat {
r := &PB_MessageFile_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileTypeEnumId:  int(m.FileTypeEnumId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int(m.HashAccess) ,
    CreatedSe:  int(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
}
return r
}

func(m *PB_Notification)ToFlat() *PB_Notification_Flat {
r := &PB_Notification_Flat{
    Id:  int(m.Id) ,
    ForUserId:  int(m.ForUserId) ,
    ActorUserId:  int(m.ActorUserId) ,
    ActionTypeId:  int(m.ActionTypeId) ,
    ObjectTypeId:  int(m.ObjectTypeId) ,
    RowId:  int(m.RowId) ,
    RootId:  int(m.RootId) ,
    RefId:  int(m.RefId) ,
    SeenStatus:  int(m.SeenStatus) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_NotificationRemoved)ToFlat() *PB_NotificationRemoved_Flat {
r := &PB_NotificationRemoved_Flat{
    NotificationId:  int(m.NotificationId) ,
    ForUserId:  int(m.ForUserId) ,
}
return r
}

func(m *PB_Offline)ToFlat() *PB_Offline_Flat {
r := &PB_Offline_Flat{
    Id:  int(m.Id) ,
    FromUserId:  int(m.FromUserId) ,
    ToUserId:  int(m.ToUserId) ,
    RpcName:  m.RpcName ,
    PBClassName:  m.PBClassName ,
    Key:  m.Key ,
    DataJson:  m.DataJson ,
    DataBlob:  []byte(m.DataBlob) ,
    DataLength:  int(m.DataLength) ,
    CreatedMs:  int(m.CreatedMs) ,
}
return r
}

func(m *PB_OldMessage)ToFlat() *PB_OldMessage_Flat {
r := &PB_OldMessage_Flat{
    Id:  int(m.Id) ,
    Uid:  int(m.Uid) ,
    UserId:  int(m.UserId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    MessageType:  int(m.MessageType) ,
    RoomType:  int(m.RoomType) ,
    MsgFileId:  int(m.MsgFileId) ,
    DataPB:  []byte(m.DataPB) ,
    Data64:  m.Data64 ,
    DataJson:  m.DataJson ,
    CreatedTimeMs:  int(m.CreatedTimeMs) ,
}
return r
}

func(m *PB_OldMsgFile)ToFlat() *PB_OldMsgFile_Flat {
r := &PB_OldMsgFile_Flat{
    Id:  int(m.Id) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileType:  int(m.FileType) ,
    MimeType:  m.MimeType ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    ThumbData:  []byte(m.ThumbData) ,
    ThumbData64:  m.ThumbData64 ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
}
return r
}

func(m *PB_OldMsgPush)ToFlat() *PB_OldMsgPush_Flat {
r := &PB_OldMsgPush_Flat{
    Id:  int(m.Id) ,
    Uid:  int(m.Uid) ,
    ToUser:  int(m.ToUser) ,
    MsgUid:  int(m.MsgUid) ,
    CreatedTimeMs:  int(m.CreatedTimeMs) ,
}
return r
}

func(m *PB_OldMsgPushEvent)ToFlat() *PB_OldMsgPushEvent_Flat {
r := &PB_OldMsgPushEvent_Flat{
    Id:  int(m.Id) ,
    Uid:  int(m.Uid) ,
    ToUserId:  int(m.ToUserId) ,
    MsgUid:  int(m.MsgUid) ,
    MsgKey:  m.MsgKey ,
    RoomKey:  m.RoomKey ,
    PeerUserId:  int(m.PeerUserId) ,
    EventType:  int(m.EventType) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_PhoneContact)ToFlat() *PB_PhoneContact_Flat {
r := &PB_PhoneContact_Flat{
    Id:  int(m.Id) ,
    PhoneDisplayName:  m.PhoneDisplayName ,
    PhoneFamilyName:  m.PhoneFamilyName ,
    PhoneNumber:  m.PhoneNumber ,
    PhoneNormalizedNumber:  m.PhoneNormalizedNumber ,
    PhoneContactRowId:  int(m.PhoneContactRowId) ,
    UserId:  int(m.UserId) ,
    DeviceUuidId:  int(m.DeviceUuidId) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedTime:  int(m.UpdatedTime) ,
}
return r
}

func(m *PB_Photo)ToFlat() *PB_Photo_Flat {
r := &PB_Photo_Flat{
    PhotoId:  int(m.PhotoId) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    AlbumId:  int(m.AlbumId) ,
    ImageTypeId:  int(m.ImageTypeId) ,
    Title:  m.Title ,
    Src:  m.Src ,
    PathSrc:  m.PathSrc ,
    BucketId:  int(m.BucketId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Ratio:  float32(m.Ratio) ,
    HashMd5:  m.HashMd5 ,
    Color:  m.Color ,
    CreatedTime:  int(m.CreatedTime) ,
    W1080:  int(m.W1080) ,
    W720:  int(m.W720) ,
    W480:  int(m.W480) ,
    W320:  int(m.W320) ,
    W160:  int(m.W160) ,
    W80:  int(m.W80) ,
}
return r
}

func(m *PB_Post)ToFlat() *PB_Post_Flat {
r := &PB_Post_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    TypeId:  int(m.TypeId) ,
    Text:  m.Text ,
    FormatedText:  m.FormatedText ,
    MediaCount:  int(m.MediaCount) ,
    SharedTo:  int(m.SharedTo) ,
    DisableComment:  int(m.DisableComment) ,
    HasTag:  int(m.HasTag) ,
    LikesCount:  int(m.LikesCount) ,
    CommentsCount:  int(m.CommentsCount) ,
    EditedTime:  int(m.EditedTime) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_PushEvent)ToFlat() *PB_PushEvent_Flat {
r := &PB_PushEvent_Flat{
    PushEventId:  int(m.PushEventId) ,
    ToUserId:  int(m.ToUserId) ,
    ToDeviceId:  int(m.ToDeviceId) ,
    MessageId:  int(m.MessageId) ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
    RoomId:  int(m.RoomId) ,
    PeerUserId:  int(m.PeerUserId) ,
    PushEventTypeEnum:  int(m.PushEventTypeEnum) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_PushMessage)ToFlat() *PB_PushMessage_Flat {
r := &PB_PushMessage_Flat{
    PushMessageId:  int(m.PushMessageId) ,
    ToUserId:  int(m.ToUserId) ,
    ToDeviceId:  int(m.ToDeviceId) ,
    MessageId:  int(m.MessageId) ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
    CreatedMs:  int(m.CreatedMs) ,
}
return r
}

func(m *PB_RecommendUser)ToFlat() *PB_RecommendUser_Flat {
r := &PB_RecommendUser_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    TargetId:  int(m.TargetId) ,
    Weight:  float32(m.Weight) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Room)ToFlat() *PB_Room_Flat {
r := &PB_Room_Flat{
    RoomId:  int(m.RoomId) ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
    UserId:  int(m.UserId) ,
    LastSeqSeen:  int(m.LastSeqSeen) ,
    LastSeqDelete:  int(m.LastSeqDelete) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedTime:  int(m.CreatedTime) ,
    CurrentSeq:  int(m.CurrentSeq) ,
}
return r
}

func(m *PB_SearchClicked)ToFlat() *PB_SearchClicked_Flat {
r := &PB_SearchClicked_Flat{
    Id:  int(m.Id) ,
    Query:  m.Query ,
    ClickType:  int(m.ClickType) ,
    TargetId:  int(m.TargetId) ,
    UserId:  int(m.UserId) ,
    CreatedAt:  int(m.CreatedAt) ,
}
return r
}

func(m *PB_Session)ToFlat() *PB_Session_Flat {
r := &PB_Session_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    SessionUuid:  m.SessionUuid ,
    ClientUuid:  m.ClientUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastActivityTime:  int(m.LastActivityTime) ,
    LastIpAddress:  m.LastIpAddress ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    LastNetworkTypeId:  int(m.LastNetworkTypeId) ,
    AppVersion:  int(m.AppVersion) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    LastSyncDirectUpdateId:  int(m.LastSyncDirectUpdateId) ,
    LastSyncGeneralUpdateId:  int(m.LastSyncGeneralUpdateId) ,
    LastSyncNotifyUpdateId:  int(m.LastSyncNotifyUpdateId) ,
}
return r
}

func(m *PB_SettingClient)ToFlat() *PB_SettingClient_Flat {
r := &PB_SettingClient_Flat{
    UserId:  int(m.UserId) ,
    AutoDownloadWifiVoice:  int(m.AutoDownloadWifiVoice) ,
    AutoDownloadWifiImage:  int(m.AutoDownloadWifiImage) ,
    AutoDownloadWifiVideo:  int(m.AutoDownloadWifiVideo) ,
    AutoDownloadWifiFile:  int(m.AutoDownloadWifiFile) ,
    AutoDownloadWifiMusic:  int(m.AutoDownloadWifiMusic) ,
    AutoDownloadWifiGif:  int(m.AutoDownloadWifiGif) ,
    AutoDownloadCellularVoice:  int(m.AutoDownloadCellularVoice) ,
    AutoDownloadCellularImage:  int(m.AutoDownloadCellularImage) ,
    AutoDownloadCellularVideo:  int(m.AutoDownloadCellularVideo) ,
    AutoDownloadCellularFile:  int(m.AutoDownloadCellularFile) ,
    AutoDownloadCellularMusic:  int(m.AutoDownloadCellularMusic) ,
    AutoDownloadCellularGif:  int(m.AutoDownloadCellularGif) ,
    AutoDownloadRoamingVoice:  int(m.AutoDownloadRoamingVoice) ,
    AutoDownloadRoamingImage:  int(m.AutoDownloadRoamingImage) ,
    AutoDownloadRoamingVideo:  int(m.AutoDownloadRoamingVideo) ,
    AutoDownloadRoamingFile:  int(m.AutoDownloadRoamingFile) ,
    AutoDownloadRoamingMusic:  int(m.AutoDownloadRoamingMusic) ,
    AutoDownloadRoamingGif:  int(m.AutoDownloadRoamingGif) ,
    SaveToGallery:  int(m.SaveToGallery) ,
}
return r
}

func(m *PB_SettingNotification)ToFlat() *PB_SettingNotification_Flat {
r := &PB_SettingNotification_Flat{
    UserId:  int(m.UserId) ,
    SocialLedOn:  int(m.SocialLedOn) ,
    SocialLedColor:  m.SocialLedColor ,
    ReqestToFollowYou:  int(m.ReqestToFollowYou) ,
    FollowedYou:  int(m.FollowedYou) ,
    AccptedYourFollowRequest:  int(m.AccptedYourFollowRequest) ,
    YourPostLiked:  int(m.YourPostLiked) ,
    YourPostCommented:  int(m.YourPostCommented) ,
    MenthenedYouInPost:  int(m.MenthenedYouInPost) ,
    MenthenedYouInComment:  int(m.MenthenedYouInComment) ,
    YourContactsJoined:  int(m.YourContactsJoined) ,
    DirectMessage:  int(m.DirectMessage) ,
    DirectAlert:  int(m.DirectAlert) ,
    DirectPerview:  int(m.DirectPerview) ,
    DirectLedOn:  int(m.DirectLedOn) ,
    DirectLedColor:  int(m.DirectLedColor) ,
    DirectVibrate:  int(m.DirectVibrate) ,
    DirectPopup:  int(m.DirectPopup) ,
    DirectSound:  int(m.DirectSound) ,
    DirectPriority:  int(m.DirectPriority) ,
}
return r
}

func(m *PB_Tag)ToFlat() *PB_Tag_Flat {
r := &PB_Tag_Flat{
    Id:  int(m.Id) ,
    Name:  m.Name ,
    Count:  int(m.Count) ,
    IsBlocked:  int(m.IsBlocked) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_TagsPost)ToFlat() *PB_TagsPost_Flat {
r := &PB_TagsPost_Flat{
    Id:  int(m.Id) ,
    TagId:  int(m.TagId) ,
    PostId:  int(m.PostId) ,
    TypeId:  int(m.TypeId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_TestChat)ToFlat() *PB_TestChat_Flat {
r := &PB_TestChat_Flat{
    Id:  int(m.Id) ,
    Id4:  int(m.Id4) ,
    TimeMs:  int(m.TimeMs) ,
    Text:  m.Text ,
    Name:  m.Name ,
    UserId:  int(m.UserId) ,
    C2:  int(m.C2) ,
    C3:  int(m.C3) ,
    C4:  int(m.C4) ,
    C5:  int(m.C5) ,
}
return r
}

func(m *PB_TriggerLog)ToFlat() *PB_TriggerLog_Flat {
r := &PB_TriggerLog_Flat{
    Id:  int(m.Id) ,
    ModelName:  m.ModelName ,
    ChangeType:  m.ChangeType ,
    TargetId:  int(m.TargetId) ,
    TargetStr:  m.TargetStr ,
    CreatedSe:  int(m.CreatedSe) ,
}
return r
}

func(m *PB_User)ToFlat() *PB_User_Flat {
r := &PB_User_Flat{
    Id:  int(m.Id) ,
    UserName:  m.UserName ,
    UserNameLower:  m.UserNameLower ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int(m.PrivacyProfile) ,
    Phone:  m.Phone ,
    Email:  m.Email ,
    IsDeleted:  int(m.IsDeleted) ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    MediaCount:  int(m.MediaCount) ,
    LikesCount:  int(m.LikesCount) ,
    ResharedCount:  int(m.ResharedCount) ,
    LastActionTime:  int(m.LastActionTime) ,
    LastPostTime:  int(m.LastPostTime) ,
    PrimaryFollowingList:  int(m.PrimaryFollowingList) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    SessionUuid:  m.SessionUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    AppVersion:  int(m.AppVersion) ,
    LastActivityTime:  int(m.LastActivityTime) ,
    LastLoginTime:  int(m.LastLoginTime) ,
    LastIpAddress:  m.LastIpAddress ,
}
return r
}

func(m *PB_UserMetaInfo)ToFlat() *PB_UserMetaInfo_Flat {
r := &PB_UserMetaInfo_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    IsNotificationDirty:  int(m.IsNotificationDirty) ,
    LastUserRecGen:  int(m.LastUserRecGen) ,
}
return r
}

func(m *PB_UserPassword)ToFlat() *PB_UserPassword_Flat {
r := &PB_UserPassword_Flat{
    UserId:  int(m.UserId) ,
    Password:  m.Password ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_UpdateGroupParticipants)ToFlat() *PB_UpdateGroupParticipants_Flat {
r := &PB_UpdateGroupParticipants_Flat{
}
return r
}

func(m *PB_UpdateNotifySettings)ToFlat() *PB_UpdateNotifySettings_Flat {
r := &PB_UpdateNotifySettings_Flat{
}
return r
}

func(m *PB_UpdateServiceNotification)ToFlat() *PB_UpdateServiceNotification_Flat {
r := &PB_UpdateServiceNotification_Flat{
}
return r
}

func(m *PB_UpdateMessageMeta)ToFlat() *PB_UpdateMessageMeta_Flat {
r := &PB_UpdateMessageMeta_Flat{
    MessageId:  int(m.MessageId) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_UpdateMessageId)ToFlat() *PB_UpdateMessageId_Flat {
r := &PB_UpdateMessageId_Flat{
    OldMessageId:  int(m.OldMessageId) ,
    NewMessageId:  int(m.NewMessageId) ,
}
return r
}

func(m *PB_UpdateMessageToEdit)ToFlat() *PB_UpdateMessageToEdit_Flat {
r := &PB_UpdateMessageToEdit_Flat{
    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_UpdateMessageToDelete)ToFlat() *PB_UpdateMessageToDelete_Flat {
r := &PB_UpdateMessageToDelete_Flat{
    MessageId:  int(m.MessageId) ,
}
return r
}

func(m *PB_UpdateRoomActionDoing)ToFlat() *PB_UpdateRoomActionDoing_Flat {
r := &PB_UpdateRoomActionDoing_Flat{
    RoomKey:  m.RoomKey ,
    
}
return r
}

func(m *PB_UpdateUserBlocked)ToFlat() *PB_UpdateUserBlocked_Flat {
r := &PB_UpdateUserBlocked_Flat{
    UserId:  int(m.UserId) ,
    Blocked:  m.Blocked ,
}
return r
}

func(m *PB_ChatView)ToFlat() *PB_ChatView_Flat {
r := &PB_ChatView_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedSe:  int(m.CreatedSe) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    LastMessageId:  int(m.LastMessageId) ,
    LastDeletedMessageId:  int(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int(m.LastSeenMessageId) ,
    LastSeqSeen:  int(m.LastSeqSeen) ,
    LastSeqDelete:  int(m.LastSeqDelete) ,
    CurrentSeq:  int(m.CurrentSeq) ,
    
    SharedMediaCount:  int(m.SharedMediaCount) ,
    UnseenCount:  int(m.UnseenCount) ,
    
    
}
return r
}

func(m *PB_MessageView)ToFlat() *PB_MessageView_Flat {
r := &PB_MessageView_Flat{
    MessageId:  int(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnumId:  int(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int(m.CreatedSe) ,
    PeerReceivedTime:  int(m.PeerReceivedTime) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int(m.DeliviryStatusEnumId) ,
    ChatKey:  m.ChatKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    IsByMe:  m.IsByMe ,
    RemoteId:  int(m.RemoteId) ,
    
}
return r
}

func(m *PB_MessageFileView)ToFlat() *PB_MessageFileView_Flat {
r := &PB_MessageFileView_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileTypeEnumId:  int(m.FileTypeEnumId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int(m.HashAccess) ,
    CreatedSe:  int(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
    ServerThumbLocalSrc:  m.ServerThumbLocalSrc ,
    RemoteMessageFileId:  int(m.RemoteMessageFileId) ,
    LocalSrc:  m.LocalSrc ,
    ThumbLocalSrc:  m.ThumbLocalSrc ,
    MessageFileStatusId:  m.MessageFileStatusId ,
}
return r
}

func(m *PB_UserView)ToFlat() *PB_UserView_Flat {
r := &PB_UserView_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int(m.PrivacyProfile) ,
    IsDeleted:  int(m.IsDeleted) ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    AppVersion:  int(m.AppVersion) ,
    LastActivityTime:  int(m.LastActivityTime) ,
    FollowingType:  int(m.FollowingType) ,
}
return r
}



///// from_flat ///

func(m *GeoLocation_Flat)ToPB() *GeoLocation {
r := &GeoLocation{
    Lat:  m.Lat ,
    Lon:  m.Lon ,
}
return r
}

func(m *RoomMessageLog_Flat)ToPB() *RoomMessageLog {
r := &RoomMessageLog{
    
    TargetUserId:  uint64(m.TargetUserId) ,
    ByUserId:  uint64(m.ByUserId) ,
}
return r
}

func(m *RoomMessageForwardFrom_Flat)ToPB() *RoomMessageForwardFrom {
r := &RoomMessageForwardFrom{
    RoomId:  uint64(m.RoomId) ,
    MessageId:  uint64(m.MessageId) ,
    RoomTypeEnum:  uint32(m.RoomTypeEnum) ,
}
return r
}

func(m *RoomDraft_Flat)ToPB() *RoomDraft {
r := &RoomDraft{
    Message:  m.Message ,
    ReplyTo:  uint64(m.ReplyTo) ,
}
return r
}

func(m *ChatRoom_Flat)ToPB() *ChatRoom {
r := &ChatRoom{
}
return r
}

func(m *Pagination_Flat)ToPB() *Pagination {
r := &Pagination{
    Offset:  uint32(m.Offset) ,
    Limit:  uint32(m.Limit) ,
}
return r
}

func(m *PB_CommandToServer_Flat)ToPB() *PB_CommandToServer {
r := &PB_CommandToServer{
    ClientCallId:  int64(m.ClientCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  m.Data ,
}
return r
}

func(m *PB_CommandToClient_Flat)ToPB() *PB_CommandToClient {
r := &PB_CommandToClient{
    ServerCallId:  int64(m.ServerCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  m.Data ,
}
return r
}

func(m *PB_CommandReachedToServer_Flat)ToPB() *PB_CommandReachedToServer {
r := &PB_CommandReachedToServer{
    ClientCallId:  int64(m.ClientCallId) ,
}
return r
}

func(m *PB_CommandReachedToClient_Flat)ToPB() *PB_CommandReachedToClient {
r := &PB_CommandReachedToClient{
    ServerCallId:  int64(m.ServerCallId) ,
}
return r
}

func(m *PB_ResponseToClient_Flat)ToPB() *PB_ResponseToClient {
r := &PB_ResponseToClient{
    ClientCallId:  int64(m.ClientCallId) ,
    PBClass:  m.PBClass ,
    RpcFullName:  m.RpcFullName ,
    Data:  m.Data ,
}
return r
}

func(m *PB_Offline_NewDirectMessage_Flat)ToPB() *PB_Offline_NewDirectMessage {
r := &PB_Offline_NewDirectMessage{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int64(m.FromMessageId) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_Offline_MessagesReachedServer_Flat)ToPB() *PB_Offline_MessagesReachedServer {
r := &PB_Offline_MessagesReachedServer{
    MessageKeys:  m.MessageKeys ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_Offline_MessagesDeliveredToUser_Flat)ToPB() *PB_Offline_MessagesDeliveredToUser {
r := &PB_Offline_MessagesDeliveredToUser{
    MessageKeys:  m.MessageKeys ,
    RoomKey:  m.RoomKey ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_Offline_MessagesSeenByPeer_Flat)ToPB() *PB_Offline_MessagesSeenByPeer {
r := &PB_Offline_MessagesSeenByPeer{
    MessageKeys:  m.MessageKeys ,
    RoomKey:  m.RoomKey ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_Offline_MessagesDeletedFromServer_Flat)ToPB() *PB_Offline_MessagesDeletedFromServer {
r := &PB_Offline_MessagesDeletedFromServer{
    MessageKeys:  m.MessageKeys ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_Offline_ChangeMessageId_Flat)ToPB() *PB_Offline_ChangeMessageId {
r := &PB_Offline_ChangeMessageId{
    MessageKey:  int64(m.MessageKey) ,
    NewMessageId:  int64(m.NewMessageId) ,
}
return r
}

func(m *PB_Offline_ChangeMessageFileId_Flat)ToPB() *PB_Offline_ChangeMessageFileId {
r := &PB_Offline_ChangeMessageFileId{
    MessageFileKey:  int64(m.MessageFileKey) ,
    NewMessageFileId:  int64(m.NewMessageFileId) ,
}
return r
}

func(m *PB_Offline_MessageToEdit_Flat)ToPB() *PB_Offline_MessageToEdit {
r := &PB_Offline_MessageToEdit{
    MessageKey:  int64(m.MessageKey) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_Offline_MessageToDelete_Flat)ToPB() *PB_Offline_MessageToDelete {
r := &PB_Offline_MessageToDelete{
    MessageKey:  int64(m.MessageKey) ,
}
return r
}

func(m *PB_Online_RoomActionDoing_Flat)ToPB() *PB_Online_RoomActionDoing {
r := &PB_Online_RoomActionDoing{
    RoomKey:  m.RoomKey ,
    
}
return r
}

func(m *PB_Offline_Messagings_Flat)ToPB() *PB_Offline_Messagings {
r := &PB_Offline_Messagings{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_UserParam_CheckUserName2_Flat)ToPB() *PB_UserParam_CheckUserName2 {
r := &PB_UserParam_CheckUserName2{
}
return r
}

func(m *PB_UserResponse_CheckUserName2_Flat)ToPB() *PB_UserResponse_CheckUserName2 {
r := &PB_UserResponse_CheckUserName2{
}
return r
}

func(m *PB_ChatParam_AddNewMessage_Flat)ToPB() *PB_ChatParam_AddNewMessage {
r := &PB_ChatParam_AddNewMessage{
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,
    Blob:  m.Blob ,
    
}
return r
}

func(m *PB_ChatResponse_AddNewMessage_Flat)ToPB() *PB_ChatResponse_AddNewMessage {
r := &PB_ChatResponse_AddNewMessage{
}
return r
}

func(m *PB_ChatParam_SetRoomActionDoing_Flat)ToPB() *PB_ChatParam_SetRoomActionDoing {
r := &PB_ChatParam_SetRoomActionDoing{
    GroupId:  int64(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,
    
}
return r
}

func(m *PB_ChatResponse_SetRoomActionDoing_Flat)ToPB() *PB_ChatResponse_SetRoomActionDoing {
r := &PB_ChatResponse_SetRoomActionDoing{
}
return r
}

func(m *PB_ChatParam_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_ChatParam_SetChatMessagesRangeAsSeen {
r := &PB_ChatParam_SetChatMessagesRangeAsSeen{
    ChatKey:  m.ChatKey ,
    BottomMessageId:  int64(m.BottomMessageId) ,
    TopMessageId:  int64(m.TopMessageId) ,
    SeenTimeMs:  int64(m.SeenTimeMs) ,
}
return r
}

func(m *PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_ChatResponse_SetChatMessagesRangeAsSeen {
r := &PB_ChatResponse_SetChatMessagesRangeAsSeen{
}
return r
}

func(m *PB_ChatParam_DeleteChatHistory_Flat)ToPB() *PB_ChatParam_DeleteChatHistory {
r := &PB_ChatParam_DeleteChatHistory{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int64(m.FromMessageId) ,
}
return r
}

func(m *PB_ChatResponse_DeleteChatHistory_Flat)ToPB() *PB_ChatResponse_DeleteChatHistory {
r := &PB_ChatResponse_DeleteChatHistory{
}
return r
}

func(m *PB_ChatParam_DeleteMessagesByIds_Flat)ToPB() *PB_ChatParam_DeleteMessagesByIds {
r := &PB_ChatParam_DeleteMessagesByIds{
    ChatKey:  m.ChatKey ,
    Both:  m.Both ,
    MessagesIds:  helper.SliceIntToInt64(m.MessagesIds) ,
}
return r
}

func(m *PB_ChatResponse_DeleteMessagesByIds_Flat)ToPB() *PB_ChatResponse_DeleteMessagesByIds {
r := &PB_ChatResponse_DeleteMessagesByIds{
}
return r
}

func(m *PB_ChatParam_SetMessagesAsReceived_Flat)ToPB() *PB_ChatParam_SetMessagesAsReceived {
r := &PB_ChatParam_SetMessagesAsReceived{
    ChatRoom:  m.ChatRoom ,
    MessageIds:  helper.SliceIntToInt64(m.MessageIds) ,
}
return r
}

func(m *PB_ChatResponse_SetMessagesAsReceived_Flat)ToPB() *PB_ChatResponse_SetMessagesAsReceived {
r := &PB_ChatResponse_SetMessagesAsReceived{
}
return r
}

func(m *PB_ChatParam_EditMessage_Flat)ToPB() *PB_ChatParam_EditMessage {
r := &PB_ChatParam_EditMessage{
    RoomKey:  m.RoomKey ,
    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_ChatResponse_EditMessage_Flat)ToPB() *PB_ChatResponse_EditMessage {
r := &PB_ChatResponse_EditMessage{
}
return r
}

func(m *PB_ChatParam_GetChatList_Flat)ToPB() *PB_ChatParam_GetChatList {
r := &PB_ChatParam_GetChatList{
}
return r
}

func(m *PB_ChatResponse_GetChatList_Flat)ToPB() *PB_ChatResponse_GetChatList {
r := &PB_ChatResponse_GetChatList{
    
    
}
return r
}

func(m *PB_ChatParam_GetChatHistoryToOlder_Flat)ToPB() *PB_ChatParam_GetChatHistoryToOlder {
r := &PB_ChatParam_GetChatHistoryToOlder{
    ChatKey:  m.ChatKey ,
    Limit:  int32(m.Limit) ,
    FromTopMessageId:  int64(m.FromTopMessageId) ,
}
return r
}

func(m *PB_ChatResponse_GetChatHistoryToOlder_Flat)ToPB() *PB_ChatResponse_GetChatHistoryToOlder {
r := &PB_ChatResponse_GetChatHistoryToOlder{
    
    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_ChatParam_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_ChatParam_GetFreshAllDirectMessagesList {
r := &PB_ChatParam_GetFreshAllDirectMessagesList{
    LowMessageId:  int64(m.LowMessageId) ,
    HighMessageId:  int64(m.HighMessageId) ,
}
return r
}

func(m *PB_ChatResponse_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_ChatResponse_GetFreshAllDirectMessagesList {
r := &PB_ChatResponse_GetFreshAllDirectMessagesList{
    
    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_MsgParam_AddNewTextMessage_Flat)ToPB() *PB_MsgParam_AddNewTextMessage {
r := &PB_MsgParam_AddNewTextMessage{
    Text:  m.Text ,
    MessageKey:  m.MessageKey ,
    ToRoomKey:  m.ToRoomKey ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,
    
}
return r
}

func(m *PB_MsgResponse_AddNewTextMessage_Flat)ToPB() *PB_MsgResponse_AddNewTextMessage {
r := &PB_MsgResponse_AddNewTextMessage{
}
return r
}

func(m *PB_MsgParam_AddNewMessage_Flat)ToPB() *PB_MsgParam_AddNewMessage {
r := &PB_MsgParam_AddNewMessage{
    Text:  m.Text ,
    MessageKey:  m.MessageKey ,
    ToRoomKey:  m.ToRoomKey ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,
    
    
    Blob:  m.Blob ,
    
}
return r
}

func(m *PB_MsgResponse_AddNewMessage_Flat)ToPB() *PB_MsgResponse_AddNewMessage {
r := &PB_MsgResponse_AddNewMessage{
}
return r
}

func(m *PB_MsgParam_SetRoomActionDoing_Flat)ToPB() *PB_MsgParam_SetRoomActionDoing {
r := &PB_MsgParam_SetRoomActionDoing{
    GroupId:  int64(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,
    
}
return r
}

func(m *PB_MsgResponse_SetRoomActionDoing_Flat)ToPB() *PB_MsgResponse_SetRoomActionDoing {
r := &PB_MsgResponse_SetRoomActionDoing{
}
return r
}

func(m *PB_MsgParam_GetMessagesByIds_Flat)ToPB() *PB_MsgParam_GetMessagesByIds {
r := &PB_MsgParam_GetMessagesByIds{
    
}
return r
}

func(m *PB_MsgResponse_GetMessagesByIds_Flat)ToPB() *PB_MsgResponse_GetMessagesByIds {
r := &PB_MsgResponse_GetMessagesByIds{
    
}
return r
}

func(m *PB_MsgParam_GetMessagesHistory_Flat)ToPB() *PB_MsgParam_GetMessagesHistory {
r := &PB_MsgParam_GetMessagesHistory{
    ChatKey:  m.ChatKey ,
    FromSeq:  int32(m.FromSeq) ,
    EndSeq:  int32(m.EndSeq) ,
}
return r
}

func(m *PB_MsgResponse_GetMessagesHistory_Flat)ToPB() *PB_MsgResponse_GetMessagesHistory {
r := &PB_MsgResponse_GetMessagesHistory{
    
}
return r
}

func(m *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_MsgParam_SetChatMessagesRangeAsSeen {
r := &PB_MsgParam_SetChatMessagesRangeAsSeen{
    ChatKey:  m.ChatKey ,
    FromSeq:  int32(m.FromSeq) ,
    EndSeq:  int32(m.EndSeq) ,
    SeenTimeMs:  int64(m.SeenTimeMs) ,
}
return r
}

func(m *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_MsgResponse_SetChatMessagesRangeAsSeen {
r := &PB_MsgResponse_SetChatMessagesRangeAsSeen{
}
return r
}

func(m *PB_MsgParam_DeleteChatHistory_Flat)ToPB() *PB_MsgParam_DeleteChatHistory {
r := &PB_MsgParam_DeleteChatHistory{
    ChatKey:  m.ChatKey ,
    ToSeq:  int32(m.ToSeq) ,
}
return r
}

func(m *PB_MsgResponse_DeleteChatHistory_Flat)ToPB() *PB_MsgResponse_DeleteChatHistory {
r := &PB_MsgResponse_DeleteChatHistory{
}
return r
}

func(m *PB_MsgParam_DeleteMessagesByIds_Flat)ToPB() *PB_MsgParam_DeleteMessagesByIds {
r := &PB_MsgParam_DeleteMessagesByIds{
    
}
return r
}

func(m *PB_MsgResponse_DeleteMessagesByIds_Flat)ToPB() *PB_MsgResponse_DeleteMessagesByIds {
r := &PB_MsgResponse_DeleteMessagesByIds{
}
return r
}

func(m *PB_MsgParam_SetMessagesAsReceived_Flat)ToPB() *PB_MsgParam_SetMessagesAsReceived {
r := &PB_MsgParam_SetMessagesAsReceived{
    
}
return r
}

func(m *PB_MsgResponse_SetMessagesAsReceived_Flat)ToPB() *PB_MsgResponse_SetMessagesAsReceived {
r := &PB_MsgResponse_SetMessagesAsReceived{
}
return r
}

func(m *PB_MsgParam_ForwardMessages_Flat)ToPB() *PB_MsgParam_ForwardMessages {
r := &PB_MsgParam_ForwardMessages{
    
    ToDirectChatKeys:  m.ToDirectChatKeys ,
    ToGroupChatIds:  helper.SliceIntToInt64(m.ToGroupChatIds) ,
}
return r
}

func(m *PB_MsgResponse_ForwardMessages_Flat)ToPB() *PB_MsgResponse_ForwardMessages {
r := &PB_MsgResponse_ForwardMessages{
}
return r
}

func(m *PB_MsgParam_EditMessage_Flat)ToPB() *PB_MsgParam_EditMessage {
r := &PB_MsgParam_EditMessage{
    ChatKey:  m.ChatKey ,
    
    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_MsgResponse_EditMessage_Flat)ToPB() *PB_MsgResponse_EditMessage {
r := &PB_MsgResponse_EditMessage{
}
return r
}

func(m *PB_MsgParam_BroadcastNewMessage_Flat)ToPB() *PB_MsgParam_BroadcastNewMessage {
r := &PB_MsgParam_BroadcastNewMessage{
    Text:  m.Text ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,
    
}
return r
}

func(m *PB_MsgResponse_BroadcastNewMessage_Flat)ToPB() *PB_MsgResponse_BroadcastNewMessage {
r := &PB_MsgResponse_BroadcastNewMessage{
}
return r
}

func(m *PB_MsgParam_GetFreshChatList_Flat)ToPB() *PB_MsgParam_GetFreshChatList {
r := &PB_MsgParam_GetFreshChatList{
}
return r
}

func(m *PB_MsgResponse_GetFreshChatList_Flat)ToPB() *PB_MsgResponse_GetFreshChatList {
r := &PB_MsgResponse_GetFreshChatList{
    
    
}
return r
}

func(m *PB_MsgParam_GetFreshRoomMessagesList_Flat)ToPB() *PB_MsgParam_GetFreshRoomMessagesList {
r := &PB_MsgParam_GetFreshRoomMessagesList{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    Last:  int64(m.Last) ,
}
return r
}

func(m *PB_MsgResponse_GetFreshRoomMessagesList_Flat)ToPB() *PB_MsgResponse_GetFreshRoomMessagesList {
r := &PB_MsgResponse_GetFreshRoomMessagesList{
    
    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_MsgParam_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_MsgParam_GetFreshAllDirectMessagesList {
r := &PB_MsgParam_GetFreshAllDirectMessagesList{
    LowMessageId:  int64(m.LowMessageId) ,
    HighMessageId:  int64(m.HighMessageId) ,
}
return r
}

func(m *PB_MsgResponse_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_MsgResponse_GetFreshAllDirectMessagesList {
r := &PB_MsgResponse_GetFreshAllDirectMessagesList{
    
    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_MessageForwardedFrom_Flat)ToPB() *PB_MessageForwardedFrom {
r := &PB_MessageForwardedFrom{
    RoomId:  int64(m.RoomId) ,
    
    MessageId:  int64(m.MessageId) ,
    MessageSeq:  int32(m.MessageSeq) ,
}
return r
}

func(m *PB_MessagesCollections_Flat)ToPB() *PB_MessagesCollections {
r := &PB_MessagesCollections{
    DirectMessagesIds:  helper.SliceIntToInt64(m.DirectMessagesIds) ,
    GroupMessagesIds:  helper.SliceIntToInt64(m.GroupMessagesIds) ,
    BroadCatMessagesIds:  helper.SliceIntToInt64(m.BroadCatMessagesIds) ,
}
return r
}

func(m *PB_MsgParam_Echo_Flat)ToPB() *PB_MsgParam_Echo {
r := &PB_MsgParam_Echo{
    Text:  m.Text ,
}
return r
}

func(m *PB_MsgResponse_PB_MsgParam_Echo_Flat)ToPB() *PB_MsgResponse_PB_MsgParam_Echo {
r := &PB_MsgResponse_PB_MsgParam_Echo{
    Text:  m.Text ,
}
return r
}

func(m *PB_SyncParam_GetDirectUpdates_Flat)ToPB() *PB_SyncParam_GetDirectUpdates {
r := &PB_SyncParam_GetDirectUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetDirectUpdates_Flat)ToPB() *PB_SyncResponse_GetDirectUpdates {
r := &PB_SyncResponse_GetDirectUpdates{
    
    
    
    
    
    
    
    
    
    
    
    
    
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_GetGeneralUpdates_Flat)ToPB() *PB_SyncParam_GetGeneralUpdates {
r := &PB_SyncParam_GetGeneralUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetGeneralUpdates_Flat)ToPB() *PB_SyncResponse_GetGeneralUpdates {
r := &PB_SyncResponse_GetGeneralUpdates{
    
    
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_GetNotifyUpdates_Flat)ToPB() *PB_SyncParam_GetNotifyUpdates {
r := &PB_SyncParam_GetNotifyUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetNotifyUpdates_Flat)ToPB() *PB_SyncResponse_GetNotifyUpdates {
r := &PB_SyncResponse_GetNotifyUpdates{
    
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_SetLastSyncDirectUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncDirectUpdateId {
r := &PB_SyncParam_SetLastSyncDirectUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncDirectUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncDirectUpdateId {
r := &PB_SyncResponse_SetLastSyncDirectUpdateId{
}
return r
}

func(m *PB_SyncParam_SetLastSyncGeneralUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncGeneralUpdateId {
r := &PB_SyncParam_SetLastSyncGeneralUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncGeneralUpdateId {
r := &PB_SyncResponse_SetLastSyncGeneralUpdateId{
}
return r
}

func(m *PB_SyncParam_SetLastSyncNotifyUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncNotifyUpdateId {
r := &PB_SyncParam_SetLastSyncNotifyUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncNotifyUpdateId {
r := &PB_SyncResponse_SetLastSyncNotifyUpdateId{
}
return r
}

func(m *PB_NotifyUpdatesView_Flat)ToPB() *PB_NotifyUpdatesView {
r := &PB_NotifyUpdatesView{
    
    
}
return r
}

func(m *PB_AllLivePushes_Flat)ToPB() *PB_AllLivePushes {
r := &PB_AllLivePushes{
    
    
    
}
return r
}

func(m *PB_UserParam_BlockUser_Flat)ToPB() *PB_UserParam_BlockUser {
r := &PB_UserParam_BlockUser{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserOfflineResponse_BlockUser_Flat)ToPB() *PB_UserOfflineResponse_BlockUser {
r := &PB_UserOfflineResponse_BlockUser{
    ByUserId:  int32(m.ByUserId) ,
    TargetUserId:  int32(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UnBlockUser_Flat)ToPB() *PB_UserParam_UnBlockUser {
r := &PB_UserParam_UnBlockUser{
    Offset:  int32(m.Offset) ,
    Limit:  int32(m.Limit) ,
}
return r
}

func(m *PB_UserOfflineResponse_UnBlockUser_Flat)ToPB() *PB_UserOfflineResponse_UnBlockUser {
r := &PB_UserOfflineResponse_UnBlockUser{
    
}
return r
}

func(m *PB_UserParam_BlockedList_Flat)ToPB() *PB_UserParam_BlockedList {
r := &PB_UserParam_BlockedList{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserResponse_BlockedList_Flat)ToPB() *PB_UserResponse_BlockedList {
r := &PB_UserResponse_BlockedList{
    ByUserId:  int32(m.ByUserId) ,
    TargetUserId:  int32(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UpdateAbout_Flat)ToPB() *PB_UserParam_UpdateAbout {
r := &PB_UserParam_UpdateAbout{
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserOfflineResponse_UpdateAbout_Flat)ToPB() *PB_UserOfflineResponse_UpdateAbout {
r := &PB_UserOfflineResponse_UpdateAbout{
    UserId:  int32(m.UserId) ,
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserParam_UpdateUserName_Flat)ToPB() *PB_UserParam_UpdateUserName {
r := &PB_UserParam_UpdateUserName{
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserOfflineResponse_UpdateUserName_Flat)ToPB() *PB_UserOfflineResponse_UpdateUserName {
r := &PB_UserOfflineResponse_UpdateUserName{
    UserId:  int32(m.UserId) ,
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserParam_ChangeAvatar_Flat)ToPB() *PB_UserParam_ChangeAvatar {
r := &PB_UserParam_ChangeAvatar{
    None:  m.None ,
    ImageData2:  m.ImageData2 ,
}
return r
}

func(m *PB_UserOfflineResponse_ChangeAvatar_Flat)ToPB() *PB_UserOfflineResponse_ChangeAvatar {
r := &PB_UserOfflineResponse_ChangeAvatar{
}
return r
}

func(m *PB_UserParam_ChangePrivacy_Flat)ToPB() *PB_UserParam_ChangePrivacy {
r := &PB_UserParam_ChangePrivacy{
    
}
return r
}

func(m *PB_UserResponseOffline_ChangePrivacy_Flat)ToPB() *PB_UserResponseOffline_ChangePrivacy {
r := &PB_UserResponseOffline_ChangePrivacy{
}
return r
}

func(m *PB_UserParam_CheckUserName_Flat)ToPB() *PB_UserParam_CheckUserName {
r := &PB_UserParam_CheckUserName{
    
}
return r
}

func(m *PB_UserResponse_CheckUserName_Flat)ToPB() *PB_UserResponse_CheckUserName {
r := &PB_UserResponse_CheckUserName{
}
return r
}

func(m *UserView_Flat)ToPB() *UserView {
r := &UserView{
}
return r
}

func(m *PB_Activity_Flat)ToPB() *PB_Activity {
r := &PB_Activity{
    Id:  int64(m.Id) ,
    ActorUserId:  int32(m.ActorUserId) ,
    ActionTypeId:  int32(m.ActionTypeId) ,
    RowId:  int32(m.RowId) ,
    RootId:  int32(m.RootId) ,
    RefId:  int64(m.RefId) ,
    CreatedAt:  int32(m.CreatedAt) ,
}
return r
}

func(m *PB_Bucket_Flat)ToPB() *PB_Bucket {
r := &PB_Bucket{
    BucketId:  int32(m.BucketId) ,
    BucketName:  m.BucketName ,
    Server1Id:  int32(m.Server1Id) ,
    Server2Id:  int32(m.Server2Id) ,
    BackupServerId:  int32(m.BackupServerId) ,
    ContentObjectTypeId:  int32(m.ContentObjectTypeId) ,
    ContentObjectCount:  int32(m.ContentObjectCount) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Chat_Flat)ToPB() *PB_Chat {
r := &PB_Chat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedSe:  int32(m.CreatedSe) ,
    StartMessageIdFrom:  int64(m.StartMessageIdFrom) ,
    LastDeletedMessageId:  int64(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int64(m.LastSeenMessageId) ,
    LastMessageId:  int64(m.LastMessageId) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
}
return r
}

func(m *PB_Comment_Flat)ToPB() *PB_Comment {
r := &PB_Comment{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    Text:  m.Text ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_DirectMessage_Flat)ToPB() *PB_DirectMessage {
r := &PB_DirectMessage{
    MessageId:  int64(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnumId:  int32(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int32(m.CreatedSe) ,
    PeerReceivedTime:  int32(m.PeerReceivedTime) ,
    PeerSeenTime:  int32(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int32(m.DeliviryStatusEnumId) ,
}
return r
}

func(m *PB_DirectOffline_Flat)ToPB() *PB_DirectOffline {
r := &PB_DirectOffline{
    DirectOfflineId:  int64(m.DirectOfflineId) ,
    ToUserId:  int32(m.ToUserId) ,
    MessageId:  int64(m.MessageId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    OtherId:  int64(m.OtherId) ,
    ChatKey:  m.ChatKey ,
    PeerUserId:  int32(m.PeerUserId) ,
    RoomLogTypeId:  int32(m.RoomLogTypeId) ,
    PBClass:  m.PBClass ,
    DataPB:  m.DataPB ,
    DataJson:  m.DataJson ,
    DataTemp:  m.DataTemp ,
    AtTimeMs:  int64(m.AtTimeMs) ,
}
return r
}

func(m *PB_DirectToMessage_Flat)ToPB() *PB_DirectToMessage {
r := &PB_DirectToMessage{
    Id:  int64(m.Id) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int64(m.MessageId) ,
    SourceEnumId:  int32(m.SourceEnumId) ,
}
return r
}

func(m *PB_DirectUpdate_Flat)ToPB() *PB_DirectUpdate {
r := &PB_DirectUpdate{
    DirectUpdateId:  int64(m.DirectUpdateId) ,
    ToUserId:  int32(m.ToUserId) ,
    MessageId:  int64(m.MessageId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    OtherId:  int64(m.OtherId) ,
    ChatKey:  m.ChatKey ,
    PeerUserId:  int32(m.PeerUserId) ,
    EventType:  int32(m.EventType) ,
    RoomLogTypeId:  int32(m.RoomLogTypeId) ,
    FromSeq:  int32(m.FromSeq) ,
    ToSeq:  int32(m.ToSeq) ,
    ExtraPB:  m.ExtraPB ,
    ExtraJson:  m.ExtraJson ,
    AtTimeMs:  int64(m.AtTimeMs) ,
}
return r
}

func(m *PB_FollowingList_Flat)ToPB() *PB_FollowingList {
r := &PB_FollowingList{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    ListType:  int32(m.ListType) ,
    Name:  m.Name ,
    Count:  int32(m.Count) ,
    IsAuto:  int32(m.IsAuto) ,
    IsPimiry:  int32(m.IsPimiry) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_FollowingListMember_Flat)ToPB() *PB_FollowingListMember {
r := &PB_FollowingListMember{
    Id:  int64(m.Id) ,
    ListId:  int32(m.ListId) ,
    UserId:  int32(m.UserId) ,
    FollowedUserId:  int32(m.FollowedUserId) ,
    FollowType:  int32(m.FollowType) ,
    UpdatedTimeMs:  int64(m.UpdatedTimeMs) ,
}
return r
}

func(m *PB_FollowingListMemberHistory_Flat)ToPB() *PB_FollowingListMemberHistory {
r := &PB_FollowingListMemberHistory{
    Id:  int64(m.Id) ,
    ListId:  int32(m.ListId) ,
    UserId:  int32(m.UserId) ,
    FollowedUserId:  int32(m.FollowedUserId) ,
    FollowType:  int32(m.FollowType) ,
    UpdatedTimeMs:  int64(m.UpdatedTimeMs) ,
    FollowId:  int32(m.FollowId) ,
}
return r
}

func(m *PB_GeneralLog_Flat)ToPB() *PB_GeneralLog {
r := &PB_GeneralLog{
    Id:  int64(m.Id) ,
    ToUserId:  int32(m.ToUserId) ,
    TargetId:  int32(m.TargetId) ,
    LogTypeId:  int32(m.LogTypeId) ,
    ExtraPB:  m.ExtraPB ,
    ExtraJson:  m.ExtraJson ,
    CreatedMs:  int64(m.CreatedMs) ,
}
return r
}

func(m *PB_Group_Flat)ToPB() *PB_Group {
r := &PB_Group{
    GroupId:  int64(m.GroupId) ,
    GroupName:  m.GroupName ,
    MembersCount:  int32(m.MembersCount) ,
    GroupPrivacyEnum:  int32(m.GroupPrivacyEnum) ,
    CreatorUserId:  int32(m.CreatorUserId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
}
return r
}

func(m *PB_GroupMember_Flat)ToPB() *PB_GroupMember {
r := &PB_GroupMember{
    Id:  int64(m.Id) ,
    GroupId:  int64(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    UserId:  int32(m.UserId) ,
    ByUserId:  int32(m.ByUserId) ,
    GroupRoleEnumId:  int32(m.GroupRoleEnumId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_GroupMessage_Flat)ToPB() *PB_GroupMessage {
r := &PB_GroupMessage{
    MessageId:  int64(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnum:  int32(m.MessageTypeEnum) ,
    Text:  m.Text ,
    CreatedMs:  int64(m.CreatedMs) ,
    DeliveryStatusEnum:  int32(m.DeliveryStatusEnum) ,
}
return r
}

func(m *PB_GroupToMessage_Flat)ToPB() *PB_GroupToMessage {
r := &PB_GroupToMessage{
    Id:  int64(m.Id) ,
    GroupId:  int64(m.GroupId) ,
    MessageId:  int64(m.MessageId) ,
    CreatedMs:  int64(m.CreatedMs) ,
    StatusEnum:  int32(m.StatusEnum) ,
}
return r
}

func(m *PB_Like_Flat)ToPB() *PB_Like {
r := &PB_Like{
    Id:  int32(m.Id) ,
    PostId:  int32(m.PostId) ,
    PostTypeId:  int32(m.PostTypeId) ,
    UserId:  int32(m.UserId) ,
    TypeId:  int32(m.TypeId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_LogChange_Flat)ToPB() *PB_LogChange {
r := &PB_LogChange{
    Id:  int32(m.Id) ,
    T:  m.T ,
}
return r
}

func(m *PB_Media_Flat)ToPB() *PB_Media {
r := &PB_Media{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    AlbumId:  int32(m.AlbumId) ,
    TypeId:  int32(m.TypeId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    Src:  m.Src ,
}
return r
}

func(m *PB_MessageFile_Flat)ToPB() *PB_MessageFile {
r := &PB_MessageFile{
    MessageFileId:  int64(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int32(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileTypeEnumId:  int32(m.FileTypeEnumId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int64(m.HashAccess) ,
    CreatedSe:  int32(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
}
return r
}

func(m *PB_Notification_Flat)ToPB() *PB_Notification {
r := &PB_Notification{
    Id:  int64(m.Id) ,
    ForUserId:  int32(m.ForUserId) ,
    ActorUserId:  int32(m.ActorUserId) ,
    ActionTypeId:  int32(m.ActionTypeId) ,
    ObjectTypeId:  int32(m.ObjectTypeId) ,
    RowId:  int32(m.RowId) ,
    RootId:  int32(m.RootId) ,
    RefId:  int32(m.RefId) ,
    SeenStatus:  int32(m.SeenStatus) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_NotificationRemoved_Flat)ToPB() *PB_NotificationRemoved {
r := &PB_NotificationRemoved{
    NotificationId:  int32(m.NotificationId) ,
    ForUserId:  int32(m.ForUserId) ,
}
return r
}

func(m *PB_Offline_Flat)ToPB() *PB_Offline {
r := &PB_Offline{
    Id:  int64(m.Id) ,
    FromUserId:  int32(m.FromUserId) ,
    ToUserId:  int32(m.ToUserId) ,
    RpcName:  m.RpcName ,
    PBClassName:  m.PBClassName ,
    Key:  m.Key ,
    DataJson:  m.DataJson ,
    DataBlob:  m.DataBlob ,
    DataLength:  int32(m.DataLength) ,
    CreatedMs:  int64(m.CreatedMs) ,
}
return r
}

func(m *PB_OldMessage_Flat)ToPB() *PB_OldMessage {
r := &PB_OldMessage{
    Id:  int64(m.Id) ,
    Uid:  int64(m.Uid) ,
    UserId:  int64(m.UserId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    MessageType:  int32(m.MessageType) ,
    RoomType:  int32(m.RoomType) ,
    MsgFileId:  int64(m.MsgFileId) ,
    DataPB:  m.DataPB ,
    Data64:  m.Data64 ,
    DataJson:  m.DataJson ,
    CreatedTimeMs:  int64(m.CreatedTimeMs) ,
}
return r
}

func(m *PB_OldMsgFile_Flat)ToPB() *PB_OldMsgFile {
r := &PB_OldMsgFile{
    Id:  int64(m.Id) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileType:  int32(m.FileType) ,
    MimeType:  m.MimeType ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    ThumbData:  m.ThumbData ,
    ThumbData64:  m.ThumbData64 ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
}
return r
}

func(m *PB_OldMsgPush_Flat)ToPB() *PB_OldMsgPush {
r := &PB_OldMsgPush{
    Id:  int64(m.Id) ,
    Uid:  int64(m.Uid) ,
    ToUser:  int64(m.ToUser) ,
    MsgUid:  int64(m.MsgUid) ,
    CreatedTimeMs:  int64(m.CreatedTimeMs) ,
}
return r
}

func(m *PB_OldMsgPushEvent_Flat)ToPB() *PB_OldMsgPushEvent {
r := &PB_OldMsgPushEvent{
    Id:  int64(m.Id) ,
    Uid:  int64(m.Uid) ,
    ToUserId:  int32(m.ToUserId) ,
    MsgUid:  int64(m.MsgUid) ,
    MsgKey:  m.MsgKey ,
    RoomKey:  m.RoomKey ,
    PeerUserId:  int32(m.PeerUserId) ,
    EventType:  int32(m.EventType) ,
    AtTime:  int32(m.AtTime) ,
}
return r
}

func(m *PB_PhoneContact_Flat)ToPB() *PB_PhoneContact {
r := &PB_PhoneContact{
    Id:  int32(m.Id) ,
    PhoneDisplayName:  m.PhoneDisplayName ,
    PhoneFamilyName:  m.PhoneFamilyName ,
    PhoneNumber:  m.PhoneNumber ,
    PhoneNormalizedNumber:  m.PhoneNormalizedNumber ,
    PhoneContactRowId:  int32(m.PhoneContactRowId) ,
    UserId:  int32(m.UserId) ,
    DeviceUuidId:  int32(m.DeviceUuidId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
}
return r
}

func(m *PB_Photo_Flat)ToPB() *PB_Photo {
r := &PB_Photo{
    PhotoId:  int32(m.PhotoId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    AlbumId:  int32(m.AlbumId) ,
    ImageTypeId:  int32(m.ImageTypeId) ,
    Title:  m.Title ,
    Src:  m.Src ,
    PathSrc:  m.PathSrc ,
    BucketId:  int32(m.BucketId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Ratio:  m.Ratio ,
    HashMd5:  m.HashMd5 ,
    Color:  m.Color ,
    CreatedTime:  int32(m.CreatedTime) ,
    W1080:  int32(m.W1080) ,
    W720:  int32(m.W720) ,
    W480:  int32(m.W480) ,
    W320:  int32(m.W320) ,
    W160:  int32(m.W160) ,
    W80:  int32(m.W80) ,
}
return r
}

func(m *PB_Post_Flat)ToPB() *PB_Post {
r := &PB_Post{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    TypeId:  int32(m.TypeId) ,
    Text:  m.Text ,
    FormatedText:  m.FormatedText ,
    MediaCount:  int32(m.MediaCount) ,
    SharedTo:  int32(m.SharedTo) ,
    DisableComment:  int32(m.DisableComment) ,
    HasTag:  int32(m.HasTag) ,
    LikesCount:  int32(m.LikesCount) ,
    CommentsCount:  int32(m.CommentsCount) ,
    EditedTime:  int32(m.EditedTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_PushEvent_Flat)ToPB() *PB_PushEvent {
r := &PB_PushEvent{
    PushEventId:  int64(m.PushEventId) ,
    ToUserId:  int32(m.ToUserId) ,
    ToDeviceId:  int64(m.ToDeviceId) ,
    MessageId:  int64(m.MessageId) ,
    RoomTypeEnum:  int32(m.RoomTypeEnum) ,
    RoomId:  int64(m.RoomId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    PushEventTypeEnum:  int32(m.PushEventTypeEnum) ,
    AtTime:  int32(m.AtTime) ,
}
return r
}

func(m *PB_PushMessage_Flat)ToPB() *PB_PushMessage {
r := &PB_PushMessage{
    PushMessageId:  int64(m.PushMessageId) ,
    ToUserId:  int32(m.ToUserId) ,
    ToDeviceId:  int64(m.ToDeviceId) ,
    MessageId:  int64(m.MessageId) ,
    RoomTypeEnum:  int32(m.RoomTypeEnum) ,
    CreatedMs:  int64(m.CreatedMs) ,
}
return r
}

func(m *PB_RecommendUser_Flat)ToPB() *PB_RecommendUser {
r := &PB_RecommendUser{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    TargetId:  int32(m.TargetId) ,
    Weight:  m.Weight ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Room_Flat)ToPB() *PB_Room {
r := &PB_Room{
    RoomId:  int64(m.RoomId) ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnum:  int32(m.RoomTypeEnum) ,
    UserId:  int32(m.UserId) ,
    LastSeqSeen:  int32(m.LastSeqSeen) ,
    LastSeqDelete:  int32(m.LastSeqDelete) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
}
return r
}

func(m *PB_SearchClicked_Flat)ToPB() *PB_SearchClicked {
r := &PB_SearchClicked{
    Id:  int64(m.Id) ,
    Query:  m.Query ,
    ClickType:  int32(m.ClickType) ,
    TargetId:  int32(m.TargetId) ,
    UserId:  int32(m.UserId) ,
    CreatedAt:  int32(m.CreatedAt) ,
}
return r
}

func(m *PB_Session_Flat)ToPB() *PB_Session {
r := &PB_Session{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    SessionUuid:  m.SessionUuid ,
    ClientUuid:  m.ClientUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    LastIpAddress:  m.LastIpAddress ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    LastNetworkTypeId:  int32(m.LastNetworkTypeId) ,
    AppVersion:  int32(m.AppVersion) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    LastSyncDirectUpdateId:  int64(m.LastSyncDirectUpdateId) ,
    LastSyncGeneralUpdateId:  int64(m.LastSyncGeneralUpdateId) ,
    LastSyncNotifyUpdateId:  int64(m.LastSyncNotifyUpdateId) ,
}
return r
}

func(m *PB_SettingClient_Flat)ToPB() *PB_SettingClient {
r := &PB_SettingClient{
    UserId:  int32(m.UserId) ,
    AutoDownloadWifiVoice:  int32(m.AutoDownloadWifiVoice) ,
    AutoDownloadWifiImage:  int32(m.AutoDownloadWifiImage) ,
    AutoDownloadWifiVideo:  int32(m.AutoDownloadWifiVideo) ,
    AutoDownloadWifiFile:  int32(m.AutoDownloadWifiFile) ,
    AutoDownloadWifiMusic:  int32(m.AutoDownloadWifiMusic) ,
    AutoDownloadWifiGif:  int32(m.AutoDownloadWifiGif) ,
    AutoDownloadCellularVoice:  int32(m.AutoDownloadCellularVoice) ,
    AutoDownloadCellularImage:  int32(m.AutoDownloadCellularImage) ,
    AutoDownloadCellularVideo:  int32(m.AutoDownloadCellularVideo) ,
    AutoDownloadCellularFile:  int32(m.AutoDownloadCellularFile) ,
    AutoDownloadCellularMusic:  int32(m.AutoDownloadCellularMusic) ,
    AutoDownloadCellularGif:  int32(m.AutoDownloadCellularGif) ,
    AutoDownloadRoamingVoice:  int32(m.AutoDownloadRoamingVoice) ,
    AutoDownloadRoamingImage:  int32(m.AutoDownloadRoamingImage) ,
    AutoDownloadRoamingVideo:  int32(m.AutoDownloadRoamingVideo) ,
    AutoDownloadRoamingFile:  int32(m.AutoDownloadRoamingFile) ,
    AutoDownloadRoamingMusic:  int32(m.AutoDownloadRoamingMusic) ,
    AutoDownloadRoamingGif:  int32(m.AutoDownloadRoamingGif) ,
    SaveToGallery:  int32(m.SaveToGallery) ,
}
return r
}

func(m *PB_SettingNotification_Flat)ToPB() *PB_SettingNotification {
r := &PB_SettingNotification{
    UserId:  int32(m.UserId) ,
    SocialLedOn:  int32(m.SocialLedOn) ,
    SocialLedColor:  m.SocialLedColor ,
    ReqestToFollowYou:  int32(m.ReqestToFollowYou) ,
    FollowedYou:  int32(m.FollowedYou) ,
    AccptedYourFollowRequest:  int32(m.AccptedYourFollowRequest) ,
    YourPostLiked:  int32(m.YourPostLiked) ,
    YourPostCommented:  int32(m.YourPostCommented) ,
    MenthenedYouInPost:  int32(m.MenthenedYouInPost) ,
    MenthenedYouInComment:  int32(m.MenthenedYouInComment) ,
    YourContactsJoined:  int32(m.YourContactsJoined) ,
    DirectMessage:  int32(m.DirectMessage) ,
    DirectAlert:  int32(m.DirectAlert) ,
    DirectPerview:  int32(m.DirectPerview) ,
    DirectLedOn:  int32(m.DirectLedOn) ,
    DirectLedColor:  int32(m.DirectLedColor) ,
    DirectVibrate:  int32(m.DirectVibrate) ,
    DirectPopup:  int32(m.DirectPopup) ,
    DirectSound:  int32(m.DirectSound) ,
    DirectPriority:  int32(m.DirectPriority) ,
}
return r
}

func(m *PB_Tag_Flat)ToPB() *PB_Tag {
r := &PB_Tag{
    Id:  int32(m.Id) ,
    Name:  m.Name ,
    Count:  int32(m.Count) ,
    IsBlocked:  int32(m.IsBlocked) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_TagsPost_Flat)ToPB() *PB_TagsPost {
r := &PB_TagsPost{
    Id:  int32(m.Id) ,
    TagId:  int32(m.TagId) ,
    PostId:  int32(m.PostId) ,
    TypeId:  int32(m.TypeId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_TestChat_Flat)ToPB() *PB_TestChat {
r := &PB_TestChat{
    Id:  int64(m.Id) ,
    Id4:  int64(m.Id4) ,
    TimeMs:  int64(m.TimeMs) ,
    Text:  m.Text ,
    Name:  m.Name ,
    UserId:  int64(m.UserId) ,
    C2:  int32(m.C2) ,
    C3:  int32(m.C3) ,
    C4:  int32(m.C4) ,
    C5:  int32(m.C5) ,
}
return r
}

func(m *PB_TriggerLog_Flat)ToPB() *PB_TriggerLog {
r := &PB_TriggerLog{
    Id:  int64(m.Id) ,
    ModelName:  m.ModelName ,
    ChangeType:  m.ChangeType ,
    TargetId:  int64(m.TargetId) ,
    TargetStr:  m.TargetStr ,
    CreatedSe:  int32(m.CreatedSe) ,
}
return r
}

func(m *PB_User_Flat)ToPB() *PB_User {
r := &PB_User{
    Id:  int32(m.Id) ,
    UserName:  m.UserName ,
    UserNameLower:  m.UserNameLower ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int32(m.PrivacyProfile) ,
    Phone:  m.Phone ,
    Email:  m.Email ,
    IsDeleted:  int32(m.IsDeleted) ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    MediaCount:  int32(m.MediaCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ResharedCount:  int32(m.ResharedCount) ,
    LastActionTime:  int32(m.LastActionTime) ,
    LastPostTime:  int32(m.LastPostTime) ,
    PrimaryFollowingList:  int32(m.PrimaryFollowingList) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    SessionUuid:  m.SessionUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    AppVersion:  int32(m.AppVersion) ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    LastLoginTime:  int32(m.LastLoginTime) ,
    LastIpAddress:  m.LastIpAddress ,
}
return r
}

func(m *PB_UserMetaInfo_Flat)ToPB() *PB_UserMetaInfo {
r := &PB_UserMetaInfo{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    IsNotificationDirty:  int32(m.IsNotificationDirty) ,
    LastUserRecGen:  int32(m.LastUserRecGen) ,
}
return r
}

func(m *PB_UserPassword_Flat)ToPB() *PB_UserPassword {
r := &PB_UserPassword{
    UserId:  int32(m.UserId) ,
    Password:  m.Password ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_UpdateGroupParticipants_Flat)ToPB() *PB_UpdateGroupParticipants {
r := &PB_UpdateGroupParticipants{
}
return r
}

func(m *PB_UpdateNotifySettings_Flat)ToPB() *PB_UpdateNotifySettings {
r := &PB_UpdateNotifySettings{
}
return r
}

func(m *PB_UpdateServiceNotification_Flat)ToPB() *PB_UpdateServiceNotification {
r := &PB_UpdateServiceNotification{
}
return r
}

func(m *PB_UpdateMessageMeta_Flat)ToPB() *PB_UpdateMessageMeta {
r := &PB_UpdateMessageMeta{
    MessageId:  int64(m.MessageId) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_UpdateMessageId_Flat)ToPB() *PB_UpdateMessageId {
r := &PB_UpdateMessageId{
    OldMessageId:  int64(m.OldMessageId) ,
    NewMessageId:  int64(m.NewMessageId) ,
}
return r
}

func(m *PB_UpdateMessageToEdit_Flat)ToPB() *PB_UpdateMessageToEdit {
r := &PB_UpdateMessageToEdit{
    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_UpdateMessageToDelete_Flat)ToPB() *PB_UpdateMessageToDelete {
r := &PB_UpdateMessageToDelete{
    MessageId:  int64(m.MessageId) ,
}
return r
}

func(m *PB_UpdateRoomActionDoing_Flat)ToPB() *PB_UpdateRoomActionDoing {
r := &PB_UpdateRoomActionDoing{
    RoomKey:  m.RoomKey ,
    
}
return r
}

func(m *PB_UpdateUserBlocked_Flat)ToPB() *PB_UpdateUserBlocked {
r := &PB_UpdateUserBlocked{
    UserId:  int32(m.UserId) ,
    Blocked:  m.Blocked ,
}
return r
}

func(m *PB_ChatView_Flat)ToPB() *PB_ChatView {
r := &PB_ChatView{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedSe:  int32(m.CreatedSe) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    LastMessageId:  int64(m.LastMessageId) ,
    LastDeletedMessageId:  int64(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int64(m.LastSeenMessageId) ,
    LastSeqSeen:  int32(m.LastSeqSeen) ,
    LastSeqDelete:  int32(m.LastSeqDelete) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
    
    SharedMediaCount:  int32(m.SharedMediaCount) ,
    UnseenCount:  int32(m.UnseenCount) ,
    
    
}
return r
}

func(m *PB_MessageView_Flat)ToPB() *PB_MessageView {
r := &PB_MessageView{
    MessageId:  int64(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnumId:  int32(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int32(m.CreatedSe) ,
    PeerReceivedTime:  int32(m.PeerReceivedTime) ,
    PeerSeenTime:  int32(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int32(m.DeliviryStatusEnumId) ,
    ChatKey:  m.ChatKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    IsByMe:  m.IsByMe ,
    RemoteId:  int64(m.RemoteId) ,
    
}
return r
}

func(m *PB_MessageFileView_Flat)ToPB() *PB_MessageFileView {
r := &PB_MessageFileView{
    MessageFileId:  int64(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int32(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileTypeEnumId:  int32(m.FileTypeEnumId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int64(m.HashAccess) ,
    CreatedSe:  int32(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
    ServerThumbLocalSrc:  m.ServerThumbLocalSrc ,
    RemoteMessageFileId:  int64(m.RemoteMessageFileId) ,
    LocalSrc:  m.LocalSrc ,
    ThumbLocalSrc:  m.ThumbLocalSrc ,
    MessageFileStatusId:  m.MessageFileStatusId ,
}
return r
}

func(m *PB_UserView_Flat)ToPB() *PB_UserView {
r := &PB_UserView{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int32(m.PrivacyProfile) ,
    IsDeleted:  int32(m.IsDeleted) ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    AppVersion:  int32(m.AppVersion) ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    FollowingType:  int32(m.FollowingType) ,
}
return r
}



///// folding ///

var GeoLocation__FOlD = &GeoLocation{
        Lat:  0.0 ,
        Lon:  0.0 ,
}


var RoomMessageLog__FOlD = &RoomMessageLog{
        
        TargetUserId:  0 ,
        ByUserId:  0 ,
}


var RoomMessageForwardFrom__FOlD = &RoomMessageForwardFrom{
        RoomId:  0 ,
        MessageId:  0 ,
        RoomTypeEnum:  0 ,
}


var RoomDraft__FOlD = &RoomDraft{
        Message:  "" ,
        ReplyTo:  0 ,
}


var ChatRoom__FOlD = &ChatRoom{
}


var Pagination__FOlD = &Pagination{
        Offset:  0 ,
        Limit:  0 ,
}


var PB_CommandToServer__FOlD = &PB_CommandToServer{
        ClientCallId:  0 ,
        Command:  "" ,
        RespondReached:  false ,
        Data:  []byte{} ,
}


var PB_CommandToClient__FOlD = &PB_CommandToClient{
        ServerCallId:  0 ,
        Command:  "" ,
        RespondReached:  false ,
        Data:  []byte{} ,
}


var PB_CommandReachedToServer__FOlD = &PB_CommandReachedToServer{
        ClientCallId:  0 ,
}


var PB_CommandReachedToClient__FOlD = &PB_CommandReachedToClient{
        ServerCallId:  0 ,
}


var PB_ResponseToClient__FOlD = &PB_ResponseToClient{
        ClientCallId:  0 ,
        PBClass:  "" ,
        RpcFullName:  "" ,
        Data:  []byte{} ,
}


var PB_Offline_NewDirectMessage__FOlD = &PB_Offline_NewDirectMessage{
        ChatKey:  "" ,
        FromMessageId:  0 ,
        AtTime:  0 ,
}


var PB_Offline_MessagesReachedServer__FOlD = &PB_Offline_MessagesReachedServer{
        MessageKeys:  "" ,
        AtTime:  0 ,
}


var PB_Offline_MessagesDeliveredToUser__FOlD = &PB_Offline_MessagesDeliveredToUser{
        MessageKeys:  "" ,
        RoomKey:  "" ,
        AtTime:  0 ,
}


var PB_Offline_MessagesSeenByPeer__FOlD = &PB_Offline_MessagesSeenByPeer{
        MessageKeys:  "" ,
        RoomKey:  "" ,
        AtTime:  0 ,
}


var PB_Offline_MessagesDeletedFromServer__FOlD = &PB_Offline_MessagesDeletedFromServer{
        MessageKeys:  "" ,
        AtTime:  0 ,
}


var PB_Offline_ChangeMessageId__FOlD = &PB_Offline_ChangeMessageId{
        MessageKey:  0 ,
        NewMessageId:  0 ,
}


var PB_Offline_ChangeMessageFileId__FOlD = &PB_Offline_ChangeMessageFileId{
        MessageFileKey:  0 ,
        NewMessageFileId:  0 ,
}


var PB_Offline_MessageToEdit__FOlD = &PB_Offline_MessageToEdit{
        MessageKey:  0 ,
        NewText:  "" ,
}


var PB_Offline_MessageToDelete__FOlD = &PB_Offline_MessageToDelete{
        MessageKey:  0 ,
}


var PB_Online_RoomActionDoing__FOlD = &PB_Online_RoomActionDoing{
        RoomKey:  "" ,
        
}


var PB_Offline_Messagings__FOlD = &PB_Offline_Messagings{
        
        
        
        
        
        
        
        
        
        
        
        
        
        LastId:  0 ,
}


var PB_UserParam_CheckUserName2__FOlD = &PB_UserParam_CheckUserName2{
}


var PB_UserResponse_CheckUserName2__FOlD = &PB_UserResponse_CheckUserName2{
}


var PB_ChatParam_AddNewMessage__FOlD = &PB_ChatParam_AddNewMessage{
        ReplyToMessageId:  0 ,
        Blob:  []byte{} ,
        
}


var PB_ChatResponse_AddNewMessage__FOlD = &PB_ChatResponse_AddNewMessage{
}


var PB_ChatParam_SetRoomActionDoing__FOlD = &PB_ChatParam_SetRoomActionDoing{
        GroupId:  0 ,
        DirectRoomKey:  "" ,
        
}


var PB_ChatResponse_SetRoomActionDoing__FOlD = &PB_ChatResponse_SetRoomActionDoing{
}


var PB_ChatParam_SetChatMessagesRangeAsSeen__FOlD = &PB_ChatParam_SetChatMessagesRangeAsSeen{
        ChatKey:  "" ,
        BottomMessageId:  0 ,
        TopMessageId:  0 ,
        SeenTimeMs:  0 ,
}


var PB_ChatResponse_SetChatMessagesRangeAsSeen__FOlD = &PB_ChatResponse_SetChatMessagesRangeAsSeen{
}


var PB_ChatParam_DeleteChatHistory__FOlD = &PB_ChatParam_DeleteChatHistory{
        ChatKey:  "" ,
        FromMessageId:  0 ,
}


var PB_ChatResponse_DeleteChatHistory__FOlD = &PB_ChatResponse_DeleteChatHistory{
}


var PB_ChatParam_DeleteMessagesByIds__FOlD = &PB_ChatParam_DeleteMessagesByIds{
        ChatKey:  "" ,
        Both:  false ,
        MessagesIds:  0 ,
}


var PB_ChatResponse_DeleteMessagesByIds__FOlD = &PB_ChatResponse_DeleteMessagesByIds{
}


var PB_ChatParam_SetMessagesAsReceived__FOlD = &PB_ChatParam_SetMessagesAsReceived{
        ChatRoom:  "" ,
        MessageIds:  0 ,
}


var PB_ChatResponse_SetMessagesAsReceived__FOlD = &PB_ChatResponse_SetMessagesAsReceived{
}


var PB_ChatParam_EditMessage__FOlD = &PB_ChatParam_EditMessage{
        RoomKey:  "" ,
        MessageId:  0 ,
        NewText:  "" ,
}


var PB_ChatResponse_EditMessage__FOlD = &PB_ChatResponse_EditMessage{
}


var PB_ChatParam_GetChatList__FOlD = &PB_ChatParam_GetChatList{
}


var PB_ChatResponse_GetChatList__FOlD = &PB_ChatResponse_GetChatList{
        
        
}


var PB_ChatParam_GetChatHistoryToOlder__FOlD = &PB_ChatParam_GetChatHistoryToOlder{
        ChatKey:  "" ,
        Limit:  0 ,
        FromTopMessageId:  0 ,
}


var PB_ChatResponse_GetChatHistoryToOlder__FOlD = &PB_ChatResponse_GetChatHistoryToOlder{
        
        HasMore:  false ,
}


var PB_ChatParam_GetFreshAllDirectMessagesList__FOlD = &PB_ChatParam_GetFreshAllDirectMessagesList{
        LowMessageId:  0 ,
        HighMessageId:  0 ,
}


var PB_ChatResponse_GetFreshAllDirectMessagesList__FOlD = &PB_ChatResponse_GetFreshAllDirectMessagesList{
        
        HasMore:  false ,
}


var PB_MsgParam_AddNewTextMessage__FOlD = &PB_MsgParam_AddNewTextMessage{
        Text:  "" ,
        MessageKey:  "" ,
        ToRoomKey:  "" ,
        PeerId:  0 ,
        Time:  0 ,
        ReplyToMessageId:  0 ,
        
}


var PB_MsgResponse_AddNewTextMessage__FOlD = &PB_MsgResponse_AddNewTextMessage{
}


var PB_MsgParam_AddNewMessage__FOlD = &PB_MsgParam_AddNewMessage{
        Text:  "" ,
        MessageKey:  "" ,
        ToRoomKey:  "" ,
        PeerId:  0 ,
        Time:  0 ,
        ReplyToMessageId:  0 ,
        
        
        Blob:  []byte{} ,
        
}


var PB_MsgResponse_AddNewMessage__FOlD = &PB_MsgResponse_AddNewMessage{
}


var PB_MsgParam_SetRoomActionDoing__FOlD = &PB_MsgParam_SetRoomActionDoing{
        GroupId:  0 ,
        DirectRoomKey:  "" ,
        
}


var PB_MsgResponse_SetRoomActionDoing__FOlD = &PB_MsgResponse_SetRoomActionDoing{
}


var PB_MsgParam_GetMessagesByIds__FOlD = &PB_MsgParam_GetMessagesByIds{
        
}


var PB_MsgResponse_GetMessagesByIds__FOlD = &PB_MsgResponse_GetMessagesByIds{
        
}


var PB_MsgParam_GetMessagesHistory__FOlD = &PB_MsgParam_GetMessagesHistory{
        ChatKey:  "" ,
        FromSeq:  0 ,
        EndSeq:  0 ,
}


var PB_MsgResponse_GetMessagesHistory__FOlD = &PB_MsgResponse_GetMessagesHistory{
        
}


var PB_MsgParam_SetChatMessagesRangeAsSeen__FOlD = &PB_MsgParam_SetChatMessagesRangeAsSeen{
        ChatKey:  "" ,
        FromSeq:  0 ,
        EndSeq:  0 ,
        SeenTimeMs:  0 ,
}


var PB_MsgResponse_SetChatMessagesRangeAsSeen__FOlD = &PB_MsgResponse_SetChatMessagesRangeAsSeen{
}


var PB_MsgParam_DeleteChatHistory__FOlD = &PB_MsgParam_DeleteChatHistory{
        ChatKey:  "" ,
        ToSeq:  0 ,
}


var PB_MsgResponse_DeleteChatHistory__FOlD = &PB_MsgResponse_DeleteChatHistory{
}


var PB_MsgParam_DeleteMessagesByIds__FOlD = &PB_MsgParam_DeleteMessagesByIds{
        
}


var PB_MsgResponse_DeleteMessagesByIds__FOlD = &PB_MsgResponse_DeleteMessagesByIds{
}


var PB_MsgParam_SetMessagesAsReceived__FOlD = &PB_MsgParam_SetMessagesAsReceived{
        
}


var PB_MsgResponse_SetMessagesAsReceived__FOlD = &PB_MsgResponse_SetMessagesAsReceived{
}


var PB_MsgParam_ForwardMessages__FOlD = &PB_MsgParam_ForwardMessages{
        
        ToDirectChatKeys:  "" ,
        ToGroupChatIds:  0 ,
}


var PB_MsgResponse_ForwardMessages__FOlD = &PB_MsgResponse_ForwardMessages{
}


var PB_MsgParam_EditMessage__FOlD = &PB_MsgParam_EditMessage{
        ChatKey:  "" ,
        
        MessageId:  0 ,
        NewText:  "" ,
}


var PB_MsgResponse_EditMessage__FOlD = &PB_MsgResponse_EditMessage{
}


var PB_MsgParam_BroadcastNewMessage__FOlD = &PB_MsgParam_BroadcastNewMessage{
        Text:  "" ,
        PeerId:  0 ,
        Time:  0 ,
        
}


var PB_MsgResponse_BroadcastNewMessage__FOlD = &PB_MsgResponse_BroadcastNewMessage{
}


var PB_MsgParam_GetFreshChatList__FOlD = &PB_MsgParam_GetFreshChatList{
}


var PB_MsgResponse_GetFreshChatList__FOlD = &PB_MsgResponse_GetFreshChatList{
        
        
}


var PB_MsgParam_GetFreshRoomMessagesList__FOlD = &PB_MsgParam_GetFreshRoomMessagesList{
        ChatKey:  "" ,
        RoomKey:  "" ,
        Last:  0 ,
}


var PB_MsgResponse_GetFreshRoomMessagesList__FOlD = &PB_MsgResponse_GetFreshRoomMessagesList{
        
        HasMore:  false ,
}


var PB_MsgParam_GetFreshAllDirectMessagesList__FOlD = &PB_MsgParam_GetFreshAllDirectMessagesList{
        LowMessageId:  0 ,
        HighMessageId:  0 ,
}


var PB_MsgResponse_GetFreshAllDirectMessagesList__FOlD = &PB_MsgResponse_GetFreshAllDirectMessagesList{
        
        HasMore:  false ,
}


var PB_MessageForwardedFrom__FOlD = &PB_MessageForwardedFrom{
        RoomId:  0 ,
        
        MessageId:  0 ,
        MessageSeq:  0 ,
}


var PB_MessagesCollections__FOlD = &PB_MessagesCollections{
        DirectMessagesIds:  0 ,
        GroupMessagesIds:  0 ,
        BroadCatMessagesIds:  0 ,
}


var PB_MsgParam_Echo__FOlD = &PB_MsgParam_Echo{
        Text:  "" ,
}


var PB_MsgResponse_PB_MsgParam_Echo__FOlD = &PB_MsgResponse_PB_MsgParam_Echo{
        Text:  "" ,
}


var PB_SyncParam_GetDirectUpdates__FOlD = &PB_SyncParam_GetDirectUpdates{
        LastId:  0 ,
}


var PB_SyncResponse_GetDirectUpdates__FOlD = &PB_SyncResponse_GetDirectUpdates{
        
        
        
        
        
        
        
        
        
        
        
        
        
        LastId:  0 ,
}


var PB_SyncParam_GetGeneralUpdates__FOlD = &PB_SyncParam_GetGeneralUpdates{
        LastId:  0 ,
}


var PB_SyncResponse_GetGeneralUpdates__FOlD = &PB_SyncResponse_GetGeneralUpdates{
        
        
        LastId:  0 ,
}


var PB_SyncParam_GetNotifyUpdates__FOlD = &PB_SyncParam_GetNotifyUpdates{
        LastId:  0 ,
}


var PB_SyncResponse_GetNotifyUpdates__FOlD = &PB_SyncResponse_GetNotifyUpdates{
        
        LastId:  0 ,
}


var PB_SyncParam_SetLastSyncDirectUpdateId__FOlD = &PB_SyncParam_SetLastSyncDirectUpdateId{
        LastId:  0 ,
}


var PB_SyncResponse_SetLastSyncDirectUpdateId__FOlD = &PB_SyncResponse_SetLastSyncDirectUpdateId{
}


var PB_SyncParam_SetLastSyncGeneralUpdateId__FOlD = &PB_SyncParam_SetLastSyncGeneralUpdateId{
        LastId:  0 ,
}


var PB_SyncResponse_SetLastSyncGeneralUpdateId__FOlD = &PB_SyncResponse_SetLastSyncGeneralUpdateId{
}


var PB_SyncParam_SetLastSyncNotifyUpdateId__FOlD = &PB_SyncParam_SetLastSyncNotifyUpdateId{
        LastId:  0 ,
}


var PB_SyncResponse_SetLastSyncNotifyUpdateId__FOlD = &PB_SyncResponse_SetLastSyncNotifyUpdateId{
}


var PB_NotifyUpdatesView__FOlD = &PB_NotifyUpdatesView{
        
        
}


var PB_AllLivePushes__FOlD = &PB_AllLivePushes{
        
        
        
}


var PB_UserParam_BlockUser__FOlD = &PB_UserParam_BlockUser{
        UserId:  0 ,
        UserName:  "" ,
}


var PB_UserOfflineResponse_BlockUser__FOlD = &PB_UserOfflineResponse_BlockUser{
        ByUserId:  0 ,
        TargetUserId:  0 ,
        TargetUserName:  "" ,
}


var PB_UserParam_UnBlockUser__FOlD = &PB_UserParam_UnBlockUser{
        Offset:  0 ,
        Limit:  0 ,
}


var PB_UserOfflineResponse_UnBlockUser__FOlD = &PB_UserOfflineResponse_UnBlockUser{
        
}


var PB_UserParam_BlockedList__FOlD = &PB_UserParam_BlockedList{
        UserId:  0 ,
        UserName:  "" ,
}


var PB_UserResponse_BlockedList__FOlD = &PB_UserResponse_BlockedList{
        ByUserId:  0 ,
        TargetUserId:  0 ,
        TargetUserName:  "" ,
}


var PB_UserParam_UpdateAbout__FOlD = &PB_UserParam_UpdateAbout{
        NewAbout:  "" ,
}


var PB_UserOfflineResponse_UpdateAbout__FOlD = &PB_UserOfflineResponse_UpdateAbout{
        UserId:  0 ,
        NewAbout:  "" ,
}


var PB_UserParam_UpdateUserName__FOlD = &PB_UserParam_UpdateUserName{
        NewUserName:  "" ,
}


var PB_UserOfflineResponse_UpdateUserName__FOlD = &PB_UserOfflineResponse_UpdateUserName{
        UserId:  0 ,
        NewUserName:  "" ,
}


var PB_UserParam_ChangeAvatar__FOlD = &PB_UserParam_ChangeAvatar{
        None:  false ,
        ImageData2:  []byte{} ,
}


var PB_UserOfflineResponse_ChangeAvatar__FOlD = &PB_UserOfflineResponse_ChangeAvatar{
}


var PB_UserParam_ChangePrivacy__FOlD = &PB_UserParam_ChangePrivacy{
        
}


var PB_UserResponseOffline_ChangePrivacy__FOlD = &PB_UserResponseOffline_ChangePrivacy{
}


var PB_UserParam_CheckUserName__FOlD = &PB_UserParam_CheckUserName{
        
}


var PB_UserResponse_CheckUserName__FOlD = &PB_UserResponse_CheckUserName{
}


var UserView__FOlD = &UserView{
}


var PB_Activity__FOlD = &PB_Activity{
        Id:  0 ,
        ActorUserId:  0 ,
        ActionTypeId:  0 ,
        RowId:  0 ,
        RootId:  0 ,
        RefId:  0 ,
        CreatedAt:  0 ,
}


var PB_Bucket__FOlD = &PB_Bucket{
        BucketId:  0 ,
        BucketName:  "" ,
        Server1Id:  0 ,
        Server2Id:  0 ,
        BackupServerId:  0 ,
        ContentObjectTypeId:  0 ,
        ContentObjectCount:  0 ,
        CreatedTime:  0 ,
}


var PB_Chat__FOlD = &PB_Chat{
        ChatKey:  "" ,
        RoomKey:  "" ,
        RoomTypeEnumId:  0 ,
        UserId:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        CreatedSe:  0 ,
        StartMessageIdFrom:  0 ,
        LastDeletedMessageId:  0 ,
        LastSeenMessageId:  0 ,
        LastMessageId:  0 ,
        UpdatedMs:  0 ,
}


var PB_Comment__FOlD = &PB_Comment{
        Id:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        Text:  "" ,
        CreatedTime:  0 ,
}


var PB_DirectMessage__FOlD = &PB_DirectMessage{
        MessageId:  0 ,
        MessageKey:  "" ,
        RoomKey:  "" ,
        UserId:  0 ,
        MessageFileId:  0 ,
        MessageTypeEnumId:  0 ,
        Text:  "" ,
        CreatedSe:  0 ,
        PeerReceivedTime:  0 ,
        PeerSeenTime:  0 ,
        DeliviryStatusEnumId:  0 ,
}


var PB_DirectOffline__FOlD = &PB_DirectOffline{
        DirectOfflineId:  0 ,
        ToUserId:  0 ,
        MessageId:  0 ,
        MessageFileId:  0 ,
        OtherId:  0 ,
        ChatKey:  "" ,
        PeerUserId:  0 ,
        RoomLogTypeId:  0 ,
        PBClass:  "" ,
        DataPB:  []byte{} ,
        DataJson:  "" ,
        DataTemp:  "" ,
        AtTimeMs:  0 ,
}


var PB_DirectToMessage__FOlD = &PB_DirectToMessage{
        Id:  0 ,
        ChatKey:  "" ,
        MessageId:  0 ,
        SourceEnumId:  0 ,
}


var PB_DirectUpdate__FOlD = &PB_DirectUpdate{
        DirectUpdateId:  0 ,
        ToUserId:  0 ,
        MessageId:  0 ,
        MessageFileId:  0 ,
        OtherId:  0 ,
        ChatKey:  "" ,
        PeerUserId:  0 ,
        EventType:  0 ,
        RoomLogTypeId:  0 ,
        FromSeq:  0 ,
        ToSeq:  0 ,
        ExtraPB:  []byte{} ,
        ExtraJson:  "" ,
        AtTimeMs:  0 ,
}


var PB_FollowingList__FOlD = &PB_FollowingList{
        Id:  0 ,
        UserId:  0 ,
        ListType:  0 ,
        Name:  "" ,
        Count:  0 ,
        IsAuto:  0 ,
        IsPimiry:  0 ,
        CreatedTime:  0 ,
}


var PB_FollowingListMember__FOlD = &PB_FollowingListMember{
        Id:  0 ,
        ListId:  0 ,
        UserId:  0 ,
        FollowedUserId:  0 ,
        FollowType:  0 ,
        UpdatedTimeMs:  0 ,
}


var PB_FollowingListMemberHistory__FOlD = &PB_FollowingListMemberHistory{
        Id:  0 ,
        ListId:  0 ,
        UserId:  0 ,
        FollowedUserId:  0 ,
        FollowType:  0 ,
        UpdatedTimeMs:  0 ,
        FollowId:  0 ,
}


var PB_GeneralLog__FOlD = &PB_GeneralLog{
        Id:  0 ,
        ToUserId:  0 ,
        TargetId:  0 ,
        LogTypeId:  0 ,
        ExtraPB:  []byte{} ,
        ExtraJson:  "" ,
        CreatedMs:  0 ,
}


var PB_Group__FOlD = &PB_Group{
        GroupId:  0 ,
        GroupName:  "" ,
        MembersCount:  0 ,
        GroupPrivacyEnum:  0 ,
        CreatorUserId:  0 ,
        CreatedTime:  0 ,
        UpdatedMs:  0 ,
        CurrentSeq:  0 ,
}


var PB_GroupMember__FOlD = &PB_GroupMember{
        Id:  0 ,
        GroupId:  0 ,
        GroupKey:  "" ,
        UserId:  0 ,
        ByUserId:  0 ,
        GroupRoleEnumId:  0 ,
        CreatedTime:  0 ,
}


var PB_GroupMessage__FOlD = &PB_GroupMessage{
        MessageId:  0 ,
        RoomKey:  "" ,
        UserId:  0 ,
        MessageFileId:  0 ,
        MessageTypeEnum:  0 ,
        Text:  "" ,
        CreatedMs:  0 ,
        DeliveryStatusEnum:  0 ,
}


var PB_GroupToMessage__FOlD = &PB_GroupToMessage{
        Id:  0 ,
        GroupId:  0 ,
        MessageId:  0 ,
        CreatedMs:  0 ,
        StatusEnum:  0 ,
}


var PB_Like__FOlD = &PB_Like{
        Id:  0 ,
        PostId:  0 ,
        PostTypeId:  0 ,
        UserId:  0 ,
        TypeId:  0 ,
        CreatedTime:  0 ,
}


var PB_LogChange__FOlD = &PB_LogChange{
        Id:  0 ,
        T:  "" ,
}


var PB_Media__FOlD = &PB_Media{
        Id:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        AlbumId:  0 ,
        TypeId:  0 ,
        CreatedTime:  0 ,
        Src:  "" ,
}


var PB_MessageFile__FOlD = &PB_MessageFile{
        MessageFileId:  0 ,
        MessageFileKey:  "" ,
        OriginalUserId:  0 ,
        Name:  "" ,
        Size:  0 ,
        FileTypeEnumId:  0 ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        HashMd5:  "" ,
        HashAccess:  0 ,
        CreatedSe:  0 ,
        ServerSrc:  "" ,
        ServerPath:  "" ,
        ServerThumbPath:  "" ,
        BucketId:  "" ,
        ServerId:  0 ,
        CanDel:  0 ,
}


var PB_Notification__FOlD = &PB_Notification{
        Id:  0 ,
        ForUserId:  0 ,
        ActorUserId:  0 ,
        ActionTypeId:  0 ,
        ObjectTypeId:  0 ,
        RowId:  0 ,
        RootId:  0 ,
        RefId:  0 ,
        SeenStatus:  0 ,
        CreatedTime:  0 ,
}


var PB_NotificationRemoved__FOlD = &PB_NotificationRemoved{
        NotificationId:  0 ,
        ForUserId:  0 ,
}


var PB_Offline__FOlD = &PB_Offline{
        Id:  0 ,
        FromUserId:  0 ,
        ToUserId:  0 ,
        RpcName:  "" ,
        PBClassName:  "" ,
        Key:  "" ,
        DataJson:  "" ,
        DataBlob:  []byte{} ,
        DataLength:  0 ,
        CreatedMs:  0 ,
}


var PB_OldMessage__FOlD = &PB_OldMessage{
        Id:  0 ,
        Uid:  0 ,
        UserId:  0 ,
        MessageKey:  "" ,
        RoomKey:  "" ,
        MessageType:  0 ,
        RoomType:  0 ,
        MsgFileId:  0 ,
        DataPB:  []byte{} ,
        Data64:  "" ,
        DataJson:  "" ,
        CreatedTimeMs:  0 ,
}


var PB_OldMsgFile__FOlD = &PB_OldMsgFile{
        Id:  0 ,
        Name:  "" ,
        Size:  0 ,
        FileType:  0 ,
        MimeType:  "" ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        ThumbData:  []byte{} ,
        ThumbData64:  "" ,
        ServerSrc:  "" ,
        ServerPath:  "" ,
        ServerId:  0 ,
        CanDel:  0 ,
}


var PB_OldMsgPush__FOlD = &PB_OldMsgPush{
        Id:  0 ,
        Uid:  0 ,
        ToUser:  0 ,
        MsgUid:  0 ,
        CreatedTimeMs:  0 ,
}


var PB_OldMsgPushEvent__FOlD = &PB_OldMsgPushEvent{
        Id:  0 ,
        Uid:  0 ,
        ToUserId:  0 ,
        MsgUid:  0 ,
        MsgKey:  "" ,
        RoomKey:  "" ,
        PeerUserId:  0 ,
        EventType:  0 ,
        AtTime:  0 ,
}


var PB_PhoneContact__FOlD = &PB_PhoneContact{
        Id:  0 ,
        PhoneDisplayName:  "" ,
        PhoneFamilyName:  "" ,
        PhoneNumber:  "" ,
        PhoneNormalizedNumber:  "" ,
        PhoneContactRowId:  0 ,
        UserId:  0 ,
        DeviceUuidId:  0 ,
        CreatedTime:  0 ,
        UpdatedTime:  0 ,
}


var PB_Photo__FOlD = &PB_Photo{
        PhotoId:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        AlbumId:  0 ,
        ImageTypeId:  0 ,
        Title:  "" ,
        Src:  "" ,
        PathSrc:  "" ,
        BucketId:  0 ,
        Width:  0 ,
        Height:  0 ,
        Ratio:  0.0 ,
        HashMd5:  "" ,
        Color:  "" ,
        CreatedTime:  0 ,
        W1080:  0 ,
        W720:  0 ,
        W480:  0 ,
        W320:  0 ,
        W160:  0 ,
        W80:  0 ,
}


var PB_Post__FOlD = &PB_Post{
        Id:  0 ,
        UserId:  0 ,
        TypeId:  0 ,
        Text:  "" ,
        FormatedText:  "" ,
        MediaCount:  0 ,
        SharedTo:  0 ,
        DisableComment:  0 ,
        HasTag:  0 ,
        LikesCount:  0 ,
        CommentsCount:  0 ,
        EditedTime:  0 ,
        CreatedTime:  0 ,
}


var PB_PushEvent__FOlD = &PB_PushEvent{
        PushEventId:  0 ,
        ToUserId:  0 ,
        ToDeviceId:  0 ,
        MessageId:  0 ,
        RoomTypeEnum:  0 ,
        RoomId:  0 ,
        PeerUserId:  0 ,
        PushEventTypeEnum:  0 ,
        AtTime:  0 ,
}


var PB_PushMessage__FOlD = &PB_PushMessage{
        PushMessageId:  0 ,
        ToUserId:  0 ,
        ToDeviceId:  0 ,
        MessageId:  0 ,
        RoomTypeEnum:  0 ,
        CreatedMs:  0 ,
}


var PB_RecommendUser__FOlD = &PB_RecommendUser{
        Id:  0 ,
        UserId:  0 ,
        TargetId:  0 ,
        Weight:  0.0 ,
        CreatedTime:  0 ,
}


var PB_Room__FOlD = &PB_Room{
        RoomId:  0 ,
        RoomKey:  "" ,
        RoomTypeEnum:  0 ,
        UserId:  0 ,
        LastSeqSeen:  0 ,
        LastSeqDelete:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        CreatedTime:  0 ,
        CurrentSeq:  0 ,
}


var PB_SearchClicked__FOlD = &PB_SearchClicked{
        Id:  0 ,
        Query:  "" ,
        ClickType:  0 ,
        TargetId:  0 ,
        UserId:  0 ,
        CreatedAt:  0 ,
}


var PB_Session__FOlD = &PB_Session{
        Id:  0 ,
        UserId:  0 ,
        SessionUuid:  "" ,
        ClientUuid:  "" ,
        DeviceUuid:  "" ,
        LastActivityTime:  0 ,
        LastIpAddress:  "" ,
        LastWifiMacAddress:  "" ,
        LastNetworkType:  "" ,
        LastNetworkTypeId:  0 ,
        AppVersion:  0 ,
        UpdatedTime:  0 ,
        CreatedTime:  0 ,
        LastSyncDirectUpdateId:  0 ,
        LastSyncGeneralUpdateId:  0 ,
        LastSyncNotifyUpdateId:  0 ,
}


var PB_SettingClient__FOlD = &PB_SettingClient{
        UserId:  0 ,
        AutoDownloadWifiVoice:  0 ,
        AutoDownloadWifiImage:  0 ,
        AutoDownloadWifiVideo:  0 ,
        AutoDownloadWifiFile:  0 ,
        AutoDownloadWifiMusic:  0 ,
        AutoDownloadWifiGif:  0 ,
        AutoDownloadCellularVoice:  0 ,
        AutoDownloadCellularImage:  0 ,
        AutoDownloadCellularVideo:  0 ,
        AutoDownloadCellularFile:  0 ,
        AutoDownloadCellularMusic:  0 ,
        AutoDownloadCellularGif:  0 ,
        AutoDownloadRoamingVoice:  0 ,
        AutoDownloadRoamingImage:  0 ,
        AutoDownloadRoamingVideo:  0 ,
        AutoDownloadRoamingFile:  0 ,
        AutoDownloadRoamingMusic:  0 ,
        AutoDownloadRoamingGif:  0 ,
        SaveToGallery:  0 ,
}


var PB_SettingNotification__FOlD = &PB_SettingNotification{
        UserId:  0 ,
        SocialLedOn:  0 ,
        SocialLedColor:  "" ,
        ReqestToFollowYou:  0 ,
        FollowedYou:  0 ,
        AccptedYourFollowRequest:  0 ,
        YourPostLiked:  0 ,
        YourPostCommented:  0 ,
        MenthenedYouInPost:  0 ,
        MenthenedYouInComment:  0 ,
        YourContactsJoined:  0 ,
        DirectMessage:  0 ,
        DirectAlert:  0 ,
        DirectPerview:  0 ,
        DirectLedOn:  0 ,
        DirectLedColor:  0 ,
        DirectVibrate:  0 ,
        DirectPopup:  0 ,
        DirectSound:  0 ,
        DirectPriority:  0 ,
}


var PB_Tag__FOlD = &PB_Tag{
        Id:  0 ,
        Name:  "" ,
        Count:  0 ,
        IsBlocked:  0 ,
        CreatedTime:  0 ,
}


var PB_TagsPost__FOlD = &PB_TagsPost{
        Id:  0 ,
        TagId:  0 ,
        PostId:  0 ,
        TypeId:  0 ,
        CreatedTime:  0 ,
}


var PB_TestChat__FOlD = &PB_TestChat{
        Id:  0 ,
        Id4:  0 ,
        TimeMs:  0 ,
        Text:  "" ,
        Name:  "" ,
        UserId:  0 ,
        C2:  0 ,
        C3:  0 ,
        C4:  0 ,
        C5:  0 ,
}


var PB_TriggerLog__FOlD = &PB_TriggerLog{
        Id:  0 ,
        ModelName:  "" ,
        ChangeType:  "" ,
        TargetId:  0 ,
        TargetStr:  "" ,
        CreatedSe:  0 ,
}


var PB_User__FOlD = &PB_User{
        Id:  0 ,
        UserName:  "" ,
        UserNameLower:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
        About:  "" ,
        FullName:  "" ,
        AvatarUrl:  "" ,
        PrivacyProfile:  0 ,
        Phone:  "" ,
        Email:  "" ,
        IsDeleted:  0 ,
        PasswordHash:  "" ,
        PasswordSalt:  "" ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        MediaCount:  0 ,
        LikesCount:  0 ,
        ResharedCount:  0 ,
        LastActionTime:  0 ,
        LastPostTime:  0 ,
        PrimaryFollowingList:  0 ,
        CreatedTime:  0 ,
        UpdatedTime:  0 ,
        SessionUuid:  "" ,
        DeviceUuid:  "" ,
        LastWifiMacAddress:  "" ,
        LastNetworkType:  "" ,
        AppVersion:  0 ,
        LastActivityTime:  0 ,
        LastLoginTime:  0 ,
        LastIpAddress:  "" ,
}


var PB_UserMetaInfo__FOlD = &PB_UserMetaInfo{
        Id:  0 ,
        UserId:  0 ,
        IsNotificationDirty:  0 ,
        LastUserRecGen:  0 ,
}


var PB_UserPassword__FOlD = &PB_UserPassword{
        UserId:  0 ,
        Password:  "" ,
        CreatedTime:  0 ,
}


var PB_UpdateGroupParticipants__FOlD = &PB_UpdateGroupParticipants{
}


var PB_UpdateNotifySettings__FOlD = &PB_UpdateNotifySettings{
}


var PB_UpdateServiceNotification__FOlD = &PB_UpdateServiceNotification{
}


var PB_UpdateMessageMeta__FOlD = &PB_UpdateMessageMeta{
        MessageId:  0 ,
        AtTime:  0 ,
}


var PB_UpdateMessageId__FOlD = &PB_UpdateMessageId{
        OldMessageId:  0 ,
        NewMessageId:  0 ,
}


var PB_UpdateMessageToEdit__FOlD = &PB_UpdateMessageToEdit{
        MessageId:  0 ,
        NewText:  "" ,
}


var PB_UpdateMessageToDelete__FOlD = &PB_UpdateMessageToDelete{
        MessageId:  0 ,
}


var PB_UpdateRoomActionDoing__FOlD = &PB_UpdateRoomActionDoing{
        RoomKey:  "" ,
        
}


var PB_UpdateUserBlocked__FOlD = &PB_UpdateUserBlocked{
        UserId:  0 ,
        Blocked:  false ,
}


var PB_ChatView__FOlD = &PB_ChatView{
        ChatKey:  "" ,
        RoomKey:  "" ,
        RoomTypeEnumId:  0 ,
        UserId:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        CreatedSe:  0 ,
        UpdatedMs:  0 ,
        LastMessageId:  0 ,
        LastDeletedMessageId:  0 ,
        LastSeenMessageId:  0 ,
        LastSeqSeen:  0 ,
        LastSeqDelete:  0 ,
        CurrentSeq:  0 ,
        
        SharedMediaCount:  0 ,
        UnseenCount:  0 ,
        
        
}


var PB_MessageView__FOlD = &PB_MessageView{
        MessageId:  0 ,
        MessageKey:  "" ,
        RoomKey:  "" ,
        UserId:  0 ,
        MessageFileId:  0 ,
        MessageTypeEnumId:  0 ,
        Text:  "" ,
        CreatedSe:  0 ,
        PeerReceivedTime:  0 ,
        PeerSeenTime:  0 ,
        DeliviryStatusEnumId:  0 ,
        ChatKey:  "" ,
        RoomTypeEnumId:  0 ,
        IsByMe:  false ,
        RemoteId:  0 ,
        
}


var PB_MessageFileView__FOlD = &PB_MessageFileView{
        MessageFileId:  0 ,
        MessageFileKey:  "" ,
        OriginalUserId:  0 ,
        Name:  "" ,
        Size:  0 ,
        FileTypeEnumId:  0 ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        HashMd5:  "" ,
        HashAccess:  0 ,
        CreatedSe:  0 ,
        ServerSrc:  "" ,
        ServerPath:  "" ,
        ServerThumbPath:  "" ,
        BucketId:  "" ,
        ServerId:  0 ,
        CanDel:  0 ,
        ServerThumbLocalSrc:  "" ,
        RemoteMessageFileId:  0 ,
        LocalSrc:  "" ,
        ThumbLocalSrc:  "" ,
        MessageFileStatusId:  "" ,
}


var PB_UserView__FOlD = &PB_UserView{
        UserId:  0 ,
        UserName:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
        About:  "" ,
        FullName:  "" ,
        AvatarUrl:  "" ,
        PrivacyProfile:  0 ,
        IsDeleted:  0 ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        UpdatedTime:  0 ,
        AppVersion:  0 ,
        LastActivityTime:  0 ,
        FollowingType:  0 ,
}



*/