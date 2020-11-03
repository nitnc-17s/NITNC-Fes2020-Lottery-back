package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	config config

	wsApp WsApp
}

// Up start server
func (server *Server) Up() {
	server.engine = gin.Default()
	server.wsApp = GenerateWsApp()

	// Load config file
	cfg := loadConfig()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Server.Host
	server.engine.Use(cors.New(corsConfig))

	server.setRoutes()

	go server.wsApp.messageSender()

	server.engine.Run(":8080")
}
