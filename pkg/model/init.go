package model

import (
	"lottery_back/pkg/config"
)

func Init() {
	loadApplicants(config.ConfigData)
	loadPrizes(config.ConfigData)
	initResults()
}
