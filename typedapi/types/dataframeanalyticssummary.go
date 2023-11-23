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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DataframeAnalyticsSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/DataframeAnalytics.ts#L306-L322
type DataframeAnalyticsSummary struct {
	AllowLazyStart *bool                            `json:"allow_lazy_start,omitempty"`
	Analysis       DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	// Authorization The security privileges that the job uses to run its queries. If Elastic
	// Stack security features were disabled at the time of the most recent update
	// to the job, this property is omitted.
	Authorization    *DataframeAnalyticsAuthorization `json:"authorization,omitempty"`
	CreateTime       *int64                           `json:"create_time,omitempty"`
	Description      *string                          `json:"description,omitempty"`
	Dest             DataframeAnalyticsDestination    `json:"dest"`
	Id               string                           `json:"id"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
	Version          *string                          `json:"version,omitempty"`
}

func (s *DataframeAnalyticsSummary) UnmarshalJSON(data []byte) error {

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

		case "allow_lazy_start":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowLazyStart = &value
			case bool:
				s.AllowLazyStart = &v
			}

		case "analysis":
			if err := dec.Decode(&s.Analysis); err != nil {
				return err
			}

		case "analyzed_fields":
			if err := dec.Decode(&s.AnalyzedFields); err != nil {
				return err
			}

		case "authorization":
			if err := dec.Decode(&s.Authorization); err != nil {
				return err
			}

		case "create_time":
			if err := dec.Decode(&s.CreateTime); err != nil {
				return err
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "dest":
			if err := dec.Decode(&s.Dest); err != nil {
				return err
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "max_num_threads":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxNumThreads = &value
			case float64:
				f := int(v)
				s.MaxNumThreads = &f
			}

		case "model_memory_limit":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelMemoryLimit = &o

		case "source":
			if err := dec.Decode(&s.Source); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeAnalyticsSummary returns a DataframeAnalyticsSummary.
func NewDataframeAnalyticsSummary() *DataframeAnalyticsSummary {
	r := &DataframeAnalyticsSummary{}

	return r
}
