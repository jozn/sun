package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func Alloc() uint64 {
	var stats runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats)
	return stats.Alloc - uint64(unsafe.Sizeof(hs[0]))*uint64(cap(hs))
}

var hs = []*map[int16]byte{}

func main() {
	hs := []*map[int16]byte{}
	n := 1000
	before := Alloc()
	for i := 0; i < n; i++ {
		h := map[int16]byte{}
		hs = append(hs, &h)
	}
	after := Alloc()
	emptyPerMap := float64(after-before) / float64(n)
	fmt.Printf("Bytes used for %d empty maps: %d, bytes/map %.1f\n", n, after-before, emptyPerMap)
	hs = nil

	k := 1
	for p := 1; p < 16; p++ {
		before = Alloc()
		for i := 0; i < n; i++ {
			h := map[int16]byte{}
			for j := 0; j < k; j++ {
				h[int16(j)] = byte(j)
			}
			hs = append(hs, &h)
		}
		after = Alloc()
		fullPerMap := float64(after-before) / float64(n)
		fmt.Printf("Bytes used for %d maps with %d entries: %d, bytes/map %.1f\n", n, k, after-before, fullPerMap)
		fmt.Printf("Bytes per entry %.1f\n", (fullPerMap-emptyPerMap)/float64(k))
		k *= 2
	}

}
