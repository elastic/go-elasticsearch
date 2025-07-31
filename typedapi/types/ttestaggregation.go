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

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ttesttype"
)

// TTestAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/metric.ts#L326-L340
type TTestAggregation struct {
	// A Test population A.
	A *TestPopulation `json:"a,omitempty"`
	// B Test population B.
	B *TestPopulation `json:"b,omitempty"`
	// Type The type of test.
	Type *ttesttype.TTestType `json:"type,omitempty"`
}

// NewTTestAggregation returns a TTestAggregation.
func NewTTestAggregation() *TTestAggregation {
	r := &TTestAggregation{}

	return r
}
