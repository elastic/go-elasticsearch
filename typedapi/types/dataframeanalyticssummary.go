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

// DataframeAnalyticsSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/DataframeAnalytics.ts#L306-L323
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
	Meta_            Metadata                         `json:"_meta,omitempty"`
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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowLazyStart", err)
				}
				s.AllowLazyStart = &value
			case bool:
				s.AllowLazyStart = &v
			}

		case "analysis":
			if err := dec.Decode(&s.Analysis); err != nil {
				return fmt.Errorf("%s | %w", "Analysis", err)
			}

		case "analyzed_fields":
			if err := dec.Decode(&s.AnalyzedFields); err != nil {
				return fmt.Errorf("%s | %w", "AnalyzedFields", err)
			}

		case "authorization":
			if err := dec.Decode(&s.Authorization); err != nil {
				return fmt.Errorf("%s | %w", "Authorization", err)
			}

		case "create_time":
			if err := dec.Decode(&s.CreateTime); err != nil {
				return fmt.Errorf("%s | %w", "CreateTime", err)
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

		case "dest":
			if err := dec.Decode(&s.Dest); err != nil {
				return fmt.Errorf("%s | %w", "Dest", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "max_num_threads":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxNumThreads", err)
				}
				s.MaxNumThreads = &value
			case float64:
				f := int(v)
				s.MaxNumThreads = &f
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "model_memory_limit":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ModelMemoryLimit", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelMemoryLimit = &o

		case "source":
			if err := dec.Decode(&s.Source); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
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
