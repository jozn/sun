package commands

import (
    "ms/sun/base"
    "ms/sun/constants"
    "ms/sun/models"
)

func NewMsgsReceivedToServer(meta *models.MessageSyncMeta) *base.Command {
    recivedCmd := base.NewCommand(constants.MsgsReceivedToServer)
    recivedCmd.AddSliceData(meta)
    recivedCmd.MakeDataReady()
    return recivedCmd
}

func NewMsgsAddNew(msg *models.MessagesTableFromClient) *base.Command {
    cmd := base.NewCommand(constants.MsgsAddNew)
    cmd.AddSliceData(msg)
    cmd.MakeDataReady()

    return cmd
}

