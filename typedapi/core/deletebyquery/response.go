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

package deletebyquery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package deletebyquery
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/delete_by_query/DeleteByQueryResponse.ts#L27-L93
type Response struct {
	// Batches The number of scroll responses pulled back by the delete by query.
	Batches *int64 `json:"batches,omitempty"`
	// Deleted The number of documents that were successfully deleted.
	Deleted *int64 `json:"deleted,omitempty"`
	// Failures An array of failures if there were any unrecoverable errors during the
	// process. If this array is not empty, the request ended abnormally because of
	// those failures. Delete by query is implemented using batches and any failures
	// cause the entire process to end but all failures in the current batch are
	// collected into the array. You can use the `conflicts` option to prevent
	// reindex from ending on version conflicts.
	Failures []types.BulkByPaginatedSearchFailure `json:"failures,omitempty"`
	// Noops This field is always equal to zero for delete by query. It exists only so
	// that delete by query, update by query, and reindex APIs return responses with
	// the same structure.
	Noops *int64 `json:"noops,omitempty"`
	// RequestsPerSecond The number of requests per second effectively run during the delete by query.
	RequestsPerSecond *float32 `json:"requests_per_second,omitempty"`
	// Retries The number of retries attempted by delete by query. `bulk` is the number of
	// bulk actions retried. `search` is the number of search actions retried.
	Retries *types.Retries `json:"retries,omitempty"`
	SliceId *int           `json:"slice_id,omitempty"`
	// Slices Status of each slice if the delete by query was sliced
	Slices    []types.ReindexStatus `json:"slices,omitempty"`
	Task      *string               `json:"task,omitempty"`
	Throttled types.Duration        `json:"throttled,omitempty"`
	// ThrottledMillis The number of milliseconds the request slept to conform to
	// `requests_per_second`.
	ThrottledMillis *int64         `json:"throttled_millis,omitempty"`
	ThrottledUntil  types.Duration `json:"throttled_until,omitempty"`
	// ThrottledUntilMillis This field should always be equal to zero in a `_delete_by_query` response.
	// It has meaning only when using the task API, where it indicates the next time
	// (in milliseconds since epoch) a throttled request will be run again in order
	// to conform to `requests_per_second`.
	ThrottledUntilMillis *int64 `json:"throttled_until_millis,omitempty"`
	// TimedOut If `true`, some requests run during the delete by query operation timed out.
	TimedOut *bool `json:"timed_out,omitempty"`
	// Took The number of milliseconds from start to end of the whole operation.
	Took *int64 `json:"took,omitempty"`
	// Total The number of documents that were successfully processed.
	Total *int64 `json:"total,omitempty"`
	// VersionConflicts The number of version conflicts that the delete by query hit.
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

		case "throttled":
			if err := dec.Decode(&s.Throttled); err != nil {
				return fmt.Errorf("%s | %w", "Throttled", err)
			}

		case "throttled_millis":
			if err := dec.Decode(&s.ThrottledMillis); err != nil {
				return fmt.Errorf("%s | %w", "ThrottledMillis", err)
			}

		case "throttled_until":
			if err := dec.Decode(&s.ThrottledUntil); err != nil {
				return fmt.Errorf("%s | %w", "ThrottledUntil", err)
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
