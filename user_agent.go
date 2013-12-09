package agency

// UserAgent represents the results from a call to Scan().
type UserAgent struct {
	Browser struct {
		Type string
		Name string
	}
}
