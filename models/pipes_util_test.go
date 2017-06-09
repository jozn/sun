package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"testing"
    "fmt"
)

func TestRoomKeyToPeerUserId(t *testing.T) {
	rs := []struct {
		key    string
		meId   int
		peerId int
		ok     bool
	}{
		{"p6_46", 6, 46, true},
		{"p6_46", 46, 6, true},
		{"p61_62", 61, 62, true},

		{"p61_61", 61, 62, false},
		{"p61_46", 61, 46, false}, //lower uid must be first
		{"pp6_46", 46, 6, false},
		{"pp6-4656", 46, 6, false},
		{"64656", 46, 6, false},
		{"ss878_sss_sxs6", 46, 6, false},
		{"lkjlkkljlkj", 46, 6, false},
	}

	for _, r := range rs {
		pid, err := RoomKeyToPeerUserId(r.key, r.meId)
		if r.ok {
			if err != nil || pid != r.peerId {
				t.Error("RoomKeyToPeerUserId err ", err, r, pid)
			}
		} else {
			if err == nil {
				t.Error("sss", err, r.key, r.peerId, " ", pid)
			}
		}

	}
}

func BenchmarkRoomKeyToPeerUserId(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RoomKeyToPeerUserId("p61_146", 61)
	}
}

func BenchmarkXox(b *testing.B) {
	b.SkipNow()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x.NewFollowingList_Selector().GetRows(base.DB)
	}
}

func BenchmarkMapInt(b *testing.B) {
	b.ReportAllocs()
	mp := make(map[int]*string)

	//s:= ""

	for i := 0; i < b.N; i++ {
		//mp[i] = helper.IntToStr(i)
		s := helper.IntToStr(i)

		mp[-i] = &s
	}
}

func ExampleHelperIntToStr() {
    s := helper.IntToStr(5)
    fmt.Println(s)

    // Output:
    // 5
}
