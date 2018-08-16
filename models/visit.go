package models

import (
	"sync"
)

type Visit struct {
	Id        uint32 `json:"id"`
	Location  uint32 `json:"location"`
	User      uint32 `json:"user"`
	VisitedAt uint32 `json:"visited_at"`
	Mark      uint8  `json:"mark"`
}

// Кэш данных посещений
type Visits struct {
	mx sync.RWMutex
	m  map[uint32]*Visit
}

func (vs *Visits) Get(id uint32) *Visit {
	vs.mx.RLock()
	defer vs.mx.RUnlock()
	return vs.m[id]
}

func (vs *Visits) Save(v *Visit) {
	vs.mx.Lock()
	vs.m[v.Id] = v
	UvlCache.Save(v.User, v.Id)
	vs.mx.Unlock()
}

func (vs *Visits) Update(id uint32, v *Visit) {
	vs.mx.Lock()
	v.Id = id
	vs.m[id] = v
	vs.mx.Unlock()
}
