package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"sync"
)

///////////////// Structs Database /////////////////////////////
// Go-Orma
type Tag struct {
	Id          int
	Name        string
	Count       int
	IsBlocked   int
	CreatedTime int
	// xo fields
	_exists, _deleted bool
}

type TagsPost struct {
	Id          int
	TagId       int
	PostId      int
	TypeId      int // text? photo? video?
	CreatedTime int
	// xo fields
	_exists, _deleted bool
}

/////////////////////////////////////////////
type tagsMap struct {
	m    sync.RWMutex
	_map map[string]*Tag
}

func (mp *tagsMap) GetTag(Name string) (*Tag, bool) {
	mp.m.RLock()
	t := mp._map[Name]

	if t == nil {
		return nil, false
	}
	return t, true
}

func newTagsMap() *tagsMap {
	tm := tagsMap{}
	tm._map = make(map[string]*Tag)
	return &tm
}

var TagsMap = newTagsMap() //new(tagsMap)
var TopTags = make([]Tag, 0, 50)
var TopTagsWithPostsResult = make([]TopTagsWithPosts, 0, 50)

//called on inits
func ReloadAllTags() {
	var tags []Tag
	base.DB.Select(&tags, "select * from tags")

	TagsMap.m.Lock()
	for _, t := range tags {
		t2 := t
		TagsMap._map[t2.Name] = &t2
	}
	TagsMap.m.Unlock()
}

func ReloadTopTags() {
	var tags []Tag
	base.DB.Select(&tags, "SELECT * FROM tags ORDER BY `count` DESC LIMIT 50 ")

	TopTags = tags
}

func ReloadTopPostsForTopTags() {

	tags := TopTags
	tagsId := []int{}

	var newTopTagsWithPosts = make([]TopTagsWithPosts, 0, 50)

	for _, t := range tags {
		tagsId = append(tagsId, t.Id) //useless
		var PostsIds []int

		base.DB.Select(&PostsIds, "SELECT PostId FROM tags_posts where TagId = ? AND TypeId = 2 ORDER BY Id DESC LIMIT 4 ", t.Id)

		post := GetPostsAndItsDetails(PostsIds, 0)

		v := TopTagsWithPosts{}
		v.Tag = t
		v.Posts = post
		newTopTagsWithPosts = append(newTopTagsWithPosts, v)
	}

	TopTagsWithPostsResult = newTopTagsWithPosts
}

//////////////////  New Apis   ///////////////////////////
func AddTagsInPost(post *Post) {
	parser := TextParser{}
	parser.Parse(post.Text)

	//create new Tags if for first time
	TagsMap.m.Lock()
	defer TagsMap.m.Unlock()

	for _, tagName := range parser.Tags {
		tg, ok := TagsMap._map[tagName]

		if !ok { // first post whit this tag with
			tg = &Tag{}
			tg.Name = tagName
			tg.CreatedTime = helper.TimeNow()
			res, _ := base.DbInsertStruct(tg, "tags")
			tid, _ := res.LastInsertId()
			tg.Id = int(tid)

			TagsMap._map[tg.Name] = tg
		}
		//add in memeory counter
		tg.Count += 1
	}

	var postTagPonters []interface{}
	for _, tagName := range parser.Tags {
		tg, ok := TagsMap._map[tagName]
		if ok {
			tagPost := TagsPost{}
			tagPost.TagId = tg.Id
			tagPost.PostId = post.Id
			tagPost.TypeId = post.TypeId
			tagPost.CreatedTime = helper.TimeNow()

			postTagPonters = append(postTagPonters, &tagPost)
		}
		//base.DbInsertStruct(&tagPost, "tags_posts")
	}
	base.DbMassReplacetStructPoninters("tags_posts", postTagPonters...)

	var tgsNames []interface{}
	for _, tagName := range parser.Tags {
		n := tagName
		tgsNames = append(tgsNames, &n)
	}

	base.DbExecute("update tags set `Count` = `Count`+1 where Name in ("+helper.DbQuestionForSqlIn(len(tgsNames))+")", tgsNames...)
}

func AddUserMentionedInPost(post *Post) {

}
