package actions

import (
    "ms/sun/base"
    "ms/sun/models"
    "ms/sun/helper"
    "net/http"
    "fmt"
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

func DBStructsTojava(a *base.Action) base.AppErr   {
    s := helper.DbStructToJava(models.UserInlineView{})
    a.SendJson(s)
    return nil
}

func DBStructsTojava2(w http.ResponseWriter, r *http.Request)  {
    str:=""
    str += helper.DbStructToJava(models.UserInlineView{})
    //str += helper.DbStructToJava(models.CommentInlineInfo{})
    //str += helper.DbStructToJava(models.PostAndDetailes{})

    fmt.Fprint(w,str)
}

