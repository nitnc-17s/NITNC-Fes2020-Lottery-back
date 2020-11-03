package model

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"lottery_back/pkg/config"
	"lottery_back/pkg/util"
	"os"
	"strconv"
)

type Prize struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var prizes []Prize

func loadPrizes(config config.Config) {
	prizes = []Prize{}

	f, err := os.Open(config.ResourcePath.Prize)
	util.CheckFatalError(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない

	log.Printf("debug: start prizes loading")

	i := 0
	for {
		record, err := reader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			util.CheckFatalError(err)
		}

		// 見出し行スキップ
		if i == 0 {
			i++
			continue
		}

		id, err := strconv.Atoi(record[0])
		util.CheckFatalError(err)

		prize := Prize{Id: id, Name: record[2]}
		prizes = append(prizes, prize)

		i++
	}

	log.Printf("debug: finished prizes loading")
}

func GetPrize(id int) (Prize, error) {
	for _, prize := range prizes {
		if prize.Id == id {
			return prize, nil
		}
	}

	return Prize{}, errors.New("invalid id")
}
