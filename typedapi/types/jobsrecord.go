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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

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
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cat/ml_jobs/types.ts#L24-L347
type JobsRecord struct {
	// AssignmentExplanation For open anomaly detection jobs only, contains messages relating to the
	// selection of a node to run the job.
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// BucketsCount The number of bucket results produced by the job.
	BucketsCount *string `json:"buckets.count,omitempty"`
	// BucketsTimeExpAvg The exponential moving average of all bucket processing times, in
	// milliseconds.
	BucketsTimeExpAvg *string `json:"buckets.time.exp_avg,omitempty"`
	// BucketsTimeExpAvgHour The exponential moving average of bucket processing times calculated in a one
	// hour time window, in milliseconds.
	BucketsTimeExpAvgHour *string `json:"buckets.time.exp_avg_hour,omitempty"`
	// BucketsTimeMax The maximum of all bucket processing times, in milliseconds.
	BucketsTimeMax *string `json:"buckets.time.max,omitempty"`
	// BucketsTimeMin The minimum of all bucket processing times, in milliseconds.
	BucketsTimeMin *string `json:"buckets.time.min,omitempty"`
	// BucketsTimeTotal The sum of all bucket processing times, in milliseconds.
	BucketsTimeTotal *string `json:"buckets.time.total,omitempty"`
	// DataBuckets The total number of buckets processed.
	DataBuckets *string `json:"data.buckets,omitempty"`
	// DataEarliestRecord The timestamp of the earliest chronologically input document.
	DataEarliestRecord *string `json:"data.earliest_record,omitempty"`
	// DataEmptyBuckets The number of buckets which did not contain any data.
	// If your data contains many empty buckets, consider increasing your
	// `bucket_span` or using functions that are tolerant to gaps in data such as
	// mean, `non_null_sum` or `non_zero_count`.
	DataEmptyBuckets *string `json:"data.empty_buckets,omitempty"`
	// DataInputBytes The number of bytes of input data posted to the anomaly detection job.
	DataInputBytes ByteSize `json:"data.input_bytes,omitempty"`
	// DataInputFields The total number of fields in input documents posted to the anomaly detection
	// job.
	// This count includes fields that are not used in the analysis.
	// However, be aware that if you are using a datafeed, it extracts only the
	// required fields from the documents it retrieves before posting them to the
	// job.
	DataInputFields *string `json:"data.input_fields,omitempty"`
	// DataInputRecords The number of input documents posted to the anomaly detection job.
	DataInputRecords *string `json:"data.input_records,omitempty"`
	// DataInvalidDates The number of input documents with either a missing date field or a date that
	// could not be parsed.
	DataInvalidDates *string `json:"data.invalid_dates,omitempty"`
	// DataLast The timestamp at which data was last analyzed, according to server time.
	DataLast *string `json:"data.last,omitempty"`
	// DataLastEmptyBucket The timestamp of the last bucket that did not contain any data.
	DataLastEmptyBucket *string `json:"data.last_empty_bucket,omitempty"`
	// DataLastSparseBucket The timestamp of the last bucket that was considered sparse.
	DataLastSparseBucket *string `json:"data.last_sparse_bucket,omitempty"`
	// DataLatestRecord The timestamp of the latest chronologically input document.
	DataLatestRecord *string `json:"data.latest_record,omitempty"`
	// DataMissingFields The number of input documents that are missing a field that the anomaly
	// detection job is configured to analyze.
	// Input documents with missing fields are still processed because it is
	// possible that not all fields are missing.
	// If you are using datafeeds or posting data to the job in JSON format, a high
	// `missing_field_count` is often not an indication of data issues.
	// It is not necessarily a cause for concern.
	DataMissingFields *string `json:"data.missing_fields,omitempty"`
	// DataOutOfOrderTimestamps The number of input documents that have a timestamp chronologically preceding
	// the start of the current anomaly detection bucket offset by the latency
	// window.
	// This information is applicable only when you provide data to the anomaly
	// detection job by using the post data API.
	// These out of order documents are discarded, since jobs require time series
	// data to be in ascending chronological order.
	DataOutOfOrderTimestamps *string `json:"data.out_of_order_timestamps,omitempty"`
	// DataProcessedFields The total number of fields in all the documents that have been processed by
	// the anomaly detection job.
	// Only fields that are specified in the detector configuration object
	// contribute to this count.
	// The timestamp is not included in this count.
	DataProcessedFields *string `json:"data.processed_fields,omitempty"`
	// DataProcessedRecords The number of input documents that have been processed by the anomaly
	// detection job.
	// This value includes documents with missing fields, since they are nonetheless
	// analyzed.
	// If you use datafeeds and have aggregations in your search query, the
	// `processed_record_count` is the number of aggregation results processed, not
	// the number of Elasticsearch documents.
	DataProcessedRecords *string `json:"data.processed_records,omitempty"`
	// DataSparseBuckets The number of buckets that contained few data points compared to the expected
	// number of data points.
	// If your data contains many sparse buckets, consider using a longer
	// `bucket_span`.
	DataSparseBuckets *string `json:"data.sparse_buckets,omitempty"`
	// ForecastsMemoryAvg The average memory usage in bytes for forecasts related to the anomaly
	// detection job.
	ForecastsMemoryAvg *string `json:"forecasts.memory.avg,omitempty"`
	// ForecastsMemoryMax The maximum memory usage in bytes for forecasts related to the anomaly
	// detection job.
	ForecastsMemoryMax *string `json:"forecasts.memory.max,omitempty"`
	// ForecastsMemoryMin The minimum memory usage in bytes for forecasts related to the anomaly
	// detection job.
	ForecastsMemoryMin *string `json:"forecasts.memory.min,omitempty"`
	// ForecastsMemoryTotal The total memory usage in bytes for forecasts related to the anomaly
	// detection job.
	ForecastsMemoryTotal *string `json:"forecasts.memory.total,omitempty"`
	// ForecastsRecordsAvg The average number of `model_forecast` documents written for forecasts
	// related to the anomaly detection job.
	ForecastsRecordsAvg *string `json:"forecasts.records.avg,omitempty"`
	// ForecastsRecordsMax The maximum number of `model_forecast` documents written for forecasts
	// related to the anomaly detection job.
	ForecastsRecordsMax *string `json:"forecasts.records.max,omitempty"`
	// ForecastsRecordsMin The minimum number of `model_forecast` documents written for forecasts
	// related to the anomaly detection job.
	ForecastsRecordsMin *string `json:"forecasts.records.min,omitempty"`
	// ForecastsRecordsTotal The total number of `model_forecast` documents written for forecasts related
	// to the anomaly detection job.
	ForecastsRecordsTotal *string `json:"forecasts.records.total,omitempty"`
	// ForecastsTimeAvg The average runtime in milliseconds for forecasts related to the anomaly
	// detection job.
	ForecastsTimeAvg *string `json:"forecasts.time.avg,omitempty"`
	// ForecastsTimeMax The maximum runtime in milliseconds for forecasts related to the anomaly
	// detection job.
	ForecastsTimeMax *string `json:"forecasts.time.max,omitempty"`
	// ForecastsTimeMin The minimum runtime in milliseconds for forecasts related to the anomaly
	// detection job.
	ForecastsTimeMin *string `json:"forecasts.time.min,omitempty"`
	// ForecastsTimeTotal The total runtime in milliseconds for forecasts related to the anomaly
	// detection job.
	ForecastsTimeTotal *string `json:"forecasts.time.total,omitempty"`
	// ForecastsTotal The number of individual forecasts currently available for the job.
	// A value of one or more indicates that forecasts exist.
	ForecastsTotal *string `json:"forecasts.total,omitempty"`
	// Id The anomaly detection job identifier.
	Id *string `json:"id,omitempty"`
	// ModelBucketAllocationFailures The number of buckets for which new entities in incoming data were not
	// processed due to insufficient model memory.
	// This situation is also signified by a `hard_limit: memory_status` property
	// value.
	ModelBucketAllocationFailures *string `json:"model.bucket_allocation_failures,omitempty"`
	// ModelByFields The number of `by` field values that were analyzed by the models.
	// This value is cumulative for all detectors in the job.
	ModelByFields *string `json:"model.by_fields,omitempty"`
	// ModelBytes The number of bytes of memory used by the models.
	// This is the maximum value since the last time the model was persisted.
	// If the job is closed, this value indicates the latest size.
	ModelBytes ByteSize `json:"model.bytes,omitempty"`
	// ModelBytesExceeded The number of bytes over the high limit for memory usage at the last
	// allocation failure.
	ModelBytesExceeded ByteSize `json:"model.bytes_exceeded,omitempty"`
	// ModelCategorizationStatus The status of categorization for the job.
	ModelCategorizationStatus *categorizationstatus.CategorizationStatus `json:"model.categorization_status,omitempty"`
	// ModelCategorizedDocCount The number of documents that have had a field categorized.
	ModelCategorizedDocCount *string `json:"model.categorized_doc_count,omitempty"`
	// ModelDeadCategoryCount The number of categories created by categorization that will never be
	// assigned again because another category’s definition makes it a superset of
	// the dead category.
	// Dead categories are a side effect of the way categorization has no prior
	// training.
	ModelDeadCategoryCount *string `json:"model.dead_category_count,omitempty"`
	// ModelFailedCategoryCount The number of times that categorization wanted to create a new category but
	// couldn’t because the job had hit its `model_memory_limit`.
	// This count does not track which specific categories failed to be created.
	// Therefore you cannot use this value to determine the number of unique
	// categories that were missed.
	ModelFailedCategoryCount *string `json:"model.failed_category_count,omitempty"`
	// ModelFrequentCategoryCount The number of categories that match more than 1% of categorized documents.
	ModelFrequentCategoryCount *string `json:"model.frequent_category_count,omitempty"`
	// ModelLogTime The timestamp when the model stats were gathered, according to server time.
	ModelLogTime *string `json:"model.log_time,omitempty"`
	// ModelMemoryLimit The upper limit for model memory usage, checked on increasing values.
	ModelMemoryLimit *string `json:"model.memory_limit,omitempty"`
	// ModelMemoryStatus The status of the mathematical models.
	ModelMemoryStatus *memorystatus.MemoryStatus `json:"model.memory_status,omitempty"`
	// ModelOverFields The number of `over` field values that were analyzed by the models.
	// This value is cumulative for all detectors in the job.
	ModelOverFields *string `json:"model.over_fields,omitempty"`
	// ModelPartitionFields The number of `partition` field values that were analyzed by the models.
	// This value is cumulative for all detectors in the job.
	ModelPartitionFields *string `json:"model.partition_fields,omitempty"`
	// ModelRareCategoryCount The number of categories that match just one categorized document.
	ModelRareCategoryCount *string `json:"model.rare_category_count,omitempty"`
	// ModelTimestamp The timestamp of the last record when the model stats were gathered.
	ModelTimestamp *string `json:"model.timestamp,omitempty"`
	// ModelTotalCategoryCount The number of categories created by categorization.
	ModelTotalCategoryCount *string `json:"model.total_category_count,omitempty"`
	// NodeAddress The network address of the assigned node.
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId The ephemeral identifier of the assigned node.
	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`
	// NodeId The uniqe identifier of the assigned node.
	NodeId *string `json:"node.id,omitempty"`
	// NodeName The name of the assigned node.
	NodeName *string `json:"node.name,omitempty"`
	// OpenedTime For open jobs only, the amount of time the job has been opened.
	OpenedTime *string `json:"opened_time,omitempty"`
	// State The status of the anomaly detection job.
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
