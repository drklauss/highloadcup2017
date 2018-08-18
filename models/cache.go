package models

import (
	"fmt"
	"time"
)

var (
	UCache   *Users
	VCache   *Visits
	LocCache *Locations
	UvlCache *UserVisitLinks
	LvlCache *LocationVisitLinks
)

func Init() {
	// Таблицы связей
	UvlCache = new(UserVisitLinks).Init()
	LvlCache = new(LocationVisitLinks).Init()
	// Основные таблицы
	t := time.Now()
	//var wg sync.WaitGroup
	//wg.Add(3)
	//go func() {
	LocCache = new(Locations).Init()
	//wg.Done()
	//}()
	//go func() {
	UCache = new(Users).Init()
	//wg.Done()
	//}()
	//go func() {
	VCache = new(Visits).Init()
	//wg.Done()
	//}()
	//wg.Wait()
	fmt.Printf("%+v", time.Since(t))
}
