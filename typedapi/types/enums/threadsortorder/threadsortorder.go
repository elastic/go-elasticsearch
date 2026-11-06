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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

// Package threadsortorder
package threadsortorder

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_types/common.ts#L333-L338
type ThreadSortOrder struct {
	Name string
}

var (

	// Total Sort by the time the thread was runnable, then by CPU time.
	Total = ThreadSortOrder{"total"}

	// Cpu Sort by CPU time only.
	Cpu = ThreadSortOrder{"cpu"}
)

func (t ThreadSortOrder) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *ThreadSortOrder) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "total":
		*t = Total
	case "cpu":
		*t = Cpu
	default:
		*t = ThreadSortOrder{string(text)}
	}

	return nil
}

func (t ThreadSortOrder) String() string {
	return t.Name
}
