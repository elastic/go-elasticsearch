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

// Package conditionoperator
package conditionoperator

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Rule.ts#L74-L79
type ConditionOperator struct {
	Name string
}

var (
	Gt = ConditionOperator{"gt"}

	Gte = ConditionOperator{"gte"}

	Lt = ConditionOperator{"lt"}

	Lte = ConditionOperator{"lte"}
)

func (c ConditionOperator) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ConditionOperator) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "gt":
		*c = Gt
	case "gte":
		*c = Gte
	case "lt":
		*c = Lt
	case "lte":
		*c = Lte
	default:
		*c = ConditionOperator{string(text)}
	}

	return nil
}

func (c ConditionOperator) String() string {
	return c.Name
}
