package controller

import (
	"github.com/gin-gonic/gin"
	"lottery_back/pkg/service"
)

func WebSocketConnector(ctx *gin.Context) {
	service.WebsocketApp.WebSocketHandler(ctx)
}
