package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=26257 dbname=bench user=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	/*structure := `
		drop database b;
		create database b;
		use bench;
	DROP TABLE IF EXISTS users;
	CREATE TABLE users (id INT PRIMARY KEY, name STRING);

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
	`*/
	c := 0
	structure := `
	drop database b;
	create database b;
	use bench;
DROP TABLE IF EXISTS users;
CREATE TABLE users (id INT PRIMARY KEY, name STRING);

DROP TABLE IF EXISTS groups;
CREATE TABLE groups (id INT PRIMARY KEY, name STRING);

DROP TABLE IF EXISTS permissions;
CREATE TABLE permissions (id INT PRIMARY KEY, rid STRING);

DROP TABLE IF EXISTS groups_permissions;
CREATE TABLE groups_permissions (id INT PRIMARY KEY, group_id INT, permission_id INT);

DROP TABLE IF EXISTS users_permissions;
CREATE TABLE users_permissions (id INT PRIMARY KEY, user_id INT, permission_id INT);

DROP TABLE IF EXISTS users_groups;
CREATE TABLE users_groups (id INT PRIMARY KEY, user_id INT, group_id INT);
`
	_, err = db.Exec(structure)
	if err != nil {
		fmt.Println(err)
	}

	numberOfUsers := 1000
	numberOfGroups := 1000
	usersPerGroup := 100
	permissionsPerUser := 100
	permissionsPerGroup := 100
	i :=0
    next := func()int{
        i++
        return i
    }
	ff := func() {
		for id := 0; id < numberOfUsers; id++ {
			c++
			fmt.Println("User ", id)
			name := fmt.Sprintf("user-%d", next())
			_, err := db.Exec("INSERT INTO users VALUES ($1, $2);", next(), name)
			if err != nil {
				//panic(err)
			}
		}

		mid := 0
		for id := 0; id < numberOfGroups; id++ {
			fmt.Println("Group ", id, c)
			name := fmt.Sprintf("group-%d", next())
			c++
			_, err := db.Exec("INSERT INTO groups VALUES ($1, $2);", next(), name)
			if err != nil {
				//panic(err)
			}
			choices := rand.Perm(numberOfUsers)[:usersPerGroup]
			for _, _ = range choices {
				c++
				_, err := db.Exec("INSERT INTO users_groups VALUES ($1, $2, $3);", next(), next(), next())
				if err != nil {
					//panic(err)
				}
				mid += 1
			}
		}

		rid := 0
		for id := 0; id < permissionsPerUser; id++ {
			c++
			fmt.Println("User Permissions ", id)
			name := fmt.Sprintf("rid-user-%d", rid)
			_, err := db.Exec("INSERT INTO permissions VALUES ($1, $2);", next(), name)
			if err != nil {
				//panic(err)
			}
			rid += 1
		}

		for id := 0; id < permissionsPerGroup; id++ {
			c++
			fmt.Println("Group Permissions ", id, c)
			name := fmt.Sprintf("rid-group-%d", id)
			_, err := db.Exec("INSERT INTO permissions VALUES ($1, $2);", next(), name)
			if err != nil {
				//panic(err)
			}
			rid += 1
		}

		id := 0
		for uid := 0; uid < numberOfUsers; uid++ {
			fmt.Println("Permissions for user ", uid, c)
			for rid := 0; rid < permissionsPerUser; rid++ {
				c++
				_, err := db.Exec("INSERT INTO users_permissions VALUES ($1, $2, $3);", next(), uid, rid)
				if err != nil {
					//panic(err)
				}
				id += 1
			}
		}

		id = 0
		for gid := 0; gid < numberOfGroups; gid++ {
			fmt.Println("Permissions for group ", gid, c)
			for rid := permissionsPerUser; rid < permissionsPerUser+permissionsPerGroup; rid++ {
				_, err := db.Exec("INSERT INTO groups_permissions VALUES ($1, $2, $3);", next(), gid, rid)
				if err != nil {
					//panic(err)
				}
				id += 1
				c++
			}
		}
	}

	go ff()
	go ff()
	go ff()
	go ff()
	ff()
}
