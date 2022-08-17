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

// SignificantTextAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L356-L374
type SignificantTextAggregation struct {
	BackgroundFilter    *QueryContainer                                              `json:"background_filter,omitempty"`
	ChiSquare           *ChiSquareHeuristic                                          `json:"chi_square,omitempty"`
	Exclude             *TermsExclude                                                `json:"exclude,omitempty"`
	ExecutionHint       *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field               *Field                                                       `json:"field,omitempty"`
	FilterDuplicateText *bool                                                        `json:"filter_duplicate_text,omitempty"`
	Gnd                 *GoogleNormalizedDistanceHeuristic                           `json:"gnd,omitempty"`
	Include             []string                                                     `json:"include,omitempty"`
	Jlh                 *EmptyObject                                                 `json:"jlh,omitempty"`
	Meta                *Metadata                                                    `json:"meta,omitempty"`
	MinDocCount         *int64                                                       `json:"min_doc_count,omitempty"`
	MutualInformation   *MutualInformationHeuristic                                  `json:"mutual_information,omitempty"`
	Name                *string                                                      `json:"name,omitempty"`
	Percentage          *PercentageScoreHeuristic                                    `json:"percentage,omitempty"`
	ScriptHeuristic     *ScriptedHeuristic                                           `json:"script_heuristic,omitempty"`
	ShardMinDocCount    *int64                                                       `json:"shard_min_doc_count,omitempty"`
	ShardSize           *int                                                         `json:"shard_size,omitempty"`
	Size                *int                                                         `json:"size,omitempty"`
	SourceFields        *Fields                                                      `json:"source_fields,omitempty"`
}

// SignificantTextAggregationBuilder holds SignificantTextAggregation struct and provides a builder API.
type SignificantTextAggregationBuilder struct {
	v *SignificantTextAggregation
}

// NewSignificantTextAggregation provides a builder for the SignificantTextAggregation struct.
func NewSignificantTextAggregationBuilder() *SignificantTextAggregationBuilder {
	r := SignificantTextAggregationBuilder{
		&SignificantTextAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SignificantTextAggregation struct
func (rb *SignificantTextAggregationBuilder) Build() SignificantTextAggregation {
	return *rb.v
}

func (rb *SignificantTextAggregationBuilder) BackgroundFilter(backgroundfilter *QueryContainerBuilder) *SignificantTextAggregationBuilder {
	v := backgroundfilter.Build()
	rb.v.BackgroundFilter = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) ChiSquare(chisquare *ChiSquareHeuristicBuilder) *SignificantTextAggregationBuilder {
	v := chisquare.Build()
	rb.v.ChiSquare = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) Exclude(exclude *TermsExcludeBuilder) *SignificantTextAggregationBuilder {
	v := exclude.Build()
	rb.v.Exclude = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *SignificantTextAggregationBuilder {
	rb.v.ExecutionHint = &executionhint
	return rb
}

func (rb *SignificantTextAggregationBuilder) Field(field Field) *SignificantTextAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *SignificantTextAggregationBuilder) FilterDuplicateText(filterduplicatetext bool) *SignificantTextAggregationBuilder {
	rb.v.FilterDuplicateText = &filterduplicatetext
	return rb
}

func (rb *SignificantTextAggregationBuilder) Gnd(gnd *GoogleNormalizedDistanceHeuristicBuilder) *SignificantTextAggregationBuilder {
	v := gnd.Build()
	rb.v.Gnd = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) Include(arg []string) *SignificantTextAggregationBuilder {
	rb.v.Include = arg
	return rb
}

func (rb *SignificantTextAggregationBuilder) Jlh(jlh *EmptyObjectBuilder) *SignificantTextAggregationBuilder {
	v := jlh.Build()
	rb.v.Jlh = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) Meta(meta *MetadataBuilder) *SignificantTextAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) MinDocCount(mindoccount int64) *SignificantTextAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *SignificantTextAggregationBuilder) MutualInformation(mutualinformation *MutualInformationHeuristicBuilder) *SignificantTextAggregationBuilder {
	v := mutualinformation.Build()
	rb.v.MutualInformation = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) Name(name string) *SignificantTextAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *SignificantTextAggregationBuilder) Percentage(percentage *PercentageScoreHeuristicBuilder) *SignificantTextAggregationBuilder {
	v := percentage.Build()
	rb.v.Percentage = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) ScriptHeuristic(scriptheuristic *ScriptedHeuristicBuilder) *SignificantTextAggregationBuilder {
	v := scriptheuristic.Build()
	rb.v.ScriptHeuristic = &v
	return rb
}

func (rb *SignificantTextAggregationBuilder) ShardMinDocCount(shardmindoccount int64) *SignificantTextAggregationBuilder {
	rb.v.ShardMinDocCount = &shardmindoccount
	return rb
}

func (rb *SignificantTextAggregationBuilder) ShardSize(shardsize int) *SignificantTextAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

func (rb *SignificantTextAggregationBuilder) Size(size int) *SignificantTextAggregationBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *SignificantTextAggregationBuilder) SourceFields(sourcefields *FieldsBuilder) *SignificantTextAggregationBuilder {
	v := sourcefields.Build()
	rb.v.SourceFields = &v
	return rb
}
