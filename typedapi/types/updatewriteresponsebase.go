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

// UpdateWriteResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/update/UpdateResponse.ts#L23-L25
type UpdateWriteResponseBase struct {
	ForcedRefresh *bool           `json:"forced_refresh,omitempty"`
	Get           *InlineGet      `json:"get,omitempty"`
	Id_           Id              `json:"_id"`
	Index_        IndexName       `json:"_index"`
	PrimaryTerm_  int64           `json:"_primary_term"`
	Result        result.Result   `json:"result"`
	SeqNo_        SequenceNumber  `json:"_seq_no"`
	Shards_       ShardStatistics `json:"_shards"`
	Version_      VersionNumber   `json:"_version"`
}

// UpdateWriteResponseBaseBuilder holds UpdateWriteResponseBase struct and provides a builder API.
type UpdateWriteResponseBaseBuilder struct {
	v *UpdateWriteResponseBase
}

// NewUpdateWriteResponseBase provides a builder for the UpdateWriteResponseBase struct.
func NewUpdateWriteResponseBaseBuilder() *UpdateWriteResponseBaseBuilder {
	r := UpdateWriteResponseBaseBuilder{
		&UpdateWriteResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the UpdateWriteResponseBase struct
func (rb *UpdateWriteResponseBaseBuilder) Build() UpdateWriteResponseBase {
	return *rb.v
}

func (rb *UpdateWriteResponseBaseBuilder) ForcedRefresh(forcedrefresh bool) *UpdateWriteResponseBaseBuilder {
	rb.v.ForcedRefresh = &forcedrefresh
	return rb
}

func (rb *UpdateWriteResponseBaseBuilder) Get(get *InlineGetBuilder) *UpdateWriteResponseBaseBuilder {
	v := get.Build()
	rb.v.Get = &v
	return rb
}

func (rb *UpdateWriteResponseBaseBuilder) Id_(id_ Id) *UpdateWriteResponseBaseBuilder {
	rb.v.Id_ = id_
	return rb
}

func (rb *UpdateWriteResponseBaseBuilder) Index_(index_ IndexName) *UpdateWriteResponseBaseBuilder {
	rb.v.Index_ = index_
	return rb
}

func (rb *UpdateWriteResponseBaseBuilder) PrimaryTerm_(primaryterm_ int64) *UpdateWriteResponseBaseBuilder {
	rb.v.PrimaryTerm_ = primaryterm_
	return rb
}

func (rb *UpdateWriteResponseBaseBuilder) Result(result result.Result) *UpdateWriteResponseBaseBuilder {
	rb.v.Result = result
	return rb
}

func (rb *UpdateWriteResponseBaseBuilder) SeqNo_(seqno_ SequenceNumber) *UpdateWriteResponseBaseBuilder {
	rb.v.SeqNo_ = seqno_
	return rb
}

func (rb *UpdateWriteResponseBaseBuilder) Shards_(shards_ *ShardStatisticsBuilder) *UpdateWriteResponseBaseBuilder {
	v := shards_.Build()
	rb.v.Shards_ = v
	return rb
}

func (rb *UpdateWriteResponseBaseBuilder) Version_(version_ VersionNumber) *UpdateWriteResponseBaseBuilder {
	rb.v.Version_ = version_
	return rb
}
