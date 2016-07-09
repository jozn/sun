package actions

import (
    "ms/sun/base"
    "ms/sun/models"
)


func MemoryStore1(a *base.Action) base.AppErr {
    s := models.UserMemoryStore
    s.ReloadAll()
    //a.SendJson(s.Map[1])
    m := make(map[int]interface{})
    for j:=0 ;j<10;j++ {
        s.ReloadAll()
        m[j] = s.Map
    }

    a.SendJson(s.GetAllAsArray())

    return nil
}

func DBStruct(a *base.Action) base.AppErr   {
    s := base.DbStructToTable2New(&models.UserTable{}, "user")
    a.SendJson(s)
    return nil
}
