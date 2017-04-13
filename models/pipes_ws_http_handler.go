package models

import (
	"ms/sun/helper"
	"net/http"

	"github.com/gorilla/websocket"
)

func ServeHttpWs(w http.ResponseWriter, r *http.Request) {
	//defer recover()
	err := r.ParseForm()
	//noErr(err)
	session := r.Form.Get("SessionUid")
	_ = session

	uid := 6
	uid = helper.StrToInt(r.Form.Get("UserId"), 6)
	//TODO add session check functiality
	if false {
		//e(session)
		return //not Autorized user
	}

	var WSUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024 * 128,
		WriteBufferSize: 1024 * 128,
	}

	ws, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		helper.Debug("WSUpgrader err", err)
		return
	}

	//TODO get userId
	AllPipesMap.ServeNewHttpWsForUser(uid, ws)

}
