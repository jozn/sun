package base

import (
	"encoding/json"
	"ms/sun/helper"
	"time"
)


//todo: maybe merge WSReq and WSRes to one
type WSReq struct {
	SessionUid string //???
	Command    string //dep
	Commands   []Command
	//RequestId  string //dep
	ReqKey     string //dep
	//Params     map[string]interface{}//dep
}

type WSRes struct {
	// BNCP
	Status    string //dep
	//ResTime   int64//dep eventing
	Commands  []*Command
	//RequestId string //dep
	ReqKey    string //dep
	//TODO remove?
	//Payload interface{} //dep no need, its eventing
}

type Command struct {
	//Command string
	Name       string
	Data       string      //marshilized json - don't set dirctly set via toJsonData
	toJsonData interface{} //`json:"-"`
	ResId int64
	CmdId int64
	//Params  map[string]interface{}//dep
}

func NewResCommand(ResId int64) *Command {
	return &Command{
		Name: "ResCmd" ,
		ResId: ResId,
		CmdId: time.Now().UnixNano()}
}

func NewCommand(ResId int64) Command {
	return Command{CmdId: time.Now().UnixNano()}
}

func (self *Command) SetData(data interface{}) {
	self.Data = helper.ToJson(data)
}

//func (self *Command) SetData(data interface{}) {
func (self *Command) ToCommandString() string {
	s, e := json.Marshal(self)
	if e != nil {
		helper.DebugPrintln("err in Command to json ")
		return ""
	}
	return string(s)
}

func (self *Command) ToCommandStringBytes() []byte {
	s, e := json.Marshal(self)
	if e != nil {
		helper.DebugPrintln("err in Command to json ")
		return []byte("")
	}
	return s
}

type CmdAction struct  {
    Req *WSReq
    UserId int
    Cmd *Command
}

var CmdMapRouter =  make(map[string]func(*CmdAction))

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


