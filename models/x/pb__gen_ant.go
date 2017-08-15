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
	HandleOfflineResult(dataPB interface{}, PBClass string, cmd PB_CommandToServer, p RPC_UserParam)
	IsUserOnlineResult(interface{}, error)
	HandelError(error)
}

var RPC_ResponseHandler RPC_ResponseHandlerInterface

//note: rpc methods cant have equal name they must be different even in different rpc services
type RPC_AllHandlersInteract interface {
	RPC_MessageReq
	RPC_MessageReqOffline
	RPC_Auth
	RPC_Msg
	RPC_UserOffline
	RPC_User
}

/////////////// Interfaces ////////////////

type RPC_MessageReq interface {
	GetLastChnagesForRoom(i *PB_ReqLastChangesForTheRoom, p RPC_UserParam) (*PB_ResponseLastChangesForTheRoom, error)
}

type RPC_MessageReqOffline interface {
	SetLastSeen(i *PB_RequestSetLastSeenMessages, p RPC_UserParam) (*PB_ResponseSetLastSeenMessages, error)
}

type RPC_Auth interface {
	CheckPhone(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error)
	SendCode(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error)
	SendCodeToSms(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error)
	SendCodeToTelgram(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error)
	SingUp(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error)
	SingIn(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error)
	LogOut(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error)
}

type RPC_Msg interface {
	AddNewTextMessage(i *PB_MsgParam_AddNewTextMessage, p RPC_UserParam) (*PB_MsgResponse_AddNewTextMessage, error)
	SetRoomActionDoing(i *PB_MsgParam_SetRoomActionDoing, p RPC_UserParam) (*PB_MsgResponse_SetRoomActionDoing, error)
	GetMessagesByIds(i *PB_MsgParam_GetMessagesByIds, p RPC_UserParam) (*PB_MsgResponse_GetMessagesByIds, error)
	GetMessagesHistory(i *PB_MsgParam_GetMessagesHistory, p RPC_UserParam) (*PB_MsgResponse_GetMessagesHistory, error)
	SetMessagesRangeAsSeen(i *PB_MsgParam_SetChatMessagesRangeAsSeen, p RPC_UserParam) (*PB_MsgResponse_SetChatMessagesRangeAsSeen, error)
	DeleteChatHistory(i *PB_MsgParam_DeleteChatHistory, p RPC_UserParam) (*PB_MsgResponse_DeleteChatHistory, error)
	DeleteMessagesByIds(i *PB_MsgParam_DeleteMessagesByIds, p RPC_UserParam) (*PB_MsgResponse_DeleteMessagesByIds, error)
	SetMessagesAsReceived(i *PB_MsgParam_SetMessagesAsReceived, p RPC_UserParam) (*PB_MsgResponse_SetMessagesAsReceived, error)
	ForwardMessages(i *PB_MsgParam_ForwardMessages, p RPC_UserParam) (*PB_MsgResponse_ForwardMessages, error)
	EditMessage(i *PB_MsgParam_EditMessage, p RPC_UserParam) (*PB_MsgResponse_EditMessage, error)
	BroadcastNewMessage(i *PB_MsgParam_BroadcastNewMessage, p RPC_UserParam) (*PB_MsgResponse_BroadcastNewMessage, error)
	Echo(i *PB_MsgParam_Echo, p RPC_UserParam) (*PB_MsgResponse_PB_MsgParam_Echo, error)
}

type RPC_UserOffline interface {
	BlockUser(i *PB_UserParam_BlockUser, p RPC_UserParam) (*PB_UserOfflineResponse_BlockUser, error)
	UnBlockUser(i *PB_UserParam_UnBlockUser, p RPC_UserParam) (*PB_UserOfflineResponse_UnBlockUser, error)
	UpdateAbout(i *PB_UserParam_UpdateAbout, p RPC_UserParam) (*PB_UserOfflineResponse_UpdateAbout, error)
	UpdateUserName(i *PB_UserParam_UpdateUserName, p RPC_UserParam) (*PB_UserOfflineResponse_UpdateUserName, error)
	ChangePrivacy(i *PB_UserParam_ChangePrivacy, p RPC_UserParam) (*PB_UserResponseOffline_ChangePrivacy, error)
	ChangeAvatar(i *PB_UserParam_ChangeAvatar, p RPC_UserParam) (*PB_UserOfflineResponse_ChangeAvatar, error)
}

