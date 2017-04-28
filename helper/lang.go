package helper

import "ms/sun/config"

func JustRecover() {
	if r := recover(); r != nil {
		if config.IS_DEBUG {
			DebugPrintln("ERROR PANICED RECOVRED - ERR:: ", r)
		}
	}
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
