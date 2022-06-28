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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// MultiTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/bucket.ts#L266-L268
type MultiTermsAggregation struct {
	Meta  *Metadata         `json:"meta,omitempty"`
	Name  *string           `json:"name,omitempty"`
	Terms []MultiTermLookup `json:"terms"`
}

// MultiTermsAggregationBuilder holds MultiTermsAggregation struct and provides a builder API.
type MultiTermsAggregationBuilder struct {
	v *MultiTermsAggregation
}

// NewMultiTermsAggregation provides a builder for the MultiTermsAggregation struct.
func NewMultiTermsAggregationBuilder() *MultiTermsAggregationBuilder {
	r := MultiTermsAggregationBuilder{
		&MultiTermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MultiTermsAggregation struct
func (rb *MultiTermsAggregationBuilder) Build() MultiTermsAggregation {
	return *rb.v
}

func (rb *MultiTermsAggregationBuilder) Meta(meta *MetadataBuilder) *MultiTermsAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *MultiTermsAggregationBuilder) Name(name string) *MultiTermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *MultiTermsAggregationBuilder) Terms(terms []MultiTermLookupBuilder) *MultiTermsAggregationBuilder {
	tmp := make([]MultiTermLookup, len(terms))
	for _, value := range terms {
		tmp = append(tmp, value.Build())
	}
	rb.v.Terms = tmp
	return rb
}
