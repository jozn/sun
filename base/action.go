package base

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	"ms/sun/helper"
	"compress/gzip"
	"fmt"
)

//parent action of all actions
type Action struct {
	Protocol BNCP
	f         http.Handler //dep?
	c         int
	_bodyText string //dep? do we need bodRes?
	_bodyRes  []byte //dep? byte beacuse to support gzip
	_payload  interface{}
	Fn        func(*Action) //sub-action
	Fn2        func(*Action) AppErr //sub-action
	Req       *http.Request
	Res       *http.ResponseWriter
	Ver         int
	// UserId    int
}

//from old EM implmn
//    @response['version'] = '0.1'
//    @response['status'] ||= RESPONSE_STATUS[:ok]
//    @response['payload'] = text
//    @response['connection_id'] = @client[:ws].__id__
//    @response['request_id'] = @request['request_id']
//    @response['server_time'] = Time.now.to_i
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
	defer recover()
	t1 := time.Now()

	//prot := BNCP{}
	c.Protocol.Status = "ok"
	c.Protocol.ServerTime = int(time.Now().Unix())
	c.Protocol.Version = 1
	//c.Protocol = prot

	c.c = c.c + 1
	c.Req = r
	c.Res = &w
	c.Req.ParseForm()
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "x-ms-uuid")
	if r.Method != "OPTIONS" { //what was this for browser CROS
		if c.Ver == 2 {
			actionErr := c.Fn2(&c) //todo: work on actionErr
			_ = actionErr
		}else {
			c.Fn(&c)
		}
	} else { //browser CROS check
//		devPrintn("callsed OPTIONS http request")
		c._payload = "OPTIONS"
	}

	// i := 5
	// w.Write([]byte(c.c))
	if rand.Intn(2) == 1 || true {
		// fmt.Fprintln(w, r.UserAgent(), c.c)
		// panic("asd")
	}
	c.Protocol.Payload = &c._payload
	c.Protocol.ResTime = time.Now().Sub(t1).Nanoseconds() / 1e6 //ms

	// fmt.Fprintln(w, r.UserAgent(), c.c, i)
	// fmt.Fprintln(*c.Res, c._bodyText) //[]byte(s))

	//TODO :mereg with SendJson
	b, err := json.Marshal(c.Protocol)
	if __DEV__ && err != nil {
		log.Fatal("json Marshaling error in send json response: ", err)
	}
	if len(b) > 1300 {//860: Akami cdn defualts
		w.Header().Set("Content-Type","text/html")
		w.Header().Set("Content-Encoding","gzip")
		bgzip,_ := gzip.NewWriterLevel(*c.Res, gzip.BestSpeed)
		bgzip.Write(b)
		bgzip.Close()
	}else {
		fmt.Fprintln(*c.Res, string(b))
	}
}

func (c *Action) IsUser() bool {
	return true
}

func (c *Action) UserId() int {
	return 6
}

func (c *Action) MustUser() {
	//if not user panic
}

func (c *Action) MustPost() {
	//for dev always true
}
func (c *Action) SendJson(res interface{}) {
	c._payload = res
	// b, err := json.Marshal(stuc)
	// if __DEV__ && err != nil {
	// 	log.Fatal("json Marshaling error in send json response: ", err)
	// }
	// c._bodyText = string(b)
	// c.SendText(string(b))
}

func (c *Action) SendText(s string) {
	// fmt.Fprintln(*c.Res, s) //[]byte(s))
	c._payload = s
}

func (c *Action) GetPage() int {
	pageStr := c.Req.Form.Get("page")
	return helper.StrToInt(pageStr,0)
}

func (c *Action) GetParamInt(param string, defulat int) int {
	val := c.Req.Form.Get(param)
	return helper.StrToInt(val,0)
}


func (c *Action) SendPalinText() {}
func (c *Action) TurnOffGzip()   {}


////////////////////////////////////////////////////////////////////////

type ActionErr struct  {
	Code string
}

func (err ActionErr) Error() string  {
	return err.Code
}