package controller

import (
	"github.com/gin-gonic/gin"
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
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	// TODO 処理を書く

	ctx.JSON(http.StatusNoContent, gin.H{
		"status": "ok",
	})
}
