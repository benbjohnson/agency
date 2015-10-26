package agency

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"testing"
)

func TestScanBrowser(t *testing.T) {
	s := NewScanner()
	b, _ := ioutil.ReadFile("fixtures/browser.csv")
	records, _ := csv.NewReader(bytes.NewBuffer(b)).ReadAll()
	for i, record := range records {
		ua, _ := s.Scan(record[2])
		if ua.Browser.Type != record[0] {
			t.Fatalf("[Line #%d], got %q, wanted %q", i+1, ua.Browser.Type, record[0])
		}
		if ua.Browser.Name != record[1] {
			t.Fatalf("[Line #%d], got %q, wanted %q", i+1, ua.Browser.Name, record[0])
		}
	}
}

func TestScanDevice(t *testing.T) {
	s := NewScanner()
	b, _ := ioutil.ReadFile("fixtures/device.csv")
	records, _ := csv.NewReader(bytes.NewBuffer(b)).ReadAll()
	for i, record := range records {
		ua, _ := s.Scan(record[1])
		if ua.Device.Type != record[0] {
			t.Fatalf("[Line #%d], got %q, wanted %q", i+1, ua.Device.Type, record[0])
		}
	}
}

func TestScanOS(t *testing.T) {
	b, _ := ioutil.ReadFile("fixtures/os.csv")
	records, _ := csv.NewReader(bytes.NewBuffer(b)).ReadAll()
	s := NewScanner()
	for i, record := range records {
		ua, _ := s.Scan(record[2])
		if ua.OS.Name != record[0] {
			t.Fatalf("[Line #%d], got %q, wanted %q", i+1, ua.OS.Name, record[0])
		}
		if ua.OS.Version != record[1] {
			t.Fatalf("[Line #%d], got %q, wanted %q", i+1, ua.OS.Version, record[1])
		}
	}
}

func BenchmarkScan(b *testing.B) {
	s := NewScanner()
	for i := 0; i < b.N; i++ {
		s.Scan("Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1667.0 Safari/537.36")
	}
}
