package actions

import (
	// "encoding/json"
	//	"fmt"
	// "strings"
	//	"strconv"
	"fmt"
	. "ms/sun/base"
	"ms/sun/helper"
	"os"
	//	"time"
	"io"
	"ms/sun/models/x"
	"net/http"
	"strconv"
	"strings"
)

func TestJson1(c *Action) AppErr {
	s := c.Req.Form.Get(":UserId")
	fmt.Println("Params: ", c.Req.Form)
	fmt.Println("Params: ", strings.Join(c.Req.Form["me"], "|"))
	sint, _ := strconv.ParseInt(s, 10, 0)
	print(s)
	print(c.Req.URL.RawQuery)
	u := x.User{}
	u22 := &x.User{}
	print(sint)
	u.CreatedTime = int(sint)
	u.Id = int(sint)
	u.FullName = c.Req.Host
	u.FirstName = c.Req.Method
	//var us []UserView
	var us []interface{}
	helper.StructToFiledsRejects(u22)
	n, m := helper.StructToFiledsRejects(&u)
	for i := 0; i < 1; i++ {
		us = append(us, u)
	}

	us = append(us, n, m)

	c.SendJson(us)
	//return ActionErr{}
	return nil
}

//func PlayUpload2(c *Action) {
func PlayUpload2(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1e6)
	//c,_:=r.Cookie("k")

	f, fd, err := r.FormFile("file")
	//f ,fd ,err :=r.MultipartForm.File["file"]
	if err != nil {
		print(err.Error())
		return
	}
	defer f.Close()
	fn, err := os.OpenFile("./upload/"+fd.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		print(2, err.Error())
		return
	}
	defer fn.Close()
	i, err := io.Copy(fn, f)

	print(i)
	if err != nil {
		print(3, err.Error())
		return
	}
}

func PlayUpload3(c *Action) {
	c.Req.ParseMultipartForm(1e6)
	f, fd, err := c.Req.FormFile("file")
	//f ,fd ,err :=r.MultipartForm.File["file"]
	if err != nil {
		print(err.Error())
		return
	}
	defer f.Close()
	fn, err := os.OpenFile("./upload/"+fd.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		print(2, err.Error())
		return
	}
	defer fn.Close()
	i, err := io.Copy(fn, f)

	print(i)
	if err != nil {
		print(3, err.Error())
		return
	}
}
