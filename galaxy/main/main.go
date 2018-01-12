package main

import (
	"ms/sun/galaxy"
	"net/http"
)

func main() {
	galaxy.Run()
	http.ListenAndServe(":4444", nil)
}
