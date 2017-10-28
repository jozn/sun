package x

import (
    "strings"
    "github.com/golang/protobuf/proto"
    "errors"
    "ms/sun/config"
    "log"
)

type RPC_UserParam interface {
    GetUserId() int
    IsUser() bool
}

type RPC_ResponseHandlerInterface interface {
    //HandleOfflineResult(dataPB interface{},PBClass string,RpcName string,cmd PB_CommandToServer,p RPC_UserParam, paramParsed interface{})
    HandleOfflineResult(resOut RpcResponseOutput)
    IsUserOnlineResult(interface{}, error)
    HandelError(error)
}

type RpcResponseOutput struct {
    UserParam       RPC_UserParam
    ResponseData    interface{}
    PBClassName     string
    RpcName         string
    CommandToServer PB_CommandToServer
    RpcParamPassed  interface{}
}

var RPC_ResponseHandler RPC_ResponseHandlerInterface

//note: rpc methods cant have equal name they must be different even in different rpc services
type RPC_AllHandlersInteract struct {

   RPC_Auth RPC_Auth
   RPC_Chat RPC_Chat
   RPC_Msg RPC_Msg
   RPC_Sync RPC_Sync
   RPC_UserOffline RPC_UserOffline
   RPC_User RPC_User
}

/////////////// Interfaces ////////////////

type RPC_Auth interface {
    CheckPhone(i *PB_UserParam_CheckUserName2, p RPC_UserParam ) (*PB_UserResponse_CheckUserName2 ,error)
    SendCode(i *PB_UserParam_CheckUserName2, p RPC_UserParam ) (*PB_UserResponse_CheckUserName2 ,error)
    SendCodeToSms(i *PB_UserParam_CheckUserName2, p RPC_UserParam ) (*PB_UserResponse_CheckUserName2 ,error)
    SendCodeToTelgram(i *PB_UserParam_CheckUserName2, p RPC_UserParam ) (*PB_UserResponse_CheckUserName2 ,error)
    SingUp(i *PB_UserParam_CheckUserName2, p RPC_UserParam ) (*PB_UserResponse_CheckUserName2 ,error)
    SingIn(i *PB_UserParam_CheckUserName2, p RPC_UserParam ) (*PB_UserResponse_CheckUserName2 ,error)
    LogOut(i *PB_UserParam_CheckUserName2, p RPC_UserParam ) (*PB_UserResponse_CheckUserName2 ,error)
}

type RPC_Chat interface {
    AddNewMessage(i *PB_ChatParam_AddNewMessage, p RPC_UserParam ) (*PB_ChatResponse_AddNewMessage ,error)
    SetRoomActionDoing(i *PB_ChatParam_SetRoomActionDoing, p RPC_UserParam ) (*PB_ChatResponse_SetRoomActionDoing ,error)
    SetMessagesRangeAsSeen(i *PB_ChatParam_SetChatMessagesRangeAsSeen, p RPC_UserParam ) (*PB_ChatResponse_SetChatMessagesRangeAsSeen ,error)
    DeleteChatHistory(i *PB_ChatParam_DeleteChatHistory, p RPC_UserParam ) (*PB_ChatResponse_DeleteChatHistory ,error)
    DeleteMessagesByIds(i *PB_ChatParam_DeleteMessagesByIds, p RPC_UserParam ) (*PB_ChatResponse_DeleteMessagesByIds ,error)
    SetMessagesAsReceived(i *PB_ChatParam_SetMessagesAsReceived, p RPC_UserParam ) (*PB_ChatResponse_SetMessagesAsReceived ,error)
    EditMessage(i *PB_ChatParam_EditMessage, p RPC_UserParam ) (*PB_ChatResponse_EditMessage ,error)
    GetChatList(i *PB_ChatParam_GetChatList, p RPC_UserParam ) (*PB_ChatResponse_GetChatList ,error)
    GetChatHistoryToOlder(i *PB_ChatParam_GetChatHistoryToOlder, p RPC_UserParam ) (*PB_ChatResponse_GetChatHistoryToOlder ,error)
    GetFreshAllDirectMessagesList(i *PB_ChatParam_GetFreshAllDirectMessagesList, p RPC_UserParam ) (*PB_ChatResponse_GetFreshAllDirectMessagesList ,error)
}

