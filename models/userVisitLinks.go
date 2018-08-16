package models

import "sync"

// Кэш связей пользователь <-> посещения
type UserVisitLinks struct {
	mx sync.RWMutex
	m  map[uint32][]uint32
}

// Get возвращает []visitId
func (uvl *UserVisitLinks) Get(userId uint32) []uint32 {
	uvl.mx.RLock()
	defer uvl.mx.RUnlock()
	return uvl.m[userId]
}

// Save сохраняет посещение пользователя
func (uvl *UserVisitLinks) Save(uId uint32, vId uint32) {
	uvl.mx.Lock()
	uvl.m[uId] = append(uvl.m[uId], vId)
	uvl.mx.Unlock()
}
