package handlers

import (
	"encoding/json"
	"github.com/drklauss/highloadcup2017/filters"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"math"
	"strconv"
)

func GetLocationAvgMark(ctx *fasthttp.RequestCtx) {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	uid := uint32(id)
	l := models.LocCache.Get(uid)
	if l == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	fu, err := filters.NewUsersFilter(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	fv, err := filters.NewVisitsFilter(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	var markSum, count uint8
	vs := models.LvlCache.Get(uint32(uid))
	oneVisit := new(models.Visit)
	oneUser := new(models.User)
	//
	for _, v := range vs {
		oneVisit = models.VCache.Get(v)
		if !fv.Run(oneVisit) {
			continue
		}
		oneUser = models.UCache.Get(oneVisit.User)
		if !fu.Run(oneUser) {
			continue
		}
		count++
		markSum += oneVisit.Mark
	}
	if count == 0 {
		ctx.SetBodyString("{\"avg\":0}")
		return
	}
	var avg struct {
		Avg float64 `json:"avg"`
	}
	avg.Avg = math.Round(float64(markSum)/float64(count)*100) / 100
	m, _ := json.Marshal(avg)
	ctx.SetBody(m)
}
