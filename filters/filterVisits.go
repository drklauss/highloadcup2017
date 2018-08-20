package filters

import (
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

type FilterVisits struct {
	ffs []filterVisitFunc
}

type filterVisitFunc interface {
	filter(*models.Visit) bool
}

type toDateFilter struct {
	toDate uint32
	f      filterVisitFunc
}

func (tdf *toDateFilter) filter(v *models.Visit) bool {
	return v.VisitedAt < tdf.toDate
}

type fromDateFilter struct {
	fromDate uint32
	f        filterVisitFunc
}

func (tdf *fromDateFilter) filter(v *models.Visit) bool {
	return v.VisitedAt > tdf.fromDate
}

// NewVisitsFilter создает экземпляр фильтра
func NewVisitsFilter(ctx *fasthttp.RequestCtx) (*FilterVisits, error) {
	f := new(FilterVisits)
	if ctx.QueryArgs().Has("fromDate") {
		fd := ctx.QueryArgs().Peek("fromDate")
		fTime, err := strconv.Atoi(string(fd))
		if err != nil {
			return nil, err
		}
		fdf := new(fromDateFilter)
		fdf.fromDate = uint32(fTime)
		f.ffs = append(f.ffs, fdf)
	}
	if ctx.QueryArgs().Has("toDate") {
		td := ctx.QueryArgs().Peek("toDate")
		tTime, err := strconv.Atoi(string(td))
		if err != nil {
			return nil, err
		}
		tdf := new(toDateFilter)
		tdf.toDate = uint32(tTime)
		f.ffs = append(f.ffs, tdf)
	}
	return f, nil
}

// Run запускает систему фильтров
// Если вернул true - удовлетворяет фильтру, false - нет
func (f *FilterVisits) Run(v *models.Visit) bool {
	for _, ff := range f.ffs {
		if !ff.filter(v) {
			return false
		}
	}

	return true
}
