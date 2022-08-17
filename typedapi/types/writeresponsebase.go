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

// WriteResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Base.ts#L36-L45
type WriteResponseBase struct {
	ForcedRefresh *bool           `json:"forced_refresh,omitempty"`
	Id_           Id              `json:"_id"`
	Index_        IndexName       `json:"_index"`
	PrimaryTerm_  int64           `json:"_primary_term"`
	Result        result.Result   `json:"result"`
	SeqNo_        SequenceNumber  `json:"_seq_no"`
	Shards_       ShardStatistics `json:"_shards"`
	Version_      VersionNumber   `json:"_version"`
}

// WriteResponseBaseBuilder holds WriteResponseBase struct and provides a builder API.
type WriteResponseBaseBuilder struct {
	v *WriteResponseBase
}

// NewWriteResponseBase provides a builder for the WriteResponseBase struct.
func NewWriteResponseBaseBuilder() *WriteResponseBaseBuilder {
	r := WriteResponseBaseBuilder{
		&WriteResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the WriteResponseBase struct
func (rb *WriteResponseBaseBuilder) Build() WriteResponseBase {
	return *rb.v
}

func (rb *WriteResponseBaseBuilder) ForcedRefresh(forcedrefresh bool) *WriteResponseBaseBuilder {
	rb.v.ForcedRefresh = &forcedrefresh
	return rb
}

func (rb *WriteResponseBaseBuilder) Id_(id_ Id) *WriteResponseBaseBuilder {
	rb.v.Id_ = id_
	return rb
}

func (rb *WriteResponseBaseBuilder) Index_(index_ IndexName) *WriteResponseBaseBuilder {
	rb.v.Index_ = index_
	return rb
}

func (rb *WriteResponseBaseBuilder) PrimaryTerm_(primaryterm_ int64) *WriteResponseBaseBuilder {
	rb.v.PrimaryTerm_ = primaryterm_
	return rb
}

func (rb *WriteResponseBaseBuilder) Result(result result.Result) *WriteResponseBaseBuilder {
	rb.v.Result = result
	return rb
}

func (rb *WriteResponseBaseBuilder) SeqNo_(seqno_ SequenceNumber) *WriteResponseBaseBuilder {
	rb.v.SeqNo_ = seqno_
	return rb
}

func (rb *WriteResponseBaseBuilder) Shards_(shards_ *ShardStatisticsBuilder) *WriteResponseBaseBuilder {
	v := shards_.Build()
	rb.v.Shards_ = v
	return rb
}

func (rb *WriteResponseBaseBuilder) Version_(version_ VersionNumber) *WriteResponseBaseBuilder {
	rb.v.Version_ = version_
	return rb
}
