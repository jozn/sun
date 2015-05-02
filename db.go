package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

//structRow: must be poniter
func dbInsertStruct(structRow interface{}, table string) (sql.Result, error) {
	return dbInsertUpdateStruct(structRow, table, true)
}

//ignors Id in all structs to set by db- or not at all
func dbUpdateStruct(structRow interface{}, table string) (sql.Result, error) {
	return dbInsertUpdateStruct(structRow, table, false)
}

//update a struct will by defult update dy Id
func dbInsertUpdateStruct(structRow interface{}, table string, isInsert bool) (sql.Result, error) {

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
		fmt.Println(query)
	}

	r, err := DB.Exec(query, vals...)

	if __DEV__ {
		if err != nil {
			log.Fatal("database err for table: ", table, " error: ", err)
		}
	}
	// noErr(err)
	return r, err

	// fmt.Println(cols, vals)
	// fmt.Println(cols, vals)
}
