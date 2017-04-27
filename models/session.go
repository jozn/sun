package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"net/http"
	"time"
)

/*
   specs:
   LastActivityTime: last time that user has made any http reqs: load, like, ...
   UpdatedTime: last update of session info changes like: Ip, MAC Address, Network type

*/

func Session_ProcessHttpReq(h *http.Request) *ReqParams {
	req := &ReqParams{
		UserId:      helper.StrToInt(h.Form.Get("user_id"), 0),
		SessionUuid: h.Form.Get("session"),
		Page:        helper.StrToInt(h.Form.Get("page"), 1),
		Last:        helper.StrToInt(h.Form.Get("last"), 0),
		Limit:       helper.StrToInt(h.Form.Get("limit"), 30),
	}

	isUser := Session_CheckAndSetUserSession(req)
	if !isUser {
        if h.Form.Get("magic") == "true" {//todo add IS_DEBUG befor production

        }else {
            req.UserId = 0
            req.SessionUuid = ""
        }
	}

	return req
}

func Session_CheckAndSetUserSession(req *ReqParams) bool {
	if req.UserId < 1 || req.SessionUuid == "" {
		return false
	}

	session, ok := Store.Session_BySessionUuid(req.SessionUuid)

	if !ok {
		return false
	}

	req.Session = session
	return true
}

var session_lastUsersActiveChan = make(chan string, 10000)

func Session_UpdatesForLastActions(ses *Session) bool {
	if ses != nil && ses.UserId > 0 {
		session_lastUsersActiveChan <- ses.SessionUuid
	}
	return true
}

//call this in the peridocs
func Session_periodicllyUpdateDB() {
	defer helper.JustRecover()

	var toUpDateUuids []string
	for {
		select {
		case uuid := <-session_lastUsersActiveChan:
			toUpDateUuids = append(toUpDateUuids, uuid)

		case <-time.Tick(time.Millisecond * 500):
			if len(toUpDateUuids) > 0 {
				cp := toUpDateUuids
				toUpDateUuids = []string{} //empty it

				NewSession_Updater().
					LastActivityTime(helper.TimeNow()).
					SessionUuid_In(cp).
					Update(base.DB)
			}
		}
	}
}
