package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

var buildCountersMap = map[string]int { }

var keyBuilder = "key_builder"
var NewBuildHandler = RequestHandler{
	Patch: "/newBuild",
	Handler: func(c *gin.Context) {
		key := c.Query(keyBuilder)
		fmt.Println("Param: " + key)
		if _, ok := buildCountersMap[key]; ok {
			buildCountersMap[key]++
		} else {
			buildCountersMap[key] = 1
		}
		c.String(http.StatusCreated, "OK")
	},
}

var BuildsCountHandler = RequestHandler{
	Patch: "/buildsCount",
	Handler: func(c *gin.Context) {
		key := c.Query(keyBuilder)
		fmt.Println("Param: " + key)
		count := buildCountersMap[key]
		c.String(http.StatusOK, "Build count: %d", count)
	},
}
