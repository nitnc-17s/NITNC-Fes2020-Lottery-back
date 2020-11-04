package controller

import (
	"github.com/gin-gonic/gin"
	"lottery_back/pkg/config"
	"lottery_back/pkg/service"
	"net/http"
)

type OperationPostRequest struct {
	ApiKey    string `json:"api_key"`
	PrizeId   int    `json:"prize_id"`
	Operation string `json:"operation"`
}

func OperationReceiver(ctx *gin.Context) {
	req := OperationPostRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		ctx.String(http.StatusBadRequest, "JSON format is wrong\n", err)
		return
	}

	if req.ApiKey != config.ConfigData.Api.Key {
		ctx.String(http.StatusForbidden, "You don't have Permission")
		return
	}

	err = service.WebSocketSender(req.PrizeId, req.Operation)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Something is wrong\n", err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"status": "ok",
	})
}
