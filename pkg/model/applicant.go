package model

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"lottery_back/pkg/server"
	"lottery_back/pkg/util"
	"os"
	"strings"
)

type Applicant struct {
	Name         string `json:"name"`
	NameFurigana string `json:"name_furigana"`
	Class        string `json:"class"`
}

var applicants []Applicant

func loadApplicants(server *server.Server) {
	applicants = []Applicant{}

	f, err := os.Open(server.Config.ResourcePath.Applicant)

	util.CheckFatalError(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない

	log.Printf("debug: start applicants loading")

	i := 0
	for {
		// 見出し行スキップ
		if i == 0 {
			continue
		}

		record, err := reader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			util.CheckFatalError(err)
		}

		class := strings.Join([]string{strings.Replace(record[5], "年", "", -1), record[6]}, "")
		applicant := Applicant{Name: record[4], NameFurigana: record[7], Class: class}
		applicants = append(applicants, applicant)

		i++
	}

	log.Printf("debug: finished applicants loading")
}

func getApplicant(id int) (Applicant, error) {
	if id >= len(applicants) || id < 0 {
		return Applicant{}, errors.New("invalid id")
	}

	res := applicants[id]

	applicants = append(applicants[:id], applicants[id+1:]...)

	return res, nil
}
