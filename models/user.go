package models

import (
	"sync"
	"time"
)

type User struct {
	Id        uint32    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
}

// Кэш данных пользователей
type Users struct {
	mx sync.RWMutex
	m  map[uint32]*User
}

func (us *Users) Get(id uint32) *User {
	us.mx.RLock()
	defer us.mx.RUnlock()
	if v, ok := us.m[id]; ok {
		return v
	}
	return nil
}

func (us *Users) Save(u *User) {
	us.mx.Lock()
	us.m[u.Id] = u
	us.mx.Unlock()
}

func (us *Users) Update(id uint32, u *User) {
	us.mx.Lock()
	u.Id = id
	us.m[id] = u
	us.mx.Unlock()
}
