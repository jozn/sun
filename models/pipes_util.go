package models

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"ms/sun/helper"
	"ms/sun/models/x"
	"strings"
	"time"
)

// room key format "p142_1569"
func RoomKeyToPeerUserId(RoomKey string, CurrentUserId int) (int, error) {
	if len(RoomKey) < 1 {
		return 0, errors.New("RoomKey wrong")
	}
	arr := strings.SplitN(RoomKey[1:], "_", -1)

	if len(arr) != 2 {
		return 0, errors.New("RoomKey wrong")
	}

	lowId := helper.StrToInt(arr[0], 0)
	highId := helper.StrToInt(arr[1], 0)

	if lowId == 0 || highId == 0 || lowId >= highId || lowId < 1 || highId < 1 {
		return 0, errors.New("RoomKey wrong")
	}

	if lowId == CurrentUserId {
		return highId, nil
	}
	return lowId, nil
}

//format "p142_1569"
func UserIdsToRoomKey(UserId1, UserId2 int) string {
	if UserId2 < UserId1 {
		UserId1, UserId2 = UserId2, UserId1
	}
	return fmt.Sprintf("p%d_%d", UserId1, UserId2)
}

//////////////////////////////////////////////////////////////
func NewPB_CommandToClient(cmd string) x.PB_CommandToClient {
	p := x.PB_CommandToClient{
		ServerCallId: int64(time.Now().UnixNano()),
		Command:      cmd,
	}
	return p
}

func NewPB_CommandToClient_WithData(cmd string, protoMsg proto.Message) x.PB_CommandToClient {
	m := NewPB_CommandToClient(cmd)
	bytes, err := proto.Marshal(protoMsg)
	if err == nil {
		m.Data = bytes
	} else {
		helper.DebugPrintln("ERROR : proto.Marshal NewPB_CommandToClient_WithData, ", err)
	}
	return m
}
