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

	router.GET("/users/:id", handlers.GetUser)
	router.GET("/users/:id/visits", handlers.GetUserVisits)
	router.GET("/locations/:id", handlers.GetUser)
	router.GET("/locations/:id/avg", handlers.GetUser)
	router.GET("/visits/:id", handlers.GetUser)

	panic(fasthttp.ListenAndServe(":8080", router.Handler))
}