type RPC_Msg interface {
    AddNewTextMessage(i *PB_MsgParam_AddNewTextMessage, p RPC_UserParam ) (*PB_MsgResponse_AddNewTextMessage ,error)
    AddNewMessage(i *PB_MsgParam_AddNewMessage, p RPC_UserParam ) (*PB_MsgResponse_AddNewMessage ,error)
    SetRoomActionDoing(i *PB_MsgParam_SetRoomActionDoing, p RPC_UserParam ) (*PB_MsgResponse_SetRoomActionDoing ,error)
    GetMessagesByIds(i *PB_MsgParam_GetMessagesByIds, p RPC_UserParam ) (*PB_MsgResponse_GetMessagesByIds ,error)
    GetMessagesHistory(i *PB_MsgParam_GetMessagesHistory, p RPC_UserParam ) (*PB_MsgResponse_GetMessagesHistory ,error)
    SetMessagesRangeAsSeen(i *PB_MsgParam_SetChatMessagesRangeAsSeen, p RPC_UserParam ) (*PB_MsgResponse_SetChatMessagesRangeAsSeen ,error)
    DeleteChatHistory(i *PB_MsgParam_DeleteChatHistory, p RPC_UserParam ) (*PB_MsgResponse_DeleteChatHistory ,error)
    DeleteMessagesByIds(i *PB_MsgParam_DeleteMessagesByIds, p RPC_UserParam ) (*PB_MsgResponse_DeleteMessagesByIds ,error)
    SetMessagesAsReceived(i *PB_MsgParam_SetMessagesAsReceived, p RPC_UserParam ) (*PB_MsgResponse_SetMessagesAsReceived ,error)
    ForwardMessages(i *PB_MsgParam_ForwardMessages, p RPC_UserParam ) (*PB_MsgResponse_ForwardMessages ,error)
    EditMessage(i *PB_MsgParam_EditMessage, p RPC_UserParam ) (*PB_MsgResponse_EditMessage ,error)
    BroadcastNewMessage(i *PB_MsgParam_BroadcastNewMessage, p RPC_UserParam ) (*PB_MsgResponse_BroadcastNewMessage ,error)
    GetFreshChatList(i *PB_MsgParam_GetFreshChatList, p RPC_UserParam ) (*PB_MsgResponse_GetFreshChatList ,error)
    GetFreshRoomMessagesList(i *PB_MsgParam_GetFreshRoomMessagesList, p RPC_UserParam ) (*PB_MsgResponse_GetFreshRoomMessagesList ,error)
    GetFreshAllDirectMessagesList(i *PB_MsgParam_GetFreshAllDirectMessagesList, p RPC_UserParam ) (*PB_MsgResponse_GetFreshAllDirectMessagesList ,error)
    Echo(i *PB_MsgParam_Echo, p RPC_UserParam ) (*PB_MsgResponse_PB_MsgParam_Echo ,error)
}

type RPC_Sync interface {
    GetDirectUpdates(i *PB_SyncParam_GetDirectUpdates, p RPC_UserParam ) (*PB_SyncResponse_GetDirectUpdates ,error)
    GetGeneralUpdates(i *PB_SyncParam_GetGeneralUpdates, p RPC_UserParam ) (*PB_SyncResponse_GetGeneralUpdates ,error)
    GetNotifyUpdates(i *PB_SyncParam_GetNotifyUpdates, p RPC_UserParam ) (*PB_SyncResponse_GetNotifyUpdates ,error)
    SetLastSyncDirectUpdateId(i *PB_SyncParam_SetLastSyncDirectUpdateId, p RPC_UserParam ) (*PB_SyncResponse_SetLastSyncDirectUpdateId ,error)
    SetLastSyncGeneralUpdateId(i *PB_SyncParam_SetLastSyncGeneralUpdateId, p RPC_UserParam ) (*PB_SyncResponse_SetLastSyncGeneralUpdateId ,error)
    SetLastSyncNotifyUpdateId(i *PB_SyncParam_SetLastSyncNotifyUpdateId, p RPC_UserParam ) (*PB_SyncResponse_SetLastSyncNotifyUpdateId ,error)
}

