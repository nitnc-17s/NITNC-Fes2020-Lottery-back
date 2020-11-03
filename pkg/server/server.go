package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Up start server
func Up() {
	server := gin.Default()
	wsApp := GenerateWsApp()

	// Load config file
	cfg := loadConfig()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Server.Host
	server.Use(cors.New(corsConfig))

	wsApp.setRoutes(server)

	go wsApp.messageSender()

	server.Run(":8080")
}
