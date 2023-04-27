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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

// Indicators type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_global/health_report/types.ts#L32-L39
type Indicators struct {
	Disk                *DiskIndicator                `json:"disk,omitempty"`
	Ilm                 *IlmIndicator                 `json:"ilm,omitempty"`
	MasterIsStable      *MasterIsStableIndicator      `json:"master_is_stable,omitempty"`
	RepositoryIntegrity *RepositoryIntegrityIndicator `json:"repository_integrity,omitempty"`
	ShardsAvailability  *ShardsAvailabilityIndicator  `json:"shards_availability,omitempty"`
	Slm                 *SlmIndicator                 `json:"slm,omitempty"`
}

// NewIndicators returns a Indicators.
func NewIndicators() *Indicators {
	r := &Indicators{}

	return r
}
