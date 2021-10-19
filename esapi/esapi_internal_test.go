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

//go:build !integration
// +build !integration

package esapi

import (
	"testing"
	"time"
)

func TestAPIHelpers(t *testing.T) {
	t.Run("BoolPtr", func(t *testing.T) {
		v := BoolPtr(false)
		if v == nil || *v != false {
			t.Errorf("Expected false, got: %v", v)
		}

		v = BoolPtr(true)
		if v == nil || *v != true {
			t.Errorf("Expected true, got: %v", v)
		}
	})

	t.Run("IntPtr", func(t *testing.T) {
		v := IntPtr(0)
		if v == nil || *v != 0 {
			t.Errorf("Expected 0, got: %v", v)
		}
	})

	t.Run("FormatDuration", func(t *testing.T) {
		var tt = []struct {
			duration time.Duration
			expected string
		}{
			{1 * time.Nanosecond, "1nanos"},
			{100 * time.Nanosecond, "100nanos"},
			{1 * time.Microsecond, "1000nanos"},
			{1 * time.Millisecond, "1ms"},
			{100 * time.Millisecond, "100ms"},
			{1 * time.Minute, "60000ms"},
			{10 * time.Minute, "600000ms"},
			{1 * time.Hour, "3600000ms"},
			{10 * time.Hour, "36000000ms"},
		}

		for _, tc := range tt {
			actual := formatDuration(tc.duration)
			if actual != tc.expected {
				t.Errorf("Unexpected output: got=%s, want=%s", actual, tc.expected)
			}
		}
	})
}
