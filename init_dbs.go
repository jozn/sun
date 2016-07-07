package main

import (
	"fmt"
	//	."ms/sun/models"
	//	."ms/sun/actions"
	. "ms/sun/base"
	. "ms/sun/models"
)

func init_dbs() {
	//	DB.Exec(dbStructToTable(&RoomListTableSpec{},"RoomListTableSpec"))
	//	DB.Exec(dbStructToTable(&GroupMembersTableSpec{},"GroupMembersTableSpec"))
	//	DB.Exec(dbStructToTable(&GroupInfoTableSpec{},"GroupInfoTableSpec"))
	//	DB.Exec(dbStructToTable(&CustomRoomSettingsTableSpec{},"CustomRoomSettingsTableSpec"))
	//	DB.Exec(dbStructToTable(&BlockedUsersTableSpec{},"BlockedUsersTableSpec"))
	//	DB.Exec(dbStructToTable(&ContactsTableSpec{},"ContactsTableSpec"))
	//	res,err:=DB.Exec(dbStructToTable(&MessagesTableSpec{},"MessagesTableSpec"))
	//	fmt.Println("DATABASE CREATIONS",res,err)

	//	DB.Exec(DbStructToTable(&ContactsTableSpec{},"ContactsTableSpec"))
	DB.Exec(DbStructToTable(&PhoneContactTable{}, "phone_contacts"))
	DB.Exec(DbStructToTable(&PhoneContact{}, "phone_contacts2"))

	DB.Exec(DbStructToTable(&User{}, "user"))

	//chats
	DB.Exec(DbStructToTable(&GroupMember{}, "chat_group_member"))
	DB.Exec(DbStructToTable(&GroupInfo{}, "chat_group_info"))
	DB.Exec(DbStructToTable(&MessagesTable{}, "chat_message"))

	//	DB.Exec(DbStructToTable(&ContactUserInfo{},"chat_group_info"))

	fmt.Print("")
}
