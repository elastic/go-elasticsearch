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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

package types

// SummaryInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/snapshot/repository_analyze/SnapshotAnalyzeRepositoryResponse.ts#L193-L202
type SummaryInfo struct {
	// Read A collection of statistics that summarise the results of the read operations
	// in the test.
	Read ReadSummaryInfo `json:"read"`
	// Write A collection of statistics that summarise the results of the write operations
	// in the test.
	Write WriteSummaryInfo `json:"write"`
}

// NewSummaryInfo returns a SummaryInfo.
func NewSummaryInfo() *SummaryInfo {
	r := &SummaryInfo{}

	return r
}

// false
