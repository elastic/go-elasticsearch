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

// Ingest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L138-L141
type Ingest struct {
	Pipelines map[string]IngestTotal `json:"pipelines,omitempty"`
	Total     *IngestTotal           `json:"total,omitempty"`
}

// IngestBuilder holds Ingest struct and provides a builder API.
type IngestBuilder struct {
	v *Ingest
}

// NewIngest provides a builder for the Ingest struct.
func NewIngestBuilder() *IngestBuilder {
	r := IngestBuilder{
		&Ingest{
			Pipelines: make(map[string]IngestTotal, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Ingest struct
func (rb *IngestBuilder) Build() Ingest {
	return *rb.v
}

func (rb *IngestBuilder) Pipelines(values map[string]*IngestTotalBuilder) *IngestBuilder {
	tmp := make(map[string]IngestTotal, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Pipelines = tmp
	return rb
}

func (rb *IngestBuilder) Total(total *IngestTotalBuilder) *IngestBuilder {
	v := total.Build()
	rb.v.Total = &v
	return rb
}
