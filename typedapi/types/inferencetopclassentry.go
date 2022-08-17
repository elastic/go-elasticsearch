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

// InferenceTopClassEntry type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L623-L627
type InferenceTopClassEntry struct {
	ClassName        FieldValue `json:"class_name"`
	ClassProbability float64    `json:"class_probability"`
	ClassScore       float64    `json:"class_score"`
}

// InferenceTopClassEntryBuilder holds InferenceTopClassEntry struct and provides a builder API.
type InferenceTopClassEntryBuilder struct {
	v *InferenceTopClassEntry
}

// NewInferenceTopClassEntry provides a builder for the InferenceTopClassEntry struct.
func NewInferenceTopClassEntryBuilder() *InferenceTopClassEntryBuilder {
	r := InferenceTopClassEntryBuilder{
		&InferenceTopClassEntry{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceTopClassEntry struct
func (rb *InferenceTopClassEntryBuilder) Build() InferenceTopClassEntry {
	return *rb.v
}

func (rb *InferenceTopClassEntryBuilder) ClassName(classname *FieldValueBuilder) *InferenceTopClassEntryBuilder {
	v := classname.Build()
	rb.v.ClassName = v
	return rb
}

func (rb *InferenceTopClassEntryBuilder) ClassProbability(classprobability float64) *InferenceTopClassEntryBuilder {
	rb.v.ClassProbability = classprobability
	return rb
}

func (rb *InferenceTopClassEntryBuilder) ClassScore(classscore float64) *InferenceTopClassEntryBuilder {
	rb.v.ClassScore = classscore
	return rb
}
