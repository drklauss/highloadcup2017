package models

import "sync"

type Location struct {
	Id       uint32 `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance uint32 `json:"distance"`
}

// Кэш данных достопримечательностей
type Locations struct {
	mx sync.RWMutex
	m  map[uint32]*Location
}

func (ls *Locations) Get(id uint32) *Location {
	ls.mx.RLock()
	defer ls.mx.RUnlock()
	return ls.m[id]
}

func (ls *Locations) Save(l *Location) {
	ls.mx.Lock()
	ls.m[l.Id] = l
	ls.mx.Unlock()
}

func (ls *Locations) Update(id uint32, l *Location) {
	ls.mx.Lock()
	l.Id = id
	ls.m[id] = l
	ls.mx.Unlock()
}
