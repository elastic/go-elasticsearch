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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

package create

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package create
//
// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/snapshot/create/SnapshotCreateRequest.ts#L24-L92
type Request struct {

	// FeatureStates Feature states to include in the snapshot. Each feature state includes one or
	// more system indices containing related data. You can view a list of eligible
	// features using the get features API. If `include_global_state` is `true`, all
	// current feature states are included by default. If `include_global_state` is
	// `false`, no feature states are included by default.
	FeatureStates []string `json:"feature_states,omitempty"`
	// IgnoreUnavailable If `true`, the request ignores data streams and indices in `indices` that are
	// missing or closed. If `false`, the request returns an error for any data
	// stream or index that is missing or closed.
	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`
	// IncludeGlobalState If `true`, the current cluster state is included in the snapshot. The cluster
	// state includes persistent cluster settings, composable index templates,
	// legacy index templates, ingest pipelines, and ILM policies. It also includes
	// data stored in system indices, such as Watches and task records (configurable
	// via `feature_states`).
	IncludeGlobalState *bool `json:"include_global_state,omitempty"`
	// Indices Data streams and indices to include in the snapshot. Supports multi-target
	// syntax. Includes all data streams and indices by default.
	Indices []string `json:"indices,omitempty"`
	// Metadata Optional metadata for the snapshot. May have any contents. Must be less than
	// 1024 bytes. This map is not automatically generated by Elasticsearch.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// Partial If `true`, allows restoring a partial snapshot of indices with unavailable
	// shards. Only shards that were successfully included in the snapshot will be
	// restored. All missing shards will be recreated as empty. If `false`, the
	// entire restore operation will fail if one or more indices included in the
	// snapshot do not have all primary shards available.
	Partial *bool `json:"partial,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Create request: %w", err)
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

		case "feature_states":
			if err := dec.Decode(&s.FeatureStates); err != nil {
				return fmt.Errorf("%s | %w", "FeatureStates", err)
			}

		case "ignore_unavailable":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreUnavailable", err)
				}
				s.IgnoreUnavailable = &value
			case bool:
				s.IgnoreUnavailable = &v
			}

		case "include_global_state":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeGlobalState", err)
				}
				s.IncludeGlobalState = &value
			case bool:
				s.IncludeGlobalState = &v
			}

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Indices", err)
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return fmt.Errorf("%s | %w", "Indices", err)
				}
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "partial":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Partial", err)
				}
				s.Partial = &value
			case bool:
				s.Partial = &v
			}

		}
	}
	return nil
}
