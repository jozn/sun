package models

func OnAppStart_Models() {

	MemoryStore_User.ReloadAll()

	ReloadTopUserIds()

	Tags_RepeatedlyJobs()

	go Session_periodicllyUpdateDB()

	/*Cache = NewCache()
	CacheModels = _cacheModels{Cache}*/

}
