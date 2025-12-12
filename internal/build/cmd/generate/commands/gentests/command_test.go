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

package gentests

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCommand_processFile_SkipOnError_RecoversAndContinues(t *testing.T) {
	EsVersion = "9.1.0"

	tmp := t.TempDir()
	outDir := filepath.Join(tmp, "out")

	// Build a path that matches the expected tests layout (used by BaseFilename()).
	fpath := filepath.Join(tmp, "elasticsearch-clients-tests", "tests", "cluster", "bad.yml")
	if err := os.MkdirAll(filepath.Dir(fpath), 0o755); err != nil {
		t.Fatalf("mkdir: %s", err)
	}

	// This YAML uses an intentionally unknown parameter (not in api_registry) to trigger
	// the same generator panic path as real-world mismatches.
	yaml := `---
"skip-on-error unknown param":
  - do:
      cluster.health:
        definitely_not_a_real_param: true
  - is_true: cluster_name
`
	if err := os.WriteFile(fpath, []byte(yaml), 0o644); err != nil {
		t.Fatalf("write yaml: %s", err)
	}

	cmd := &Command{
		Input:       fpath,
		Output:      outDir,
		Gofmt:       false,
		SkipOnError: true,
	}

	if err := cmd.processFile(fpath); err != nil {
		t.Fatalf("expected no error (skip-on-error), got: %s", err)
	}

	// Since generation failed, the output directory should not exist (or be empty).
	if _, err := os.Stat(outDir); err == nil {
		t.Fatalf("expected output dir to not be created on skipped generation, but it exists: %s", outDir)
	}
}
