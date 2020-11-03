package main

import (
	"log"
	"lottery_back/pkg/server"

	"github.com/comail/colog"
)

func main() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	new(server.Server).Up()
}
