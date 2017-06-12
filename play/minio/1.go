// Install Minio library.
//
// $ go get -u github.com/minio/minio-go
//
package main

import (
	"fmt"
	"github.com/minio/minio-go" // Import Minio library.
	"log"
)

func main() {
	// Use a secure connection.
	ssl := false

	// Initialize minio client object.
	minioClient, err := minio.New("127.0.0.1:9000", "JRUUBXOVQ6EYD77XNMGP", "Z39RhJSQo1oxp3QUfDU1M0u714CUUqzw6I4jDnXy", ssl)
	if err != nil {
		log.Fatalln(err)
	}

	// Creates bucket with name mybucket.
	bu := "my/buc/ket5"
	err = minioClient.MakeBucket(bu, "us-east-1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully created mybucket.")

	// Upload an object 'myobject.txt' with contents from '/home/joe/myfilename.txt'
	n, err := minioClient.FPutObject(bu, "myobject.2txt", "./1.go", "application/text")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Uploaded myobject.txt of size %d\n", n)
}
