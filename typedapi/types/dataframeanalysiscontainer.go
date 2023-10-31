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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

// DataframeAnalysisContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeAnalytics.ts#L84-L101
type DataframeAnalysisContainer struct {
	// Classification The configuration information necessary to perform classification.
	Classification *DataframeAnalysisClassification `json:"classification,omitempty"`
	// OutlierDetection The configuration information necessary to perform outlier detection. NOTE:
	// Advanced parameters are for fine-tuning classification analysis. They are set
	// automatically by hyperparameter optimization to give the minimum validation
	// error. It is highly recommended to use the default values unless you fully
	// understand the function of these parameters.
	OutlierDetection *DataframeAnalysisOutlierDetection `json:"outlier_detection,omitempty"`
	// Regression The configuration information necessary to perform regression. NOTE: Advanced
	// parameters are for fine-tuning regression analysis. They are set
	// automatically by hyperparameter optimization to give the minimum validation
	// error. It is highly recommended to use the default values unless you fully
	// understand the function of these parameters.
	Regression *DataframeAnalysisRegression `json:"regression,omitempty"`
}

// NewDataframeAnalysisContainer returns a DataframeAnalysisContainer.
func NewDataframeAnalysisContainer() *DataframeAnalysisContainer {
	r := &DataframeAnalysisContainer{}

	return r
}
