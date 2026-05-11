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

// Binary-size fixture for [elasticsearch.Client]. Built by CI and
// compared against the PR's base branch to detect unintended growth.
// Must not import typedapi: a regression like elastic/go-elasticsearch#1472
// (a method on *Client transitively pulling in the typedapi tree) shows
// up as a delta here.
//
// The reflect.ValueOf(c).NumMethod() loop below is deliberate. Go's
// linker performs per-method dead-code elimination on concrete types
// when it can prove no reachable code path uses a given method. A
// minimal program that calls only c.Info() lets the linker eliminate
// every other *Client method (including any that accidentally reach
// into typedapi), which hides exactly the regression class this
// fixture exists to detect. Real-world apps that hit #1472 use
// reflection-based libraries (DI containers, observability SDKs,
// generic serialization) which keep the full *Client method set live;
// the reflect walk simulates that and forces the linker to retain
// every method body so a transitive heavy-package edge is visible.
package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	c, err := elasticsearch.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	v := reflect.ValueOf(c)
	for i := 0; i < v.NumMethod(); i++ {
		_ = v.Type().Method(i).Name
	}
	res, err := c.Info()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	res.Body.Close()
}
