package handlers

import (
	"encoding/json"
	"github.com/drklauss/highloadcup2017/filters"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

func GetUserVisits(ctx *fasthttp.RequestCtx) {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	u := models.UCache.Get(uint32(id))
	if u == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	fv, err := filters.NewVisitsFilter(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	fl, err := filters.NewLocationsFilter(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	vs := models.UvlCache.Get(uint32(id))
	resp := new(models.VisitStats)
	visitStat := new(models.VisitStat)
	oneVisit := new(models.Visit)
	oneLoc := new(models.Location)
	for _, v := range vs {
		oneVisit = models.VCache.Get(v)
		if !fv.Run(oneVisit) {
			continue

		}
		oneLoc = models.LocCache.Get(oneVisit.Location)
		if !fl.Run(oneLoc) {
			continue
		}
		visitStat.Mark = oneVisit.Mark
		visitStat.VisitedAt = oneVisit.VisitedAt
		visitStat.Place = models.LocCache.Get(oneVisit.Location).Place
		resp.Visits = append(resp.Visits, *visitStat)
	}

	m, _ := json.Marshal(resp.Sort())
	ctx.SetBody(m)
}
