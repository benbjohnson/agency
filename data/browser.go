package data

import (
	"bytes"
	"encoding/csv"
	"strconv"
)

type Browser struct {
	Rank int
	Type string
	Name string
	Token []byte
}

var Browsers []*Browser

func init() {
	records, err := csv.NewReader(bytes.NewBuffer(browser_csv())).ReadAll()
	if err != nil {
		panic("parse error (browser.csv): " + err.Error())
	}
	for _, record := range records {
		rank, _ := strconv.Atoi(record[0])
		Browsers = append(Browsers, &Browser{rank, record[1], record[2], []byte(record[3])})
	}
}
