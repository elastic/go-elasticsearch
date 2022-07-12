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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// SlowlogTresholds type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/indices/_types/IndexSettings.ts#L478-L487
type SlowlogTresholds struct {
	Fetch *SlowlogTresholdLevels `json:"fetch,omitempty"`
	// Index The indexing slow log, similar in functionality to the search slow log. The
	// log file name ends with `_index_indexing_slowlog.json`.
	// Log and the thresholds are configured in the same way as the search slowlog.
	Index *SlowlogTresholdLevels `json:"index,omitempty"`
	Query *SlowlogTresholdLevels `json:"query,omitempty"`
}

// SlowlogTresholdsBuilder holds SlowlogTresholds struct and provides a builder API.
type SlowlogTresholdsBuilder struct {
	v *SlowlogTresholds
}

// NewSlowlogTresholds provides a builder for the SlowlogTresholds struct.
func NewSlowlogTresholdsBuilder() *SlowlogTresholdsBuilder {
	r := SlowlogTresholdsBuilder{
		&SlowlogTresholds{},
	}

	return &r
}

// Build finalize the chain and returns the SlowlogTresholds struct
func (rb *SlowlogTresholdsBuilder) Build() SlowlogTresholds {
	return *rb.v
}

func (rb *SlowlogTresholdsBuilder) Fetch(fetch *SlowlogTresholdLevelsBuilder) *SlowlogTresholdsBuilder {
	v := fetch.Build()
	rb.v.Fetch = &v
	return rb
}

// Index The indexing slow log, similar in functionality to the search slow log. The
// log file name ends with `_index_indexing_slowlog.json`.
// Log and the thresholds are configured in the same way as the search slowlog.

func (rb *SlowlogTresholdsBuilder) Index(index *SlowlogTresholdLevelsBuilder) *SlowlogTresholdsBuilder {
	v := index.Build()
	rb.v.Index = &v
	return rb
}

func (rb *SlowlogTresholdsBuilder) Query(query *SlowlogTresholdLevelsBuilder) *SlowlogTresholdsBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}
