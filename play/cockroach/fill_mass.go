package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"
    "time"
)

const (
    numberOfUsers = 1000
    numberOfGroups = 1000
    usersPerGroup = 100
    permissionsPerUser = 100
    permissionsPerGroup = 100
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=26257 dbname=b user=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	structure := `
	drop database b;
	create database b;
	use bench;
DROP TABLE IF EXISTS users;
CREATE TABLE if not EXISTS users (id INT PRIMARY KEY, name STRING ,name2 STRING);

DROP TABLE IF EXISTS groups;
CREATE TABLE groups (id INT PRIMARY KEY, name STRING);

DROP TABLE IF EXISTS permissions;
CREATE TABLE permissions (id INT PRIMARY KEY, rid STRING);

DROP TABLE IF EXISTS groups_permissions;
CREATE TABLE groups_permissions (id INT PRIMARY KEY, group_id INT, permission_id INT);
CREATE INDEX ON groups_permissions (group_id, permission_id);
CREATE INDEX ON groups_permissions (permission_id, group_id);

DROP TABLE IF EXISTS users_permissions;
CREATE TABLE users_permissions (id INT PRIMARY KEY, user_id INT, permission_id INT);
CREATE INDEX ON users_permissions (user_id, permission_id);
CREATE INDEX ON users_permissions (permission_id, user_id);

DROP TABLE IF EXISTS users_groups;
CREATE TABLE users_groups (id INT PRIMARY KEY, user_id INT, group_id INT);
CREATE INDEX ON users_groups (user_id, group_id);
CREATE INDEX ON users_groups (group_id, user_id);
`
	_, err = db.Exec(structure)
	if err != nil {
		panic(err)
	}

    mass2(db)

    id := 1
	for {
        s := ""
        for id2:=0; id2 < numberOfUsers; id2++ {
            //fmt.Println("UserView ", id)
            id++
            name := fmt.Sprintf("user-%d", id)
            //id2 := id
            s += fmt.Sprintf("INSERT INTO users VALUES (%d, '%s');", id, name)

        }

        _, err = db.Exec(s)
        if err != nil {
            panic(err)
        }
        fmt.Println("groups : ", id)
    }


	mid := 0
	for id := 0; id < numberOfGroups; id++ {
		fmt.Println("Group ", id)
		name := fmt.Sprintf("group-%d", id)
		_, err := db.Exec("INSERT INTO groups VALUES ($1, $2);", id, name)
		if err != nil {
			panic(err)
		}
		choices := rand.Perm(numberOfUsers)[:usersPerGroup]
		for _, uu := range choices {
			_, err := db.Exec("INSERT INTO users_groups VALUES ($1, $2, $3);", mid, uu, id)
			if err != nil {
				panic(err)
			}
			mid += 1
		}
	}

	rid := 0
	for id := 0; id < permissionsPerUser; id++ {
		fmt.Println("UserView Permissions ", id)
		name := fmt.Sprintf("rid-user-%d", rid)
		_, err := db.Exec("INSERT INTO permissions VALUES ($1, $2);", id, name)
		if err != nil {
			panic(err)
		}
		rid += 1
	}

	for id := 0; id < permissionsPerGroup; id++ {
		fmt.Println("Group Permissions ", id)
		name := fmt.Sprintf("rid-group-%d", id)
		_, err := db.Exec("INSERT INTO permissions VALUES ($1, $2);", rid, name)
		if err != nil {
			panic(err)
		}
		rid += 1
	}

	id2 := 0
	for uid := 0; uid < numberOfUsers; uid++ {
		fmt.Println("Permissions for user ", uid)
		for rid := 0; rid < permissionsPerUser; rid++ {
			_, err := db.Exec("INSERT INTO users_permissions VALUES ($1, $2, $3);", id2, uid, rid)
			if err != nil {
				panic(err)
			}
			id += 1
		}
	}

	id = 0
	for gid := 0; gid < numberOfGroups; gid++ {
		fmt.Println("Permissions for group ", gid)
		for rid := permissionsPerUser; rid < permissionsPerUser+permissionsPerGroup; rid++ {
			_, err := db.Exec("INSERT INTO groups_permissions VALUES ($1, $2, $3);", id, gid, rid)
			if err != nil {
				panic(err)
			}
			id += 1
		}
	}
}

func mass2(db *sql.DB)  {
    t :=time.Now()
    lastT :=time.Now()
    lastId :=0
    id := 1000000000
    for {
        o:=""
        for id2:=0; id2 < 1000; id2++ {
            //fmt.Println("UserView ", id)
            id++
            name := fmt.Sprintf("user-%d-sd yeryey ee eyery sads شسشسسش سیسسیی سیسییس ثقلثلثقل صشصفشص فقفف", id)
            name2 := fmt.Sprintf("useasd  فقفف", id)
            //id2 := id
            o += fmt.Sprintf("(%d, '%s' ,'%s'), ", id, name ,name2)

        }

        o = o[:len(o)-2]
        s := "INSERT INTO users (id, name, name2) VALUES " + o + ";"

        _, err := db.Exec(s)
        if err != nil {
            fmt.Println(s,err)
            panic(err)
        }
        if id %100000 == 0 {
            qpsAll:= id / int(time.Now().Sub(t).Seconds())
            qpslast:= (id-lastId) *1000 / int(time.Now().Sub(lastT).Nanoseconds()/1000000)

            fmt.Printf("id: %d -- qps:%d  - last: %d -- last rows count: %d \n", id,qpsAll,qpslast,(id-lastId))

            lastT = time.Now()
            lastId = id

        }
    }

}

