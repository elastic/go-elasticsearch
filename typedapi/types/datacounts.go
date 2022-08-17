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

// DataCounts type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Job.ts#L129-L149
type DataCounts struct {
	BucketCount                 int64  `json:"bucket_count"`
	EarliestRecordTimestamp     *int64 `json:"earliest_record_timestamp,omitempty"`
	EmptyBucketCount            int64  `json:"empty_bucket_count"`
	InputBytes                  int64  `json:"input_bytes"`
	InputFieldCount             int64  `json:"input_field_count"`
	InputRecordCount            int64  `json:"input_record_count"`
	InvalidDateCount            int64  `json:"invalid_date_count"`
	JobId                       Id     `json:"job_id"`
	LastDataTime                *int64 `json:"last_data_time,omitempty"`
	LatestBucketTimestamp       *int64 `json:"latest_bucket_timestamp,omitempty"`
	LatestEmptyBucketTimestamp  *int64 `json:"latest_empty_bucket_timestamp,omitempty"`
	LatestRecordTimestamp       *int64 `json:"latest_record_timestamp,omitempty"`
	LatestSparseBucketTimestamp *int64 `json:"latest_sparse_bucket_timestamp,omitempty"`
	LogTime                     *int64 `json:"log_time,omitempty"`
	MissingFieldCount           int64  `json:"missing_field_count"`
	OutOfOrderTimestampCount    int64  `json:"out_of_order_timestamp_count"`
	ProcessedFieldCount         int64  `json:"processed_field_count"`
	ProcessedRecordCount        int64  `json:"processed_record_count"`
	SparseBucketCount           int64  `json:"sparse_bucket_count"`
}

// DataCountsBuilder holds DataCounts struct and provides a builder API.
type DataCountsBuilder struct {
	v *DataCounts
}

// NewDataCounts provides a builder for the DataCounts struct.
func NewDataCountsBuilder() *DataCountsBuilder {
	r := DataCountsBuilder{
		&DataCounts{},
	}

	return &r
}

// Build finalize the chain and returns the DataCounts struct
func (rb *DataCountsBuilder) Build() DataCounts {
	return *rb.v
}

func (rb *DataCountsBuilder) BucketCount(bucketcount int64) *DataCountsBuilder {
	rb.v.BucketCount = bucketcount
	return rb
}

func (rb *DataCountsBuilder) EarliestRecordTimestamp(earliestrecordtimestamp int64) *DataCountsBuilder {
	rb.v.EarliestRecordTimestamp = &earliestrecordtimestamp
	return rb
}

func (rb *DataCountsBuilder) EmptyBucketCount(emptybucketcount int64) *DataCountsBuilder {
	rb.v.EmptyBucketCount = emptybucketcount
	return rb
}

func (rb *DataCountsBuilder) InputBytes(inputbytes int64) *DataCountsBuilder {
	rb.v.InputBytes = inputbytes
	return rb
}

func (rb *DataCountsBuilder) InputFieldCount(inputfieldcount int64) *DataCountsBuilder {
	rb.v.InputFieldCount = inputfieldcount
	return rb
}

func (rb *DataCountsBuilder) InputRecordCount(inputrecordcount int64) *DataCountsBuilder {
	rb.v.InputRecordCount = inputrecordcount
	return rb
}

func (rb *DataCountsBuilder) InvalidDateCount(invaliddatecount int64) *DataCountsBuilder {
	rb.v.InvalidDateCount = invaliddatecount
	return rb
}

func (rb *DataCountsBuilder) JobId(jobid Id) *DataCountsBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *DataCountsBuilder) LastDataTime(lastdatatime int64) *DataCountsBuilder {
	rb.v.LastDataTime = &lastdatatime
	return rb
}

func (rb *DataCountsBuilder) LatestBucketTimestamp(latestbuckettimestamp int64) *DataCountsBuilder {
	rb.v.LatestBucketTimestamp = &latestbuckettimestamp
	return rb
}

func (rb *DataCountsBuilder) LatestEmptyBucketTimestamp(latestemptybuckettimestamp int64) *DataCountsBuilder {
	rb.v.LatestEmptyBucketTimestamp = &latestemptybuckettimestamp
	return rb
}

func (rb *DataCountsBuilder) LatestRecordTimestamp(latestrecordtimestamp int64) *DataCountsBuilder {
	rb.v.LatestRecordTimestamp = &latestrecordtimestamp
	return rb
}

func (rb *DataCountsBuilder) LatestSparseBucketTimestamp(latestsparsebuckettimestamp int64) *DataCountsBuilder {
	rb.v.LatestSparseBucketTimestamp = &latestsparsebuckettimestamp
	return rb
}

func (rb *DataCountsBuilder) LogTime(logtime int64) *DataCountsBuilder {
	rb.v.LogTime = &logtime
	return rb
}

func (rb *DataCountsBuilder) MissingFieldCount(missingfieldcount int64) *DataCountsBuilder {
	rb.v.MissingFieldCount = missingfieldcount
	return rb
}

func (rb *DataCountsBuilder) OutOfOrderTimestampCount(outofordertimestampcount int64) *DataCountsBuilder {
	rb.v.OutOfOrderTimestampCount = outofordertimestampcount
	return rb
}

func (rb *DataCountsBuilder) ProcessedFieldCount(processedfieldcount int64) *DataCountsBuilder {
	rb.v.ProcessedFieldCount = processedfieldcount
	return rb
}

func (rb *DataCountsBuilder) ProcessedRecordCount(processedrecordcount int64) *DataCountsBuilder {
	rb.v.ProcessedRecordCount = processedrecordcount
	return rb
}

func (rb *DataCountsBuilder) SparseBucketCount(sparsebucketcount int64) *DataCountsBuilder {
	rb.v.SparseBucketCount = sparsebucketcount
	return rb
}
