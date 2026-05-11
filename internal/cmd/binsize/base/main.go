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

// Binary-size fixture for [elasticsearch.BaseClient]. Built by CI and
// compared against the PR's base branch to detect unintended growth.
// Must not transitively reach esapi or typedapi.
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	c, err := elasticsearch.NewBase()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	res, err := c.Perform(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	res.Body.Close()
}
