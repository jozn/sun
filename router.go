package main

import (
	//		"ms/sun/fact"
	. "ms/sun/actions"
	"ms/sun/fact"
	"net/http"
	//"github.com/drone/routes"
	"ms/sun/routes"
"ms/sun/sync"
	"ms/sun/ctrl"
)

func registerRoutes() {
	mux := routes.NewPrefix("/v1")

	//http.Handle("/posts", actioner(GetPostsAction))
	//http.Handle("/add-post", actioner(AddPostAction))
	//http.Handle("/likes", actioner(GetLikesAction))
	//http.Handle("/comments", actioner(GetCommentsAction))
	//http.Handle("/add-post-like", actioner(PostAddLikeAction))
	//http.Handle("/add-comment", actioner(PostAddCommentAction))

	http.Handle("/upload-avatar", actioner(UploadAvatarAction))
	http.Handle("/remove-avatar", actioner(RemoveAvatarAction))

	http.Handle("/profile", actioner(GetProfileAction))

	http.Handle("/search", actioner(SearchAllAction))
	http.Handle("/tag", actioner(SearchTagsAction))

	//http.Handle("/follow", actioner(FollowAction))
	//http.Handle("/unfollow", actioner(UnfollowAction))
	//http.Handle("/followers", actioner(GetFollowersAction))
	//http.Handle("/followings", actioner(GetFollowingsAction))

	//http.Handle("/sync-all-contacts", actioner(SyncAllContactsAction))
	mux.Post("/sync-all-contacts", actionToFunc(SyncAllContactsAction))
	mux.Post("/sync-followings", actionToFunc(ctrl.SyncFollowingsAction))
	mux.Get("/sync-followings", actionToFunc(ctrl.SyncFollowingsAction))

	// http.Handle("/2", actioner(h2))
	//	http.Handle("/hello", actioner(helo))
	//	http.Handle("/x", actioner(helo))
	//	http.Handle("/json1", actioner(playJsonAction))
	//	http.Handle("/test2", actioner(playTagsAction))
	//
	//	http.Handle("/play1", actioner(play1))
	//	http.Handle("/play2", actioner(play2))
	//	http.Handle("/play3", actioner(play3))
	//	http.Handle("/play4", actioner(play4))
	//	http.Handle("/play5", actioner(play5))
	//	http.Handle("/play6", actioner(play6))
	//	http.Handle("/redis1", actioner(redisAction1))
	//	http.Handle("/redis2", actioner(redisAction2))
	//	http.Handle("/panic", actioner(panicAction))
	//
	http.Handle("/fact/user1", actioner(fact.FactUser1))
	//	http.Handle("/fact/post1", actioner(factPost1))
	//	http.Handle("/fact/post2", actioner(factPost2))
	//	http.Handle("/fact/like1", actioner(factLike1))
		http.Handle("/fact/like2", actioner(fact.FactLike2))
		http.Handle("/fact/like3", actioner(fact.FactLike3))
	//	http.Handle("/fact/comment1", actioner(factComment1))
	http.Handle("/fact/follow1", actioner(fact.FactFollow1))
	http.Handle("/fact/follow2", actioner(fact.FactFollow2))
	// http.Handle("/fact/comment1", actioner(factC))
	 http.Handle("/fact/comment2", actioner(fact.FactComment2))

	//	http.Handle("/fact/p1", actioner(factPlay1))
	//	http.Handle("/fact/p2", actioner(factPlay2))
	//	http.Handle("/fact/p3", actioner(factPlay3))
	//	http.Handle("/fact/p4", actioner(factPlay4))

	http.Handle("/fact/user_real", actioner(fact.FactRealUser))

	// http.HandleFunc("/ws", serveWs)
	// http.HandleFunc("/ws2", serveWs2)
	// http.HandleFunc("/ws3", serveWs3)
	//	http.HandleFunc("/ws", serveWs2)

	//http.HandleFunc("/", Home)

	//http.Handle("/json1", actioner(TestJson1))
	//http.Handle("/json1", actioner(TestJson1))
	//mux.GET()
	mux.Get("/json2/:UserId",actionToFunc(TestJson1))
	mux.Del("/json2/:UserId",actionToFunc(TestJson1))

	http.Handle("/upload1", actioner(TestUpload1))
	http.Handle("/upload3", actioner(PlayUpload3))
	http.HandleFunc("/upload2", PlayUpload2)

	http.HandleFunc("/ws", sync.ServeUserWs)

	http.Handle("/mysql1", actioner(fact.IsamPlay))
	http.Handle("/f/msg", actioner(fact.ChatMsgFact1))
	http.Handle("/f/gm", actioner(fact.GroupMemFact1))
	http.Handle("/f/ginfo", actioner(fact.GroupInfoFact1))
	http.Handle("/f/1", actioner(fact.Run1))
	http.Handle("/f/2", actioner(fact.Run2))
	//	http.Handle("/f/all_chat", actioner(fact.ChatFact))

	//phone dbs
	http.Handle("/mf/contacts", actioner(fact.FactPhoneContacts))

	http.Handle("/ping", actionToFunc(ctrl.PingAction))

	http.Handle("/i/msg", actionToFunc(SendSampleMesgTable))
	http.Handle("/i/msg2", actionToFunc(SendSampleMesgTable2))
	http.Handle("/i/redis", actionToFunc(RedisSavePlay))
	http.Handle("/i/play", actionToFunc(PlaySomething))
	http.Handle("/i/store1", actionToFunc(MemoryStore1))
	http.Handle("/i/db", actionToFunc(DBStruct))
	http.Handle("/i/java2", actionToFunc(DBStructsTojava))
	http.HandleFunc("/i/java", DBStructsTojava2)

	//messages
	http.HandleFunc("/MsgUpload", MsgUpload)

	//New V1 apis
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


	http.Handle("/", mux)

}
