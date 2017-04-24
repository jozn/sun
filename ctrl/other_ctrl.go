package ctrl

import "ms/sun/base"

func PingAction(c *base.Action) base.AppErr {
	c.SendText("PONG")
	return nil
}
