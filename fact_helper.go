package main

import (
	// "encoding/json"
	// "fmt"
	// "reflect"
	//"strings"
	"math/rand"
	"strconv"
)

func randSilceString(sl []string) string {
	n := len(sl)
	return sl[rand.Intn(n)]
}

func _randomEmail() string {
	n := rand.Intn(1000000)
	return "eamil-" + strconv.Itoa(n) + "@gmail.com"
}
