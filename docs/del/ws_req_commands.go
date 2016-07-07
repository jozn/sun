package main

import (
	// "fmt"
	// "github.com/gorilla/websocket"
	// "net/http"
	// // "log"
	// "encoding/json"
	"time"
)

func helloCommand(c *WSAction) {
	x := User{}
	time.Sleep(time.Microsecond * 10)
	c.Res.Payload = x
}

func wsErrorCommand(c *WSAction) {
	c.Res.Status = "ERROR"
}

func wsPanicCommand(c *WSAction) {
	panic("ads")

	//c.Res.Status = "ERROR"
}

func cmd_getAllGroups(c *WSAction) {
	// c.Res.Status = "ERROR"
	//	cv := WSReq{}

}
