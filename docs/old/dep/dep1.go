
func GetPostsActionPlay1(c *Action) {
	devPrintn("action post")
	var rs []Post
	DB.Select(&rs, "select * from post") //, rand.Intn(1000000))
	// noErr(err)
	// print(len(rs))
	// b, _ := json.Marshal(rs)
	// c.SendText(string(b))

	var viw []PostAndDetailes

	for _, p := range rs {
		v := &PostAndDetailes{}
		v.Post = p
		v.Sender = GetUserView(p.UserId)
		v.Comments = GetPostLastComments(p.Id)
		v.Likes = GetPostLastLikes(p.Id)
		v.AmIlike = false
		viw = append(viw, *v)
	}

	c.SendJson(viw)

	// var rs []test5
	// for i := 0; i < 50; i++ {
	// 	DB.Select(&rs, "select * from test whers id = ?", rand.Intn(1000000))
	// }
	// b, _ := json.Marshal(rs)
	// c.SendText(string(b))
}

func GetUserView2(uid int) UserInlineView {
	debug("GetUserView: ", uid)
	// u := GetUserById(uid)
	v := UserInlineView{}
	v.FullName = "asdasdasdasdas" //u.FirstName + " " + u.LastName
	v.UserId = 15                 //u.Id
	v.UserName = "u.UserName"
	v.AvatarSrc = "u.AvatarSrc" //u.AvatarSrc
	return v
}

// modified copied from db.go
//returns "UserName, UserId," + stripd Id
func structSqlNamesValues(stuct interface{}) (string, []interface{}) {
	s := reflect.ValueOf(stuct).Elem()
	var cols []string
	var vals []interface{}
	// id := 0

	stut := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		filed := stut.Field(i).Name
		value := f.Interface()
		if filed != "Id" {
			cols = append(cols, filed)
			vals = append(vals, value)
		}
		// else {
		// 	// id = value.(int) //only used for update
		// }
	}
	cl := strings.Join(cols, ", ")
	return cl, vals
	// cl := strings.Join(cols, ", ")
	// vl := strings.Join(vals, ", ")
	// sql := "(" + cl + ") Values (" + vl + ")"
	// return sql
}

