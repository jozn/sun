package models

import (
	"ms/sun/models/x"
)

var CallMapRouter = make(map[string]func(*x.PB_CommandToServer, *UserDevicePipe))

//command events that clinet invokes
func init() {
	registerCmdRouters()
}

func registerCmdRouters() {
	mp := CallMapRouter //reference
	mp["echo"] = EchoCmd
	mp["Echo"] = EchoCmd

	//mp["MsgsAddOne"] = CallReceive_MsgsAddOne
	mp["MsgsAddMany"] = CallReceive_MsgsAddMany
	mp["MsgsSeenMany"] = CallRecive_MsgSeenByPeer

}
