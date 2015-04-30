package main

import (
	"net/http"
)

func registerRoutes() {
	http.Handle("/2", actioner(h2))
	http.Handle("/hello", actioner(helo))
	http.Handle("/x", actioner(helo))

	http.Handle("/play1", actioner(play1))
	http.Handle("/play2", actioner(play2))
	http.Handle("/play3", actioner(play3))
	http.Handle("/play4", actioner(play4))
	http.Handle("/play5", actioner(play5))

	http.Handle("/fact/p1", actioner(factPlay1))
	http.Handle("/fact/p2", actioner(factPlay2))
	http.Handle("/fact/p3", actioner(factPlay3))
	http.Handle("/fact/p4", actioner(factPlay4))

	http.HandleFunc("/", Home)
}
