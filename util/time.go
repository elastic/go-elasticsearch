package util

import "time"

// ParseDuration wraps time.ParseDuration returning a single value.
func ParseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}
