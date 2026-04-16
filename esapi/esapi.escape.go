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

package esapi

import (
	"net/url"
	"strings"
)

// escapePathPart percent-escapes a single URL path segment while preserving
// commas. Commas act as separators inside comma-joined list path params
// (e.g. "index1,index2"), so each element is escaped independently and the
// commas are re-joined verbatim.
func escapePathPart(s string) string {
	if !strings.Contains(s, ",") {
		return url.PathEscape(s)
	}
	parts := strings.Split(s, ",")
	for i, p := range parts {
		parts[i] = url.PathEscape(p)
	}
	return strings.Join(parts, ",")
}
