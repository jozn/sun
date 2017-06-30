package helper

import (
	"testing"
)
//50us
func BenchmarkParseText(b *testing.B) {
	t6 := `#hi #666 this is #one #we all like #At_Iran8 #من_را  888dsa sd

    for th e thing in #tag1

    jlk  #ایرانی
    5445
    #😫 #😩😤 😠 😡 😶 #ol😐
    `

	for i := 0; i < b.N; i++ {
		ParseText(t6)
	}
}
