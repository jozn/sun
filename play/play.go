package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type test5 struct {
	Id   int
	Text string
}

type testView5 struct {
	BlogId int
	TextMe string
	Load   *test5
}

type testEmbed5 struct {
	*testView5
	Timestamp int
}

func play1(c *Action) {
	for i := 0; i < 500; i++ {
		DB.MustExec("INSERT INTO test (text) VALUES (?)", "dssdcsdcs udsbc sducbsdcc scybscysc scscbsdss csdcybsycbsc scsycbsc scsycb cssssssssss")

	}
}

func play2(c *Action) {
	var re []test5
	for i := 0; i < 50; i++ {
		DB.Select(&re, "select * from test where id = ?", rand.Intn(1000000))
	}
	b, _ := json.Marshal(re)
	c.SendText(string(b))
}

func play3(c *Action) {
	var re []test5
	var in []string
	for i := 0; i < 50; i++ {
		in = append(in, strconv.Itoa(rand.Intn(1000000)))
		// DB.Select(&re, "select * from test where id = ?", rand.Intn(1000000))
	}
	ins := strings.Join(in, ",")
	print(ins, "\n")
	DB.Select(&re, "select * from test where id in("+ins+")")
	//fmt.Println(len(re))
	b, err := json.Marshal(re)
	// fmt.Println(re)
	noErr(err)
	c.SendText(string(b))
}

func play4(c *Action) {
	var re []test5
	var in []string
	for i := 0; i < 50; i++ {
		in = append(in, strconv.Itoa(rand.Intn(1000000)))
		// DB.Select(&re, "select * from test where id = ?", rand.Intn(1000000))
	}
	ins := strings.Join(in, ",")
	// print(ins, "\n")
	DB.Select(&re, "select * from test where id in("+ins+")")
	//fmt.Println(len(re))
	var js []testView5
	for i := 0; i < len(re); i++ {
		l := testView5{25, "تایشسعااشس", &re[i]}
		js = append(js, l)
	}

	b, err := json.Marshal(js)
	// fmt.Println(re)
	noErr(err)
	c.SendText(string(b))
}

func play5(c *Action) {
	var re []test5
	var in []string
	for i := 0; i < 50; i++ {
		in = append(in, strconv.Itoa(rand.Intn(1000000)))
		// DB.Select(&re, "select * from test where id = ?", rand.Intn(1000000))
	}
	ins := strings.Join(in, ",")
	// print(ins, "\n")
	DB.Select(&re, "select * from test where id in("+ins+")")
	//fmt.Println(len(re))emvar js []testView5
	var js []testView5
	for i := 0; i < len(re); i++ {
		l := testView5{25, "تایشسعااشس", &re[i]}
		js = append(js, l)
	}

	var em []testEmbed5
	for i := 0; i < len(js); i++ {
		l := testEmbed5{&js[i], 888888}
		em = append(em, l)
	}

	b, err := json.Marshal(em)
	// fmt.Println(re)
	noErr(err)
	c.SendText(string(b))
}

//play with r.Forms
func play6(c *Action) {
	// c.Req.ParseForm()
	f := c.Req.Form.Get("key") + c.Req.RemoteAddr + c.Req.RequestURI //+ "  " + string(c.Req.Body)
	// n := c.Req.Form.

	c.SendText(f)
}

func dumy1111() {
	fmt.Println("...")
}
