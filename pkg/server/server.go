package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Up is Server Start
func (app *App) Up(url string) {
	server := gin.Default()

	// Load config file
	cfg := loadConfig()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		cfg.Server.Host,
	}
	server.Use(cors.New(corsConfig))

	//TODO ルーティング設定

	server.GET(url, func(ctx *gin.Context) {
		app.handleConnections(ctx.Writer, ctx.Request)
	})

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	go app.handleMessages()
	server.Run(":80")
}

func (app *App) handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := app.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer ws.Close()

	app.clients[ws] = true

	// WebSocket維持用ループ
	for {
		_, _, err := ws.ReadMessage() // 無を受け取る
		if err != nil {
			break
		}
	}
}

func (app *App) handleMessages() {
	// 送信処理維持用ループ
	for {
		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range app.clients {
			time.Sleep(time.Second / time.Duration((10 * len(app.clients))))
			err := client.WriteJSON(app.lotteryData.GetJSON())
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(app.clients, client)
			}
		}
	}
}
