//package base
//
//import (
//    "github.com/gorilla/websocket"
//    "sync"
//    "fmt"
//    //"time"
//    "encoding/json"
//    "os"
//    "ms/sun/helper"
//)
//
//type UserDevicePipe struct  {
//    UserId int
//    IsOpen bool
//    Ws *websocket.Conn
//    ToDeviceChan chan *WSRes
//    FromDeviceChan chan *WSRes // NOT NEEDED?
//    Mutex sync.Mutex
//}
//
//func(pipe *UserDevicePipe) ServeIncomingReqs() {
//    go func() {
//        defer func() {
//            if r := recover(); r != nil {
//                fmt.Println("Recovered in ws messaging clinet request", r)
//            }
//        }()
//        //FIXME:this is sequence message proccessing do we need multi process
//        for {
//
//            var req WSReq
//            messageType, bytes, err := pipe.Ws.ReadMessage() //blocking
//            //cnt++
//            fmt.Println("messageType: "," ::",messageType, string(bytes))
//            if messageType == websocket.CloseMessage || err != nil {
//                pipe.IsOpen = false
//                close(pipe.ToDeviceChan)
//                fmt.Println("closeding ",messageType," error: ",err)
//                return
//            }
//            //TODO: another go fn here for many command proccessing
//            if messageType == websocket.TextMessage {
//                //t1 := time.Now().UnixNano()
//                json.Unmarshal(bytes, &req)
//                //res := handleWSCommand(req,resChan)
//                serverWSReqCmds(req,pipe)
//                //					devPrintn("WS: Command: ", req.Command, " Res: ", res)
//                //res.ResTime = (time.Now().UnixNano() - t1) / 1e6 //to miliscond
//                //resChan <- res
//            }
//
//            if messageType == websocket.BinaryMessage {
//                f,_ :=os.Create("./upload/ws_"+ helper.RandString(8))
//                wd,_:=os.Getwd()
//                fmt.Println("ws: "+" wd:" +wd + "binary:  ",bytes)
//                f.Write(bytes)
//                f.Close()
//            }
//        }
//    }()
//}
//
//func(pipe *UserDevicePipe) ServeSendToUserDevice() {
//    go func() {
//        //if we panic somehow? in writing json to writer, just close ws
//        defer func() {
//            if r := recover(); r != nil {
//                pipe.Ws.Close()
//            }
//        }()
//        for r := range pipe.ToDeviceChan {
//            pipe.Ws.WriteJSON(r)
//        }
//        //after closing chanel
//        err := pipe.Ws.Close()
//        pipe.IsOpen = false
//        fmt.Println("closed: ", err)
//    }()
//}
//
//func(pipe *UserDevicePipe) SendToUser(res WSRes) {
//    if pipe.IsOpen{
//        pipe.ToDeviceChan <- &res
//    }
//}
