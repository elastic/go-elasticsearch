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

// TransformProgress type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/get_transform_stats/types.ts#L40-L46
type TransformProgress struct {
	DocsIndexed     int64   `json:"docs_indexed"`
	DocsProcessed   int64   `json:"docs_processed"`
	DocsRemaining   int64   `json:"docs_remaining"`
	PercentComplete float64 `json:"percent_complete"`
	TotalDocs       int64   `json:"total_docs"`
}

// TransformProgressBuilder holds TransformProgress struct and provides a builder API.
type TransformProgressBuilder struct {
	v *TransformProgress
}

// NewTransformProgress provides a builder for the TransformProgress struct.
func NewTransformProgressBuilder() *TransformProgressBuilder {
	r := TransformProgressBuilder{
		&TransformProgress{},
	}

	return &r
}

// Build finalize the chain and returns the TransformProgress struct
func (rb *TransformProgressBuilder) Build() TransformProgress {
	return *rb.v
}

func (rb *TransformProgressBuilder) DocsIndexed(docsindexed int64) *TransformProgressBuilder {
	rb.v.DocsIndexed = docsindexed
	return rb
}

func (rb *TransformProgressBuilder) DocsProcessed(docsprocessed int64) *TransformProgressBuilder {
	rb.v.DocsProcessed = docsprocessed
	return rb
}

func (rb *TransformProgressBuilder) DocsRemaining(docsremaining int64) *TransformProgressBuilder {
	rb.v.DocsRemaining = docsremaining
	return rb
}

func (rb *TransformProgressBuilder) PercentComplete(percentcomplete float64) *TransformProgressBuilder {
	rb.v.PercentComplete = percentcomplete
	return rb
}

func (rb *TransformProgressBuilder) TotalDocs(totaldocs int64) *TransformProgressBuilder {
	rb.v.TotalDocs = totaldocs
	return rb
}
