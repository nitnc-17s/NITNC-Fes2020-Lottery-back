package model

import "lottery_back/pkg/server"

func Init(server *server.Server) {
	loadApplicants(server)
	loadPrizes(server)
	initResults()
}
