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

// ReindexStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/reindex_rethrottle/types.ts#L37-L51
type ReindexStatus struct {
	Batches              int64                   `json:"batches"`
	Created              int64                   `json:"created"`
	Deleted              int64                   `json:"deleted"`
	Noops                int64                   `json:"noops"`
	RequestsPerSecond    float32                 `json:"requests_per_second"`
	Retries              Retries                 `json:"retries"`
	Throttled            *Duration               `json:"throttled,omitempty"`
	ThrottledMillis      DurationValueUnitMillis `json:"throttled_millis"`
	ThrottledUntil       *Duration               `json:"throttled_until,omitempty"`
	ThrottledUntilMillis DurationValueUnitMillis `json:"throttled_until_millis"`
	Total                int64                   `json:"total"`
	Updated              int64                   `json:"updated"`
	VersionConflicts     int64                   `json:"version_conflicts"`
}

// ReindexStatusBuilder holds ReindexStatus struct and provides a builder API.
type ReindexStatusBuilder struct {
	v *ReindexStatus
}

// NewReindexStatus provides a builder for the ReindexStatus struct.
func NewReindexStatusBuilder() *ReindexStatusBuilder {
	r := ReindexStatusBuilder{
		&ReindexStatus{},
	}

	return &r
}

// Build finalize the chain and returns the ReindexStatus struct
func (rb *ReindexStatusBuilder) Build() ReindexStatus {
	return *rb.v
}

func (rb *ReindexStatusBuilder) Batches(batches int64) *ReindexStatusBuilder {
	rb.v.Batches = batches
	return rb
}

func (rb *ReindexStatusBuilder) Created(created int64) *ReindexStatusBuilder {
	rb.v.Created = created
	return rb
}

func (rb *ReindexStatusBuilder) Deleted(deleted int64) *ReindexStatusBuilder {
	rb.v.Deleted = deleted
	return rb
}

func (rb *ReindexStatusBuilder) Noops(noops int64) *ReindexStatusBuilder {
	rb.v.Noops = noops
	return rb
}

func (rb *ReindexStatusBuilder) RequestsPerSecond(requestspersecond float32) *ReindexStatusBuilder {
	rb.v.RequestsPerSecond = requestspersecond
	return rb
}

func (rb *ReindexStatusBuilder) Retries(retries *RetriesBuilder) *ReindexStatusBuilder {
	v := retries.Build()
	rb.v.Retries = v
	return rb
}

func (rb *ReindexStatusBuilder) Throttled(throttled *DurationBuilder) *ReindexStatusBuilder {
	v := throttled.Build()
	rb.v.Throttled = &v
	return rb
}

func (rb *ReindexStatusBuilder) ThrottledMillis(throttledmillis *DurationValueUnitMillisBuilder) *ReindexStatusBuilder {
	v := throttledmillis.Build()
	rb.v.ThrottledMillis = v
	return rb
}

func (rb *ReindexStatusBuilder) ThrottledUntil(throttleduntil *DurationBuilder) *ReindexStatusBuilder {
	v := throttleduntil.Build()
	rb.v.ThrottledUntil = &v
	return rb
}

func (rb *ReindexStatusBuilder) ThrottledUntilMillis(throttleduntilmillis *DurationValueUnitMillisBuilder) *ReindexStatusBuilder {
	v := throttleduntilmillis.Build()
	rb.v.ThrottledUntilMillis = v
	return rb
}

func (rb *ReindexStatusBuilder) Total(total int64) *ReindexStatusBuilder {
	rb.v.Total = total
	return rb
}

func (rb *ReindexStatusBuilder) Updated(updated int64) *ReindexStatusBuilder {
	rb.v.Updated = updated
	return rb
}

func (rb *ReindexStatusBuilder) VersionConflicts(versionconflicts int64) *ReindexStatusBuilder {
	rb.v.VersionConflicts = versionconflicts
	return rb
}
