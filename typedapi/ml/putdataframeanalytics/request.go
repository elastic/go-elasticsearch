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

package putdataframeanalytics

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putdataframeanalytics
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/put_data_frame_analytics/MlPutDataFrameAnalyticsRequest.ts#L30-L139
type Request struct {

	// AllowLazyStart Specifies whether this job can start when there is insufficient machine
	// learning node capacity for it to be immediately assigned to a node. If
	// set to `false` and a machine learning node with capacity to run the job
	// cannot be immediately found, the API returns an error. If set to `true`,
	// the API does not return an error; the job waits in the `starting` state
	// until sufficient machine learning node capacity is available. This
	// behavior is also affected by the cluster-wide
	// `xpack.ml.max_lazy_ml_nodes` setting.
	AllowLazyStart *bool `json:"allow_lazy_start,omitempty"`
	// Analysis The analysis configuration, which contains the information necessary to
	// perform one of the following types of analysis: classification, outlier
	// detection, or regression.
	Analysis types.DataframeAnalysisContainer `json:"analysis"`
	// AnalyzedFields Specifies `includes` and/or `excludes` patterns to select which fields
	// will be included in the analysis. The patterns specified in `excludes`
	// are applied last, therefore `excludes` takes precedence. In other words,
	// if the same field is specified in both `includes` and `excludes`, then
	// the field will not be included in the analysis. If `analyzed_fields` is
	// not set, only the relevant fields will be included. For example, all the
	// numeric fields for outlier detection.
	// The supported fields vary for each type of analysis. Outlier detection
	// requires numeric or `boolean` data to analyze. The algorithms don’t
	// support missing values therefore fields that have data types other than
	// numeric or boolean are ignored. Documents where included fields contain
	// missing values, null values, or an array are also ignored. Therefore the
	// `dest` index may contain documents that don’t have an outlier score.
	// Regression supports fields that are numeric, `boolean`, `text`,
	// `keyword`, and `ip` data types. It is also tolerant of missing values.
	// Fields that are supported are included in the analysis, other fields are
	// ignored. Documents where included fields contain an array with two or
	// more values are also ignored. Documents in the `dest` index that don’t
	// contain a results field are not included in the regression analysis.
	// Classification supports fields that are numeric, `boolean`, `text`,
	// `keyword`, and `ip` data types. It is also tolerant of missing values.
	// Fields that are supported are included in the analysis, other fields are
	// ignored. Documents where included fields contain an array with two or
	// more values are also ignored. Documents in the `dest` index that don’t
	// contain a results field are not included in the classification analysis.
	// Classification analysis can be improved by mapping ordinal variable
	// values to a single number. For example, in case of age ranges, you can
	// model the values as `0-14 = 0`, `15-24 = 1`, `25-34 = 2`, and so on.
	AnalyzedFields *types.DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	// Description A description of the job.
	Description *string `json:"description,omitempty"`
	// Dest The destination configuration.
	Dest    types.DataframeAnalyticsDestination `json:"dest"`
	Headers map[string][]string                 `json:"headers,omitempty"`
	// MaxNumThreads The maximum number of threads to be used by the analysis. Using more
	// threads may decrease the time necessary to complete the analysis at the
	// cost of using more CPU. Note that the process may use additional threads
	// for operational functionality other than the analysis itself.
	MaxNumThreads *int `json:"max_num_threads,omitempty"`
	// ModelMemoryLimit The approximate maximum amount of memory resources that are permitted for
	// analytical processing. If your `elasticsearch.yml` file contains an
	// `xpack.ml.max_model_memory_limit` setting, an error occurs when you try
	// to create data frame analytics jobs that have `model_memory_limit` values
	// greater than that setting.
	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`
	// Source The configuration of how to source the analysis data.
	Source  types.DataframeAnalyticsSource `json:"source"`
	Version *string                        `json:"version,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Putdataframeanalytics request: %w", err)
	}

	return &req, nil
}
