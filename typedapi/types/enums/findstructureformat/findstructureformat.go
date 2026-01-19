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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

// Package findstructureformat
package findstructureformat

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/text_structure/find_structure/FindStructureRequest.ts#L213-L218
type FindStructureFormat struct {
	Name string
}

var (
	Ndjson = FindStructureFormat{"ndjson"}

	Xml = FindStructureFormat{"xml"}

	Delimited = FindStructureFormat{"delimited"}

	Semistructuredtext = FindStructureFormat{"semi_structured_text"}
)

func (f FindStructureFormat) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FindStructureFormat) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "ndjson":
		*f = Ndjson
	case "xml":
		*f = Xml
	case "delimited":
		*f = Delimited
	case "semi_structured_text":
		*f = Semistructuredtext
	default:
		*f = FindStructureFormat{string(text)}
	}

	return nil
}

func (f FindStructureFormat) String() string {
	return f.Name
}
