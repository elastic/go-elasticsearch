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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/reasoningeffort"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/reasoningsummary"
)

// The reasoning configuration to use for the completion request. Currently
// supported only for `elastic` provider.
//
// https://github.com/elastic/elasticsearch-specification/blob/836fca874204ca4173ae5c36fb6b5107d28d2fc0/specification/inference/_types/CommonTypes.ts#L216-L246
type Reasoning struct {
	// Effort The level of effort the model should put into reasoning. This is a hint that
	// guides the model in how much effort to put into reasoning, with `xhigh` being
	// the most effort and `none` being no effort.
	Effort *reasoningeffort.ReasoningEffort `json:"effort,omitempty"`
	// Enabled Whether to enable reasoning with default settings. This is a shortcut for
	// enabling reasoning without having to specify the other parameters. If
	// `enabled` is set to `true`, then reasoning at the `medium` effort level is
	// enabled. Ignored if `effort` is specified, in which case that parameter will
	// control the reasoning process instead.
	Enabled *bool `json:"enabled,omitempty"`
	// Exclude Whether to exclude reasoning information from the response. If `true`, the
	// response will not include any reasoning details.
	Exclude *bool `json:"exclude,omitempty"`
	// Summary The level of detail included in the reasoning summary returned in the
	// response. This is a hint on how much detail to include in the summary of the
	// reasoning that is returned in the response, with `auto` being the default
	// level of detail, `concise` being less detail, and `detailed` being more
	// detail.
	Summary *reasoningsummary.ReasoningSummary `json:"summary,omitempty"`
}

func (s *Reasoning) UnmarshalJSON(data []byte) error {

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

		case "effort":
			if err := dec.Decode(&s.Effort); err != nil {
				return fmt.Errorf("%s | %w", "Effort", err)
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		case "exclude":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Exclude", err)
				}
				s.Exclude = &value
			case bool:
				s.Exclude = &v
			}

		case "summary":
			if err := dec.Decode(&s.Summary); err != nil {
				return fmt.Errorf("%s | %w", "Summary", err)
			}

		}
	}
	return nil
}

// NewReasoning returns a Reasoning.
func NewReasoning() *Reasoning {
	r := &Reasoning{}

	return r
}

type ReasoningVariant interface {
	ReasoningCaster() *Reasoning
}

func (s *Reasoning) ReasoningCaster() *Reasoning {
	return s
}
