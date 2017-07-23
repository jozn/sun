package ir.ms.pb

public class RPC_INTERFACES {

public interface RPC_MessageReq {
    GetLastChnagesForRoom( PB_ResponseLastChangesForTheRoom );}
public interface RPC_MessageReqOffline {
    SetLastSeen( PB_ResponseSetLastSeenMessages );}
public interface Greeter {
    SayHello( PB_UserWithMe );}
public interface RPC_Msg {
    AddNewTextMessage( PB_MsgResponse_AddNewTextMessage );
    SetRoomActionDoing( PB_MsgResponse_SetRoomActionDoing );
    GetMessagesByIds( PB_MsgResponse_GetMessagesByIds );
    GetMessagesHistory( PB_MsgResponse_GetMessagesHistory );
    SetMessagesRangeAsSeen( PB_MsgResponse_SetMessagesRangeAsSeen );
    DeleteRoomHistory( PB_MsgResponse_DeleteRoomHistory );
    DeleteMessagesByIds( PB_MsgResponse_DeleteMessagesByIds );
    SetMessagesAsReceived( PB_MsgResponse_SetMessagesAsReceived );
    ForwardMessages( PB_MsgResponse_ForwardMessages );
    EditMessage( PB_MsgResponse_EditMessage );
    BroadcastNewMessage( PB_MsgResponse_BroadcastNewMessage );}
public interface RpcMsgs {
    UploadNewMsg( PB_ResRpcAddMsg );}
public interface RPC_UserOffline {
    BlockUser( PB_UserOfflineResponse_BlockUser );
    UnBlockUser( PB_UserOfflineResponse_UnBlockUser );
    UpdateAbout( PB_UserOfflineResponse_UpdateAbout );
    UpdateUserName( PB_UserOfflineResponse_UpdateUserName );
    ChangePrivacy( PB_UserResponseOffline_ChangePrivacy );
    ChangeAvatar( PB_UserOfflineResponse_ChangeAvatar );}
public interface RPC_User {
    CheckUserName( PB_UserResponse_CheckUserName );
    GetBlockedList( PB_UserResponse_BlockedList );}
public interface RPC_Auth {
    CheckPhone( PB_UserResponse_CheckUserName );
    SendCode( PB_UserResponse_CheckUserName );
    SendCodeToSms( PB_UserResponse_CheckUserName );
    SendCodeToTelgram( PB_UserResponse_CheckUserName );
    SingUp( PB_UserResponse_CheckUserName );
    SingIn( PB_UserResponse_CheckUserName );
    LogOut( PB_UserResponse_CheckUserName );}
	
}
/*

RPC_INTERFACES.RPC_MessageReq RPC_MessageReq_Handeler = null;
RPC_INTERFACES.RPC_MessageReqOffline RPC_MessageReqOffline_Handeler = null;
RPC_INTERFACES.Greeter Greeter_Handeler = null;
RPC_INTERFACES.RPC_Msg RPC_Msg_Handeler = null;
RPC_INTERFACES.RpcMsgs RpcMsgs_Handeler = null;
RPC_INTERFACES.RPC_UserOffline RPC_UserOffline_Handeler = null;
RPC_INTERFACES.RPC_User RPC_User_Handeler = null;
RPC_INTERFACES.RPC_Auth RPC_Auth_Handeler = null;
	
*/