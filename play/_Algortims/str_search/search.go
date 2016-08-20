package main

import (
    "fmt"
    "strconv"
    "strings"
)

func Search(src, p string) int{
    l := len(src)
    lp := len(p)
    for i:=0 ; i < l ; i++{
        j:= i
        f := true
        for n:=0 ; n < lp && j < l ; n++  {
            if src[j] != p[n]{
                f = false
                break
            }
            j++
        }

        if f {
            return i
        }
    }
    return -1
}


func main() {
    s := "They are both theoretically dictionary coders. LZ77 maintains a sliding window during compression. This was later shown to be equivalent to the explicit dictionary constructed by LZ78—however, they are only equivalent when the entire data is intended to be decompressed. LZ78 decompression allows random access to the input as long as the entire dictionary is available,[dubious – discuss] while LZ77 decompression must always start at the beginning of the input."
    fmt.Println("Hello, playground")
    fmt.Println( Search(s,"always") )
    strings.Join()
}
