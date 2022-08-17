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

// SnapshotLifecycle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/slm/_types/SnapshotLifecycle.ts#L38-L49
type SnapshotLifecycle struct {
	InProgress          *InProgress         `json:"in_progress,omitempty"`
	LastFailure         *Invocation         `json:"last_failure,omitempty"`
	LastSuccess         *Invocation         `json:"last_success,omitempty"`
	ModifiedDate        *DateTime           `json:"modified_date,omitempty"`
	ModifiedDateMillis  EpochTimeUnitMillis `json:"modified_date_millis"`
	NextExecution       *DateTime           `json:"next_execution,omitempty"`
	NextExecutionMillis EpochTimeUnitMillis `json:"next_execution_millis"`
	Policy              Policy              `json:"policy"`
	Stats               Statistics          `json:"stats"`
	Version             VersionNumber       `json:"version"`
}

// SnapshotLifecycleBuilder holds SnapshotLifecycle struct and provides a builder API.
type SnapshotLifecycleBuilder struct {
	v *SnapshotLifecycle
}

// NewSnapshotLifecycle provides a builder for the SnapshotLifecycle struct.
func NewSnapshotLifecycleBuilder() *SnapshotLifecycleBuilder {
	r := SnapshotLifecycleBuilder{
		&SnapshotLifecycle{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotLifecycle struct
func (rb *SnapshotLifecycleBuilder) Build() SnapshotLifecycle {
	return *rb.v
}

func (rb *SnapshotLifecycleBuilder) InProgress(inprogress *InProgressBuilder) *SnapshotLifecycleBuilder {
	v := inprogress.Build()
	rb.v.InProgress = &v
	return rb
}

func (rb *SnapshotLifecycleBuilder) LastFailure(lastfailure *InvocationBuilder) *SnapshotLifecycleBuilder {
	v := lastfailure.Build()
	rb.v.LastFailure = &v
	return rb
}

func (rb *SnapshotLifecycleBuilder) LastSuccess(lastsuccess *InvocationBuilder) *SnapshotLifecycleBuilder {
	v := lastsuccess.Build()
	rb.v.LastSuccess = &v
	return rb
}

func (rb *SnapshotLifecycleBuilder) ModifiedDate(modifieddate *DateTimeBuilder) *SnapshotLifecycleBuilder {
	v := modifieddate.Build()
	rb.v.ModifiedDate = &v
	return rb
}

func (rb *SnapshotLifecycleBuilder) ModifiedDateMillis(modifieddatemillis *EpochTimeUnitMillisBuilder) *SnapshotLifecycleBuilder {
	v := modifieddatemillis.Build()
	rb.v.ModifiedDateMillis = v
	return rb
}

func (rb *SnapshotLifecycleBuilder) NextExecution(nextexecution *DateTimeBuilder) *SnapshotLifecycleBuilder {
	v := nextexecution.Build()
	rb.v.NextExecution = &v
	return rb
}

func (rb *SnapshotLifecycleBuilder) NextExecutionMillis(nextexecutionmillis *EpochTimeUnitMillisBuilder) *SnapshotLifecycleBuilder {
	v := nextexecutionmillis.Build()
	rb.v.NextExecutionMillis = v
	return rb
}

func (rb *SnapshotLifecycleBuilder) Policy(policy *PolicyBuilder) *SnapshotLifecycleBuilder {
	v := policy.Build()
	rb.v.Policy = v
	return rb
}

func (rb *SnapshotLifecycleBuilder) Stats(stats *StatisticsBuilder) *SnapshotLifecycleBuilder {
	v := stats.Build()
	rb.v.Stats = v
	return rb
}

func (rb *SnapshotLifecycleBuilder) Version(version VersionNumber) *SnapshotLifecycleBuilder {
	rb.v.Version = version
	return rb
}
