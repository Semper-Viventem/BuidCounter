package handlers

import "github.com/gin-gonic/gin"

type RequestHandler struct {
	Patch string
	Handler func(c *gin.Context)
}
