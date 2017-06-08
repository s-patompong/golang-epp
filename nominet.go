package main

import "github.com/qiangxue/fasthttp-routing"

func nominetRoutes(router *routing.Router) {
	nominet := router.Group("/nominet")

	nominetUkRoutes(nominet)

}
