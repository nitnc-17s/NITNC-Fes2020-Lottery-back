package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) setRoutes(server *gin.Engine) {
	server.GET("/ws", func(ctx *gin.Context) {
		app.handleConnections(ctx.Writer, ctx.Request)
	})

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
