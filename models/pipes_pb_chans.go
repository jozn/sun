package models

import (
	"ms/sun/models/x"
	"time"
)

var newChatMsgsChanelBuffer = make(chan newChatMsgBuffer, 10000)

type newChatMsgBuffer struct {
	msg        *x.PB_Message
	fromUserId int
	toUserId   int
	roomKey    string
	hashId     int
}

func batcheNewMsgsBufferProceess() {
	const siz = 1000
	arr := make([]newChatMsgBuffer, 0, siz)
	go func() {
		for m := range newChatMsgsChanelBuffer {
			arr = append(arr, m)
		}
	}()

	for {
		time.Sleep(time.Microsecond * 5)
		if len(arr) > 0 {
			pre := arr
			arr = make([]newChatMsgBuffer, 0, siz)
			processNewChatMsgBuffer(pre)
		}
	}
}

func processNewChatMsgBuffer(msgsBuff []newChatMsgBuffer) {
	mp := make(map[int][]newChatMsgBuffer)

	for _, mb := range msgsBuff {
		mp[mb.toUserId] = append(mp[mb.toUserId], mb)
	}

}