type RPC_UserOffline interface {
    BlockUser(i *PB_UserParam_BlockUser, p RPC_UserParam ) (*PB_UserOfflineResponse_BlockUser ,error)
    UnBlockUser(i *PB_UserParam_UnBlockUser, p RPC_UserParam ) (*PB_UserOfflineResponse_UnBlockUser ,error)
    UpdateAbout(i *PB_UserParam_UpdateAbout, p RPC_UserParam ) (*PB_UserOfflineResponse_UpdateAbout ,error)
    UpdateUserName(i *PB_UserParam_UpdateUserName, p RPC_UserParam ) (*PB_UserOfflineResponse_UpdateUserName ,error)
    ChangePrivacy(i *PB_UserParam_ChangePrivacy, p RPC_UserParam ) (*PB_UserResponseOffline_ChangePrivacy ,error)
    ChangeAvatar(i *PB_UserParam_ChangeAvatar, p RPC_UserParam ) (*PB_UserOfflineResponse_ChangeAvatar ,error)
}

type RPC_User interface {
    CheckUserName(i *PB_UserParam_CheckUserName, p RPC_UserParam ) (*PB_UserResponse_CheckUserName ,error)
    GetBlockedList(i *PB_UserParam_BlockedList, p RPC_UserParam ) (*PB_UserResponse_BlockedList ,error)
}


func noDevErr(err error)  {
    if config.IS_DEBUG && err != nil {
        log.Panic(err)
    }
}

