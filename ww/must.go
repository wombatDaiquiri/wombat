package ww

// Must is a wrapper for nicer syntax when an initialization is required on system startup.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
