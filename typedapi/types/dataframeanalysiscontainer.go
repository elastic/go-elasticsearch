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

import (
	"encoding/json"
	"fmt"
)

// DataframeAnalysisContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/DataframeAnalytics.ts#L84-L101
type DataframeAnalysisContainer struct {
	AdditionalDataframeAnalysisContainerProperty map[string]json.RawMessage `json:"-"`
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

// MarhsalJSON overrides marshalling for types with additional properties
func (s DataframeAnalysisContainer) MarshalJSON() ([]byte, error) {
	type opt DataframeAnalysisContainer
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalDataframeAnalysisContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalDataframeAnalysisContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewDataframeAnalysisContainer returns a DataframeAnalysisContainer.
func NewDataframeAnalysisContainer() *DataframeAnalysisContainer {
	r := &DataframeAnalysisContainer{
		AdditionalDataframeAnalysisContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}
