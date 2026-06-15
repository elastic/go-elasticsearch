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
	"testing"
)

func TestDenseVectorMarshal(t *testing.T) {
	v := DenseVectorF32{1.0, -2.5, 0}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		t.Fatalf("expected marshalled JSON to be string, got: %v", err)
	}

	raw, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		t.Fatalf("invalid base64: %v", err)
	}
	if len(raw) != 12 {
		t.Fatalf("unexpected raw length: %d", len(raw))
	}
}

func TestDenseVectorNullAndEmpty(t *testing.T) {
	var v DenseVectorF32 = nil
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal null error: %v", err)
	}
	if string(b) != "null" {
		t.Fatalf("expected null, got %s", string(b))
	}

	v = DenseVectorF32{}
	b, err = json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal empty error: %v", err)
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		t.Fatalf("expected marshalled empty to be string, got: %v", err)
	}
	if s != "" {
		t.Fatalf("expected empty base64 string, got %q", s)
	}
}

func TestDenseVectorKnownBase64(t *testing.T) {
	v := DenseVectorF32{1.0, 2.0}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		t.Fatalf("expected marshalled JSON to be string, got: %v", err)
	}

	want := "P4AAAEAAAAA="
	if s != want {
		t.Fatalf("unexpected base64 encoding: want %q, got %q", want, s)
	}
}
