package models

// todo при вычитывании данных визита предусмотреть сохранение в таблицу locationVisitLinks и userVisitLinks
type VisitStat struct {
	Mark      uint8  `json:"mark"`
	VisitedAt uint32 `json:"visited_at"`
	Place     string `json:"place"`
}
type VisitStats struct {
	Visits []VisitStat `json:"visits"`
}

func (vStats *VisitStats) Sort() *VisitStats {

	return vStats
}
