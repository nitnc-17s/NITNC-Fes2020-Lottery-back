package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) setRoutes() {
	server.engine.GET("/ws", func(ctx *gin.Context) {
		server.wsApp.wsHandler(ctx.Writer, ctx.Request)
	})

	server.engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
