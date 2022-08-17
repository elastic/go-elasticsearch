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

// InProgress type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/slm/_types/SnapshotLifecycle.ts#L131-L136
type InProgress struct {
	Name            Name                `json:"name"`
	StartTimeMillis EpochTimeUnitMillis `json:"start_time_millis"`
	State           string              `json:"state"`
	Uuid            Uuid                `json:"uuid"`
}

// InProgressBuilder holds InProgress struct and provides a builder API.
type InProgressBuilder struct {
	v *InProgress
}

// NewInProgress provides a builder for the InProgress struct.
func NewInProgressBuilder() *InProgressBuilder {
	r := InProgressBuilder{
		&InProgress{},
	}

	return &r
}

// Build finalize the chain and returns the InProgress struct
func (rb *InProgressBuilder) Build() InProgress {
	return *rb.v
}

func (rb *InProgressBuilder) Name(name Name) *InProgressBuilder {
	rb.v.Name = name
	return rb
}

func (rb *InProgressBuilder) StartTimeMillis(starttimemillis *EpochTimeUnitMillisBuilder) *InProgressBuilder {
	v := starttimemillis.Build()
	rb.v.StartTimeMillis = v
	return rb
}

func (rb *InProgressBuilder) State(state string) *InProgressBuilder {
	rb.v.State = state
	return rb
}

func (rb *InProgressBuilder) Uuid(uuid Uuid) *InProgressBuilder {
	rb.v.Uuid = uuid
	return rb
}
