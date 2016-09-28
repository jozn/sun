package base

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	//"bytes"
	"ms/sun/config"
	"ms/sun/helper"
)

//structRow: must be poniter
func DbInsertStruct(structRow interface{}, table string) (sql.Result, error) {
	//return DbInsertUpdateStruct(structRow, table, true)
	keys, values := helper.StructToFiledsRejectsEscape(structRow, "Id")

	ks := make([]string, len(keys))
	for i, s := range keys {
		ks[i] = "`" + s + "`"
	}

	//q := "Insert into " + table + " ("+strings.Join(ks,",") +") values (" +strings.Join(values,",") +")"
	q := "Insert " + table + " (" + strings.Join(ks, ",") + ") values (" + DbQuestionForIn(len(values)) + ")"

	vals := make([]interface{}, len(values))
	for i, s := range values {
		//vals[i] = "`"+s+"`"
		vals[i] = s
	}

	r, err := DB.Exec(q, vals...)
	//r, err := DB.Exec(q)

	if config.DEBUG_LOG_SQL {
		helper.DebugPrintln(q, vals)
		if err != nil {
			fmt.Println("DATABASE ERROR: for table: ", table, " error: ", err, " query ", q, " PARAMS: ", vals)
		}
	}
	return r, err
}

//ignors Id in all structs to set by db- or not at all
func DbUpdateStruct(structRow interface{}, table string) (sql.Result, error) {
	//return DbInsertUpdateStruct(structRow, table, false)
	keys, values := helper.StructToFiledsRejectsEscape(structRow, "Id")
	_, idavls := helper.StructToFiledsCollect(structRow, "Id")

	ks := make([]string, len(keys))
	for i, s := range keys {
		ks[i] = "`" + s + "`"
	}
	idval := ""
	if len(idavls) == 1 {
		idval = fmt.Sprint(idavls[0])
	} else {
		log.Panic("In Update statemant Id of struct is not exist - struct: ", structRow)
	}

	q := "Update " + table + " (" + strings.Join(ks, ",") + ") values (" + DbQuestionForIn(len(values)) + ") WHERE Id = " + idval

	vals := make([]interface{}, len(values))
	for i, s := range values {
		//vals[i] = "`"+s+"`"
		vals[i] = s
	}

	r, err := DB.Exec(q, vals...)

	if config.DEBUG_LOG_SQL {
		helper.DebugPrintln(q, vals)
		if err != nil {
			fmt.Println("DATABASE ERROR: for table: ", table, " error: ", err, " query ", q, " PARAMS: ", vals)
		}
	}
	return r, err
}

func DbExecute(query string, params ...interface{}) (sql.Result, error) {
	r, err := DB.Exec(query, params...)

	if config.DEBUG_LOG_SQL {
		helper.DebugPrintln(query, params)
		if err != nil {
			fmt.Println("DATABASE ERROR: error: ", err, " query ", query, " PARAMS: ", params)
		}
	}
	return r, err
}

//structRows must be an [] of POINTERS
func DbMassReplacetStructPoninters(table string, structRows ...interface{}) (sql.Result, error) {
	//fmt.Println("Inside DbMassReplacetStruct")
	if len(structRows) == 0 {
		w := errors.New("length of structRows must be grater than zero")
		helper.DebugPrintln(w)
		return nil, w
	}
	structRow := structRows[0]

	keys, values := helper.StructToFiledsRejectsEscape(structRow, "Id")

	ks := make([]string, len(keys))
	for i, s := range keys {
		ks[i] = "`" + s + "`"
	}

	// build "? ,? ,? ,? " for use of Mysql escape
	var insAll []string //[ "(?,?,?)" , "?,?,?" , ... ]
	var qus, colsSql, query string

	//TODO:optimize and simpilify beacus it is just "(?,?,?), ..." we can build one and repeat it
	for i := 0; i < len(structRows); i++ {
		var insRow []string
		for j := 0; j < len(values); j++ {
			insRow = append(insRow, "?")
		}
		insAll = append(insAll, "("+strings.Join(insRow, ",")+")") // building: "(?,?,?)"
	}

	qus = strings.Join(insAll, ", ") // "(?,?,?) , (?,?,?) ..."
	colsSql = strings.Join(ks, ",")  //columsn " Id ,Nmae, UserName"
	query = "REPLACE INTO " + table + " (" + colsSql + ") VALUES " + qus + " "

	//fmt.Println(query)
	//	var valRaw []interface{}
	var valHolders []interface{}
	for n := 0; n < len(structRows); n++ {
		_, values2 := helper.StructToFiledsRejectsEscape(structRows[n], "Id")

		for j := 0; j < len(values2); j++ {
			valHolders = append(valHolders, values2[j])
		}
	}

	r, err := DB.Exec(query, valHolders...)

	if config.DEBUG_LOG_SQL {
		helper.DebugPrintln(query, valHolders)
		if err != nil {
			fmt.Println("DATABASE ERROR: error: ", err, " query ", query, " PARAMS: ", valHolders)
		}
	}

	return r, err

}

//update a struct will by defult update dy Id
// deprecated
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
		fmt.Println("-db query: ", query, vals)
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

//deprecated use DbMassReplacetStruct
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

func DbStructToTable2New(structRow interface{}, table string) string {
	arr_str, arr_vals := helper.StructToFiledsRejects(structRow, "nothing")
	var cols []string
	_ = cols
	//stut := s.Type()
	//	fmt.Printf("struc %T %v ", stut, stut)
	//	fmt.Println("struc  ", stut)
	for i := 0; i < len(arr_str); i++ {
		col := arr_str[i]
		val := arr_vals[i]
		//f := s.Field(i)
		//f := col
		//filedName := stut.Field(i).Name
		filedName := col
		//		fmt.Printf("f: %v - %T\n",f,f)
		//		fmt.Printf("filed: %v --- %T \n", filedName, filedName)
		sqlCol := ""
		//		f.t
		switch val.(type) {
		case int:
			sqlCol = fmt.Sprintf("  `%v` int(10) DEFAULT '0' ", filedName)
		case string:
			sqlCol = fmt.Sprintf("  `%v` varchar(250) DEFAULT ''", filedName)
		}
		cols = append(cols, sqlCol)

	}
	//	fmt.Println(cols)
	sql1 := fmt.Sprintf(strings.Join(cols, ",\n"))

	sql2 := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%v` (\n%v \n) ENGINE=Aria DEFAULT CHARSET=utf8mb4;\n", table, sql1)
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
	for _, s := range str {
		//buf.WriteString(s)
		buf = append(buf, "'"+MySqlEscape(s)+"'")
	}
	return strings.Join(buf, ",")
}
