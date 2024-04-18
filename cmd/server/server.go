package server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type server struct {
	http   *fasthttp.Server
	router *router.Router
}
