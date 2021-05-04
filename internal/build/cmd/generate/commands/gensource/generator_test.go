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

package gensource_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/internal/build/cmd/generate/commands/gensource"
)

func TestGenerator(t *testing.T) {
	t.Run("Output", func(t *testing.T) {
		f, err := os.Open("testdata/info.json")
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		endpoint, err := gensource.NewEndpoint(f)
		if err != nil {
			t.Fatalf("Error creating endpoint for %q: %s", f.Name(), err)
		}

		gen := gensource.Generator{Endpoint: endpoint}

		out, err := gen.OutputFormatted()
		if err != nil {
			t.Fatalf("Error generating output for %q: %s", f.Name(), err)
		}

		s, err := ioutil.ReadAll(out)
		if err != nil {
			t.Fatalf("Error reading output for %q: %s", f.Name(), err)
		}
		// t.Logf("\n%s\n", s)

		if !strings.Contains(string(s), "func newInfoFunc(t Transport) Info {") {
			t.Error("Incorrect output")
		}
	})
}
