/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package generator

import (
	"bytes"
	"io"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/aryann/difflib"
)

const (
	testSpecDir = "testdata"
)

type dummyCloser struct {
	io.Reader
}

func (d *dummyCloser) Close() error {
	return nil
}

func newIndexMethod(t *testing.T) *method {
	templates, err := template.ParseFiles(filepath.Join("..", DefaultTemplatesDir, "method.tmpl"),
		filepath.Join("..", DefaultTemplatesDir, "option.tmpl"))
	if err != nil {
		t.Fatal(err)
	}
	m, err := newMethod(testSpecDir, "index.json", nil, templates, true)
	if err != nil {
		t.Fatal(err)
	}
	body := `<div class="chapter"><div class="titlepage"><div><div><h2 class="title"><a id="docs-index_"></a>
	Index API<a href="https://github.com/elastic/elasticsearch/edit/master/docs/reference/docs/index_.asciidoc"
	class="edit_me" title="Edit this page on GitHub" rel="nofollow">edit</a></h2></div></div></div>
	<p>The index API adds or updates a typed JSON document in a specific index,
making it searchable. The following example inserts the JSON document
into the "twitter" index, under a type called "tweet" with an id of 1:</p>`
	m.HTTPCache["http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html"] =
		&dummyCloser{bytes.NewBufferString(body)}
	return m
}

func splitAndTrim(s string) []string {
	lines := []string{}
	for _, l := range strings.Split(s, "\n") {
		lines = append(lines, strings.Trim(l, " \t"))
	}
	return lines
}

func diff(t *testing.T, expected, actual string) []difflib.DiffRecord {
	expectedLines := splitAndTrim(expected)
	actualLines := splitAndTrim(actual)
	d := difflib.Diff(expectedLines, actualLines)
	diffs := []difflib.DiffRecord{}
	for _, delta := range d {
		if delta.Delta != difflib.Common {
			diffs = append(diffs, delta)
		}
	}
	if len(diffs) > 0 {
		for _, delta := range d {
			t.Log(delta)
		}
	}
	return diffs
}
