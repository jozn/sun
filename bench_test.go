package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"ms/sun/models/x"
	"testing"
	"time"
)

func BenchmarkPushView_directLogsTo_PB_ChangesHolderView(b *testing.B) {
	go prepareTest()
	time.Sleep(time.Second * 2)
	rows, err := x.NewDirectLog_Selector().ToUserId_Eq(6).GetRows(base.DB)
	helper.NoErr(err)

	v := models.PushView_directLogsTo_PB_ChangesHolderView(6, rows)
	bs, _ := proto.Marshal(v)
	fmt.Println("size :", len(bs))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		models.PushView_directLogsTo_PB_ChangesHolderView(6, rows)
	}
}

func BenchmarkPushView_directLogsTo_PB_ChangesHolderView2(b *testing.B) {
	go prepareTest()
	time.Sleep(time.Second * 2)
	rows, err := x.NewDirectLog_Selector().ToUserId_Eq(6).GetRows(base.DB)
	helper.NoErr(err)

	models.PushView_directLogsTo_PB_ChangesHolderView(6, rows)

	b.ResetTimer()

	view := models.PushView_directLogsTo_PB_ChangesHolderView(6, rows)
	for i := 0; i < b.N; i++ {
		proto.Marshal(view)
	}
}

func BenchmarkPushView_directLogsTo_PB_ChangesHolderView3(b *testing.B) {
	go prepareTest()
	time.Sleep(time.Second * 1)
	rows, err := x.NewDirectLog_Selector().ToUserId_Eq(6).GetRows(base.DB)
	helper.NoErr(err)

	models.PushView_directLogsTo_PB_ChangesHolderView(6, rows)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		view := models.PushView_directLogsTo_PB_ChangesHolderView(6, rows)
		proto.Marshal(view)
	}
}

func BenchmarkProtoBuff(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p := &x.PB_ChatView{
			ChatKey:              "dasdasda",
			ChatId:               194965,
			RoomTypeEnumId:       94965,
			UserId:               94965,
			PeerUserId:           94965,
			GroupId:              94965,
			CreatedTime:          58,
			UpdatedMs:            94965,
			LastMessageId:        94965,
			LastDeletedMessageId: 94965,
			LastSeenMessageId:    787,
			LastSeqSeen:          94965,
			LastSeqDelete:        94965,
			CurrentSeq:           94965,
		}

		proto.Marshal(p)
	}
}
