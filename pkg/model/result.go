package model

import (
	"log"
	"math/rand"
	"time"
)

type Result struct {
	Prize  Prize     `json:"prize"`
	Winner Applicant `json:"winner"`
}

var results map[int]Result

func initResults() {
	results = make(map[int]Result)
}

func getResult(prizeId int) (Result, error) {
	res, ok := results[prizeId]

	// 既に抽選結果があるとき
	if ok {
		return res, nil
	}

	// 抽選結果がないとき
	prize, err := getPrize(prizeId)
	if err != nil {
		return Result{}, err
	}

	res = Result{Prize: prize}
	res.lottery()

	return res, nil
}

func (result *Result) lottery() {
	l := len(applicants)

	rand.Seed(time.Now().UnixNano())
	winnerId := rand.Intn(l)

	applicant, err := getApplicant(winnerId)
	if err != nil {
		log.Printf("warn: %v", err)
		return
	}

	result.Winner = applicant
}
