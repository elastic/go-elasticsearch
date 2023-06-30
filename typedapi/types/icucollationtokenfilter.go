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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationstrength"
)

// IcuCollationTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/analysis/icu-plugin.ts#L51-L65
type IcuCollationTokenFilter struct {
	Alternate              *icucollationalternate.IcuCollationAlternate         `json:"alternate,omitempty"`
	CaseFirst              *icucollationcasefirst.IcuCollationCaseFirst         `json:"caseFirst,omitempty"`
	CaseLevel              *bool                                                `json:"caseLevel,omitempty"`
	Country                *string                                              `json:"country,omitempty"`
	Decomposition          *icucollationdecomposition.IcuCollationDecomposition `json:"decomposition,omitempty"`
	HiraganaQuaternaryMode *bool                                                `json:"hiraganaQuaternaryMode,omitempty"`
	Language               *string                                              `json:"language,omitempty"`
	Numeric                *bool                                                `json:"numeric,omitempty"`
	Rules                  *string                                              `json:"rules,omitempty"`
	Strength               *icucollationstrength.IcuCollationStrength           `json:"strength,omitempty"`
	Type                   string                                               `json:"type,omitempty"`
	VariableTop            *string                                              `json:"variableTop,omitempty"`
	Variant                *string                                              `json:"variant,omitempty"`
	Version                *string                                              `json:"version,omitempty"`
}

func (s *IcuCollationTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "alternate":
			if err := dec.Decode(&s.Alternate); err != nil {
				return err
			}

		case "caseFirst":
			if err := dec.Decode(&s.CaseFirst); err != nil {
				return err
			}

		case "caseLevel":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.CaseLevel = &value
			case bool:
				s.CaseLevel = &v
			}

		case "country":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Country = &o

		case "decomposition":
			if err := dec.Decode(&s.Decomposition); err != nil {
				return err
			}

		case "hiraganaQuaternaryMode":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.HiraganaQuaternaryMode = &value
			case bool:
				s.HiraganaQuaternaryMode = &v
			}

		case "language":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Language = &o

		case "numeric":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Numeric = &value
			case bool:
				s.Numeric = &v
			}

		case "rules":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Rules = &o

		case "strength":
			if err := dec.Decode(&s.Strength); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "variableTop":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VariableTop = &o

		case "variant":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Variant = &o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIcuCollationTokenFilter returns a IcuCollationTokenFilter.
func NewIcuCollationTokenFilter() *IcuCollationTokenFilter {
	r := &IcuCollationTokenFilter{}

	r.Type = "icu_collation"

	return r
}
