package models

func init() {

	go impleOfPeriodaclyUpdateLastActivityOfUser()

}

func OnAppStart_Models() {

	UserMemoryStore.ReloadAll()
	MemoryStore_User.ReloadAll()

	ReloadTopUserIds()

	Tags_RepeatedlyJobs()

	go PeriodicReloadTopPostsForTopTags()
	go Session_periodicllyUpdateDB()

	Cache = NewCache()
	CacheModels = _cacheModels{Cache}

}
