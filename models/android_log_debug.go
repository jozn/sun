package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"ms/sun/helper"
	"net/http"
	"os"
	"time"
)

type androidDebugLogger struct {
	Module   string
	Time     int
	Severity string
	Text     string
	Version  int
}

var androidLoggerMapper = make(map[string]*os.File)

var startTime = time.Now()

func ServeHttpAndroidLogger(w http.ResponseWriter, r *http.Request) {
	defer helper.JustRecover()
	helper.DebugPrintln("ServeHttpAndroidLogger")
	err := r.ParseForm()
	var WSUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024 * 128,
		WriteBufferSize: 1024 * 128,
	}

	ws, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		helper.Debug("WSUpgrader err", err)
		return
	}

	for {
		messageType, bytes, err := ws.ReadMessage() //blocking

		if messageType == websocket.CloseMessage || err != nil {
			helper.Debugf("closeing android debug logger, err: %v", err)
			return
		}

		if messageType == websocket.BinaryMessage {

		}

		//helper.DebugPrintln("ServeHttpAndroidLogger got: ", string(bytes))

		if messageType == websocket.TextMessage {
			adl := &androidDebugLogger{}
			err = json.Unmarshal(bytes, adl)
			if adl.Time < 1 {
				adl.Time = helper.TimeNow()
			}
			t := time.Unix(int64(adl.Time), 0)
			if err == nil {
				//logFileName := fmt.Sprintf("%s_%d_%d_%d.javal", adl.Module, t.Year(), t.Day(), t.Hour())
				//logFileName := fmt.Sprintf("%s_%d.javal", adl.Module, startTime.Unix())
				logFileName := fmt.Sprintf("%s_%d.javal", adl.Module, adl.Version)
				file, ok := androidLoggerMapper[logFileName]
				if !ok {
					os.MkdirAll("./logs/android/", os.ModeDir)
					out := "./logs/android/" + logFileName //+ ".txt"
					file, err = os.Create(out)
					if err != nil {
						helper.DebugPrintln("ServeHttpAndroidLogger: error in file genrating: ", err)
						return
					}
					androidLoggerMapper[logFileName] = file
				}

				_, err = file.WriteString(fmt.Sprintf("%d:%d:%d - %s\n", t.Hour(), t.Minute(), t.Second(), adl.Text))
				file.Sync()
				helper.NoErr(err)
			} else {
				helper.DebugPrintln("ServeHttpAndroidLogger:json unmarsh error :  ", err)

			}
		}
	}

}
