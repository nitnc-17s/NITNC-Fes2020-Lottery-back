package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// App is websocket app data
type App struct {
	clients  map[*websocket.Conn]bool
	upgrader websocket.Upgrader
	sender   chan JSONData
	sync.RWMutex
}

// Generate is websocket app generate
func Generate() App {
	return App{
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

// handleConnections is called when websocket connection started
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

// handleMessages send message with websocket
func (app *App) handleMessages() {
	// 送信処理維持用ループ
	for {
		/*
			jsonData := <- app.sender

			jsonMsg, err := json.Marshal(jsonData)
			if err != nil {
				log.Printf("error: %v", err)
				continue
			}
		*/

		data := <-app.sender

		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range app.clients {
			err := client.WriteJSON(data)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(app.clients, client)
			}
		}
	}
}
