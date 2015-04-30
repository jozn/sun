package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

//parent action of all actions
type Action struct {
	f      http.Handler
	c      int
	Fn     func(*Action) //sub-action
	Req    *http.Request
	Res    *http.ResponseWriter
	UserId int
}

//send a copy(not pointer*) of Actions
func (h Action) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.c = h.c + 1
	h.Req = r
	h.Res = &w
	h.Fn(&h)
	// i := 5
	// w.Write([]byte(h.c))
	if rand.Intn(2) == 1 || true {
		// fmt.Fprintln(w, r.UserAgent(), h.c)
		// panic("asd")
	}
	// fmt.Fprintln(w, r.UserAgent(), h.c, i)
}

func (h *Action) IsUser()   {}
func (h *Action) MustUser() {}
func (h *Action) MustPost() {}
func (h *Action) SendJson() {}

func (h *Action) SendText(s string) {
	fmt.Fprintln(*h.Res, s) //[]byte(s))
}

func (h *Action) SendPalinText() {}
func (h *Action) TurnOffGzip()   {}
