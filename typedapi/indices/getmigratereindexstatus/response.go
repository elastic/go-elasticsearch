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

package getmigratereindexstatus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package getmigratereindexstatus
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/get_migrate_reindex_status/MigrateGetReindexStatusResponse.ts#L23-L36
type Response struct {
	Complete                     bool                     `json:"complete"`
	Errors                       []types.StatusError      `json:"errors"`
	Exception                    *string                  `json:"exception,omitempty"`
	InProgress                   []types.StatusInProgress `json:"in_progress"`
	Pending                      int                      `json:"pending"`
	StartTime                    types.DateTime           `json:"start_time,omitempty"`
	StartTimeMillis              int64                    `json:"start_time_millis"`
	Successes                    int                      `json:"successes"`
	TotalIndicesInDataStream     int                      `json:"total_indices_in_data_stream"`
	TotalIndicesRequiringUpgrade int                      `json:"total_indices_requiring_upgrade"`
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

		case "complete":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Complete", err)
				}
				s.Complete = value
			case bool:
				s.Complete = v
			}

		case "errors":
			if err := dec.Decode(&s.Errors); err != nil {
				return fmt.Errorf("%s | %w", "Errors", err)
			}

		case "exception":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Exception", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Exception = &o

		case "in_progress":
			if err := dec.Decode(&s.InProgress); err != nil {
				return fmt.Errorf("%s | %w", "InProgress", err)
			}

		case "pending":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Pending", err)
				}
				s.Pending = value
			case float64:
				f := int(v)
				s.Pending = f
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return fmt.Errorf("%s | %w", "StartTime", err)
			}

		case "start_time_millis":
			if err := dec.Decode(&s.StartTimeMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartTimeMillis", err)
			}

		case "successes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Successes", err)
				}
				s.Successes = value
			case float64:
				f := int(v)
				s.Successes = f
			}

		case "total_indices_in_data_stream":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalIndicesInDataStream", err)
				}
				s.TotalIndicesInDataStream = value
			case float64:
				f := int(v)
				s.TotalIndicesInDataStream = f
			}

		case "total_indices_requiring_upgrade":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalIndicesRequiringUpgrade", err)
				}
				s.TotalIndicesRequiringUpgrade = value
			case float64:
				f := int(v)
				s.TotalIndicesRequiringUpgrade = f
			}

		}
	}
	return nil
}
