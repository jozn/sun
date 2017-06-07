package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	// "github.com/nfnt/resize"
	"log"
	"ms/sun/base"
	"ms/sun/models"
	"net/http"

	_ "github.com/garyburd/redigo/redis"
	//_ "net/http/pprof"
	_ "net/http/pprof"

	"github.com/mediocregopher/radix.v2/pool"
	"ms/sun/helper"
)

var redisPool *pool.Pool

func main() {
	print("Try To Start The App")
	startApp()

	/*go func() {
		models.SERVER_GRPC()
	}()

	go func() {
		models.Client_GRPC()
	}()*/

	/*go func() {
		helper.JustRecover()
		//http.ListenAndServe(":5000", nil)
	}()*/

	http.ListenAndServe(":5000", nil)

	//http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	//runtime.MemProfileRecord{}.
}

func startApp() {
	var err error
	//DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3308)/ms?charset=utf8mb4,utf8&collation=utf8mb4_general_ci")
	base.DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms?charset=utf8mb4")
	//DB, err = sqlx.Connect("mysql", "root:123456@localhost:3307/ms?charset=utf8mb4")
	//DB.Exec("SET NAMES 'utf8mb4';")
	noErr(err)
	base.DB.MapperFunc(func(s string) string { return s })
	redisInit()

	fmt.Println("main start")
	v1Tree := registerRoutes()
	//base.RegisterGlobTypes()

	//// Inits ///////////////
	//registerCmdRouters()

	http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	//in models
	models.OnAppStart_Models()

	_ = v1Tree

	go func() {
		helper.JustRecover()
		//http.ListenAndServe(":6000", v1Tree)
	}()

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

func noErr(err error) {
	if err != nil {
		// log.Fatalln(err) //painc
		log.Panic("***** PANIC ****: ", err)
	}
}

func actioner(action func(*base.Action)) http.Handler {
	return &base.Action{Fn: action}
}

//for Version 2 of Action -- that returns ActionErr
func actionToFunc(action func(*base.Action) base.AppErr) http.HandlerFunc {
	return (&base.Action{Fn2: action, Ver: 2}).ServeHTTP
}
