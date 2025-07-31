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

package createcrossclusterapikey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package createcrossclusterapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/create_cross_cluster_api_key/CreateCrossClusterApiKeyRequest.ts#L25-L80
type Request struct {

	// Access The access to be granted to this API key.
	// The access is composed of permissions for cross-cluster search and
	// cross-cluster replication.
	// At least one of them must be specified.
	//
	// NOTE: No explicit privileges should be specified for either search or
	// replication access.
	// The creation process automatically converts the access specification to a
	// role descriptor which has relevant privileges assigned accordingly.
	Access types.Access `json:"access"`
	// Expiration Expiration time for the API key.
	// By default, API keys never expire.
	Expiration types.Duration `json:"expiration,omitempty"`
	// Metadata Arbitrary metadata that you want to associate with the API key.
	// It supports nested data structure.
	// Within the metadata object, keys beginning with `_` are reserved for system
	// usage.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// Name Specifies the name for this API key.
	Name string `json:"name"`
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
		return nil, fmt.Errorf("could not deserialise json into Createcrossclusterapikey request: %w", err)
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

		case "access":
			if err := dec.Decode(&s.Access); err != nil {
				return fmt.Errorf("%s | %w", "Access", err)
			}

		case "expiration":
			if err := dec.Decode(&s.Expiration); err != nil {
				return fmt.Errorf("%s | %w", "Expiration", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		}
	}
	return nil
}
