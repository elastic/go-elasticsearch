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
)

// DataCounts type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Job.ts#L352-L372
type DataCounts struct {
	BucketCount                 int64  `json:"bucket_count"`
	EarliestRecordTimestamp     *int64 `json:"earliest_record_timestamp,omitempty"`
	EmptyBucketCount            int64  `json:"empty_bucket_count"`
	InputBytes                  int64  `json:"input_bytes"`
	InputFieldCount             int64  `json:"input_field_count"`
	InputRecordCount            int64  `json:"input_record_count"`
	InvalidDateCount            int64  `json:"invalid_date_count"`
	JobId                       string `json:"job_id"`
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

func (s *DataCounts) UnmarshalJSON(data []byte) error {

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

		case "bucket_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.BucketCount = value
			case float64:
				f := int64(v)
				s.BucketCount = f
			}

		case "earliest_record_timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.EarliestRecordTimestamp = &value
			case float64:
				f := int64(v)
				s.EarliestRecordTimestamp = &f
			}

		case "empty_bucket_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.EmptyBucketCount = value
			case float64:
				f := int64(v)
				s.EmptyBucketCount = f
			}

		case "input_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.InputBytes = value
			case float64:
				f := int64(v)
				s.InputBytes = f
			}

		case "input_field_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.InputFieldCount = value
			case float64:
				f := int64(v)
				s.InputFieldCount = f
			}

		case "input_record_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.InputRecordCount = value
			case float64:
				f := int64(v)
				s.InputRecordCount = f
			}

		case "invalid_date_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.InvalidDateCount = value
			case float64:
				f := int64(v)
				s.InvalidDateCount = f
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "last_data_time":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LastDataTime = &value
			case float64:
				f := int64(v)
				s.LastDataTime = &f
			}

		case "latest_bucket_timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LatestBucketTimestamp = &value
			case float64:
				f := int64(v)
				s.LatestBucketTimestamp = &f
			}

		case "latest_empty_bucket_timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LatestEmptyBucketTimestamp = &value
			case float64:
				f := int64(v)
				s.LatestEmptyBucketTimestamp = &f
			}

		case "latest_record_timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LatestRecordTimestamp = &value
			case float64:
				f := int64(v)
				s.LatestRecordTimestamp = &f
			}

		case "latest_sparse_bucket_timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LatestSparseBucketTimestamp = &value
			case float64:
				f := int64(v)
				s.LatestSparseBucketTimestamp = &f
			}

		case "log_time":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LogTime = &value
			case float64:
				f := int64(v)
				s.LogTime = &f
			}

		case "missing_field_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MissingFieldCount = value
			case float64:
				f := int64(v)
				s.MissingFieldCount = f
			}

		case "out_of_order_timestamp_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.OutOfOrderTimestampCount = value
			case float64:
				f := int64(v)
				s.OutOfOrderTimestampCount = f
			}

		case "processed_field_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ProcessedFieldCount = value
			case float64:
				f := int64(v)
				s.ProcessedFieldCount = f
			}

		case "processed_record_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ProcessedRecordCount = value
			case float64:
				f := int64(v)
				s.ProcessedRecordCount = f
			}

		case "sparse_bucket_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SparseBucketCount = value
			case float64:
				f := int64(v)
				s.SparseBucketCount = f
			}

		}
	}
	return nil
}

// NewDataCounts returns a DataCounts.
func NewDataCounts() *DataCounts {
	r := &DataCounts{}

	return r
}
