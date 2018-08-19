package filters

import (
	"errors"
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
func NewLocationsFilter(ctx *fasthttp.RequestCtx) (*FilterLocations, error) {
	f := new(FilterLocations)
	if ctx.QueryArgs().Has("country") {
		country := ctx.QueryArgs().Peek("country")
		if len(country) == 0 {
			return nil, errors.New("empty county")
		}
		cF := new(countryFilter)
		cF.country = string(country)
		f.ffs = append(f.ffs, cF)
	}
	if ctx.QueryArgs().Has("toDistance") {
		toDist := ctx.QueryArgs().Peek("toDistance")
		td, err := strconv.Atoi(string(toDist))
		if err != nil {
			return nil, err
		}
		tdf := new(toDistFilter)
		tdf.toDate = uint32(td)
		f.ffs = append(f.ffs, tdf)
	}

	return f, nil
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
