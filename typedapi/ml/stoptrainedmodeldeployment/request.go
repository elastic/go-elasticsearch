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
// https://github.com/elastic/elasticsearch-specification/tree/224e96968e3ab27c2d1d33f015783b44ed183c1f

package stoptrainedmodeldeployment

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Request holds the request body struct for the package stoptrainedmodeldeployment
//
// https://github.com/elastic/elasticsearch-specification/blob/224e96968e3ab27c2d1d33f015783b44ed183c1f/specification/ml/stop_trained_model_deployment/MlStopTrainedModelDeploymentRequest.ts#L23-L83
type Request struct {
	// AllowNoMatch Specifies what to do when the request: contains wildcard expressions and
	// there are no deployments that match;
	// contains the  `_all` string or no identifiers and there are no matches; or
	// contains wildcard expressions and
	// there are only partial matches. By default, it returns an empty array when
	// there are no matches and the subset of results when there are partial
	// matches.
	// If `false`, the request returns a 404 status code when there are no matches
	// or only partial matches.
	AllowNoMatch *bool `json:"allow_no_match,omitempty"`
	// Force Forcefully stops the deployment, even if it is used by ingest pipelines. You
	// can't use these pipelines until you
	// restart the model deployment.
	Force *bool `json:"force,omitempty"`
	// Id If provided, must be the same identifier as in the path.
	Id *string `json:"id,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Stoptrainedmodeldeployment request: %w", err)
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

		case "allow_no_match":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowNoMatch", err)
				}
				s.AllowNoMatch = &value
			case bool:
				s.AllowNoMatch = &v
			}

		case "force":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Force", err)
				}
				s.Force = &value
			case bool:
				s.Force = &v
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		}
	}
	return nil
}
