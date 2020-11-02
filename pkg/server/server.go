package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Up start server
func (app *App) Up() {
	server := gin.Default()

	// Load config file
	cfg := loadConfig()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Server.Host
	server.Use(cors.New(corsConfig))

	app.setRoutes(server)

	go app.handleMessages()

	server.Run(":8080")
}
