package models

import "sync"

// Кэш связей достопримечателльности <-> посещения
type LocationVisitLinks struct {
	mx sync.RWMutex
	m  map[uint32][]uint32
}

func (lvl *LocationVisitLinks) Get(locId uint32) []uint32 {
	lvl.mx.RLock()
	defer lvl.mx.RUnlock()
	return lvl.m[locId]
}

func (lvl *LocationVisitLinks) Save(lId uint32, vId uint32) {
	lvl.mx.Lock()
	lvl.m[lId] = append(lvl.m[lId], vId)
	lvl.mx.Unlock()
}
