package agency

// UserAgent represents the results from a call to Scan().
type UserAgent struct {
	Browser struct {
		Type string
		Name string
	}
	Device struct {
		Type string
	}
	OS struct {
		Name string
		Version string
	}
}
