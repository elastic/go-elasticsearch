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

// Category type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Category.ts#L23-L49
type Category struct {
	// CategoryId A unique identifier for the category. category_id is unique at the job level,
	// even when per-partition categorization is enabled.
	CategoryId uint64 `json:"category_id"`
	// Examples A list of examples of actual values that matched the category.
	Examples []string `json:"examples"`
	// GrokPattern [experimental] A Grok pattern that could be used in Logstash or an ingest
	// pipeline to extract fields from messages that match the category. This field
	// is experimental and may be changed or removed in a future release. The Grok
	// patterns that are found are not optimal, but are often a good starting point
	// for manual tweaking.
	GrokPattern *string `json:"grok_pattern,omitempty"`
	// JobId Identifier for the anomaly detection job.
	JobId Id `json:"job_id"`
	// MaxMatchingLength The maximum length of the fields that matched the category. The value is
	// increased by 10% to enable matching for similar fields that have not been
	// analyzed.
	MaxMatchingLength uint64 `json:"max_matching_length"`
	Mlcategory        string `json:"mlcategory"`
	// NumMatches The number of messages that have been matched by this category. This is only
	// guaranteed to have the latest accurate count after a job _flush or _close
	NumMatches *int64  `json:"num_matches,omitempty"`
	P          *string `json:"p,omitempty"`
	// PartitionFieldName If per-partition categorization is enabled, this property identifies the
	// field used to segment the categorization. It is not present when
	// per-partition categorization is disabled.
	PartitionFieldName *string `json:"partition_field_name,omitempty"`
	// PartitionFieldValue If per-partition categorization is enabled, this property identifies the
	// value of the partition_field_name for the category. It is not present when
	// per-partition categorization is disabled.
	PartitionFieldValue *string `json:"partition_field_value,omitempty"`
	// PreferredToCategories A list of category_id entries that this current category encompasses. Any new
	// message that is processed by the categorizer will match against this category
	// and not any of the categories in this list. This is only guaranteed to have
	// the latest accurate list of categories after a job _flush or _close
	PreferredToCategories []Id `json:"preferred_to_categories,omitempty"`
	// Regex A regular expression that is used to search for values that match the
	// category.
	Regex      string `json:"regex"`
	ResultType string `json:"result_type"`
	// Terms A space separated list of the common tokens that are matched in values of the
	// category.
	Terms string `json:"terms"`
}

// CategoryBuilder holds Category struct and provides a builder API.
type CategoryBuilder struct {
	v *Category
}

// NewCategory provides a builder for the Category struct.
func NewCategoryBuilder() *CategoryBuilder {
	r := CategoryBuilder{
		&Category{},
	}

	return &r
}

// Build finalize the chain and returns the Category struct
func (rb *CategoryBuilder) Build() Category {
	return *rb.v
}

// CategoryId A unique identifier for the category. category_id is unique at the job level,
// even when per-partition categorization is enabled.

func (rb *CategoryBuilder) CategoryId(categoryid uint64) *CategoryBuilder {
	rb.v.CategoryId = categoryid
	return rb
}

// Examples A list of examples of actual values that matched the category.

func (rb *CategoryBuilder) Examples(examples ...string) *CategoryBuilder {
	rb.v.Examples = examples
	return rb
}

// GrokPattern [experimental] A Grok pattern that could be used in Logstash or an ingest
// pipeline to extract fields from messages that match the category. This field
// is experimental and may be changed or removed in a future release. The Grok
// patterns that are found are not optimal, but are often a good starting point
// for manual tweaking.

func (rb *CategoryBuilder) GrokPattern(grokpattern string) *CategoryBuilder {
	rb.v.GrokPattern = &grokpattern
	return rb
}

// JobId Identifier for the anomaly detection job.

func (rb *CategoryBuilder) JobId(jobid Id) *CategoryBuilder {
	rb.v.JobId = jobid
	return rb
}

// MaxMatchingLength The maximum length of the fields that matched the category. The value is
// increased by 10% to enable matching for similar fields that have not been
// analyzed.

func (rb *CategoryBuilder) MaxMatchingLength(maxmatchinglength uint64) *CategoryBuilder {
	rb.v.MaxMatchingLength = maxmatchinglength
	return rb
}

func (rb *CategoryBuilder) Mlcategory(mlcategory string) *CategoryBuilder {
	rb.v.Mlcategory = mlcategory
	return rb
}

// NumMatches The number of messages that have been matched by this category. This is only
// guaranteed to have the latest accurate count after a job _flush or _close

func (rb *CategoryBuilder) NumMatches(nummatches int64) *CategoryBuilder {
	rb.v.NumMatches = &nummatches
	return rb
}

func (rb *CategoryBuilder) P(p string) *CategoryBuilder {
	rb.v.P = &p
	return rb
}

// PartitionFieldName If per-partition categorization is enabled, this property identifies the
// field used to segment the categorization. It is not present when
// per-partition categorization is disabled.

func (rb *CategoryBuilder) PartitionFieldName(partitionfieldname string) *CategoryBuilder {
	rb.v.PartitionFieldName = &partitionfieldname
	return rb
}

// PartitionFieldValue If per-partition categorization is enabled, this property identifies the
// value of the partition_field_name for the category. It is not present when
// per-partition categorization is disabled.

func (rb *CategoryBuilder) PartitionFieldValue(partitionfieldvalue string) *CategoryBuilder {
	rb.v.PartitionFieldValue = &partitionfieldvalue
	return rb
}

// PreferredToCategories A list of category_id entries that this current category encompasses. Any new
// message that is processed by the categorizer will match against this category
// and not any of the categories in this list. This is only guaranteed to have
// the latest accurate list of categories after a job _flush or _close

func (rb *CategoryBuilder) PreferredToCategories(preferred_to_categories ...Id) *CategoryBuilder {
	rb.v.PreferredToCategories = preferred_to_categories
	return rb
}

// Regex A regular expression that is used to search for values that match the
// category.

func (rb *CategoryBuilder) Regex(regex string) *CategoryBuilder {
	rb.v.Regex = regex
	return rb
}

func (rb *CategoryBuilder) ResultType(resulttype string) *CategoryBuilder {
	rb.v.ResultType = resulttype
	return rb
}

// Terms A space separated list of the common tokens that are matched in values of the
// category.

func (rb *CategoryBuilder) Terms(terms string) *CategoryBuilder {
	rb.v.Terms = terms
	return rb
}
