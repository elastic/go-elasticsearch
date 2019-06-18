// +build !integration

package esutil

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

type errReader struct{}

func (errReader) Len() int                           { return 1 }
func (errReader) Read(p []byte) (int, error)         { return 1, errors.New("MOCK ERROR") }
func (errReader) Write(p []byte) (int, error)        { return 0, errors.New("MOCK ERROR") }
func (errReader) WriteTo(w io.Writer) (int64, error) { return 0, errors.New("MOCK ERROR") }

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

func TestJSONReader(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		out, _ := ioutil.ReadAll(NewJSONReader(map[string]string{"foo": "bar"}))
		if string(out) != `{"foo":"bar"}`+"\n" {
			t.Fatalf("Unexpected output: %s", out)
		}
	})

	t.Run("Custom", func(t *testing.T) {
		out, _ := ioutil.ReadAll(NewJSONReader(Foo{Bar: "baz"}))
		if string(out) != `{"bar":"BAZ"}`+"\n" {
			t.Fatalf("Unexpected output: %s", out)
		}
	})

	t.Run("WriteTo", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})
		r := JSONReader{val: map[string]string{"foo": "bar"}}
		r.WriteTo(b)
		if b.String() != `{"foo":"bar"}`+"\n" {
			t.Fatalf("Unexpected output: %s", b.String())
		}
	})

	t.Run("Len", func(t *testing.T) {
		r := JSONReader{val: map[string]string{"foo": "bar"}}
		if n := r.Len(); n != 14 {
			t.Fatalf("Unexpected length: %d", n)
		}
	})

	t.Run("Read error", func(t *testing.T) {
		b := []byte{}
		r := JSONReader{val: map[string]string{"foo": "bar"}, buf: errReader{}}
		size, err := r.Read(b)
		if err == nil {
			t.Fatalf("Expected error, got: %#v", err)
		}
		if n := r.Len(); n != size {
			t.Fatalf("Expected length to be 1, got %d", n)
		}
	})
}
