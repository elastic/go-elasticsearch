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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

// RollupJob type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/rollup/get_jobs/types.ts#L28-L43
type RollupJob struct {
	// Config The rollup job configuration.
	Config RollupJobConfiguration `json:"config"`
	// Stats Transient statistics about the rollup job, such as how many documents have
	// been processed and how many rollup summary docs have been indexed.
	// These stats are not persisted.
	// If a node is restarted, these stats are reset.
	Stats RollupJobStats `json:"stats"`
	// Status The current status of the indexer for the rollup job.
	Status RollupJobStatus `json:"status"`
}

// NewRollupJob returns a RollupJob.
func NewRollupJob() *RollupJob {
	r := &RollupJob{}

	return r
}
