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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)

// IndexResultSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L271-L277
type IndexResultSummary struct {
	Created bool          `json:"created"`
	Id      Id            `json:"id"`
	Index   IndexName     `json:"index"`
	Result  result.Result `json:"result"`
	Version VersionNumber `json:"version"`
}

// IndexResultSummaryBuilder holds IndexResultSummary struct and provides a builder API.
type IndexResultSummaryBuilder struct {
	v *IndexResultSummary
}

// NewIndexResultSummary provides a builder for the IndexResultSummary struct.
func NewIndexResultSummaryBuilder() *IndexResultSummaryBuilder {
	r := IndexResultSummaryBuilder{
		&IndexResultSummary{},
	}

	return &r
}

// Build finalize the chain and returns the IndexResultSummary struct
func (rb *IndexResultSummaryBuilder) Build() IndexResultSummary {
	return *rb.v
}

func (rb *IndexResultSummaryBuilder) Created(created bool) *IndexResultSummaryBuilder {
	rb.v.Created = created
	return rb
}

func (rb *IndexResultSummaryBuilder) Id(id Id) *IndexResultSummaryBuilder {
	rb.v.Id = id
	return rb
}

func (rb *IndexResultSummaryBuilder) Index(index IndexName) *IndexResultSummaryBuilder {
	rb.v.Index = index
	return rb
}

func (rb *IndexResultSummaryBuilder) Result(result result.Result) *IndexResultSummaryBuilder {
	rb.v.Result = result
	return rb
}

func (rb *IndexResultSummaryBuilder) Version(version VersionNumber) *IndexResultSummaryBuilder {
	rb.v.Version = version
	return rb
}
