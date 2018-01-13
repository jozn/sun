package galaxy

import (
	"fmt"
	"io/ioutil"
	"log"
    "math/rand"
)

var cnt int = 1
var size int = 0

func insert() {
	name, bs, err := RandImage()
	if err != nil {
		return
	}
	cas := &cassandraRow{
		Id:       cnt,
		Data:     bs,
		FileType: name,
	}
    cnt++
	putToCassandra(cas)
}

func Insert_many(num int) {
	for i := 0; i < num; i++ {
		insert()
		fmt.Println("cnt: ", i, " size:(mb)", size/1000000)
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
