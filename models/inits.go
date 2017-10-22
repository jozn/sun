package models

import "fmt"

func OnAppStart_Models() {
    fmt.Println("OnAppStart_Models")
	MemoryStore_User.ReloadAll()

	ReloadTopUserIds()

	Tags_RepeatedlyJobs()
	Recommend_Jobs()

	go Session_periodicllyUpdateDB()

	/*Cache = NewCache()
	CacheModels = _cacheModels{Cache}*/

}
