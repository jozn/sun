package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"time"
)

type TopTagsWithPostsView struct {
	Tag   *Tag
	Posts []*PostView
}

func Tags_RepeatedlyJobs() {

	//top tags
	go func() {
		for {
			ReloadTopTags()
			ReloadTopPostsForTopTags()
			time.Sleep(time.Minute * 15)
		}
	}()

}

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

func AddTagsInPost(post *Post) {
	parser := TextParser{}
	parser.Parse(post.Text)

	if len(post.Text) == 0 {
		return
	}

	Store.PreLoadTag_ByNames(parser.Tags)

	tags := []*Tag{}
	tagPosts := []TagsPost{}
	tagsIds := []int{}
	for _, tagName := range parser.Tags {
		tg, ok := Store.Tag_ByName(tagName)
		if !ok {
			tg = &Tag{
				Name:        tagName,
				CreatedTime: helper.TimeNow(),
			}
			tg.Save(base.DB)
		}
		tgp := TagsPost{
			TagId:       tg.Id,
			PostId:      post.Id,
			TypeId:      post.TypeId,
			CreatedTime: helper.TimeNow(),
		}

		tags = append(tags, tg)
		tagPosts = append(tagPosts, tgp)
		tagsIds = append(tagsIds, tg.Id)
	}

	MassInsert_TagsPost(tagPosts, base.DB)
	NewTag_Updater().Count_Increment(1).Id_In(tagsIds).Update(base.DB)
}

func AddUserMentionedInPost(post *Post) {

}
