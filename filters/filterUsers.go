package filters

import (
	"errors"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type FilterUsers struct {
	ffs []filterUserFunc
}

type filterUserFunc interface {
	filter(*models.User) bool
}

type toAgeFilter struct {
	toAge int32
	f     filterUserFunc
}

func (taf *toAgeFilter) filter(u *models.User) bool {
	return u.BirthDate < taf.toAge
}

type fromAgeFilter struct {
	fromAge int32
	f       filterUserFunc
}

func (faf *fromAgeFilter) filter(u *models.User) bool {
	return u.BirthDate > faf.fromAge
}

type genderFilter struct {
	gender string
	f      filterUserFunc
}

func (faf *genderFilter) filter(u *models.User) bool {
	return u.Gender == faf.gender
}

// NewVisitsFilter создает экземпляр фильтра
func NewUsersFilter(ctx *fasthttp.RequestCtx) (*FilterUsers, error) {
	time.Now()
	f := new(FilterUsers)
	if ctx.QueryArgs().Has("fromAge") {
		fd := ctx.QueryArgs().Peek("fromAge")
		fTime, err := strconv.Atoi(string(fd))
		if err != nil {
			return nil, err
		}
		faf := new(fromAgeFilter)
		faf.fromAge = int32(fTime)
		f.ffs = append(f.ffs, faf)
	}
	if ctx.QueryArgs().Has("toAge") {
		td := ctx.QueryArgs().Peek("toAge")
		tTime, err := strconv.Atoi(string(td))
		if err != nil {
			return nil, err
		}
		taf := new(toAgeFilter)
		taf.toAge = int32(tTime)
		f.ffs = append(f.ffs, taf)
	}
	if ctx.QueryArgs().Has("gender") {
		g := string(ctx.QueryArgs().Peek("gender"))
		if !(g == "m" || g == "f") {
			return nil, errors.New("wrong gender")
		}
		tdf := new(genderFilter)
		tdf.gender = string(g)
		f.ffs = append(f.ffs, tdf)
	}
	return f, nil
}

// Run запускает систему фильтров
// Если вернул true - удовлетворяет фильтру, false - нет
func (f *FilterUsers) Run(v *models.User) bool {
	for _, ff := range f.ffs {
		if !ff.filter(v) {
			return false
		}
	}

	return true
}
