package main

import (
	"github.com/gin-gonic/gin"
	"./handlers"
)

var api = gin.Default()

var methods = []*handlers.RequestHandler{
	&handlers.NewBuildHandler,
	&handlers.BuildsCountHandler,
}

func main() {

	for _, method := range methods {
		api.GET(method.Patch, method.Handler)
	}

	api.Run()
}
