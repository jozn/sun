package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//Note don't use "" to wrap commands in args string it dosn't work, just plain text
func main() {

	//cmd := exec.Command(`xox mysql://root:123456@localhost:3307/ms -o "../models/x" --package x --template-path C:\Go\_gopath\src\ms\xox\templates`)
	//args := []string{"xox", "mysql://root:123456@localhost:3307/ms" , "-o", "../models/x" , "--package"}
	args := strings.Split(`mysql://root:123456@localhost:3306/ms -o ../models/x --package x --template-path C:\Go\_gopath\src\ms\xox\templates`, " ")
	fmt.Println("")
	fmt.Println(strings.Join(args, " * "))
	cmd := exec.Command(`xox`, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

}
