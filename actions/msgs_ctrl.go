package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"ms/sun/commands"
	"ms/sun/models"
	"ms/sun/sync"
	"net/http"
	"os"
	"strings"
	"time"
)

//TODO: add session handlerer/checker
func MsgUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1e6)

	print("called MsgUpload: ", r)

	upladedFile, fd, err := r.FormFile("file")
	message := r.Form.Get("message")
	//f ,fd ,err :=r.MultipartForm.File["file"]
	if message == "" {
		return
	}

	if err != nil {
		print(err.Error())
		return
	}
	defer upladedFile.Close()

	msgStr := r.Form.Get("message")

	var msgRecived models.MessagesTableFromClient
	err = json.Unmarshal([]byte(msgStr), &msgRecived)
	if err != nil || msgStr == "" {
		return
	}
	meUserId := msgRecived.UserId
	t := time.Now()
	//dirName := "./upload/"+t.Year()+
	dirName := fmt.Sprintf("./upload/msgs/%v/%d/%v/%v/", t.Year(), t.Month(), t.Day(), t.Hour())
	up_filename := strings.Replace(fd.Filename, ":", "-", -1)       //remove 17:24:56 file format -- just for windowse
	fileName := dirName + msgRecived.MessageKey + "_" + up_filename //msg.MediaName + "." + msg.MediaExtension

	err = os.MkdirAll(dirName, 0666)
	newFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		print(2, err.Error())
		return
	}
	defer newFile.Close()
	i, err := io.Copy(newFile, upladedFile)

	toUid := models.RoomKeyToUserId(msgRecived.RoomKey)
	msgToPeer := models.CreateNewMsgRecivedForSendingToPeer(&msgRecived, toUid, msgRecived.UserId) //copy
	msgToPeer.MediaServerSrc = "http://localhost:5000/" + fileName[2:]                             //remove: "./"

	meta := models.CreateMsgRecivedToServerMetaResponse(&msgRecived)

	recivedCmd := commands.NewMsgsReceivedToServer(meta)

	sync.AllPipesMap.SendAndStoreCmdToUser(meUserId, recivedCmd)

	//send msg to peer
	cmd := commands.NewMsgsAddNew(msgToPeer)
	sync.AllPipesMap.SendAndStoreCmdToUser(toUid, cmd)

	print("Upladed media Size Write: ", i)
	if err != nil {
		print(3, err.Error())
		return
	}
}

//c,_:=r.Cookie("k")
//raw_out,err := os.OpenFile("./__MsgUpload.txt", os.O_WRONLY|os.O_CREATE, 0666)
//defer raw_out.Close()
//raw_out.Write([]byte{"asdasd"})
//if err != nil {
//    print(err.Error())
//    return
//}

//i2, err := io.Copy(raw_out,r.Body)

//msg.UserId = c.UserId
//msg.RoomKey = "u"+helper.IntToStr(c.UserId)

//send MsgReceivedToServer to this user
//meta := models.MessageSyncMeta{}
//meta.RoomKey = m.RoomKey
//meta.ByUserId = -1
//meta.MessageKey = m.MessageKey
//meta.AtTimeMs = helper.TimeNowMs()
