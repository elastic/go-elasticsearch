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

// Package sourcemode
package sourcemode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/_types/IndexSettings.ts#L509-L513
type SourceMode struct {
	Name string
}

var (
	Disabled = SourceMode{"disabled"}

	Stored = SourceMode{"stored"}

	Synthetic = SourceMode{"synthetic"}
)

func (s SourceMode) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SourceMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "disabled":
		*s = Disabled
	case "stored":
		*s = Stored
	case "synthetic":
		*s = Synthetic
	default:
		*s = SourceMode{string(text)}
	}

	return nil
}

func (s SourceMode) String() string {
	return s.Name
}
