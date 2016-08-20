package main

import (
    "fmt"
    "time"
)
var arr = make([]byte,1000000000)
func Loop2() {
    l := len(arr)
    s:=0
    i:=0
    for ; i < l; i++ {
        if i < l {
            arr[i]= byte(i)
        }
        //j++
    }
    fmt.Println("Set: ",i, " ", s)

}
func Loop() {
    l := len(arr)
    s:=0
    i:=0
    for ; i < l; i++ {
         if i < l {
             s += int(arr[i])
         }
        //j++
    }
    fmt.Println("1 Billion: ",i, " ", s)

}


func main() {
    Loop2()
    t1:= time.Now().UnixNano()

    Loop()

    t2:= time.Now().UnixNano()
    fmt.Println("T: ", (t2-t1)/1e6)
}
