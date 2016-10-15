package main

import (
	"ms/sun/base"
	//"ms/sun/pipes"
	"ms/sun/models"
)

//command events that clinet invokes
func registerCmdRouters() {
	/*
		mp := base.CmdMapRouter

		//commands reciced to client
		mp["CommandsReceived"] = pipesold.CommandsReceived_Dep
		mp["CommandsReceivedToClient"] = pipesold.CommandsReceived

		//Messages flow
		mp["MsgsAddNew"] = pipesold.MsgAddNew
		mp["MsgsReceivedToServer"] = pipesold.MsgReceivedToServer
		mp["MsgsReceivedToPeer"] = pipesold.MsgReceivedToPeer
		mp["MsgsSeenByPeer"] = pipesold.MsgSeenByPeer
		mp["MsgsDeletedFromServer"] = pipesold.MsgDeletedFromServer

		mp["GetUserForTable"] = pipesold.GetUserForTable

		//play
		mp["hello"] = pipesold.HelloCmd
		mp["echo"] = pipesold.EchoCmd

		mp["AddNewMsg"] = pipesold.AddNewMsg

		mp["EchoRes"] = pipesold.EchoRes
	*/

	mp2 := base.CallMapRouter
	mp2["echo"] = models.EchoCmd
	mp2["Echo"] = models.EchoCmd

	mp2["MsgsAddOne"] = models.CallReceive_MsgsAddOne
	mp2["MsgsAddMany"] = models.CallReceive_MsgsAddMany
	mp2["MsgsSeenMany"] = models.CallRecive_MsgSeenByPeer

}

/*
//commands reciced to client
mp["CommandsReceived"] = cmd.CommandsReceived

//Messages flow
mp["MsgsAddNew"] = cmd.MsgAddNew
mp["MsgsReceivedToServer"] = cmd.MsgReceivedToServer
mp["MsgsReceivedToPeer"] = cmd.MsgReceivedToPeer
mp["MsgsSeenByPeer"] = cmd.MsgSeenByPeer
mp["MsgsDeletedFromServer"] = cmd.MsgDeletedFromServer

mp["GetUserForTable"] = cmd.GetUserForTable

//play
mp["hello"] = cmd.HelloCmd
mp["echo"] = cmd.EchoCmd

mp["AddNewMsg"] = cmd.AddNewMsg

mp["EchoRes"] = cmd.EchoRes
*/
