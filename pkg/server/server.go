package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"lottery_back/pkg/config"
	"lottery_back/pkg/model"
	"time"
)

type Server struct {
	Engine *gin.Engine
	WsApp  WsApp
}

// Up start server
func (server *Server) Up(cfg config.Config) {
	server.Engine = gin.Default()
	server.WsApp = GenerateWsApp()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Server.Host
	server.Engine.Use(cors.New(corsConfig))

	server.setRoutes()

	go server.WsApp.messageSender()

	go server.testSender() // For test Run

	server.Engine.Run(":8080")
}

func (server *Server) testSender() {
	time.Sleep(time.Second * 5)

	for i := 1; i < 100; i++ {
		result, err := model.GetResult(i)
		if err != nil {
			log.Printf("warn: %v", err)
		}

		server.WsApp.sender <- result
		time.Sleep(time.Second * 10)
	}
}
