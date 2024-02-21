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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package reloadsecuresettings

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package reloadsecuresettings
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/nodes/reload_secure_settings/ReloadSecureSettingsResponse.ts#L30-L32
type Response struct {
	ClusterName string `json:"cluster_name"`
	// NodeStats Contains statistics about the number of nodes selected by the request’s node
	// filters.
	NodeStats *types.NodeStatistics             `json:"_nodes,omitempty"`
	Nodes     map[string]types.NodeReloadResult `json:"nodes"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Nodes: make(map[string]types.NodeReloadResult, 0),
	}
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

		case "cluster_name":
			if err := dec.Decode(&s.ClusterName); err != nil {
				return err
			}

		case "_nodes":
			if err := dec.Decode(&s.NodeStats); err != nil {
				return err
			}

		case "nodes":
			if s.Nodes == nil {
				s.Nodes = make(map[string]types.NodeReloadResult, 0)
			}
			if err := dec.Decode(&s.Nodes); err != nil {
				return err
			}

		}
	}
	return nil
}
