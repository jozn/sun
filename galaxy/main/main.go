package main

import (
	"ms/sun/galaxy"
	"net/http"
)

func main() {
	galaxy.Run()
	galaxy.Insert_many(100000)
	http.ListenAndServe(":4444", nil)
}
