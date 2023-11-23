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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package childscoremode
package childscoremode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/joining.ts#L25-L39
type ChildScoreMode struct {
	Name string
}

var (
	None = ChildScoreMode{"none"}

	Avg = ChildScoreMode{"avg"}

	Sum = ChildScoreMode{"sum"}

	Max = ChildScoreMode{"max"}

	Min = ChildScoreMode{"min"}
)

func (c ChildScoreMode) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ChildScoreMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "none":
		*c = None
	case "avg":
		*c = Avg
	case "sum":
		*c = Sum
	case "max":
		*c = Max
	case "min":
		*c = Min
	default:
		*c = ChildScoreMode{string(text)}
	}

	return nil
}

func (c ChildScoreMode) String() string {
	return c.Name
}
