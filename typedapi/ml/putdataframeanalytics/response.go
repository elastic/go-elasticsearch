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
// https://github.com/elastic/elasticsearch-specification/tree/50c316c036cf0c3f567011c2bc24e7d2e1b8c781

package putdataframeanalytics

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package putdataframeanalytics
//
// https://github.com/elastic/elasticsearch-specification/blob/50c316c036cf0c3f567011c2bc24e7d2e1b8c781/specification/ml/put_data_frame_analytics/MlPutDataFrameAnalyticsResponse.ts#L31-L46
type Response struct {
	AllowLazyStart   bool                                   `json:"allow_lazy_start"`
	Analysis         types.DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields   *types.DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	Authorization    *types.DataframeAnalyticsAuthorization `json:"authorization,omitempty"`
	CreateTime       int64                                  `json:"create_time"`
	Description      *string                                `json:"description,omitempty"`
	Dest             types.DataframeAnalyticsDestination    `json:"dest"`
	Id               string                                 `json:"id"`
	MaxNumThreads    int                                    `json:"max_num_threads"`
	ModelMemoryLimit string                                 `json:"model_memory_limit"`
	Source           types.DataframeAnalyticsSource         `json:"source"`
	Version          string                                 `json:"version"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
