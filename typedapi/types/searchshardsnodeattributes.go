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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"
)

// SearchShardsNodeAttributes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search_shards/SearchShardsResponse.ts#L42-L60
type SearchShardsNodeAttributes struct {
	// Attributes Lists node attributes.
	Attributes map[string]string `json:"attributes"`
	// EphemeralId The ephemeral ID of the node.
	EphemeralId     string `json:"ephemeral_id"`
	ExternalId      string `json:"external_id"`
	MaxIndexVersion int    `json:"max_index_version"`
	MinIndexVersion int    `json:"min_index_version"`
	// Name The human-readable identifier of the node.
	Name  string              `json:"name"`
	Roles []noderole.NodeRole `json:"roles"`
	// TransportAddress The host and port where transport HTTP connections are accepted.
	TransportAddress string `json:"transport_address"`
	Version          string `json:"version"`
}

func (s *SearchShardsNodeAttributes) UnmarshalJSON(data []byte) error {

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

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return fmt.Errorf("%s | %w", "Attributes", err)
			}

		case "ephemeral_id":
			if err := dec.Decode(&s.EphemeralId); err != nil {
				return fmt.Errorf("%s | %w", "EphemeralId", err)
			}

		case "external_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ExternalId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ExternalId = o

		case "max_index_version":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxIndexVersion", err)
				}
				s.MaxIndexVersion = value
			case float64:
				f := int(v)
				s.MaxIndexVersion = f
			}

		case "min_index_version":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinIndexVersion", err)
				}
				s.MinIndexVersion = value
			case float64:
				f := int(v)
				s.MinIndexVersion = f
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return fmt.Errorf("%s | %w", "TransportAddress", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewSearchShardsNodeAttributes returns a SearchShardsNodeAttributes.
func NewSearchShardsNodeAttributes() *SearchShardsNodeAttributes {
	r := &SearchShardsNodeAttributes{
		Attributes: make(map[string]string),
	}

	return r
}
