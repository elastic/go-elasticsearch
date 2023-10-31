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

// Package memorystatus
package memorystatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Model.ts#L88-L92
type MemoryStatus struct {
	Name string
}

var (
	Ok = MemoryStatus{"ok"}

	Softlimit = MemoryStatus{"soft_limit"}

	Hardlimit = MemoryStatus{"hard_limit"}
)

func (m MemoryStatus) MarshalText() (text []byte, err error) {
	return []byte(m.String()), nil
}

func (m *MemoryStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "ok":
		*m = Ok
	case "soft_limit":
		*m = Softlimit
	case "hard_limit":
		*m = Hardlimit
	default:
		*m = MemoryStatus{string(text)}
	}

	return nil
}

func (m MemoryStatus) String() string {
	return m.Name
}
