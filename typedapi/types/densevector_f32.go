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
	"encoding/base64"
	"encoding/json"
	"math"
)

// DenseVectorF32 is a compact representation of a vector of float32 values
// that serializes to JSON as a big-endian base64 encoded string of raw
// IEEE-754 float32 bytes. Use a nil DenseVectorF32 to emit JSON null; an
// empty (non-nil) DenseVectorF32 serializes to the empty string (base64 of
// zero bytes).
//
// This is a static implementation included in the generated typed API
// standard library to speed up ingestion paths that can provide a
// pre-encoded vector as a []float32.
type DenseVectorF32 []float32

// MarshalJSON implements json.Marshaler for float32-backed vectors.
func (d DenseVectorF32) MarshalJSON() ([]byte, error) {
	if d == nil {
		return []byte("null"), nil
	}

	if len(d) == 0 {
		return json.Marshal("")
	}

	b := make([]byte, 4*len(d))
	j := 0
	for _, f := range d {
		bits := math.Float32bits(f)
		b[j+0] = byte(bits >> 24)
		b[j+1] = byte(bits >> 16)
		b[j+2] = byte(bits >> 8)
		b[j+3] = byte(bits)
		j += 4
	}

	enc := base64.StdEncoding.EncodeToString(b)
	return json.Marshal(enc)
}
