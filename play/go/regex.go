package main

import (
    "regexp"
    "fmt"
)

func main()  {
    r := regexp.MustCompile(`\w+_?\.\w`)
    f :=r.MatchString("6511.jp")
    r.
    fmt.Println(f)
}