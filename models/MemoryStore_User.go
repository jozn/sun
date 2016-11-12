package models

import (
	"ms/sun/base"
	"ms/sun/helper"
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
	Row User
}

func (m *mapMemoryStore_UserImpl) ReloadAll() {
	m.MapUserToData = make(map[int]*userMemRowData, 10000)
	m.MapUsernameToId = make(map[string]int, 10000)
	m.MapSessionToId = make(map[string]int, 10000)

	users, err := NewUser_Selector().GetRows(base.DB)
	if err != nil {
		helper.DebugPrintln(err)
	}
	m.Lock()
	defer m.Unlock()

	for _, u := range users {
		memData := &userMemRowData{u}
		m.MapUserToData[u.Id] = memData
		if u.UserName != "" {
			m.MapUsernameToId[u.UserName] = u.Id
		}

		if u.SessionUuid != "" {
			m.MapSessionToId[u.SessionUuid] = u.Id
		}
	}
	helper.DebugPrintln("Loaded User to models.MemoryStore_User counts: ", len(m.MapUserToData))
}

func (m mapMemoryStore_UserImpl) GetUser(UserId int) (User, bool) {
	m.RLock()
	defer m.RUnlock()
	md, ok := m.MapUserToData[UserId]
	if ok {
		return md.Row, true
	}
	return User{}, false
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

func (m *mapMemoryStore_UserImpl) GetUserByUserName(UserName string) (User, bool) {
	m.RLock()
	defer m.RUnlock()
	uid, ok := m.MapUsernameToId[UserName]
	if ok {
		return m.GetUser(uid)
	}
    return User{}, false
}

func (m *mapMemoryStore_UserImpl) GetUserBySession(Session string) (User, bool) {
	m.RLock()
	defer m.RUnlock()
	uid, ok := m.MapSessionToId[Session]
	if ok {
		return m.GetUser(uid)
	}
    return User{}, false
}
