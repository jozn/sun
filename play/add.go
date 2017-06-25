package main

import (
    "time"
    "fmt"
)

func main()  {
    t1 := time.Now()
    sum := 0
    for i:=0 ; i< 1000000000 ; i++ {
        sum += i
    }

    fmt.Printf("%d - %d ", time.Now().Sub(t1).Nanoseconds() / 1000000 ,sum)
}