type RPC_User interface {
	CheckUserName(i *PB_UserParam_CheckUserName, p RPC_UserParam) (*PB_UserResponse_CheckUserName, error)
	GetBlockedList(i *PB_UserParam_BlockedList, p RPC_UserParam) (*PB_UserResponse_BlockedList, error)
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

	case "RPC_MessageReq":

		rpc, ok := rpcHandler.(RPC_MessageReq)
		if !ok {
			e := errors.New("rpcHandler could not be cast to : RPC_MessageReq")
			noDevErr(e)
			RPC_ResponseHandler.HandelError(e)
			return
		}

		switch splits[1] {
		case "GetLastChnagesForRoom": //each pb_service_method
			load := &PB_ReqLastChangesForTheRoom{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetLastChnagesForRoom(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_ResponseLastChangesForTheRoom", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_MessageReqOffline":

		rpc, ok := rpcHandler.(RPC_MessageReqOffline)
		if !ok {
			e := errors.New("rpcHandler could not be cast to : RPC_MessageReqOffline")
			noDevErr(e)
			RPC_ResponseHandler.HandelError(e)
			return
		}

		switch splits[1] {
		case "SetLastSeen": //each pb_service_method
			load := &PB_RequestSetLastSeenMessages{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetLastSeen(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_ResponseSetLastSeenMessages", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_Auth":

		rpc, ok := rpcHandler.(RPC_Auth)
		if !ok {
			e := errors.New("rpcHandler could not be cast to : RPC_Auth")
			noDevErr(e)
			RPC_ResponseHandler.HandelError(e)
			return
		}

		switch splits[1] {
		case "CheckPhone": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.CheckPhone(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_CheckUserName2", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_CheckUserName2", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_CheckUserName2", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_CheckUserName2", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_CheckUserName2", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_CheckUserName2", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_CheckUserName2", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_Msg":

		rpc, ok := rpcHandler.(RPC_Msg)
		if !ok {
			e := errors.New("rpcHandler could not be cast to : RPC_Msg")
			noDevErr(e)
			RPC_ResponseHandler.HandelError(e)
			return
		}

		switch splits[1] {
		case "AddNewTextMessage": //each pb_service_method
			load := &PB_MsgParam_AddNewTextMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.AddNewTextMessage(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_AddNewTextMessage", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetRoomActionDoing": //each pb_service_method
			load := &PB_MsgParam_SetRoomActionDoing{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetRoomActionDoing(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_SetRoomActionDoing", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetMessagesByIds": //each pb_service_method
			load := &PB_MsgParam_GetMessagesByIds{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetMessagesByIds(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_GetMessagesByIds", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetMessagesHistory": //each pb_service_method
			load := &PB_MsgParam_GetMessagesHistory{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetMessagesHistory(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_GetMessagesHistory", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetMessagesRangeAsSeen": //each pb_service_method
			load := &PB_MsgParam_SetChatMessagesRangeAsSeen{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetMessagesRangeAsSeen(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_SetChatMessagesRangeAsSeen", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "DeleteChatHistory": //each pb_service_method
			load := &PB_MsgParam_DeleteChatHistory{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteChatHistory(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_DeleteChatHistory", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "DeleteMessagesByIds": //each pb_service_method
			load := &PB_MsgParam_DeleteMessagesByIds{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteMessagesByIds(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_DeleteMessagesByIds", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetMessagesAsReceived": //each pb_service_method
			load := &PB_MsgParam_SetMessagesAsReceived{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetMessagesAsReceived(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_SetMessagesAsReceived", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "ForwardMessages": //each pb_service_method
			load := &PB_MsgParam_ForwardMessages{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ForwardMessages(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_ForwardMessages", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "EditMessage": //each pb_service_method
			load := &PB_MsgParam_EditMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.EditMessage(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_EditMessage", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "BroadcastNewMessage": //each pb_service_method
			load := &PB_MsgParam_BroadcastNewMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.BroadcastNewMessage(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_BroadcastNewMessage", cmd, params)
				} else {
					RPC_ResponseHandler.HandelError(err)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "Echo": //each pb_service_method
			load := &PB_MsgParam_Echo{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.Echo(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_MsgResponse_PB_MsgParam_Echo", cmd, params)
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

		rpc, ok := rpcHandler.(RPC_UserOffline)
		if !ok {
			e := errors.New("rpcHandler could not be cast to : RPC_UserOffline")
			noDevErr(e)
			RPC_ResponseHandler.HandelError(e)
			return
		}

		switch splits[1] {
		case "BlockUser": //each pb_service_method
			load := &PB_UserParam_BlockUser{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.BlockUser(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserOfflineResponse_BlockUser", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserOfflineResponse_UnBlockUser", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserOfflineResponse_UpdateAbout", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserOfflineResponse_UpdateUserName", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponseOffline_ChangePrivacy", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserOfflineResponse_ChangeAvatar", cmd, params)
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

		rpc, ok := rpcHandler.(RPC_User)
		if !ok {
			e := errors.New("rpcHandler could not be cast to : RPC_User")
			noDevErr(e)
			RPC_ResponseHandler.HandelError(e)
			return
		}

		switch splits[1] {
		case "CheckUserName": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.CheckUserName(load, params)
				if err == nil {
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_CheckUserName", cmd, params)
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
					RPC_ResponseHandler.HandleOfflineResult(res, "PB_UserResponse_BlockedList", cmd, params)
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


 RPC_MessageReq.GetLastChnagesForRoom



 RPC_MessageReqOffline.SetLastSeen



 RPC_Auth.CheckPhone
 RPC_Auth.SendCode
 RPC_Auth.SendCodeToSms
 RPC_Auth.SendCodeToTelgram
 RPC_Auth.SingUp
 RPC_Auth.SingIn
 RPC_Auth.LogOut



 RPC_Msg.AddNewTextMessage
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
 RPC_Msg.Echo



 RPC_UserOffline.BlockUser
 RPC_UserOffline.UnBlockUser
 RPC_UserOffline.UpdateAbout
 RPC_UserOffline.UpdateUserName
 RPC_UserOffline.ChangePrivacy
 RPC_UserOffline.ChangeAvatar



 RPC_User.CheckUserName
 RPC_User.GetBlockedList


*/
