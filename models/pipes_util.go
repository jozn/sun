package models

import (
	"errors"
	"ms/sun/helper"
	"strings"
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
