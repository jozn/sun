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
	HandleOfflineResult(interface{}, PB_CommandToServer, RPC_UserParam)
	IsUserOnlineResult(interface{}, error)
	HandelError(error)
}

var RPC_ResponseHandler RPC_ResponseHandlerInterface

//note: rpc methods cant have equal name they must be different even in different rpc services
type RPC_AllHandlersInteract interface {
	RPC_MessageReq
	RPC_MessageReqOffline
	Greeter
	RPC_Msg
	RpcMsgs
	RPC_UserOffline
	RPC_User
	RPC_Auth
}

/////////////// Interfaces ////////////////

type RPC_MessageReq interface {
	GetLastChnagesForRoom(i *PB_ReqLastChangesForTheRoom) (*PB_ResponseLastChangesForTheRoom, error)
}

type RPC_MessageReqOffline interface {
	SetLastSeen(i *PB_RequestSetLastSeenMessages) (*PB_ResponseSetLastSeenMessages, error)
}

type Greeter interface {
	SayHello(i *PB_Message) (*PB_UserWithMe, error)
}

type RPC_Msg interface {
	AddNewTextMessage(i *PB_MsgParam_AddNewTextMessage) (*PB_MsgResponse_AddNewTextMessage, error)
	SetRoomActionDoing(i *PB_MsgParam_SetRoomActionDoing) (*PB_MsgResponse_SetRoomActionDoing, error)
	GetMessagesByIds(i *PB_MsgParam_GetMessagesByIds) (*PB_MsgResponse_GetMessagesByIds, error)
	GetMessagesHistory(i *PB_MsgParam_GetMessagesHistory) (*PB_MsgResponse_GetMessagesHistory, error)
	SetMessagesRangeAsSeen(i *PB_MsgParam_SetMessagesRangeAsSeen) (*PB_MsgResponse_SetMessagesRangeAsSeen, error)
	DeleteRoomHistory(i *PB_MsgParam_DeleteRoomHistory) (*PB_MsgResponse_DeleteRoomHistory, error)
	DeleteMessagesByIds(i *PB_MsgParam_DeleteMessagesByIds) (*PB_MsgResponse_DeleteMessagesByIds, error)
	SetMessagesAsReceived(i *PB_MsgParam_SetMessagesAsReceived) (*PB_MsgResponse_SetMessagesAsReceived, error)
	ForwardMessages(i *PB_MsgParam_ForwardMessages) (*PB_MsgResponse_ForwardMessages, error)
	EditMessage(i *PB_MsgParam_EditMessage) (*PB_MsgResponse_EditMessage, error)
	BroadcastNewMessage(i *PB_MsgParam_BroadcastNewMessage) (*PB_MsgResponse_BroadcastNewMessage, error)
}

type RpcMsgs interface {
	UploadNewMsg(i *PB_Message) (*PB_ResRpcAddMsg, error)
}

type RPC_UserOffline interface {
	BlockUser(i *PB_UserParam_BlockUser) (*PB_UserOfflineResponse_BlockUser, error)
	UnBlockUser(i *PB_UserParam_UnBlockUser) (*PB_UserOfflineResponse_UnBlockUser, error)
	UpdateAbout(i *PB_UserParam_UpdateAbout) (*PB_UserOfflineResponse_UpdateAbout, error)
	UpdateUserName(i *PB_UserParam_UpdateUserName) (*PB_UserOfflineResponse_UpdateUserName, error)
	ChangePrivacy(i *PB_UserParam_ChangePrivacy) (*PB_UserResponseOffline_ChangePrivacy, error)
	ChangeAvatar(i *PB_UserParam_ChangeAvatar) (*PB_UserOfflineResponse_ChangeAvatar, error)
}

type RPC_User interface {
	CheckUserName(i *PB_UserParam_CheckUserName) (*PB_UserResponse_CheckUserName, error)
	GetBlockedList(i *PB_UserParam_BlockedList) (*PB_UserResponse_BlockedList, error)
}

