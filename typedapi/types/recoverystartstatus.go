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

// RecoveryStartStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/indices/recovery/types.ts#L85-L88
type RecoveryStartStatus struct {
	CheckIndexTime    int64  `json:"check_index_time"`
	TotalTimeInMillis string `json:"total_time_in_millis"`
}

// RecoveryStartStatusBuilder holds RecoveryStartStatus struct and provides a builder API.
type RecoveryStartStatusBuilder struct {
	v *RecoveryStartStatus
}

// NewRecoveryStartStatus provides a builder for the RecoveryStartStatus struct.
func NewRecoveryStartStatusBuilder() *RecoveryStartStatusBuilder {
	r := RecoveryStartStatusBuilder{
		&RecoveryStartStatus{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryStartStatus struct
func (rb *RecoveryStartStatusBuilder) Build() RecoveryStartStatus {
	return *rb.v
}

func (rb *RecoveryStartStatusBuilder) CheckIndexTime(checkindextime int64) *RecoveryStartStatusBuilder {
	rb.v.CheckIndexTime = checkindextime
	return rb
}

func (rb *RecoveryStartStatusBuilder) TotalTimeInMillis(totaltimeinmillis string) *RecoveryStartStatusBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
