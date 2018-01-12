package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//Note don't use "" to wrap commands in args string it dosn't work, just plain text

//see java-grpc : https://github.com/grpc/grpc-java/tree/master/compiler

func main() {
	fmt.Println(os.Getwd())
	os.Chdir("../models/protos/")
	fmt.Println(os.Getwd())

	cd := []string{
		`ants.exe -alaki`,
		`protoc --go_out=plugins=grpc:../x -I ./ *.proto`,
		// `protoc --java_out=D:\dev_working2\MS_Native\app\src\main\java --grpc-java_out=D:\dev_working2\MS_Native\app\src\main\java *.proto`,
		// `protoc --plugin=protoc-gen-grpc-java=C:\Go\_gopath\src\ms\sun\scripts\protoc-gen-grpc-java-1.4.0-windows-x86_64.exe --java_out=D:\dev_working2\MS_Native\proto\src\main\java --grpc-java_out=lite:D:\dev_working2\MS_Native\proto\src\main\java *.proto`,
		`protoc --plugin=protoc-gen-grpc-java=C:\Go\_gopath\src\ms\sun\scripts\protoc-gen-grpc-java-1.4.0-windows-x86_64.exe --java_out=D:\dev_working2\social\proto\src\main\java --grpc-java_out=lite:D:\dev_working2\social\proto\src\main\java *.proto`,
		// `protoc --grpc-java_out=lite:=D:\dev_working2\MS_Native\app\src\main\java --grpc-java_out=D:\dev_working2\MS_Native\app\src\main\java *.proto`,//java jsut grpc
		`goimports -w ../x/`,
		`javafmt -alaki`,
	}

	for _, c := range cd {
		args := strings.Split(c, " ")
		fmt.Println(strings.Join(args, " * "))
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Errorf("Error: %s", err)
		}
	}

}
