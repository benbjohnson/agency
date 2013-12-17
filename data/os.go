package data

import (
	"bytes"
	"encoding/csv"
	"strconv"
)

type OS struct {
	Rank int
	Name string
	Version string
	Token []byte
}

var OSes []*OS

func init() {
	records, err := csv.NewReader(bytes.NewBuffer(os_csv())).ReadAll()
	if err != nil {
		panic("parse error (os.csv): " + err.Error())
	}
	for _, record := range records {
		rank, _ := strconv.Atoi(record[0])
		OSes = append(OSes, &OS{rank, record[1], record[2], []byte(record[3])})
	}
}
