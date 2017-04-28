package ctrl

import (
	"math/rand"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

func InsertActivity(c *base.Action) base.AppErr {

	for i := 0; i < 10000; i++ {
		arr := make([]x.Activity, 0, 1000)
		for i := 0; i < 1000; i++ {
			arr = append(arr, x.Activity{
				ActorUserId:  25,
				ActionTypeId: 592,
				TargetId:     rand.Intn(10000000),
				RefId:        rand.Intn(10000000),
				CreatedAt:    helper.TimeNow(),
			})
		}
		x.MassInsert_Activity(arr, base.DB)
	}

	return nil
}
