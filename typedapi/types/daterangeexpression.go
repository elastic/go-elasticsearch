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

// DateRangeExpression type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L146-L150
type DateRangeExpression struct {
	From *FieldDateMath `json:"from,omitempty"`
	Key  *string        `json:"key,omitempty"`
	To   *FieldDateMath `json:"to,omitempty"`
}

// DateRangeExpressionBuilder holds DateRangeExpression struct and provides a builder API.
type DateRangeExpressionBuilder struct {
	v *DateRangeExpression
}

// NewDateRangeExpression provides a builder for the DateRangeExpression struct.
func NewDateRangeExpressionBuilder() *DateRangeExpressionBuilder {
	r := DateRangeExpressionBuilder{
		&DateRangeExpression{},
	}

	return &r
}

// Build finalize the chain and returns the DateRangeExpression struct
func (rb *DateRangeExpressionBuilder) Build() DateRangeExpression {
	return *rb.v
}

func (rb *DateRangeExpressionBuilder) From(from *FieldDateMathBuilder) *DateRangeExpressionBuilder {
	v := from.Build()
	rb.v.From = &v
	return rb
}

func (rb *DateRangeExpressionBuilder) Key(key string) *DateRangeExpressionBuilder {
	rb.v.Key = &key
	return rb
}

func (rb *DateRangeExpressionBuilder) To(to *FieldDateMathBuilder) *DateRangeExpressionBuilder {
	v := to.Build()
	rb.v.To = &v
	return rb
}
