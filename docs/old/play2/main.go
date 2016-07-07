package main
import (
	"reflect"
	"fmt"
	"strings"
)
func main()  {
	println("das")
	dbStructToTable(&Sample{},"SAMLE")
}

type Sample  struct{

	UserId int
	UserName string
}

func dbStructToTable(structRow interface{}, table string) string{
	s := reflect.ValueOf(structRow).Elem()
	var cols []string
	_ = cols
	stut := s.Type()
	//	fmt.Printf("struc %T %v ", stut, stut)
	//	fmt.Println("struc  ", stut)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		filedName := stut.Field(i).Name
		//		fmt.Printf("f: %v - %T\n",f,f)
		//		fmt.Printf("filed: %v --- %T \n", filedName, filedName)
		sqlCol := ""
		//		f.t
		switch f.Type().Kind() {
		case reflect.Int:
			sqlCol = fmt.Sprintf("  `%v` int(10) DEFAULT '0' ",filedName)
		case reflect.String:
			sqlCol = fmt.Sprintf("  `%v` varchar(250) DEFAULT ''",filedName)
		}
		cols = append(cols,sqlCol)

	}
	//	fmt.Println(cols)
	sql1:= fmt.Sprintf(strings.Join(cols, ",\n") )

	sql2:= fmt.Sprintf("CREATE TABLE IF NOT EXIST `%v` (\n%v \n) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4;\n",table,sql1)
	fmt.Println(sql2)
	return sql2
}
