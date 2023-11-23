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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sourcefieldmode"
)

// SourceField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/mapping/meta-fields.ts#L58-L65
type SourceField struct {
	Compress          *bool                            `json:"compress,omitempty"`
	CompressThreshold *string                          `json:"compress_threshold,omitempty"`
	Enabled           *bool                            `json:"enabled,omitempty"`
	Excludes          []string                         `json:"excludes,omitempty"`
	Includes          []string                         `json:"includes,omitempty"`
	Mode              *sourcefieldmode.SourceFieldMode `json:"mode,omitempty"`
}

func (s *SourceField) UnmarshalJSON(data []byte) error {

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

		case "compress":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Compress = &value
			case bool:
				s.Compress = &v
			}

		case "compress_threshold":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CompressThreshold = &o

		case "enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		case "excludes":
			if err := dec.Decode(&s.Excludes); err != nil {
				return err
			}

		case "includes":
			if err := dec.Decode(&s.Includes); err != nil {
				return err
			}

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSourceField returns a SourceField.
func NewSourceField() *SourceField {
	r := &SourceField{}

	return r
}
