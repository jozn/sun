package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

func main() {

	DB, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms?charset=utf8mb4")

	if err != nil {
		log.Panic(err)
	}
	sql := "INSERT into test_chat (`Id`, `TimeMs`,`Text` ,`Name`, `UserId`, `c2`, `c3`, `c4`, `c5`) VALUES (?,?,?,?,?,?,?,?,?) "
	//(`Id`, `TimeMs`, `Name`, `UserId`, `c2`, `c3`, `c4`, `c5`) values ('1', '5445', 'sad', '25', '1', '2', '6', '5')
	_, err = DB.Prepare(sql)

	go func() {
		for i := 0; i < 1; i++ {
			/*p := &x.PostCopy{
				Id:             0,
				UserId:         rand.Intn(50000),
				TypeId:         0,
				Text:           helper.FactRandStr(200),
				FormatedText:   "",
				MediaCount:     0,
				SharedTo:       0,
				DisableComment: 20,
				HasTag:         50,
				LikesCount:     10,
				CommentsCount:  0,
				EditedTime:     1110,
				CreatedTime:    helper.TimeNow(),
			}*/
			p := x.LikesCopy{
				Id:          0,
				PostId:      rand.Intn(50000),
				PostTypeId:  1,
				UserId:      rand.Intn(50000),
				TypeId:      2,
				CreatedTime: helper.TimeNow(),
			}

			err = p.Insert(DB)

			//helper.NoErr(err)

			//DB.Exec("SELECT * FROM likes AS l JOIN `user` AS u WHERE l.UserId = u.Id AND l.PostId < 100 LIMIT 100 ;")
			//DB.Exec("SELECT * FROM likes LIMIT 100;")
			//time.Sleep(time.Second * 20)
			if i%100 == 0 {
				fmt.Println(i)
			}
		}
	}()

	go func() {

		c := 0
		t1 := time.Now()
		t2 := time.Now()
		for i := 0; i < 1000000; i++ {
			/*arr := make([]x.PostCopy, 100)
			for j := 0; j < 100; j++ {
				p := x.PostCopy{
					Id:             0,
					UserId:         rand.Intn(50000),
					TypeId:         0,
					Text:           helper.FactRandStr(200),
					FormatedText:   "",
					MediaCount:     0,
					SharedTo:       0,
					DisableComment: 20,
					HasTag:         50,
					LikesCount:     10,
					CommentsCount:  0,
					EditedTime:     1110,
					CreatedTime:    helper.TimeNow(),
				}

				arr[j] = p
			}

						x.MassInsert_PostCopy(arr, DB)

			*/

			const N = 500
			arr := make([]x.LikesCopy, N)
			for j := 0; j < N; j++ {
				p := x.LikesCopy{
					Id:          0,
					PostId:      rand.Intn(1000000),
					PostTypeId:  1,
					UserId:      rand.Intn(100000),
					TypeId:      2,
					CreatedTime: helper.TimeNow(),
				}

				arr[j] = p
			}
			c = c + N
			x.MassInsert_LikesCopy(arr, DB)

			//DB.Exec("SELECT * FROM likes AS l JOIN `user` AS u WHERE l.UserId = u.Id AND l.PostId < 100 LIMIT 100 ;")
			//DB.Exec("SELECT * FROM likes LIMIT 100;")
			//time.Sleep(time.Second * 20)
			if i%100 == 0 {
				t2 = time.Now()
				tt := t2.Sub(t1)
				t2.Nanosecond()
				o := c / (int(tt.Nanoseconds()) / 1e6)
				fmt.Println("M ", i, " ", o)
				t1 = time.Now()
				c = 0
			}
		}
	}()
	if err != nil {
		log.Panic(err)
	}

	time.Sleep(time.Hour * 10)

}
