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

package restore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package restore
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/restore/SnapshotRestoreRequest.ts#L25-L78
type Request struct {
	FeatureStates       []string             `json:"feature_states,omitempty"`
	IgnoreIndexSettings []string             `json:"ignore_index_settings,omitempty"`
	IgnoreUnavailable   *bool                `json:"ignore_unavailable,omitempty"`
	IncludeAliases      *bool                `json:"include_aliases,omitempty"`
	IncludeGlobalState  *bool                `json:"include_global_state,omitempty"`
	IndexSettings       *types.IndexSettings `json:"index_settings,omitempty"`
	Indices             []string             `json:"indices,omitempty"`
	Partial             *bool                `json:"partial,omitempty"`
	RenamePattern       *string              `json:"rename_pattern,omitempty"`
	RenameReplacement   *string              `json:"rename_replacement,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Restore request: %w", err)
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

		case "ignore_index_settings":
			if err := dec.Decode(&s.IgnoreIndexSettings); err != nil {
				return fmt.Errorf("%s | %w", "IgnoreIndexSettings", err)
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

		case "include_aliases":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeAliases", err)
				}
				s.IncludeAliases = &value
			case bool:
				s.IncludeAliases = &v
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

		case "index_settings":
			if err := dec.Decode(&s.IndexSettings); err != nil {
				return fmt.Errorf("%s | %w", "IndexSettings", err)
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

		case "rename_pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RenamePattern", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RenamePattern = &o

		case "rename_replacement":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RenameReplacement", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RenameReplacement = &o

		}
	}
	return nil
}
