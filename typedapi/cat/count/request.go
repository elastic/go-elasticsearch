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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package count

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Request holds the request body struct for the package count
//
// https://github.com/elastic/elasticsearch-specification/blob/b1811e10a0722431d79d1c234dd412ff47d8656f/specification/cat/count/CatCountRequest.ts#L23-L83
type Request struct {
	// ProjectRouting Specifies a subset of projects to target using project
	// metadata tags in a subset of Lucene query syntax.
	// Allowed Lucene queries: the _alias tag and a single value (possibly
	// wildcarded).
	// Examples:
	//  _alias:my-project
	//  _alias:_origin
	//  _alias:*pr*
	// Supported in serverless only.
	ProjectRouting *string `json:"project_routing,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Count request: %w", err)
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

		case "project_routing":
			if err := dec.Decode(&s.ProjectRouting); err != nil {
				return fmt.Errorf("%s | %w", "ProjectRouting", err)
			}

		}
	}
	return nil
}
