package handlers

import (
	"encoding/json"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

func CreateUpdateLocation(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("param").(string)
	if param == "new" {
		createLocation(ctx)
		return
	}
	updateLocation(ctx, &param)
}

func createLocation(ctx *fasthttp.RequestCtx) {
	l := new(models.Location)
	err := json.Unmarshal(ctx.PostBody(), l)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	uCache := models.VCache.Get(l.Id)
	if uCache != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	models.LocCache.Save(l)
	ctx.SetBodyString("{}")
}

func updateLocation(ctx *fasthttp.RequestCtx, param *string) {
	id, _ := strconv.Atoi(*param)
	l := models.LocCache.Get(uint32(id))
	if l == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	err := json.Unmarshal(ctx.PostBody(), &l)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	models.LocCache.Update(l)
}
