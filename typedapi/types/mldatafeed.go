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

// MLDatafeed type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Datafeed.ts#L37-L58
type MLDatafeed struct {
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`
	// Authorization The security privileges that the datafeed uses to run its queries. If Elastic
	// Stack security features were disabled at the time of the most recent update
	// to the datafeed, this property is omitted.
	Authorization          *DatafeedAuthorization `json:"authorization,omitempty"`
	ChunkingConfig         *ChunkingConfig        `json:"chunking_config,omitempty"`
	DatafeedId             string                 `json:"datafeed_id"`
	DelayedDataCheckConfig DelayedDataCheckConfig `json:"delayed_data_check_config"`
	Frequency              Duration               `json:"frequency,omitempty"`
	Indexes                []string               `json:"indexes,omitempty"`
	Indices                []string               `json:"indices"`
	IndicesOptions         *IndicesOptions        `json:"indices_options,omitempty"`
	JobId                  string                 `json:"job_id"`
	MaxEmptySearches       *int                   `json:"max_empty_searches,omitempty"`
	Query                  Query                  `json:"query"`
	QueryDelay             Duration               `json:"query_delay,omitempty"`
	RuntimeMappings        RuntimeFields          `json:"runtime_mappings,omitempty"`
	ScriptFields           map[string]ScriptField `json:"script_fields,omitempty"`
	ScrollSize             *int                   `json:"scroll_size,omitempty"`
}

func (s *MLDatafeed) UnmarshalJSON(data []byte) error {

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

		case "aggregations", "aggs":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]Aggregations, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return err
			}

		case "authorization":
			if err := dec.Decode(&s.Authorization); err != nil {
				return err
			}

		case "chunking_config":
			if err := dec.Decode(&s.ChunkingConfig); err != nil {
				return err
			}

		case "datafeed_id":
			if err := dec.Decode(&s.DatafeedId); err != nil {
				return err
			}

		case "delayed_data_check_config":
			if err := dec.Decode(&s.DelayedDataCheckConfig); err != nil {
				return err
			}

		case "frequency":
			if err := dec.Decode(&s.Frequency); err != nil {
				return err
			}

		case "indexes":
			if err := dec.Decode(&s.Indexes); err != nil {
				return err
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "indices_options":
			if err := dec.Decode(&s.IndicesOptions); err != nil {
				return err
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "max_empty_searches":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxEmptySearches = &value
			case float64:
				f := int(v)
				s.MaxEmptySearches = &f
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "query_delay":
			if err := dec.Decode(&s.QueryDelay); err != nil {
				return err
			}

		case "runtime_mappings":
			if err := dec.Decode(&s.RuntimeMappings); err != nil {
				return err
			}

		case "script_fields":
			if s.ScriptFields == nil {
				s.ScriptFields = make(map[string]ScriptField, 0)
			}
			if err := dec.Decode(&s.ScriptFields); err != nil {
				return err
			}

		case "scroll_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ScrollSize = &value
			case float64:
				f := int(v)
				s.ScrollSize = &f
			}

		}
	}
	return nil
}

// NewMLDatafeed returns a MLDatafeed.
func NewMLDatafeed() *MLDatafeed {
	r := &MLDatafeed{
		Aggregations: make(map[string]Aggregations, 0),
		ScriptFields: make(map[string]ScriptField, 0),
	}

	return r
}
