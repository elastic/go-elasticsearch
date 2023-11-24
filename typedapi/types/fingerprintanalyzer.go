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

// FingerprintAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/analysis/analyzers.ts#L37-L45
type FingerprintAnalyzer struct {
	MaxOutputSize    int      `json:"max_output_size"`
	PreserveOriginal bool     `json:"preserve_original"`
	Separator        string   `json:"separator"`
	Stopwords        []string `json:"stopwords,omitempty"`
	StopwordsPath    *string  `json:"stopwords_path,omitempty"`
	Type             string   `json:"type,omitempty"`
	Version          *string  `json:"version,omitempty"`
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

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxOutputSize = value
			case float64:
				f := int(v)
				s.MaxOutputSize = f
			}

		case "preserve_original":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.PreserveOriginal = value
			case bool:
				s.PreserveOriginal = v
			}

		case "separator":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Separator = o

		case "stopwords":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Stopwords = append(s.Stopwords, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Stopwords); err != nil {
					return err
				}
			}

		case "stopwords_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StopwordsPath = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
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
