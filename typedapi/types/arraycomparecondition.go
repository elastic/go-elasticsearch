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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/quantifier"
)

// ArrayCompareCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Conditions.ts#L25-L31
type ArrayCompareCondition struct {
	ArrayPath  string                `json:"array_path"`
	Comparison string                `json:"comparison"`
	Path       string                `json:"path"`
	Quantifier quantifier.Quantifier `json:"quantifier"`
	Value      interface{}           `json:"value,omitempty"`
}

// ArrayCompareConditionBuilder holds ArrayCompareCondition struct and provides a builder API.
type ArrayCompareConditionBuilder struct {
	v *ArrayCompareCondition
}

// NewArrayCompareCondition provides a builder for the ArrayCompareCondition struct.
func NewArrayCompareConditionBuilder() *ArrayCompareConditionBuilder {
	r := ArrayCompareConditionBuilder{
		&ArrayCompareCondition{},
	}

	return &r
}

// Build finalize the chain and returns the ArrayCompareCondition struct
func (rb *ArrayCompareConditionBuilder) Build() ArrayCompareCondition {
	return *rb.v
}

func (rb *ArrayCompareConditionBuilder) ArrayPath(arraypath string) *ArrayCompareConditionBuilder {
	rb.v.ArrayPath = arraypath
	return rb
}

func (rb *ArrayCompareConditionBuilder) Comparison(comparison string) *ArrayCompareConditionBuilder {
	rb.v.Comparison = comparison
	return rb
}

func (rb *ArrayCompareConditionBuilder) Path(path string) *ArrayCompareConditionBuilder {
	rb.v.Path = path
	return rb
}

func (rb *ArrayCompareConditionBuilder) Quantifier(quantifier quantifier.Quantifier) *ArrayCompareConditionBuilder {
	rb.v.Quantifier = quantifier
	return rb
}

func (rb *ArrayCompareConditionBuilder) Value(value interface{}) *ArrayCompareConditionBuilder {
	rb.v.Value = value
	return rb
}
