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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Package include
package include

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Include.ts#L20-L47
type Include struct {
	Name string
}

var (
	Definition = Include{"definition"}

	Featureimportancebaseline = Include{"feature_importance_baseline"}

	Hyperparameters = Include{"hyperparameters"}

	Totalfeatureimportance = Include{"total_feature_importance"}

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
