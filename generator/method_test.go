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
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	m, err := newMethod(filepath.Join("..", DefaultSpecDir, "index.json"))
	if err != nil {
		t.Fatal(err)
	}
	var writer bytes.Buffer
	err = m.generate(filepath.Join("..", templatesDir), &writer)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `// generated by github.com/elastic/elasticsearch-go/cmd/generator; DO NOT EDIT

package client

// Index is documented at http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html
func (c *Client) Index(documentType string, index string, opt ...Option) (*http.Response, error){
	_ = []Option{

//WithID,


//WithOpType,
//WithParent,
//WithPipeline,
//WithRefresh,
//WithRouting,
//WithTTL,
//WithTimeout,
//WithTimestamp,
//WithVersion,
//WithVersionType,
//WithWaitForActiveShards,

	}
	req := &http.Request{
		Method: "POST",
	}
	return c.client.Do(req)
}
`
	if writer.String() != expectedCode {
		t.Fatalf("Expected the generation of:\n\n%q\n\nbut got:\n\n%q", expectedCode, writer.String())
	}
}
