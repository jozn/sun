package base

import (
	// "strings"
	"strconv"
)

var GlobTypes map[string]string

// GlobTypes := make(map[string]string) - basicaly it's an enum
func RegisterGlobTypes() {
	GlobTypes = map[string]string{
		//shared posts types
		"text":  "10",
		"10":    "text",
		"photo": "11",
		"11":    "photo",
		"video": "12",
		"12":    "video",
	}

}

func TypeIdToName(id int) string {
	s := GlobTypes[intToStr(id)]
	return s
}

func TypeNameToId(typ string) int {
	s := GlobTypes[typ]
	i, _ := strconv.Atoi(s)
	return i
}

func intToStr(i int) string {
	return strconv.Itoa(i)
}
