package models

import (
	"ms/sun/helper"
	"time"
)

var _bufferOfSessionLastActivity = make([]int, 0, 10000)

func PeriodaclyUpdateLastActivityOfUser(UserId int) {
	_bufferOfSessionLastActivity = append(_bufferOfSessionLastActivity, UserId)
}

func impleOfPeriodaclyUpdateLastActivityOfUser() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	for {
		time.Sleep(time.Second * 1)
		arr := _bufferOfSessionLastActivity
		_bufferOfSessionLastActivity = make([]int, 0, 10000)
		QueryUpdateSessionLastActivitiesOfUsers(arr)
	}
}

func PeriodicReloadTopPostsForTopTags() {
	defer helper.JustRecover()

	for {
		ReloadTopPostsForTopTags()
		time.Sleep(time.Second * 60)
	}
}
