package models

import (
	"ms/sun/helper"
	"net/http"

	"github.com/gorilla/websocket"
    "os"
    "fmt"
    "ms/sun/config"
    "ms/sun/models/x"
    "github.com/golang/protobuf/proto"
    "encoding/base64"
)

func ServeHttpWsPB(w http.ResponseWriter, r *http.Request) {
	//defer recover()
	err := r.ParseForm()
	//noErr(err)
	session := r.Form.Get("SessionUid")
	_ = session

	uid := 6
	uid = helper.StrToInt(r.Form.Get("UserId"), 6)
	//TODO add session check functiality
	if false {
		//e(session)
		return //not Autorized user
	}

	var WSUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024 * 128,
		WriteBufferSize: 1024 * 128,
	}

	ws, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		helper.Debug("WSUpgrader err", err)
		return
	}

	//TODO get userId
	//AllPipesMap.ServeNewHttpWsForUser(uid, ws)

    go func() {
        defer func() {
            if r := recover(); r != nil {
                helper.Debug("Recovered in ws messaging clinet request", r ,uid)
            }
        }()
        for {
            messageType, bytes, err := ws.ReadMessage() //blocking

            helper.Debug("messageType: ", " ::", messageType, string(bytes))
            //wsPbDebugLog("messageType: ", " ::", messageType, string(bytes))
            if messageType == websocket.CloseMessage || err != nil {
                helper.Debugf("closeing pip for userId: %v , messageType:%v , err: %v")
                return
            }

            if messageType == websocket.TextMessage {

            }

            //b :=`CI+x+IOQkeLgFBIHTVNHX0FERBqoAQoOMTBAOmdNRHBUWHVoeEgSBXA2XzEwGAooCjJsaWQ6IDE0IPCfkanigI3wn5Gp4oCN8J+RpyDZgtmI24zigIzYqtix24zZhiDZhtmH2YbaryAj2qnZiNqG2qnigIzYqtix24zZhiDwn5G18J+PvCAg8J+Rt/Cfj7wg2LTYr9mHINiq2YjYs9i3UKyppKrBK1iQuKSqwStg4fuFg6351N8UgAGs9O/IBQ==`
            if messageType == websocket.BinaryMessage {


                    /*c := &x.PB_CommandToServer{
                        Command: proto.String("MSG_ADD"),
                        CallId: proto.Int64(256841654),
                    }*/
                /*

                c := &x.PB_CommandToServer{
                    Command:("MSG_ADD"),
                    CallId: (256841654),
                }
                    c.String()
                    m := &x.PB_Message{Text:"ساتمنن حخخح شینشسینشس",IsByMe: 2, UserId: 256,MessageTypeId:int32(i)}
                    c.Data,err  = proto.Marshal(m)
                    if err !=nil{
                        wsDebugLog("ERR IN  proto.Marshal(m) : ", err)
                    }
                    f1 , _ := os.Create("./logs/msg/" + strconv.Itoa(i) + "_" )
                    od, err := proto.Marshal(c)
                    if err !=nil{
                        wsDebugLog("ERR IN  proto.Marshal(c) : ", err)
                    }
                    f1.Write(od)
                    f1.Close()

                    wsPbDebugLog("+++ CMD: ", c.GetCommand() ," ", c.GetCallId(), " Data: ", helper.ToJson(m))
                    if rand.Intn(3) == 2 {
                        bytes = od
                    }

                    c3 := &x.PB_CommandToServer{}
                    if err = proto.Unmarshal(od,c3 ); err == nil{
                        wsPbDebugLog("$$$$$ CMD: ", c3.GetCommand() ," ", c3.GetCallId())

                        m3 := &x.PB_Message{}
                        if err = proto.Unmarshal(c3.GetData(), m3 ); err == nil{
                            wsPbDebugLog(">>>>>>  msg: ", m3.GetText() ," ", m3.GetIsByMe(),m3)
                        }else {
                            wsPbDebugLog("$$$$$ ERR: ", err)
                        }

                    }else {
                        wsPbDebugLog("$$$$$ ERR: ", err)
                    }

                    bytes = od


                */

               /* i++
                f , _ := os.Create("./logs/msg/" + strconv.Itoa(i) )
                f.Write(bytes)
                f.Close()

                if rand.Intn(3) == 2 {
                    bb ,err:= base64.StdEncoding.DecodeString(b)
                    if err == nil{
                        bytes = bb
                    }else {
                        wsDebugLog("base: ", err)
                    }
                }*/
                cmd := &x.PB_CommandToServer{}
                err = proto.Unmarshal(bytes,cmd)

                if err !=nil{
                    wsDebugLog("ERR", err)
                }

                msg:= &x.PB_Message{}
                err = proto.Unmarshal(cmd.GetData(),msg)

                if err !=nil{
                    wsDebugLog("ERR",err)
                }

                b64 := base64.StdEncoding.EncodeToString(bytes)
                _ = b64
                //wsPbDebugLog("D: ", b64)
                wsPbDebugLog("CMD: "+ cmd.GetCommand() ," ", cmd.GetCallId(), " Data: ", helper.ToJson(msg))

            }

        }
    }()

}

var _wsPbLogFile *os.File
var wsPbDebugLog = func(strings ...interface{}) {
    if config.IS_DEBUG {
        if _wsPbLogFile == nil {
            _wsPbLogFile, _ = os.OpenFile("./logs/ws_pb_"+helper.IntToStr(helper.TimeNow())+".json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        }
        _wsPbLogFile.WriteString(fmt.Sprintln(strings...))
        _wsPbLogFile.Sync()
    }
}

