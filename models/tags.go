package models

import (
    "ms/sun/base"
    "sync"
    "ms/sun/helper"
)

///////////////// Structs Database /////////////////////////////
type Tag struct {
    Id          int
    Name        string
    Count       int
    IsBlocked   int
    CreatedTime int
}

type TagPost struct {
    Id          int
    TagId       int
    PostId      int
    TypeId      int // text? photo? video?
    CreatedTime int
}

/////////////////////////////////////////////
type tagsMap struct {
    m sync.RWMutex
    _map map[string]*Tag
}

func newTagsMap() *tagsMap {
    tm := tagsMap{}
    tm._map = make(map[string]*Tag)
    return  &tm
}

var TagsMap = newTagsMap() //new(tagsMap)
var TopTags = make([]Tag,0,50)

//called on inits
func ReloadAllTags()  {
    var tags []Tag
    base.DB.Select(&tags ,"select * from tags")

    TagsMap.m.Lock()
    for _, t:= range tags {
        t2 := t
        TagsMap._map[t2.Name] = &t2
    }
    TagsMap.m.Unlock()
}

func ReloadTopTags()  {
    var tags []Tag
    base.DB.Select(&tags ,"SELECT * FROM tags ORDER BY `count` DESC LIMIT 50 ")

    TopTags = tags
}


//////////////////  New Apis   ///////////////////////////
func AddTagsInPost(post *Post) {
    parser := TextParser{}
    parser.Parse(post.Text)

    //create new Tags if for first time
    TagsMap.m.Lock()
    defer TagsMap.m.Unlock()

    for _, tagName := range parser.Tags {
        tg, ok :=TagsMap._map[tagName]

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
        tg, ok :=TagsMap._map[tagName]
        if ok {
            tagPost := TagPost{}
            tagPost.TagId = tg.Id
            tagPost.PostId = post.Id
            tagPost.TypeId = post.TypeId
            tagPost.CreatedTime = helper.TimeNow()

            postTagPonters = append(postTagPonters, &tagPost)
        }
        //base.DbInsertStruct(&tagPost, "tags_posts")
    }
    base.DbMassReplacetStructPoninters("tags_posts",postTagPonters...)

    var tgsNames []interface{}
    for _, tagName := range parser.Tags {
        n := tagName
        tgsNames = append(tgsNames, &n)
    }

    base.DbExecute("update tags set `Count` = `Count`+1 where Name in (" + helper.DbQuestionForSqlIn(len(tgsNames)) +")" ,tgsNames... )
}

func AddUserMentionedInPost(post *Post) {

}

func AddTagsInPost_OLD_DEP(post Post) {
    parser := TextParser{}
    parser.Parse(post.Text)
    for _, tag := range parser.Tags {
        var dbTags []Tag
        var dbTag Tag
        base.DB.Select(&dbTags, "select * from tags where Name = ? ", tag)
        if len(dbTags) == 0 { //not exist ,insert it
            dbTag = Tag{}
            dbTag.Name = tag
            dbTag.CreatedTime = now()
            res, _ := base.DbInsertStruct(&dbTag, "tags")
            tid, _ := res.LastInsertId()
            dbTag.Id = int(tid)
        } else {
            dbTag = dbTags[0]
        }

        tagPost := TagPost{}
        tagPost.TagId = dbTag.Id
        tagPost.PostId = post.Id
        tagPost.TypeId = post.TypeId
        tagPost.CreatedTime = now()

        base.DbInsertStruct(&tagPost, "tags_posts")
        //TODO increment dbTags.Count
    }
}
