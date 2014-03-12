package agency

import (
	"bytes"
	"encoding/csv"
	"strconv"
)

const MaxRank = 5

type Browser struct {
	Rank  int
	Type  string
	Name  string
	Token []byte
}

var Browsers []*Browser

func init() {
	data, _ := browser_csv()
	records, err := csv.NewReader(bytes.NewBuffer(data)).ReadAll()
	if err != nil {
		panic("parse error (browser.csv): " + err.Error())
	}
	for _, record := range records {
		rank, _ := strconv.Atoi(record[0])
		Browsers = append(Browsers, &Browser{rank, record[1], record[2], []byte(record[3])})
	}
}

type Device struct {
	Rank  int
	Type  string
	Token []byte
}

var Devices []*Device

func init() {
	data, _ := device_csv()
	records, err := csv.NewReader(bytes.NewBuffer(data)).ReadAll()
	if err != nil {
		panic("parse error (device.csv): " + err.Error())
	}
	for _, record := range records {
		rank, _ := strconv.Atoi(record[0])
		Devices = append(Devices, &Device{rank, record[1], []byte(record[2])})
	}
}

type Mobile struct {
	Token []byte
}

var Mobiles []*Mobile

func init() {
	data, _ := mobile_csv()
	records, err := csv.NewReader(bytes.NewBuffer(data)).ReadAll()
	if err != nil {
		panic("parse error (mobile.csv): " + err.Error())
	}
	for _, record := range records {
		Mobiles = append(Mobiles, &Mobile{[]byte(record[0])})
	}
}

type OS struct {
	Rank    int
	Name    string
	Version string
	Token   []byte
}

var OSes []*OS

func init() {
	data, _ := os_csv()
	records, err := csv.NewReader(bytes.NewBuffer(data)).ReadAll()
	if err != nil {
		panic("parse error (os.csv): " + err.Error())
	}
	for _, record := range records {
		rank, _ := strconv.Atoi(record[0])
		OSes = append(OSes, &OS{rank, record[1], record[2], []byte(record[3])})
	}
}
