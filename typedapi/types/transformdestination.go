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
// https://github.com/elastic/elasticsearch-specification/tree/964a36594f01c23463551aa7d09d17e514f361d1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TransformDestination type.
//
// https://github.com/elastic/elasticsearch-specification/blob/964a36594f01c23463551aa7d09d17e514f361d1/specification/transform/_types/Transform.ts#L40-L62
type TransformDestination struct {
	// Aliases The aliases that the destination index for the transform should have. Aliases
	// are manipulated using the stored credentials of the transform, which means
	// the secondary credentials supplied at creation time (if both primary and
	// secondary credentials are specified).
	//
	// The destination index is added to the aliases regardless of whether the
	// destination index was created by the transform or pre-created by the user.
	Aliases []DestinationAlias `json:"aliases,omitempty"`
	// Index The destination index for the transform. The mappings of the destination
	// index are deduced based on the source fields when possible. If alternate
	// mappings are required, use the create index API prior to starting the
	// transform.
	Index *string `json:"index,omitempty"`
	// Pipeline The unique identifier for an ingest pipeline.
	Pipeline *string `json:"pipeline,omitempty"`
}

func (s *TransformDestination) UnmarshalJSON(data []byte) error {

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

		case "aliases":
			if err := dec.Decode(&s.Aliases); err != nil {
				return fmt.Errorf("%s | %w", "Aliases", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "pipeline":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Pipeline", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pipeline = &o

		}
	}
	return nil
}

// NewTransformDestination returns a TransformDestination.
func NewTransformDestination() *TransformDestination {
	r := &TransformDestination{}

	return r
}

type TransformDestinationVariant interface {
	TransformDestinationCaster() *TransformDestination
}

func (s *TransformDestination) TransformDestinationCaster() *TransformDestination {
	return s
}
