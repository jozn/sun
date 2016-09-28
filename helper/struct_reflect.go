package helper

import (
	//"fmt"
	"fmt"
	"reflect"
	"unicode"
)

// isCollect: true => just collect fileds , false:  rejects fileds
func structToFiledsGeneral(structPointer interface{}, isCollect bool, filedsFilter ...string) (fileds []string, values []interface{}) {
	s := reflect.ValueOf(structPointer).Elem().Interface() //for pointer
	rjs := make(map[string]bool)

	for _, r := range filedsFilter {
		rjs[r] = true
	}

	//should run @A if bloc
	shouldRunBlock := func(existis bool) bool {
		switch {
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

	vls, fls := DeepFields(s) //len(fls) == len(valeus)
	for i, filed := range fls {
		//@A
		if inMap := rjs[filed]; shouldRunBlock(inMap) {
			fileds = append(fileds, fls[i])
			values = append(values, vls[i])
		}
	}

	return //fileds, values
}
func DeepFields(iface interface{}) (valeus []interface{}, fields []string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		n := ift.Field(i).Name
		if unicode.IsLower(rune(n[0])) { //will painc for un-exported filds
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			vls, cls := DeepFields(v.Interface())
			valeus = append(valeus, vls...)
			fields = append(fields, cls...)
		default:
			valeus = append(valeus, v.Interface())
			fields = append(fields, n)
		}
	}
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

func StructToFiledsCollect(structPointer interface{}, collectFileds ...string) (fileds []string, values []interface{}) {
	return structToFiledsGeneral(structPointer, true, collectFileds...)
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
func StructToFiledsRejects(structPointer interface{}, rejecttFileds ...string) (fileds []string, values []interface{}) {
	return structToFiledsGeneral(structPointer, false, rejecttFileds...)
}

func StructToFiledsRejectsEscape(structPointer interface{}, rejecttFileds ...string) (fileds []string, values []string) {
	fileds, vls := StructToFiledsRejects(structPointer, rejecttFileds...)
	values = StructValuesToMysqlEscape(vls...)
	return
}

func StructValuesToMysqlEscape(values ...interface{}) []string {
	//if len(fileds)
	valuesEscaped := make([]string, 0, len(values))
	for _, val := range values {
		valuesEscaped = append(valuesEscaped, MySqlEscape(fmt.Sprintf("%v", val)))
	}
	return valuesEscaped
}

///////////////////////////////////////////////////

func DbStructToJava(structRow interface{}) string {
	_, fileds := DeepFields(structRow)
	refVal := reflect.ValueOf(structRow)
	javaFilsdStr := ""
	for _, field := range fileds {
		f := refVal.FieldByName(field)
		javaType := ""
		switch f.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			javaType = "int"
		case reflect.Float32, reflect.Float64:
			javaType = "float"
		case reflect.Bool:
			javaType = "boolean"
		case reflect.String:
			javaType = "String"
		}
		javaFilsdStr += " public " + javaType + " " + field + ";\n"
	}

	res := "public class V {\n" + javaFilsdStr + "\n}"
	return res
}
