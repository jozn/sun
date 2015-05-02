package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

//parent action of all actions
type Action struct {
	f         http.Handler
	c         int
	_bodyText string        //do we need bodRes?
	_bodyRes  []byte        //byte beacuse to support gzip
	Fn        func(*Action) //sub-action
	Req       *http.Request
	Res       *http.ResponseWriter
	UserId    int
}

//send a copy(not pointer*) of Actions
func (c Action) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.c = c.c + 1
	c.Req = r
	c.Res = &w
	w.Header().Add("Access-Control-Allow-Origin", "*")
	c.Fn(&c)
	// i := 5
	// w.Write([]byte(c.c))
	if rand.Intn(2) == 1 || true {
		// fmt.Fprintln(w, r.UserAgent(), c.c)
		// panic("asd")
	}
	// fmt.Fprintln(w, r.UserAgent(), c.c, i)
	fmt.Fprintln(*c.Res, c._bodyText) //[]byte(s))
}

func (c *Action) IsUser()   {}
func (c *Action) MustUser() {}
func (c *Action) MustPost() {}
func (c *Action) SendJson(stuc interface{}) {
	b, err := json.Marshal(stuc)
	if __DEV__ && err != nil {
		log.Fatal("json Marshaling error in send json response: ", err)
	}
	c._bodyText = string(b)
	// c.SendText(string(b))
}

func (c *Action) SendText(s string) {
	// fmt.Fprintln(*c.Res, s) //[]byte(s))
	c._bodyText = s
}

func (c *Action) SendPalinText() {}
func (c *Action) TurnOffGzip()   {}
