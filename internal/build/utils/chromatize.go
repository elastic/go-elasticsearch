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
