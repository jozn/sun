package base

import (
	"ms/sun/helper"
	"time"
)

type Call struct {
	Name            string
	UserId          int
	ClientCallId    int64
	ServerCallId    int64
	Data            string        //marshilized json - don't set dirctly set via toJsonData
	toJsonSliceData []interface{} //`json:"-"`
	Meta            string        //marshilized json - don't set dirctly set via toJsonData
    TimeMs          int
}

func NewCall(Name string) Call {
	return Call{
        Name: Name,
        ServerCallId: time.Now().UnixNano(),
        TimeMs: int(time.Now().UnixNano()/1000000),
    }
}

func NewCallWithData(Name string, Data interface{}) Call {
	c := NewCall(Name)
	c.SetData(Data)
	return c
}

func (self *Call) SetData(data interface{}) {
	self.Data = helper.ToJson(data)
}

func (self *Call) SetMeta(meta interface{}) {
	self.Meta = helper.ToJson(meta)
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
