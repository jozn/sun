package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var urls = []string{
	"http://www.google.com/",
	// "http://www.google.tr/",
	"http://www.google.us/",
	"http://www.google.fr/",
	"http://www.bbc.co.uk/",
	"http://www.rk.com",
	"http://www.hqll.com/",
	"http://www.sextury.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
	"http://blogfa.com/",
	"http://mardomsara.com/",
	"http://mardomsara.ir/",
	"http://sex.com/",
}

func main() {
	// Execute an HTTP HEAD request for all url’s
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		go act(url)
		time.Sleep(1 * 1e8)
	}

	time.Sleep(8 * 1e9)
}

func act(url string) {
	defer recover()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", url, err)
	}
	v, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(url, ": ", resp.Status, " - ", len(v), "\n")
	// fmt.Print(url, ": ", resp.Status, resp.Body, "\n")
}