type RPC_Auth interface {
	CheckPhone(i *PB_UserParam_CheckUserName) (*PB_UserResponse_CheckUserName, error)
	SendCode(i *PB_UserParam_CheckUserName) (*PB_UserResponse_CheckUserName, error)
	SendCodeToSms(i *PB_UserParam_CheckUserName) (*PB_UserResponse_CheckUserName, error)
	SendCodeToTelgram(i *PB_UserParam_CheckUserName) (*PB_UserResponse_CheckUserName, error)
	SingUp(i *PB_UserParam_CheckUserName) (*PB_UserResponse_CheckUserName, error)
	SingIn(i *PB_UserParam_CheckUserName) (*PB_UserResponse_CheckUserName, error)
	LogOut(i *PB_UserParam_CheckUserName) (*PB_UserResponse_CheckUserName, error)
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
				res, err := rpc.GetLastChnagesForRoom(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
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
				res, err := rpc.SetLastSeen(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		}
	case "Greeter":
		rpc, ok := rpcHandler.(Greeter)
		if !ok {
			e := errors.New("rpcHandler could not be cast to : Greeter")
			noDevErr(e)
			RPC_ResponseHandler.HandelError(e)
			return
		}

		switch splits[1] {
		case "SayHello": //each pb_service_method
			load := &PB_Message{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SayHello(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
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
				res, err := rpc.AddNewTextMessage(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetRoomActionDoing": //each pb_service_method
			load := &PB_MsgParam_SetRoomActionDoing{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetRoomActionDoing(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetMessagesByIds": //each pb_service_method
			load := &PB_MsgParam_GetMessagesByIds{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetMessagesByIds(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetMessagesHistory": //each pb_service_method
			load := &PB_MsgParam_GetMessagesHistory{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetMessagesHistory(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetMessagesRangeAsSeen": //each pb_service_method
			load := &PB_MsgParam_SetMessagesRangeAsSeen{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetMessagesRangeAsSeen(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "DeleteRoomHistory": //each pb_service_method
			load := &PB_MsgParam_DeleteRoomHistory{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteRoomHistory(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "DeleteMessagesByIds": //each pb_service_method
			load := &PB_MsgParam_DeleteMessagesByIds{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteMessagesByIds(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SetMessagesAsReceived": //each pb_service_method
			load := &PB_MsgParam_SetMessagesAsReceived{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetMessagesAsReceived(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "ForwardMessages": //each pb_service_method
			load := &PB_MsgParam_ForwardMessages{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ForwardMessages(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "EditMessage": //each pb_service_method
			load := &PB_MsgParam_EditMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.EditMessage(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "BroadcastNewMessage": //each pb_service_method
			load := &PB_MsgParam_BroadcastNewMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.BroadcastNewMessage(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		}
	case "RpcMsgs":
		rpc, ok := rpcHandler.(RpcMsgs)
		if !ok {
			e := errors.New("rpcHandler could not be cast to : RpcMsgs")
			noDevErr(e)
			RPC_ResponseHandler.HandelError(e)
			return
		}

		switch splits[1] {
		case "UploadNewMsg": //each pb_service_method
			load := &PB_Message{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UploadNewMsg(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
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
				res, err := rpc.BlockUser(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "UnBlockUser": //each pb_service_method
			load := &PB_UserParam_UnBlockUser{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UnBlockUser(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "UpdateAbout": //each pb_service_method
			load := &PB_UserParam_UpdateAbout{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UpdateAbout(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "UpdateUserName": //each pb_service_method
			load := &PB_UserParam_UpdateUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UpdateUserName(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "ChangePrivacy": //each pb_service_method
			load := &PB_UserParam_ChangePrivacy{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ChangePrivacy(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "ChangeAvatar": //each pb_service_method
			load := &PB_UserParam_ChangeAvatar{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ChangeAvatar(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
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
				res, err := rpc.CheckUserName(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "GetBlockedList": //each pb_service_method
			load := &PB_UserParam_BlockedList{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetBlockedList(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
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
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.CheckPhone(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SendCode": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCode(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SendCodeToSms": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCodeToSms(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SendCodeToTelgram": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCodeToTelgram(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SingUp": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SingUp(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "SingIn": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SingIn(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		case "LogOut": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.LogOut(load)
				if err == nil {
					RPC_ResponseHandler.HandelError(err)
				} else {
					RPC_ResponseHandler.HandleOfflineResult(res, cmd, params)
				}
			} else {
				RPC_ResponseHandler.HandelError(err)
			}
		}
	default:
		noDevErr(errors.New("HandleRpcs: splic is not 2 parts"))
	}
}
