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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

// SignificantTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L338-L354
type SignificantTermsAggregation struct {
	BackgroundFilter  *QueryContainer                                              `json:"background_filter,omitempty"`
	ChiSquare         *ChiSquareHeuristic                                          `json:"chi_square,omitempty"`
	Exclude           *TermsExclude                                                `json:"exclude,omitempty"`
	ExecutionHint     *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field             *Field                                                       `json:"field,omitempty"`
	Gnd               *GoogleNormalizedDistanceHeuristic                           `json:"gnd,omitempty"`
	Include           *TermsInclude                                                `json:"include,omitempty"`
	Jlh               *EmptyObject                                                 `json:"jlh,omitempty"`
	Meta              *Metadata                                                    `json:"meta,omitempty"`
	MinDocCount       *int64                                                       `json:"min_doc_count,omitempty"`
	MutualInformation *MutualInformationHeuristic                                  `json:"mutual_information,omitempty"`
	Name              *string                                                      `json:"name,omitempty"`
	Percentage        *PercentageScoreHeuristic                                    `json:"percentage,omitempty"`
	ScriptHeuristic   *ScriptedHeuristic                                           `json:"script_heuristic,omitempty"`
	ShardMinDocCount  *int64                                                       `json:"shard_min_doc_count,omitempty"`
	ShardSize         *int                                                         `json:"shard_size,omitempty"`
	Size              *int                                                         `json:"size,omitempty"`
}

// SignificantTermsAggregationBuilder holds SignificantTermsAggregation struct and provides a builder API.
type SignificantTermsAggregationBuilder struct {
	v *SignificantTermsAggregation
}

// NewSignificantTermsAggregation provides a builder for the SignificantTermsAggregation struct.
func NewSignificantTermsAggregationBuilder() *SignificantTermsAggregationBuilder {
	r := SignificantTermsAggregationBuilder{
		&SignificantTermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SignificantTermsAggregation struct
func (rb *SignificantTermsAggregationBuilder) Build() SignificantTermsAggregation {
	return *rb.v
}

func (rb *SignificantTermsAggregationBuilder) BackgroundFilter(backgroundfilter *QueryContainerBuilder) *SignificantTermsAggregationBuilder {
	v := backgroundfilter.Build()
	rb.v.BackgroundFilter = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) ChiSquare(chisquare *ChiSquareHeuristicBuilder) *SignificantTermsAggregationBuilder {
	v := chisquare.Build()
	rb.v.ChiSquare = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Exclude(exclude *TermsExcludeBuilder) *SignificantTermsAggregationBuilder {
	v := exclude.Build()
	rb.v.Exclude = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *SignificantTermsAggregationBuilder {
	rb.v.ExecutionHint = &executionhint
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Field(field Field) *SignificantTermsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Gnd(gnd *GoogleNormalizedDistanceHeuristicBuilder) *SignificantTermsAggregationBuilder {
	v := gnd.Build()
	rb.v.Gnd = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Include(include *TermsIncludeBuilder) *SignificantTermsAggregationBuilder {
	v := include.Build()
	rb.v.Include = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Jlh(jlh *EmptyObjectBuilder) *SignificantTermsAggregationBuilder {
	v := jlh.Build()
	rb.v.Jlh = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Meta(meta *MetadataBuilder) *SignificantTermsAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) MinDocCount(mindoccount int64) *SignificantTermsAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *SignificantTermsAggregationBuilder) MutualInformation(mutualinformation *MutualInformationHeuristicBuilder) *SignificantTermsAggregationBuilder {
	v := mutualinformation.Build()
	rb.v.MutualInformation = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Name(name string) *SignificantTermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Percentage(percentage *PercentageScoreHeuristicBuilder) *SignificantTermsAggregationBuilder {
	v := percentage.Build()
	rb.v.Percentage = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) ScriptHeuristic(scriptheuristic *ScriptedHeuristicBuilder) *SignificantTermsAggregationBuilder {
	v := scriptheuristic.Build()
	rb.v.ScriptHeuristic = &v
	return rb
}

func (rb *SignificantTermsAggregationBuilder) ShardMinDocCount(shardmindoccount int64) *SignificantTermsAggregationBuilder {
	rb.v.ShardMinDocCount = &shardmindoccount
	return rb
}

func (rb *SignificantTermsAggregationBuilder) ShardSize(shardsize int) *SignificantTermsAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

func (rb *SignificantTermsAggregationBuilder) Size(size int) *SignificantTermsAggregationBuilder {
	rb.v.Size = &size
	return rb
}
