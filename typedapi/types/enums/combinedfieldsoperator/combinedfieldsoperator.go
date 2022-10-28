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
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


// Package combinedfieldsoperator
package combinedfieldsoperator

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_types/query_dsl/abstractions.ts#L202-L205
type CombinedFieldsOperator struct {
	name string
}

var (
	Or = CombinedFieldsOperator{"or"}

	And = CombinedFieldsOperator{"and"}
)

func (c CombinedFieldsOperator) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CombinedFieldsOperator) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "or":
		*c = Or
	case "and":
		*c = And
	default:
		*c = CombinedFieldsOperator{string(text)}
	}

	return nil
}

func (c CombinedFieldsOperator) String() string {
	return c.name
}
