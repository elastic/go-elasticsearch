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
)

// DenseVectorBytes is a byte-backed dense vector representation. It
// serializes to a base64 string of the raw bytes (big-endian is not
// relevant here since the payload is already bytes). This is useful when
// the client already has the vector encoded as bytes and wants to avoid
// re-encoding into float32 on the client-side.
type DenseVectorBytes []byte

// MarshalJSON implements json.Marshaler for byte-backed vectors.
func (d DenseVectorBytes) MarshalJSON() ([]byte, error) {
	if d == nil {
		return []byte("null"), nil
	}
	if len(d) == 0 {
		return json.Marshal("")
	}
	enc := base64.StdEncoding.EncodeToString([]byte(d))
	return json.Marshal(enc)
}
