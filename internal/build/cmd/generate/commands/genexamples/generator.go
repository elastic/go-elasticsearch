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

package genexamples

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// SrcGenerator represents the generator for Go source files.
//
type SrcGenerator struct {
	b       bytes.Buffer
	Example Example
}

// DocGenerator represents the generator for Go source files.
//
type DocGenerator struct {
	b        bytes.Buffer
	Source   io.Reader
	Filename string
	Digest   string
}

func (g SrcGenerator) Filename() string {
	return fmt.Sprintf(
		"%s_%s_test.go",
		strings.ReplaceAll(strings.TrimSuffix(g.Example.SourceLocation.File, ".asciidoc"), "/", "-"),
		g.Example.Digest)
}

func (g SrcGenerator) Output() (io.Reader, error) {
	var out bytes.Buffer

	out.WriteString(`// Licensed to Elasticsearch B.V. under one or more contributor
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
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

`)

	fmt.Fprintf(&out, "// <%s>\n//\n", g.Example.GithubURL())
	out.WriteString("// " + strings.Repeat("-", 80) + "\n")
	for _, l := range strings.Split(g.Example.Source, "\n") {
		out.WriteString("// " + l + "\n")
	}
	out.WriteString("// " + strings.Repeat("-", 80) + "\n\n")

	fmt.Fprintf(&out, `func Test_%s_%s(t *testing.T) {`+"\n", g.Example.Chapter(), g.Example.Digest)
	out.WriteString("\t")
	if !g.Example.IsTranslated() {
		out.WriteString("// ")
	}
	out.WriteString(`es, _ := elasticsearch.NewDefaultClient()` + "\n")

	if !g.Example.IsTranslated() {
		out.WriteString("\n\tt.Error(`Missing implementation for: ")

		out.WriteString(strings.Split(g.Example.Source, "\n")[0])
		out.WriteString("`)\n")
	} else {
		src, err := g.Example.Translated()
		if err != nil {
			return nil, err
		} else {
			out.WriteString(src)
		}
	}

	out.WriteString("}\n")

	return &out, nil
}

func (g DocGenerator) Output() (io.Reader, error) {
	var out bytes.Buffer

	fmt.Fprintf(&out, "// Generated from %s\n//\n", g.Filename)

	out.WriteString("[source, go]\n----\n")

	_, err := out.ReadFrom(g.Source)
	if err != nil {
		return nil, err
	}

	out.WriteString("----\n")

	return &out, nil
}
