package models

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"ms/sun/config"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

//var RpcAll = rpcAll{}
var RpcAll = x.RPC_AllHandlersInteract{}

func init() {
	//RpcAll.RPC_Msg = rpcMsg(0)
	RpcAll.RPC_Sync = rpcSync(0)
	RpcAll.RPC_Chat = rpcChat(0)
	RpcAll.RPC_Other = rpcOther(74)
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

var rpcResHand = rpcResHandeler(7)

func init() {
	//x.RPC_ResponseHandler = rpcResHand
}

func (rpcResHandeler) HandleOfflineResult(resOut x.RpcResponseOutput) {
    //fmt.Println("HandleOfflineResult",resOut)
	//fmt.Println("implement me", i, c, p)
	resOfRpcFunc, ok := resOut.ResponseData.(proto.Message)
	if ok {
		data, err := proto.Marshal(resOfRpcFunc)

		//todo: if response has error we must call the clint with and error: like: SERVER_ERR
		if err != nil {
            if config.IS_DEBUG {
                panic("can nto Marshal pb in HandleOfflineResult")
            }
			return
		}
		resToClient := &x.PB_ResponseToClient{
			ClientCallId: int64(resOut.CommandToServer.ClientCallId),
			PBClass:      resOut.PBClassName,
			RpcFullName:  resOut.RpcName,
			Data:         data,
		}

		if config.IS_DEBUG {
			//logRpc.Println("debuging loggin " + RpcName)
            //fmt.Println("implement me")
			param2, _ := resOut.RpcParamPassed.(proto.Message)
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

			logRpc.Printf(`"%s - %s"`, resOut.RpcName, t.Format("3:04:04"))
			logRpc.Println("Param = ", helper.ToJsonPerety2(param2))
			logRpc.Println("Result = ", helper.ToJsonPerety2(resOfRpcFunc))
			logRpc.Println("PB_ResponseToClient = ", helper.ToJsonPerety2(resToClient))
			logRpc.Println("ResponseDataSize = ", helper.ToJsonPerety2(len(resToClient.Data)))
			logRpc.Println(s)

		}

		//wsDebugLog(fmt.Sprintf("%s %s", "HandleOfflineResult: "+PBClass+" ", i))
		cmd := NewPB_CommandToClient_WithData("PB_ResponseToClient", resToClient)
		AllPipesMap.SendToUser(resOut.UserParam.GetUserId(), cmd)
		//_ = cmd
	}else if config.IS_DEBUG{
	    panic("rpc HandleOfflineResult response is not of type PB - ants must wroks wrong!")
    }

}

func (rpcResHandeler) IsUserOnlineResult(i interface{}, erro error) {
	fmt.Println("zzzzzzzzzzzzzzzzz", i, erro)
}

func (rpcResHandeler) HandelError(erro error) {
	fmt.Println("=============== Error in Rpc:: (rpcResHandeler) HandelError(erro error) ", erro)
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

