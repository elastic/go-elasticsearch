// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package esutil_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/esutil"
)

var _ = fmt.Print

type Foo struct {
	Bar string
}

func (f Foo) EncodeJSON(w io.Writer) error {
	var b bytes.Buffer
	b.WriteString(`{"bar":"`)
	b.WriteString(strings.ToUpper(f.Bar))
	b.WriteString(`"}`)
	b.WriteString("\n")
	_, err := b.WriteTo(w)
	if err != nil {
		return err
	}
	return nil
}

func BenchmarkJSONReader(b *testing.B) {
	b.ReportAllocs()

	b.Run("None", func(b *testing.B) {
		b.ResetTimer()

		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			json.NewEncoder(&buf).Encode(map[string]string{"foo": "bar"})
			if string(buf.String()) != `{"foo":"bar"}`+"\n" {
				b.Fatalf("Unexpected output: %q", buf.String())
			}
			buf.Reset()
		}
	})

	b.Run("Default", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			out, _ := ioutil.ReadAll(esutil.NewJSONReader(map[string]string{"foo": "bar"}))
			if string(out) != `{"foo":"bar"}`+"\n" {
				b.Fatalf("Unexpected output: %q", out)
			}
		}
	})

	b.Run("Default-Copy", func(b *testing.B) {
		b.ResetTimer()

		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			io.Copy(&buf, esutil.NewJSONReader(map[string]string{"foo": "bar"}))
			if buf.String() != `{"foo":"bar"}`+"\n" {
				b.Fatalf("Unexpected output: %q", buf.String())
			}
			buf.Reset()
		}
	})

	b.Run("Custom", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			out, _ := ioutil.ReadAll(esutil.NewJSONReader(Foo{Bar: "baz"}))
			if string(out) != `{"bar":"BAZ"}`+"\n" {
				b.Fatalf("Unexpected output: %q", out)
			}
		}
	})
}
