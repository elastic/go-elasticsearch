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
	"strings"
	"testing"

	"github.com/aryann/difflib"
)

func splitAndTrim(s string) []string {
	lines := []string{}
	for _, l := range strings.Split(s, "\n") {
		lines = append(lines, strings.Trim(l, " \t"))
	}
	return lines
}

func diff(t *testing.T, expected, actual string) []difflib.DiffRecord {
	expectedLines := splitAndTrim(expected)
	actualLines := splitAndTrim(actual)
	d := difflib.Diff(expectedLines, actualLines)
	diffs := []difflib.DiffRecord{}
	for _, delta := range d {
		if delta.Delta != difflib.Common {
			diffs = append(diffs, delta)
		}
	}
	if len(diffs) > 0 {
		for _, delta := range d {
			t.Log(delta)
		}
	}
	return diffs
}
