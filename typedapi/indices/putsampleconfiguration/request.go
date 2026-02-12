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

package putsampleconfiguration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package putsampleconfiguration
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/indices/put_sample_configuration/IndicesPutSampleConfigurationRequest.ts#L26-L92
type Request struct {
	// If An optional condition script that sampled documents must satisfy.
	If *string `json:"if,omitempty"`
	// MaxSamples The maximum number of documents to sample. Must be greater than 0 and less
	// than or equal to 10,000.
	MaxSamples *int `json:"max_samples,omitempty"`
	// MaxSize The maximum total size of sampled documents. Must be greater than 0 and less
	// than or equal to 5GB.
	MaxSize types.ByteSize `json:"max_size,omitempty"`
	// Rate The fraction of documents to sample. Must be greater than 0 and less than or
	// equal to 1.
	// Can be specified as a number or a string.
	Rate types.Stringifieddouble `json:"rate"`
	// TimeToLive The duration for which the sampled documents should be retained.
	// Must be greater than 0 and less than or equal to 30 days.
	TimeToLive types.Duration `json:"time_to_live,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putsampleconfiguration request: %w", err)
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
				s.MaxSamples = &value
			case float64:
				f := int(v)
				s.MaxSamples = &f
			}

		case "max_size":
			if err := dec.Decode(&s.MaxSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxSize", err)
			}

		case "rate":
			if err := dec.Decode(&s.Rate); err != nil {
				return fmt.Errorf("%s | %w", "Rate", err)
			}

		case "time_to_live":
			if err := dec.Decode(&s.TimeToLive); err != nil {
				return fmt.Errorf("%s | %w", "TimeToLive", err)
			}

		}
	}
	return nil
}
