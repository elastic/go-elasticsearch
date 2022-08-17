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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jobstate"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/memorystatus"
)

// JobsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/ml_jobs/types.ts#L24-L325
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
	Id *Id `json:"id,omitempty"`
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
	NodeEphemeralId *NodeId `json:"node.ephemeral_id,omitempty"`
	// NodeId id of the assigned node
	NodeId *NodeId `json:"node.id,omitempty"`
	// NodeName name of the assigned node
	NodeName *string `json:"node.name,omitempty"`
	// OpenedTime the amount of time the job has been opened
	OpenedTime *string `json:"opened_time,omitempty"`
	// State the job state
	State *jobstate.JobState `json:"state,omitempty"`
}

// JobsRecordBuilder holds JobsRecord struct and provides a builder API.
type JobsRecordBuilder struct {
	v *JobsRecord
}

// NewJobsRecord provides a builder for the JobsRecord struct.
func NewJobsRecordBuilder() *JobsRecordBuilder {
	r := JobsRecordBuilder{
		&JobsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the JobsRecord struct
func (rb *JobsRecordBuilder) Build() JobsRecord {
	return *rb.v
}

// AssignmentExplanation why the job is or is not assigned to a node

func (rb *JobsRecordBuilder) AssignmentExplanation(assignmentexplanation string) *JobsRecordBuilder {
	rb.v.AssignmentExplanation = &assignmentexplanation
	return rb
}

// BucketsCount bucket count

func (rb *JobsRecordBuilder) BucketsCount(bucketscount string) *JobsRecordBuilder {
	rb.v.BucketsCount = &bucketscount
	return rb
}

// BucketsTimeExpAvg exponential average bucket processing time (milliseconds)

func (rb *JobsRecordBuilder) BucketsTimeExpAvg(bucketstimeexpavg string) *JobsRecordBuilder {
	rb.v.BucketsTimeExpAvg = &bucketstimeexpavg
	return rb
}

// BucketsTimeExpAvgHour exponential average bucket processing time by hour (milliseconds)

func (rb *JobsRecordBuilder) BucketsTimeExpAvgHour(bucketstimeexpavghour string) *JobsRecordBuilder {
	rb.v.BucketsTimeExpAvgHour = &bucketstimeexpavghour
	return rb
}

// BucketsTimeMax maximum bucket processing time

func (rb *JobsRecordBuilder) BucketsTimeMax(bucketstimemax string) *JobsRecordBuilder {
	rb.v.BucketsTimeMax = &bucketstimemax
	return rb
}

// BucketsTimeMin minimum bucket processing time

func (rb *JobsRecordBuilder) BucketsTimeMin(bucketstimemin string) *JobsRecordBuilder {
	rb.v.BucketsTimeMin = &bucketstimemin
	return rb
}

// BucketsTimeTotal total bucket processing time

func (rb *JobsRecordBuilder) BucketsTimeTotal(bucketstimetotal string) *JobsRecordBuilder {
	rb.v.BucketsTimeTotal = &bucketstimetotal
	return rb
}

// DataBuckets total bucket count

func (rb *JobsRecordBuilder) DataBuckets(databuckets string) *JobsRecordBuilder {
	rb.v.DataBuckets = &databuckets
	return rb
}

// DataEarliestRecord earliest record time

func (rb *JobsRecordBuilder) DataEarliestRecord(dataearliestrecord string) *JobsRecordBuilder {
	rb.v.DataEarliestRecord = &dataearliestrecord
	return rb
}

// DataEmptyBuckets number of empty buckets

func (rb *JobsRecordBuilder) DataEmptyBuckets(dataemptybuckets string) *JobsRecordBuilder {
	rb.v.DataEmptyBuckets = &dataemptybuckets
	return rb
}

// DataInputBytes total input bytes

func (rb *JobsRecordBuilder) DataInputBytes(datainputbytes *ByteSizeBuilder) *JobsRecordBuilder {
	v := datainputbytes.Build()
	rb.v.DataInputBytes = &v
	return rb
}

// DataInputFields total field count

func (rb *JobsRecordBuilder) DataInputFields(datainputfields string) *JobsRecordBuilder {
	rb.v.DataInputFields = &datainputfields
	return rb
}

// DataInputRecords total record count

func (rb *JobsRecordBuilder) DataInputRecords(datainputrecords string) *JobsRecordBuilder {
	rb.v.DataInputRecords = &datainputrecords
	return rb
}

// DataInvalidDates number of records with invalid dates

func (rb *JobsRecordBuilder) DataInvalidDates(datainvaliddates string) *JobsRecordBuilder {
	rb.v.DataInvalidDates = &datainvaliddates
	return rb
}

// DataLast last time data was seen

func (rb *JobsRecordBuilder) DataLast(datalast string) *JobsRecordBuilder {
	rb.v.DataLast = &datalast
	return rb
}

// DataLastEmptyBucket last time an empty bucket occurred

func (rb *JobsRecordBuilder) DataLastEmptyBucket(datalastemptybucket string) *JobsRecordBuilder {
	rb.v.DataLastEmptyBucket = &datalastemptybucket
	return rb
}

// DataLastSparseBucket last time a sparse bucket occurred

func (rb *JobsRecordBuilder) DataLastSparseBucket(datalastsparsebucket string) *JobsRecordBuilder {
	rb.v.DataLastSparseBucket = &datalastsparsebucket
	return rb
}

// DataLatestRecord latest record time

func (rb *JobsRecordBuilder) DataLatestRecord(datalatestrecord string) *JobsRecordBuilder {
	rb.v.DataLatestRecord = &datalatestrecord
	return rb
}

// DataMissingFields number of records with missing fields

func (rb *JobsRecordBuilder) DataMissingFields(datamissingfields string) *JobsRecordBuilder {
	rb.v.DataMissingFields = &datamissingfields
	return rb
}

// DataOutOfOrderTimestamps number of records handled out of order

func (rb *JobsRecordBuilder) DataOutOfOrderTimestamps(dataoutofordertimestamps string) *JobsRecordBuilder {
	rb.v.DataOutOfOrderTimestamps = &dataoutofordertimestamps
	return rb
}

// DataProcessedFields number of processed fields

func (rb *JobsRecordBuilder) DataProcessedFields(dataprocessedfields string) *JobsRecordBuilder {
	rb.v.DataProcessedFields = &dataprocessedfields
	return rb
}

// DataProcessedRecords number of processed records

func (rb *JobsRecordBuilder) DataProcessedRecords(dataprocessedrecords string) *JobsRecordBuilder {
	rb.v.DataProcessedRecords = &dataprocessedrecords
	return rb
}

// DataSparseBuckets number of sparse buckets

func (rb *JobsRecordBuilder) DataSparseBuckets(datasparsebuckets string) *JobsRecordBuilder {
	rb.v.DataSparseBuckets = &datasparsebuckets
	return rb
}

// ForecastsMemoryAvg average memory used by forecasts

func (rb *JobsRecordBuilder) ForecastsMemoryAvg(forecastsmemoryavg string) *JobsRecordBuilder {
	rb.v.ForecastsMemoryAvg = &forecastsmemoryavg
	return rb
}

// ForecastsMemoryMax maximum memory used by forecasts

func (rb *JobsRecordBuilder) ForecastsMemoryMax(forecastsmemorymax string) *JobsRecordBuilder {
	rb.v.ForecastsMemoryMax = &forecastsmemorymax
	return rb
}

// ForecastsMemoryMin minimum memory used by forecasts

func (rb *JobsRecordBuilder) ForecastsMemoryMin(forecastsmemorymin string) *JobsRecordBuilder {
	rb.v.ForecastsMemoryMin = &forecastsmemorymin
	return rb
}

// ForecastsMemoryTotal total memory used by all forecasts

func (rb *JobsRecordBuilder) ForecastsMemoryTotal(forecastsmemorytotal string) *JobsRecordBuilder {
	rb.v.ForecastsMemoryTotal = &forecastsmemorytotal
	return rb
}

// ForecastsRecordsAvg average record count for forecasts

func (rb *JobsRecordBuilder) ForecastsRecordsAvg(forecastsrecordsavg string) *JobsRecordBuilder {
	rb.v.ForecastsRecordsAvg = &forecastsrecordsavg
	return rb
}

// ForecastsRecordsMax maximum record count for forecasts

func (rb *JobsRecordBuilder) ForecastsRecordsMax(forecastsrecordsmax string) *JobsRecordBuilder {
	rb.v.ForecastsRecordsMax = &forecastsrecordsmax
	return rb
}

// ForecastsRecordsMin minimum record count for forecasts

func (rb *JobsRecordBuilder) ForecastsRecordsMin(forecastsrecordsmin string) *JobsRecordBuilder {
	rb.v.ForecastsRecordsMin = &forecastsrecordsmin
	return rb
}

// ForecastsRecordsTotal total record count for all forecasts

func (rb *JobsRecordBuilder) ForecastsRecordsTotal(forecastsrecordstotal string) *JobsRecordBuilder {
	rb.v.ForecastsRecordsTotal = &forecastsrecordstotal
	return rb
}

// ForecastsTimeAvg average runtime for all forecasts (milliseconds)

func (rb *JobsRecordBuilder) ForecastsTimeAvg(forecaststimeavg string) *JobsRecordBuilder {
	rb.v.ForecastsTimeAvg = &forecaststimeavg
	return rb
}

// ForecastsTimeMax maximum run time for forecasts

func (rb *JobsRecordBuilder) ForecastsTimeMax(forecaststimemax string) *JobsRecordBuilder {
	rb.v.ForecastsTimeMax = &forecaststimemax
	return rb
}

// ForecastsTimeMin minimum runtime for forecasts

func (rb *JobsRecordBuilder) ForecastsTimeMin(forecaststimemin string) *JobsRecordBuilder {
	rb.v.ForecastsTimeMin = &forecaststimemin
	return rb
}

// ForecastsTimeTotal total runtime for all forecasts

func (rb *JobsRecordBuilder) ForecastsTimeTotal(forecaststimetotal string) *JobsRecordBuilder {
	rb.v.ForecastsTimeTotal = &forecaststimetotal
	return rb
}

// ForecastsTotal total number of forecasts

func (rb *JobsRecordBuilder) ForecastsTotal(forecaststotal string) *JobsRecordBuilder {
	rb.v.ForecastsTotal = &forecaststotal
	return rb
}

// Id the job_id

func (rb *JobsRecordBuilder) Id(id Id) *JobsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// ModelBucketAllocationFailures number of bucket allocation failures

func (rb *JobsRecordBuilder) ModelBucketAllocationFailures(modelbucketallocationfailures string) *JobsRecordBuilder {
	rb.v.ModelBucketAllocationFailures = &modelbucketallocationfailures
	return rb
}

// ModelByFields count of 'by' fields

func (rb *JobsRecordBuilder) ModelByFields(modelbyfields string) *JobsRecordBuilder {
	rb.v.ModelByFields = &modelbyfields
	return rb
}

// ModelBytes model size

func (rb *JobsRecordBuilder) ModelBytes(modelbytes *ByteSizeBuilder) *JobsRecordBuilder {
	v := modelbytes.Build()
	rb.v.ModelBytes = &v
	return rb
}

// ModelBytesExceeded how much the model has exceeded the limit

func (rb *JobsRecordBuilder) ModelBytesExceeded(modelbytesexceeded *ByteSizeBuilder) *JobsRecordBuilder {
	v := modelbytesexceeded.Build()
	rb.v.ModelBytesExceeded = &v
	return rb
}

// ModelCategorizationStatus current categorization status

func (rb *JobsRecordBuilder) ModelCategorizationStatus(modelcategorizationstatus categorizationstatus.CategorizationStatus) *JobsRecordBuilder {
	rb.v.ModelCategorizationStatus = &modelcategorizationstatus
	return rb
}

// ModelCategorizedDocCount count of categorized documents

func (rb *JobsRecordBuilder) ModelCategorizedDocCount(modelcategorizeddoccount string) *JobsRecordBuilder {
	rb.v.ModelCategorizedDocCount = &modelcategorizeddoccount
	return rb
}

// ModelDeadCategoryCount count of dead categories

func (rb *JobsRecordBuilder) ModelDeadCategoryCount(modeldeadcategorycount string) *JobsRecordBuilder {
	rb.v.ModelDeadCategoryCount = &modeldeadcategorycount
	return rb
}

// ModelFailedCategoryCount count of failed categories

func (rb *JobsRecordBuilder) ModelFailedCategoryCount(modelfailedcategorycount string) *JobsRecordBuilder {
	rb.v.ModelFailedCategoryCount = &modelfailedcategorycount
	return rb
}

// ModelFrequentCategoryCount count of frequent categories

func (rb *JobsRecordBuilder) ModelFrequentCategoryCount(modelfrequentcategorycount string) *JobsRecordBuilder {
	rb.v.ModelFrequentCategoryCount = &modelfrequentcategorycount
	return rb
}

// ModelLogTime when the model stats were gathered

func (rb *JobsRecordBuilder) ModelLogTime(modellogtime string) *JobsRecordBuilder {
	rb.v.ModelLogTime = &modellogtime
	return rb
}

// ModelMemoryLimit model memory limit

func (rb *JobsRecordBuilder) ModelMemoryLimit(modelmemorylimit string) *JobsRecordBuilder {
	rb.v.ModelMemoryLimit = &modelmemorylimit
	return rb
}

// ModelMemoryStatus current memory status

func (rb *JobsRecordBuilder) ModelMemoryStatus(modelmemorystatus memorystatus.MemoryStatus) *JobsRecordBuilder {
	rb.v.ModelMemoryStatus = &modelmemorystatus
	return rb
}

// ModelOverFields count of 'over' fields

func (rb *JobsRecordBuilder) ModelOverFields(modeloverfields string) *JobsRecordBuilder {
	rb.v.ModelOverFields = &modeloverfields
	return rb
}

// ModelPartitionFields count of 'partition' fields

func (rb *JobsRecordBuilder) ModelPartitionFields(modelpartitionfields string) *JobsRecordBuilder {
	rb.v.ModelPartitionFields = &modelpartitionfields
	return rb
}

// ModelRareCategoryCount count of rare categories

func (rb *JobsRecordBuilder) ModelRareCategoryCount(modelrarecategorycount string) *JobsRecordBuilder {
	rb.v.ModelRareCategoryCount = &modelrarecategorycount
	return rb
}

// ModelTimestamp the time of the last record when the model stats were gathered

func (rb *JobsRecordBuilder) ModelTimestamp(modeltimestamp string) *JobsRecordBuilder {
	rb.v.ModelTimestamp = &modeltimestamp
	return rb
}

// ModelTotalCategoryCount count of categories

func (rb *JobsRecordBuilder) ModelTotalCategoryCount(modeltotalcategorycount string) *JobsRecordBuilder {
	rb.v.ModelTotalCategoryCount = &modeltotalcategorycount
	return rb
}

// NodeAddress network address of the assigned node

func (rb *JobsRecordBuilder) NodeAddress(nodeaddress string) *JobsRecordBuilder {
	rb.v.NodeAddress = &nodeaddress
	return rb
}

// NodeEphemeralId ephemeral id of the assigned node

func (rb *JobsRecordBuilder) NodeEphemeralId(nodeephemeralid NodeId) *JobsRecordBuilder {
	rb.v.NodeEphemeralId = &nodeephemeralid
	return rb
}

// NodeId id of the assigned node

func (rb *JobsRecordBuilder) NodeId(nodeid NodeId) *JobsRecordBuilder {
	rb.v.NodeId = &nodeid
	return rb
}

// NodeName name of the assigned node

func (rb *JobsRecordBuilder) NodeName(nodename string) *JobsRecordBuilder {
	rb.v.NodeName = &nodename
	return rb
}

// OpenedTime the amount of time the job has been opened

func (rb *JobsRecordBuilder) OpenedTime(openedtime string) *JobsRecordBuilder {
	rb.v.OpenedTime = &openedtime
	return rb
}

// State the job state

func (rb *JobsRecordBuilder) State(state jobstate.JobState) *JobsRecordBuilder {
	rb.v.State = &state
	return rb
}
