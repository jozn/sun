package models

import (
    "database/sql"
    "ms/sun/helper"
    "strings"
    "strconv"
    "github.com/jmoiron/sqlx"
)

type deleter struct  {
    where []whereClause
    whereSep string
    table string
}

type whereClause struct  {
    condition string
    args      []interface{}
}

//////
type _User2Deleter struct  {
    wheres   []whereClause
    whereSep string
}

type _User2Updater struct  {
    wheres   []whereClause
    updates   map[string]interface{}
    whereSep string
}

type _User2Selector struct  {
    wheres   []whereClause
    selects  []string
    whereSep  string
    orderBy string//" order by id desc //for ints
    limit int
    offset int
}


func User_Deleter()  *_User2Deleter {
    d := &_User2Deleter{whereSep: " AND "}
    return d
}


func (u *_User2Deleter)Id_In (ins []int) *_User2Deleter {
    w := whereClause{}
    var insWhere []interface{}
    for _, i:= range ins {
        insWhere = append(insWhere,i)
    }
    w.args = insWhere
    w.condition = " Id IN("+helper.DbQuestionForSqlIn(len(ins))+") "
    u.wheres = append(u.wheres, w)

    return u
}

func (u *_User2Deleter)Id_NotIn (ins []int) *_User2Deleter {
    w := whereClause{}
    var insWhere []interface{}
    for _, i:= range ins {
        insWhere = append(insWhere,i)
    }
    w.args = insWhere
    w.condition = " Id NOT IN("+helper.DbQuestionForSqlIn(len(ins))+") "
    u.wheres = append(u.wheres, w)

    return u
}

func (u *_User2Deleter)Id_EQ (val int) *_User2Deleter {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Id = ? "
    u.wheres = append(u.wheres, w)

    return u
}

func (u *_User2Deleter)Id_NotEQ (val int) *_User2Deleter {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Id != ? "
    u.wheres = append(u.wheres, w)

    return u
}

/////////////////// selecter

func (s *_User2Selector)Limit(num int) *_User2Selector {
    s.limit = num
    return s
}

func (s *_User2Selector)Offset(num int) *_User2Selector {
    s.offset = num
    return s
}

func (s *_User2Selector)OrderBy_Id_Desc(num int) *_User2Selector {
    s.orderBy = " order by Id DESC "
    return s
}

func (u *_User2Selector)Get(db sqlx.DB) *_User2Selector {
    q:=""
    if u.orderBy != ""{
        q += u.orderBy
    }

    if u.limit != 0 {
        q += " LIMIT " + strconv.Itoa(u.limit)
    }

    if u.limit != 0 {
        q += " OFFSET " + strconv.Itoa(u.offset)
    }

    u.orderBy = " order by Id DESC "

    //db.Select()
    return u
}

///updater///////////////////////////////////
func (u *_User2Updater)Id (newVal int) *_User2Updater {
    u.updates["Id"] = newVal
    return u
}

func (u *_User2Updater)Name (newVal string) *_User2Deleter {
    u.updates[" ( Name = ? )"] = newVal
    return u
}

func (u *_User2Updater)Update (db sql.DB)(int,error) {
    var err error

    var updateArgs []interface{}
    sqlUpdate := ""
    for up, newVal := range u.updates {
        sqlUpdate += up
        updateArgs = append(updateArgs, newVal)
    }

    sqlWherrs, whereArgs := whereClusesToSql(u.wheres,u.whereSep)

    var allArgs []interface{}
    allArgs = append(allArgs, updateArgs...)
    allArgs = append(allArgs, whereArgs...)

    sqlstr := `UPDATE user2 ` + sqlUpdate +" " +sqlWherrs

    //XOLog(sqlstr,allArgs...)
    _, err = db.Exec(sqlstr, allArgs)
    return err
}

func (u *_User2Deleter)Delete (db sql.DB) (int,error) {
    var err error
    var wheresArr []string
    for _,w := range u.wheres{
        wheresArr = append(wheresArr,w.condition)
    }
    wheresStr := strings.Join(wheresArr, u.whereSep)

    var args []interface{}
    for _,w := range u.wheres{
        args = append(args,w.args...)
    }

    sqlstr := "DELETE FROM _user2 WHERE " + wheresStr

    // run query
    //XOLog(sqlstr, args...)
    res, err := db.Exec(sqlstr, args...)
    if err != nil {
        return 0,err
    }

    // retrieve id
    num, err := res.RowsAffected()
    if err != nil {
        return 0,err
    }

    return int(num),nil

}

func whereClusesToSql(wheres []whereClause, whereSep string ) (string, []interface{}) {
    var wheresArr []string
    for _,w := range wheres{
        wheresArr = append(wheresArr,w.condition)
    }
    wheresStr := strings.Join(wheresArr, whereSep)

    var args []interface{}
    for _,w := range wheres{
        args = append(args,w.args...)
    }
    return wheresStr , args
}

