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

//go:build !integration
// +build !integration

package esutil

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

type Foo struct {
	Bar string
}

func (f Foo) EncodeJSON(w io.Writer) error {
	_, err := w.Write([]byte(`{"bar":"` + strings.ToUpper(f.Bar) + `"}` + "\n"))
	if err != nil {
		return err
	}
	return nil
}

// errEncoder is a JSONEncoder that always returns an error, used to test
// encode-error paths.
type errEncoder struct{}

func (errEncoder) EncodeJSON(_ io.Writer) error {
	return errors.New("MOCK ERROR")
}

func TestJSONReader(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		out, err := io.ReadAll(NewJSONReader(map[string]string{"foo": "bar"}))
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if string(out) != `{"foo":"bar"}`+"\n" {
			t.Fatalf("Unexpected output: %s", out)
		}
	})

	t.Run("Custom", func(t *testing.T) {
		out, err := io.ReadAll(NewJSONReader(Foo{Bar: "baz"}))
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if string(out) != `{"bar":"BAZ"}`+"\n" {
			t.Fatalf("Unexpected output: %s", out)
		}
	})

	t.Run("WriteTo", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})
		r := JSONReader{val: map[string]string{"foo": "bar"}}
		_, _ = r.WriteTo(b)
		if b.String() != `{"foo":"bar"}`+"\n" {
			t.Fatalf("Unexpected output: %s", b.String())
		}
	})

	t.Run("Read error", func(t *testing.T) {
		r := NewJSONReader(errEncoder{})
		_, err := r.Read(make([]byte, 10))
		if err == nil {
			t.Fatalf("Expected error, got: %#v", err)
		}
	})

	t.Run("Seek", func(t *testing.T) {
		r, ok := NewJSONReader(map[string]string{"foo": "bar"}).(io.ReadSeeker)
		if !ok {
			t.Fatal("NewJSONReader did not return an io.ReadSeeker")
		}

		// Read all content.
		first, err := io.ReadAll(r)
		if err != nil {
			t.Fatalf("Unexpected error reading: %s", err)
		}

		// Seek back to start.
		if _, seekErr := r.Seek(0, io.SeekStart); seekErr != nil {
			t.Fatalf("Unexpected error seeking: %s", seekErr)
		}

		// Read again – should get identical content.
		second, err := io.ReadAll(r)
		if err != nil {
			t.Fatalf("Unexpected error reading after seek: %s", err)
		}

		if string(first) != string(second) {
			t.Fatalf("Content after seek differs: %q vs %q", first, second)
		}
	})
}
