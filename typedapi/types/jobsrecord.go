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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/categorizationstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jobstate"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/memorystatus"
)

// JobsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/ml_jobs/types.ts#L24-L325
type JobsRecord struct {
	// AssignmentExplanation why the job is or is not assigned to a node
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// BucketsCount bucket count
	BucketsCount *string `json:"buckets.count,omitempty"`
	// BucketsTimeExpAvg exponential average bucket processing time (milliseconds)
	BucketsTimeExpAvg *string `json:"buckets.time.exp_avg,omitempty"`
	// BucketsTimeExpAvgHour exponential average bucket processing time by hour (milliseconds)
	BucketsTimeExpAvgHour *string `json:"buckets.time.exp_avg_hour,omitempty"`
	// BucketsTimeMax maximum bucket processing time
	BucketsTimeMax *string `json:"buckets.time.max,omitempty"`
	// BucketsTimeMin minimum bucket processing time
	BucketsTimeMin *string `json:"buckets.time.min,omitempty"`
	// BucketsTimeTotal total bucket processing time
	BucketsTimeTotal *string `json:"buckets.time.total,omitempty"`
	// DataBuckets total bucket count
	DataBuckets *string `json:"data.buckets,omitempty"`
	// DataEarliestRecord earliest record time
	DataEarliestRecord *string `json:"data.earliest_record,omitempty"`
	// DataEmptyBuckets number of empty buckets
	DataEmptyBuckets *string `json:"data.empty_buckets,omitempty"`
	// DataInputBytes total input bytes
	DataInputBytes *ByteSize `json:"data.input_bytes,omitempty"`
	// DataInputFields total field count
	DataInputFields *string `json:"data.input_fields,omitempty"`
	// DataInputRecords total record count
	DataInputRecords *string `json:"data.input_records,omitempty"`
	// DataInvalidDates number of records with invalid dates
	DataInvalidDates *string `json:"data.invalid_dates,omitempty"`
	// DataLast last time data was seen
	DataLast *string `json:"data.last,omitempty"`
	// DataLastEmptyBucket last time an empty bucket occurred
	DataLastEmptyBucket *string `json:"data.last_empty_bucket,omitempty"`
	// DataLastSparseBucket last time a sparse bucket occurred
	DataLastSparseBucket *string `json:"data.last_sparse_bucket,omitempty"`
	// DataLatestRecord latest record time
	DataLatestRecord *string `json:"data.latest_record,omitempty"`
	// DataMissingFields number of records with missing fields
	DataMissingFields *string `json:"data.missing_fields,omitempty"`
	// DataOutOfOrderTimestamps number of records handled out of order
	DataOutOfOrderTimestamps *string `json:"data.out_of_order_timestamps,omitempty"`
	// DataProcessedFields number of processed fields
	DataProcessedFields *string `json:"data.processed_fields,omitempty"`
	// DataProcessedRecords number of processed records
	DataProcessedRecords *string `json:"data.processed_records,omitempty"`
	// DataSparseBuckets number of sparse buckets
	DataSparseBuckets *string `json:"data.sparse_buckets,omitempty"`
	// ForecastsMemoryAvg average memory used by forecasts
	ForecastsMemoryAvg *string `json:"forecasts.memory.avg,omitempty"`
	// ForecastsMemoryMax maximum memory used by forecasts
	ForecastsMemoryMax *string `json:"forecasts.memory.max,omitempty"`
	// ForecastsMemoryMin minimum memory used by forecasts
	ForecastsMemoryMin *string `json:"forecasts.memory.min,omitempty"`
	// ForecastsMemoryTotal total memory used by all forecasts
	ForecastsMemoryTotal *string `json:"forecasts.memory.total,omitempty"`
	// ForecastsRecordsAvg average record count for forecasts
	ForecastsRecordsAvg *string `json:"forecasts.records.avg,omitempty"`
	// ForecastsRecordsMax maximum record count for forecasts
	ForecastsRecordsMax *string `json:"forecasts.records.max,omitempty"`
	// ForecastsRecordsMin minimum record count for forecasts
	ForecastsRecordsMin *string `json:"forecasts.records.min,omitempty"`
	// ForecastsRecordsTotal total record count for all forecasts
	ForecastsRecordsTotal *string `json:"forecasts.records.total,omitempty"`
	// ForecastsTimeAvg average runtime for all forecasts (milliseconds)
	ForecastsTimeAvg *string `json:"forecasts.time.avg,omitempty"`
	// ForecastsTimeMax maximum run time for forecasts
	ForecastsTimeMax *string `json:"forecasts.time.max,omitempty"`
	// ForecastsTimeMin minimum runtime for forecasts
	ForecastsTimeMin *string `json:"forecasts.time.min,omitempty"`
	// ForecastsTimeTotal total runtime for all forecasts
	ForecastsTimeTotal *string `json:"forecasts.time.total,omitempty"`
	// ForecastsTotal total number of forecasts
	ForecastsTotal *string `json:"forecasts.total,omitempty"`
	// Id the job_id
	Id *string `json:"id,omitempty"`
	// ModelBucketAllocationFailures number of bucket allocation failures
	ModelBucketAllocationFailures *string `json:"model.bucket_allocation_failures,omitempty"`
	// ModelByFields count of 'by' fields
	ModelByFields *string `json:"model.by_fields,omitempty"`
	// ModelBytes model size
	ModelBytes *ByteSize `json:"model.bytes,omitempty"`
	// ModelBytesExceeded how much the model has exceeded the limit
	ModelBytesExceeded *ByteSize `json:"model.bytes_exceeded,omitempty"`
	// ModelCategorizationStatus current categorization status
	ModelCategorizationStatus *categorizationstatus.CategorizationStatus `json:"model.categorization_status,omitempty"`
	// ModelCategorizedDocCount count of categorized documents
	ModelCategorizedDocCount *string `json:"model.categorized_doc_count,omitempty"`
	// ModelDeadCategoryCount count of dead categories
	ModelDeadCategoryCount *string `json:"model.dead_category_count,omitempty"`
	// ModelFailedCategoryCount count of failed categories
	ModelFailedCategoryCount *string `json:"model.failed_category_count,omitempty"`
	// ModelFrequentCategoryCount count of frequent categories
	ModelFrequentCategoryCount *string `json:"model.frequent_category_count,omitempty"`
	// ModelLogTime when the model stats were gathered
	ModelLogTime *string `json:"model.log_time,omitempty"`
	// ModelMemoryLimit model memory limit
	ModelMemoryLimit *string `json:"model.memory_limit,omitempty"`
	// ModelMemoryStatus current memory status
	ModelMemoryStatus *memorystatus.MemoryStatus `json:"model.memory_status,omitempty"`
	// ModelOverFields count of 'over' fields
	ModelOverFields *string `json:"model.over_fields,omitempty"`
	// ModelPartitionFields count of 'partition' fields
	ModelPartitionFields *string `json:"model.partition_fields,omitempty"`
	// ModelRareCategoryCount count of rare categories
	ModelRareCategoryCount *string `json:"model.rare_category_count,omitempty"`
	// ModelTimestamp the time of the last record when the model stats were gathered
	ModelTimestamp *string `json:"model.timestamp,omitempty"`
	// ModelTotalCategoryCount count of categories
	ModelTotalCategoryCount *string `json:"model.total_category_count,omitempty"`
	// NodeAddress network address of the assigned node
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId ephemeral id of the assigned node
	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`
	// NodeId id of the assigned node
	NodeId *string `json:"node.id,omitempty"`
	// NodeName name of the assigned node
	NodeName *string `json:"node.name,omitempty"`
	// OpenedTime the amount of time the job has been opened
	OpenedTime *string `json:"opened_time,omitempty"`
	// State the job state
	State *jobstate.JobState `json:"state,omitempty"`
}

// NewJobsRecord returns a JobsRecord.
func NewJobsRecord() *JobsRecord {
	r := &JobsRecord{}

	return r
}
