package helper

import "ms/sun/config"

//NOTE: proper use is defer helper.JustRecover()
func JustRecover() {
	if r := recover(); r != nil {
		if config.IS_DEBUG {
			DebugPrintln("ERROR PANICED RECOVRED - ERR:: ", r)
		}
	}
}

func SliceInt64ToInt(ins []int64) []int {
	res := make([]int, len(ins))
	for i := 0; i < len(ins); i++ {
		res[i] = int(ins[i])
	}
	return res
}

func SliceIntToInt64(ins []int) []int64 {
	res := make([]int64, len(ins))
	for i := 0; i < len(ins); i++ {
		res[i] = int64(ins[i])
	}
	return res
}

type R struct {
	Start, End int
}

func MaxPageLimit(max, page, limit int) (R, bool) {
	r := R{0, 0}
	if max < 0 || page < 0 || limit < 0 {
		return r, false
	}
	if page == 0 {
		page = 1
	}

	start := (page - 1) * limit
	end := start + limit

	if end > max {
		end = max
	}
	if end < start {
		return r, false
	}
	r.Start = start
	r.End = end
	return r, true
}
