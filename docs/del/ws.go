package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	// "log"
	"encoding/json"
	"time"
	"os"
	"ms/sun/helper"
)

type WSReq struct {
	SessionUid string
	Command    string
	Params     map[string]interface{}
	RequestId  string
}

type WSRes struct {
	// BNCP
	Status    string
	Payload   interface{}
	ResTime   int64
	Commands  []WSResCommand
	RequestId string
}

type WSResCommand struct {
	Command string
	Params  map[string]interface{}
}

var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 128,
	WriteBufferSize: 1024 * 128,
}

//
func serveWs2(w http.ResponseWriter, r *http.Request) {
	//defer recover()
	err := r.ParseForm()
	noErr(err)
	session := r.Form.Get("SessionUid")
	//TODO add session check functiality
	if false {
		e(session)
		return //not Autorized user
	}

	ws, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WSUpgrader err", err)
		return
	}

	//function: wsWorker
	go func() {
		//TODO: extarct wsWorker to its func and then in panic-recover start agin another wsWorker
		//if we panic somehow? in writing json to writer, just close ws
		defer func() {
			if r := recover(); r != nil {
				ws.Close()
			}
		}()

		resChan := make(chan WSRes, 2)
		isOpen := true
		//Routine of clinet command request
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in ws messaging clinet request", r)
				}
			}()
			//FIXME:this is sequence message proccessing do we need multi process
			for isOpen == true {
				var req WSReq
				messageType, bytes, _ := ws.ReadMessage() //blocking
				if messageType == websocket.CloseMessage {
					close(resChan)
					isOpen = false
					return
				}
				//TODO: another go fn here for many command proccessing
				if messageType == websocket.TextMessage {
					t1 := time.Now().UnixNano()
					json.Unmarshal(bytes, &req)
					res := handleWSCommand(req)
					devPrintn("WS: Command: ", req.Command, " Res: ", res)
					res.ResTime = (time.Now().UnixNano() - t1) / 1e6 //to miliscond
					resChan <- res
				}
				if messageType == websocket.BinaryMessage {
					f,_ :=os.Create("./ws_"+ helper.RandString(8))
					wd,_:=os.Getwd()
					fmt.Println("ws: "+" wd:" +wd + "binary:  ",bytes)
					f.Write(bytes)
					f.Close()
				}
			}
		}()

		for r := range resChan {
			ws.WriteJSON(r)
		}
		//after closing chanel
		ws.Close()
		// for {
		// 	<-time.After(time.Second * 1)
		// 	ws.WriteJSON(User{UserName: "as"})

		// }
	}()
}

func handleWSCommand(req WSReq) (res WSRes) {
	//defer recover()
	res = WSRes{}
	defer func() { //if command fails
		if r := recover(); r != nil {
			fmt.Println("Recovered in ws proccesing Command"+req.Command, r)
			res.Status = "ERROR"
			//return res
		}
	}()
	res.RequestId = req.RequestId
	// res.RequestId = req.Command
	res.Status = "OK"
	fn := WSCommandMap[req.Command]

	action := WSAction{Req: &req, Res: &res}
	if fn != nil {
		fn(&action)
	} else {
		wsErrorCommand(&action)
	}
	return res
}

// WSCommandMap := make( map[string]func(*WSAction) )
var WSCommandMap map[string]func(*WSAction)

type WSAction struct {
	Req *WSReq
	Res *WSRes
	//holder [string]func(*WSAction)
}
