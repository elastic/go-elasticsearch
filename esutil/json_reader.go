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

// NewJSONReader encodes v into JSON and returns it as an io.ReadSeeker.
//
// The returned value satisfies io.ReadSeeker, making it directly usable
// as BulkIndexerItem.Body without additional buffering.
func NewJSONReader(v interface{}) io.Reader {
	return &JSONReader{val: v}
}

// JSONEncoder defines the interface for custom JSON encoders.
type JSONEncoder interface {
	EncodeJSON(io.Writer) error
}

// JSONReader represents a reader which takes an interface value,
// encodes it into JSON, and wraps it in an io.ReadSeeker.
//
// It implements io.ReadSeeker so it can be used directly as
// BulkIndexerItem.Body.
type JSONReader struct {
	val interface{}
	buf *bytes.Reader
}

// ensureEncoded lazily encodes the value on the first Read or Seek call.
func (r *JSONReader) ensureEncoded() error {
	if r.buf != nil {
		return nil
	}
	var b bytes.Buffer
	if err := r.encode(&b); err != nil {
		return err
	}
	r.buf = bytes.NewReader(b.Bytes())
	return nil
}

// Read implements the io.Reader interface.
func (r *JSONReader) Read(p []byte) (int, error) {
	if err := r.ensureEncoded(); err != nil {
		return 0, err
	}
	return r.buf.Read(p)
}

// Seek implements the io.Seeker interface.
func (r *JSONReader) Seek(offset int64, whence int) (int64, error) {
	if err := r.ensureEncoded(); err != nil {
		return 0, err
	}
	return r.buf.Seek(offset, whence)
}

// WriteTo implements the io.WriterTo interface.
func (r *JSONReader) WriteTo(w io.Writer) (int64, error) {
	if r.buf != nil {
		// Already encoded; seek to start so WriteTo emits the full content.
		if _, err := r.buf.Seek(0, io.SeekStart); err != nil {
			return 0, err
		}
		return r.buf.WriteTo(w)
	}
	cw := countingWriter{Writer: w}
	err := r.encode(&cw)
	return int64(cw.n), err
}

func (r *JSONReader) encode(w io.Writer) error {
	if e, ok := r.val.(JSONEncoder); ok {
		return e.EncodeJSON(w)
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
