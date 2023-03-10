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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// PerPartitionCategorization type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/Analysis.ts#L93-L102
type PerPartitionCategorization struct {
	// Enabled To enable this setting, you must also set the `partition_field_name` property
	// to the same value in every detector that uses the keyword `mlcategory`.
	// Otherwise, job creation fails.
	Enabled *bool `json:"enabled,omitempty"`
	// StopOnWarn This setting can be set to true only if per-partition categorization is
	// enabled. If true, both categorization and subsequent anomaly detection stops
	// for partitions where the categorization status changes to warn. This setting
	// makes it viable to have a job where it is expected that categorization works
	// well for some partitions but not others; you do not pay the cost of bad
	// categorization forever in the partitions where it works badly.
	StopOnWarn *bool `json:"stop_on_warn,omitempty"`
}

// NewPerPartitionCategorization returns a PerPartitionCategorization.
func NewPerPartitionCategorization() *PerPartitionCategorization {
	r := &PerPartitionCategorization{}

	return r
}
