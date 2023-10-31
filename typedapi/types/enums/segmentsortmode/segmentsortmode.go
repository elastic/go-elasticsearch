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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Package segmentsortmode
package segmentsortmode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/indices/_types/IndexSegmentSort.ts#L36-L41
type SegmentSortMode struct {
	Name string
}

var (
	Min = SegmentSortMode{"min"}

	Max = SegmentSortMode{"max"}
)

func (s SegmentSortMode) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SegmentSortMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "min":
		*s = Min
	case "max":
		*s = Max
	default:
		*s = SegmentSortMode{string(text)}
	}

	return nil
}

func (s SegmentSortMode) String() string {
	return s.Name
}
