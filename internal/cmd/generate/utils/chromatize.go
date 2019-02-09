package utils

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

// Chromatize returns a syntax highlighted Go code.
//
func Chromatize(r io.Reader) (io.Reader, error) {
	var b bytes.Buffer
	lexer := lexers.Get("go")

	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	it, err := lexer.Tokenise(nil, string(contents))
	if err != nil {
		return nil, err
	}

	style, _ := styles.Get("pygments").Builder().Build()
	formatter := formatters.Get("terminal256")
	err = formatter.Format(&b, style, it)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
