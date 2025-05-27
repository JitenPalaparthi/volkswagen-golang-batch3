package handlers

import (
	"github.com/gin-gonic/gin"
)

func init() {
	println("handler calling --1 in handlers")
}

func init() {
	println("Handler calling --2 in handlers")
}

func Root(ctx *gin.Context) {
	ctx.String(200, "Hello World!")
}
func Ping(ctx *gin.Context) {
	ctx.String(200, "Pong")
}

func Health(ctx *gin.Context) {
	//ctx.JSON(200, map[string]any{"health": "ok"})
	ctx.JSON(200, gin.H{"health": "ok"})
}
