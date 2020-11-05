package server

import (
	"github.com/gin-gonic/gin"
	"lottery_back/pkg/controller"
	"net/http"
)

func (server *Server) setRoutes() {
	server.Engine.GET("/ws", controller.WebSocketConnector)
	server.Engine.POST("/operation", controller.OperationReceiver)

	server.Engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
