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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Sampling configuration as returned by the API.
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/indices/_types/SampleConfiguration.ts#L24-L64
type SamplingConfiguration struct {
	// CreationTime The time when the sampling configuration was created.
	CreationTime DateTime `json:"creation_time,omitempty"`
	// CreationTimeInMillis The time when the sampling configuration was created, in milliseconds since
	// epoch.
	CreationTimeInMillis int64 `json:"creation_time_in_millis"`
	// If An optional condition script that sampled documents must satisfy.
	If *string `json:"if,omitempty"`
	// MaxSamples The maximum number of documents to sample.
	MaxSamples int `json:"max_samples"`
	// MaxSize The maximum total size of sampled documents.
	MaxSize ByteSize `json:"max_size,omitempty"`
	// MaxSizeInBytes The maximum total size of sampled documents in bytes.
	MaxSizeInBytes int64 `json:"max_size_in_bytes"`
	// Rate The fraction of documents to sample between 0 and 1.
	Rate Float64 `json:"rate"`
	// TimeToLive The duration for which the sampled documents should be retained.
	TimeToLive Duration `json:"time_to_live,omitempty"`
	// TimeToLiveInMillis The duration for which the sampled documents should be retained, in
	// milliseconds.
	TimeToLiveInMillis int64 `json:"time_to_live_in_millis"`
}

func (s *SamplingConfiguration) UnmarshalJSON(data []byte) error {

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

		case "creation_time":
			if err := dec.Decode(&s.CreationTime); err != nil {
				return fmt.Errorf("%s | %w", "CreationTime", err)
			}

		case "creation_time_in_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CreationTimeInMillis", err)
				}
				s.CreationTimeInMillis = value
			case float64:
				f := int64(v)
				s.CreationTimeInMillis = f
			}

		case "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "If", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.If = &o

		case "max_samples":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxSamples", err)
				}
				s.MaxSamples = value
			case float64:
				f := int(v)
				s.MaxSamples = f
			}

		case "max_size":
			if err := dec.Decode(&s.MaxSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxSize", err)
			}

		case "max_size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxSizeInBytes", err)
				}
				s.MaxSizeInBytes = value
			case float64:
				f := int64(v)
				s.MaxSizeInBytes = f
			}

		case "rate":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Rate", err)
				}
				f := Float64(value)
				s.Rate = f
			case float64:
				f := Float64(v)
				s.Rate = f
			}

		case "time_to_live":
			if err := dec.Decode(&s.TimeToLive); err != nil {
				return fmt.Errorf("%s | %w", "TimeToLive", err)
			}

		case "time_to_live_in_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TimeToLiveInMillis", err)
				}
				s.TimeToLiveInMillis = value
			case float64:
				f := int64(v)
				s.TimeToLiveInMillis = f
			}

		}
	}
	return nil
}

// NewSamplingConfiguration returns a SamplingConfiguration.
func NewSamplingConfiguration() *SamplingConfiguration {
	r := &SamplingConfiguration{}

	return r
}
