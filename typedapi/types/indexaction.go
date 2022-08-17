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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

// IndexAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L256-L265
type IndexAction struct {
	DocId              *Id              `json:"doc_id,omitempty"`
	ExecutionTimeField *Field           `json:"execution_time_field,omitempty"`
	Index              IndexName        `json:"index"`
	OpType             *optype.OpType   `json:"op_type,omitempty"`
	Refresh            *refresh.Refresh `json:"refresh,omitempty"`
	Timeout            *Duration        `json:"timeout,omitempty"`
}

// IndexActionBuilder holds IndexAction struct and provides a builder API.
type IndexActionBuilder struct {
	v *IndexAction
}

// NewIndexAction provides a builder for the IndexAction struct.
func NewIndexActionBuilder() *IndexActionBuilder {
	r := IndexActionBuilder{
		&IndexAction{},
	}

	return &r
}

// Build finalize the chain and returns the IndexAction struct
func (rb *IndexActionBuilder) Build() IndexAction {
	return *rb.v
}

func (rb *IndexActionBuilder) DocId(docid Id) *IndexActionBuilder {
	rb.v.DocId = &docid
	return rb
}

func (rb *IndexActionBuilder) ExecutionTimeField(executiontimefield Field) *IndexActionBuilder {
	rb.v.ExecutionTimeField = &executiontimefield
	return rb
}

func (rb *IndexActionBuilder) Index(index IndexName) *IndexActionBuilder {
	rb.v.Index = index
	return rb
}

func (rb *IndexActionBuilder) OpType(optype optype.OpType) *IndexActionBuilder {
	rb.v.OpType = &optype
	return rb
}

func (rb *IndexActionBuilder) Refresh(refresh refresh.Refresh) *IndexActionBuilder {
	rb.v.Refresh = &refresh
	return rb
}

func (rb *IndexActionBuilder) Timeout(timeout *DurationBuilder) *IndexActionBuilder {
	v := timeout.Build()
	rb.v.Timeout = &v
	return rb
}
