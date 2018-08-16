package handlers

import (
	"fmt"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

func GetUser(ctx *fasthttp.RequestCtx) {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	u := models.UCache.Get(uint32(id))
	if nil != u {
		fmt.Fprintf(ctx, "Welcome!\n %+v", u)
	} else {
		fmt.Fprint(ctx, "kjfhasdlfjhkljfh")
	}
}
