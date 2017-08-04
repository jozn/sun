package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	// "github.com/nfnt/resize"
	"log"
	"ms/sun/base"
	"ms/sun/models"
	"net/http"

	_ "github.com/garyburd/redigo/redis"
	//_ "net/http/pprof"
	_ "net/http/pprof"

	. "ms/sun/actions"
	"ms/sun/config"
	"ms/sun/ctrl"
	"ms/sun/fact"
	"ms/sun/helper"

	"github.com/dimfeld/httptreemux"
	"github.com/mediocregopher/radix.v2/pool"
)

var redisPool *pool.Pool

func main() {
	print("Try To Start The App")
	startApp()

	/*go func() {
		models.SERVER_GRPC()
	}()

	go func() {
		models.Client_GRPC()
	}()*/

	/*go func() {
		helper.JustRecover()
		//http.ListenAndServe(":5000", nil)
	}()*/

	http.ListenAndServe(":5000", nil)

	//http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	//runtime.MemProfileRecord{}.
}

func startApp() {
	var err error
	//DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3308)/ms?charset=utf8mb4,utf8&collation=utf8mb4_general_ci")
	base.DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms?charset=utf8mb4")
	//DB, err = sqlx.Connect("mysql", "root:123456@localhost:3307/ms?charset=utf8mb4")
	//DB.Exec("SET NAMES 'utf8mb4';")
	noErr(err)
	base.DB.MapperFunc(func(s string) string { return s })
	redisInit()

	fmt.Println("main start")
	v1Tree := registerRoutes()
	//base.RegisterGlobTypes()

	//// Inits ///////////////
	//registerCmdRouters()

	http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	//in models
	models.OnAppStart_Models()

	_ = v1Tree

	go func() {
		defer helper.JustRecover()
		//http.ListenAndServe(":6000", v1Tree)
	}()

}

func redisInit() {
	var err error
	redisPool, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		fmt.Println("redis failed")
		return
	}
	fmt.Println("redis online")
	// In another go-routine

	conn, err := redisPool.Get()
	if err != nil {
		// handle error
	}

	if conn.Cmd("SET", "KEYYY", "Val").Err != nil {
		// handle error
	}

	redisPool.Put(conn)
}

func noErr(err error) {
	if err != nil {
		// log.Fatalln(err) //painc
		log.Panic("***** PANIC ****: ", err)
	}
}

func actioner(action func(*base.Action)) http.Handler {
	return &base.Action{Fn: action}
}

//for Version 2 of Action -- that returns ActionErr
func actionToFunc(action func(*base.Action) base.AppErr) http.HandlerFunc {
	return (&base.Action{Fn2: action, Ver: 2}).ServeHTTP
}

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
	v1.GET("/post/stream", toV1(ctrl.GetPostsStreamAction))
	//v1.GET("/post/stream2", toV1(ctrl.GetPostsStraemAction_NEW))
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

	//http.HandleFunc("/ws_call", models.ServeHttpWs)
	http.HandleFunc("/ws_pb_call", models.ServeHttpWsPB)
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

		http.Handle("/upload3", actioner(PlayUpload3))
		http.HandleFunc("/upload2", PlayUpload2)

		//phone dbs

		http.Handle("/ping", actionToFunc(ctrl.PingAction))

		//http.Handle("/i/msg", actionToFunc(ctrl.SendSampleMesgTable3_v04))
		http.Handle("/i/msg2", actionToFunc(ctrl.SendSampleMesg))
		http.Handle("/i/redis", actionToFunc(RedisSavePlay))
		http.Handle("/i/play", actionToFunc(PlaySomething))
		http.Handle("/i/cache", actionToFunc(ShowCached))
		http.Handle("/i/cacher", actionToFunc(ShowCacher))
		http.Handle("/i/cache_row", actionToFunc(ShowCacheKeys))
		http.Handle("/i/cache_row_keys", actionToFunc(ShowCacheRowKeys))
		http.Handle("/i/cache_index", actionToFunc(ShowCacheIndex))
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
		http.Handle("/fact/unlike", actioner(fact.FactUnLike))
		http.Handle("/fact/like_fast", actioner(fact.FactLikeFast))
		http.Handle("/fact/comment", actioner(fact.FactComment))
		http.Handle("/fact/post", actioner(fact.FactPosts))
		//http.Handle("/fact/post2", actioner(fact.FactPosts2))
		http.Handle("/fact/mix", actioner(fact.FactMix))
		http.Handle("/fact/mix_delay", actioner(fact.FactDelayMix))
		http.Handle("/fact/about", actioner(fact.FactUpdateAboutMe))

	}

	return v1Tree

}

func toV1(fn func(*base.Action) base.AppErr) func(http.ResponseWriter, *http.Request, map[string]string) {
	return func(rw http.ResponseWriter, r *http.Request, parms map[string]string) {
		(base.Action{Fn2: fn, Ver: 2}).ServeHTTP(rw, r)
	}
}
