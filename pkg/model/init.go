package model

import (
	"lottery_back/pkg/config"
)

func Init(config config.Config) {
	loadApplicants(config)
	loadPrizes(config)
	initResults()
}
