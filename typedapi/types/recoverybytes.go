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

// RecoveryBytes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/recovery/types.ts#L38-L48
type RecoveryBytes struct {
	Percent                      Percentage `json:"percent"`
	Recovered                    *ByteSize  `json:"recovered,omitempty"`
	RecoveredFromSnapshot        *ByteSize  `json:"recovered_from_snapshot,omitempty"`
	RecoveredFromSnapshotInBytes *ByteSize  `json:"recovered_from_snapshot_in_bytes,omitempty"`
	RecoveredInBytes             ByteSize   `json:"recovered_in_bytes"`
	Reused                       *ByteSize  `json:"reused,omitempty"`
	ReusedInBytes                ByteSize   `json:"reused_in_bytes"`
	Total                        *ByteSize  `json:"total,omitempty"`
	TotalInBytes                 ByteSize   `json:"total_in_bytes"`
}

// RecoveryBytesBuilder holds RecoveryBytes struct and provides a builder API.
type RecoveryBytesBuilder struct {
	v *RecoveryBytes
}

// NewRecoveryBytes provides a builder for the RecoveryBytes struct.
func NewRecoveryBytesBuilder() *RecoveryBytesBuilder {
	r := RecoveryBytesBuilder{
		&RecoveryBytes{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryBytes struct
func (rb *RecoveryBytesBuilder) Build() RecoveryBytes {
	return *rb.v
}

func (rb *RecoveryBytesBuilder) Percent(percent *PercentageBuilder) *RecoveryBytesBuilder {
	v := percent.Build()
	rb.v.Percent = v
	return rb
}

func (rb *RecoveryBytesBuilder) Recovered(recovered *ByteSizeBuilder) *RecoveryBytesBuilder {
	v := recovered.Build()
	rb.v.Recovered = &v
	return rb
}

func (rb *RecoveryBytesBuilder) RecoveredFromSnapshot(recoveredfromsnapshot *ByteSizeBuilder) *RecoveryBytesBuilder {
	v := recoveredfromsnapshot.Build()
	rb.v.RecoveredFromSnapshot = &v
	return rb
}

func (rb *RecoveryBytesBuilder) RecoveredFromSnapshotInBytes(recoveredfromsnapshotinbytes *ByteSizeBuilder) *RecoveryBytesBuilder {
	v := recoveredfromsnapshotinbytes.Build()
	rb.v.RecoveredFromSnapshotInBytes = &v
	return rb
}

func (rb *RecoveryBytesBuilder) RecoveredInBytes(recoveredinbytes *ByteSizeBuilder) *RecoveryBytesBuilder {
	v := recoveredinbytes.Build()
	rb.v.RecoveredInBytes = v
	return rb
}

func (rb *RecoveryBytesBuilder) Reused(reused *ByteSizeBuilder) *RecoveryBytesBuilder {
	v := reused.Build()
	rb.v.Reused = &v
	return rb
}

func (rb *RecoveryBytesBuilder) ReusedInBytes(reusedinbytes *ByteSizeBuilder) *RecoveryBytesBuilder {
	v := reusedinbytes.Build()
	rb.v.ReusedInBytes = v
	return rb
}

func (rb *RecoveryBytesBuilder) Total(total *ByteSizeBuilder) *RecoveryBytesBuilder {
	v := total.Build()
	rb.v.Total = &v
	return rb
}

func (rb *RecoveryBytesBuilder) TotalInBytes(totalinbytes *ByteSizeBuilder) *RecoveryBytesBuilder {
	v := totalinbytes.Build()
	rb.v.TotalInBytes = v
	return rb
}
