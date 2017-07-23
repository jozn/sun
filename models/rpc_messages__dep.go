package models

import "ms/sun/models/x"

type RPC_MessageReqImple int
type RPC2 int

func (RPC2) GetLastChnagesForRoom(context.Context, *x.PB_ReqLastChangesForTheRoom) (*x.PB_ResponseLastChangesForTheRoom, error) {
    panic("implement me")
}

func mm()  {
    c := struct {
        RPC_MessageReqImple
        RPC2
    }{
        c2 :x.Activity__{},
    }
    c.GeLa


}



func (pv RPC_MessageReqImple) GetLastChnagesForRoom2(x.PB_ReqLastChangesForTheRoom) (x.PB_ResponseLastChangesForTheRoom, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) SetLastSeen(x.PB_RequestSetLastSeenMessages) (x.PB_ResponseSetLastSeenMessages, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) SayHello(x.PB_Message) (x.PB_UserWithMe, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) UploadNewMsg(x.PB_Message) (x.PB_ResRpcAddMsg, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) BlockUser(x.PB_UserParam_BlockUser) (x.PB_UserOfflineResponse_BlockUser, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) UnBlockUser(x.PB_UserParam_UnBlockUser) (x.PB_UserOfflineResponse_UnBlockUser, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) UpdateAbout(x.PB_UserParam_UpdateAbout) (x.PB_UserOfflineResponse_UpdateAbout, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) UpdateUserName(x.PB_UserParam_UpdateUserName) (x.PB_UserOfflineResponse_UpdateUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) ChangePrivacy(x.PB_UserParam_ChangePrivacy) (x.PB_UserResponseOffline_ChangePrivacy, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) ChangeAvatar(x.PB_UserParam_ChangeAvatar) (x.PB_UserOfflineResponse_ChangeAvatar, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) CheckPhone(x.PB_UserParam_CheckUserName) (x.PB_UserResponse_CheckUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) SendCode(x.PB_UserParam_CheckUserName) (x.PB_UserResponse_CheckUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) SendCodeToSms(x.PB_UserParam_CheckUserName) (x.PB_UserResponse_CheckUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) SendCodeToTelgram(x.PB_UserParam_CheckUserName) (x.PB_UserResponse_CheckUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) SingUp(x.PB_UserParam_CheckUserName) (x.PB_UserResponse_CheckUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) SingIn(x.PB_UserParam_CheckUserName) (x.PB_UserResponse_CheckUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) LogOut(x.PB_UserParam_CheckUserName) (x.PB_UserResponse_CheckUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) CheckUserName(x.PB_UserParam_CheckUserName) (x.PB_UserResponse_CheckUserName, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple) GetBlockedList(x.PB_UserParam_BlockedList) (x.PB_UserResponse_BlockedList, error) {
    panic("implement me")
}

func (pv RPC_MessageReqImple)s() {}

