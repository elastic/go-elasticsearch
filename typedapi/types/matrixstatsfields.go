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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// MatrixStatsFields type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/aggregations/Aggregate.ts#L754-L763
type MatrixStatsFields struct {
	Correlation map[string]float64 `json:"correlation"`
	Count       int64              `json:"count"`
	Covariance  map[string]float64 `json:"covariance"`
	Kurtosis    float64            `json:"kurtosis"`
	Mean        float64            `json:"mean"`
	Name        string             `json:"name"`
	Skewness    float64            `json:"skewness"`
	Variance    float64            `json:"variance"`
}

// NewMatrixStatsFields returns a MatrixStatsFields.
func NewMatrixStatsFields() *MatrixStatsFields {
	r := &MatrixStatsFields{
		Correlation: make(map[string]float64, 0),
		Covariance:  make(map[string]float64, 0),
	}

	return r
}
