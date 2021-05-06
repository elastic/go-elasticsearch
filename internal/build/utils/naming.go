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

package utils

import "strings"

var (
	// Options: boolean, enum, list, number, string, time
	typesToGo = map[string]string{
		"boolean": "*bool",
		"enum":    "string", // TODO: Custom "enum" type
		"list":    "[]string",
		"number":  "*int",
		"int":     "*int",
		"long":    "*int",
		"string":  "string",
		"time":    "time.Duration",
	}
)

// APIToGo returns the Go version of API call, eg. cluster.health => ClusterHealth
//
func APIToGo(s string) string {
	ep := strings.Split(s, ".")
	ns := make([]string, len(ep))
	for _, v := range ep {
		m := strings.Split(v, "_")
		mn := make([]string, len(m))
		for _, vv := range m {
			mn = append(mn, NameToGo(vv))
		}
		ns = append(ns, strings.Join(mn, ""))
	}
	return strings.Join(ns, "")
}

// NameToGo returns a Go version of name, eg. node_id => NodeID.
//
func NameToGo(s string, api ...string) string {
	exceptions := map[string]string{
		"index": "Index",
		"id":    "DocumentID",
		"type":  "DocumentType",
	}

	acronyms := map[string]string{
		"ip":  "IP",
		"id":  "ID",
		"ttl": "TTL",

		"api":   "API",
		"ccr":   "CCR",
		"ilm":   "ILM",
		"ml":    "ML",
		"sql":   "SQL",
		"ssl":   "SSL",
		"xpack": "XPack",
	}

	specialMappingsForID := map[string]string{
		"DeleteScript":         "ScriptID",
		"GetScript":            "ScriptID",
		"PutScript":            "ScriptID",
		"IngestDeletePipeline": "PipelineID",
		"IngestGetPipeline":    "PipelineID",
		"IngestPutPipeline":    "PipelineID",
		"IngestSimulate":       "PipelineID",
		"RenderSearchTemplate": "TemplateID",

		"MLDeleteDataFrameAnalytics":   "ID",
		"MLGetDataFrameAnalytics":      "ID",
		"MLGetDataFrameAnalyticsStats": "ID",
		"MLPutDataFrameAnalytics":      "ID",
		"MLStartDataFrameAnalytics":    "ID",
		"MLStopDataFrameAnalytics":     "ID",

		"RollupDeleteJob":     "JobID",
		"RollupGetJobs":       "JobID",
		"RollupGetRollupCaps": "Index",
		"RollupPutJob":        "JobID",
		"RollupStartJob":      "JobID",
		"RollupStopJob":       "JobID",

		"SecurityGetAPIKey": "ID",

		"WatcherDeleteWatch":  "WatchID",
		"WatcherExecuteWatch": "WatchID",
		"WatcherGetWatch":     "WatchID",
		"WatcherPutWatch":     "WatchID",
	}

	if s == "id" && api != nil && len(api) > 0 && api[0] != "" && specialMappingsForID[api[0]] != "" {
		return specialMappingsForID[api[0]]
	}

	if value, ok := exceptions[s]; ok {
		return value
	}

	ep := strings.Split(s, "_")
	ns := make([]string, len(ep))
	for _, v := range ep {
		if value, ok := acronyms[v]; ok {
			v = value
		} else if v == "uuid" {
			v = "UUID"
		}
		ns = append(ns, strings.Title(v))
	}
	return strings.Join(ns, "")
}

// TypeToGo returns a Go version of type, eg. boolean => *bool.
//
func TypeToGo(s string, comment ...bool) string {
	// If the string contains a pipe character, it's a polymorphic parameter,
	// ie. it takes heterogeous values, such as "boolean" and "number"
	if strings.Contains(s, "|") {
		return "interface{}"
	}

	if v, ok := typesToGo[s]; ok {
		return v
	}

	return "interface{}"
}
