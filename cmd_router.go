package main
import (
    "ms/sun/base"
    "ms/sun/cmd"
)

//command events that clinet invokes
func registerCmdRouters() {
    mp:= base.CmdMapRouter

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

}

