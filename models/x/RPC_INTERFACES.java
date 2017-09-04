package ir.ms.pb;

import android.util.Log;

public class RPC_INTERFACES {

public interface RPC_MessageReq {
    void GetLastChnagesForRoom( PB_ResponseLastChangesForTheRoom pb, boolean handled);}
public interface RPC_MessageReqOffline {
    void SetLastSeen( PB_ResponseSetLastSeenMessages pb, boolean handled);}
public interface RpcMsgs {
    void UploadNewMsg( PB_ResRpcAddMsg pb, boolean handled);}
public interface RPC_Auth {
    void CheckPhone( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SendCode( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SendCodeToSms( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SendCodeToTelgram( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SingUp( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SingIn( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void LogOut( PB_UserResponse_CheckUserName2 pb, boolean handled);}
public interface RPC_Msg {
    void AddNewTextMessage( PB_MsgResponse_AddNewTextMessage pb, boolean handled);
    void SetRoomActionDoing( PB_MsgResponse_SetRoomActionDoing pb, boolean handled);
    void GetMessagesByIds( PB_MsgResponse_GetMessagesByIds pb, boolean handled);
    void GetMessagesHistory( PB_MsgResponse_GetMessagesHistory pb, boolean handled);
    void SetMessagesRangeAsSeen( PB_MsgResponse_SetChatMessagesRangeAsSeen pb, boolean handled);
    void DeleteChatHistory( PB_MsgResponse_DeleteChatHistory pb, boolean handled);
    void DeleteMessagesByIds( PB_MsgResponse_DeleteMessagesByIds pb, boolean handled);
    void SetMessagesAsReceived( PB_MsgResponse_SetMessagesAsReceived pb, boolean handled);
    void ForwardMessages( PB_MsgResponse_ForwardMessages pb, boolean handled);
    void EditMessage( PB_MsgResponse_EditMessage pb, boolean handled);
    void BroadcastNewMessage( PB_MsgResponse_BroadcastNewMessage pb, boolean handled);
    void Echo( PB_MsgResponse_PB_MsgParam_Echo pb, boolean handled);}
public interface RPC_UserOffline {
    void BlockUser( PB_UserOfflineResponse_BlockUser pb, boolean handled);
    void UnBlockUser( PB_UserOfflineResponse_UnBlockUser pb, boolean handled);
    void UpdateAbout( PB_UserOfflineResponse_UpdateAbout pb, boolean handled);
    void UpdateUserName( PB_UserOfflineResponse_UpdateUserName pb, boolean handled);
    void ChangePrivacy( PB_UserResponseOffline_ChangePrivacy pb, boolean handled);
    void ChangeAvatar( PB_UserOfflineResponse_ChangeAvatar pb, boolean handled);}
public interface RPC_User {
    void CheckUserName( PB_UserResponse_CheckUserName pb, boolean handled);
    void GetBlockedList( PB_UserResponse_BlockedList pb, boolean handled);}


  public static class RPC_MessageReq_Empty implements RPC_MessageReq{
  
  	@Override
    public void GetLastChnagesForRoom( PB_ResponseLastChangesForTheRoom pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_MessageReq.GetLastChnagesForRoom' ");
    }
  }
  public static class RPC_MessageReqOffline_Empty implements RPC_MessageReqOffline{
  
  	@Override
    public void SetLastSeen( PB_ResponseSetLastSeenMessages pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_MessageReqOffline.SetLastSeen' ");
    }
  }
  public static class RpcMsgs_Empty implements RpcMsgs{
  
  	@Override
    public void UploadNewMsg( PB_ResRpcAddMsg pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RpcMsgs.UploadNewMsg' ");
    }
  }
  public static class RPC_Auth_Empty implements RPC_Auth{
  
  	@Override
    public void CheckPhone( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.CheckPhone' ");
    }
  	@Override
    public void SendCode( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SendCode' ");
    }
  	@Override
    public void SendCodeToSms( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SendCodeToSms' ");
    }
  	@Override
    public void SendCodeToTelgram( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SendCodeToTelgram' ");
    }
  	@Override
    public void SingUp( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SingUp' ");
    }
  	@Override
    public void SingIn( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SingIn' ");
    }
  	@Override
    public void LogOut( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.LogOut' ");
    }
  }
  public static class RPC_Msg_Empty implements RPC_Msg{
  
  	@Override
    public void AddNewTextMessage( PB_MsgResponse_AddNewTextMessage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.AddNewTextMessage' ");
    }
  	@Override
    public void SetRoomActionDoing( PB_MsgResponse_SetRoomActionDoing pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.SetRoomActionDoing' ");
    }
  	@Override
    public void GetMessagesByIds( PB_MsgResponse_GetMessagesByIds pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.GetMessagesByIds' ");
    }
  	@Override
    public void GetMessagesHistory( PB_MsgResponse_GetMessagesHistory pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.GetMessagesHistory' ");
    }
  	@Override
    public void SetMessagesRangeAsSeen( PB_MsgResponse_SetChatMessagesRangeAsSeen pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.SetMessagesRangeAsSeen' ");
    }
  	@Override
    public void DeleteChatHistory( PB_MsgResponse_DeleteChatHistory pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.DeleteChatHistory' ");
    }
  	@Override
    public void DeleteMessagesByIds( PB_MsgResponse_DeleteMessagesByIds pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.DeleteMessagesByIds' ");
    }
  	@Override
    public void SetMessagesAsReceived( PB_MsgResponse_SetMessagesAsReceived pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.SetMessagesAsReceived' ");
    }
  	@Override
    public void ForwardMessages( PB_MsgResponse_ForwardMessages pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.ForwardMessages' ");
    }
  	@Override
    public void EditMessage( PB_MsgResponse_EditMessage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.EditMessage' ");
    }
  	@Override
    public void BroadcastNewMessage( PB_MsgResponse_BroadcastNewMessage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.BroadcastNewMessage' ");
    }
  	@Override
    public void Echo( PB_MsgResponse_PB_MsgParam_Echo pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Msg.Echo' ");
    }
  }
  public static class RPC_UserOffline_Empty implements RPC_UserOffline{
  
  	@Override
    public void BlockUser( PB_UserOfflineResponse_BlockUser pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_UserOffline.BlockUser' ");
    }
  	@Override
    public void UnBlockUser( PB_UserOfflineResponse_UnBlockUser pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_UserOffline.UnBlockUser' ");
    }
  	@Override
    public void UpdateAbout( PB_UserOfflineResponse_UpdateAbout pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_UserOffline.UpdateAbout' ");
    }
  	@Override
    public void UpdateUserName( PB_UserOfflineResponse_UpdateUserName pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_UserOffline.UpdateUserName' ");
    }
  	@Override
    public void ChangePrivacy( PB_UserResponseOffline_ChangePrivacy pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_UserOffline.ChangePrivacy' ");
    }
  	@Override
    public void ChangeAvatar( PB_UserOfflineResponse_ChangeAvatar pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_UserOffline.ChangeAvatar' ");
    }
  }
  public static class RPC_User_Empty implements RPC_User{
  
  	@Override
    public void CheckUserName( PB_UserResponse_CheckUserName pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.CheckUserName' ");
    }
  	@Override
    public void GetBlockedList( PB_UserResponse_BlockedList pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.GetBlockedList' ");
    }
  }
	
}
/*

RPC_INTERFACES.RPC_MessageReq RPC_MessageReq_Empty_Handler = new RPC_INTERFACES.RPC_MessageReq RPC_MessageReq_Empty();
RPC_INTERFACES.RPC_MessageReqOffline RPC_MessageReqOffline_Empty_Handler = new RPC_INTERFACES.RPC_MessageReqOffline RPC_MessageReqOffline_Empty();
RPC_INTERFACES.RpcMsgs RpcMsgs_Empty_Handler = new RPC_INTERFACES.RpcMsgs RpcMsgs_Empty();
RPC_INTERFACES.RPC_Auth RPC_Auth_Empty_Handler = new RPC_INTERFACES.RPC_Auth RPC_Auth_Empty();
RPC_INTERFACES.RPC_Msg RPC_Msg_Empty_Handler = new RPC_INTERFACES.RPC_Msg RPC_Msg_Empty();
RPC_INTERFACES.RPC_UserOffline RPC_UserOffline_Empty_Handler = new RPC_INTERFACES.RPC_UserOffline RPC_UserOffline_Empty();
RPC_INTERFACES.RPC_User RPC_User_Empty_Handler = new RPC_INTERFACES.RPC_User RPC_User_Empty();
	
*/