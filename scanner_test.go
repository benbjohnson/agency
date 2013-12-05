package agency

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var scantests = []struct {
        typ string
        category string
        browser string
        ua string
}{
        {"Windows", "Desktop", "Chrome", `Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1667.0 Safari/537.36`},
        {"Mac", "Desktop", "Safari", `Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_7; en-us) AppleWebKit/533.21.1 (KHTML, like Gecko) Version/5.0.5 Safari/533.21.1`},
}

func TestScan(t *testing.T) {
	for i, test := range scantests {
		name := fmt.Sprintf("TEST LINE %d", i+1)
		ua, _ := Scan(test.ua)
		assert.Equal(t, ua.Type, test.typ, name)
		assert.Equal(t, ua.Category, test.category, name)
		assert.Equal(t, ua.Browser, test.browser, name)
	}
}


func BenchmarkScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Scan("Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1667.0 Safari/537.36")
	}
}
