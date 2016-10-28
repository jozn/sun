package main

import (
	// "ms/sun/fact/user_fact"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	// "github.com/nfnt/resize"
	cacheDrive "github.com/pmylund/go-cache"
	"log"
	//	"math/rand"
	"net/http"
	// "runtime"
	//	"./commands"
	"bytes"
	_ "github.com/garyburd/redigo/redis"
	//redis "github.com/mediocregopher/radix.v2"
	"github.com/gorilla/websocket"
	"math/rand"
	//. "ms/sun/base"
	"ms/sun/base"
	//. "ms/sun/models"
	"ms/sun/models"
	"net"
	//_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"
	//"github.com/pkg/profile"
	//"runtime"
	//"github.com/garyburd/redigo/redis"
	"github.com/mediocregopher/radix.v2/pool"
)

var cashe *cacheDrive.Cache

var __DEV__ bool

//var DB *sqlx.DB
var users []models.User

var redisPool *pool.Pool
var ws websocket.Conn

//var WSCommandMap map[string]func(*WSAction)

func main() {
	__DEV__ = true
	var err error
	//DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3308)/ms?charset=utf8mb4,utf8&collation=utf8mb4_general_ci")
	base.DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms?charset=utf8mb4")
	//DB, err = sqlx.Connect("mysql", "root:123456@localhost:3307/ms?charset=utf8mb4")
	//DB.Exec("SET NAMES 'utf8mb4';")
	// DB, err = sqlx.Connect("mysql", "root:123456@/ms3?charset=ascii")
	noErr(err)
	//xxx := DB.MapperFunc
	//DB.MapperFunc = xxx
	base.DB.MapperFunc(func(s string) string { return s })
	redisInit()

	//x := re
	//	commands.PR("asdas")
	//	commands.PR("k")

	// runtime.GOMAXPROCS(runtime.NumCPU())
	// DB.MustExec("set general_log_file=E:/Program Files/MySQL/MySQL Server 5.6/data/logs/log.txt;")
	// DB.Exec("SET global general_log = 1;")
	// debug(e)
	fmt.Println("main start")
	//	var err error
	///////////////redis////////////////////////
	//	redisServer, err = redis.Dial("tcp", "127.0.0.1:6379")
	//	fmt.Println("redisServer: ", redisServer)
	//	if err == nil {
	//		// handle error
	//		// redisServer.Do("SET", "GO", "t546sdSSA")
	//		// playRedis()
	//	}
	//defer redisServer.Close()

	////////////////MYSQL///////////////////
	//	DB, err = sqlx.Connect("mysql", "root:123456@/ms3?charset=utf8mb4")
	// DB, err = sqlx.Connect("mysql", "root:123456@/ms3?charset=ascii")
	//	noErr(err)
	//	DB.MapperFunc(func(s string) string { return s })
	v1Tree := registerRoutes()
	base.RegisterGlobTypes()

	//registerWSRoutes()
	//init_dbs()

	//// Inits ///////////////
	registerCmdRouters()

	// _casheCom = make(map[int][]Comment, 100)
	// _casheLike = make(map[int][]Like, 100)
	cashe = cacheDrive.New(5*time.Minute, 30*time.Second)

	// http.Handle("/2", actioner(h2))
	// http.Handle("/hello", actioner(helo))
	// http.HandleFunc("/", Home)
	// GetMedias()
	//	runMysql()
	http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	//go serverEcho()
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	//in models
	models.OnAppStart_Models()

    _ = v1Tree
	http.ListenAndServe(":5000", nil)
	//runtime.MemProfileRecord{}.
}

func redisInit() {
	var err error
	redisPool, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		fmt.Println("redis failed")
		return
	}
	fmt.Println("redis online")
	// In another go-routine

	conn, err := redisPool.Get()
	if err != nil {
		// handle error
	}

	if conn.Cmd("SET", "KEYYY", "Val").Err != nil {
		// handle error
	}

	redisPool.Put(conn)
}

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Println("from Home")
	fmt.Println(r.URL.RequestURI())
	fmt.Fprintln(w, r.Host, r.UserAgent())

}

func noErr(err error) {
	if err != nil {
		// log.Fatalln(err) //painc
		log.Panic("***** PANIC ****: ", err)
	}
}

//func noDevErr(err error) {
//	if __DEV__ && err != nil {
//		// log.Fatalln(err) //painc
//		log.Panic("***** PANIC ****: ", err)
//	}
//}

