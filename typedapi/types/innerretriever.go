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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scorenormalizer"
)

// InnerRetriever type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Retriever.ts#L85-L89
type InnerRetriever struct {
	Normalizer scorenormalizer.ScoreNormalizer `json:"normalizer"`
	Retriever  RetrieverContainer              `json:"retriever"`
	Weight     float32                         `json:"weight"`
}

func (s *InnerRetriever) UnmarshalJSON(data []byte) error {

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

		case "normalizer":
			if err := dec.Decode(&s.Normalizer); err != nil {
				return fmt.Errorf("%s | %w", "Normalizer", err)
			}

		case "retriever":
			if err := dec.Decode(&s.Retriever); err != nil {
				return fmt.Errorf("%s | %w", "Retriever", err)
			}

		case "weight":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Weight", err)
				}
				f := float32(value)
				s.Weight = f
			case float64:
				f := float32(v)
				s.Weight = f
			}

		}
	}
	return nil
}

// NewInnerRetriever returns a InnerRetriever.
func NewInnerRetriever() *InnerRetriever {
	r := &InnerRetriever{}

	return r
}

type InnerRetrieverVariant interface {
	InnerRetrieverCaster() *InnerRetriever
}

func (s *InnerRetriever) InnerRetrieverCaster() *InnerRetriever {
	return s
}
