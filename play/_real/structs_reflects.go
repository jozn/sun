package main

import (
	"fmt"
	"reflect"
	"unicode"
	"ms/sun/helper"
)

type U struct {
	In2 int
	Id2 int64
	Name2 string
	Ok bool
	fi2 int

}
type User struct {
	U
	In int
	Id int64
	Name string
	Ok bool
	fi int

}


// isCollect: true => just collect fileds , false:  rejects fileds
func structToFiledsGeneral(structPointer interface{}, isCollect bool , filedsFilter ...string )(fileds []string, values []interface{} ) {
	//func structToFiledsGeneral(structPointer interface{}, isCollect bool , filedsFilter ...string )([]string, []interface{} ) {
	s := reflect.ValueOf(structPointer).Elem().Interface()//for pointer
	rjs := make(map[string]bool)
	//var fileds []string
	//var values []interface{}

	for _ , r := range filedsFilter {
		rjs[r] = true
	}
	//fmt.Println(rjs)

	//should run @A if bloc
	shouldRunBlock := func(existis bool) bool {
		switch  {
		case isCollect && existis:
			return true
		case isCollect && !existis:
			return false

		//it is reject list
		case !isCollect && existis:
			return false
		case !isCollect && !existis:
			return true
		}

		return false
	}
	//_ = shouldRunBlock
	//_ =s

	vls, fls := DeepFields(s)//len(fls) == len(valeus)
	//fmt.Println("Depps: ",vls,fls)
	for i,filed := range fls {
		if  inMap := rjs[filed]; shouldRunBlock(inMap) {
			fileds = append(fileds, fls[i])
			values = append(values, vls[i])
		}
	}

	//fmt.Println(DeepFields(s))

	//allFileds := make([]string,10)//for embeded

	//stut := s.Type()
/*	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		filed := stut.Field(i).Name
		if unicode.IsLower(rune(filed[0])) {//will painc for un-exported filds
			continue
		}
		switch f.Kind() {
		case reflect.Struct:
			fields = append(fields, DeepFields(v.Interface())...)
		default:
			fields = append(fields, v)
		}
		allFileds = append(fileds, filed)*/
		/*
		value := f.Interface()
		if reflect
		//if filed != "Id" {
		//@A
		if  inMap := rjs[filed]; shouldRunBlock(inMap)  {
			fileds = append(fileds, filed)
			values = append(values, value)
		} else {
			//id = value.(int) //only used for update
		}*/
	//}

	//fmt.Println(fileds)
	//fmt.Println(values)
	return //fileds, values
}

func DeepFields(iface interface{}) (valeus []interface{} , fields []string) {
	//reflect.ValueOf(iface)
	//fields := make([]reflect.Value, 0)
	//fields := make([]interface{}, 0)
	//cols:= make([]string, 0)
	//ifv := reflect.ValueOf(iface).Elem()//.Elem() for pointer
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	//ift := ifv.Type()

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		n := ift.Field(i).Name
		if unicode.IsLower(rune(n[0])) {//will painc for un-exported filds
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			vls ,cls := DeepFields(v.Interface())
			valeus = append(valeus, vls...)
			fields = append(fields, cls...)
		default:
			//fields = append(fields, v)
			valeus = append(valeus, v.Interface())
			fields = append(fields, n)
		}
	}
	//fmt.Println(cols)
	//ss := fmt.Sprintf("%t",cols[0])
	//fmt.Println(ss)
	return valeus, fields
}
/*
example
u := User{}
u.In = 5888
u.Name = "asdasdasdasdسسس"
u.fi =444

StructToFiledsCollect(&u,"name","Name","In")

[5888 asdasdasdasdسسس]
[Id Ok]
*/

func StructToFiledsCollect(structPointer interface{} ,collectFileds ...string) (fileds []string, values []interface{} )  {
	return structToFiledsGeneral(structPointer, true , collectFileds...)
}

/*
example
u := User{}
u.In = 5888
u.Name = "asdasdasdasdسسس"
u.fi =444

StructToFiledsRejects(&u,"name","Name","In")

[Id Ok]
[0 false]
*/
func StructToFiledsRejects(structPointer interface{} ,rejecttFileds ...string) (fileds []string, values []interface{} )  {
	return structToFiledsGeneral(structPointer, false , rejecttFileds...)
}

func StructValuesToMysqlEscape(values ...interface{} ) ([]string) {
	//if len(fileds)
	valuesEscaped := make([]string,len(values))
	for _,val := range values{
		valuesEscaped = append(valuesEscaped, helper.MySqlEscape(fmt.Sprintf("%v",val)))
	}
	return valuesEscaped
}


func main() {
	u := User{}
	u.In = 5888
	u.Name = "asdasl'd\"asd`asd`سسس"
	u.fi =444
	//fmt.Println(DeepFields(u))
	//DeepFields(u)
	//StructToFiledsCollect(u,"name","Name","In")
	//StructToFiledsCollect(&u,"name","Name","In")

	fmt.Println(StructToFiledsCollect(&u,"Name","Id"))
	//structToFiledsGeneral(&u, false,  "Id")
	//structToFiledsGeneral(&u, true,  "Id")
	////structToFiledsGeneral(u, true,  "Id")
	//
	//fmt.Println(StructToFiledsRejects(&u))
	//StructToFiledsCollect(&u,"name","Name","In")
	//_,n:=StructToFiledsRejects(&u)//,"name","Name","In")
	//fmt.Println(StructValuesToMysqlEscape(n...))
}
