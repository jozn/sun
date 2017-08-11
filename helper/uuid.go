package helper

import (
	"fmt"
	"sync"
)

//fixme:: cosider times when clock has come back (in those situations we must add another serverid )
var defGen = &uuidGen{
	timeMs:   TimeNowMs(),
	counter:  1,
	serverId: 1,
}

type uuidGen struct {
	sync.RWMutex
	timeMs   int
	counter  int
	serverId int
}

//must be an 19digit - max 64 bit signed is 9223372036854775807 ____  2*63 - 1
func (gen *uuidGen) next() int {
	now := TimeNowMs()
	gen.Lock()
	if gen.timeMs < now {
		gen.timeMs = now
		//gen.counter = 1
	}

	if gen.counter >= 9999 {
		gen.counter = 0
	}

	// "13time milliscon + 2 serverId + 4 random  == 19digits
	t := gen.timeMs
	id := gen.serverId
	cnt := gen.counter
	gen.counter++
	gen.Unlock() //no defer this is faster fmt.Sprintf took 700us

	s := fmt.Sprintf("%d%02d%04d", t, id, cnt) //19
	i := StrToInt(s, 0)
	return i
}

func NextRowsSeqId() int {
	return defGen.next()
}
