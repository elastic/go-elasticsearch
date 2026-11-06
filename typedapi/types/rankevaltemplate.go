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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// RankEvalTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/rank_eval/types.ts#L112-L121
type RankEvalTemplate struct {
	// Id The ID of the template. A rated request refers to it by this ID in its
	// `template_id`.
	Id string `json:"id"`
	// Template The search template, either an inline `source` or the `id` of a stored search
	// template. It must be a `mustache` template (the default `lang`), and it must
	// render to a search request body. Any `params` set here are ignored: each
	// rated request provides its own `params`, which are used to render the
	// template.
	Template Script `json:"template"`
}

func (s *RankEvalTemplate) UnmarshalJSON(data []byte) error {

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

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return fmt.Errorf("%s | %w", "Template", err)
			}

		}
	}
	return nil
}

// NewRankEvalTemplate returns a RankEvalTemplate.
func NewRankEvalTemplate() *RankEvalTemplate {
	r := &RankEvalTemplate{}

	return r
}

type RankEvalTemplateVariant interface {
	RankEvalTemplateCaster() *RankEvalTemplate
}

func (s *RankEvalTemplate) RankEvalTemplateCaster() *RankEvalTemplate {
	return s
}
