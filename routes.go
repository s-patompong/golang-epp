package main

import "github.com/qiangxue/fasthttp-routing"

func routes() *routing.Router {
	router := routing.New()

	nominetRoutes(router)

	return router
}
