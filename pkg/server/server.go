package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Config config

	WsApp WsApp
}

// Up start server
func (server *Server) Up() {
	server.Engine = gin.Default()
	server.WsApp = GenerateWsApp()

	// Load Config file
	cfg := loadConfig()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Server.Host
	server.Engine.Use(cors.New(corsConfig))

	server.setRoutes()

	go server.WsApp.messageSender()

	server.Engine.Run(":8080")
}
