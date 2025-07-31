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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/connectorfieldtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/displaytype"
)

// ConnectorConfigProperties type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/_types/Connector.ts#L83-L99
type ConnectorConfigProperties struct {
	Category       *string                                `json:"category,omitempty"`
	DefaultValue   ScalarValue                            `json:"default_value"`
	DependsOn      []Dependency                           `json:"depends_on"`
	Display        displaytype.DisplayType                `json:"display"`
	Label          string                                 `json:"label"`
	Options        []SelectOption                         `json:"options"`
	Order          *int                                   `json:"order,omitempty"`
	Placeholder    *string                                `json:"placeholder,omitempty"`
	Required       bool                                   `json:"required"`
	Sensitive      bool                                   `json:"sensitive"`
	Tooltip        *string                                `json:"tooltip,omitempty"`
	Type           *connectorfieldtype.ConnectorFieldType `json:"type,omitempty"`
	UiRestrictions []string                               `json:"ui_restrictions,omitempty"`
	Validations    []Validation                           `json:"validations,omitempty"`
	Value          json.RawMessage                        `json:"value,omitempty"`
}

func (s *ConnectorConfigProperties) UnmarshalJSON(data []byte) error {

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

		case "category":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Category", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Category = &o

		case "default_value":
			if err := dec.Decode(&s.DefaultValue); err != nil {
				return fmt.Errorf("%s | %w", "DefaultValue", err)
			}

		case "depends_on":
			if err := dec.Decode(&s.DependsOn); err != nil {
				return fmt.Errorf("%s | %w", "DependsOn", err)
			}

		case "display":
			if err := dec.Decode(&s.Display); err != nil {
				return fmt.Errorf("%s | %w", "Display", err)
			}

		case "label":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Label", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Label = o

		case "options":
			if err := dec.Decode(&s.Options); err != nil {
				return fmt.Errorf("%s | %w", "Options", err)
			}

		case "order":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}
				s.Order = &value
			case float64:
				f := int(v)
				s.Order = &f
			}

		case "placeholder":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Placeholder", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Placeholder = &o

		case "required":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Required", err)
				}
				s.Required = value
			case bool:
				s.Required = v
			}

		case "sensitive":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Sensitive", err)
				}
				s.Sensitive = value
			case bool:
				s.Sensitive = v
			}

		case "tooltip":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tooltip", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tooltip = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "ui_restrictions":
			if err := dec.Decode(&s.UiRestrictions); err != nil {
				return fmt.Errorf("%s | %w", "UiRestrictions", err)
			}

		case "validations":

			buf := []json.RawMessage{}
			dec.Decode(&buf)
			for _, rawMsg := range buf {

				source := bytes.NewReader(rawMsg)
				localDec := json.NewDecoder(source)
				kind := make(map[string]string, 0)
				localDec.Decode(&kind)
				source.Seek(0, io.SeekStart)

				switch kind["type"] {

				case "less_than":
					o := NewLessThanValidation()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "less_than", err)
					}
					s.Validations = append(s.Validations, *o)
				case "greater_than":
					o := NewGreaterThanValidation()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "greater_than", err)
					}
					s.Validations = append(s.Validations, *o)
				case "list_type":
					o := NewListTypeValidation()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "list_type", err)
					}
					s.Validations = append(s.Validations, *o)
				case "included_in":
					o := NewIncludedInValidation()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "included_in", err)
					}
					s.Validations = append(s.Validations, *o)
				case "regex":
					o := NewRegexValidation()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "regex", err)
					}
					s.Validations = append(s.Validations, *o)
				default:
					o := new(any)
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("Validations | %w", err)
					}
					s.Validations = append(s.Validations, *o)
				}
			}

		case "value":
			if err := dec.Decode(&s.Value); err != nil {
				return fmt.Errorf("%s | %w", "Value", err)
			}

		}
	}
	return nil
}

// NewConnectorConfigProperties returns a ConnectorConfigProperties.
func NewConnectorConfigProperties() *ConnectorConfigProperties {
	r := &ConnectorConfigProperties{}

	return r
}
