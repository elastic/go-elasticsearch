// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

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
		s = `"Infinity"`
	case math.IsInf(float64(f), -1):
		s = `"-Infinity"`
	case math.IsNaN(float64(f)):
		s = `"NaN"`
	default:
		s = strconv.FormatFloat(float64(f), 'f', -1, 64)
	}
	return []byte(s), nil
}

// UnmarshalJSON implements Unmarshaler interface.
func (f *Float64) UnmarshalJSON(data []byte) error {
	switch {
	case bytes.Equal(data, []byte(`"NaN"`)):
		*f = Float64(math.NaN())
	case bytes.Equal(data, []byte(`"Infinity"`)):
		*f = Float64(math.Inf(1))
	case bytes.Equal(data, []byte(`"-Infinity"`)):
		*f = Float64(math.Inf(-1))
	case bytes.Equal(data, []byte(`null`)):
		return nil
	default:
		n, err := strconv.ParseFloat(string(data), 64)
		if err != nil {
			return err
		}
		*f = Float64(n)
	}
	return nil
}
