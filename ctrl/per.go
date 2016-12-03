package ctrl

import (
    "ms/sun/base"
    "ms/sun/models"
    "ms/sun/helper"
    "math/rand"
)

func InsertActivity(c *base.Action) base.AppErr  {

    for i:=0; i< 10000; i++ {
        arr := make([]models.Activity,0,1000)
        for i:=0; i< 1000; i++ {
            arr = append(arr, models.Activity{
                ActorUserId:25,
                ActionTypeId: 592,
                TargetId: rand.Intn(10000000),
                RefId: rand.Intn(10000000),
                CreatedAt: helper.TimeNow(),
            })
        }
        models.MassInsert_Activity(arr,base.DB)
    }

    return nil
}
