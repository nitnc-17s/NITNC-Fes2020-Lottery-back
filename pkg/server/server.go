package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lottery_back/pkg/config"
	"lottery_back/pkg/service"
)

type Server struct {
	Engine *gin.Engine
}

// Up start server
func (server *Server) Up() {
	server.Engine = gin.Default()
	service.GenerateWsApp()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.ConfigData.Server.Host
	server.Engine.Use(cors.New(corsConfig))

	server.setRoutes()

	server.Engine.Run(":8080")
}
