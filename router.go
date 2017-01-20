package main

import (
	. "ms/sun/actions"
	"ms/sun/ctrl"
	"ms/sun/fact"
	"ms/sun/models"
	"net/http"
	//"ms/sun/routes"
	"fmt"
	"github.com/dimfeld/httptreemux"
	"ms/sun/base"
	"ms/sun/config"
)

func registerRoutes() *httptreemux.TreeMux {
	v1Tree := httptreemux.New()
	v1Tree.PanicHandler = func(w http.ResponseWriter, r *http.Request, p interface{}) {
		fmt.Println("PANIC: ", p)
		w.Write([]byte(fmt.Sprintf("%v", p)))
	}
	v1 := v1Tree.NewGroup("/v1")

	//////////////////// New V1 apis //////////////////////////
	http.Handle("/v1/grab_contacts", actionToFunc(ctrl.GrabAllUserContactsCtrl))

	v1.POST("/post/add", toV1(ctrl.AddPostAction))
	v1.GET("/post/add", toV1(ctrl.AddPostAction))
	v1.GET("/post/get", toV1(ctrl.GetSinglePostAction))
	v1.GET("/post/stream", toV1(ctrl.GetPostsStraemAction))
	v1.GET("/post/latest", toV1(ctrl.GetPostsLatestAction))
	v1.GET("/post/delete", toV1(ctrl.PostDeleteAction))
	v1.GET("/post/update", toV1(ctrl.PostUpdateAction))

	v1.POST("/follow", toV1(ctrl.FollowAction))
	v1.POST("/unfollow", toV1(ctrl.UnfollowAction))
	v1.GET("/followers", toV1(ctrl.GetFollowersListAction))
	v1.GET("/following", toV1(ctrl.GetFollowingsListAction))

	//mux.GET("/likes", toV1(ctrl.GetFollowingsListAction))
	v1.POST("/like", toV1(ctrl.PostAddLikeCtrl))
	v1.POST("/unlike", toV1(ctrl.PostRemoveLikeCtrl))
	v1.GET("/likes", toV1(ctrl.GetLikesAction))

	v1.POST("/comments/add", toV1(ctrl.PostAddCommentAction))
	v1.GET("/comments/list", toV1(ctrl.GetCommentsAction))
	v1.POST("/comments/delete", toV1(ctrl.RemoveCommentAction))

	v1.POST("/session/info", toV1(ctrl.SessionGetUserInfo))
	v1.GET("/session/info", toV1(ctrl.SessionGetUserInfo))

	v1.GET("/profile/posts", toV1(ctrl.GetPostsForProfileAction))
	v1.GET("/profile/info", toV1(ctrl.GetProfileInfoAction))

	v1.GET("/notify", toV1(ctrl.NotifyCtrl))
	v1.GET("/notify_add_remove", toV1(ctrl.NotifyAddRemoveCtrl))
	v1.GET("/activity", toV1(ctrl.ActivityListCtrl))

	v1.GET("/recommend/top_posts", toV1(ctrl.RecommendPostsCtrl))
	v1.GET("/recommend/users", toV1(ctrl.RecommendUsersCtrl))
	v1.GET("/recommend/top_tags", toV1(ctrl.RecommendTagsCtrl))
	v1.GET("/recommend/top_tags_discover", toV1(ctrl.RecommendTagsWithPostsCtrl))

	v1.GET("/tags/list", toV1(ctrl.TagsPostsListCtrl))

	v1.GET("/search/tags", toV1(ctrl.SearchTagsCtrl))
	v1.GET("/search/users", toV1(ctrl.SearchUsersCtrl))
	v1.GET("/search/cliked", toV1(ctrl.SearchClickedCtrl))

	v1.GET("/sync_users", toV1(ctrl.SyncUsersCtrl))

	http.HandleFunc("/ws_call", models.ServeHttpWs)
	///// v0.4 Msgs
	http.HandleFunc("/msgs/v1/add_one", ctrl.MsgUploadV1)

	http.Handle("/", v1Tree)

	if config.IS_DEBUG || true {

		http.Handle("/upload-avatar", actioner(UploadAvatarAction))
		http.Handle("/remove-avatar", actioner(RemoveAvatarAction))

		http.Handle("/fact/user1", actioner(fact.FactUser1))

		http.Handle("/fact/user_real", actioner(fact.FactRealUser))

		v1.GET("/json2/:UserId", toV1(TestJson1))
		v1.DELETE("/json2/:UserId", toV1(TestJson1))

		http.Handle("/upload1", actioner(TestUpload1))
		http.Handle("/upload3", actioner(PlayUpload3))
		http.HandleFunc("/upload2", PlayUpload2)

		http.Handle("/mysql1", actioner(fact.IsamPlay))
		http.Handle("/f/msg", actioner(fact.ChatMsgFact1))
		http.Handle("/f/gm", actioner(fact.GroupMemFact1))
		http.Handle("/f/ginfo", actioner(fact.GroupInfoFact1))

		//phone dbs
		http.Handle("/mf/contacts", actioner(fact.FactPhoneContacts))

		http.Handle("/ping", actionToFunc(ctrl.PingAction))

		http.Handle("/i/msg", actionToFunc(ctrl.SendSampleMesgTable3_v04))
		http.Handle("/i/redis", actionToFunc(RedisSavePlay))
		http.Handle("/i/play", actionToFunc(PlaySomething))
		http.Handle("/i/store1", actionToFunc(MemoryStore1))
		http.Handle("/i/cache", actionToFunc(ShowCached))
		http.Handle("/i/cacher", actionToFunc(ShowCacher))
		http.Handle("/i/cache_row", actionToFunc(ShowCacheKeys))
		http.Handle("/i/cache_row_keys", actionToFunc(ShowCacheRowKeys))
		http.Handle("/i/db", actionToFunc(DBStruct))
		http.Handle("/i/java2", actionToFunc(DBStructsTojava))
		http.Handle("/i/table", actionToFunc(DBStructsToTable))
		//fixes
		http.Handle("/i/fix_counts", actionToFunc(ctrl.FixAllCountsCounts))
		http.Handle("/i/mem_user", actionToFunc(ctrl.DebugMemUser_ctrl))
		http.Handle("/i/mem_user2", actionToFunc(ctrl.DebugMemUser_ctrl2))
		http.HandleFunc("/i/java", DBStructsTojava2)

		http.HandleFunc("/f/activity", actionToFunc(ctrl.InsertActivity))

		////////////// New Facts from v0.4 /////////////////
		http.Handle("/fact/follow", actioner(fact.FactFollow))
		http.Handle("/fact/unfollow", actioner(fact.FactUnFollow))
		http.Handle("/fact/avatar", actioner(fact.FactUserAvatars))
		http.Handle("/fact/like", actioner(fact.FactLike))
		http.Handle("/fact/like_fast", actioner(fact.FactLikeFast))
		http.Handle("/fact/comment", actioner(fact.FactComment))
		http.Handle("/fact/post", actioner(fact.FactPost))
		http.Handle("/fact/mix", actioner(fact.FactMix))
		http.Handle("/fact/about", actioner(fact.FactUpdateAboutMe))

	}

	return v1Tree

}

func toV1(fn func(*base.Action) base.AppErr) func(http.ResponseWriter, *http.Request, map[string]string) {
	return func(rw http.ResponseWriter, r *http.Request, parms map[string]string) {
		(base.Action{Fn2: fn, Ver: 2}).ServeHTTP(rw, r)
	}
}
