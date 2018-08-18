package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
)

func NewVisit(ctx *fasthttp.RequestCtx) {
	var v models.Visit
	fmt.Printf("%+v", string(ctx.PostBody()))
	err := json.Unmarshal(ctx.PostBody(), &v)
	if err != nil {
		fmt.Printf("%+v", err.Error())
	}
	models.VCache.Save(&v)
	fmt.Printf("\n%+v", models.UvlCache.Get(v.User))
	fmt.Printf("\n%+v", models.LvlCache.Get(v.Location))
}
