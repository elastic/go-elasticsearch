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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"

	"fmt"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// DocumentSimulation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ingest/simulate/types.ts#L47-L60
type DocumentSimulation struct {
	DocumentSimulation map[string]string          `json:"-"`
	Id_                string                     `json:"_id"`
	Index_             string                     `json:"_index"`
	Ingest_            SimulateIngest             `json:"_ingest"`
	Routing_           *string                    `json:"_routing,omitempty"`
	Source_            map[string]json.RawMessage `json:"_source"`
	VersionType_       *versiontype.VersionType   `json:"_version_type,omitempty"`
	Version_           StringifiedVersionNumber   `json:"_version,omitempty"`
}

func (s *DocumentSimulation) UnmarshalJSON(data []byte) error {

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

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return err
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return err
			}

		case "_ingest":
			if err := dec.Decode(&s.Ingest_); err != nil {
				return err
			}

		case "_routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Routing_ = &o

		case "_source":
			if s.Source_ == nil {
				s.Source_ = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Source_); err != nil {
				return err
			}

		case "_version_type":
			if err := dec.Decode(&s.VersionType_); err != nil {
				return err
			}

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return err
			}

		default:

			if key, ok := t.(string); ok {
				if s.DocumentSimulation == nil {
					s.DocumentSimulation = make(map[string]string, 0)
				}
				raw := new(string)
				if err := dec.Decode(&raw); err != nil {
					return err
				}
				s.DocumentSimulation[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DocumentSimulation) MarshalJSON() ([]byte, error) {
	type opt DocumentSimulation
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.DocumentSimulation {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "DocumentSimulation")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewDocumentSimulation returns a DocumentSimulation.
func NewDocumentSimulation() *DocumentSimulation {
	r := &DocumentSimulation{
		DocumentSimulation: make(map[string]string, 0),
		Source_:            make(map[string]json.RawMessage, 0),
	}

	return r
}
