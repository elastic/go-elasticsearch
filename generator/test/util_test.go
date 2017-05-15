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

package test

import (
	"bytes"
	"io"
	"testing"
	"text/template"

	"github.com/elastic/go-elasticsearch/generator/api"
)

type dummyCloser struct {
	io.Reader
}

func (d *dummyCloser) Close() error {
	return nil
}

// newIndexMethod creates an index method for testing.
// TODO: this should be shared with the generator/api package.
func newIndexMethod(t *testing.T) *api.Method {
	templates, err := template.ParseFiles("../api/templates/method.tmpl", "../api/templates/option.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	m, err := api.NewMethod("../api/testdata", "index.json", nil, templates, true)
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
