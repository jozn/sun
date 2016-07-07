package base

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	// "log"
	"encoding/json"
	"time"
	//"ms/sun/helper"
)


type WSAction struct {
	Req    *WSReq
	Res    *WSRes
	Chanel chan WSRes
	//holder [string]func(*WSAction)
}

func (c *WSAction) SendRes() {
	c.Chanel <- *c.Res
}

func (c *WSAction) Send(res WSRes) {
	c.Chanel <- res
}

func (c *WSAction) AddCommand(com Command) {
	//if com.ToJson != nil {
	//	//com.Data =
	//}
}

var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 128,
	WriteBufferSize: 1024 * 128,
}

// WSCommandMap := make( map[string]func(*WSAction) )
var WSCommandMap map[string]func(*WSAction)


func wsErrorCommand(c *WSAction) {
	c.Res.Status = "ERROR"
}


var cnt int = 0
//dep
func ServeWs2(w http.ResponseWriter, r *http.Request) {
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
				messageType, bytes, err := ws.ReadMessage() //blocking
				cnt++
				fmt.Println("messageType: ",cnt," ::",messageType)
				//if err != nil {
				//	close(resChan)
				//	fmt.Println("err ",messageType)
				//	isOpen = false
				//	return
				//}
				if messageType == websocket.CloseMessage || err != nil {
					close(resChan)
					fmt.Println("closeding ",messageType," error: ",err)
					isOpen = false
					return
				}
				//TODO: another go fn here for many command proccessing
				if messageType == websocket.TextMessage {
					t1 := time.Now().UnixNano()
					json.Unmarshal(bytes, &req)
					res := handleWSCommand(req,resChan)
					//					devPrintn("WS: Command: ", req.Command, " Res: ", res)
					res.ResTime = (time.Now().UnixNano() - t1) / 1e6 //to miliscond
					//resChan <- res
				}
			}
		}()

		for r := range resChan {
			ws.WriteJSON(r)
			//ws.WriteMessage()
		}
		//after closing chanel
		err:=ws.Close()
		fmt.Println("closed: ",err)
		// for {
		// 	<-time.After(time.Second * 1)
		// 	ws.WriteJSON(User{UserName: "as"})

		// }
	}()
}

func handleWSCommand(req WSReq, chanel chan WSRes) (res WSRes) {
	//defer recover()
	//fmt.Println("WS: ",req)
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

	action := WSAction{Req: &req, Res: &res, Chanel: chanel}
	if fn != nil {
		fn(&action)
	} else {
		wsErrorCommand(&action)
	}
	return res
}
