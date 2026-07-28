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
// https://github.com/elastic/elasticsearch-specification/tree/7fd0bd13eaf28bd179fc906f57da09e852eb818e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Represents a dataset definition stored in cluster state. A dataset is a named
// reference to external data that participates in the index namespace alongside
// indices, aliases, and views. Datasets inherit credentials from their
// referenced data source at query time.
//
// https://github.com/elastic/elasticsearch-specification/blob/7fd0bd13eaf28bd179fc906f57da09e852eb818e/specification/esql/_types/types.ts#L147-L175
type ESQLDataset struct {
	// DataSource The name of the referenced data source.
	DataSource string `json:"data_source"`
	// Description A free-text description.
	Description *string `json:"description,omitempty"`
	// Mappings The user-declared mapping on the dataset definition.
	Mappings *DatasetMapping `json:"mappings,omitempty"`
	// Name The dataset name.
	Name string `json:"name"`
	// Resource The URI that identifies the data to read, resolved against the referenced
	// data source. It can include glob patterns, for example a recursive pattern
	// that matches Parquet files under `s3://logs-bucket/access`.
	Resource string `json:"resource"`
	// Settings Format- and parsing-specific settings that configure how the resource is
	// read. Common keys include `format` and `partition_detection`. Additional keys
	// depend on the format reader; compression can be inferred from the resource
	// URI.
	Settings map[string]json.RawMessage `json:"settings,omitempty"`
}

func (s *ESQLDataset) UnmarshalJSON(data []byte) error {

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

		case "data_source":
			if err := dec.Decode(&s.DataSource); err != nil {
				return fmt.Errorf("%s | %w", "DataSource", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "resource":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Resource", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Resource = o

		case "settings":
			if s.Settings == nil {
				s.Settings = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		}
	}
	return nil
}

// NewESQLDataset returns a ESQLDataset.
func NewESQLDataset() *ESQLDataset {
	r := &ESQLDataset{
		Settings: make(map[string]json.RawMessage),
	}

	return r
}
