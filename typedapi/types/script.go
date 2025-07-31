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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
)

// Script type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Scripting.ts#L75-L99
type Script struct {
	// Id The `id` for a stored script.
	Id *string `json:"id,omitempty"`
	// Lang Specifies the language the script is written in.
	Lang    *scriptlanguage.ScriptLanguage `json:"lang,omitempty"`
	Options map[string]string              `json:"options,omitempty"`
	// Params Specifies any named parameters that are passed into the script as variables.
	// Use parameters instead of hard-coded values to decrease compile time.
	Params map[string]json.RawMessage `json:"params,omitempty"`
	// Source The script source.
	Source *string `json:"source,omitempty"`
}

func (s *Script) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		if !bytes.HasPrefix(data, []byte(`"`)) {
			data = append([]byte{'"'}, data...)
			data = append(data, []byte{'"'}...)
		}
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Source)
		if err != nil {
			return err
		}
		return nil
	}

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "lang":
			if err := dec.Decode(&s.Lang); err != nil {
				return fmt.Errorf("%s | %w", "Lang", err)
			}

		case "options":
			if s.Options == nil {
				s.Options = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Options); err != nil {
				return fmt.Errorf("%s | %w", "Options", err)
			}

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "source":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Source = &o

		}
	}
	return nil
}

// NewScript returns a Script.
func NewScript() *Script {
	r := &Script{
		Options: make(map[string]string),
		Params:  make(map[string]json.RawMessage),
	}

	return r
}
