package pipesold

import (
	"encoding/json"
	"fmt"
	"ms/sun/base"
)

func CommandsReceived(c *base.CmdAction) {
	d := []byte(c.Cmd.Data)
	fmt.Println("CommandsReceived: user:", c.UserId, " data:", string(d))
	cmdIds := make([]int64, 0)
	err := json.Unmarshal(d, &cmdIds)
	if err == nil && len(cmdIds) > 0 {
		//var min,max int64
		min, max := cmdIds[0], cmdIds[0]
		for _, n := range cmdIds {
			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}
		RemoveCommandsFromRedis(c.UserId, min, max)
	}
}

func CommandsReceived_Dep(c *base.CmdAction) {
	d := []byte(c.Cmd.Data)
	fmt.Println("CommandsReceived: user:", c.UserId, " data:", string(d))
	cmdIds := make([]int64, 0)
	err := json.Unmarshal(d, &cmdIds)
	if err == nil && len(cmdIds) > 0 {
		//var min,max int64
		min, max := cmdIds[0], cmdIds[0]
		for _, n := range cmdIds {
			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}
		RemoveCommandsFromRedis(c.UserId, min, max)
	}
}
