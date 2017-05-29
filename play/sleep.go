package main

import (
    "time"
    "fmt"
)

func main() {
    ch := make(chan int,100000)
    go func() {
        i:=0
        for {
            ch <- i
            i++
            time.Sleep(time.Microsecond * 1)
        }

    }()

    for o := range ch {
        if o%1000 == 0 {
            fmt.Println(o)
        }
    }

}
