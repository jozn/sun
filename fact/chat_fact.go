package fact

import (
	. "ms/sun/base"
	. "ms/sun/models"
	//	"strings"
	//	"strconv"
	//	"time"
	//	"fmt"
	"math/rand"
)

func ChatMsgFact1(c *Action) {
	uid := rand.Intn(50)
	t := MessagesTable{}

	t.MessageKey = "u" + intToStr(uid) + intToStr(now())
	t.RoomKey = intToStr(uid)
	t.UserId = uid
	t.Text = factRnddStr(100)
	t.MessageTypeId = 1
	t.CreatedTimestampMs = now()

	DbInsertStruct(&t, "chat_message")
	c.SendJson(t)
}

func GroupMemFact1(c *Action) {
	uid := rand.Intn(50)
	t := GroupMember{}
	t.UserId = uid
	t.ByUserId = rand.Intn(50)
	t.RoleId = 2
	t.CreatedTimestamp = now()
	t.RoomKey = "g-" + intToStr(rand.Intn(10)+1)

	DbInsertStruct(&t, "chat_group_member")
	c.SendJson(t)

}

func GroupInfoFact1(c *Action) {
	uid := rand.Intn(50)
	t := GroupInfo{}

	t.CreatorUserId = uid
	t.GroupPrivacyId = 1
	t.MembersCount = rand.Intn(100)
	t.RoomName = factRnddStr(25)
	t.CreatedTimestamp = now()
	t.RoomKey = "g-" + intToStr(rand.Intn(10)+1)

	DbInsertStruct(&t, "chat_group_info")
	c.SendJson(t)

}
