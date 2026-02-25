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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package include
package include

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/ml/_types/Include.ts#L20-L47
type Include struct {
	Name string
}

var (

	// Definition Includes the model definition.
	Definition = Include{"definition"}

	// Featureimportancebaseline Includes the baseline for feature importance values.
	Featureimportancebaseline = Include{"feature_importance_baseline"}

	// Hyperparameters Includes the information about hyperparameters used to train the model. This
	// information consists of the value, the absolute and relative importance of
	// the hyperparameter as well as an indicator of whether it was specified by the
	// user or tuned during hyperparameter optimization.
	Hyperparameters = Include{"hyperparameters"}

	// Totalfeatureimportance Includes the total feature importance for the training data set. The baseline
	// and total feature importance values are returned in the metadata field in the
	// response body.
	Totalfeatureimportance = Include{"total_feature_importance"}

	// Definitionstatus Includes the model definition status.
	Definitionstatus = Include{"definition_status"}
)

func (i Include) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *Include) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "definition":
		*i = Definition
	case "feature_importance_baseline":
		*i = Featureimportancebaseline
	case "hyperparameters":
		*i = Hyperparameters
	case "total_feature_importance":
		*i = Totalfeatureimportance
	case "definition_status":
		*i = Definitionstatus
	default:
		*i = Include{string(text)}
	}

	return nil
}

func (i Include) String() string {
	return i.Name
}
