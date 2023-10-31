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

package postdata

// Response holds the response body struct for the package postdata
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/post_data/MlPostJobDataResponse.ts#L23-L41

type Response struct {
	BucketCount              int64  `json:"bucket_count"`
	EarliestRecordTimestamp  int64  `json:"earliest_record_timestamp"`
	EmptyBucketCount         int64  `json:"empty_bucket_count"`
	InputBytes               int64  `json:"input_bytes"`
	InputFieldCount          int64  `json:"input_field_count"`
	InputRecordCount         int64  `json:"input_record_count"`
	InvalidDateCount         int64  `json:"invalid_date_count"`
	JobId                    string `json:"job_id"`
	LastDataTime             int    `json:"last_data_time"`
	LatestRecordTimestamp    int64  `json:"latest_record_timestamp"`
	MissingFieldCount        int64  `json:"missing_field_count"`
	OutOfOrderTimestampCount int64  `json:"out_of_order_timestamp_count"`
	ProcessedFieldCount      int64  `json:"processed_field_count"`
	ProcessedRecordCount     int64  `json:"processed_record_count"`
	SparseBucketCount        int64  `json:"sparse_bucket_count"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
