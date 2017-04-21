package actions

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"net/http"
)

func MemoryStore1(a *base.Action) base.AppErr {
	s := models.UserMemoryStore
	s.ReloadAll()
	//a.SendJson(s.Map[1])
	m := make(map[int]interface{})
	for j := 0; j < 10; j++ {
		s.ReloadAll()
		m[j] = s.Map
	}

	a.SendJson(s.GetAllAsArray())

	return nil
}

func DBStruct(a *base.Action) base.AppErr {
	s := base.DbStructToTable2New(&models.UserTable{}, "user")
	a.SendJson(s)
	return nil
}

func DBStructsTojava(a *base.Action) base.AppErr {
	s := helper.DbStructToJava(models.UserInlineView{})
	a.SendJson(s)
	return nil
}

func DBStructsTojava2(w http.ResponseWriter, r *http.Request) {
	str := ""
	str += helper.DbStructToJava(models.UserInlineView{})
	//str += helper.DbStructToJava(models.CommentInlineInfo{})
	//str += helper.DbStructToJava(models.PostAndDetailes{})

	fmt.Fprint(w, str)
}

func DBStructsToTable(a *base.Action) base.AppErr {
	str := ""
	str += base.DbStructToTable(&models.Notification{}, "notification")

	a.SendJson(str)
	return nil
}

func ShowCached(a *base.Action) base.AppErr {

	a.SendJson(models.Cache.Items())
	return nil
}

func ShowCacher(a *base.Action) base.AppErr {

	a.SendJson(models.Cacher.Items())
	return nil
}

func ShowCacheKeys(a *base.Action) base.AppErr {

	a.SendJson(models.RowCache.Items())
	return nil
}

func ShowCacheRowKeys(a *base.Action) base.AppErr {
	var arr []string
	m := models.RowCache.Items()
	for k, _ := range m {
		arr = append(arr, k)
	}
	a.SendJson(arr)
	return nil
}

func ShowCacheIndex(a *base.Action) base.AppErr {
    /*if rand.Intn(3) == 2 {
        i:=rand.Intn(1000000000000)
        e:=i+1000000
        for ;i<e;i++ {
            models.RowCacheIndex.Set("kjslkjlkjlkjlkjljlj:"+strconv.Itoa(i),&models.Photo{},0)
        }
    }

    for i:=0;i<100000;i++ {
        models.RowCacheIndex.Get("kjslkjlkjlkjlkjljlj:"+strconv.Itoa(i))
    }
	a.SendJson(len(models.RowCacheIndex.Items()))*/
    a.SendJson(models.RowCacheIndex.Items())
	return nil
}
