package types

import (
	"bytes"
	"math"
	"strconv"
)

// Float64 custom type for Inf & NaN handling.
type Float64 float64

// MarshalJSON implements Marshaler interface.
func (f Float64) MarshalJSON() ([]byte, error) {
	var s string
	switch {
	case math.IsInf(float64(f), 1):
		s = "Infinity"
	case math.IsInf(float64(f), -1):
		s = "-Infinity"
	case math.IsNaN(float64(f)):
		s = "NaN"
	default:
		s = strconv.FormatFloat(float64(f), 'f', -1, 64)
	}
	return []byte(`"` + s + `"`), nil
}

// UnmarshalJSON implements Unmarshaler interface.
func (f *Float64) UnmarshalJSON(data []byte) error {
	switch {
	case bytes.Equal(data, []byte(`null`)):
		f = nil
	default:
		n, err := strconv.ParseFloat(string(data), 64)
		if err != nil {
			return err
		}
		*f = Float64(n)
	}
	return nil
}
