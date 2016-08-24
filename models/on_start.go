package models

func OnAppStart_Models() {

    UserMemoryStore.ReloadAll()

    ReloadTopUserIds()

    ReloadAllTags()
    ReloadTopTags()

    go PeriodicReloadTopPostsForTopTags()

}


