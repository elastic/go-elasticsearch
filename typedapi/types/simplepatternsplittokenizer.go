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
// https://github.com/elastic/elasticsearch-specification/tree/0f6f3696eb685db8b944feefb6a209ad7e385b9c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SimplePatternSplitTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0f6f3696eb685db8b944feefb6a209ad7e385b9c/specification/_types/analysis/tokenizers.ts#L116-L119
type SimplePatternSplitTokenizer struct {
	Pattern *string `json:"pattern,omitempty"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *SimplePatternSplitTokenizer) UnmarshalJSON(data []byte) error {

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

		case "pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Pattern", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pattern = &o

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
func (s SimplePatternSplitTokenizer) MarshalJSON() ([]byte, error) {
	type innerSimplePatternSplitTokenizer SimplePatternSplitTokenizer
	tmp := innerSimplePatternSplitTokenizer{
		Pattern: s.Pattern,
		Type:    s.Type,
		Version: s.Version,
	}

	tmp.Type = "simple_pattern_split"

	return json.Marshal(tmp)
}

// NewSimplePatternSplitTokenizer returns a SimplePatternSplitTokenizer.
func NewSimplePatternSplitTokenizer() *SimplePatternSplitTokenizer {
	r := &SimplePatternSplitTokenizer{}

	return r
}

// true

type SimplePatternSplitTokenizerVariant interface {
	SimplePatternSplitTokenizerCaster() *SimplePatternSplitTokenizer
}

func (s *SimplePatternSplitTokenizer) SimplePatternSplitTokenizerCaster() *SimplePatternSplitTokenizer {
	return s
}
