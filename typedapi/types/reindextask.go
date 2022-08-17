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

// ReindexTask type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/reindex_rethrottle/types.ts#L53-L64
type ReindexTask struct {
	Action             string                 `json:"action"`
	Cancellable        bool                   `json:"cancellable"`
	Description        string                 `json:"description"`
	Headers            HttpHeaders            `json:"headers"`
	Id                 int64                  `json:"id"`
	Node               Name                   `json:"node"`
	RunningTimeInNanos DurationValueUnitNanos `json:"running_time_in_nanos"`
	StartTimeInMillis  EpochTimeUnitMillis    `json:"start_time_in_millis"`
	Status             ReindexStatus          `json:"status"`
	Type               string                 `json:"type"`
}

// ReindexTaskBuilder holds ReindexTask struct and provides a builder API.
type ReindexTaskBuilder struct {
	v *ReindexTask
}

// NewReindexTask provides a builder for the ReindexTask struct.
func NewReindexTaskBuilder() *ReindexTaskBuilder {
	r := ReindexTaskBuilder{
		&ReindexTask{},
	}

	return &r
}

// Build finalize the chain and returns the ReindexTask struct
func (rb *ReindexTaskBuilder) Build() ReindexTask {
	return *rb.v
}

func (rb *ReindexTaskBuilder) Action(action string) *ReindexTaskBuilder {
	rb.v.Action = action
	return rb
}

func (rb *ReindexTaskBuilder) Cancellable(cancellable bool) *ReindexTaskBuilder {
	rb.v.Cancellable = cancellable
	return rb
}

func (rb *ReindexTaskBuilder) Description(description string) *ReindexTaskBuilder {
	rb.v.Description = description
	return rb
}

func (rb *ReindexTaskBuilder) Headers(headers *HttpHeadersBuilder) *ReindexTaskBuilder {
	v := headers.Build()
	rb.v.Headers = v
	return rb
}

func (rb *ReindexTaskBuilder) Id(id int64) *ReindexTaskBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ReindexTaskBuilder) Node(node Name) *ReindexTaskBuilder {
	rb.v.Node = node
	return rb
}

func (rb *ReindexTaskBuilder) RunningTimeInNanos(runningtimeinnanos *DurationValueUnitNanosBuilder) *ReindexTaskBuilder {
	v := runningtimeinnanos.Build()
	rb.v.RunningTimeInNanos = v
	return rb
}

func (rb *ReindexTaskBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *ReindexTaskBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}

func (rb *ReindexTaskBuilder) Status(status *ReindexStatusBuilder) *ReindexTaskBuilder {
	v := status.Build()
	rb.v.Status = v
	return rb
}

func (rb *ReindexTaskBuilder) Type_(type_ string) *ReindexTaskBuilder {
	rb.v.Type = type_
	return rb
}
