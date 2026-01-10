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
	"encoding/json"
	"testing"
)

func TestDenseVectorBytesKnownBase64(t *testing.T) {
	// raw big-endian float32 bytes for 1.0 and 2.0: 0x3f800000, 0x40000000
	raw := []byte{0x3f, 0x80, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00}

	dbytes := DenseVectorBytes(raw)
	b, err := json.Marshal(dbytes)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	want := "\"P4AAAEAAAAA=\""
	if s := string(b); s != want {
		t.Fatalf("unexpected base64 encoding for bytes: want %q, got %q", want, s)
	}
}
