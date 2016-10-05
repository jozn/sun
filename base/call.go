package base

import (
	"ms/sun/helper"
	"time"
)

type Call struct {
	Name            string
	UserId          int
	ClientNanoId    int64
	ServerNanoId    int64
	Data            string        //marshilized json - don't set dirctly set via toJsonData
	toJsonSliceData []interface{} //`json:"-"`
}

func NewCall(Name string) *Call {
	return &Call{Name: Name, ServerNanoId: time.Now().UnixNano()}
}

func (self *Call) SetData(data interface{}) {
	self.Data = helper.ToJson(data)
}

func (self *Call) MakeDataReady() {
	if self.toJsonSliceData != nil && len(self.toJsonSliceData) > 0 {
		self.Data = helper.ToJson(self.toJsonSliceData)
	}
	//self.Data = helper.ToJson(data)
}

var CallMapRouter = make(map[string]func(Call))

/*
func (self *Call) AddSliceData(data interface{}) {
	self.toJsonSliceData = append(self.toJsonSliceData, data)
	//self.Data = helper.ToJson(data)
}
*/
