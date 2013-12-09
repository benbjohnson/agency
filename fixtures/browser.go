package fixtures

import (
	"bytes"
	"encoding/csv"
)

type Browser struct {
	Type string
	Name string
	UserAgentString string
}

// Browsers returns a list of browser data.
func Browsers() []*Browser {
	records, err := csv.NewReader(bytes.NewBuffer(browser_csv())).ReadAll()
	if err != nil {
		panic("fixture parse error (browser.csv): " + err.Error())
	}

	items := make([]*Browser, 1)
	for _, record := range records {
		items = append(items, &Browser{record[0], record[1], record[2]})
	}
	return items
}
