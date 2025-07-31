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

package evaluatedataframe

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package evaluatedataframe
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/evaluate_data_frame/MlEvaluateDataFrameResponse.ts#L26-L44
type Response struct {

	// Classification Evaluation results for a classification analysis.
	// It outputs a prediction that identifies to which of the classes each document
	// belongs.
	Classification *types.DataframeClassificationSummary `json:"classification,omitempty"`
	// OutlierDetection Evaluation results for an outlier detection analysis.
	// It outputs the probability that each document is an outlier.
	OutlierDetection *types.DataframeOutlierDetectionSummary `json:"outlier_detection,omitempty"`
	// Regression Evaluation results for a regression analysis which outputs a prediction of
	// values.
	Regression *types.DataframeRegressionSummary `json:"regression,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
