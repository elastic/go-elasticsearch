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
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestExecuteTemplate(t *testing.T) {
	api := "index"
	methodName := "Index"
	docURL := "http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-" + api + "_.html"
	spec := map[string]interface{}{
		api: map[string]interface{}{

			"documentation": docURL,
		},
	}
	apiDir, err := ioutil.TempDir("", "client")
	if err != nil {
		t.Fatal(err)
	}
	err = executeTemplate(spec, filepath.Join("..", templatesDir), apiDir)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := fmt.Sprintf(`package client

// %s
// See: %s
func (c *Client) %s() {
}
`, methodName, docURL, methodName)
	code, err := ioutil.ReadFile(filepath.Join(apiDir, api) + ".go")
	if err != nil {
		t.Fatal(err)
	}
	if string(code) != expectedCode {
		t.Fatalf("Expected the generation of:\n\n%q\n\nbut got:\n\n%q", expectedCode, code)
	}
}
