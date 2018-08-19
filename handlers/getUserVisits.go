package handlers

import (
	"github.com/drklauss/highloadcup2017/filters"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

func GetUserVisits(ctx *fasthttp.RequestCtx) {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	fv := filters.NewVisitFilter(ctx)
	vs := models.UvlCache.Get(uint32(id))
	for _, v := range vs {
		if fv.Run(models.VCache.Get(v)) {

		}
	}
}
