package server

import (
	"github.com/gorilla/websocket"
	"lottery_back/pkg/lottery"
	"net/http"
	"sync"
)

// App is websocket app data
type App struct {
	lotteryData *lottery.Data
	clients     map[*websocket.Conn]bool
	upgrader    websocket.Upgrader
	sync.RWMutex
}

// Generate is websocket app generate
func Generate() App {
	return App{
		lotteryData: lottery.Generate(),
		clients:     make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
