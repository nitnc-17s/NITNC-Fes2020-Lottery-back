package model

import (
	"encoding/csv"
	"io"
	"log"
	"lottery_back/pkg/server"
	"lottery_back/pkg/util"
	"os"
	"strconv"
)

type Prize struct {
	id   int
	name string
}

var prizes []Prize

func loadPrizes(server server.Server) {
	prizes = []Prize{}

	f, err := os.Open(server.Config.ResourcePath.Prize)

	util.CheckFatalError(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない

	log.Printf("info: start prizes loading")

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

		id, err := strconv.Atoi(record[0])
		util.CheckFatalError(err)

		prize := Prize{id: id, name: record[2]}
		prizes = append(prizes, prize)

		i++
	}

	log.Printf("info: finished prizes loading")
}
