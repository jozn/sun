package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println(os.Getwd())
	os.Chdir(`C:\Go\_gopath\src\ms\sun\scripts`)
	fmt.Println(os.Getwd())

	cd := []string{
		`java -jar ./google-java-format-1.3-all-deps.jar -r  D:\dev_working2\MS_Native\app\src\main\java\ir\ms\pb\* `,
		`java -jar ./google-java-format-1.3-all-deps.jar -r  D:\dev_working2\MS_Native\app\src\main\java\com\mardomsara\* `,
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
