package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/drklauss/highloadcup2017/handlers"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
)

func main() {
	models.Init()
	router := fasthttprouter.New()
	// Пользователи
	router.GET("/users/:id", jsonResponse(handlers.GetUser))
	router.GET("/users/:id/visits", jsonResponse(handlers.GetUserVisits))
	router.POST("/users/:param", jsonResponse(handlers.GetUser))
	// Достопримечательности
	router.GET("/locations/:id", jsonResponse(handlers.GetUser))
	router.POST("/locations/:param", jsonResponse(handlers.GetUser))
	router.GET("/locations/:id/avg", jsonResponse(handlers.GetUser))
	// Посещения
	router.GET("/visits/:id", jsonResponse(handlers.GetUser))
	router.POST("/visits/:param", jsonResponse(handlers.GetUser))

	panic(fasthttp.ListenAndServe(":8080", router.Handler))
}

func jsonResponse(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")
		h(ctx)
	})
}
