package main

import (
	//		"ms/sun/fact"
	. "ms/sun/actions"
	"ms/sun/fact"
	"net/http"
	//"github.com/drone/routes"
	"ms/sun/ctrl"
	"ms/sun/pipes"
	"ms/sun/routes"
	"ms/sun/sync"
)

func registerRoutes() {
	mux := routes.NewPrefix("/v1")

	http.Handle("/upload-avatar", actioner(UploadAvatarAction))
	http.Handle("/remove-avatar", actioner(RemoveAvatarAction))

	//http.Handle("/profile", actioner(GetProfileAction))

	http.Handle("/search", actioner(SearchAllAction))
	http.Handle("/tag", actioner(SearchTagsAction))
	mux.Post("/sync-all-contacts", actionToFunc(SyncAllContactsAction))
	mux.Post("/sync-followings", actionToFunc(ctrl.SyncFollowingsAction))
	mux.Get("/sync-followings", actionToFunc(ctrl.SyncFollowingsAction))

	http.Handle("/fact/user1", actioner(fact.FactUser1))

	http.Handle("/fact/user_real", actioner(fact.FactRealUser))

	mux.Get("/json2/:UserId", actionToFunc(TestJson1))
	mux.Del("/json2/:UserId", actionToFunc(TestJson1))

	http.Handle("/upload1", actioner(TestUpload1))
	http.Handle("/upload3", actioner(PlayUpload3))
	http.HandleFunc("/upload2", PlayUpload2)

	http.HandleFunc("/ws1", sync.ServeUserWs) //Deprectaed
	http.HandleFunc("/ws_call", pipes.ServeHttpWs)

	http.Handle("/mysql1", actioner(fact.IsamPlay))
	http.Handle("/f/msg", actioner(fact.ChatMsgFact1))
	http.Handle("/f/gm", actioner(fact.GroupMemFact1))
	http.Handle("/f/ginfo", actioner(fact.GroupInfoFact1))

	//phone dbs
	http.Handle("/mf/contacts", actioner(fact.FactPhoneContacts))

	http.Handle("/ping", actionToFunc(ctrl.PingAction))

	http.Handle("/i/msg", actionToFunc(SendSampleMesgTable))
	http.Handle("/i/msg2", actionToFunc(SendSampleMesgTable2))
	http.Handle("/i/redis", actionToFunc(RedisSavePlay))
	http.Handle("/i/play", actionToFunc(PlaySomething))
	http.Handle("/i/store1", actionToFunc(MemoryStore1))
	http.Handle("/i/cache", actionToFunc(ShowCached))
	http.Handle("/i/db", actionToFunc(DBStruct))
	http.Handle("/i/java2", actionToFunc(DBStructsTojava))
	http.Handle("/i/table", actionToFunc(DBStructsToTable))
	http.HandleFunc("/i/java", DBStructsTojava2)

	////////////// New Facts from v0.4 /////////////////
	http.Handle("/fact/follow", actioner(fact.FactFollow))
	http.Handle("/fact/unfollow", actioner(fact.FactUnFollow))
	http.Handle("/fact/avatar", actioner(fact.FactUserAvatars))
	http.Handle("/fact/like", actioner(fact.FactLike))
	http.Handle("/fact/comment", actioner(fact.FactComment))
	http.Handle("/fact/post", actioner(fact.FactPost))

	//messages
	http.HandleFunc("/MsgUpload", MsgUpload)

	//////////////////// New V1 apis //////////////////////////
	//mux.Post("/grab_contacts", actionToFunc(ctrl.GrabAllUserContactsCtrl))
	//mux.Get("/grab_contacts", actionToFunc(ctrl.GrabAllUserContactsCtrl))
	http.Handle("/v1/grab_contacts", actionToFunc(ctrl.GrabAllUserContactsCtrl))

	mux.Post("/post/add", actionToFunc(ctrl.AddPostAction))
	mux.Get("/post/add", actionToFunc(ctrl.AddPostAction))
	mux.Get("/post/get", actionToFunc(ctrl.GetSinglePostAction))
	mux.Get("/post/stream", actionToFunc(ctrl.GetPostsStraemAction))
	mux.Get("/post/latest", actionToFunc(ctrl.GetPostsLatestAction))
	mux.Get("/post/delete", actionToFunc(ctrl.PostDeleteAction))
	mux.Get("/post/update", actionToFunc(ctrl.PostUpdateAction))

	mux.Get("/follow", actionToFunc(ctrl.FollowAction))
	mux.Get("/unfollow", actionToFunc(ctrl.UnfollowAction))
	mux.Get("/followers", actionToFunc(ctrl.GetFollowersListAction))
	mux.Get("/following", actionToFunc(ctrl.GetFollowingsListAction))

	//mux.Get("/likes", actionToFunc(ctrl.GetFollowingsListAction))
	mux.Get("/like", actionToFunc(ctrl.PostAddLikeAction))
	mux.Get("/unlike", actionToFunc(ctrl.PostRemoveLikeAction))
	mux.Get("/likes", actionToFunc(ctrl.GetLikesAction))

	mux.Get("/comments/add", actionToFunc(ctrl.PostAddCommentAction))
	mux.Post("/comments/add", actionToFunc(ctrl.PostAddCommentAction))
	mux.Get("/comments/list", actionToFunc(ctrl.GetCommentsAction))
	mux.Post("/comments/delete", actionToFunc(ctrl.RemoveCommentAction))

	mux.Post("/session/info", actionToFunc(ctrl.SessionGetUserInfo))
	mux.Get("/session/info", actionToFunc(ctrl.SessionGetUserInfo))

	mux.Get("/profile/all", actionToFunc(ctrl.GetPostsForProfileAction))

	mux.Get("/notifications", actionToFunc(ctrl.VNoftificationsCtrl))
	//mux.Get("notifications", actionToFunc(ctrl.NoftificationsCtrl))

	mux.Get("/recommend/top_posts", actionToFunc(ctrl.RecommendPostsCtrl))
	mux.Get("/recommend/users", actionToFunc(ctrl.RecommendUsersCtrl))
	mux.Get("/recommend/top_tags", actionToFunc(ctrl.RecommendTagsCtrl))
	mux.Get("/recommend/top_tags_discover", actionToFunc(ctrl.RecommendTagsWithPostsCtrl))

	mux.Get("/tags/list", actionToFunc(ctrl.TagsPostsListCtrl))
	mux.Get("/search", actionToFunc(ctrl.SearchCtrl))
	mux.Get("/search/tags", actionToFunc(ctrl.SearchTagsCtrl))
	mux.Get("/search/users", actionToFunc(ctrl.SearchUsersCtrl))

	mux.Get("/noti", actionToFunc(ctrl.NoftificationsCtrl2))

	mux.Get("/not", actionToFunc(ctrl.VNoftificationsCtrl))
	http.Handle("/not2", actioner(ctrl.VNoftificationsCtrl2))

	mux.Get("/sync_users", actionToFunc(ctrl.SyncUsersCtrl))

	http.Handle("/", mux)

}
