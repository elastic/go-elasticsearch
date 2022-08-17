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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/categorizationstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/memorystatus"
)

// ModelSizeStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Model.ts#L56-L78
type ModelSizeStats struct {
	AssignmentMemoryBasis         *string                                   `json:"assignment_memory_basis,omitempty"`
	BucketAllocationFailuresCount int64                                     `json:"bucket_allocation_failures_count"`
	CategorizationStatus          categorizationstatus.CategorizationStatus `json:"categorization_status"`
	CategorizedDocCount           int                                       `json:"categorized_doc_count"`
	DeadCategoryCount             int                                       `json:"dead_category_count"`
	FailedCategoryCount           int                                       `json:"failed_category_count"`
	FrequentCategoryCount         int                                       `json:"frequent_category_count"`
	JobId                         Id                                        `json:"job_id"`
	LogTime                       DateTime                                  `json:"log_time"`
	MemoryStatus                  memorystatus.MemoryStatus                 `json:"memory_status"`
	ModelBytes                    ByteSize                                  `json:"model_bytes"`
	ModelBytesExceeded            *ByteSize                                 `json:"model_bytes_exceeded,omitempty"`
	ModelBytesMemoryLimit         *ByteSize                                 `json:"model_bytes_memory_limit,omitempty"`
	PeakModelBytes                *ByteSize                                 `json:"peak_model_bytes,omitempty"`
	RareCategoryCount             int                                       `json:"rare_category_count"`
	ResultType                    string                                    `json:"result_type"`
	Timestamp                     *int64                                    `json:"timestamp,omitempty"`
	TotalByFieldCount             int64                                     `json:"total_by_field_count"`
	TotalCategoryCount            int                                       `json:"total_category_count"`
	TotalOverFieldCount           int64                                     `json:"total_over_field_count"`
	TotalPartitionFieldCount      int64                                     `json:"total_partition_field_count"`
}

// ModelSizeStatsBuilder holds ModelSizeStats struct and provides a builder API.
type ModelSizeStatsBuilder struct {
	v *ModelSizeStats
}

// NewModelSizeStats provides a builder for the ModelSizeStats struct.
func NewModelSizeStatsBuilder() *ModelSizeStatsBuilder {
	r := ModelSizeStatsBuilder{
		&ModelSizeStats{},
	}

	return &r
}

// Build finalize the chain and returns the ModelSizeStats struct
func (rb *ModelSizeStatsBuilder) Build() ModelSizeStats {
	return *rb.v
}

func (rb *ModelSizeStatsBuilder) AssignmentMemoryBasis(assignmentmemorybasis string) *ModelSizeStatsBuilder {
	rb.v.AssignmentMemoryBasis = &assignmentmemorybasis
	return rb
}

func (rb *ModelSizeStatsBuilder) BucketAllocationFailuresCount(bucketallocationfailurescount int64) *ModelSizeStatsBuilder {
	rb.v.BucketAllocationFailuresCount = bucketallocationfailurescount
	return rb
}

func (rb *ModelSizeStatsBuilder) CategorizationStatus(categorizationstatus categorizationstatus.CategorizationStatus) *ModelSizeStatsBuilder {
	rb.v.CategorizationStatus = categorizationstatus
	return rb
}

func (rb *ModelSizeStatsBuilder) CategorizedDocCount(categorizeddoccount int) *ModelSizeStatsBuilder {
	rb.v.CategorizedDocCount = categorizeddoccount
	return rb
}

func (rb *ModelSizeStatsBuilder) DeadCategoryCount(deadcategorycount int) *ModelSizeStatsBuilder {
	rb.v.DeadCategoryCount = deadcategorycount
	return rb
}

func (rb *ModelSizeStatsBuilder) FailedCategoryCount(failedcategorycount int) *ModelSizeStatsBuilder {
	rb.v.FailedCategoryCount = failedcategorycount
	return rb
}

func (rb *ModelSizeStatsBuilder) FrequentCategoryCount(frequentcategorycount int) *ModelSizeStatsBuilder {
	rb.v.FrequentCategoryCount = frequentcategorycount
	return rb
}

func (rb *ModelSizeStatsBuilder) JobId(jobid Id) *ModelSizeStatsBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *ModelSizeStatsBuilder) LogTime(logtime *DateTimeBuilder) *ModelSizeStatsBuilder {
	v := logtime.Build()
	rb.v.LogTime = v
	return rb
}

func (rb *ModelSizeStatsBuilder) MemoryStatus(memorystatus memorystatus.MemoryStatus) *ModelSizeStatsBuilder {
	rb.v.MemoryStatus = memorystatus
	return rb
}

func (rb *ModelSizeStatsBuilder) ModelBytes(modelbytes *ByteSizeBuilder) *ModelSizeStatsBuilder {
	v := modelbytes.Build()
	rb.v.ModelBytes = v
	return rb
}

func (rb *ModelSizeStatsBuilder) ModelBytesExceeded(modelbytesexceeded *ByteSizeBuilder) *ModelSizeStatsBuilder {
	v := modelbytesexceeded.Build()
	rb.v.ModelBytesExceeded = &v
	return rb
}

func (rb *ModelSizeStatsBuilder) ModelBytesMemoryLimit(modelbytesmemorylimit *ByteSizeBuilder) *ModelSizeStatsBuilder {
	v := modelbytesmemorylimit.Build()
	rb.v.ModelBytesMemoryLimit = &v
	return rb
}

func (rb *ModelSizeStatsBuilder) PeakModelBytes(peakmodelbytes *ByteSizeBuilder) *ModelSizeStatsBuilder {
	v := peakmodelbytes.Build()
	rb.v.PeakModelBytes = &v
	return rb
}

func (rb *ModelSizeStatsBuilder) RareCategoryCount(rarecategorycount int) *ModelSizeStatsBuilder {
	rb.v.RareCategoryCount = rarecategorycount
	return rb
}

func (rb *ModelSizeStatsBuilder) ResultType(resulttype string) *ModelSizeStatsBuilder {
	rb.v.ResultType = resulttype
	return rb
}

func (rb *ModelSizeStatsBuilder) Timestamp(timestamp int64) *ModelSizeStatsBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

func (rb *ModelSizeStatsBuilder) TotalByFieldCount(totalbyfieldcount int64) *ModelSizeStatsBuilder {
	rb.v.TotalByFieldCount = totalbyfieldcount
	return rb
}

func (rb *ModelSizeStatsBuilder) TotalCategoryCount(totalcategorycount int) *ModelSizeStatsBuilder {
	rb.v.TotalCategoryCount = totalcategorycount
	return rb
}

func (rb *ModelSizeStatsBuilder) TotalOverFieldCount(totaloverfieldcount int64) *ModelSizeStatsBuilder {
	rb.v.TotalOverFieldCount = totaloverfieldcount
	return rb
}

func (rb *ModelSizeStatsBuilder) TotalPartitionFieldCount(totalpartitionfieldcount int64) *ModelSizeStatsBuilder {
	rb.v.TotalPartitionFieldCount = totalpartitionfieldcount
	return rb
}
