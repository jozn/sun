package main

import (
	"fmt"
	"math/rand"
	"ms/sun/helper"
	"time"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, helper.TimeNowMs(), helper.TimeNowMs() * 5}

	for _, n := range ints {
		fmt.Printf("%b \n", n)
	}

	fmt.Println(len("1011001010000110101111000101111"))
	fmt.Println(len("10101110010101111001010001000001111111100"))

	fmt.Println("", "", "", "", "")

	s := helper.TimeNowMs()
	s2 := s << 23
	fmt.Printf("%64b \n", s)
	fmt.Printf("%d \n", s)

	fmt.Printf("%64b \n", s2)
	fmt.Printf("%d \n", s2)

	fmt.Printf("\n\n\n\n\n\n\n\n\n")

	m := 1
	m = m << 62
	ints2 := []int{helper.TimeNow(), helper.TimeNowMs(), m}

	for _, n := range ints2 {
		fmt.Printf("%d \n", n)
	}

	m2 := uint64(7)
	m2 = m2 << 63
	fmt.Println(len(fmt.Sprintf("%d", m)))
	fmt.Println(len(fmt.Sprintf("%d", m2)))

	t := helper.TimeNowNano()
	fmt.Printf("\n\n\n\n")
	fmt.Printf("%d \n", t)
	fmt.Println(len(fmt.Sprintf("%d", t)))

	for n := 0; n < 100; n++ {
		t := helper.TimeNowMs() //13
		id := 1                 //2
		rnd := rand.Intn(10000) //4
		_ = rnd

		s := fmt.Sprintf("%d%02d%04d", t, id, n) //19

		si := helper.StrToInt(s, 0)
		fmt.Println(si, len(s))
	}

	fmt.Println(len("9223372036854775807"))

	for {
		time.Sleep(time.Millisecond)
		fmt.Println(next())
	}
}

//must be an 19digit - max 64 bit signed is 9223372036854775807 ____  2*63 - 1

func next() string {
	SERVER_ID := 1              //2
	t := helper.TimeNowMs()     //13
	rnd := rand.Intn(10000 - 1) //4
	_ = rnd

	// "13time milliscon + 2 serverId + 4 random
	s := fmt.Sprintf("%d%02d%04d", t, SERVER_ID, rnd) //19
	return s
}

func next_old() {
	/*d := time.Date(2017,0,0,0,0,0,0, time.UTC)
	  fmt.*/
	time.Unix(1000000000, 0)

}
