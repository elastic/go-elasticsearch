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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DeleteInferenceEndpointResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/inference/_types/Results.ts#L91-L96
type DeleteInferenceEndpointResult struct {
	// Acknowledged For a successful response, this value is always true. On failure, an
	// exception is returned instead.
	Acknowledged bool     `json:"acknowledged"`
	Pipelines    []string `json:"pipelines"`
}

func (s *DeleteInferenceEndpointResult) UnmarshalJSON(data []byte) error {

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

		case "acknowledged":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Acknowledged", err)
				}
				s.Acknowledged = value
			case bool:
				s.Acknowledged = v
			}

		case "pipelines":
			if err := dec.Decode(&s.Pipelines); err != nil {
				return fmt.Errorf("%s | %w", "Pipelines", err)
			}

		}
	}
	return nil
}

// NewDeleteInferenceEndpointResult returns a DeleteInferenceEndpointResult.
func NewDeleteInferenceEndpointResult() *DeleteInferenceEndpointResult {
	r := &DeleteInferenceEndpointResult{}

	return r
}
