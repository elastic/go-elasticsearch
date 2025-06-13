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
// https://github.com/elastic/elasticsearch-specification/tree/3a94b6715915b1e9311724a2614c643368eece90

// Package coheretruncatetype
package coheretruncatetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3a94b6715915b1e9311724a2614c643368eece90/specification/inference/_types/CommonTypes.ts#L717-L721
type CohereTruncateType struct {
	Name string
}

var (
	END = CohereTruncateType{"END"}

	NONE = CohereTruncateType{"NONE"}

	START = CohereTruncateType{"START"}
)

func (c CohereTruncateType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CohereTruncateType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "END":
		*c = END
	case "NONE":
		*c = NONE
	case "START":
		*c = START
	default:
		*c = CohereTruncateType{string(text)}
	}

	return nil
}

func (c CohereTruncateType) String() string {
	return c.Name
}
