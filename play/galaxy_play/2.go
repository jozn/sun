package main

import (
	//"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"image/jpeg"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type Media struct {
	Id   int    `db:"id"`
	Key  string `db:"key"`
	Load []byte `db:"load"`
}

var db *sqlx.DB
var err error
var cnt = 0
var size = 0

func main() {
	db, err = sqlx.Open("postgres", "postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable")

	if err != nil {
		panic(err)
	}

	//many_select(10000)
	//insert_many(10000)
	//write()
	sample_bench_simple_writes_set()
	go sample_bench_simple_writes()
	go sample_bench_simple_writes()
	go sample_bench_simple_writes()
	go sample_bench_simple_writes()
	go sample_bench_simple_writes()
	go sample_bench_simple_writes()
	go sample_bench_simple_writes()
	go sample_bench_simple_writes()
	go sample_bench_simple_writes()
	sample_bench_simple_writes()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	go sample_bench_simple_read()
//	sample_bench_simple_read()
}

func write() {
	for i := 0; i < 10000; i++ {
		write_one(i)
	}
}

func write_one(id int) {
	var med Media

	err = db.Get(&med, "select * from galaxy.msg WHERE id = $1 limit 1", id)
	if err != nil {
		log.Println("EEEEEE ", err)
		return
	}
	err = ioutil.WriteFile(fmt.Sprintf(`E:\_cockroachdb\out2/%d.jpg`, id), med.Load, os.ModePerm)

	if err != nil {
		log.Println("EEEEEE ", err)
		log.Println(os.Getwd())
		return
	}
	fmt.Println("", id, " ", len(med.Load))

}

func sample_bench_simple_writes_set() {
	db.MustExec("CREATE database if not EXISTS galaxy ")
	db.MustExec("Drop table if EXISTS galaxy.simple ")
	db.MustExec("CREATE TABLE if not EXISTS galaxy.simple (id INT PRIMARY KEY, key string)")
}

var i = 0

func sample_bench_simple_writes() {
	for ; i < 100000; i++ {
		db.Exec("INSERT into galaxy.simple (id,key) VALUES ($1,$2)", i, "dasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdas")
		if i%1000 == 0 {
			fmt.Println("cnt: ", i, " ", size/1000000)
		}
	}
}

func sample_bench_simple_read() {
	for ; i < 100000; i++ {
		db.QueryRow("SELECT * from galaxy.simple WHERE id = $1 ", i).Scan()
		if i%1000 == 0 {
			fmt.Println("cnt: ", i, " ", size/1000000)
		}
	}
}
func many_select(num int) {
	for i := 0; i < num; i++ {
		select_one(i)
		//fmt.Println("cnt: ", i)
	}
}
func select_one(i int) {
	var med Media

	err = db.Get(&med, "select * from galaxy.msg WHERE id = $1 limit 1", rand.Intn(9999)+1)
	if err != nil {
		log.Println("EEEEEE ", err)
	}
	fmt.Println("", i, " ", len(med.Load))
}
func for_insert() {
	db.MustExec("CREATE database if not EXISTS galaxy ")
	db.MustExec("Drop table if EXISTS galaxy.msg ")
	db.MustExec("CREATE TABLE if not EXISTS galaxy.msg (id INT PRIMARY KEY, key string, load bytes)")
}

func insert() {
	name, bs, err := RandImage()
	if err != nil {
		return
	}
	cnt++
	m := Media{
		Id:   cnt, //int(time.Now().UnixNano()),
		Key:  name,
		Load: bs,
	}
	db.MustExec("INSERT into galaxy.msg (id,key,load) VALUES ($1,$2,$3)", m.Id, m.Key, m.Load)
}

func insert_many(num int) {
	for_insert()
	for i := 0; i < num; i++ {
		insert()
		fmt.Println("cnt: ", i, " ", size/1000000)
	}
}

func RandImage() (fn string, bs []byte, err error) {
	const dir = `C:\Go\_gopath\src\ms\sun\upload\samples`
	imageFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fn = dir + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()

	bs, err = ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	size += len(bs)
	return fn[2:], bs, nil

}

func RandImage_bk() (string, int, int, []byte) {
	const dir = `C:\Go\_gopath\src\ms\sun\upload\samples`
	imageFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range imageFiles {
		_ = f
		//fmt.Println(f.Name())
	}
	//fn := "./upload/posts1" + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()
	fn := dir + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()
	file, err := os.Open(fn)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		//log.Fatal(err)
	}
	b := img.Bounds().Size()

	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	return fn[2:], b.X, b.Y, bs

}
