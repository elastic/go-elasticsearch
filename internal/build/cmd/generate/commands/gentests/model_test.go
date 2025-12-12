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

import "testing"

func TestNewTestSuite_RequiresStackFalse_SkipsSuite(t *testing.T) {
	payloads := []TestPayload{
		{
			Filepath: "tests/some/stack_false.yml",
			Payload: map[interface{}]interface{}{
				"requires": map[interface{}]interface{}{
					"serverless": true,
					"stack":      false,
				},
			},
		},
		{
			Filepath: "tests/some/stack_false.yml",
			Payload: map[interface{}]interface{}{
				"some test": []interface{}{
					map[interface{}]interface{}{
						"do": map[interface{}]interface{}{
							"cluster.health": map[interface{}]interface{}{},
						},
					},
				},
			},
		},
	}

	ts := NewTestSuite("tests/some/stack_false.yml", payloads)
	if !ts.Skip {
		t.Fatalf("expected suite to be skipped when requires.stack=false")
	}
	if ts.SkipInfo == "" {
		t.Fatalf("expected SkipInfo to be set")
	}
	if len(ts.Tests) != 0 {
		t.Fatalf("expected no tests to be generated when suite is skipped, got: %d", len(ts.Tests))
	}
}
