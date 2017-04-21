package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"sync"
)

type TopTagsWithPosts struct {
	Tag   Tag
	Posts []*PostAndDetailes
	//Posts []PostAndDetailes
}

type TopTagsWithPostsView struct {
	Tag   *Tag
	Posts []*PostView
}

func Tags_RepeatedlyJobs() {
	//ReloadAllTags()
	ReloadTopTags()
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
var TopTags = make([]*Tag, 0, 50)
var TopTagsWithPostsResult = make([]*TopTagsWithPostsView, 0, 50)

func ReloadTopTags() {
	tags, err := NewTag_Selector().OrderBy_Count_Desc().Limit(50).GetRows(base.DB)
	if err == nil {
		TopTags = tags
	}
}

func ReloadTopPostsForTopTags() {
	tags := TopTags
	var newTopTagsWithPosts = make([]*TopTagsWithPostsView, 0, 50)

	for _, t := range tags {
		postsIds, err := NewTagsPost_Selector().Select_PostId().TagId_EQ(t.Id).Limit(4).OrderBy_Id_Desc().GetIntSlice(base.DB)
		if err != nil {
			continue
		}

		v := &TopTagsWithPostsView{
			Tag:   t,
			Posts: Views.PostsViewsForPostIds(postsIds, 0),
		}

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
