package base

import (
    "net/http"
    "fmt"
    "github.com/gorilla/websocket"
)

func ServeUserWs(w http.ResponseWriter, r *http.Request) {
    //defer recover()
    err := r.ParseForm()
    noErr(err)
    session := r.Form.Get("SessionUid")

    //TODO add session check functiality
    if false {
        e(session)
        return //not Autorized user
    }

    var WSUpgrader = websocket.Upgrader{
        ReadBufferSize:  1024 * 128,
        WriteBufferSize: 1024 * 128,
    }

    ws, err := WSUpgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("WSUpgrader err", err)
        return
    }

    //TODO get userId
    uid := 6
    AllPipesMap.ServeUserWs(uid,ws)
}