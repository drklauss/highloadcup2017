package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type User struct {
	Id        uint32 `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int32  `json:"birth_date"`
}

// Кэш данных пользователей
type Users struct {
	mx sync.RWMutex
	m  map[uint32]User
}

func (us *Users) Init() *Users {
	us.m = make(map[uint32]User)
	us.readData()
	return us
}

func (us *Users) Get(id uint32) *User {
	us.mx.RLock()
	defer us.mx.RUnlock()
	if v, ok := us.m[id]; ok {
		return &v
	}
	return nil
}

func (us *Users) Save(u *User) {
	us.mx.Lock()
	us.m[u.Id] = *u
	us.mx.Unlock()
}

func (us *Users) Update(u *User) {
	us.mx.Lock()
	us.m[u.Id] = *u
	us.mx.Unlock()
}

func (us *Users) readData() {
	t := time.Now()
	count := 1
	type users struct {
		Users []User `json:"users"`
	}
	for {
		fName := fmt.Sprintf("data/users_%d.json", count)
		fmt.Println(fName)
		b, err := ioutil.ReadFile(fName)
		if err != nil {
			println(err.Error())
			break
		}
		var users users
		err = json.Unmarshal(b, &users)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		for _, u := range users.Users {
			us.m[u.Id] = u
		}
		count++
	}
	fmt.Printf("All Users: %d\nUsers Time:%+v\n", len(us.m), time.Since(t))
}
