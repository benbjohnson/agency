package agency

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanBrowser(t *testing.T) {
	s := NewScanner()
	b, _ := ioutil.ReadFile("fixtures/browser.csv")
	records, _ := csv.NewReader(bytes.NewBuffer(b)).ReadAll()
	for i, record := range records {
		ua, _ := s.Scan(record[2])
		istype := assert.Equal(t, record[0], ua.Browser.Type, fmt.Sprintf("Line #%d", i+1))
		isname := assert.Equal(t, record[1], ua.Browser.Name, fmt.Sprintf("Line #%d", i+1))
		if !isname || !istype {
			break
		}
	}
}

func TestScanDevice(t *testing.T) {
	s := NewScanner()
	b, _ := ioutil.ReadFile("fixtures/device.csv")
	records, _ := csv.NewReader(bytes.NewBuffer(b)).ReadAll()
	for i, record := range records {
		ua, _ := s.Scan(record[1])
		if !assert.Equal(t, record[0], ua.Device.Type, fmt.Sprintf("Line #%d", i+1)) {
			break
		}
	}
}

func TestScanOS(t *testing.T) {
	s := NewScanner()
	b, _ := ioutil.ReadFile("fixtures/os.csv")
	records, _ := csv.NewReader(bytes.NewBuffer(b)).ReadAll()
	for i, record := range records {
		ua, _ := s.Scan(record[2])
		isname := assert.Equal(t, record[0], ua.OS.Name, fmt.Sprintf("Line #%d", i+1))
		isversion := assert.Equal(t, record[1], ua.OS.Version, fmt.Sprintf("Line #%d", i+1))
		if !isname || !isversion {
			break
		}
	}
}



func BenchmarkScan(b *testing.B) {
	s := NewScanner()
	for i := 0; i < b.N; i++ {
		s.Scan("Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1667.0 Safari/537.36")
	}
}
