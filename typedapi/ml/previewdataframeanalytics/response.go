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

package previewdataframeanalytics

// Response holds the response body struct for the package previewdataframeanalytics
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/preview_data_frame_analytics/MlPreviewDataFrameAnalyticsResponse.ts#L23-L28
type Response struct {

	// FeatureValues An array of objects that contain feature name and value pairs. The features
	// have been processed and indicate what will be sent to the model for training.
	FeatureValues []map[string]string `json:"feature_values"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
