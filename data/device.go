package data

import (
	"bytes"
	"encoding/csv"
	"strconv"
)

type Device struct {
	Rank int
	Type string
	Token []byte
}

var Devices []*Device

func init() {
	records, err := csv.NewReader(bytes.NewBuffer(device_csv())).ReadAll()
	if err != nil {
		panic("parse error (device.csv): " + err.Error())
	}
	for _, record := range records {
		rank, _ := strconv.Atoi(record[0])
		Devices = append(Devices, &Device{rank, record[1], []byte(record[2])})
	}
}
