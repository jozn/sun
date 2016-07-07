package base

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	//"bytes"
)

//structRow: must be poniter
func DbInsertStruct(structRow interface{}, table string) (sql.Result, error) {
	return DbInsertUpdateStruct(structRow, table, true)
}

//ignors Id in all structs to set by db- or not at all
func DbUpdateStruct(structRow interface{}, table string) (sql.Result, error) {
	return DbInsertUpdateStruct(structRow, table, false)
}

//update a struct will by defult update dy Id
func DbInsertUpdateStruct(structRow interface{}, table string, isInsert bool) (sql.Result, error) {
	s := reflect.ValueOf(structRow).Elem()
	var cols []string
	var vals []interface{}
	id := 0

	stut := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		filed := stut.Field(i).Name
		value := f.Interface()
		if filed != "Id" {
			cols = append(cols, filed)
			vals = append(vals, value)
		} else {
			id = value.(int) //only used for update
		}
	}

	// build "? ,? ,? ,? " for use of Mysql escape
	var ins []string
	var qus, cls, query string

	if isInsert { //building: insert into user (Name, UserName) Values(?,?)
		for i := 0; i < len(vals); i++ {
			ins = append(ins, "? ")
		}
		// strings.Repeat("?")
		qus = strings.Join(ins, ",")
		cls = strings.Join(cols, ",") //columsn " Id ,Nmae, UserName"
		query = "INSERT INTO " + table + " (" + cls + ") VALUES (" + qus + ")"
	} else { //building: UPDATE user set Name=? UserName=? WHERE Id = ?
		var sets []string
		for i := 0; i < len(cols); i++ {
			sets = append(sets, cols[i]+"= ? ")
		}
		setstr := strings.Join(sets, ",") //Id = ? ,Nmae = ? , UserName = ?
		query = "UPDATE " + table + " SET " + setstr + " WHERE Id = ?"
		vals = append(vals, id)
	}

	if __DEV__ {
		// fmt.Println("-db query: ", query)
	}

	r, err := DB.Exec(query, vals...)
	// DB.Elem

	if __DEV__ {
		if err != nil {
			fmt.Println("DATABASE ERROR: for table: ", table, " error: ", err)
		}
	}
	// noErr(err)
	return r, err

	// fmt.Println(cols, vals)
	// fmt.Println(cols, vals)
}

//update a struct will by defult update dy Id
func DbMassInsertStruct(table string, structRows ...interface{}) (sql.Result, error) {
	if len(structRows) == 0 {
		return nil, errors.New("length of structRows must be grater than zero")
	}
	structRow := structRows[0]
	s := reflect.ValueOf(structRow) //.Elem()
	var fieldStruct []string

	stut := s.Type()
	for i := 0; i < s.NumField(); i++ {
		filed := stut.Field(i).Name
		if filed != "Id" {
			fieldStruct = append(fieldStruct, filed)
		}
	}

	//	var valRaw []interface{}
	var valHolders []interface{}
	for i := 0; i < len(structRows); i++ {
		st2 := reflect.ValueOf(structRows[i]) //.Elem()
		for j := 0; j < len(fieldStruct); j++ {
			v := st2.FieldByName(fieldStruct[j])
			valHolders = append(valHolders, v.Interface())
		}
	}

	//fmt.Println(valHolders)

	// build "? ,? ,? ,? " for use of Mysql escape
	var insAll []string //[ "(?,?,?)" , "?,?,?" , ... ]
	var qus, colsSql, query string

	//TODO:optimize and simpilify beacus it is just "(?,?,?), ..." we can build one and repeat it
	for i := 0; i < len(structRows); i++ {
		var insRow []string
		for j := 0; j < len(fieldStruct); j++ {
			insRow = append(insRow, "?")
		}
		insAll = append(insAll, "("+strings.Join(insRow, ",")+")") // building: "(?,?,?)"
	}

	qus = strings.Join(insAll, ", ")         // "(?,?,?) , (?,?,?) ..."
	colsSql = strings.Join(fieldStruct, ",") //columsn " Id ,Nmae, UserName"
	query = "INSERT INTO " + table + " (" + colsSql + ") VALUES " + qus + " "

	if __DEV__ {
		// fmt.Println("-db query: ", query)
	}

	r, err := DB.Exec(query, valHolders...)

	if __DEV__ {
		if err != nil {
			log.Println("database err for table: ", table, " error: ", err)
		}
	}

	return r, err

}

func dbQuestion(size int) string {
	s := strings.Repeat("?,", size)
	s = s[0 : len(s)-2] //remove last ','
	return s
}

func DbStructToTable(structRow interface{}, table string) string {
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
			sqlCol = fmt.Sprintf("  `%v` int(10) DEFAULT '0' ", filedName)
		case reflect.String:
			sqlCol = fmt.Sprintf("  `%v` varchar(250) DEFAULT ''", filedName)
		}
		cols = append(cols, sqlCol)

	}
	//	fmt.Println(cols)
	sql1 := fmt.Sprintf(strings.Join(cols, ",\n"))

	sql2 := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%v` (\n%v \n) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4;\n", table, sql1)
	fmt.Println(sql2)
	return sql2
}

func DbQuestionForIn(size int) string {
	s := strings.Repeat("?,", size)
	s = s[0 : len(s)-1] //remove last ','
	return s
}

//["ad'sa","qwe","opi"] -returns-> MysqlEsapeted of `"ad\'sa","qwe","opi"`
func DbSliceStringToSafeIns(str []string) string {
	//var buf bytes.Buffer
	var buf []string
	for _,s := range str{
		//buf.WriteString(s)
		buf = append(buf, "'"+MySqlEscape(s)+"'" )
	}
	return strings.Join(buf,",")
}



