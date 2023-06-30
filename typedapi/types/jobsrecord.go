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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/categorizationstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jobstate"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/memorystatus"
)

// JobsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cat/ml_jobs/types.ts#L24-L325
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
	DataInputBytes ByteSize `json:"data.input_bytes,omitempty"`
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
	ModelBytes ByteSize `json:"model.bytes,omitempty"`
	// ModelBytesExceeded how much the model has exceeded the limit
	ModelBytesExceeded ByteSize `json:"model.bytes_exceeded,omitempty"`
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

func (s *JobsRecord) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "assignment_explanation", "ae":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AssignmentExplanation = &o

		case "buckets.count", "bc", "bucketsCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BucketsCount = &o

		case "buckets.time.exp_avg", "btea", "bucketsTimeExpAvg":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BucketsTimeExpAvg = &o

		case "buckets.time.exp_avg_hour", "bteah", "bucketsTimeExpAvgHour":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BucketsTimeExpAvgHour = &o

		case "buckets.time.max", "btmax", "bucketsTimeMax":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BucketsTimeMax = &o

		case "buckets.time.min", "btmin", "bucketsTimeMin":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BucketsTimeMin = &o

		case "buckets.time.total", "btt", "bucketsTimeTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BucketsTimeTotal = &o

		case "data.buckets", "db", "dataBuckets":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataBuckets = &o

		case "data.earliest_record", "der", "dataEarliestRecord":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataEarliestRecord = &o

		case "data.empty_buckets", "deb", "dataEmptyBuckets":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataEmptyBuckets = &o

		case "data.input_bytes", "dib", "dataInputBytes":
			if err := dec.Decode(&s.DataInputBytes); err != nil {
				return err
			}

		case "data.input_fields", "dif", "dataInputFields":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataInputFields = &o

		case "data.input_records", "dir", "dataInputRecords":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataInputRecords = &o

		case "data.invalid_dates", "did", "dataInvalidDates":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataInvalidDates = &o

		case "data.last", "dl", "dataLast":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataLast = &o

		case "data.last_empty_bucket", "dleb", "dataLastEmptyBucket":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataLastEmptyBucket = &o

		case "data.last_sparse_bucket", "dlsb", "dataLastSparseBucket":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataLastSparseBucket = &o

		case "data.latest_record", "dlr", "dataLatestRecord":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataLatestRecord = &o

		case "data.missing_fields", "dmf", "dataMissingFields":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataMissingFields = &o

		case "data.out_of_order_timestamps", "doot", "dataOutOfOrderTimestamps":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataOutOfOrderTimestamps = &o

		case "data.processed_fields", "dpf", "dataProcessedFields":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataProcessedFields = &o

		case "data.processed_records", "dpr", "dataProcessedRecords":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataProcessedRecords = &o

		case "data.sparse_buckets", "dsb", "dataSparseBuckets":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DataSparseBuckets = &o

		case "forecasts.memory.avg", "fmavg", "forecastsMemoryAvg":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsMemoryAvg = &o

		case "forecasts.memory.max", "fmmax", "forecastsMemoryMax":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsMemoryMax = &o

		case "forecasts.memory.min", "fmmin", "forecastsMemoryMin":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsMemoryMin = &o

		case "forecasts.memory.total", "fmt", "forecastsMemoryTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsMemoryTotal = &o

		case "forecasts.records.avg", "fravg", "forecastsRecordsAvg":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsRecordsAvg = &o

		case "forecasts.records.max", "frmax", "forecastsRecordsMax":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsRecordsMax = &o

		case "forecasts.records.min", "frmin", "forecastsRecordsMin":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsRecordsMin = &o

		case "forecasts.records.total", "frt", "forecastsRecordsTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsRecordsTotal = &o

		case "forecasts.time.avg", "ftavg", "forecastsTimeAvg":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsTimeAvg = &o

		case "forecasts.time.max", "ftmax", "forecastsTimeMax":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsTimeMax = &o

		case "forecasts.time.min", "ftmin", "forecastsTimeMin":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsTimeMin = &o

		case "forecasts.time.total", "ftt", "forecastsTimeTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsTimeTotal = &o

		case "forecasts.total", "ft", "forecastsTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ForecastsTotal = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "model.bucket_allocation_failures", "mbaf", "modelBucketAllocationFailures":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelBucketAllocationFailures = &o

		case "model.by_fields", "mbf", "modelByFields":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelByFields = &o

		case "model.bytes", "mb", "modelBytes":
			if err := dec.Decode(&s.ModelBytes); err != nil {
				return err
			}

		case "model.bytes_exceeded", "mbe", "modelBytesExceeded":
			if err := dec.Decode(&s.ModelBytesExceeded); err != nil {
				return err
			}

		case "model.categorization_status", "mcs", "modelCategorizationStatus":
			if err := dec.Decode(&s.ModelCategorizationStatus); err != nil {
				return err
			}

		case "model.categorized_doc_count", "mcdc", "modelCategorizedDocCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelCategorizedDocCount = &o

		case "model.dead_category_count", "mdcc", "modelDeadCategoryCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelDeadCategoryCount = &o

		case "model.failed_category_count", "mfcc", "modelFailedCategoryCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelFailedCategoryCount = &o

		case "model.frequent_category_count", "modelFrequentCategoryCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelFrequentCategoryCount = &o

		case "model.log_time", "mlt", "modelLogTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelLogTime = &o

		case "model.memory_limit", "mml", "modelMemoryLimit":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelMemoryLimit = &o

		case "model.memory_status", "mms", "modelMemoryStatus":
			if err := dec.Decode(&s.ModelMemoryStatus); err != nil {
				return err
			}

		case "model.over_fields", "mof", "modelOverFields":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelOverFields = &o

		case "model.partition_fields", "mpf", "modelPartitionFields":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelPartitionFields = &o

		case "model.rare_category_count", "mrcc", "modelRareCategoryCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelRareCategoryCount = &o

		case "model.timestamp", "mt", "modelTimestamp":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelTimestamp = &o

		case "model.total_category_count", "mtcc", "modelTotalCategoryCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelTotalCategoryCount = &o

		case "node.address", "na", "nodeAddress":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeAddress = &o

		case "node.ephemeral_id", "ne", "nodeEphemeralId":
			if err := dec.Decode(&s.NodeEphemeralId); err != nil {
				return err
			}

		case "node.id", "ni", "nodeId":
			if err := dec.Decode(&s.NodeId); err != nil {
				return err
			}

		case "node.name", "nn", "nodeName":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeName = &o

		case "opened_time", "ot":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.OpenedTime = &o

		case "state", "s":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewJobsRecord returns a JobsRecord.
func NewJobsRecord() *JobsRecord {
	r := &JobsRecord{}

	return r
}
