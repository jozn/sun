package main

import (
	// "./fact/user_fact"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"net/http"
	"runtime"
)

var DB *sqlx.DB
var users []User

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("main start")
	var err error
	DB, err = sqlx.Connect("mysql", "root:123456@/ms2")
	noErr(err)
	DB.MapperFunc(func(s string) string { return s })
	registerRoutes()

	// http.Handle("/2", actioner(h2))
	// http.Handle("/hello", actioner(helo))
	// http.HandleFunc("/", Home)

	http.ListenAndServe(":5000", nil)

}

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Println("from Home")
	fmt.Println(r.URL.RequestURI())
	fmt.Fprintln(w, r.Host, r.UserAgent())

}

func noErr(err error) {
	if err != nil {
		log.Fatalln(err) //painc
	}
}

func actioner(action func(*Action)) http.Handler {
	return &Action{Fn: action}
}

func dummy() {
	fmt.Println(rand.Intn(2))
}

//empty for fuell go nasty varible must used
func e(...interface{}) {

}

func helo(c *Action) {
	c.SendText("das")
}

func h2(c *Action) {
	c.SendText(c.Req.UserAgent() + c.Req.Method)
}
