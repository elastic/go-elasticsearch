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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// The final result of a completed reindex operation, as stored in the task
// result. This is the serialized form of `BulkByScrollResponse`.
//
// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/_types/Reindex.ts#L90-L151
type ReindexTaskResult struct {
	// Batches The number of scroll responses pulled back by the reindex.
	Batches *int64 `json:"batches,omitempty"`
	// Created The number of documents that were successfully created.
	Created *int64 `json:"created,omitempty"`
	// Deleted The number of documents that were successfully deleted.
	Deleted *int64 `json:"deleted,omitempty"`
	// Failures Any failures encountered during the reindex. If non-empty, the reindex ended
	// because of these failures.
	Failures []BulkIndexByScrollFailure `json:"failures,omitempty"`
	// Noops The number of documents that were ignored because the script returned a
	// `noop` value for `ctx.op`.
	Noops *int64 `json:"noops,omitempty"`
	// RequestsPerSecond The number of requests per second effectively executed during the reindex.
	RequestsPerSecond *float32 `json:"requests_per_second,omitempty"`
	// Retries The number of retries attempted by reindex.
	Retries *Retries `json:"retries,omitempty"`
	// ThrottledMillis Number of milliseconds the request slept to conform to `requests_per_second`.
	ThrottledMillis *int64 `json:"throttled_millis,omitempty"`
	// ThrottledUntilMillis This field should always be equal to zero in a completed reindex result.
	ThrottledUntilMillis *int64 `json:"throttled_until_millis,omitempty"`
	// TimedOut Whether any of the requests executed during the reindex timed out.
	TimedOut *bool `json:"timed_out,omitempty"`
	// Took The total milliseconds the entire operation took.
	Took *int64 `json:"took,omitempty"`
	// Total The number of documents that were successfully processed.
	Total *int64 `json:"total,omitempty"`
	// Updated The number of documents that were successfully updated.
	Updated *int64 `json:"updated,omitempty"`
	// VersionConflicts The number of version conflicts that occurred.
	VersionConflicts *int64 `json:"version_conflicts,omitempty"`
}

func (s *ReindexTaskResult) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.Failures); err != nil {
				return fmt.Errorf("%s | %w", "Failures", err)
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

// NewReindexTaskResult returns a ReindexTaskResult.
func NewReindexTaskResult() *ReindexTaskResult {
	r := &ReindexTaskResult{}

	return r
}
