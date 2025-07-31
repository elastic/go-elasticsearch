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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

// Package eventtype
package eventtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/86f41834c7bb975159a38a73be8a9d930010d673/specification/search_application/_types/AnalyticsEvent.ts#L22-L26
type EventType struct {
	Name string
}

var (
	PageView = EventType{"page_view"}

	Search = EventType{"search"}

	SearchClick = EventType{"search_click"}
)

func (e EventType) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *EventType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "page_view":
		*e = PageView
	case "search":
		*e = Search
	case "search_click":
		*e = SearchClick
	default:
		*e = EventType{string(text)}
	}

	return nil
}

func (e EventType) String() string {
	return e.Name
}