func actioner(action func(*base.Action)) http.Handler {
	return &base.Action{Fn: action}
}

func actioner2(action func(*base.Action)) http.HandlerFunc {
	return (&base.Action{Fn: action}).ServeHTTP
}

//for Version 2 of Action -- that returns ActionErr
func actionToFunc(action func(*base.Action) base.AppErr) http.HandlerFunc {
	return (&base.Action{Fn2: action, Ver: 2}).ServeHTTP
}

//func dummy() {
//	fmt.Println(rand.Intn(2))
//}

//empty for fuell go nasty varible must used
func e(...interface{}) {

}

func helo(c *base.Action) {
	c.SendText("das")
	e(ws)
}

func runMysql() {
	var DB *sqlx.DB
	DB, err := sqlx.Connect("mysql", "root:123456@/ms3?charset=utf8mb4")
	// DB, err = sqlx.Connect("mysql", "root:123456@/ms3?charset=ascii")
	noErr(err)
	DB.MapperFunc(func(s string) string { return s })
	val := "( 152, 'text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 ', 'chars 1' )"
	var a []string
	for i := 0; i < 1000; i++ {
		a = append(a, val)
	}
	vals := strings.Join(a, ",")
	t2 := time.Now().UnixNano()
	if DB == nil {
		fmt.Println("NILLLLLLLLLLLLLLLLLLLLLL")
	}
	for i := 0; i < 10000; i++ {
		t2 := time.Now().UnixNano()
		DB.Exec("insert into test2 (Id, Text, Str) values " + vals)
		r := ("time: " + strconv.Itoa(int((time.Now().UnixNano()-t2)/1e6)))
		if i%10 == 0 {
			fmt.Println("RES ", strconv.Itoa(i), " ::", r)
		}
	}
	x, err := DB.Exec("insert into test2 (Id, Text, Str) values " + vals)
	fmt.Println("RES MYSQL:: ", x, err)
	r := ("time: " + strconv.Itoa(int((time.Now().UnixNano()-t2)/1e6)))
	fmt.Println("RES:: ", r)
}

func serverEcho() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		//go p_handleEcho(conn)
		go p_handleConnection(conn)
	}
}

func p_handleConnection(con net.Conn) {
	//var  ss := []byte(`HTTP/1.1 200 OK
	ss := (`HTTP/1.1 200 OK
Date: Mon, 23 May 2005 22:38:34 GMT
Content-Type: text/html; charset=UTF-8
Content-Encoding: UTF-8
Content-Length: 138
Last-Modified: Wed, 08 Jan 2003 23:11:55 GMT
Server: Apache/1.3.3.7 (Unix) (Red-Hat/Linux)
ETag: "3f80f-1b6-3e1cb03b"
Accept-Ranges: bytes
Connection: close

<html>
<head>
  <title>An Example Page</title>
</head>
<body>
  Hello World, this is a very simple HTML document.
`)
	r := make([]byte, 5000000)
	time.Sleep(time.Millisecond * 4000)
	i, _ := con.Read(r)
	print(i)
	s2 := ss + string(r) + `</body>
</html>`
	con.Write([]byte(s2))
	//for {
	//	time.Sleep(1e7)
	//	con.Write([]byte("AS"))
	//}
	f, _ := os.Create("./___w1" + strconv.Itoa(rand.Intn(10000)))
	//f.Write(r[:bytes.IndexByte(r, 0x00)])
	f.Write(r)
	f.Close()
	//con.Close()
	//con.

}

func p_handleEcho(con net.Conn) {
	for {
		r := make([]byte, 1000000)
		time.Sleep(time.Millisecond * 1000)
		i, _ := con.Read(r)
		print(i, "\n")
		//fmt.Println("Echo got size:",i," load: ",string(r))
		print("Echo got: ", len(string(r[:i])), " ", string(r[:i]), "\n")
		//strings.
		r = append(r, '\n')
		//con.Write(r[:i])
		print("bb: ", bytes.IndexByte(r, 0x00), " ", r[:bytes.IndexByte(r, 0x00)])
		f, _ := os.Create("./fsd/_w1" + strconv.Itoa(rand.Intn(10000)))
		f.Write(r[:bytes.IndexByte(r, 0x00)])
	}
}

//func h2(c *Action) {
//	c.SendText(c.Req.UserAgent() + c.Req.Method)
//}
