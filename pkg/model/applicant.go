package model

import (
	"encoding/csv"
	"io"
	"log"
	"lottery_back/pkg/server"
	"lottery_back/pkg/util"
	"os"
	"strings"
)

type Applicant struct {
	name         string
	nameFurigana string
	class        string
}

var applicants []Applicant

func loadApplicants(server server.Server) {
	applicants = []Applicant{}

	f, err := os.Open(server.Config.ResourcePath.Applicant)

	util.CheckFatalError(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない

	log.Printf("info: start applicants loading")

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
		applicant := Applicant{name: record[4], nameFurigana: record[7], class: class}
		applicants = append(applicants, applicant)

		i++
	}

	log.Printf("info: finished applicants loading")
}
