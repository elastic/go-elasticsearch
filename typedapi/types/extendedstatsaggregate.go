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

// ExtendedStatsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L267-L283
type ExtendedStatsAggregate struct {
	Avg                        float64                          `json:"avg,omitempty"`
	AvgAsString                *string                          `json:"avg_as_string,omitempty"`
	Count                      int64                            `json:"count"`
	Max                        float64                          `json:"max,omitempty"`
	MaxAsString                *string                          `json:"max_as_string,omitempty"`
	Meta                       *Metadata                        `json:"meta,omitempty"`
	Min                        float64                          `json:"min,omitempty"`
	MinAsString                *string                          `json:"min_as_string,omitempty"`
	StdDeviation               float64                          `json:"std_deviation,omitempty"`
	StdDeviationAsString       *string                          `json:"std_deviation_as_string,omitempty"`
	StdDeviationBounds         *StandardDeviationBounds         `json:"std_deviation_bounds,omitempty"`
	StdDeviationBoundsAsString *StandardDeviationBoundsAsString `json:"std_deviation_bounds_as_string,omitempty"`
	Sum                        float64                          `json:"sum"`
	SumAsString                *string                          `json:"sum_as_string,omitempty"`
	SumOfSquares               float64                          `json:"sum_of_squares,omitempty"`
	SumOfSquaresAsString       *string                          `json:"sum_of_squares_as_string,omitempty"`
	Variance                   float64                          `json:"variance,omitempty"`
	VarianceAsString           *string                          `json:"variance_as_string,omitempty"`
	VariancePopulation         float64                          `json:"variance_population,omitempty"`
	VariancePopulationAsString *string                          `json:"variance_population_as_string,omitempty"`
	VarianceSampling           float64                          `json:"variance_sampling,omitempty"`
	VarianceSamplingAsString   *string                          `json:"variance_sampling_as_string,omitempty"`
}

// ExtendedStatsAggregateBuilder holds ExtendedStatsAggregate struct and provides a builder API.
type ExtendedStatsAggregateBuilder struct {
	v *ExtendedStatsAggregate
}

// NewExtendedStatsAggregate provides a builder for the ExtendedStatsAggregate struct.
func NewExtendedStatsAggregateBuilder() *ExtendedStatsAggregateBuilder {
	r := ExtendedStatsAggregateBuilder{
		&ExtendedStatsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the ExtendedStatsAggregate struct
func (rb *ExtendedStatsAggregateBuilder) Build() ExtendedStatsAggregate {
	return *rb.v
}

func (rb *ExtendedStatsAggregateBuilder) Avg(avg float64) *ExtendedStatsAggregateBuilder {
	rb.v.Avg = avg
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) AvgAsString(avgasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.AvgAsString = &avgasstring
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) Count(count int64) *ExtendedStatsAggregateBuilder {
	rb.v.Count = count
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) Max(max float64) *ExtendedStatsAggregateBuilder {
	rb.v.Max = max
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) MaxAsString(maxasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.MaxAsString = &maxasstring
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) Meta(meta *MetadataBuilder) *ExtendedStatsAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) Min(min float64) *ExtendedStatsAggregateBuilder {
	rb.v.Min = min
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) MinAsString(minasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.MinAsString = &minasstring
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) StdDeviation(stddeviation float64) *ExtendedStatsAggregateBuilder {
	rb.v.StdDeviation = stddeviation
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) StdDeviationAsString(stddeviationasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.StdDeviationAsString = &stddeviationasstring
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) StdDeviationBounds(stddeviationbounds *StandardDeviationBoundsBuilder) *ExtendedStatsAggregateBuilder {
	v := stddeviationbounds.Build()
	rb.v.StdDeviationBounds = &v
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) StdDeviationBoundsAsString(stddeviationboundsasstring *StandardDeviationBoundsAsStringBuilder) *ExtendedStatsAggregateBuilder {
	v := stddeviationboundsasstring.Build()
	rb.v.StdDeviationBoundsAsString = &v
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) Sum(sum float64) *ExtendedStatsAggregateBuilder {
	rb.v.Sum = sum
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) SumAsString(sumasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.SumAsString = &sumasstring
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) SumOfSquares(sumofsquares float64) *ExtendedStatsAggregateBuilder {
	rb.v.SumOfSquares = sumofsquares
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) SumOfSquaresAsString(sumofsquaresasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.SumOfSquaresAsString = &sumofsquaresasstring
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) Variance(variance float64) *ExtendedStatsAggregateBuilder {
	rb.v.Variance = variance
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) VarianceAsString(varianceasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.VarianceAsString = &varianceasstring
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) VariancePopulation(variancepopulation float64) *ExtendedStatsAggregateBuilder {
	rb.v.VariancePopulation = variancepopulation
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) VariancePopulationAsString(variancepopulationasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.VariancePopulationAsString = &variancepopulationasstring
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) VarianceSampling(variancesampling float64) *ExtendedStatsAggregateBuilder {
	rb.v.VarianceSampling = variancesampling
	return rb
}

func (rb *ExtendedStatsAggregateBuilder) VarianceSamplingAsString(variancesamplingasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.VarianceSamplingAsString = &variancesamplingasstring
	return rb
}
