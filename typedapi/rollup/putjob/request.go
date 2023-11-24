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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package putjob

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putjob
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/rollup/put_job/CreateRollupJobRequest.ts#L27-L89
type Request struct {

	// Cron A cron string which defines the intervals when the rollup job should be
	// executed. When the interval
	// triggers, the indexer attempts to rollup the data in the index pattern. The
	// cron pattern is unrelated
	// to the time interval of the data being rolled up. For example, you may wish
	// to create hourly rollups
	// of your document but to only run the indexer on a daily basis at midnight, as
	// defined by the cron. The
	// cron pattern is defined just like a Watcher cron schedule.
	Cron string `json:"cron"`
	// Groups Defines the grouping fields and aggregations that are defined for this rollup
	// job. These fields will then be
	// available later for aggregating into buckets. These aggs and fields can be
	// used in any combination. Think of
	// the groups configuration as defining a set of tools that can later be used in
	// aggregations to partition the
	// data. Unlike raw data, we have to think ahead to which fields and
	// aggregations might be used. Rollups provide
	// enough flexibility that you simply need to determine which fields are needed,
	// not in what order they are needed.
	Groups  types.Groupings   `json:"groups"`
	Headers types.HttpHeaders `json:"headers,omitempty"`
	// IndexPattern The index or index pattern to roll up. Supports wildcard-style patterns
	// (`logstash-*`). The job attempts to
	// rollup the entire index or index-pattern.
	IndexPattern string `json:"index_pattern"`
	// Metrics Defines the metrics to collect for each grouping tuple. By default, only the
	// doc_counts are collected for each
	// group. To make rollup useful, you will often add metrics like averages, mins,
	// maxes, etc. Metrics are defined
	// on a per-field basis and for each field you configure which metric should be
	// collected.
	Metrics []types.FieldMetric `json:"metrics,omitempty"`
	// PageSize The number of bucket results that are processed on each iteration of the
	// rollup indexer. A larger value tends
	// to execute faster, but requires more memory during processing. This value has
	// no effect on how the data is
	// rolled up; it is merely used for tweaking the speed or memory cost of the
	// indexer.
	PageSize int `json:"page_size"`
	// RollupIndex The index that contains the rollup results. The index can be shared with
	// other rollup jobs. The data is stored so that it doesn’t interfere with
	// unrelated jobs.
	RollupIndex string `json:"rollup_index"`
	// Timeout Time to wait for the request to complete.
	Timeout types.Duration `json:"timeout,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putjob request: %w", err)
	}

	return &req, nil
}
