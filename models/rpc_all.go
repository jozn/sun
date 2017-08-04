package models

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"ms/sun/models/x"
)

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

//////////////////////
type rpcResHandeler int

var rpcResHand = rpcResHandeler(0)

func init() {
	x.RPC_ResponseHandler = rpcResHand
}

func (rpcResHandeler) HandleOfflineResult(i interface{}, c x.PB_CommandToServer, p x.RPC_UserParam) {
	fmt.Println("implement me", i, c, p)
	res, ok := i.(proto.Message)
	if ok {
		cmd := NewPB_CommandToClient_WithData("cmd_res", res)
		AllPipesMap.SendToUser(p.GetUserId(), cmd)
	}

}

func (rpcResHandeler) IsUserOnlineResult(i interface{}, erro error) {
	fmt.Println("zzzzzzzzzzzzzzzzz", i, erro)
}

func (rpcResHandeler) HandelError(erro error) {
	fmt.Println("===============", erro)
}
