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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// IngestDocumentSimulation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cbfcc73d01310bed2a480ec35aaef98138b598e5/specification/simulate/ingest/SimulateIngestResponse.ts#L35-L78
type IngestDocumentSimulation struct {
	// Error Any error resulting from simulatng ingest on this doc. This can be an error
	// generated by
	// executing a processor, or a mapping validation error when simulating indexing
	// the resulting
	// doc.
	Error *ErrorCause `json:"error,omitempty"`
	// ExecutedPipelines A list of the names of the pipelines executed on this document.
	ExecutedPipelines []string `json:"executed_pipelines"`
	// Id_ Identifier for the document.
	Id_ string `json:"_id"`
	// IgnoredFields A list of the fields that would be ignored at the indexing step. For example,
	// a field whose
	// value is larger than the allowed limit would make it through all of the
	// pipelines, but
	// would not be indexed into Elasticsearch.
	IgnoredFields []map[string]string `json:"ignored_fields,omitempty"`
	// Index_ Name of the index that the document would be indexed into if this were not a
	// simulation.
	Index_                   string            `json:"_index"`
	IngestDocumentSimulation map[string]string `json:"-"`
	// Source_ JSON body for the document.
	Source_  map[string]json.RawMessage `json:"_source"`
	Version_ StringifiedVersionNumber   `json:"_version"`
}

func (s *IngestDocumentSimulation) UnmarshalJSON(data []byte) error {

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

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}

		case "executed_pipelines":
			if err := dec.Decode(&s.ExecutedPipelines); err != nil {
				return fmt.Errorf("%s | %w", "ExecutedPipelines", err)
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "ignored_fields":
			if err := dec.Decode(&s.IgnoredFields); err != nil {
				return fmt.Errorf("%s | %w", "IgnoredFields", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "_source":
			if s.Source_ == nil {
				s.Source_ = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Source_); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return fmt.Errorf("%s | %w", "Version_", err)
			}

		default:

			if key, ok := t.(string); ok {
				if s.IngestDocumentSimulation == nil {
					s.IngestDocumentSimulation = make(map[string]string, 0)
				}
				raw := new(string)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "IngestDocumentSimulation", err)
				}
				s.IngestDocumentSimulation[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s IngestDocumentSimulation) MarshalJSON() ([]byte, error) {
	type opt IngestDocumentSimulation
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.IngestDocumentSimulation {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "IngestDocumentSimulation")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewIngestDocumentSimulation returns a IngestDocumentSimulation.
func NewIngestDocumentSimulation() *IngestDocumentSimulation {
	r := &IngestDocumentSimulation{
		IngestDocumentSimulation: make(map[string]string),
		Source_:                  make(map[string]json.RawMessage),
	}

	return r
}

// false
