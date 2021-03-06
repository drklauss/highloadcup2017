package handlers

import (
	"encoding/json"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

func GetVisit(ctx *fasthttp.RequestCtx) {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	u := models.VCache.Get(uint32(id))
	if u == nil {
		ctx.Error("Not Found", fasthttp.StatusNotFound)
		return
	}
	m, _ := json.Marshal(u)
	ctx.SetBody(m)
}
