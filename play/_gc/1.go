//Uses 2-3 Gig of ram
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

type User struct {
	id   int
	name string
	u    *User
}
type MM struct {
	sync.RWMutex
	m map[int]User
}

var uuu *User
var str string

func main() {
	debug.SetGCPercent(500)

	mm := MM{}
	mm.m = make(map[int]User)
	N := 100000
	index := 0
	deUser := User{}
	deUser2 := &User{}
	deUser.u = deUser2

	lenMM := func() int {
		mm.Lock()
		l := len(mm.m)
		mm.Unlock()
		return l
	}

	doned := make(chan bool)
	//real refrence
	go func() {
		for {
			for i := 0; i < N*100; i++ {
				index++
				mm.Lock()
				mm.m[index] = User{id: i, name: strconv.Itoa(rand.Intn(5000000)), u: &deUser}
				mm.Unlock()
			}
			for j := 0; j < N; j++ {
				r := rand.Intn(lenMM())
				r2 := rand.Intn(lenMM())
				mm.Lock()
				us, ok := mm.m[r]
				us2, ok2 := mm.m[r2]
				mm.Unlock()
				if ok && ok2 {
					us.u = &us2
				}

				//def user sequnce refs
				u5 := User{}
				u5.name = "asdasdasd" + strconv.Itoa(rand.Intn(5000000))
				deUser2.u = &u5
				deUser2 = &u5
			}
			time.Sleep(1 * time.Second)
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++=")
			break
		}
		//doned <- true

	}()

	go func() {
		<-doned
		for {
			for i := 0; i < 1000; i++ {
				index++
				mm.Lock()
				mm.m[index] = User{id: i, name: strconv.Itoa(rand.Intn(5000000)), u: &deUser}
				mm.Unlock()
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	go func() {
		//<- doned
		for {
			for i := 0; i < 10; i++ {
				r := rand.Intn(index)
				mm.Lock()
				delete(mm.m, r)
				mm.Unlock()
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	//////////////////////////////////////
	//garbege producer
	g := 0
	for k := 0; k < 10; k++ {
		go func() {
			for {
				for i := 0; i < 20000; i++ {
					u := User{}
					u.name = "lkaskjas"
					uuu = &u
					//str = "nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa"
				}
				g++
				time.Sleep(10 * time.Microsecond)
			}
		}()
		time.Sleep(40 * time.Microsecond)
	}

	go func() {
		for {
			for i := 0; i < N; i++ {
				str = "nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa"
			}
			time.Sleep(10 * time.Microsecond)
		}
	}()
	/////////////////////////////////////////

	/////loger
	z := 0
	go func() {
		for {
			z++
			fmt.Println("\n\n")
			fmt.Println("N: ", z)
			fmt.Println("g: ", g)
			fmt.Println("Map Len: ", lenMM())
			printMem()
			time.Sleep(5 * time.Second)
		}
	}()
	time.Sleep(1000 * time.Second)
}

var t1 int64

func init() {
	t1 = time.Now().Unix()
}

func printMem() {

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)

	fmt.Println("Alloc: ", mem.Alloc)
	fmt.Println("EnableGC: ", mem.EnableGC)
	fmt.Println("Frees: ", mem.Frees)
	fmt.Println("HeapAlloc: ", mem.HeapAlloc)
	fmt.Println("HeapIdle: ", mem.HeapIdle)
	fmt.Println("HeapObjects: ", mem.HeapObjects)
	fmt.Println("HeapReleased: ", mem.HeapReleased)
	fmt.Println("TotalAlloc: ", mem.TotalAlloc)
	fmt.Println("StackInuse: ", mem.StackInuse)
	fmt.Println("Lookups: ", mem.Lookups)

	fmt.Println("LastGC: ", mem.LastGC/1e9)
	fmt.Println("NumGC: ", mem.NumGC)
	fmt.Println("NextGC: ", mem.NextGC/1e9)
	fmt.Println("Mallocs: ", mem.Mallocs)
	fmt.Println("PauseTotalNs: ", (mem.PauseTotalNs / 1e6))

	fmt.Println("GCCPUFraction: ", mem.GCCPUFraction)
	fmt.Println("GCSys: ", mem.GCSys)

	fmt.Println("Time: ", (time.Now().Unix() - t1))
	fmt.Println("Now: ", (time.Now().Unix()))

}
