package main

import (
	// "encoding/json"
	// "fmt"
	// "reflect"
	//"strings"
	"math/rand"
)

func randSilce(sl []interface{}) interface{} {
	n := len(sl)
	return sl[rand.Intn(n)]
}
