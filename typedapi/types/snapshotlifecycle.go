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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// SnapshotLifecycle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/slm/_types/SnapshotLifecycle.ts#L32-L43
type SnapshotLifecycle struct {
	InProgress          *InProgress   `json:"in_progress,omitempty"`
	LastFailure         *Invocation   `json:"last_failure,omitempty"`
	LastSuccess         *Invocation   `json:"last_success,omitempty"`
	ModifiedDate        *DateString   `json:"modified_date,omitempty"`
	ModifiedDateMillis  EpochMillis   `json:"modified_date_millis"`
	NextExecution       *DateString   `json:"next_execution,omitempty"`
	NextExecutionMillis EpochMillis   `json:"next_execution_millis"`
	Policy              Policy        `json:"policy"`
	Stats               Statistics    `json:"stats"`
	Version             VersionNumber `json:"version"`
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

func (rb *SnapshotLifecycleBuilder) ModifiedDate(modifieddate DateString) *SnapshotLifecycleBuilder {
	rb.v.ModifiedDate = &modifieddate
	return rb
}

func (rb *SnapshotLifecycleBuilder) ModifiedDateMillis(modifieddatemillis *EpochMillisBuilder) *SnapshotLifecycleBuilder {
	v := modifieddatemillis.Build()
	rb.v.ModifiedDateMillis = v
	return rb
}

func (rb *SnapshotLifecycleBuilder) NextExecution(nextexecution DateString) *SnapshotLifecycleBuilder {
	rb.v.NextExecution = &nextexecution
	return rb
}

func (rb *SnapshotLifecycleBuilder) NextExecutionMillis(nextexecutionmillis *EpochMillisBuilder) *SnapshotLifecycleBuilder {
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
