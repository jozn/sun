package main

import (
	"ms/sun/actions"
	"ms/sun/base"
	"ms/sun/commands"
)

//command events that clinet invokes
func registerWSRoutes() {
	base.WSCommandMap = make(map[string]func(*base.WSAction))

	base.WSCommandMap["hello"] = actions.HelloCommand
	base.WSCommandMap["hello2"] = actions.Hello2Command
	base.WSCommandMap["time"] = commands.TimeCommand2
	base.WSCommandMap["time1"] = commands.TimeCommand
	base.WSCommandMap["panic"] = actions.WsPanicCommand

}

