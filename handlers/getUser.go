package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

func GetUser(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	u := models.UCache.Get(uint32(id))
	if u == nil {
		ctx.Error("Not Found", fasthttp.StatusNotFound)
		return
	}

	fmt.Printf("\n%+v", u)
	visits := models.UvlCache.Get(u.Id)
	for _, oneV := range visits {
		fmt.Printf("\n%+v", models.VCache.Get(oneV))
	}

	m, _ := json.Marshal(u)
	ctx.SetBody(m)
}
