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
// https://github.com/elastic/elasticsearch-specification/tree/7560c979602e6941815872bdaec801200bc7ec4e

package putdataset

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package putdataset
//
// https://github.com/elastic/elasticsearch-specification/blob/7560c979602e6941815872bdaec801200bc7ec4e/specification/esql/put_dataset/PutDatasetRequest.ts#L27-L86
type Request struct {
	// DataSource The name of the referenced data source. The data source must already exist.
	DataSource string `json:"data_source"`
	// Description A free-text description of the dataset.
	Description *string `json:"description,omitempty"`
	// Mappings User-declared mapping on the dataset definition
	Mappings *types.DatasetMapping `json:"mappings,omitempty"`
	// Resource The URI that identifies the data to read, resolved against the referenced
	// data source. It can include glob patterns. For example, a recursive pattern
	// can match all Parquet files under the `s3://logs-bucket/access` prefix.
	Resource string `json:"resource"`
	// Settings Format and parsing-specific settings that configure how the resource is read.
	// Common keys include `format`, which explicitly selects a registered format,
	// and `partition_detection`, which accepts `auto`, `hive`, `template`, or
	// `none`. Additional keys depend on the format reader. Compression can be
	// inferred from the resource URI.
	Settings map[string]json.RawMessage `json:"settings,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Settings: make(map[string]json.RawMessage, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putdataset request: %w", err)
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
