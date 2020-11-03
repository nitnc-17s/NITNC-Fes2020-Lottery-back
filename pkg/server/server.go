package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lottery_back/pkg/config"
)

type Server struct {
	Engine *gin.Engine

	WsApp WsApp
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

	server.Engine.Run(":8080")
}
