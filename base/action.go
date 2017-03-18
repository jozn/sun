package base

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"ms/sun/config"
	"ms/sun/constants"
	"ms/sun/helper"
	"net/http"
	"time"
)

//parent action of all actions
type Action struct {
	Protocol  BNCP
	f         http.Handler //dep?
	c         int
	_bodyText string //dep? do we need bodRes?
	_bodyRes  []byte //dep? byte beacuse to support gzip
	_payload  interface{}
	Fn        func(*Action)        //sub-action
	Fn2       func(*Action) AppErr //sub-action
	Req       *http.Request
	Res       http.ResponseWriter
	Ver       int
	_userId   int
	// UserId    int
}

type BNCP struct {
	Status     string
	Error      string
	Payload    *interface{}
	Meta       string
	ServerTime int
	Version    int
	ResTime    int64
	// UserId    int
}

//send a copy(not pointer*) of Actions
//c standfor controller
func (c Action) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//defer recover()
	defer func() {
		e := recover()
		if e != nil {
			if e == constants.HttpIsNotUser {
				c.Protocol.Error = "باید وارد شوید."
				c.Protocol.Status = "Error"
				c.SendJson(nil)
				setResponseBody(&c, w, time.Now())
			}

			fmt.Println("PANIC 2: ", e)
			w.Write([]byte(fmt.Sprintf("%v", e)))
		}

	}()
	t1 := time.Now()

	//prot := BNCP{}
	c.Protocol.Status = "ok"
	c.Protocol.ServerTime = int(time.Now().Unix())
	c.Protocol.Version = 1
	//c.Protocol = prot

	c.c = c.c + 1
	c.Req = r
	c.Res = w
	c.Req.ParseMultipartForm(1e6)//c.Req.ParseForm()
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "x-ms-uuid")
	if r.Method != "OPTIONS" { //what was this for browser CROS
		if c.Ver == 2 {
			actionErr := c.Fn2(&c) //todo: work on actionErr
			_ = actionErr
		} else {
			c.Fn(&c)
		}
	} else { //browser CROS check
		c._payload = "OPTIONS"
	}

	if false {
		fmt.Println(r.RemoteAddr)
		fmt.Println(r.RequestURI)
		fmt.Println(r.Referer())
		fmt.Println(r.Host)
		fmt.Println(r.Header)
		fmt.Println(r.Method)
		fmt.Println(r.UserAgent())
		// panic("asd")
	}
	setResponseBody(&c, w, t1)
}

func setResponseBody(c *Action, w http.ResponseWriter, t1 time.Time) {
	c.Protocol.Payload = &c._payload
	c.Protocol.ResTime = time.Now().Sub(t1).Nanoseconds() / 1e6 //ms

	// fmt.Fprintln(w, r.UserAgent(), c.c, i)
	// fmt.Fprintln(*c.Res, c._bodyText) //[]byte(s))

	//TODO :mereg with SendJson
	b, err := json.Marshal(c.Protocol)
	if config.IS_DEBUG && err != nil {
		log.Fatal("json Marshaling error in send json response: ", err)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-ResTime", fmt.Sprint(time.Now().Sub(t1).Nanoseconds()/1e6))
	w.Header().Set("X-Time", fmt.Sprint(time.Now().UnixNano()/1e9))

	if len(b) > 1300 { //860: Akami cdn defualts
		//w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Content-Encoding", "gzip")
		bgzip, _ := gzip.NewWriterLevel(c.Res, gzip.BestSpeed)
		bgzip.Write(b)
		bgzip.Close()
	} else {
		fmt.Fprintln(c.Res, string(b))
	}

}

func (c *Action) IsUser() bool {
	return true
}

func (c *Action) UserId() int {
	return c._userId
}

func (c *Action) SetUserId(UserId int) {
	c._userId = UserId
}

func (c *Action) MustUser() {
	//if not user panic

}

func (c *Action) MustPost() {
	//for dev always true
}
func (c *Action) SendJson(res interface{}) {
	c._payload = res
}

func (c *Action) SendText(s string) {
	// fmt.Fprintln(*c.Res, s) //[]byte(s))
	c._payload = s
}

func (c *Action) GetPage() int {
	pageStr := c.Req.Form.Get("page")
	return helper.StrToInt(pageStr, 0)
}

func (c *Action) GetParamInt(param string, defulat int) int {
	val := c.Req.Form.Get(param)
	intVal := helper.StrToInt(val, defulat)
	if intVal == 0 {
		return defulat
	}
	return intVal
	//return helper.StrToInt(val, defulat)
}

func (c *Action) SendPalinText() {}
func (c *Action) TurnOffGzip()   {}

////////////////////////////////////////////////////////////////////////

type ActionErr struct {
	Code string
}

func (err ActionErr) Error() string {
	return err.Code
}
