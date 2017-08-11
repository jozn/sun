package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//Note don't use "" to wrap commands in args string it dosn't work, just plain text
func main() {
	fmt.Println(os.Getwd())
	os.Chdir("../")
	fmt.Println(os.Getwd())

	cd := []string{
		`set GOOS=linux`,
		`go install`,
		`go build -o ./build/sun_linux main.go`,
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
