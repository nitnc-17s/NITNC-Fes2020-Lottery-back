package main

import (
	"github.com/comail/colog"
	"log"
	"lottery_back/pkg/config"
	"lottery_back/pkg/model"
	"lottery_back/pkg/server"
)

func main() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	// Load Config file
	cfg := config.LoadConfig()

	// Model Initialization
	model.Init(cfg)

	new(server.Server).Up(cfg)
}
