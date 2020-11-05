package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// WsApp is websocket app data
type WsApp struct {
	clients  map[*websocket.Conn]bool
	upgrader websocket.Upgrader
	sender   chan JSONData
	sync.RWMutex
}

// JSONData is JSON formatted data
type JSONData interface{}

var WebsocketApp WsApp

// GenerateWsApp is websocket app generate
func GenerateWsApp() {
	WebsocketApp = WsApp{
		clients: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		sender: make(chan JSONData),
	}

	go WebsocketApp.messageSender()
}

// wsHandler is websocket connection handler
func (wsApp *WsApp) WebSocketHandler(ctx *gin.Context) {
	ws, err := wsApp.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("error: %v", err)
		ctx.String(http.StatusBadRequest, "Can't upgrade websocket\n", err)
		return
	}

	defer ws.Close()

	wsApp.clients[ws] = true

	// WebSocket維持用ループ
	for {
		_, _, err := ws.ReadMessage() // 無を受け取る
		if err != nil {
			log.Printf("warn: %v", err)
			break
		}
	}
}

// messageSender send message with websocket
func (wsApp *WsApp) messageSender() {
	// 送信処理維持用ループ
	for {
		/*
			jsonData := <- WsApp.sender

			jsonMsg, err := json.Marshal(jsonData)
			if err != nil {
				log.Printf("warn: %v", err)
				continue
			}
		*/

		data := <-wsApp.sender

		log.Println("debug: ", data) // For debug

		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range wsApp.clients {
			err := client.WriteJSON(data)
			if err != nil {
				log.Printf("warn: %v", err)
				client.Close()
				delete(wsApp.clients, client)
			}
		}
	}
}
