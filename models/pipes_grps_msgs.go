package models

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"os"
	"strings"
	"time"
)

type _msgGrpc struct {
}

func (g _msgGrpc) UploadNewMsg(ctx context.Context, msgPb *x.PB_Message) (res *x.PB_ResRpcAddMsg, err error) {
	//TODO session handling in here

	filePb := msgPb.File

	msgBuf := newChatMsgDelayer{
		msgPB:        msgPb,
		fromUserId:   6,                 //FIXME
		toUserId:     int(msgPb.PeerId), //FIXME
		roomKey:      msgPb.RoomKey,
		hashId:       1,
		uid:          helper.RandomSeqUid(),
		msgFileRowId: -1,
	}

	if isMsgOfMediaType(msgPb.MessageTypeId) && filePb != nil {

		t := time.Now()
		dirName := fmt.Sprintf("./upload/msgs/%v/%d/%v/%v/", t.Year(), t.Month(), t.Day(), t.Hour())
		//up_filename := strings.Replace(filePb.Name, ":", "-", -1)       //remove 17:24:56 file format -- just for windowse
		up_filename := helper.RandString(10)
		fileName := dirName + msgPb.MessageKey + "_" + up_filename //msg.MediaName + "." + msg.MediaExtension

		fileName = strings.Replace(fileName, ":", "-", -1) //Windows dosn't accept ':'

		err = os.MkdirAll(dirName, 0666)
		newFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			print(2, err.Error())
			return
		}

		defer newFile.Close()

		_, err = io.Copy(newFile, bytes.NewBuffer(filePb.Data))
		if err != nil {
			print(2, err.Error())
			return
		}

		rowFile := PBConv_PB_MsgFile_toNew_MsgFile(filePb)
		rowFile.Insert(base.DB)

		res.ServerSrc = rowFile.ServerSrc

		msgBuf.msgFileRowId = rowFile.Id

	} else {
		//text types , we should never land in here use ws for texts
	}

	chanNewChatMsgsBuffer <- msgBuf

	return
}

func isMsgOfMediaType(typMsg int32) bool {
	r := false

	switch typMsg {
	case MESSAGE_TEXT:
		r = false
	case MESSAGE_CONTACT:
		r = false
	case MESSAGE_lOCCASION:
		r = false
	case MESSAGE_STICKER:
		r = false
	case MESSAGE_POST:
		r = false
	case MESSAGE_IMAGE:
		r = true
	case MESSAGE_VIDEO:
		r = true
	case MESSAGE_AUDIO:
		r = true
	case MESSAGE_FILE:
		r = true
	default:
		logProblems.Errorf("Unkonen Message type id: \n", typMsg)
	}

	return r

}
