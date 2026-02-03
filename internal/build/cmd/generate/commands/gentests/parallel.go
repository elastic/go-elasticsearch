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
	"path"
	"path/filepath"
	"sort"
	"strings"
)

// ParallelCleanup lists scoped resources to clear for parallel-safe suites.
type ParallelCleanup struct {
	Indices     []string
	DataStreams []string
}

var parallelUnsafeGlobs = []string{
	"autoscaling/*",
	"cat/*",
	"cluster/*",
	"connector/*",
	"dangling_indices/*",
	"enrich/*",
	"fleet/*",
	"graph/*",
	"health_report.yml",
	"ilm/*",
	"info_*.yml",
	"inference/*",
	"license/*",
	"logstash/*",
	"machine_learning/*",
	"migration/*",
	"nodes/*",
	"profiling/*",
	"query_rules/*",
	"rollup/*",
	"search_application/*",
	"searchable_snapshots/*",
	"security/*",
	"slm/*",
	"snapshot/*",
	"ssl.yml",
	"streams/*",
	"tasks*.yml",
	"terms_enum/*",
	"text_structure/*",
	"transform/*",
	"watcher/*",
	"xpack_info.yml",
}

var parallelUnsafePrefixes = map[string]struct{}{
	"cat":                  {},
	"cluster":              {},
	"connector":            {},
	"dangling_indices":     {},
	"enrich":               {},
	"fleet":                {},
	"graph":                {},
	"ilm":                  {},
	"ingest":               {},
	"inference":            {},
	"license":              {},
	"logstash":             {},
	"ml":                   {},
	"migration":            {},
	"nodes":                {},
	"profiling":            {},
	"query_rules":          {},
	"rollup":               {},
	"search_application":   {},
	"searchable_snapshots": {},
	"security":             {},
	"shutdown":             {},
	"slm":                  {},
	"snapshot":             {},
	"streams":              {},
	"synonyms":             {},
	"tasks":                {},
	"text_structure":       {},
	"transform":            {},
	"watcher":              {},
}

var parallelUnsafeMethods = map[string]struct{}{
	"health_report": {},
	"info":          {},
	"ping":          {},
	"search_shards": {},
	"xpack.info":    {},
}

func isParallelUnsafeByPath(rel string) bool {
	rel = path.Clean(strings.ReplaceAll(rel, string(filepath.Separator), "/"))
	for _, glob := range parallelUnsafeGlobs {
		if ok, _ := path.Match(glob, rel); ok {
			return true
		}
	}
	return false
}

func (ts *TestSuite) classifyParallelSafety() {
	if ts.Skip {
		return
	}

	if ts.Type == "xpack" {
		ts.ParallelSafe = false
		ts.ParallelReason = "xpack suite"
		return
	}

	if isParallelUnsafeByPath(ts.BaseFilename()) {
		ts.ParallelSafe = false
		ts.ParallelReason = "denylist"
		return
	}

	actions := ts.collectActions()
	if len(actions) == 0 {
		ts.ParallelSafe = false
		ts.ParallelReason = "no actions"
		return
	}

	cleanup := parallelCleanupSet{
		indices:     make(map[string]struct{}),
		dataStreams: make(map[string]struct{}),
	}

	for _, a := range actions {
		method := strings.TrimSpace(a.method)
		if method == "" {
			continue
		}

		if isUnsafeMethod(method) {
			ts.ParallelSafe = false
			ts.ParallelReason = "unsafe endpoint: " + method
			return
		}

		indexes := extractIndexValues(a.params, "index", "indices")
		dataStreams := extractIndexValues(a.params, "data_stream", "data_streams", "name")

		if len(indexes) == 0 && len(dataStreams) == 0 {
			ts.ParallelSafe = false
			ts.ParallelReason = "no scoped index/data stream in " + method
			return
		}

		if containsUnsafeResource(indexes) || containsUnsafeResource(dataStreams) {
			ts.ParallelSafe = false
			ts.ParallelReason = "wildcard/unknown resource in " + method
			return
		}

		cleanup.addAll(cleanup.indices, indexes)
		cleanup.addAll(cleanup.dataStreams, dataStreams)
	}

	if len(cleanup.indices) == 0 && len(cleanup.dataStreams) == 0 {
		ts.ParallelSafe = false
		ts.ParallelReason = "no cleanup resources"
		return
	}

	ts.ParallelSafe = true
	ts.ParallelCleanup = cleanup.toCleanup()
}

func isUnsafeMethod(method string) bool {
	if _, ok := parallelUnsafeMethods[method]; ok {
		return true
	}

	parts := strings.Split(method, ".")
	if len(parts) > 0 {
		if _, ok := parallelUnsafePrefixes[parts[0]]; ok {
			return true
		}
	}

	if strings.Contains(method, "template") || strings.Contains(method, "alias") {
		return true
	}

	return false
}

func extractIndexValues(params map[interface{}]interface{}, keys ...string) []string {
	if len(params) == 0 {
		return nil
	}

	var out []string
	for _, key := range keys {
		for kk, vv := range params {
			if ks, ok := kk.(string); ok && ks == key {
				out = append(out, normalizeResourceValues(vv)...)
			}
		}
	}
	return out
}

func normalizeResourceValues(v interface{}) []string {
	var out []string
	switch vv := v.(type) {
	case string:
		out = append(out, splitCSV(vv)...)
	case []interface{}:
		for _, item := range vv {
			if s, ok := item.(string); ok {
				out = append(out, splitCSV(s)...)
			}
		}
	}
	return out
}

func splitCSV(value string) []string {
	var out []string
	for _, part := range strings.Split(value, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func containsUnsafeResource(values []string) bool {
	for _, value := range values {
		if value == "" {
			return true
		}
		if strings.HasPrefix(value, "$") || strings.Contains(value, "{") {
			return true
		}
		if value == "_all" || strings.ContainsAny(value, "*?") {
			return true
		}
	}
	return false
}

type parallelCleanupSet struct {
	indices     map[string]struct{}
	dataStreams map[string]struct{}
}

func (s *parallelCleanupSet) addAll(target map[string]struct{}, values []string) {
	for _, value := range values {
		if value == "" {
			continue
		}
		target[value] = struct{}{}
	}
}

func (s parallelCleanupSet) toCleanup() ParallelCleanup {
	cleanup := ParallelCleanup{
		Indices:     keysFromSet(s.indices),
		DataStreams: keysFromSet(s.dataStreams),
	}
	return cleanup
}

func keysFromSet(values map[string]struct{}) []string {
	if len(values) == 0 {
		return nil
	}
	out := make([]string, 0, len(values))
	for value := range values {
		out = append(out, value)
	}
	sort.Strings(out)
	return out
}
