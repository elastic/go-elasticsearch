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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package reindex

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package reindex
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/reindex/ReindexResponse.ts#L27-L97
type Response struct {
	// Batches The number of scroll responses that were pulled back by the reindex.
	Batches *int64 `json:"batches,omitempty"`
	// Created The number of documents that were successfully created.
	Created *int64 `json:"created,omitempty"`
	// Deleted The number of documents that were successfully deleted.
	Deleted *int64 `json:"deleted,omitempty"`
	// Failures If there were any unrecoverable errors during the process, it is an array of
	// those failures. If this array is not empty, the request ended because of
	// those failures. Reindex is implemented using batches and any failure causes
	// the entire process to end but all failures in the current batch are collected
	// into the array. You can use the `conflicts` option to prevent the reindex
	// from ending on version conflicts.
	Failures []types.BulkByPaginatedSearchFailure `json:"failures,omitempty"`
	// Noops The number of documents that were ignored because the script used for the
	// reindex returned a `noop` value for `ctx.op`.
	Noops *int64 `json:"noops,omitempty"`
	// RequestsPerSecond The number of requests per second effectively run during the reindex.
	RequestsPerSecond *float32 `json:"requests_per_second,omitempty"`
	// Retries The number of retries attempted by reindex.
	Retries *types.Retries `json:"retries,omitempty"`
	SliceId *int           `json:"slice_id,omitempty"`
	// Slices Status of each slice if the reindex was sliced
	Slices []types.ReindexStatus `json:"slices,omitempty"`
	Task   *string               `json:"task,omitempty"`
	// ThrottledMillis The number of milliseconds the request slept to conform to
	// `requests_per_second`.
	ThrottledMillis *int64 `json:"throttled_millis,omitempty"`
	// ThrottledUntilMillis This field should always be equal to zero in a reindex response. It has
	// meaning only when using the task API, where it indicates the next time (in
	// milliseconds since epoch) that a throttled request will be run again in order
	// to conform to `requests_per_second`.
	ThrottledUntilMillis *int64 `json:"throttled_until_millis,omitempty"`
	// TimedOut If any of the requests that ran during the reindex timed out, it is `true`.
	TimedOut *bool `json:"timed_out,omitempty"`
	// Took The total milliseconds the entire operation took.
	Took *int64 `json:"took,omitempty"`
	// Total The number of documents that were successfully processed.
	Total *int64 `json:"total,omitempty"`
	// Updated The number of documents that were successfully updated. That is to say, a
	// document with the same ID already existed before the reindex updated it.
	Updated *int64 `json:"updated,omitempty"`
	// VersionConflicts The number of version conflicts that occurred.
	VersionConflicts *int64 `json:"version_conflicts,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "batches":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Batches", err)
				}
				s.Batches = &value
			case float64:
				f := int64(v)
				s.Batches = &f
			}

		case "created":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Created", err)
				}
				s.Created = &value
			case float64:
				f := int64(v)
				s.Created = &f
			}

		case "deleted":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Deleted", err)
				}
				s.Deleted = &value
			case float64:
				f := int64(v)
				s.Deleted = &f
			}

		case "failures":
			messageArray := []json.RawMessage{}
			if err := dec.Decode(&messageArray); err != nil {
				return fmt.Errorf("%s | %w", "Failures", err)
			}
		failures_field:
			for _, message := range messageArray {
				keyDec := json.NewDecoder(bytes.NewReader(message))
				for {
					t, err := keyDec.Token()
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return fmt.Errorf("%s | %w", "Failures", err)
					}

					switch t {

					case "cause", "id":
						o := types.NewBulkIndexByScrollFailure()
						localDec := json.NewDecoder(bytes.NewReader(message))
						if err := localDec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Failures", err)
						}
						s.Failures = append(s.Failures, o)
						continue failures_field

					case "node", "primary", "reason", "shard":
						o := types.NewShardFailure()
						localDec := json.NewDecoder(bytes.NewReader(message))
						if err := localDec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Failures", err)
						}
						s.Failures = append(s.Failures, o)
						continue failures_field

					}
				}

				var o any
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Failures", err)
				}
				s.Failures = append(s.Failures, o)
			}

		case "noops":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Noops", err)
				}
				s.Noops = &value
			case float64:
				f := int64(v)
				s.Noops = &f
			}

		case "requests_per_second":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "RequestsPerSecond", err)
				}
				f := float32(value)
				s.RequestsPerSecond = &f
			case float64:
				f := float32(v)
				s.RequestsPerSecond = &f
			}

		case "retries":
			if err := dec.Decode(&s.Retries); err != nil {
				return fmt.Errorf("%s | %w", "Retries", err)
			}

		case "slice_id":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SliceId", err)
				}
				s.SliceId = &value
			case float64:
				f := int(v)
				s.SliceId = &f
			}

		case "slices":
			if err := dec.Decode(&s.Slices); err != nil {
				return fmt.Errorf("%s | %w", "Slices", err)
			}

		case "task":
			if err := dec.Decode(&s.Task); err != nil {
				return fmt.Errorf("%s | %w", "Task", err)
			}

		case "throttled_millis":
			if err := dec.Decode(&s.ThrottledMillis); err != nil {
				return fmt.Errorf("%s | %w", "ThrottledMillis", err)
			}

		case "throttled_until_millis":
			if err := dec.Decode(&s.ThrottledUntilMillis); err != nil {
				return fmt.Errorf("%s | %w", "ThrottledUntilMillis", err)
			}

		case "timed_out":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TimedOut", err)
				}
				s.TimedOut = &value
			case bool:
				s.TimedOut = &v
			}

		case "took":
			if err := dec.Decode(&s.Took); err != nil {
				return fmt.Errorf("%s | %w", "Took", err)
			}

		case "total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = &value
			case float64:
				f := int64(v)
				s.Total = &f
			}

		case "updated":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Updated", err)
				}
				s.Updated = &value
			case float64:
				f := int64(v)
				s.Updated = &f
			}

		case "version_conflicts":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "VersionConflicts", err)
				}
				s.VersionConflicts = &value
			case float64:
				f := int64(v)
				s.VersionConflicts = &f
			}

		}
	}
	return nil
}
