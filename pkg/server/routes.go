package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) setRoutes() {
	server.Engine.GET("/ws", func(ctx *gin.Context) {
		server.WsApp.wsHandler(ctx.Writer, ctx.Request)
	})

	server.Engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
