package actions

import (
	// "encoding/json"
	//	"fmt"
	// "strings"
	//	"strconv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "ms/sun/base"
	. "ms/sun/models"
	"ms/sun/helper"
	"os"
	//	"time"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func TestJson1(c *Action) AppErr {
	s :=c.Req.Form.Get(":UserId")
	fmt.Println("Params: ", c.Req.Form)
	fmt.Println("Params: ", strings.Join(c.Req.Form["me"],"|") )
	sint,_ := strconv.ParseInt(s,10,0)
	print(s)
	print(c.Req.URL.RawQuery)
	u := User{}
	u22 := &User{}
	print(sint)
	u.CreatedTimestamp =int(sint)
	u.Id =int(sint)
	u.FullName = c.Req.Host
	u.FirstName = c.Req.Method
	//var us []User
	var us []interface{}
	helper.StructToFiledsRejects(u22)
	n,m:= helper.StructToFiledsRejects(&u)
	for i := 0; i < 1; i++ {
		us = append(us, u)
	}

	us = append(us, n ,m)

	c.SendJson(us)
	//return ActionErr{}
	return nil
}

func TestUpload1(c *Action) {

	c.Req.ParseForm()
	d := c.Req.Form.Get("data2")
	println(d)
	f, err := os.OpenFile("./con.json", 0, 0666)

	s, err2 := f.Write([]byte("dasdasdas"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err, s, err2)
	ioutil.WriteFile("./con2.json", []byte(d), 0666)
	var contacts []PhoneContactTable
	json.Unmarshal([]byte(d), &contacts)
	println(contacts)

	var interfaceSlice []interface{} = make([]interface{}, len(contacts))
	for i, d := range contacts {
		row := d
		//		row :=PhoneContactTable{}
		//		row.PhoneContact = d
		row.UserId = c.UserId()
		row.CreatedTimeStamp = now()
		row.UpdatedTimeStamp = now()
		interfaceSlice[i] = row //DbMassInsertStruct only accets real []interface{}
		fmt.Println(i)
		for i := 0; i < 1000; i++ {
			// DbMassInsertStruct("phone_contacts2", interfaceSlice...)
			// fmt.Println(i)
			DbInsertStruct(&row, "phone_contacts2")
		}

	}

	// DbMassInsertStruct("phone_contacts2", interfaceSlice...)
	for i := 0; i < 1000; i++ {
		// DbMassInsertStruct("phone_contacts2", interfaceSlice...)
		// fmt.Println(i)
		_ = i
	}
	f.Close()
	c.SendJson(d)
	//	for i:= 0;i<3 ; i++  {
	////		interfaceSlice = append(interfaceSlice,interfaceSlice...)
	//	}
	//
	//	for i:= 0;i<100000 ; i++  {
	//		 DbMassInsertStruct("phone_contacts",interfaceSlice...)
	//		fmt.Println(i)
	////		time.Sleep(200 *time.Millisecond)
	//	}
	//	DbMassInsertStruct("phone_contacts",interfaceSlice...)
	//	for i:=0;i<len(contacts);i++ {
	//		DbInsertStruct(&contacts[i],"phone_contacts2")
	//
	//	}
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
