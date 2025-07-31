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

package syncjobpost

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syncjobtriggermethod"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syncjobtype"
)

// Request holds the request body struct for the package syncjobpost
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/sync_job_post/SyncJobPostRequest.ts#L23-L51
type Request struct {

	// Id The id of the associated connector
	Id            string                                     `json:"id"`
	JobType       *syncjobtype.SyncJobType                   `json:"job_type,omitempty"`
	TriggerMethod *syncjobtriggermethod.SyncJobTriggerMethod `json:"trigger_method,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Syncjobpost request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "job_type":
			if err := dec.Decode(&s.JobType); err != nil {
				return fmt.Errorf("%s | %w", "JobType", err)
			}

		case "trigger_method":
			if err := dec.Decode(&s.TriggerMethod); err != nil {
				return fmt.Errorf("%s | %w", "TriggerMethod", err)
			}

		}
	}
	return nil
}
