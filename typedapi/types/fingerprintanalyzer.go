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

// FingerprintAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/analyzers.ts#L37-L45
type FingerprintAnalyzer struct {
	MaxOutputSize    int       `json:"max_output_size"`
	PreserveOriginal bool      `json:"preserve_original"`
	Separator        string    `json:"separator"`
	Stopwords        StopWords `json:"stopwords,omitempty"`
	StopwordsPath    *string   `json:"stopwords_path,omitempty"`
	Type             string    `json:"type,omitempty"`
	Version          *string   `json:"version,omitempty"`
}

func (s *FingerprintAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "max_output_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxOutputSize", err)
				}
				s.MaxOutputSize = value
			case float64:
				f := int(v)
				s.MaxOutputSize = f
			}

		case "preserve_original":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PreserveOriginal", err)
				}
				s.PreserveOriginal = value
			case bool:
				s.PreserveOriginal = v
			}

		case "separator":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Separator", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Separator = o

		case "stopwords":
			if err := dec.Decode(&s.Stopwords); err != nil {
				return fmt.Errorf("%s | %w", "Stopwords", err)
			}

		case "stopwords_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "StopwordsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StopwordsPath = &o

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
func (s FingerprintAnalyzer) MarshalJSON() ([]byte, error) {
	type innerFingerprintAnalyzer FingerprintAnalyzer
	tmp := innerFingerprintAnalyzer{
		MaxOutputSize:    s.MaxOutputSize,
		PreserveOriginal: s.PreserveOriginal,
		Separator:        s.Separator,
		Stopwords:        s.Stopwords,
		StopwordsPath:    s.StopwordsPath,
		Type:             s.Type,
		Version:          s.Version,
	}

	tmp.Type = "fingerprint"

	return json.Marshal(tmp)
}

// NewFingerprintAnalyzer returns a FingerprintAnalyzer.
func NewFingerprintAnalyzer() *FingerprintAnalyzer {
	r := &FingerprintAnalyzer{}

	return r
}
