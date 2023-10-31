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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Configuration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/slm/_types/SnapshotLifecycle.ts#L99-L129
type Configuration struct {
	// FeatureStates A list of feature states to be included in this snapshot. A list of features
	// available for inclusion in the snapshot and their descriptions be can be
	// retrieved using the get features API.
	// Each feature state includes one or more system indices containing data
	// necessary for the function of that feature. Providing an empty array will
	// include no feature states in the snapshot, regardless of the value of
	// include_global_state. By default, all available feature states will be
	// included in the snapshot if include_global_state is true, or no feature
	// states if include_global_state is false.
	FeatureStates []string `json:"feature_states,omitempty"`
	// IgnoreUnavailable If false, the snapshot fails if any data stream or index in indices is
	// missing or closed. If true, the snapshot ignores missing or closed data
	// streams and indices.
	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`
	// IncludeGlobalState If true, the current global state is included in the snapshot.
	IncludeGlobalState *bool `json:"include_global_state,omitempty"`
	// Indices A comma-separated list of data streams and indices to include in the
	// snapshot. Multi-index syntax is supported.
	// By default, a snapshot includes all data streams and indices in the cluster.
	// If this argument is provided, the snapshot only includes the specified data
	// streams and clusters.
	Indices []string `json:"indices,omitempty"`
	// Metadata Attaches arbitrary metadata to the snapshot, such as a record of who took the
	// snapshot, why it was taken, or any other useful data. Metadata must be less
	// than 1024 bytes.
	Metadata Metadata `json:"metadata,omitempty"`
	// Partial If false, the entire snapshot will fail if one or more indices included in
	// the snapshot do not have all primary shards available.
	Partial *bool `json:"partial,omitempty"`
}

func (s *Configuration) UnmarshalJSON(data []byte) error {

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
				return err
			}

		case "ignore_unavailable":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreUnavailable = &value
			case bool:
				s.IgnoreUnavailable = &v
			}

		case "include_global_state":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
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
					return err
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return err
				}
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return err
			}

		case "partial":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Partial = &value
			case bool:
				s.Partial = &v
			}

		}
	}
	return nil
}

// NewConfiguration returns a Configuration.
func NewConfiguration() *Configuration {
	r := &Configuration{}

	return r
}
