package main

import (
	"github.com/valyala/fasthttp"
)

func main() {
	router := routes()

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}
