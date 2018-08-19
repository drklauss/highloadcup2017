package filters

import (
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

type FilterLocations struct {
	ffs []filterLocationFunc
}

type filterLocationFunc interface {
	filter(*models.Location) bool
}

type toDistFilter struct {
	toDate uint32
	f      filterLocationFunc
}

func (tdf *toDistFilter) filter(l *models.Location) bool {
	return l.Distance < tdf.toDate
}

type countryFilter struct {
	country string
	f       filterLocationFunc
}

func (cF *countryFilter) filter(v *models.Location) bool {
	return v.Country == cF.country
}

// NewVisitsFilter создает экземпляр фильтра
func NewLocationsFilter(ctx *fasthttp.RequestCtx) *FilterLocations {
	f := new(FilterLocations)
	country := ctx.QueryArgs().Peek("country")
	toDist := ctx.QueryArgs().Peek("toDistance")
	if len(country) > 0 {
		cF := new(countryFilter)
		cF.country = string(country)
		f.ffs = append(f.ffs, cF)
	}
	if len(toDist) > 0 {
		td, _ := strconv.Atoi(string(toDist))
		tdf := new(toDistFilter)
		tdf.toDate = uint32(td)
		f.ffs = append(f.ffs, tdf)
	}
	return f
}

// Run запускает систему фильтров
// Если вернул true - удовлетворяет фильтру, false - нет
func (f *FilterLocations) Run(v *models.Location) bool {
	for _, ff := range f.ffs {
		if !ff.filter(v) {
			return false
		}
	}

	return true
}
