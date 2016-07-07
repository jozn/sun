package actions

import (
	// "fmt"
	// "github.com/gorilla/websocket"
	// "net/http"
	// // "log"
	// "encoding/json"
	. "ms/sun/base"
	. "ms/sun/models"
	"time"
)

func HelloCommand(c *WSAction) {
	//	_ = User{}
	time.Sleep(time.Microsecond * 10)
	//	c.Req.Params
	c.Res.Payload = User{}
	c.SendRes()
}

func Hello2Command(c *WSAction) {
	//	_ = User{}
	time.Sleep(time.Microsecond * 10)
	//	c.Req.Params
	c.Res.Payload = Post{}
}

func WsErrorCommand(c *WSAction) {
	c.Res.Status = "ERROR"
}

func WsPanicCommand(c *WSAction) {
	panic("ads")

	//c.Res.Status = "ERROR"
}

func Cmd_getAllGroups(c *WSAction) {
	// c.Res.Status = "ERROR"
	//	cv := WSReq{}

}
