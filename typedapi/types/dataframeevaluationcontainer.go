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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"encoding/json"
	"fmt"
)

// DataframeEvaluationContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/ml/_types/DataframeEvaluation.ts#L25-L33
type DataframeEvaluationContainer struct {
	AdditionalDataframeEvaluationContainerProperty map[string]json.RawMessage `json:"-"`
	// Classification Classification evaluation evaluates the results of a classification analysis
	// which outputs a prediction that identifies to which of the classes each
	// document belongs.
	Classification *DataframeEvaluationClassification `json:"classification,omitempty"`
	// OutlierDetection Outlier detection evaluates the results of an outlier detection analysis
	// which outputs the probability that each document is an outlier.
	OutlierDetection *DataframeEvaluationOutlierDetection `json:"outlier_detection,omitempty"`
	// Regression Regression evaluation evaluates the results of a regression analysis which
	// outputs a prediction of values.
	Regression *DataframeEvaluationRegression `json:"regression,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DataframeEvaluationContainer) MarshalJSON() ([]byte, error) {
	type opt DataframeEvaluationContainer
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
	for key, value := range s.AdditionalDataframeEvaluationContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalDataframeEvaluationContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewDataframeEvaluationContainer returns a DataframeEvaluationContainer.
func NewDataframeEvaluationContainer() *DataframeEvaluationContainer {
	r := &DataframeEvaluationContainer{
		AdditionalDataframeEvaluationContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}

// true

type DataframeEvaluationContainerVariant interface {
	DataframeEvaluationContainerCaster() *DataframeEvaluationContainer
}

func (s *DataframeEvaluationContainer) DataframeEvaluationContainerCaster() *DataframeEvaluationContainer {
	return s
}
