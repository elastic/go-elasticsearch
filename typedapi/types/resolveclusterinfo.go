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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ResolveClusterInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/resolve_cluster/ResolveClusterResponse.ts#L29-L55
type ResolveClusterInfo struct {
	// Connected Whether the remote cluster is connected to the local (querying) cluster.
	Connected bool `json:"connected"`
	// Error Provides error messages that are likely to occur if you do a search with this
	// index expression
	// on the specified cluster (for example, lack of security privileges to query
	// an index).
	Error *string `json:"error,omitempty"`
	// MatchingIndices Whether the index expression provided in the request matches any indices,
	// aliases or data streams
	// on the cluster.
	MatchingIndices *bool `json:"matching_indices,omitempty"`
	// SkipUnavailable The `skip_unavailable` setting for a remote cluster.
	SkipUnavailable bool `json:"skip_unavailable"`
	// Version Provides version information about the cluster.
	Version *ElasticsearchVersionMinInfo `json:"version,omitempty"`
}

func (s *ResolveClusterInfo) UnmarshalJSON(data []byte) error {

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

		case "connected":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Connected", err)
				}
				s.Connected = value
			case bool:
				s.Connected = v
			}

		case "error":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Error = &o

		case "matching_indices":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MatchingIndices", err)
				}
				s.MatchingIndices = &value
			case bool:
				s.MatchingIndices = &v
			}

		case "skip_unavailable":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SkipUnavailable", err)
				}
				s.SkipUnavailable = value
			case bool:
				s.SkipUnavailable = v
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewResolveClusterInfo returns a ResolveClusterInfo.
func NewResolveClusterInfo() *ResolveClusterInfo {
	r := &ResolveClusterInfo{}

	return r
}
