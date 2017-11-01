package x

import (
	"errors"
	"log"
	"ms/sun/config"
	"strings"

	"github.com/golang/protobuf/proto"
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
	RPC_Auth        RPC_Auth
	RPC_Chat        RPC_Chat
	RPC_Other       RPC_Other
	RPC_Sync        RPC_Sync
	RPC_UserOffline RPC_UserOffline
	RPC_User        RPC_User
}

/////////////// Interfaces ////////////////

type RPC_Auth interface {
	CheckPhone(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SendCode(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SendCodeToSms(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SendCodeToTelgram(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SingUp(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SingIn(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	LogOut(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
}

type RPC_Chat interface {
	AddNewMessage(param *PB_ChatParam_AddNewMessage, userParam RPC_UserParam) (res PB_ChatResponse_AddNewMessage, err error)
	SetRoomActionDoing(param *PB_ChatParam_SetRoomActionDoing, userParam RPC_UserParam) (res PB_ChatResponse_SetRoomActionDoing, err error)
	SetMessagesRangeAsSeen(param *PB_ChatParam_SetChatMessagesRangeAsSeen, userParam RPC_UserParam) (res PB_ChatResponse_SetChatMessagesRangeAsSeen, err error)
	DeleteChatHistory(param *PB_ChatParam_DeleteChatHistory, userParam RPC_UserParam) (res PB_ChatResponse_DeleteChatHistory, err error)
	DeleteMessagesByIds(param *PB_ChatParam_DeleteMessagesByIds, userParam RPC_UserParam) (res PB_ChatResponse_DeleteMessagesByIds, err error)
	SetMessagesAsReceived(param *PB_ChatParam_SetMessagesAsReceived, userParam RPC_UserParam) (res PB_ChatResponse_SetMessagesAsReceived, err error)
	EditMessage(param *PB_ChatParam_EditMessage, userParam RPC_UserParam) (res PB_ChatResponse_EditMessage, err error)
	GetChatList(param *PB_ChatParam_GetChatList, userParam RPC_UserParam) (res PB_ChatResponse_GetChatList, err error)
	GetChatHistoryToOlder(param *PB_ChatParam_GetChatHistoryToOlder, userParam RPC_UserParam) (res PB_ChatResponse_GetChatHistoryToOlder, err error)
	GetFreshAllDirectMessagesList(param *PB_ChatParam_GetFreshAllDirectMessagesList, userParam RPC_UserParam) (res PB_ChatResponse_GetFreshAllDirectMessagesList, err error)
}

type RPC_Other interface {
	Echo(param *PB_OtherParam_Echo, userParam RPC_UserParam) (res PB_OtherResponse_Echo, err error)
}

type RPC_Sync interface {
	GetGeneralUpdates(param *PB_SyncParam_GetGeneralUpdates, userParam RPC_UserParam) (res PB_SyncResponse_GetGeneralUpdates, err error)
	GetNotifyUpdates(param *PB_SyncParam_GetNotifyUpdates, userParam RPC_UserParam) (res PB_SyncResponse_GetNotifyUpdates, err error)
	SetLastSyncDirectUpdateId(param *PB_SyncParam_SetLastSyncDirectUpdateId, userParam RPC_UserParam) (res PB_SyncResponse_SetLastSyncDirectUpdateId, err error)
	SetLastSyncGeneralUpdateId(param *PB_SyncParam_SetLastSyncGeneralUpdateId, userParam RPC_UserParam) (res PB_SyncResponse_SetLastSyncGeneralUpdateId, err error)
	SetLastSyncNotifyUpdateId(param *PB_SyncParam_SetLastSyncNotifyUpdateId, userParam RPC_UserParam) (res PB_SyncResponse_SetLastSyncNotifyUpdateId, err error)
}

type RPC_UserOffline interface {
	BlockUser(param *PB_UserParam_BlockUser, userParam RPC_UserParam) (res PB_UserOfflineResponse_BlockUser, err error)
	UnBlockUser(param *PB_UserParam_UnBlockUser, userParam RPC_UserParam) (res PB_UserOfflineResponse_UnBlockUser, err error)
	UpdateAbout(param *PB_UserParam_UpdateAbout, userParam RPC_UserParam) (res PB_UserOfflineResponse_UpdateAbout, err error)
	UpdateUserName(param *PB_UserParam_UpdateUserName, userParam RPC_UserParam) (res PB_UserOfflineResponse_UpdateUserName, err error)
	ChangePrivacy(param *PB_UserParam_ChangePrivacy, userParam RPC_UserParam) (res PB_UserResponseOffline_ChangePrivacy, err error)
	ChangeAvatar(param *PB_UserParam_ChangeAvatar, userParam RPC_UserParam) (res PB_UserOfflineResponse_ChangeAvatar, err error)
}

type RPC_User interface {
	CheckUserName(param *PB_UserParam_CheckUserName, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName, err error)
	GetBlockedList(param *PB_UserParam_BlockedList, userParam RPC_UserParam) (res PB_UserResponse_BlockedList, err error)
}

func noDevErr(err error) {
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

		switch splits[1] {
		case "CheckPhone": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.CheckPhone(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.CheckPhone",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.CheckPhone",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SendCode": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCode(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SendCode",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCode",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SendCodeToSms": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCodeToSms(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SendCodeToSms",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCodeToSms",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SendCodeToTelgram": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCodeToTelgram(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SendCodeToTelgram",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCodeToTelgram",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SingUp": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SingUp(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SingUp",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SingUp",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SingIn": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SingIn(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SingIn",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SingIn",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "LogOut": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.LogOut(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.LogOut",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.LogOut",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
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

		switch splits[1] {
		case "AddNewMessage": //each pb_service_method
			load := &PB_ChatParam_AddNewMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.AddNewMessage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.AddNewMessage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_AddNewMessage",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_AddNewMessage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_AddNewMessage","RPC_Chat.AddNewMessage",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetRoomActionDoing": //each pb_service_method
			load := &PB_ChatParam_SetRoomActionDoing{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetRoomActionDoing(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.SetRoomActionDoing",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_SetRoomActionDoing",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetRoomActionDoing",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetRoomActionDoing","RPC_Chat.SetRoomActionDoing",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetMessagesRangeAsSeen": //each pb_service_method
			load := &PB_ChatParam_SetChatMessagesRangeAsSeen{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetMessagesRangeAsSeen(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.SetMessagesRangeAsSeen",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_SetChatMessagesRangeAsSeen",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetChatMessagesRangeAsSeen",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetChatMessagesRangeAsSeen","RPC_Chat.SetMessagesRangeAsSeen",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "DeleteChatHistory": //each pb_service_method
			load := &PB_ChatParam_DeleteChatHistory{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteChatHistory(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.DeleteChatHistory",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_DeleteChatHistory",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteChatHistory",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteChatHistory","RPC_Chat.DeleteChatHistory",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "DeleteMessagesByIds": //each pb_service_method
			load := &PB_ChatParam_DeleteMessagesByIds{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteMessagesByIds(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.DeleteMessagesByIds",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_DeleteMessagesByIds",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteMessagesByIds",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteMessagesByIds","RPC_Chat.DeleteMessagesByIds",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetMessagesAsReceived": //each pb_service_method
			load := &PB_ChatParam_SetMessagesAsReceived{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetMessagesAsReceived(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.SetMessagesAsReceived",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_SetMessagesAsReceived",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetMessagesAsReceived",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetMessagesAsReceived","RPC_Chat.SetMessagesAsReceived",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "EditMessage": //each pb_service_method
			load := &PB_ChatParam_EditMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.EditMessage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.EditMessage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_EditMessage",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_EditMessage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_EditMessage","RPC_Chat.EditMessage",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetChatList": //each pb_service_method
			load := &PB_ChatParam_GetChatList{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetChatList(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.GetChatList",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_GetChatList",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatList",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatList","RPC_Chat.GetChatList",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetChatHistoryToOlder": //each pb_service_method
			load := &PB_ChatParam_GetChatHistoryToOlder{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetChatHistoryToOlder(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.GetChatHistoryToOlder",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_GetChatHistoryToOlder",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatHistoryToOlder",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatHistoryToOlder","RPC_Chat.GetChatHistoryToOlder",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetFreshAllDirectMessagesList": //each pb_service_method
			load := &PB_ChatParam_GetFreshAllDirectMessagesList{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetFreshAllDirectMessagesList(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.GetFreshAllDirectMessagesList",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_GetFreshAllDirectMessagesList",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetFreshAllDirectMessagesList",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetFreshAllDirectMessagesList","RPC_Chat.GetFreshAllDirectMessagesList",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_Other":

		//rpc,ok := rpcHandler.RPC_Other
		rpc := rpcHandler.RPC_Other
		/*if !ok {
		    e:=errors.New("rpcHandler could not be cast to : RPC_Other")
		    noDevErr(e)
		    RPC_ResponseHandler.HandelError(e)
		    return
		}*/

		switch splits[1] {
		case "Echo": //each pb_service_method
			load := &PB_OtherParam_Echo{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.Echo(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Other.Echo",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_OtherResponse_Echo",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_OtherResponse_Echo",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_OtherResponse_Echo","RPC_Other.Echo",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
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

		switch splits[1] {
		case "GetGeneralUpdates": //each pb_service_method
			load := &PB_SyncParam_GetGeneralUpdates{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetGeneralUpdates(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Sync.GetGeneralUpdates",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SyncResponse_GetGeneralUpdates",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetGeneralUpdates",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetGeneralUpdates","RPC_Sync.GetGeneralUpdates",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetNotifyUpdates": //each pb_service_method
			load := &PB_SyncParam_GetNotifyUpdates{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetNotifyUpdates(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Sync.GetNotifyUpdates",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SyncResponse_GetNotifyUpdates",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetNotifyUpdates",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_GetNotifyUpdates","RPC_Sync.GetNotifyUpdates",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetLastSyncDirectUpdateId": //each pb_service_method
			load := &PB_SyncParam_SetLastSyncDirectUpdateId{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetLastSyncDirectUpdateId(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Sync.SetLastSyncDirectUpdateId",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SyncResponse_SetLastSyncDirectUpdateId",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncDirectUpdateId",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncDirectUpdateId","RPC_Sync.SetLastSyncDirectUpdateId",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetLastSyncGeneralUpdateId": //each pb_service_method
			load := &PB_SyncParam_SetLastSyncGeneralUpdateId{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetLastSyncGeneralUpdateId(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Sync.SetLastSyncGeneralUpdateId",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SyncResponse_SetLastSyncGeneralUpdateId",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncGeneralUpdateId",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncGeneralUpdateId","RPC_Sync.SetLastSyncGeneralUpdateId",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetLastSyncNotifyUpdateId": //each pb_service_method
			load := &PB_SyncParam_SetLastSyncNotifyUpdateId{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetLastSyncNotifyUpdateId(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Sync.SetLastSyncNotifyUpdateId",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SyncResponse_SetLastSyncNotifyUpdateId",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncNotifyUpdateId",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SyncResponse_SetLastSyncNotifyUpdateId","RPC_Sync.SetLastSyncNotifyUpdateId",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
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

		switch splits[1] {
		case "BlockUser": //each pb_service_method
			load := &PB_UserParam_BlockUser{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.BlockUser(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_UserOffline.BlockUser",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserOfflineResponse_BlockUser",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_BlockUser",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_BlockUser","RPC_UserOffline.BlockUser",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "UnBlockUser": //each pb_service_method
			load := &PB_UserParam_UnBlockUser{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UnBlockUser(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_UserOffline.UnBlockUser",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserOfflineResponse_UnBlockUser",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UnBlockUser",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UnBlockUser","RPC_UserOffline.UnBlockUser",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "UpdateAbout": //each pb_service_method
			load := &PB_UserParam_UpdateAbout{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UpdateAbout(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_UserOffline.UpdateAbout",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserOfflineResponse_UpdateAbout",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UpdateAbout",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UpdateAbout","RPC_UserOffline.UpdateAbout",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "UpdateUserName": //each pb_service_method
			load := &PB_UserParam_UpdateUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UpdateUserName(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_UserOffline.UpdateUserName",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserOfflineResponse_UpdateUserName",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UpdateUserName",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_UpdateUserName","RPC_UserOffline.UpdateUserName",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "ChangePrivacy": //each pb_service_method
			load := &PB_UserParam_ChangePrivacy{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ChangePrivacy(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_UserOffline.ChangePrivacy",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponseOffline_ChangePrivacy",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponseOffline_ChangePrivacy",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponseOffline_ChangePrivacy","RPC_UserOffline.ChangePrivacy",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "ChangeAvatar": //each pb_service_method
			load := &PB_UserParam_ChangeAvatar{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ChangeAvatar(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_UserOffline.ChangeAvatar",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserOfflineResponse_ChangeAvatar",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_ChangeAvatar",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserOfflineResponse_ChangeAvatar","RPC_UserOffline.ChangeAvatar",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
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

		switch splits[1] {
		case "CheckUserName": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.CheckUserName(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.CheckUserName",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName","RPC_User.CheckUserName",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetBlockedList": //each pb_service_method
			load := &PB_UserParam_BlockedList{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetBlockedList(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.GetBlockedList",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_BlockedList",
						ResponseData:    res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_BlockedList",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_BlockedList","RPC_User.GetBlockedList",cmd, params , load)
					RPC_ResponseHandler.HandleOfflineResult(out)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	default:
		noDevErr(errors.New("rpc dosent exisit for: " + cmd.Command))
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



 RPC_Other.Echo



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
