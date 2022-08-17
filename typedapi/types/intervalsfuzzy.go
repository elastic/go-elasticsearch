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

// IntervalsFuzzy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L88-L97
type IntervalsFuzzy struct {
	Analyzer       *string    `json:"analyzer,omitempty"`
	Fuzziness      *Fuzziness `json:"fuzziness,omitempty"`
	PrefixLength   *int       `json:"prefix_length,omitempty"`
	Term           string     `json:"term"`
	Transpositions *bool      `json:"transpositions,omitempty"`
	UseField       *Field     `json:"use_field,omitempty"`
}

// IntervalsFuzzyBuilder holds IntervalsFuzzy struct and provides a builder API.
type IntervalsFuzzyBuilder struct {
	v *IntervalsFuzzy
}

// NewIntervalsFuzzy provides a builder for the IntervalsFuzzy struct.
func NewIntervalsFuzzyBuilder() *IntervalsFuzzyBuilder {
	r := IntervalsFuzzyBuilder{
		&IntervalsFuzzy{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsFuzzy struct
func (rb *IntervalsFuzzyBuilder) Build() IntervalsFuzzy {
	return *rb.v
}

func (rb *IntervalsFuzzyBuilder) Analyzer(analyzer string) *IntervalsFuzzyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *IntervalsFuzzyBuilder) Fuzziness(fuzziness *FuzzinessBuilder) *IntervalsFuzzyBuilder {
	v := fuzziness.Build()
	rb.v.Fuzziness = &v
	return rb
}

func (rb *IntervalsFuzzyBuilder) PrefixLength(prefixlength int) *IntervalsFuzzyBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

func (rb *IntervalsFuzzyBuilder) Term(term string) *IntervalsFuzzyBuilder {
	rb.v.Term = term
	return rb
}

func (rb *IntervalsFuzzyBuilder) Transpositions(transpositions bool) *IntervalsFuzzyBuilder {
	rb.v.Transpositions = &transpositions
	return rb
}

func (rb *IntervalsFuzzyBuilder) UseField(usefield Field) *IntervalsFuzzyBuilder {
	rb.v.UseField = &usefield
	return rb
}
