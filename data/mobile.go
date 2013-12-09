package data

import (
	"bytes"
	"encoding/csv"
)

type Mobile struct {
	Token []byte
}

var Mobiles []*Mobile

func init() {
	records, err := csv.NewReader(bytes.NewBuffer(mobile_csv())).ReadAll()
	if err != nil {
		panic("parse error (mobile.csv): " + err.Error())
	}
	for _, record := range records {
		Mobiles = append(Mobiles, &Mobile{[]byte(record[0])})
	}
}
