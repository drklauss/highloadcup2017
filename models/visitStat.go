package models

import (
	"sort"
)

// VisitStat урезанная структура посещений, используемая в операции получения песещенных мест пользователем
type VisitStat struct {
	Mark      uint8  `json:"mark"`
	VisitedAt uint32 `json:"visited_at"`
	Place     string `json:"place"`
}
type VisitStats struct {
	Visits []VisitStat `json:"visits"`
}

// Sort сортирует посещения по дате
func (vStats *VisitStats) Sort() *VisitStats {
	dTime := func(v1, v2 *VisitStat) bool {
		return v1.VisitedAt < v2.VisitedAt
	}
	By(dTime).Sort(vStats)
	return vStats
}

// By определяет сигнатуру функции сортировки
type By func(p1, p2 *VisitStat) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(visitStats *VisitStats) {
	ps := &visitsSorter{
		visitStats: *visitStats,
		by:         by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Сортировщик посещений
type visitsSorter struct {
	visitStats VisitStats
	by         func(p1, p2 *VisitStat) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *visitsSorter) Len() int {
	return len(s.visitStats.Visits)
}

// Swap is part of sort.Interface.
func (s *visitsSorter) Swap(i, j int) {
	s.visitStats.Visits[i], s.visitStats.Visits[j] = s.visitStats.Visits[j], s.visitStats.Visits[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *visitsSorter) Less(i, j int) bool {
	return s.by(&s.visitStats.Visits[i], &s.visitStats.Visits[j])
}
