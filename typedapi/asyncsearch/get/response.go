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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package get

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package get
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/async_search/get/AsyncSearchGetResponse.ts#L22-L24
type Response struct {

	// CompletionTime Indicates when the async search completed.
	// It is present only when the search has completed.
	CompletionTime         types.DateTime `json:"completion_time,omitempty"`
	CompletionTimeInMillis *int64         `json:"completion_time_in_millis,omitempty"`
	// ExpirationTime Indicates when the async search will expire.
	ExpirationTime         types.DateTime `json:"expiration_time,omitempty"`
	ExpirationTimeInMillis int64          `json:"expiration_time_in_millis"`
	Id                     *string        `json:"id,omitempty"`
	// IsPartial When the query is no longer running, this property indicates whether the
	// search failed or was successfully completed on all shards.
	// While the query is running, `is_partial` is always set to `true`.
	IsPartial bool `json:"is_partial"`
	// IsRunning Indicates whether the search is still running or has completed.
	//
	// > info
	// > If the search failed after some shards returned their results or the node
	// that is coordinating the async search dies, results may be partial even
	// though `is_running` is `false`.
	IsRunning         bool              `json:"is_running"`
	Response          types.AsyncSearch `json:"response"`
	StartTime         types.DateTime    `json:"start_time,omitempty"`
	StartTimeInMillis int64             `json:"start_time_in_millis"`
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

		case "completion_time":
			if err := dec.Decode(&s.CompletionTime); err != nil {
				return fmt.Errorf("%s | %w", "CompletionTime", err)
			}

		case "completion_time_in_millis":
			if err := dec.Decode(&s.CompletionTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "CompletionTimeInMillis", err)
			}

		case "expiration_time":
			if err := dec.Decode(&s.ExpirationTime); err != nil {
				return fmt.Errorf("%s | %w", "ExpirationTime", err)
			}

		case "expiration_time_in_millis":
			if err := dec.Decode(&s.ExpirationTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "ExpirationTimeInMillis", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "is_partial":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsPartial", err)
				}
				s.IsPartial = value
			case bool:
				s.IsPartial = v
			}

		case "is_running":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsRunning", err)
				}
				s.IsRunning = value
			case bool:
				s.IsRunning = v
			}

		case "response":
			if err := dec.Decode(&s.Response); err != nil {
				return fmt.Errorf("%s | %w", "Response", err)
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return fmt.Errorf("%s | %w", "StartTime", err)
			}

		case "start_time_in_millis":
			if err := dec.Decode(&s.StartTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartTimeInMillis", err)
			}

		}
	}
	return nil
}
