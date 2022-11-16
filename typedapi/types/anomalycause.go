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
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// AnomalyCause type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/ml/_types/Anomaly.ts#L49-L64
type AnomalyCause struct {
	Actual                 []float64   `json:"actual"`
	ByFieldName            string      `json:"by_field_name"`
	ByFieldValue           string      `json:"by_field_value"`
	CorrelatedByFieldValue string      `json:"correlated_by_field_value"`
	FieldName              string      `json:"field_name"`
	Function               string      `json:"function"`
	FunctionDescription    string      `json:"function_description"`
	Influencers            []Influence `json:"influencers"`
	OverFieldName          string      `json:"over_field_name"`
	OverFieldValue         string      `json:"over_field_value"`
	PartitionFieldName     string      `json:"partition_field_name"`
	PartitionFieldValue    string      `json:"partition_field_value"`
	Probability            float64     `json:"probability"`
	Typical                []float64   `json:"typical"`
}

// NewAnomalyCause returns a AnomalyCause.
func NewAnomalyCause() *AnomalyCause {
	r := &AnomalyCause{}

	return r
}
