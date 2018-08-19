package handlers

import (
	"encoding/json"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

func CreateUpdateUser(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("param").(string)
	if param == "new" {
		createUser(ctx)
		return
	}
	updateUser(ctx, &param)
}

func createUser(ctx *fasthttp.RequestCtx) {
	u := new(models.User)
	err := json.Unmarshal(ctx.PostBody(), u)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	uCache := models.UCache.Get(u.Id)
	if uCache != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	models.UCache.Save(u)
	ctx.SetBodyString("{}")
}

func updateUser(ctx *fasthttp.RequestCtx, param *string) {
	id, _ := strconv.Atoi(*param)
	u := models.UCache.Get(uint32(id))
	if u == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	err := json.Unmarshal(ctx.PostBody(), &u)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	models.UCache.Update(u)
}
