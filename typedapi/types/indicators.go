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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package types

// Indicators type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cbfcc73d01310bed2a480ec35aaef98138b598e5/specification/_global/health_report/types.ts#L32-L42
type Indicators struct {
	DataStreamLifecycle *DataStreamLifecycleIndicator `json:"data_stream_lifecycle,omitempty"`
	Disk                *DiskIndicator                `json:"disk,omitempty"`
	FileSettings        *FileSettingsIndicator        `json:"file_settings,omitempty"`
	Ilm                 *IlmIndicator                 `json:"ilm,omitempty"`
	MasterIsStable      *MasterIsStableIndicator      `json:"master_is_stable,omitempty"`
	RepositoryIntegrity *RepositoryIntegrityIndicator `json:"repository_integrity,omitempty"`
	ShardsAvailability  *ShardsAvailabilityIndicator  `json:"shards_availability,omitempty"`
	ShardsCapacity      *ShardsCapacityIndicator      `json:"shards_capacity,omitempty"`
	Slm                 *SlmIndicator                 `json:"slm,omitempty"`
}

// NewIndicators returns a Indicators.
func NewIndicators() *Indicators {
	r := &Indicators{}

	return r
}

// false
