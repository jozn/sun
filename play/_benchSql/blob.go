package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"fmt"
	"github.com/jozn/protobuf/proto"
	"ms/sun/helper"
	"ms/sun/models/x"
)

var DB *sqlx.DB
var err error

func main() {
	fmt.Println("1")
	DB, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/ms?charset=utf8mb4")
	fmt.Println("2")

    DB.MapperFunc(func(s string) string { return s })

	helper.NoErr(err)

	d := &x.DirectUpdate{
		DirectUpdateId: helper.TimeNowMs(),
		ToUserId:       0,
		MessageId:      0,
		MessageFileId:  0,
		OtherId:        0,
		ChatKey:        "",
		PeerUserId:     0,
		EventType:      0,
		RoomLogTypeId:  0,
		FromSeq:        0,
		ToSeq:          0,
		ExtraPB:        []byte{},
		ExtraJson:      "",
		AtTimeMs:       0,
	}

	b, _ := proto.Marshal(&x.PB_Chat{
		ChatKey:    "sdsd",
		UserId:     1554545,
		CurrentSeq: 8784,
		RoomKey:    "asdas",
	})

	_ = b
	d.ExtraPB = b
	fmt.Println("%s", b)

	err = d.Save(DB)

	helper.NoErr(err)

	rows, err := x.NewDirectUpdate_Selector().DirectUpdateId_Eq(1509032130479).GetRows(DB)
	//rows, err := x.NewDirectUpdate_Selector().GetRows(DB)

    helper.NoErr(err)
	fmt.Println("size: ", len(rows))
	fmt.Println("size pb: ", len(rows[0].ExtraPB))
	fmt.Println("1=> ", rows[0])
	fmt.Println("1=> ", rows[0].DirectUpdateId)

	helper.PertyPrint(rows)

    rows2, err := x.NewDirectMessage_Selector().Limit(5).GetRows(DB)


    helper.PertyPrint(rows2)

	time.Sleep(time.Second * 100000)

}
