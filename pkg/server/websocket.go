package server

import (
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

// GenerateWsApp is websocket app generate
func GenerateWsApp() WsApp {
	return WsApp{
		clients: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		sender: make(chan JSONData),
	}
}

// JSONData is JSON formatted data
type JSONData interface{} // もう少しマシな書き方がある気がする

// wsHandler is websocket connection handler
func (wsApp *WsApp) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := wsApp.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer ws.Close()

	wsApp.clients[ws] = true

	// WebSocket維持用ループ
	for {
		_, _, err := ws.ReadMessage() // 無を受け取る
		if err != nil {
			break
		}
	}
}

// messageSender send message with websocket
func (wsApp *WsApp) messageSender() {
	// 送信処理維持用ループ
	for {
		/*
			jsonData := <- wsApp.sender

			jsonMsg, err := json.Marshal(jsonData)
			if err != nil {
				log.Printf("error: %v", err)
				continue
			}
		*/

		data := <-wsApp.sender

		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range wsApp.clients {
			err := client.WriteJSON(data)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(wsApp.clients, client)
			}
		}
	}
}
