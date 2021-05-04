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

package esutil

import (
	"bytes"
	"encoding/json"
	"io"
)

// NewJSONReader encodes v into JSON and returns it as an io.Reader.
//
func NewJSONReader(v interface{}) io.Reader {
	return &JSONReader{val: v, buf: nil}
}

// JSONEncoder defines the interface for custom JSON encoders.
//
type JSONEncoder interface {
	EncodeJSON(io.Writer) error
}

// JSONReader represents a reader which takes an interface value,
// encodes it into JSON, and wraps it in an io.Reader.
//
type JSONReader struct {
	val interface{}
	buf interface {
		io.ReadWriter
		io.WriterTo
	}
}

// Read implements the io.Reader interface.
//
func (r *JSONReader) Read(p []byte) (int, error) {
	if r.buf == nil {
		r.buf = new(bytes.Buffer)
		if err := r.encode(r.buf); err != nil {
			return 0, err
		}
	}

	return r.buf.Read(p)
}

// WriteTo implements the io.WriterTo interface.
//
func (r *JSONReader) WriteTo(w io.Writer) (int64, error) {
	cw := countingWriter{Writer: w}
	err := r.encode(&cw)
	return int64(cw.n), err
}

func (r *JSONReader) encode(w io.Writer) error {
	var err error

	if e, ok := r.val.(JSONEncoder); ok {
		err = e.EncodeJSON(w)
		if err != nil {
			return err
		}
		return nil
	}

	return json.NewEncoder(w).Encode(r.val)
}

type countingWriter struct {
	io.Writer
	n int
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.Writer.Write(p)
	if n > 0 {
		cw.n += n
	}
	return n, err
}
