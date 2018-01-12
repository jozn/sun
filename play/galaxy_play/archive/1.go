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
	"time"
)

type Media struct {
	Id   int    `db:"id"`
	Key  string `db:"key"`
	Load []byte `db:"load"`
}

var db *sqlx.DB
var err error

func main() {
	db, err = sqlx.Open("postgres", "postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable")

	if err != nil {
		panic(err)
	}
	/*rows, err := db.Query("select * from galaxy.msg_files")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		fmt.Println("22")
	}

	var meds []Media

	err = db.Select(&meds, "select * from galaxy.msg_files")

	if err != nil {
		panic(err)
	}
	fmt.Println(meds)*/

	//insert()
	insert_many()
}

func insert() {
	name, bs, err := RandImage()
	if err != nil {
		return
	}
	m := Media{
		Id:   int(time.Now().UnixNano()),
		Key:  name,
		Load: bs,
	}
	/*m := Media{
		Id:   12,
		Key:  "one_key",
		Load: []byte("gsdgs"),
	}*/
	db.MustExec("INSERT into galaxy.msg_files (id,key,load) VALUES ($1,$2,$3)", m.Id, m.Key, m.Load)
}

func insert_many() {
	for i := 0; i < 1000; i++ {
		insert()
		fmt.Println("cnt: ", i)
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
