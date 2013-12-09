package agency

import (
	"fmt"
	"testing"

	"github.com/benbjohnson/agency/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestScanBrowser(t *testing.T) {
	s := NewScanner()
	for i, browser := range fixtures.Browsers() {
		lineNum := fmt.Sprintf("Line #%d", i+1)
		ua, _ := s.Scan(browser.UserAgentString)
		isname := assert.Equal(t, ua.Browser.Name, browser.Name, lineNum)
		istype := assert.Equal(t, ua.Browser.Type, browser.Type, lineNum)
		if !isname || !istype {
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