////////////// map of rpc methods to all
func HandleRpcs(cmd PB_CommandToServer, params RPC_UserParam, rpcHandler RPC_AllHandlersInteract) {

    splits := strings.Split(cmd.Command, ".")

    if len(splits) != 2 {
        noDevErr(errors.New("HandleRpcs: splic is not 2 parts"))
        return
    }

    switch splits[0] {

    case "RPC_Auth":
            
            //rpc,ok := rpcHandler.RPC_Auth
            rpc := rpcHandler.RPC_Auth
            /*if !ok {
                e:=errors.New("rpcHandler could not be cast to : RPC_Auth")
                noDevErr(e)
                RPC_ResponseHandler.HandelError(e)
                return
            }*/

            switch splits[1]  {
                case "CheckPhone": //each pb_service_method
                    load := &PB_UserParam_CheckUserName2{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.CheckPhone(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Auth.CheckPhone",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_CheckUserName2",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.CheckPhone",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SendCode": //each pb_service_method
                    load := &PB_UserParam_CheckUserName2{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SendCode(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Auth.SendCode",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_CheckUserName2",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCode",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SendCodeToSms": //each pb_service_method
                    load := &PB_UserParam_CheckUserName2{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SendCodeToSms(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Auth.SendCodeToSms",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_CheckUserName2",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCodeToSms",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SendCodeToTelgram": //each pb_service_method
                    load := &PB_UserParam_CheckUserName2{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SendCodeToTelgram(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Auth.SendCodeToTelgram",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_CheckUserName2",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCodeToTelgram",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SingUp": //each pb_service_method
                    load := &PB_UserParam_CheckUserName2{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SingUp(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Auth.SingUp",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_CheckUserName2",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SingUp",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SingIn": //each pb_service_method
                    load := &PB_UserParam_CheckUserName2{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SingIn(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Auth.SingIn",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_CheckUserName2",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SingIn",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "LogOut": //each pb_service_method
                    load := &PB_UserParam_CheckUserName2{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.LogOut(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Auth.LogOut",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_CheckUserName2",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.LogOut",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                default:
                    noDevErr(errors.New("rpc method is does not exist: "+cmd.Command))
            }
    case "RPC_Chat":
            
            //rpc,ok := rpcHandler.RPC_Chat
            rpc := rpcHandler.RPC_Chat
            /*if !ok {
                e:=errors.New("rpcHandler could not be cast to : RPC_Chat")
                noDevErr(e)
                RPC_ResponseHandler.HandelError(e)
                return
            }*/

            switch splits[1]  {
                case "AddNewMessage": //each pb_service_method
                    load := &PB_ChatParam_AddNewMessage{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.AddNewMessage(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.AddNewMessage",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_AddNewMessage",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_AddNewMessage",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_AddNewMessage","RPC_Chat.AddNewMessage",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetRoomActionDoing": //each pb_service_method
                    load := &PB_ChatParam_SetRoomActionDoing{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetRoomActionDoing(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.SetRoomActionDoing",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_SetRoomActionDoing",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetRoomActionDoing",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetRoomActionDoing","RPC_Chat.SetRoomActionDoing",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetMessagesRangeAsSeen": //each pb_service_method
                    load := &PB_ChatParam_SetChatMessagesRangeAsSeen{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetMessagesRangeAsSeen(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.SetMessagesRangeAsSeen",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_SetChatMessagesRangeAsSeen",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetChatMessagesRangeAsSeen",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetChatMessagesRangeAsSeen","RPC_Chat.SetMessagesRangeAsSeen",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "DeleteChatHistory": //each pb_service_method
                    load := &PB_ChatParam_DeleteChatHistory{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.DeleteChatHistory(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.DeleteChatHistory",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_DeleteChatHistory",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteChatHistory",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteChatHistory","RPC_Chat.DeleteChatHistory",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "DeleteMessagesByIds": //each pb_service_method
                    load := &PB_ChatParam_DeleteMessagesByIds{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.DeleteMessagesByIds(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.DeleteMessagesByIds",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_DeleteMessagesByIds",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteMessagesByIds",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteMessagesByIds","RPC_Chat.DeleteMessagesByIds",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetMessagesAsReceived": //each pb_service_method
                    load := &PB_ChatParam_SetMessagesAsReceived{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetMessagesAsReceived(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.SetMessagesAsReceived",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_SetMessagesAsReceived",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetMessagesAsReceived",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetMessagesAsReceived","RPC_Chat.SetMessagesAsReceived",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "EditMessage": //each pb_service_method
                    load := &PB_ChatParam_EditMessage{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.EditMessage(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.EditMessage",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_EditMessage",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_EditMessage",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_EditMessage","RPC_Chat.EditMessage",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetChatList": //each pb_service_method
                    load := &PB_ChatParam_GetChatList{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetChatList(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.GetChatList",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_GetChatList",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatList",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatList","RPC_Chat.GetChatList",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetChatHistoryToOlder": //each pb_service_method
                    load := &PB_ChatParam_GetChatHistoryToOlder{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetChatHistoryToOlder(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.GetChatHistoryToOlder",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_GetChatHistoryToOlder",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatHistoryToOlder",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatHistoryToOlder","RPC_Chat.GetChatHistoryToOlder",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetFreshAllDirectMessagesList": //each pb_service_method
                    load := &PB_ChatParam_GetFreshAllDirectMessagesList{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetFreshAllDirectMessagesList(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Chat.GetFreshAllDirectMessagesList",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_ChatResponse_GetFreshAllDirectMessagesList",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetFreshAllDirectMessagesList",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetFreshAllDirectMessagesList","RPC_Chat.GetFreshAllDirectMessagesList",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                default:
                    noDevErr(errors.New("rpc method is does not exist: "+cmd.Command))
            }
    case "RPC_Msg":
            
            //rpc,ok := rpcHandler.RPC_Msg
            rpc := rpcHandler.RPC_Msg
            /*if !ok {
                e:=errors.New("rpcHandler could not be cast to : RPC_Msg")
                noDevErr(e)
                RPC_ResponseHandler.HandelError(e)
                return
            }*/

            switch splits[1]  {
                case "AddNewTextMessage": //each pb_service_method
                    load := &PB_MsgParam_AddNewTextMessage{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.AddNewTextMessage(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.AddNewTextMessage",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_AddNewTextMessage",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_AddNewTextMessage",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_AddNewTextMessage","RPC_Msg.AddNewTextMessage",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "AddNewMessage": //each pb_service_method
                    load := &PB_MsgParam_AddNewMessage{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.AddNewMessage(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.AddNewMessage",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_AddNewMessage",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_AddNewMessage",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_AddNewMessage","RPC_Msg.AddNewMessage",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetRoomActionDoing": //each pb_service_method
                    load := &PB_MsgParam_SetRoomActionDoing{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetRoomActionDoing(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.SetRoomActionDoing",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_SetRoomActionDoing",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_SetRoomActionDoing",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_SetRoomActionDoing","RPC_Msg.SetRoomActionDoing",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetMessagesByIds": //each pb_service_method
                    load := &PB_MsgParam_GetMessagesByIds{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetMessagesByIds(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.GetMessagesByIds",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_GetMessagesByIds",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetMessagesByIds",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetMessagesByIds","RPC_Msg.GetMessagesByIds",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetMessagesHistory": //each pb_service_method
                    load := &PB_MsgParam_GetMessagesHistory{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetMessagesHistory(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.GetMessagesHistory",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_GetMessagesHistory",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetMessagesHistory",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetMessagesHistory","RPC_Msg.GetMessagesHistory",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetMessagesRangeAsSeen": //each pb_service_method
                    load := &PB_MsgParam_SetChatMessagesRangeAsSeen{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetMessagesRangeAsSeen(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.SetMessagesRangeAsSeen",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_SetChatMessagesRangeAsSeen",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_SetChatMessagesRangeAsSeen",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_SetChatMessagesRangeAsSeen","RPC_Msg.SetMessagesRangeAsSeen",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "DeleteChatHistory": //each pb_service_method
                    load := &PB_MsgParam_DeleteChatHistory{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.DeleteChatHistory(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.DeleteChatHistory",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_DeleteChatHistory",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_DeleteChatHistory",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_DeleteChatHistory","RPC_Msg.DeleteChatHistory",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "DeleteMessagesByIds": //each pb_service_method
                    load := &PB_MsgParam_DeleteMessagesByIds{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.DeleteMessagesByIds(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.DeleteMessagesByIds",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_DeleteMessagesByIds",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_DeleteMessagesByIds",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_DeleteMessagesByIds","RPC_Msg.DeleteMessagesByIds",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetMessagesAsReceived": //each pb_service_method
                    load := &PB_MsgParam_SetMessagesAsReceived{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetMessagesAsReceived(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.SetMessagesAsReceived",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_SetMessagesAsReceived",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_SetMessagesAsReceived",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_SetMessagesAsReceived","RPC_Msg.SetMessagesAsReceived",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "ForwardMessages": //each pb_service_method
                    load := &PB_MsgParam_ForwardMessages{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.ForwardMessages(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.ForwardMessages",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_ForwardMessages",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_ForwardMessages",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_ForwardMessages","RPC_Msg.ForwardMessages",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "EditMessage": //each pb_service_method
                    load := &PB_MsgParam_EditMessage{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.EditMessage(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.EditMessage",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_EditMessage",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_EditMessage",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_EditMessage","RPC_Msg.EditMessage",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "BroadcastNewMessage": //each pb_service_method
                    load := &PB_MsgParam_BroadcastNewMessage{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.BroadcastNewMessage(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.BroadcastNewMessage",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_BroadcastNewMessage",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_BroadcastNewMessage",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_BroadcastNewMessage","RPC_Msg.BroadcastNewMessage",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetFreshChatList": //each pb_service_method
                    load := &PB_MsgParam_GetFreshChatList{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetFreshChatList(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.GetFreshChatList",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_GetFreshChatList",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetFreshChatList",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetFreshChatList","RPC_Msg.GetFreshChatList",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetFreshRoomMessagesList": //each pb_service_method
                    load := &PB_MsgParam_GetFreshRoomMessagesList{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetFreshRoomMessagesList(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.GetFreshRoomMessagesList",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_GetFreshRoomMessagesList",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetFreshRoomMessagesList",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetFreshRoomMessagesList","RPC_Msg.GetFreshRoomMessagesList",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetFreshAllDirectMessagesList": //each pb_service_method
                    load := &PB_MsgParam_GetFreshAllDirectMessagesList{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetFreshAllDirectMessagesList(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.GetFreshAllDirectMessagesList",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_GetFreshAllDirectMessagesList",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetFreshAllDirectMessagesList",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_GetFreshAllDirectMessagesList","RPC_Msg.GetFreshAllDirectMessagesList",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "Echo": //each pb_service_method
                    load := &PB_MsgParam_Echo{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.Echo(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Msg.Echo",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_MsgResponse_PB_MsgParam_Echo",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_PB_MsgParam_Echo",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_MsgResponse_PB_MsgParam_Echo","RPC_Msg.Echo",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                default:
                    noDevErr(errors.New("rpc method is does not exist: "+cmd.Command))
            }
    case "RPC_Sync":
            
            //rpc,ok := rpcHandler.RPC_Sync
            rpc := rpcHandler.RPC_Sync
            /*if !ok {
                e:=errors.New("rpcHandler could not be cast to : RPC_Sync")
                noDevErr(e)
                RPC_ResponseHandler.HandelError(e)
                return
            }*/

            switch splits[1]  {
                case "GetDirectUpdates": //each pb_service_method
                    load := &PB_SyncParam_GetDirectUpdates{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetDirectUpdates(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Sync.GetDirectUpdates",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_SyncResponse_GetDirectUpdates",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetDirectUpdates",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetDirectUpdates","RPC_Sync.GetDirectUpdates",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetGeneralUpdates": //each pb_service_method
                    load := &PB_SyncParam_GetGeneralUpdates{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetGeneralUpdates(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Sync.GetGeneralUpdates",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_SyncResponse_GetGeneralUpdates",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetGeneralUpdates",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetGeneralUpdates","RPC_Sync.GetGeneralUpdates",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetNotifyUpdates": //each pb_service_method
                    load := &PB_SyncParam_GetNotifyUpdates{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetNotifyUpdates(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Sync.GetNotifyUpdates",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_SyncResponse_GetNotifyUpdates",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetNotifyUpdates",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetNotifyUpdates","RPC_Sync.GetNotifyUpdates",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetLastSyncDirectUpdateId": //each pb_service_method
                    load := &PB_SyncParam_SetLastSyncDirectUpdateId{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetLastSyncDirectUpdateId(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Sync.SetLastSyncDirectUpdateId",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_SyncResponse_SetLastSyncDirectUpdateId",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncDirectUpdateId",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncDirectUpdateId","RPC_Sync.SetLastSyncDirectUpdateId",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetLastSyncGeneralUpdateId": //each pb_service_method
                    load := &PB_SyncParam_SetLastSyncGeneralUpdateId{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetLastSyncGeneralUpdateId(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Sync.SetLastSyncGeneralUpdateId",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_SyncResponse_SetLastSyncGeneralUpdateId",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncGeneralUpdateId",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncGeneralUpdateId","RPC_Sync.SetLastSyncGeneralUpdateId",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "SetLastSyncNotifyUpdateId": //each pb_service_method
                    load := &PB_SyncParam_SetLastSyncNotifyUpdateId{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetLastSyncNotifyUpdateId(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_Sync.SetLastSyncNotifyUpdateId",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_SyncResponse_SetLastSyncNotifyUpdateId",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncNotifyUpdateId",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncNotifyUpdateId","RPC_Sync.SetLastSyncNotifyUpdateId",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                default:
                    noDevErr(errors.New("rpc method is does not exist: "+cmd.Command))
            }
    case "RPC_UserOffline":
            
            //rpc,ok := rpcHandler.RPC_UserOffline
            rpc := rpcHandler.RPC_UserOffline
            /*if !ok {
                e:=errors.New("rpcHandler could not be cast to : RPC_UserOffline")
                noDevErr(e)
                RPC_ResponseHandler.HandelError(e)
                return
            }*/

            switch splits[1]  {
                case "BlockUser": //each pb_service_method
                    load := &PB_UserParam_BlockUser{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.BlockUser(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_UserOffline.BlockUser",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserOfflineResponse_BlockUser",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_BlockUser",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_BlockUser","RPC_UserOffline.BlockUser",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "UnBlockUser": //each pb_service_method
                    load := &PB_UserParam_UnBlockUser{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.UnBlockUser(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_UserOffline.UnBlockUser",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserOfflineResponse_UnBlockUser",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UnBlockUser",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UnBlockUser","RPC_UserOffline.UnBlockUser",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "UpdateAbout": //each pb_service_method
                    load := &PB_UserParam_UpdateAbout{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.UpdateAbout(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_UserOffline.UpdateAbout",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserOfflineResponse_UpdateAbout",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UpdateAbout",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UpdateAbout","RPC_UserOffline.UpdateAbout",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "UpdateUserName": //each pb_service_method
                    load := &PB_UserParam_UpdateUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.UpdateUserName(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_UserOffline.UpdateUserName",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserOfflineResponse_UpdateUserName",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UpdateUserName",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UpdateUserName","RPC_UserOffline.UpdateUserName",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "ChangePrivacy": //each pb_service_method
                    load := &PB_UserParam_ChangePrivacy{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.ChangePrivacy(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_UserOffline.ChangePrivacy",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponseOffline_ChangePrivacy",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponseOffline_ChangePrivacy",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponseOffline_ChangePrivacy","RPC_UserOffline.ChangePrivacy",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "ChangeAvatar": //each pb_service_method
                    load := &PB_UserParam_ChangeAvatar{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.ChangeAvatar(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_UserOffline.ChangeAvatar",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserOfflineResponse_ChangeAvatar",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_ChangeAvatar",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_ChangeAvatar","RPC_UserOffline.ChangeAvatar",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                default:
                    noDevErr(errors.New("rpc method is does not exist: "+cmd.Command))
            }
    case "RPC_User":
            
            //rpc,ok := rpcHandler.RPC_User
            rpc := rpcHandler.RPC_User
            /*if !ok {
                e:=errors.New("rpcHandler could not be cast to : RPC_User")
                noDevErr(e)
                RPC_ResponseHandler.HandelError(e)
                return
            }*/

            switch splits[1]  {
                case "CheckUserName": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.CheckUserName(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_User.CheckUserName",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_CheckUserName",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName","RPC_User.CheckUserName",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                case "GetBlockedList": //each pb_service_method
                    load := &PB_UserParam_BlockedList{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetBlockedList(load,params)
                        if err == nil {
                            out:= RpcResponseOutput{
                                RpcName: "RPC_User.GetBlockedList",
                                UserParam: params,
                                CommandToServer: cmd,
                                PBClassName: "PB_UserResponse_BlockedList",
                                ResponseData: res,
                                RpcParamPassed: load,
                            }
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_BlockedList",cmd, params)
                            //RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_BlockedList","RPC_User.GetBlockedList",cmd, params , load)
                            RPC_ResponseHandler.HandleOfflineResult(out)
                        }else{
                            RPC_ResponseHandler.HandelError(err)
                        }
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                default:
                    noDevErr(errors.New("rpc method is does not exist: "+cmd.Command))
            }
    default:
    noDevErr(errors.New("rpc dosent exisit for: "+cmd.Command))
}
}

/////////////// Direct in PB_CommandToClient /////////////
/*


 RPC_Auth.CheckPhone
 RPC_Auth.SendCode
 RPC_Auth.SendCodeToSms
 RPC_Auth.SendCodeToTelgram
 RPC_Auth.SingUp
 RPC_Auth.SingIn
 RPC_Auth.LogOut



 RPC_Chat.AddNewMessage
 RPC_Chat.SetRoomActionDoing
 RPC_Chat.SetMessagesRangeAsSeen
 RPC_Chat.DeleteChatHistory
 RPC_Chat.DeleteMessagesByIds
 RPC_Chat.SetMessagesAsReceived
 RPC_Chat.EditMessage
 RPC_Chat.GetChatList
 RPC_Chat.GetChatHistoryToOlder
 RPC_Chat.GetFreshAllDirectMessagesList



 RPC_Msg.AddNewTextMessage
 RPC_Msg.AddNewMessage
 RPC_Msg.SetRoomActionDoing
 RPC_Msg.GetMessagesByIds
 RPC_Msg.GetMessagesHistory
 RPC_Msg.SetMessagesRangeAsSeen
 RPC_Msg.DeleteChatHistory
 RPC_Msg.DeleteMessagesByIds
 RPC_Msg.SetMessagesAsReceived
 RPC_Msg.ForwardMessages
 RPC_Msg.EditMessage
 RPC_Msg.BroadcastNewMessage
 RPC_Msg.GetFreshChatList
 RPC_Msg.GetFreshRoomMessagesList
 RPC_Msg.GetFreshAllDirectMessagesList
 RPC_Msg.Echo



 RPC_Sync.GetDirectUpdates
 RPC_Sync.GetGeneralUpdates
 RPC_Sync.GetNotifyUpdates
 RPC_Sync.SetLastSyncDirectUpdateId
 RPC_Sync.SetLastSyncGeneralUpdateId
 RPC_Sync.SetLastSyncNotifyUpdateId



 RPC_UserOffline.BlockUser
 RPC_UserOffline.UnBlockUser
 RPC_UserOffline.UpdateAbout
 RPC_UserOffline.UpdateUserName
 RPC_UserOffline.ChangePrivacy
 RPC_UserOffline.ChangeAvatar



 RPC_User.CheckUserName
 RPC_User.GetBlockedList


*/