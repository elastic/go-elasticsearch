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

// QueryWatch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Watch.ts#L58-L64
type QueryWatch struct {
	Id_          Id              `json:"_id"`
	PrimaryTerm_ *int            `json:"_primary_term,omitempty"`
	SeqNo_       *SequenceNumber `json:"_seq_no,omitempty"`
	Status       *WatchStatus    `json:"status,omitempty"`
	Watch        *Watch          `json:"watch,omitempty"`
}

// QueryWatchBuilder holds QueryWatch struct and provides a builder API.
type QueryWatchBuilder struct {
	v *QueryWatch
}

// NewQueryWatch provides a builder for the QueryWatch struct.
func NewQueryWatchBuilder() *QueryWatchBuilder {
	r := QueryWatchBuilder{
		&QueryWatch{},
	}

	return &r
}

// Build finalize the chain and returns the QueryWatch struct
func (rb *QueryWatchBuilder) Build() QueryWatch {
	return *rb.v
}

func (rb *QueryWatchBuilder) Id_(id_ Id) *QueryWatchBuilder {
	rb.v.Id_ = id_
	return rb
}

func (rb *QueryWatchBuilder) PrimaryTerm_(primaryterm_ int) *QueryWatchBuilder {
	rb.v.PrimaryTerm_ = &primaryterm_
	return rb
}

func (rb *QueryWatchBuilder) SeqNo_(seqno_ SequenceNumber) *QueryWatchBuilder {
	rb.v.SeqNo_ = &seqno_
	return rb
}

func (rb *QueryWatchBuilder) Status(status *WatchStatusBuilder) *QueryWatchBuilder {
	v := status.Build()
	rb.v.Status = &v
	return rb
}

func (rb *QueryWatchBuilder) Watch(watch *WatchBuilder) *QueryWatchBuilder {
	v := watch.Build()
	rb.v.Watch = &v
	return rb
}
