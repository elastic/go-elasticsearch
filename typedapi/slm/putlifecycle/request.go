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

package putlifecycle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putlifecycle
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/slm/put_lifecycle/PutSnapshotLifecycleRequest.ts#L26-L89
type Request struct {

	// Config Configuration for each snapshot created by the policy.
	Config *types.Configuration `json:"config,omitempty"`
	// Name Name automatically assigned to each snapshot created by the policy. Date math
	// is supported. To prevent conflicting snapshot names, a UUID is automatically
	// appended to each snapshot name.
	Name *string `json:"name,omitempty"`
	// Repository Repository used to store snapshots created by this policy. This repository
	// must exist prior to the policy’s creation. You can create a repository using
	// the snapshot repository API.
	Repository *string `json:"repository,omitempty"`
	// Retention Retention rules used to retain and delete snapshots created by the policy.
	Retention *types.Retention `json:"retention,omitempty"`
	// Schedule Periodic or absolute schedule at which the policy creates snapshots. SLM
	// applies schedule changes immediately.
	Schedule *string `json:"schedule,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Putlifecycle request: %w", err)
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

		case "config":
			if err := dec.Decode(&s.Config); err != nil {
				return fmt.Errorf("%s | %w", "Config", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "repository":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Repository", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Repository = &o

		case "retention":
			if err := dec.Decode(&s.Retention); err != nil {
				return fmt.Errorf("%s | %w", "Retention", err)
			}

		case "schedule":
			if err := dec.Decode(&s.Schedule); err != nil {
				return fmt.Errorf("%s | %w", "Schedule", err)
			}

		}
	}
	return nil
}
