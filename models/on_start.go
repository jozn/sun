package models

func OnAppStart_Models() {

	UserMemoryStore.ReloadAll()

	ReloadTopUserIds()

	ReloadAllTags()
	ReloadTopTags()

	go PeriodicReloadTopPostsForTopTags()

	Cache = NewCache()
	CacheModels = _cacheModels{Cache}

	//todo: del this
	p := Post{}
	p.Text = "sadasdasd"

	Cache.Set("post_12", &p, 0)

}
