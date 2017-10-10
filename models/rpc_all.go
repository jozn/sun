package models

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"ms/sun/config"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

type rpcAll struct {
	//x.RPC_AllHandlersInteract
	x.EmptyRPC_RPC_MessageReq
	x.EmptyRPC_RPC_MessageReqOffline
	x.EmptyRPC_RPC_Auth
	//x.EmptyRPC_RPC_Msg
	x.EmptyRPC_RPC_UserOffline
	x.EmptyRPC_RPC_User
	//x.EmptyRPC_RpcMsgs

	/// ouer imple
	rpcMsg
	rpcSync
}

var RpcAll = rpcAll{}

func init() {
	RpcAll.rpcMsg = rpcMsg(0)
	RpcAll.rpcSync = rpcSync(0)
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

func (rpcResHandeler) HandleOfflineResult(i interface{}, PBClass string, RpcName string, c x.PB_CommandToServer, p x.RPC_UserParam, paramParsed interface{}) {
	//fmt.Println("implement me", i, c, p)
	resOfRpcFunc, ok := i.(proto.Message)
	if ok {
		data, err := proto.Marshal(resOfRpcFunc)

		//todo: if response has error we must call the clint with and error: like: SERVER_ERR
		if err != nil {
			return
		}
		resToClient := &x.PB_ResponseToClient{
			ClientCallId: int64(c.ClientCallId),
			PBClass:      PBClass,
			RpcFullName:  RpcName,
			Data:         data,
		}

		if config.IS_DEBUG {
			//logRpc.Println("debuging loggin " + RpcName)
			param2, _ := paramParsed.(proto.Message)
			t := time.Now()
			s := "//======================================================================================================================="
			//s = "////////////////////////////////////////////////////////////////////////////////////"
			//logRpc.Println(s)
			/*oT :=
							`"%s - %s"
			Param = %s
			Result = %s
			`*/
			//logRpc.Printf(oT, RpcName, t.Format("3:04:04"), helper.ToJsonPerety2(param2), helper.ToJsonPerety2(resOfRpcFunc))
			//logRpc.Println(s)

			logRpc.Printf(`"%s - %s"`, RpcName, t.Format("3:04:04"))
			logRpc.Println("Param = ", helper.ToJsonPerety2(param2))
			logRpc.Println("Result = ", helper.ToJsonPerety2(resOfRpcFunc))
			logRpc.Println("PB_ResponseToClient = ", helper.ToJsonPerety2(resToClient))
			logRpc.Println("ResponseDataSize = ", helper.ToJsonPerety2(len(resToClient.Data)))
			logRpc.Println(s)

		}

		//wsDebugLog(fmt.Sprintf("%s %s", "HandleOfflineResult: "+PBClass+" ", i))
		cmd := NewPB_CommandToClient_WithData("PB_ResponseToClient", resToClient)
		AllPipesMap.SendToUser(p.GetUserId(), cmd)
		//_ = cmd
	}

}

func (rpcResHandeler) IsUserOnlineResult(i interface{}, erro error) {
	fmt.Println("zzzzzzzzzzzzzzzzz", i, erro)
}

func (rpcResHandeler) HandelError(erro error) {
	fmt.Println("===============", erro)
}

func PushToUserLiveData(UserId int, pbAllLivePush x.PB_AllLivePushes) {
	if config.IS_DEBUG {
		t := time.Now()
		s := "//======================================================================================================================="
		logLivePush.Printf(`"%s"`, t.Format("3:04:04"))
		logLivePush.Println("LivePushes = ", helper.ToJsonPerety2(pbAllLivePush))
		logLivePush.Println(s)
	}

	cmd := NewPB_CommandToClient_WithData("PB_AllLivePushes", &pbAllLivePush)
	AllPipesMap.SendToUser(UserId, cmd)
}
