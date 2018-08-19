package handlers

import (
	"encoding/json"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

func CreateUpdateVisit(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("param").(string)
	if param == "new" {
		createVisit(ctx)
		return
	}
	updateVisit(ctx, &param)
}

func createVisit(ctx *fasthttp.RequestCtx) {
	v := new(models.Visit)
	err := json.Unmarshal(ctx.PostBody(), v)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	uCache := models.VCache.Get(v.Id)
	if uCache != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	models.VCache.Save(v)
	ctx.SetBodyString("{}")
}

func updateVisit(ctx *fasthttp.RequestCtx, param *string) {
	id, _ := strconv.Atoi(*param)
	v := models.VCache.Get(uint32(id))
	if v == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	err := json.Unmarshal(ctx.PostBody(), &v)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	models.VCache.Update(v)
}
