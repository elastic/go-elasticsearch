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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package putdatasource

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package putdatasource
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/esql/put_data_source/PutDataSourceRequest.ts#L26-L79
type Request struct {
	// Description A free-text description of the data source.
	Description *string `json:"description,omitempty"`
	// Settings Type-specific connection and authentication settings. For `s3`, connection
	// settings include `region` and `endpoint`. Authentication settings include
	// `auth` and the credentials required by the selected authentication method.
	Settings map[string]json.RawMessage `json:"settings,omitempty"`
	// Type The data source type. Currently, `s3` is supported. The value must be
	// lowercase and contain no whitespace.
	Type string `json:"type"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Settings: make(map[string]json.RawMessage, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putdatasource request: %w", err)
	}

	return &req, nil
}
