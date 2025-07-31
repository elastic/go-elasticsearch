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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/snowballlanguage"
)

// SnowballAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/analyzers.ts#L374-L380
type SnowballAnalyzer struct {
	Language  snowballlanguage.SnowballLanguage `json:"language"`
	Stopwords StopWords                         `json:"stopwords,omitempty"`
	Type      string                            `json:"type,omitempty"`
	Version   *string                           `json:"version,omitempty"`
}

func (s *SnowballAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "language":
			if err := dec.Decode(&s.Language); err != nil {
				return fmt.Errorf("%s | %w", "Language", err)
			}

		case "stopwords":
			if err := dec.Decode(&s.Stopwords); err != nil {
				return fmt.Errorf("%s | %w", "Stopwords", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s SnowballAnalyzer) MarshalJSON() ([]byte, error) {
	type innerSnowballAnalyzer SnowballAnalyzer
	tmp := innerSnowballAnalyzer{
		Language:  s.Language,
		Stopwords: s.Stopwords,
		Type:      s.Type,
		Version:   s.Version,
	}

	tmp.Type = "snowball"

	return json.Marshal(tmp)
}

// NewSnowballAnalyzer returns a SnowballAnalyzer.
func NewSnowballAnalyzer() *SnowballAnalyzer {
	r := &SnowballAnalyzer{}

	return r
}

type SnowballAnalyzerVariant interface {
	SnowballAnalyzerCaster() *SnowballAnalyzer
}

func (s *SnowballAnalyzer) SnowballAnalyzerCaster() *SnowballAnalyzer {
	return s
}

func (s *SnowballAnalyzer) AnalyzerCaster() *Analyzer {
	o := Analyzer(s)
	return &o
}
