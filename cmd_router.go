package main

import (
	"ms/sun/base"
	//"ms/sun/pipes"
	"ms/sun/models"
)

//command events that clinet invokes
func registerCmdRouters() {

	mp2 := base.CallMapRouter
	mp2["echo"] = models.EchoCmd
	mp2["Echo"] = models.EchoCmd

	mp2["MsgsAddOne"] = models.CallReceive_MsgsAddOne
	mp2["MsgsAddMany"] = models.CallReceive_MsgsAddMany
	mp2["MsgsSeenMany"] = models.CallRecive_MsgSeenByPeer

}
