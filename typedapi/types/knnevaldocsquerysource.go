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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Samples the query vectors from the stored documents.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/knn_eval/types.ts#L33-L46
type KnnEvalDocsQuerySource struct {
	From string `json:"from,omitempty"`
	// Seed The seed of the sampling, to sample the same documents across requests.
	Seed *int `json:"seed,omitempty"`
	// Size The number of documents to sample, from 1 to 1000.
	Size int `json:"size"`
}

func (s *KnnEvalDocsQuerySource) UnmarshalJSON(data []byte) error {

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

		case "from":
			if err := dec.Decode(&s.From); err != nil {
				return fmt.Errorf("%s | %w", "From", err)
			}

		case "seed":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Seed", err)
				}
				s.Seed = &value
			case float64:
				f := int(v)
				s.Seed = &f
			}

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = value
			case float64:
				f := int(v)
				s.Size = f
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s KnnEvalDocsQuerySource) MarshalJSON() ([]byte, error) {
	type innerKnnEvalDocsQuerySource KnnEvalDocsQuerySource
	tmp := innerKnnEvalDocsQuerySource{
		From: s.From,
		Seed: s.Seed,
		Size: s.Size,
	}

	tmp.From = "docs"

	return json.Marshal(tmp)
}

// NewKnnEvalDocsQuerySource returns a KnnEvalDocsQuerySource.
func NewKnnEvalDocsQuerySource() *KnnEvalDocsQuerySource {
	r := &KnnEvalDocsQuerySource{}

	return r
}

type KnnEvalDocsQuerySourceVariant interface {
	KnnEvalDocsQuerySourceCaster() *KnnEvalDocsQuerySource
}

func (s *KnnEvalDocsQuerySource) KnnEvalDocsQuerySourceCaster() *KnnEvalDocsQuerySource {
	return s
}

func (s *KnnEvalDocsQuerySource) KnnEvalQuerySourceCaster() *KnnEvalQuerySource {
	if s == nil {
		return nil
	}
	o := KnnEvalQuerySource(s)
	return &o
}
