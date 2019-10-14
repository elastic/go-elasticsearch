// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package gensource_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/internal/cmd/generate/commands/gensource"
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
