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
)

// DataframeAnalyticsSource type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeAnalytics.ts#L39-L53
type DataframeAnalyticsSource struct {
	// Index Index or indices on which to perform the analysis. It can be a single index
	// or index pattern as well as an array of indices or patterns. NOTE: If your
	// source indices contain documents with the same IDs, only the document that is
	// indexed last appears in the destination index.
	Index []string `json:"index"`
	// Query The Elasticsearch query domain-specific language (DSL). This value
	// corresponds to the query object in an Elasticsearch search POST body. All the
	// options that are supported by Elasticsearch can be used, as this object is
	// passed verbatim to Elasticsearch. By default, this property has the following
	// value: {"match_all": {}}.
	Query *Query `json:"query,omitempty"`
	// RuntimeMappings Definitions of runtime fields that will become part of the mapping of the
	// destination index.
	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`
	// Source_ Specify `includes` and/or `excludes patterns to select which fields will be
	// present in the destination. Fields that are excluded cannot be included in
	// the analysis.
	Source_ *DataframeAnalysisAnalyzedFields `json:"_source,omitempty"`
}

func (s *DataframeAnalyticsSource) UnmarshalJSON(data []byte) error {

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

		case "index":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Index = append(s.Index, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Index); err != nil {
					return err
				}
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "runtime_mappings":
			if err := dec.Decode(&s.RuntimeMappings); err != nil {
				return err
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeAnalyticsSource returns a DataframeAnalyticsSource.
func NewDataframeAnalyticsSource() *DataframeAnalyticsSource {
	r := &DataframeAnalyticsSource{}

	return r
}
