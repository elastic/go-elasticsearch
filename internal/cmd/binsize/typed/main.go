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

// Binary-size fixture for [elasticsearch.TypedClient]. Built by CI and
// compared against the PR's base branch to detect unintended growth.
//
// The reflect.ValueOf(c).NumMethod() loop below is deliberate. Go's
// linker performs per-method dead-code elimination on concrete types
// when it can prove no reachable code path uses a given method. A
// fixture that calls only c.Info().Do(ctx) lets the linker eliminate
// every other typed namespace method, which keeps the binary at ~9 MB
// with ~38 typedapi symbols out of ~31,000: a regression in any typed
// endpoint other than Info would be invisible. The reflect walk forces
// the linker to retain the full typed surface (~55 MB / ~31,155
// symbols) so that growth in any typed namespace is tracked here.
package main

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	c, err := elasticsearch.NewTyped()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	v := reflect.ValueOf(c)
	for i := 0; i < v.NumMethod(); i++ {
		_ = v.Type().Method(i).Name
	}
	if _, err := c.Info().Do(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
