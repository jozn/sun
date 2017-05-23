package main

import (
    "os/exec"
    "fmt"
    "os"
    "strings"
)

//Note don't use "" to wrap commands in args string it dosn't work, just plain text
func main()  {
    fmt.Println(os.Getwd())
    os.Chdir("../models/protos/")
    fmt.Println(os.Getwd())

    cd := []string{
        `protoc --go_out=plugins=grpc:../x -I ./ *.proto`,
        `protoc --java_out=D:\dev_working2\MS_Native\app\src\main\java --grpc-java_out=D:\dev_working2\MS_Native\app\src\main\java *.proto`,
        `goimports -w ../x/`,
    }

    for _,c := range cd {
        args := strings.Split(c," ")
        fmt.Println(strings.Join(args, " * "))
        cmd := exec.Command(args[0], args[1:]...)
        cmd.Stdin = os.Stdin
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout
        err := cmd.Run()
        if err != nil{
            fmt.Errorf("Error: %s",err)
        }
    }

}