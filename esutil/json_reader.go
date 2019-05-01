package esutil

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSONReader is an utility function which encodes v into JSON and returns it as a reader.
//
func JSONReader(v interface{}) io.Reader {
	return &jsonReader{val: v, buf: nil}
}

// JSONEncoder defines the interface for custom JSON encoders.
//
type JSONEncoder interface {
	EncodeJSON(io.Writer) error
}

type jsonReader struct {
	val interface{}
	buf interface {
		io.ReadWriter
		io.WriterTo
	}
}

func (r *jsonReader) Read(p []byte) (int, error) {
	if r.buf == nil {
		r.buf = new(bytes.Buffer)
		if err := r.encode(r.buf); err != nil {
			return 0, err
		}
	}

	return r.buf.Read(p)
}

func (r *jsonReader) WriteTo(w io.Writer) (int64, error) {
	cw := countingWriter{Writer: w}
	err := r.encode(&cw)
	return int64(cw.n), err
}

func (r *jsonReader) encode(w io.Writer) error {
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
