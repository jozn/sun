package main

import (
	"ms/sun/base"
	"ms/sun/pipes"
)

//command events that clinet invokes
func registerCmdRouters() {
	mp := base.CmdMapRouter

	//commands reciced to client
	mp["CommandsReceived"] = pipes.CommandsReceived

	//Messages flow
	mp["MsgsAddNew"] = pipes.MsgAddNew
	mp["MsgsReceivedToServer"] = pipes.MsgReceivedToServer
	mp["MsgsReceivedToPeer"] = pipes.MsgReceivedToPeer
	mp["MsgsSeenByPeer"] = pipes.MsgSeenByPeer
	mp["MsgsDeletedFromServer"] = pipes.MsgDeletedFromServer

	mp["GetUserForTable"] = pipes.GetUserForTable

	//play
	mp["hello"] = pipes.HelloCmd
	mp["echo"] = pipes.EchoCmd

	mp["AddNewMsg"] = pipes.AddNewMsg

	mp["EchoRes"] = pipes.EchoRes

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
