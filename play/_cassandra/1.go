package main

import (
    "github.com/gocql/gocql"
    "log"
    "fmt"
)

func main()  {
    cluster := gocql.NewCluster("127.0.0.1")//, "192.168.1.2", "192.168.1.3")
    cluster.Keyspace = "store"
    session, err := cluster.CreateSession()
    if(err != nil){
        log.Fatal(err)
    }
    fmt.Println(session)
    fmt.Println(session.Query("insert into one (id,body) values (?,?)", 12,"sd"))
    //fmt.Println(session.Query("select * from files.media"))
    r := session.Query("select * from one")
    fmt.Println(r.Exec())
    //session
}