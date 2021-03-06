package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"strings"
	"sync"
)

type mapMemoryStore_UserImpl struct {
	MapUserToData   map[int]*userMemRowData
	MapUsernameToId map[string]int
	MapSessionToId  map[string]int
	sync.RWMutex
}

var MemoryStore_User *mapMemoryStore_UserImpl

func init() {
	MemoryStore_User = &mapMemoryStore_UserImpl{}
	MemoryStore_User.MapUserToData = make(map[int]*userMemRowData)
}

type userMemRowData struct {
	Row x.User
}

func (m *mapMemoryStore_UserImpl) ReloadAll() {
	m.MapUserToData = make(map[int]*userMemRowData, 10000)
	m.MapUsernameToId = make(map[string]int, 10000)
	m.MapSessionToId = make(map[string]int, 10000)

	users, err := x.NewUser_Selector().GetRows2(base.DB)
	if err != nil {
		helper.DebugPrintln(err)
	}
	m.Lock()
	defer m.Unlock()

	for _, u := range users {
		memData := &userMemRowData{u}
		m.MapUserToData[u.Id] = memData
		if u.UserName != "" {
			m.MapUsernameToId[strings.ToLower(u.UserName)] = u.Id
		}

		if u.SessionUuid != "" {
			m.MapSessionToId[u.SessionUuid] = u.Id
		}
	}
	helper.DebugPrintln("Loaded UserView to models.MemoryStore_User counts: ", len(m.MapUserToData))
}

func (m mapMemoryStore_UserImpl) GetUser(UserId int) (x.User, bool) {
	m.RLock()
	defer m.RUnlock()
	md, ok := m.MapUserToData[UserId]
	if ok {
		return md.Row, true
	}
	return x.User{}, false
}

func (m *mapMemoryStore_UserImpl) GetMemRow(UserId int) (*userMemRowData, bool) {
	m.RLock()
	defer m.RUnlock()
	md, ok := m.MapUserToData[UserId]
	if ok {
		return md, true
	}
	return nil, false
}

func (m *mapMemoryStore_UserImpl) GetUserByUserName(UserName string) (x.User, bool) {
	m.RLock()
	defer m.RUnlock()
	UserName = strings.Replace(UserName, "@", "", 1)
	uid, ok := m.MapUsernameToId[strings.ToLower(UserName)]
	if ok {
		return m.GetUser(uid)
	}
	return x.User{}, false
}

func (m *mapMemoryStore_UserImpl) GetUserBySession(Session string) (x.User, bool) {
	m.RLock()
	defer m.RUnlock()
	uid, ok := m.MapSessionToId[Session]
	if ok {
		return m.GetUser(uid)
	}
	return x.User{}, false
}
