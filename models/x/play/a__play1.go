package play

import (
	"strings"

	"github.com/golang/protobuf/proto"
)

type RPC_UserParam interface {
	GetUserId() int
	IsUser() bool
}

type RPC_ResponseHandlerInterface interface {
	HandleOfflineResult(interface{}, error) int
	IsUserOnlineResult(interface{}, error) bool
}

var RPC_ResponseHandler RPC_ResponseHandlerInterface

type RPC_MessageOfflineSample interface {
	GetUserId2() int
	IsUser1() bool
	AddtextMessage(*PB_MsgEvent) (*PB_Push, error)
}

type RFC_AllHandlersInteract interface {
	RPC_MessageOfflineSample
	RPC_UserParam
}

func HandleWsRpcsSample(cmd PB_CommandToClient, params RPC_UserParam, rpcHandler RFC_AllHandlersInteract) {

	splits := strings.Split(cmd.Command, ".")

	if len(splits) != 2 {
		return
	}

	switch splits[0] {
	case "RPC_MessageOfflineSample": //each pb_service
		rpc, ok := rpcHandler.(RPC_MessageOfflineSample)
		if !ok {
			return
		}
		switch cmd.Command {
		case "AddtextMessage": //each pb_service_method
			load := &PB_MsgEvent{}
			err := proto.Unmarshal(cmd.Data, load)
			if err != nil {
				res, err := rpc.AddtextMessage(load)
				if err != nil && res != nil {
					handleOfflineSample(res)
					handleOnlineSample(res)
				}
			}
		}
	}

}

func handleOfflineSample(err interface{}) {

}
func handleOnlineSample(err interface{}) {

}
