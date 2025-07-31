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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/phoneticencoder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/phoneticlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/phoneticnametype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/phoneticruletype"
)

// PhoneticTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/phonetic-plugin.ts#L64-L72
type PhoneticTokenFilter struct {
	Encoder     phoneticencoder.PhoneticEncoder     `json:"encoder"`
	Languageset []phoneticlanguage.PhoneticLanguage `json:"languageset,omitempty"`
	MaxCodeLen  *int                                `json:"max_code_len,omitempty"`
	NameType    *phoneticnametype.PhoneticNameType  `json:"name_type,omitempty"`
	Replace     *bool                               `json:"replace,omitempty"`
	RuleType    *phoneticruletype.PhoneticRuleType  `json:"rule_type,omitempty"`
	Type        string                              `json:"type,omitempty"`
	Version     *string                             `json:"version,omitempty"`
}

func (s *PhoneticTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "encoder":
			if err := dec.Decode(&s.Encoder); err != nil {
				return fmt.Errorf("%s | %w", "Encoder", err)
			}

		case "languageset":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := &phoneticlanguage.PhoneticLanguage{}
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Languageset", err)
				}

				s.Languageset = append(s.Languageset, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Languageset); err != nil {
					return fmt.Errorf("%s | %w", "Languageset", err)
				}
			}

		case "max_code_len":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxCodeLen", err)
				}
				s.MaxCodeLen = &value
			case float64:
				f := int(v)
				s.MaxCodeLen = &f
			}

		case "name_type":
			if err := dec.Decode(&s.NameType); err != nil {
				return fmt.Errorf("%s | %w", "NameType", err)
			}

		case "replace":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Replace", err)
				}
				s.Replace = &value
			case bool:
				s.Replace = &v
			}

		case "rule_type":
			if err := dec.Decode(&s.RuleType); err != nil {
				return fmt.Errorf("%s | %w", "RuleType", err)
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
func (s PhoneticTokenFilter) MarshalJSON() ([]byte, error) {
	type innerPhoneticTokenFilter PhoneticTokenFilter
	tmp := innerPhoneticTokenFilter{
		Encoder:     s.Encoder,
		Languageset: s.Languageset,
		MaxCodeLen:  s.MaxCodeLen,
		NameType:    s.NameType,
		Replace:     s.Replace,
		RuleType:    s.RuleType,
		Type:        s.Type,
		Version:     s.Version,
	}

	tmp.Type = "phonetic"

	return json.Marshal(tmp)
}

// NewPhoneticTokenFilter returns a PhoneticTokenFilter.
func NewPhoneticTokenFilter() *PhoneticTokenFilter {
	r := &PhoneticTokenFilter{}

	return r
}
