package main

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
