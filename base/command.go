package base

import (
	"encoding/json"
	"ms/sun/helper"
	"time"
)

//todo: maybe merge WSReq and WSRes to one
type WSReq struct {
	//SessionUid string //???
	//Command    string //dep
	Commands   []Command
    NanoId uint64
	//RequestId  string //dep
	//ReqKey string //dep
	//Params     map[string]interface{}//dep
}

type WSRes struct {
	// BNCP
    Status       string //dep
	//ResTime   int64//dep eventing
    Commands     []*Command
    SyncedNanoId int64
	//RequestId string //dep
    ReqKey       string //dep
	//TODO remove?
	//Payload interface{} //dep no need, its eventing
}

type Command struct {
	//Command string
    Name            string
    ResId           int64
    ClientNanoId    int64
    ServerNanoId    int64
    Data            string        //marshilized json - don't set dirctly set via toJsonData
    toJsonSliceData []interface{} //`json:"-"`
	//Params  map[string]interface{}//dep
}

func NewResponseCommand(ResId int64) *Command {
	return &Command{
		Name:  "ResCmd",
		ResId: ResId,
		ServerNanoId: time.Now().UnixNano()}
}

func NewCommand(Name string) *Command {
	return &Command{Name: Name, ServerNanoId: time.Now().UnixNano()}
}

func (self *Command) SetData(data interface{}) {
	self.Data = helper.ToJson(data)
}

func (self *Command) AddSliceData(data interface{}) {
	self.toJsonSliceData = append(self.toJsonSliceData, data)
	//self.Data = helper.ToJson(data)
}

func (self *Command) MakeDataReady() {
	if self.toJsonSliceData != nil && len(self.toJsonSliceData) > 0 {
		self.Data = helper.ToJson(self.toJsonSliceData)
	}
	//self.Data = helper.ToJson(data)
}

//dep?
//func (self *Command) SetData(data interface{}) {
func (self *Command) ToCommandString() string {
	s, e := json.Marshal(self)
	if e != nil {
		helper.DebugPrintln("err in Command to json ")
		return ""
	}
	return string(s)
}

//dep?
func (self *Command) ToCommandStringBytes() []byte {
	s, e := json.Marshal(self)
	if e != nil {
		helper.DebugPrintln("err in Command to json ")
		return []byte("")
	}
	return s
}

type CmdAction struct {
	Req    *WSReq
	UserId int
	Cmd    *Command
}

var CmdMapRouter = make(map[string]func(*CmdAction))

//
//func serverWSReqCmds(req WSReq, pipe *UserDevicePipe) {
//    for _, cmd := range req.Commands {
//        fncmd := CmdMapRouter[cmd.Name]
//		fmt.Println("serving Cmd: ",cmd.Name, " Userid: ", pipe.UserId)
//        action := CmdAction{Req: &req, UserId: pipe.UserId , Cmd: &cmd}
//        if fncmd != nil {
//            fncmd(&action)
//        } else {//send cmd not found -- in debug
//            //wsErrorCommand(&action)
//        }
//    }
//}
