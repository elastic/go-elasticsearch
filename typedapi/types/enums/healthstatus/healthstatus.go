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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package healthstatus
package healthstatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/common.ts#L223-L245
type HealthStatus struct {
	Name string
}

var (
	Green = HealthStatus{"green"}

	Yellow = HealthStatus{"yellow"}

	Red = HealthStatus{"red"}

	Unknown = HealthStatus{"unknown"}

	Unavailable = HealthStatus{"unavailable"}
)

func (h HealthStatus) MarshalText() (text []byte, err error) {
	return []byte(h.String()), nil
}

func (h *HealthStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "green":
		*h = Green
	case "yellow":
		*h = Yellow
	case "red":
		*h = Red
	case "unknown":
		*h = Unknown
	case "unavailable":
		*h = Unavailable
	default:
		*h = HealthStatus{string(text)}
	}

	return nil
}

func (h HealthStatus) String() string {
	return h.Name
}
