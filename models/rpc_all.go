package models

import "ms/sun/models/x"

type rpcAll struct {
	//x.RPC_AllHandlersInteract
	x.EmptyRPC_RPC_MessageReq
	x.EmptyRPC_RPC_MessageReqOffline
	x.EmptyRPC_RPC_Auth
	//x.EmptyRPC_RPC_Msg
	x.EmptyRPC_RPC_UserOffline
	x.EmptyRPC_RPC_User

	/// ouer imple
	rpcMsg
}

var RpcAll = rpcAll{}

func init() {
	RpcAll.rpcMsg = rpcMsg(0)
}

///////////////////////////////
type RPC_UserParam_Imple struct {
	userId int
}

func (s RPC_UserParam_Imple) GetUserId() int {
	return s.userId
}

func (s RPC_UserParam_Imple) IsUser() bool {
	return s.userId > 0
}
