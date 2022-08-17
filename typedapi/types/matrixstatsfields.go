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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// MatrixStatsFields type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L714-L723
type MatrixStatsFields struct {
	Correlation map[Field]float64 `json:"correlation"`
	Count       int64             `json:"count"`
	Covariance  map[Field]float64 `json:"covariance"`
	Kurtosis    float64           `json:"kurtosis"`
	Mean        float64           `json:"mean"`
	Name        Field             `json:"name"`
	Skewness    float64           `json:"skewness"`
	Variance    float64           `json:"variance"`
}

// MatrixStatsFieldsBuilder holds MatrixStatsFields struct and provides a builder API.
type MatrixStatsFieldsBuilder struct {
	v *MatrixStatsFields
}

// NewMatrixStatsFields provides a builder for the MatrixStatsFields struct.
func NewMatrixStatsFieldsBuilder() *MatrixStatsFieldsBuilder {
	r := MatrixStatsFieldsBuilder{
		&MatrixStatsFields{
			Correlation: make(map[Field]float64, 0),
			Covariance:  make(map[Field]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MatrixStatsFields struct
func (rb *MatrixStatsFieldsBuilder) Build() MatrixStatsFields {
	return *rb.v
}

func (rb *MatrixStatsFieldsBuilder) Correlation(value map[Field]float64) *MatrixStatsFieldsBuilder {
	rb.v.Correlation = value
	return rb
}

func (rb *MatrixStatsFieldsBuilder) Count(count int64) *MatrixStatsFieldsBuilder {
	rb.v.Count = count
	return rb
}

func (rb *MatrixStatsFieldsBuilder) Covariance(value map[Field]float64) *MatrixStatsFieldsBuilder {
	rb.v.Covariance = value
	return rb
}

func (rb *MatrixStatsFieldsBuilder) Kurtosis(kurtosis float64) *MatrixStatsFieldsBuilder {
	rb.v.Kurtosis = kurtosis
	return rb
}

func (rb *MatrixStatsFieldsBuilder) Mean(mean float64) *MatrixStatsFieldsBuilder {
	rb.v.Mean = mean
	return rb
}

func (rb *MatrixStatsFieldsBuilder) Name(name Field) *MatrixStatsFieldsBuilder {
	rb.v.Name = name
	return rb
}

func (rb *MatrixStatsFieldsBuilder) Skewness(skewness float64) *MatrixStatsFieldsBuilder {
	rb.v.Skewness = skewness
	return rb
}

func (rb *MatrixStatsFieldsBuilder) Variance(variance float64) *MatrixStatsFieldsBuilder {
	rb.v.Variance = variance
	return rb
}
