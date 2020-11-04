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

func GetResult(prizeId int) (Result, error) {
	res, ok := results[prizeId]

	// 既に抽選結果があるとき
	if ok {
		return res, nil
	}

	// 抽選結果がないとき
	prize, err := GetPrize(prizeId)
	if err != nil {
		return Result{}, err
	}

	res = Result{Prize: prize}

	err = res.Lottery()
	if err != nil {
		return Result{}, err
	}

	return res, nil
}

func GetEmptyResult() Result {
	return Result{Prize: Prize{Id: -1, Name: ""}, Winner: Applicant{Name: "", NameFurigana: "", Class: ""}}
}

func (result *Result) Lottery() error {
	l := len(applicants)

	rand.Seed(time.Now().UnixNano())
	winnerId := rand.Intn(l)

	applicant, err := GetApplicant(winnerId)
	if err != nil {
		log.Printf("warn: %v", err)
		return err
	}

	result.Winner = applicant
	results[result.Prize.Id] = *result
	return nil
}

func (result *Result) GetPrizeMaskedResult() Result {
	response := result.GetWinnerMaskedResult()

	response.Prize.Name = ""
	return response
}

func (result *Result) GetWinnerMaskedResult() Result {
	response := *result

	response.Winner = Applicant{}
	return response
}
