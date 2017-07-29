package main

import (
    "github.com/syncthing/syncthing/lib/sync"
    "fmt"
    "time"
)

func main()  {
    g:=sync.NewWaitGroup()
    t1:=time.Now()
    N:= 1000000
    g.Add(N)
    for i:=0;i<N ; i++  {
        go func() {
            g.Done()
        }()
    }

    t2 := time.Now()
    g.Wait()
    fmt.Println("done ", t2.Sub(t1).Nanoseconds()/1e6)
}
