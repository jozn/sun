package ctrl

import (
    "io"
    "os"
    "strings"
    "fmt"
    "time"
    "encoding/json"
    "net/http"
    "ms/sun/models"
)

//TODO: add session handlerer/checker
func MsgUploadV1(w http.ResponseWriter, r *http.Request) {
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
        print(2, err.Error())
        return
    }
    meUserId := msgRecived.UserId
    t := time.Now()
    //dirName := "./upload/"+t.Year()+
    dirName := fmt.Sprintf("./upload/msgs/%v/%d/%v/%v/", t.Year(), t.Month(), t.Day(), t.Hour())
    up_filename := strings.Replace(fd.Filename, ":", "-", -1)       //remove 17:24:56 file format -- just for windowse
    fileName := dirName + msgRecived.MessageKey + "_" + up_filename //msg.MediaName + "." + msg.MediaExtension

    fileName = strings.Replace(fileName,":","-",-1)//Windows dosn't accept ':'
    //fileName = strings.Replace(fileName,"@","-",-1)

    err = os.MkdirAll(dirName, 0666)
    newFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        print(2, err.Error())
        return
    }
    defer newFile.Close()
    i, err := io.Copy(newFile, upladedFile)

    msgToSend := msgRecived
    msgToSend.MediaServerSrc = "http://localhost:5000/" + fileName[2:]

    toUid,err := models.RoomKeyToPeerUserId(msgToSend.RoomKey,meUserId)
    if err != nil {
        print(2, err.Error())
        return
    }
    //msgToPeer := models.CreateNewMsgRecivedForSendingToPeer(&msgRecived, toUid, msgRecived.UserId) //copy
    msgToSend.MediaServerSrc = "http://localhost:5000/" + fileName[2:]//remove: "./"

    models.MessageModel.SendAndStoreMessage(toUid,msgToSend)

    /*meta := models.CreateMsgRecivedToServerMetaResponse(&msgRecived)

    recivedCmd := commands.NewMsgsReceivedToServer(meta)

    sync.AllPipesMap.SendAndStoreCmdToUser(meUserId, recivedCmd)

    //send msg to peer
    cmd := commands.NewMsgsAddNew(msgToPeer)
    sync.AllPipesMap.SendAndStoreCmdToUser(toUid, cmd)*/

    print("Upladed media Size Write: ", i)
    if err != nil {
        print(3, err.Error())
        return
    }
}

